package models

import (
	"log"
	"time"
)

type Image_Tags struct {
	ID        int       `json:"id"`
	ImageID   int       `json:"image_id"`
	TagID     int       `json:"tag_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateImageTag(image_id, tag_id int) error {
	cmd := `insert into image_tags (
		image_id,
		tag_id,
		updated_at,
		created_at) values ($1, $2, $3, $4)`

	_, err = Db.Exec(cmd, image_id, tag_id, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}

	return err
}

func GetImageTagIDFromImageID(image_id int) (image_tags []Image_Tags, err error) {
	cmd := `select id, image_id, tag_id, updated_at, created_at from image_tags where image_id = $1 order by image_id desc`
	rows, err := Db.Query(cmd, image_id)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var image_tag Image_Tags
		err = rows.Scan(
			&image_tag.ID,
			&image_tag.ImageID,
			&image_tag.TagID,
			&image_tag.UpdatedAt,
			&image_tag.CreatedAt)

		if err != nil {
			log.Println(err)
		}

		image_tags = append(image_tags, image_tag)
	}
	rows.Close()

	return image_tags, err
}
func GetImageTagIDFromTagID(tag_id int) (image_tags []Image_Tags, err error) {
	cmd := `select id, image_id, tag_id, updated_at, created_at from image_tags where tag_id = $1 order by image_id desc`
	rows, err := Db.Query(cmd, tag_id)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var image_tag Image_Tags
		err = rows.Scan(
			&image_tag.ID,
			&image_tag.ImageID,
			&image_tag.TagID,
			&image_tag.UpdatedAt,
			&image_tag.CreatedAt)

		if err != nil {
			log.Println(err)
		}

		image_tags = append(image_tags, image_tag)
	}
	rows.Close()

	return image_tags, err
}

func DeleteAllImageTagsByImageID(id int) (err error) {
	cmd := `delete from image_tags where image_id = $1`
	_, err = Db.Exec(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func DeleteAllImageTagsByTagID(id int) (err error) {
	cmd := `delete from image_tags where tag_id = $1`
	_, err = Db.Exec(cmd, id)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}
