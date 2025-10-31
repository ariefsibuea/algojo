from datetime import date


class Report:
    def __init__(self, parking_date: date, total_income: int) -> None:
        self.parking_date: date = parking_date
        self.total_income: int = total_income
