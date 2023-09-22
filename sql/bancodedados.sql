CREATE TABLE "user" (
  "id" serial PRIMARY KEY,
  "name" varchar,
  "cpf" varchar,
  "balance" float,
  "created_at" timestamp
);

CREATE TABLE "balance_extract" (
  "id" serial PRIMARY KEY,
  "id_user" int,
  "balance" float,
  "id_user_receive" int
);

COMMENT ON COLUMN "user"."cpf" IS 'Documento do usuario';

ALTER TABLE "balance_extract" ADD FOREIGN KEY ("id_user") REFERENCES "user" ("id");

ALTER TABLE "balance_extract" ADD FOREIGN KEY ("id_user_receive") REFERENCES "user" ("id");
