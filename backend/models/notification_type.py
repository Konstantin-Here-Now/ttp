from enum import Enum

from pydantic import BaseModel


class NotificationType(Enum, BaseModel):
    EMAIL = 0
    TELEGRAM = 1