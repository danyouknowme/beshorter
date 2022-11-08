CREATE TABLE "shorturl" (
  "id" bigserial PRIMARY KEY,
  "full_url" varchar NOT NULL,
  "url" varchar NOT NULL,
  "click" bigint NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
)
