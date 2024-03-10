from typing import Optional
from uuid import UUID

from pydantic import BaseModel

from .notification import Notification

from .occupation import Occupation


class User(BaseModel):
    id: UUID
    username: str
    password: str
    lastname: str
    firstname: str
    is_admin: bool
    notifications: Optional[list[Notification]]
    occupations: Optional[list[Occupation]]
