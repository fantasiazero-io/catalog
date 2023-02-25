-- Table: public.Boosters

-- DROP TABLE IF EXISTS public."Boosters";

CREATE TABLE IF NOT EXISTS public."Boosters"
(
    "Id" text COLLATE pg_catalog."default" NOT NULL,
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "RarityId" integer NOT NULL,
    "Action" jsonb NOT NULL,
    CONSTRAINT "Boosters_pkey" PRIMARY KEY ("Id")
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Boosters"
    OWNER to postgres;