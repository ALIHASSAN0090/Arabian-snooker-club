 CREATE TABLE "centuries" (
    "id" BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    "price_per_minute" NUMERIC(10, 2) NOT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP 
);