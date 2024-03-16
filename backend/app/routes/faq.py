from uuid import UUID
from sqlalchemy import select, update
from sqlalchemy.ext.asyncio import AsyncSession

from ..database.database import get_db_session
from ..database import models as db_models
from fastapi import APIRouter, Depends, HTTPException, status

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
async def read_faq(
    id: int,
    session: AsyncSession = Depends(get_db_session),
) -> FAQ:
    faq = await session.get(db_models.FAQ, id)
    if faq is None:
        raise HTTPException(
            status_code=status.HTTP_404_NOT_FOUND,
            detail="Ingredient does not exist",
        )
    return faq


@router.get("/", status_code=status.HTTP_200_OK)
async def read_faqs(
    session: AsyncSession = Depends(get_db_session),
) -> list[FAQ]:
    faqs = await session.scalars(select(db_models.FAQ))
    return faqs


@router.put("/{id}")
async def update_faq(
    id: int,
    data: FAQ,
    session: AsyncSession = Depends(get_db_session),
) -> None:
    faq = await session.get(db_models.FAQ, id)
    faq.question = data.question
    faq.answer = data.answer
    await session.commit()
    return status.HTTP_200_OK


@router.delete("/{id}")
async def delete_faq(
    id: int,
    session: AsyncSession = Depends(get_db_session),
) -> None:
    faq = await session.get(db_models.FAQ, id)
    await session.delete(faq)
    await session.commit()
    return status.HTTP_200_OK
