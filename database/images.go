package database //TODO разобрать пакет по сущностям

import (
	"github.com/lib/pq"
	"katinka/services/container"
	"path/filepath"
	"strings"
)

type Image struct {
	Id        int    `json:"id"`
	ImageBlob []byte `json:"image_blob,omitempty"`
	Extension string `json:"extension"`
	OrigName  string `json:"orig_name,omitempty"`
}

func SaveImage(imgData []byte, fileName string) (err error) {
	ext := filepath.Ext(fileName)
	_, err = container.DB().Exec(
		"INSERT INTO images (image_blob, extension, orig_name) VALUES ($1, $2, $3)",
		imgData, strings.ToLower(ext), fileName,
	)

	return
}

func GetImage(id int) (img Image, err error) {
	rows, err := container.DB().Query(`SELECT id, image_blob, extension, orig_name FROM images WHERE id = $1`, id)

	if err != nil {
		return img, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&img.Id, &img.ImageBlob, &img.Extension, &img.OrigName)
		if err != nil {
			return img, err
		}
		break
	}

	return img, err
}

func GetRandImages(count int) ([]Image, error) {

	rows, err := container.DB().Query(`SELECT id, extension FROM images ORDER BY random() LIMIT $1`, count)

	images := make([]Image, 0, count)

	if err != nil {
		return images, err
	}

	defer rows.Close()

	for rows.Next() {
		var img Image
		err = rows.Scan(&img.Id, &img.Extension)
		if err != nil {
			return images, err
		}

		images = append(images, img)
	}

	return images, err
}

func GetLastUploaded(count int) ([]Image, error) {

	rows, err := container.DB().Query(`SELECT id, orig_name FROM images ORDER BY id DESC LIMIT $1`, count)

	images := make([]Image, 0, count)

	if err != nil {
		return images, err
	}

	defer rows.Close()

	for rows.Next() {
		var img Image
		err = rows.Scan(&img.Id, &img.OrigName)
		if err != nil {
			return images, err
		}

		images = append(images, img)
	}

	return images, err
}

func SearchImages(tagIds []int, page int) ([]Image, error) {

	rows, err := container.DB().Query(`
		SELECT
		    id, extension 
		FROM 
		    images 
		JOIN 
			image_tag ON image_tag.image_id = images.id
		WHERE
		    image_tag.tag_id = ANY ($1)
		GROUP BY 
		    images.id
		HAVING count(image_tag.tag_id) = $2
		LIMIT 20`, pq.Array(tagIds), len(tagIds))

	images := make([]Image, 0)

	if err != nil {
		return images, err
	}

	defer rows.Close()

	for rows.Next() {
		var img Image
		err = rows.Scan(&img.Id, &img.Extension)
		if err != nil {
			return images, err
		}

		images = append(images, img)
	}

	return images, err
}
