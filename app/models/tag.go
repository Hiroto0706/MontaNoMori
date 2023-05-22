package models

import (
	"log"
	"time"
)

type Tag struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateTag(name string) error {
	cmd := `insert into tags (
		name,
		updated_at,
		created_at) values ($1, $2, $3)`

	_, err = Db.Exec(cmd, name, time.Now(), time.Now())
	if err != nil {
		log.Println(err)
	}

	return err
}

func GetTags() (tags []Tag, err error) {
	cmd := `select id, name, updated_at, created_at from tags order by id`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var tag Tag
		err = rows.Scan(
			&tag.ID,
			&tag.Name,
			&tag.UpdatedAt,
			&tag.CreatedAt)

		if err != nil {
			log.Println(err)
		}

		tags = append(tags, tag)
	}
	rows.Close()

	return tags, err
}

func GetTag(id int) (tag Tag, err error) {
	tag = Tag{}
	cmd := `select id, name, updated_at, created_at from tags  where id = $1`
	err = Db.QueryRow(cmd, id).Scan(
		&tag.ID,
		&tag.Name,
		&tag.UpdatedAt,
		&tag.CreatedAt)
	if err != nil {
		return tag, nil
	}

	return tag, err
}

func GetTagByTagName(name string) (tag Tag, err error) {
	tag = Tag{}
	cmd := `select id, name, updated_at, created_at from tags  where name = $1`
	err = Db.QueryRow(cmd, name).Scan(
		&tag.ID,
		&tag.Name,
		&tag.UpdatedAt,
		&tag.CreatedAt)
	if err != nil {
		return tag, nil
	}

	return tag, err
}

func (t *Tag) UpdateTag() (err error) {
	cmd := `update tags set name = $1, updated_at = $2 where id = $3`
	_, _ = Db.Exec(cmd, t.Name, time.Now(), t.ID)
	if err != nil {
		return err
	}

	return err
}

func (t *Tag) DeleteTag() (err error) {
	err = DeleteAllImageTagsByTagID(t.ID)
	if err != nil {
		log.Fatalln(err)
	}

	cmd := `delete from tags where id = $1`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		return err
	}

	return err
}
