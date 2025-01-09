 CREATE TABLE "users" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR(100) NOT NULL,
    "email" VARCHAR(320),
    "password" VARCHAR(255),
    "role" VARCHAR(50) NOT NULL
);