CREATE TABLE "user_balance" (
  "id" serial PRIMARY KEY,
  "name" string,
  "document" string,
  "dateCreate" timestamp,
  "dateDelete" timestamp
);

CREATE TABLE "extract_user" (
  "id" serial PRIMARY KEY,
  "typeOperation" string,
  "idUserOrigem" int,
  "valueOrigem" numeric(15,2),
  "idUserDestiny" int,
  "valueDestiny" numeric(15,2),
  "dateInsert" timestamp
);

CREATE TABLE "balance" (
  "id" serial,
  "idUser" int,
  "value" numeric(15,2)
);

ALTER TABLE "extract_user" ADD FOREIGN KEY ("idUserOrigem") REFERENCES "user_balance" ("id");

ALTER TABLE "extract_user" ADD FOREIGN KEY ("idUserDestiny") REFERENCES "user_balance" ("id");

ALTER TABLE "balance" ADD FOREIGN KEY ("idUser") REFERENCES "user_balance" ("id");
