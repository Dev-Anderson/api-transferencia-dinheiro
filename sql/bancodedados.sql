CREATE TABLE "user" (
  "id" serial PRIMARY KEY,
  "name" string,
  "document" string,
  "dateCreate" timestamp,
  "dateDelete" timestamp
);

CREATE TABLE "extractUser" (
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

ALTER TABLE "extractUser" ADD FOREIGN KEY ("idUserOrigem") REFERENCES "user" ("id");

ALTER TABLE "extractUser" ADD FOREIGN KEY ("idUserDestiny") REFERENCES "user" ("id");

ALTER TABLE "balance" ADD FOREIGN KEY ("idUser") REFERENCES "user" ("id");
