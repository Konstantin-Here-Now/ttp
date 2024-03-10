import datetime
from fastapi import APIRouter

from ..models.day import Day

router = APIRouter(
    prefix="/days",
    tags=["days"],
    responses={404: {"description": "Not found"}},
)

@router.get("/{date}")
async def read_day(date: datetime.date) -> Day:
    return Day