from uuid import UUID
from pydantic import BaseModel


class Lesson(BaseModel):
    id: UUID
    desc: str
    is_approved: bool

    class ConfigDict:
        from_attributes = True
