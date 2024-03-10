from fastapi import APIRouter

from ..models.lesson import Lesson

router = APIRouter(
    prefix="/lessons",
    tags=["lessons"]
)

@router.get("/")
async def read_lessons() -> list[Lesson]:
    return list[Lesson]