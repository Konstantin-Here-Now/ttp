CREATE TABLE
  "timetable" (
    "id" SMALLINT PRIMARY KEY,
    "day" TEXT UNIQUE NOT NULL,
    "at" TEXT NOT NULL,
    "date" DATE NOT NULL,
  );

CREATE TABLE
  "material" (
    "id" SERIAL PRIMARY KEY,
    "name" TEXT UNIQUE NOT NULL,
    "desc" TEXT NOT NULL,
    "url" TEXT NOT NULL,
  );