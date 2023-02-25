-- Table: public.Equipments

-- DROP TABLE IF EXISTS public."Equipments";

CREATE TABLE IF NOT EXISTS public."Equipments"
(
    "Id" text COLLATE pg_catalog."default" NOT NULL,
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "Action" jsonb NOT NULL,
    CONSTRAINT "Equipments_pkey" PRIMARY KEY ("Id")
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Equipments"
    OWNER to root;