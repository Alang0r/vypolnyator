CREATE TABLE IF NOT EXISTS "user" (
    id bigserial PRIMARY KEY,
    telegram_id varchar UNIQUE NOT NULL,
    name varchar,
    created_on TIMESTAMP NOT NULL
)