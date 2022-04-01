CREATE TABLE public.tags
(
    id          bigserial    NOT NULL,
    name        varchar(255) NOT NULL,
    description varchar(255) NULL,
    status      varchar(100) NOT NULL,
    created_at  timestamp(0) NULL,
    updated_at  timestamp(0) NULL,
    CONSTRAINT tags_pkey PRIMARY KEY (id)
);