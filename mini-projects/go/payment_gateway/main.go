package main

import (
	"context"
	"fmt"
	"sync"
)

// Implement an in-memory payment gateway that supports creating payments, capturing payments, and retrieving payment details.
//
// Assume that the PaymentGateway methods are called from a HTTP handler:
// network call --> HTTP handler --> PaymentGateway code --> ReportingService --> network call
//
// The CreatePayment operation involves authorizing the payment and reporting it via an external service. Assume that the reporting service makes an HTTP call inside it.
//
// Payment can not be Captured if it was not reported yet.
//
// Your implementation should handle all scenarios that can occur during payment processing, including successful and unsuccessful operations, in a production-ready way.
func main() {
	g := PaymentGateway{}
	// after implementing the code below, call all methods from the main function, to see that it compiles and runs successfully
}

// / Here is some code template that you should fix + write implementation by using those existing pieces of the code ///
// type PaymentStatus type // enum, should have values: Authorized, Canceled, Captured
type PaymentStatus int

const (
	PaymentStatusAutorized = iota
	PaymentStatusCanceled
	PaymentStatusCaptured
)

type Payment struct {
	ID                int
	AccountBankNumber string
	BankCode          string
	Amount            int
	PaymentType       string
	PaymentStatus     PaymentStatus
	Reported          bool
}

type PaymentGateway struct {
	BankService      BankService
	ReportingService ReportingService
	PaymentRepo      PaymentRepository
}

type CreatePaymentRequest struct {
	AccountBankNumber string
	BankCode          string
	Amount            int
	PaymentType       string
}

type CreatePaymentResponse struct {
	ID            int
	PaymentStatus PaymentStatus
}

func (s PaymentGateway) CreatePayment(ctx context.Context, req CreatePaymentRequest) (CreatePaymentResponse, error) {
	var result CreatePaymentResponse

	// validate the bank
	if err := s.BankService.ValidateBankAccount(ctx, req.AccountBankNumber); err != nil {
		return result, err
	}

	// request to create the payment to bank
	var count = 0
	var errCreate error

	for count < 3 {
		errCreate = s.BankService.CreatePayment(ctx)
		if errCreate == nil {
			break
		}
	}

	if errCreate != nil {
		return result, errCreate
	}

	// store the payment data to repository
	res, err := s.PaymentRepo.Create(ctx, req)
	if err != nil {
		return result, err
	}

	// create reporting
	err = s.ReportingService.ReportAuthorizedPayment(ctx, ReportingRequest{
		PaymentID: res.ID,
		Amount:    req.Amount,
	})
	if err != nil {
		return result, err
	}

	err = s.PaymentRepo.Update(ctx, UpdatePayment{
		PaymentID:     res.ID,
		PaymentStatus: res.PaymentStatus,
		Reported:      true,
	})
	if err != nil {
		return result, err
	}

	result.ID = res.ID
	result.PaymentStatus = res.PaymentStatus

	return result, nil
}

type GetPaymentRequest struct {
	PaymentID int
}

type GetPaymentResponse struct {
	PaymentID     int
	Amount        int
	PaymentType   string
	PaymentStatus PaymentStatus
	Reported      bool
}

func (s PaymentGateway) GetPayment(ctx context.Context, req GetPaymentRequest) (GetPaymentResponse, error) {
	var result GetPaymentResponse

	paymentData, err := s.PaymentRepo.Get(ctx, req.PaymentID)
	if err != nil {
		return result, err
	}

	return GetPaymentResponse{
		PaymentID:     paymentData.ID,
		Amount:        paymentData.Amount,
		PaymentType:   paymentData.PaymentType,
		PaymentStatus: paymentData.PaymentStatus,
		Reported:      paymentData.Reported,
	}, nil
}

func (s PaymentGateway) CancelPayment(ctx context.Context, paymentID int) error {
	paymentData, err := s.GetPayment(ctx, GetPaymentRequest{PaymentID: paymentID})
	if err != nil {
		return err
	}

	if paymentData.PaymentStatus == PaymentStatusCaptured {
		return fmt.Errorf("cancel payment not allowed")
	}

	return s.PaymentRepo.Update(ctx, UpdatePayment{
		PaymentID:     paymentID,
		PaymentStatus: PaymentStatusCanceled, // update here
		Reported:      paymentData.Reported,
	})
}

