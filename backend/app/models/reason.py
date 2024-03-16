from enum import Enum


class Reason(str, Enum):
    LESSON = "Занятие"
    HOLIDAY = "Выходной"

    class ConfigDict:
        from_attributes = True
