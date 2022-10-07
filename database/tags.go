package database

import (
	"database/sql"
	"katinka/services/container"
)

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func AddTag(name string) (tag Tag, err error) {
	err = container.DB().QueryRow(
		`INSERT INTO tags ("name") VALUES ($1) ON CONFLICT DO NOTHING RETURNING id`,
		name,
	).Scan(&tag.Id)

	if err == nil {
		tag.Name = name
	}

	return tag, err
}

func GetTagsList() (tags []Tag, err error) {
	rows, err := container.DB().Query(`SELECT id, "name" FROM tags ORDER BY "name"`)
	if err != nil {
		return tags, err
	}

	return fetchManyTags(rows)
}

func GetTag(tagId int) (tag Tag, err error) {
	rows, err := container.DB().Query(
		`SELECT id, "name" FROM tags WHERE id = $1 LIMIT 1`, tagId)

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return
		}
		break
	}
	return
}

func EditTag(tag Tag) error {
	_, err := container.DB().Exec(`UPDATE tags set "name"=$1 WHERE id = $2`, tag.Name, tag.Id)
	return err
}

func SearchTags(name string) (tags []Tag, err error) {
	rows, err := container.DB().Query(`SELECT id, "name" FROM tags WHERE "name" LIKE  $1 || '%' ORDER BY "name" LIMIT 20`, name)
	if err != nil {
		return tags, err
	}

	return fetchManyTags(rows)
}

func GetTagsByImageId(imageId int) (tags []Tag, err error) {
	rows, err := container.DB().Query(
		`SELECT tags.* FROM tags JOIN image_tag ON image_tag.tag_id = tags.id WHERE image_tag.image_id = $1`,
		imageId,
	)

	if err != nil {
		return tags, err
	}

	return fetchManyTags(rows)
}

func AssignTag(imageId int, tagId int) error {
	_, err := container.DB().Exec(
		`INSERT INTO image_tag(image_id, tag_id) VALUES($1, $2)  ON CONFLICT DO NOTHING`, imageId, tagId)
	return err
}

func DetachTag(imageId int, tagId int) error {
	_, err := container.DB().Exec(
		`DELETE FROM image_tag WHERE image_id = $1 AND tag_id = $2`, imageId, tagId)
	return err
}

func RemoveTag(tagId int) error {
	_, err := container.DB().Exec(
		`DELETE FROM image_tag WHERE tag_id = $1`, tagId)

	if err != nil {
		return err
	}
	_, err = container.DB().Exec(
		`DELETE FROM tags WHERE id = $1`, tagId)

	return err
}

func fetchManyTags(rows *sql.Rows) (tags []Tag, err error) {
	defer rows.Close()

	for rows.Next() {
		var tag Tag
		err = rows.Scan(&tag.Id, &tag.Name)
		if err != nil {
			return tags, err
		}

		tags = append(tags, tag)
	}

	return
}
