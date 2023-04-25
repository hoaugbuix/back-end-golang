-- +migrate Up
CREATE table "users" (
    "id" text not null PRIMARY KEY,
    "username" text not null,
    "email" text not null unique,
    "password" text not null,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
create TABLE "roles" (
    "id" text not null PRIMARY KEY,
    "role_name" text not null unique,
    "description" text,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
create TABLE "user_roles" (
    "id" text not null PRIMARY KEY,
    "role_id" text not null unique,
    "user_id" text not null unique,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
create TABLE "tags" (
    "id" text not null PRIMARY KEY,
    "tag_name" text not null unique,
    "description" text,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
create TABLE "posts" (
    "id" text not null PRIMARY KEY,
    "content" text not null unique,
    "thumbnail" text not null,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
create TABLE "post_tags" (
    "id" text not null PRIMARY KEY,
    "post_id" text not null unique,
    "tag_id" text not null unique,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
create TABLE "post_cates" (
    "id" text not null PRIMARY KEY,
    "post_id" text not null unique,
    "cate_id" text not null unique,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
create TABLE "categories" (
    "id" text not null PRIMARY KEY,
    "category_name" text not null unique,
    "category_code" text not null unique,
    "description" text,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
CREATE TABLE "languages" (
    "id" text not null PRIMARY KEY,
    "en" text not null unique,
    "vn" text not null unique,
    "location" text,
    "active_flag" int DEFAULT 1 not null,
    "created_date" TIMESTAMPTZ NOT NULL,
    "updated_date" TIMESTAMPTZ NOT NULL
);
-- +migrate Down
DROP TABLE "users";
DROP TABLE "roles";
DROP TABLE "user_roles";
DROP TABLE "tags";
DROP TABLE "post_tags";
DROP TABLE "posts";
DROP TABLE "categories";
DROP TABLE "post_cates";
DROP TABLE "languages";