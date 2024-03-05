import sys
sys.path.append("..")

import datetime
from uuid import UUID

from fastapi import FastAPI, status
from backend.models.faq import FAQ
from backend.models.material import Material
from backend.models.day import Day
from backend.models.lesson import Lesson
from backend.models.occupation import Occupation
from backend.models.timetable import Timetable

app = FastAPI()


@app.get("/timetable")
async def get_timetable() -> Timetable:
    return Timetable


@app.get("/occupation")
async def read_occupation(id: UUID) -> Occupation:
    return Occupation


@app.post("/occupation", status_code=status.HTTP_201_CREATED)
async def create_occupation(lesson: Occupation) -> Occupation:
    return Occupation


@app.put("/occupation")
async def update_occupation(id: UUID) -> None:
    return status.HTTP_200_OK


@app.delete("/occupation")
async def delete_occupation(id: UUID) -> None:
    return status.HTTP_200_OK


@app.get("/lessons")
async def read_lessons() -> list[Lesson]:
    return list[Lesson]


@app.get("/day")
async def read_day(date: datetime.date) -> Day:
    return Day


@app.get("/materials")
async def read_materials() -> list[Material]:
    return list[Material]


@app.post("/faq", status_code=status.HTTP_201_CREATED)
async def create_faq(faq: FAQ) -> FAQ:
    return FAQ


@app.get("/faq")
async def read_faq(id: UUID) -> FAQ:
    return FAQ


@app.put("/faq")
async def update_faq(id: UUID) -> None:
    return status.HTTP_200_OK


@app.delete("/faq")
async def delete_faq(id: UUID) -> None:
    return status.HTTP_200_OK
