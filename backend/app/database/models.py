from sqlalchemy import (
    TIME,
    TIMESTAMP,
    UUID,
    Boolean,
    Column,
    Date,
    ForeignKey,
    Integer,
    String,
)
from sqlalchemy.orm import relationship
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class User(Base):
    __tablename__ = "user"

    id = Column(UUID, primary_key=True)
    username = Column(String)
    password = Column(String)
    lastname = Column(String)
    firstname = Column(String)
    is_admin = Column(Boolean, default=False)
    notifications = relationship("Notification", back_populates="user")
    occupations = relationship("Occupation", back_populates="user")


class Notification(Base):
    __tablename__ = "notification"

    id = Column(Integer, primary_key=True)
    type = Column(String)
    link = Column(String)
    user_id = Column(UUID, ForeignKey("user.id"))
    user = relationship("User", back_populates="notifications")


class Occupation(Base):
    __tablename__ = "occupation"

    id = Column(Integer, primary_key=True)
    date = Column(Date)
    start = Column(TIME)
    end = Column(TIME)
    reason = Column(String)
    created_at = Column(TIMESTAMP)
    user_id = Column(UUID, ForeignKey("user.id"))
    lesson = relationship("Lesson", back_populates="occupation", uselist=False)
    user = relationship("User", back_populates="occupations")


class Lesson(Base):
    __tablename__ = "lesson"
    id = Column(Integer, primary_key=True)
    occupation_id = Column(Integer, ForeignKey("occupation.id"))
    occupation = relationship("Occupation", back_populates="lesson")


class Material(Base):
    __tablename__ = "material"
    id = Column(Integer, primary_key=True)
    name = Column(String)
    desc = Column(String)
    link = Column(String)


class FAQ(Base):
    __tablename__ = "faq"
    id = Column(Integer, primary_key=True)
    question = Column(String)
    answer = Column(String)
