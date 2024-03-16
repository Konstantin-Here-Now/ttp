from pydantic import BaseModel
from .timerange import TimeRange


class AvailableTime(BaseModel):
    at: TimeRange

    class ConfigDict:
        from_attributes = True
