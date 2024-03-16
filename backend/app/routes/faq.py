from uuid import UUID
from sqlalchemy.ext.asyncio import AsyncSession

from ..database.database import get_db_session
from ..database import models as db_models
from fastapi import APIRouter, Depends, status

from ..models.faq import FAQ

router = APIRouter(
    prefix="/faq",
    tags=["faq"],
    responses={404: {"description": "Not found"}},
)


@router.post("/", status_code=status.HTTP_201_CREATED)
async def create_faq(
    data: FAQ,
    session: AsyncSession = Depends(get_db_session),
) -> FAQ:
    faq = db_models.FAQ(**data.model_dump())
    faq.question.replace("?", "")
    session.add(faq)
    await session.commit()
    await session.refresh(faq)
    return faq


@router.get("/{id}")
async def read_faq(id: UUID) -> FAQ:
    return FAQ


@router.put("/{id}")
async def update_faq(id: UUID) -> None:
    return status.HTTP_200_OK


@router.delete("/{id}")
async def delete_faq(id: UUID) -> None:
    return status.HTTP_200_OK
