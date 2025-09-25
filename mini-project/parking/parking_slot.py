from .vehicle import VehicleType


class ParkingSlot:
    def __init__(self, vehicle_type: VehicleType, slot: int) -> None:
        self.vehicle_type: VehicleType = vehicle_type
        self.slot: int = slot

    def increase_slot(self, n: int | None = None):
        self.slot = self.slot + n if n else self.slot + 1

    def decrease_slot(self, n: int | None = None):
        self.slot = self.slot - n if n else self.slot - 1

    def is_available(self):
        return self.slot > 0
