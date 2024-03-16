from fastapi import FastAPI
from fastapi.concurrency import asynccontextmanager
from fastapi.responses import RedirectResponse

from .migrate import migrate_tables

from .routes import days, faq, lessons, materials, occupations, timetable


@asynccontextmanager
async def lifespan(app: FastAPI):
    # Before start of app
    await migrate_tables()
    yield
    # Before end of app


app = FastAPI(lifespan=lifespan)
app.include_router(days.router)
app.include_router(faq.router)
app.include_router(lessons.router)
app.include_router(materials.router)
app.include_router(occupations.router)
app.include_router(timetable.router)


@app.get("/")
async def main_function():
    return RedirectResponse(url="/docs/")
