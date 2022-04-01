CREATE TABLE public.categories
(
    id          bigserial    NOT NULL,
    "name"      varchar(120) NOT NULL,
    description varchar(500) NULL,
    is_default  boolean      NOT NULL DEFAULT false,
    status      varchar(100) NOT NULL,
    created_at  timestamp(0) NULL,
    updated_at  timestamp(0) NULL,
    CONSTRAINT categories_pkey PRIMARY KEY (id)
);