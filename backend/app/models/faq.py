from pydantic import BaseModel


class FAQ(BaseModel):
    id: int = None
    # without question mark at the end
    question: str
    answer: str

    class ConfigDict:
        from_attributes = True
