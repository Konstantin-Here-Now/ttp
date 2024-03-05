from pydantic import BaseModel
from .day import Day


class Timetable(BaseModel):
    mon: Day
    tue: Day
    wed: Day
    thu: Day
    fri: Day
    sat: Day
    sun: Day