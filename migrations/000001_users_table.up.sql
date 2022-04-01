CREATE TABLE public.users
(
    id            bigserial    NOT NULL,
    username      varchar(100) NOT NULL,
    dob           date NULL,
    about         varchar(255) NULL,
    avatar        varchar(255) NULL,
    name          varchar(255) NOT NULL,
    email         varchar(255) NULL,
    password      varchar(255) NOT NULL,
    gender        varchar(255) NOT NULL DEFAULT 'male':: character varying,
    last_login_at timestamp(0) NULL,
    created_at    timestamp(0) NULL,
    updated_at    timestamp(0) NULL,
    CONSTRAINT users_gender_check CHECK (((gender)::text = ANY ((ARRAY['male':: character varying, 'female':: character varying])::text[])
) ),
	CONSTRAINT users_pkey PRIMARY KEY (id),
	CONSTRAINT users_username_unique UNIQUE (username)
);

INSERT INTO "users"
    ("username","dob","about","avatar","name","email","password","gender","last_login_at","created_at","updated_at")
    VALUES
   ('supersu','1997-04-15','about','','Giang Nguyen','giang@gmail.com','$2a$14$H4g2bAIPI7SYNJHrgbZhTu9IoD9/SwMFbFC3aqI3LtEfZiYu5b4xS','male','2021-11-10 18:02:53.769','2021-11-10 18:02:53.769','2021-11-10 18:02:53.769')