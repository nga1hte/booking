CREATE TYPE "usertype" AS ENUM (
  'admin',
  'normal'
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar NOT NULL,
  "email" varchar NOT NULL,
  "mobile_number" varchar NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "userType" usertype NOT NULL
);

CREATE TABLE "bookings" (
  "bookingId" bigserial PRIMARY KEY,
  "bookedBy" int NOT NULL,
  "bookOn" timestamptz NOT NULL DEFAULT (now()),
  "bookStarts" timestamptz NOT NULL,
  "bookEnds" timestamptz NOT NULL
);

CREATE TABLE "payments" (
  "paymentId" bigserial PRIMARY KEY,
  "bookingId" int NOT NULL,
  "paymentDate" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("id");

CREATE INDEX ON "bookings" ("bookEnds");

CREATE INDEX ON "bookings" ("bookingId");

CREATE INDEX ON "payments" ("paymentDate");

ALTER TABLE "bookings" ADD FOREIGN KEY ("bookedBy") REFERENCES "users" ("id");

ALTER TABLE "payments" ADD FOREIGN KEY ("bookingId") REFERENCES "bookings" ("bookingId");
