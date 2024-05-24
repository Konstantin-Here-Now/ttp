CREATE TABLE IF NOT EXISTS "default_occupation" (
  "id" SERIAL PRIMARY KEY,
  "day" TEXT UNIQUE NOT NULL,
  "at" TEXT NOT NULL,
  "date" DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS "default_occupation_changes" (
  "id" SERIAL PRIMARY KEY,
  "at" TEXT NOT NULL,
  "date" DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS "occupation_type" (
  "id" SERIAL PRIMARY KEY,
  "type" TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS "occupation" (
  "id" UUID PRIMARY KEY,
  "type_id" INT REFERENCES "occupation_type",
  "date" DATE NOT NULL,
  "start" TIME NOT NULL,
  "end" TIME NOT NULL,
  "desc" TEXT,
  "created_at" TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS "material" (
  "id" SERIAL PRIMARY KEY,
  "name" TEXT UNIQUE NOT NULL,
  "desc" TEXT NOT NULL,
  "url" TEXT NOT NULL
);