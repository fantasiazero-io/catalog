-- Table: public.Characters

-- DROP TABLE IF EXISTS public."Characters";

CREATE TABLE IF NOT EXISTS public."Characters"
(
    "Id" text COLLATE pg_catalog."default" NOT NULL,
    "MangaId" text COLLATE pg_catalog."default" NOT NULL,
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "ClassId" integer NOT NULL,
    "PersonalityId" integer NOT NULL,
    "LeadershipId" integer NOT NULL,
    "RarityId" integer NOT NULL,
    CONSTRAINT "Characters_pkey" PRIMARY KEY ("Id"),
    CONSTRAINT "Characters_MangaId_fkey" FOREIGN KEY ("MangaId")
    REFERENCES public."Mangas" ("Id") MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    )

    TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."Characters"
    OWNER to root;