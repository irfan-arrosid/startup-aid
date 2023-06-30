CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "occupation" varchar,
  "email" varchar NOT NULL,
  "password_hash" varchar NOT NULL,
  "avatar_filename" varchar,
  "role" varchar,
  "token" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "campaigns" (
  "id" bigserial PRIMARY KEY,
  "user_id" integer,
  "name" varchar NOT NULL,
  "short_description" varchar,
  "description" text,
  "goal_amount" integer,
  "current_amount" integer,
  "perks" text,
  "backer_count" integer,
  "slug" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "campaign_images" (
  "id" bigserial PRIMARY KEY,
  "campaign_id" integer,
  "filename" varchar NOT NULL,
  "is_primary" smallint,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

CREATE TABLE "transactions" (
  "id" bigserial PRIMARY KEY,
  "campaign_id" integer,
  "user_id" integer,
  "amount" integer,
  "status" varchar,
  "code" varchar,
  "created_at" timestamptz DEFAULT (now()),
  "updated_at" timestamptz DEFAULT (now())
);

ALTER TABLE "campaigns" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "campaign_images" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("campaign_id") REFERENCES "campaigns" ("id");

ALTER TABLE "transactions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
