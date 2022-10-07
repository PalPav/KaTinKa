-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.images
(
    id serial NOT NULL,
    image_blob bytea NOT NULL,
    extension text COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    orig_name text COLLATE pg_catalog."default",
    CONSTRAINT images_pkey PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.images;
-- +goose StatementEnd
