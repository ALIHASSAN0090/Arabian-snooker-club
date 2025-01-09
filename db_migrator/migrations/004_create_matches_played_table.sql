 CREATE TABLE "matches_played" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "user_id" BIGINT NOT NULL,
    "price_id" BIGINT NOT NULL,
    "status" SMALLINT NOT NULL,
    "played_at" TIMESTAMP NOT NULL,
    "paid_at" TIMESTAMP,
    "paid_by" BIGINT,
    "no_of_matches" INTEGER NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" TIMESTAMP
);

ALTER TABLE "matches_played"
    ADD CONSTRAINT "matches_played_paid_by_foreign" FOREIGN KEY("paid_by") REFERENCES "users"("id");

ALTER TABLE "matches_played"
    ADD CONSTRAINT "matches_played_user_id_foreign" FOREIGN KEY("user_id") REFERENCES "users"("id");

ALTER TABLE "matches_played"
    ADD CONSTRAINT "matches_played_price_id_foreign" FOREIGN KEY("price_id") REFERENCES "matches"("id");