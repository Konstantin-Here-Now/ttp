from fastapi import FastAPI
from fastapi.responses import RedirectResponse
from .database_models import Base
from .database import engine

from .routes import days, faq, lessons, materials, occupations, timetable

Base.metadata.create_all(bind=engine)

app = FastAPI()
app.include_router(days.router)
app.include_router(faq.router)
app.include_router(lessons.router)
app.include_router(materials.router)
app.include_router(occupations.router)
app.include_router(timetable.router)


@app.get("/")
def main_function():
    return RedirectResponse(url="/docs/")
