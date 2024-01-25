CREATE TABLE IF NOT EXISTS "users" (
  "id" BIGSERIAL PRIMARY KEY,
  "phone_number" VARCHAR(16) UNIQUE NOT NULL,
  "full_name" VARCHAR(60) NOT NULL,
  "password_hash" VARCHAR(255) NOT NULL,
  "successful_logins" BIGINT DEFAULT 0 NOT NULL,
  "created_at" TIMESTAMP DEFAULT 'now()'::TIMESTAMP WITHOUT TIME ZONE NOT NULL,
  "last_login_at" TIMESTAMP DEFAULT 'now()'::TIMESTAMP WITHOUT TIME ZONE NOT NULL
);

CREATE INDEX IF NOT EXISTS "users_phone_number_index" ON "users" ("phone_number");

CREATE TABLE IF NOT EXISTS "login_attempts" (
  "id" BIGSERIAL PRIMARY KEY,
  "user_id" BIGSERIAL REFERENCES "users"("id") NOT NULL,
  "success" BOOLEAN,
  "attempted_at" TIMESTAMP DEFAULT 'now()'::TIMESTAMP WITHOUT TIME ZONE NOT NULL
);
