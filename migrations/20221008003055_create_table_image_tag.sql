-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS public.image_tag
(
    image_id integer NOT NULL,
    tag_id integer NOT NULL,
    CONSTRAINT image_tag_pkey PRIMARY KEY (image_id, tag_id),
    CONSTRAINT image_fkey FOREIGN KEY (image_id)
        REFERENCES public.images (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT tag_fkey FOREIGN KEY (tag_id)
        REFERENCES public.tags (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS public.image_tag;
-- +goose StatementEnd
