package database

import "katinka/services/container"

func GetTotalHits() (count int, err error) {
	rows, err := container.DB().Query(`SELECT count(*) FROM hit_log`)

	if err != nil {
		return count, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return count, err
		}
		break
	}

	return count, err
}

func RegisterHit(imageId int) (err error) {
	_, err = container.DB().Exec(
		"INSERT INTO hit_log (image_id) VALUES ($1)",
		imageId,
	)

	return
}
