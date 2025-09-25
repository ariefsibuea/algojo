from enum import Enum


class TicketPaymentStatus(Enum):
    UNPAID = 1
    PAID = 2
    FAILED = 3


class TicketPayment:
    def __init__(self, ticket_no: str, fee: int) -> None:
        self.payment_no: str = f"PAY-{ticket_no}"
        self.ticket_no: str = ticket_no
        self.fee: int = fee
        self.fine: int = 0
        self.total: int = fee
        self.status: TicketPaymentStatus = TicketPaymentStatus.UNPAID

    def update_status_paid(self):
        self.status = TicketPaymentStatus.PAID
