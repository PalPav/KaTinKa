-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.tags
(
    id serial NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT tags_pkey PRIMARY KEY (id),
    CONSTRAINT tag_name UNIQUE (name)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.tags;
-- +goose StatementEnd
