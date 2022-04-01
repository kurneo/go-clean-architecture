CREATE TABLE public.posts
(
    id          bigserial    NOT NULL,
    name        varchar(255) NOT NULL,
    description varchar(500) NOT NULL,
    "content"   text         NOT NULL,
    status      varchar(100) NOT NULL,
    type        varchar(100) NOT NULL,
    author_id   int8         NOT NULL,
    thumbnail   varchar(255) NOT NULL,
    "views"     int4         NOT NULL DEFAULT 0,
    feature     int2         NOT NULL DEFAULT '0':: smallint,
    created_at  timestamp(0) NULL,
    updated_at  timestamp(0) NULL,
    CONSTRAINT posts_pkey PRIMARY KEY (id)
);