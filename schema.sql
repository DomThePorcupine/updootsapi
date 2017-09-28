-- Adminer 4.3.1 PostgreSQL dump

DROP TABLE IF EXISTS "messages";
CREATE SEQUENCE messages_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 9223372036854775807 START 1 CACHE 1;

CREATE TABLE "public"."messages" (
    "id" integer DEFAULT nextval('messages_id_seq') NOT NULL,
    "message" character varying(250),
    "userid" character varying(20),
    "created" timestamp DEFAULT now(),
    CONSTRAINT "messages_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "votes";
CREATE TABLE "public"."votes" (
    "message" integer NOT NULL,
    "userid" character varying(20),
    "updoot" smallint DEFAULT 0 NOT NULL,
    "downdoot" smallint DEFAULT 0 NOT NULL
) WITH (oids = false);


DROP TABLE IF EXISTS "users";
CREATE TABLE "public"."users" (
    "created" timestamp DEFAULT now(),
    "userid" character varying(20) NOT NULL,
    "admin" smallint DEFAULT 0 NOT NULL,
    CONSTRAINT "users_userid_key" UNIQUE ("userid")
) WITH (oids = false);


-- 2017-08-09 12:13:59.450851+00