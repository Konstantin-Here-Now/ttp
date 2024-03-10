from fastapi import FastAPI
from fastapi.responses import RedirectResponse

from .routes import days, faq, lessons, materials, occupations, timetable


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
