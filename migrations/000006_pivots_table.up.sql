CREATE TABLE public.post_tag
(
    post_id int8 NOT NULL,
    tag_id  int8 NOT NULL,
    CONSTRAINT post_tag_pkey PRIMARY KEY (post_id, tag_id)
);

ALTER TABLE public.post_tag
    ADD CONSTRAINT post_tag_post_id_foreign FOREIGN KEY (post_id) REFERENCES public.posts (id) ON DELETE CASCADE;
ALTER TABLE public.post_tag
    ADD CONSTRAINT post_tag_tag_id_foreign FOREIGN KEY (tag_id) REFERENCES public.tags (id) ON DELETE CASCADE;


CREATE TABLE public.category_post
(
    category_id int8 NOT NULL,
    post_id     int8 NOT NULL,
    CONSTRAINT category_post_pkey PRIMARY KEY (category_id, post_id)
);


ALTER TABLE public.category_post
    ADD CONSTRAINT category_post_category_id_foreign FOREIGN KEY (category_id) REFERENCES public.categories (id) ON DELETE CASCADE;
ALTER TABLE public.category_post
    ADD CONSTRAINT category_post_post_id_foreign FOREIGN KEY (post_id) REFERENCES public.posts (id) ON DELETE CASCADE;

ALTER TABLE public.posts
    ADD CONSTRAINT posts_author_id_foreign FOREIGN KEY (author_id) REFERENCES public.authors (id) ON DELETE CASCADE;