import datetime
from typing import Optional
from uuid import UUID

from pydantic import BaseModel

from .reason import Reason

from .lesson import Lesson


class Occupation(BaseModel):
    id: UUID
    date: datetime.date
    start: datetime.time
    end: datetime.time
    reason: Reason
    created_at: datetime.datetime
    lesson: Optional[Lesson]

    class ConfigDict:
        from_attributes = True
