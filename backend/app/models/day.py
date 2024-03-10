import datetime
from pydantic import BaseModel


class Day(BaseModel):
    date: datetime.date
    # string of AvailableTime
    at: str