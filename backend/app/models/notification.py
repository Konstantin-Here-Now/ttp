from pydantic import BaseModel

from .notification_type import NotificationType


class Notification(BaseModel):
    id: int
    type: NotificationType
    link: str