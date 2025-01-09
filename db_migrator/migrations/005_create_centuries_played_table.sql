 CREATE TABLE "centuries_played" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "minutes_played" INTEGER NOT NULL,
    "total_price" NUMERIC(10, 2) NOT NULL,
    "status" SMALLINT NOT NULL,
    "created_by" BIGINT NOT NULL,
    "updated_by" BIGINT,
    "deleted_by" BIGINT
);

ALTER TABLE "centuries_played"
    ADD CONSTRAINT "centuries_played_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");