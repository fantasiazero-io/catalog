-- Table: public.Matches

-- DROP TABLE IF EXISTS public."Matches";

CREATE TABLE IF NOT EXISTS public."Matches"
(
    "Id" text COLLATE pg_catalog."default" NOT NULL,
    "MyTeamId" text COLLATE pg_catalog."default" NOT NULL,
    "OpponentTeamId" text COLLATE pg_catalog."default" NOT NULL,
    "MyScore" numeric NOT NULL,
    "OpponentScore" numeric NOT NULL,
    "Status" integer NOT NULL,
    CONSTRAINT "Matches_pkey" PRIMARY KEY ("Id")
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Matches"
    OWNER to root;