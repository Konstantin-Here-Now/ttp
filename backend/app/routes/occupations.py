from uuid import UUID
from fastapi import APIRouter, status

from ..models.occupation import Occupation

router = APIRouter(
    prefix="/occupations",
    tags=["occupations"],
    responses={404: {"description": "Not found"}},
)

@router.get("/")
async def read_occupation(id: UUID) -> Occupation:
    return Occupation


@router.post("/", status_code=status.HTTP_201_CREATED)
async def create_occupation(lesson: Occupation) -> Occupation:
    return Occupation


@router.put("/{id}")
async def update_occupation(id: UUID) -> None:
    return status.HTTP_200_OK


@router.delete("/{id}")
async def delete_occupation(id: UUID) -> None:
    return status.HTTP_200_OK