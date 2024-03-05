from pydantic import BaseModel


class FAQ(BaseModel):
    id: int
    # TODO remove question mark if exists at the end
    question: str
    answer: str