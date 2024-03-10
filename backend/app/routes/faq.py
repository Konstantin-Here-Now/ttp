from uuid import UUID
from fastapi import APIRouter, status

from ..models.faq import FAQ

router = APIRouter(
    prefix="/faq",
    tags=["faq"],
    responses={404: {"description": "Not found"}},
)


@router.post("/", status_code=status.HTTP_201_CREATED)
async def create_faq(faq: FAQ) -> FAQ:
    return FAQ


@router.get("/{id}")
async def read_faq(id: UUID) -> FAQ:
    return FAQ


@router.put("/{id}")
async def update_faq(id: UUID) -> None:
    return status.HTTP_200_OK


@router.delete("/{id}")
async def delete_faq(id: UUID) -> None:
    return status.HTTP_200_OK
