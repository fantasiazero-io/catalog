-- Table: public.Teams

-- DROP TABLE IF EXISTS public."Teams";

CREATE TABLE IF NOT EXISTS public."Teams"
(
    "Id" text COLLATE pg_catalog."default" NOT NULL,
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "Characters" text[] COLLATE pg_catalog."default" NOT NULL,
    "Boosters" text[] COLLATE pg_catalog."default",
    "Equipments" text[] COLLATE pg_catalog."default",
    CONSTRAINT "Teams_pkey" PRIMARY KEY ("Id")
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Teams"
    OWNER to root;