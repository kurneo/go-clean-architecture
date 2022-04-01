CREATE TABLE public.authors
(
    id         bigserial    NOT NULL,
    name       varchar(255) NOT NULL,
    status     varchar(100) NOT NULL,
    avatar     varchar(255) NULL,
    created_at timestamp(0) NULL,
    updated_at timestamp(0) NULL,
    CONSTRAINT authors_pkey PRIMARY KEY (id)
);