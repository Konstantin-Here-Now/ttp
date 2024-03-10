from ..models.timetable import Timetable
from fastapi import APIRouter

router = APIRouter(
    prefix="/timetable",
    tags=["timetable"]
)


@router.get("/")
async def read_timetable() -> Timetable:
    return Timetable
