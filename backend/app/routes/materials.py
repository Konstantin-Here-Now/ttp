from fastapi import APIRouter

from ..models.material import Material

router = APIRouter(
    prefix="/materials",
    tags=["materials"]
)

@router.get("/")
async def read_materials() -> list[Material]:
    return list[Material]