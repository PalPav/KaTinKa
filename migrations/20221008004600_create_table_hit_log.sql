-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.hit_log
(
    id serial NOT NULL,
    image_id integer,
    created_at timestamp with time zone DEFAULT CURRENT_TIMESTAMP,
    "desc" text COLLATE pg_catalog."default",
    CONSTRAINT hit_log_pkey PRIMARY KEY (id),
    CONSTRAINT image_fkey FOREIGN KEY (image_id)
        REFERENCES public.images (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.hit_log;
-- +goose StatementEnd
