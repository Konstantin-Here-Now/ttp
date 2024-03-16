from uuid import UUID

from pydantic import BaseModel


class Material(BaseModel):
    id: UUID
    name: str
    desc: str
    link: str

    class ConfigDict:
        from_attributes = True
