CREATE TABLE public.files
(
    id         bigserial    NOT NULL,
    path       text         NOT NULL UNIQUE,
    size       varchar(100) NOT NULL,
    mime       varchar(100) NOT NULL,
    created_at timestamp(0) NULL,
    updated_at timestamp(0) NULL,
    CONSTRAINT files_pkey PRIMARY KEY (id)
);