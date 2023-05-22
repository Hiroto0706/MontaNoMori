package models

import (
	"log"
	"time"
)

type Image struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URLPath     string    `json:"url_path"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`

	PrevID int `json:"prev_id"`
	NextID int `json:"next_id"`

	Tags []Tag `json:"tags"`
}

// 画像をGCSに保存する処理
func CreateImage(title, description, url_path string) error {
	cmd := `insert into images (
		title,
		description,
		url_path,
		updated_at,
		created_at) values ($1, $2, $3, $4, $5)`

	_, err = Db.Exec(cmd, title, description, url_path, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}

	return err
}

// 画像の一覧を取得する処理
func GetImagePaths() (images []Image, err error) {
	cmd := `select id, title, description, url_path, updated_at, created_at from images order by id desc`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var image Image
		err = rows.Scan(
			&image.ID,
			&image.Title,
			&image.Description,
			&image.URLPath,
			&image.UpdatedAt,
			&image.CreatedAt)

		if err != nil {
			log.Println(err)
		}

		images = append(images, image)
	}
	rows.Close()

	return images, err
}

// 画像のURLを取得する処理
func GetImagePath(id int) (image Image, err error) {
	image = Image{}
	cmd := `select id, title, description, url_path, updated_at, created_at from images  where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&image.ID,
		&image.Title,
		&image.Description,
		&image.URLPath,
		&image.UpdatedAt,
		&image.CreatedAt)
	if err != nil {
		return image, nil
	}

	return image, err
}
func GetNextImagePath(id int) (image Image, err error) {
	image = Image{}
	cmd := `select id, title, description, url_path, updated_at, created_at from images  WHERE id > $1 ORDER BY id ASC LIMIT 1`
	err = Db.QueryRow(cmd, id).Scan(
		&image.ID,
		&image.Title,
		&image.Description,
		&image.URLPath,
		&image.UpdatedAt,
		&image.CreatedAt)
	if err != nil {
		return image, nil
	}

	return image, err
}
func GetPrevImagePath(id int) (image Image, err error) {
	image = Image{}
	cmd := `select id, title, description, url_path, updated_at, created_at from images  WHERE id < $1 ORDER BY id DESC LIMIT 1`
	err = Db.QueryRow(cmd, id).Scan(
		&image.ID,
		&image.Title,
		&image.Description,
		&image.URLPath,
		&image.UpdatedAt,
		&image.CreatedAt)
	if err != nil {
		return image, nil
	}

	return image, err
}

func GetImageByImagePath(path string) (image Image, err error) {
	image = Image{}
	cmd := `select id, title, description, url_path, updated_at, created_at from images  where url_path = $1`
	err = Db.QueryRow(cmd, path).Scan(
		&image.ID,
		&image.Title,
		&image.Description,
		&image.URLPath,
		&image.UpdatedAt,
		&image.CreatedAt)
	if err != nil {
		return image, nil
	}

	return image, err
}

func (i *Image) UpdateImage() (err error) {
	cmd := `update images set title = $1, description = $2, url_path = $3, updated_at = $4 where id = $5`
	_, _ = Db.Exec(cmd, i.Title, i.Description, i.URLPath, time.Now(), i.ID)
	if err != nil {
		return
	}

	return err
}

func (i *Image) DeleteImage() (err error) {
	err = DeleteAllImageTagsByImageID(i.ID)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := `delete from images where id = $1`
	_, err = Db.Exec(cmd, i.ID)
	if err != nil {
		log.Fatalln(err)
	}

	return err
}

func SearchImages(target string) (images []Image, err error) {
	cmd := `select id, title, description, url_path, updated_at, created_at from images where title LIKE '%' || $1 || '%' or description LIKE '%' || $1 || '%' order by id`
	rows, err := Db.Query(cmd, target)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var image Image
		err = rows.Scan(
			&image.ID,
			&image.Title,
			&image.Description,
			&image.URLPath,
			&image.UpdatedAt,
			&image.CreatedAt)

		if err != nil {
			log.Println(err)
		}

		images = append(images, image)
	}
	rows.Close()

	return images, err
}