type CapturePaymentRequest struct {
	PaymentID         int
	AccountBankNumber string
	Amount            int
}

func (s PaymentGateway) CapturePayment(ctx context.Context, req CapturePaymentRequest) error {
	paymentData, err := s.GetPayment(ctx, GetPaymentRequest{PaymentID: req.PaymentID})
	if err != nil {
		return err
	}

	if !paymentData.Reported && paymentData.PaymentStatus != PaymentStatusAutorized {
		return fmt.Errorf("payment now allowed to capture")
	}

	err = s.BankService.CapturePayment(ctx, req.AccountBankNumber, req.Amount)
	if err != nil {
		return err
	}

	err = s.PaymentRepo.Update(ctx, UpdatePayment{
		PaymentID:     paymentData.PaymentID,
		PaymentStatus: PaymentStatusCaptured, // update this part
		Reported:      paymentData.Reported,
	})
	return err
}

type ValidatePaymentRequest struct {
	PaymentID         int
	AccountBankNumber string
	Amount            int
	ValidateType      string // capture, canceled
}

func (s PaymentGateway) ValidatePayment(ctx context.Context, req ValidatePaymentRequest) {

}

type PaymentRepository struct {
	LatestID int
	// Data     []Payment
	Data map[int]Payment
}

func NewPaymentRepository() PaymentRepository {
	return PaymentRepository{
		LatestID: 0,
		Data:     map[int]Payment{},
		Mx:       &sync.Mutex{},
	}
}

func (r *PaymentRepository) Create(ctx context.Context, req CreatePaymentRequest) (CreatePaymentResponse, error) {
	newPayment := Payment{
		ID:                r.LatestID + 1,
		AccountBankNumber: req.AccountBankNumber,
		BankCode:          req.BankCode,
		Amount:            req.Amount,
		PaymentType:       req.PaymentType,
		PaymentStatus:     PaymentStatusAutorized,
	}

	r.Data[newPayment.ID] = newPayment
	r.LatestID = newPayment.ID

	return CreatePaymentResponse{
		ID:            newPayment.ID,
		PaymentStatus: newPayment.PaymentStatus,
	}, nil
}

func (r *PaymentRepository) Get(ctx context.Context, paymentID int) (Payment, error) {
	res, ok := r.Data[paymentID]
	if !ok {
		return Payment{}, fmt.Errorf("data not found")
	}
	return res, nil
}

type UpdatePayment struct {
	PaymentID     int
	PaymentStatus PaymentStatus
	Reported      bool
}

func (r *PaymentRepository) Update(ctx context.Context, req UpdatePayment) error {
	data, ok := r.Data[req.PaymentID]
	if !ok {
		return fmt.Errorf("data not found")
	}

	data.PaymentStatus = req.PaymentStatus
	data.Reported = req.Reported

	r.Data[req.PaymentID] = data

	return nil
}

// func (r *PaymentRepository) UpdateStatus(ctx context.Context, paymentID int, status PaymentStatus) error {
// 	data, ok := r.Data[paymentID]
// 	if !ok {
// 		return fmt.Errorf("data not found")
// 	}

// 	data.PaymentStatus = status
// 	r.Data[paymentID] = data

// 	return nil
// }

type BankService struct{}

func (s *BankService) ValidateBankAccount(ctx context.Context, accountNumber string) error {
	return nil
}

func (s *BankService) CreatePayment(ctx context.Context) error {
	return nil
}

func (s *BankService) CapturePayment(ctx context.Context, accountBankNumber string, amount int) error {
	return nil
}

// / ----------- Please do not change any code below this line except types ----------- ///
type ReportingRequest struct {
	PaymentID int
	Amount    int
}

type ReportingService struct{}

func (s ReportingService) ReportAuthorizedPayment(ctx context.Context, request ReportingRequest) error {
	// Consider that it is implemented already but the code of it is hidden
	// Returning nil here is just a dummy line of code to let it compile, in reality anything can happen, similar to
	// real-life scenarios of calling a 3rd party HTTP service
	return nil
}
