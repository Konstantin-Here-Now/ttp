import datetime
from pydantic import BaseModel


class TimeRange(BaseModel):
    start: datetime.time
    end: datetime.time

    class ConfigDict:
        from_attributes = True
