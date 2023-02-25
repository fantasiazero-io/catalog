-- Table: public.Mangas

-- DROP TABLE public."Mangas";

CREATE TABLE public."Mangas"
(
    "Id" text COLLATE pg_catalog."default" NOT NULL,
    "Name" text COLLATE pg_catalog."default" NOT NULL,
    "CreationDate" timestamp without time zone NOT NULL,
    "LastUpdate" timestamp without time zone,
    CONSTRAINT "Manga_pkey" PRIMARY KEY ("Id")
)

    TABLESPACE pg_default;

ALTER TABLE public."Mangas"
    OWNER to postgres;