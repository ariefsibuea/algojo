from enum import Enum


class VehicleType(Enum):
    MOTORCYCLE = 1
    CAR = 2
    TRUCK = 3
    BUS = 4


class Vehicle:
    def __init__(self, type_: VehicleType, license_no: str) -> None:
        self.type_: VehicleType = type_
        self.license_no: str = license_no
