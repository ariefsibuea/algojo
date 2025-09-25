from .vehicle import VehicleType


class ParkingFeeRule:
    def __init__(self, vehicle_type: VehicleType, fee: int) -> None:
        self.vehicle_type: VehicleType = vehicle_type
        self.fee: int = fee

    def calculate_fee(self) -> int:
        return self.fee
