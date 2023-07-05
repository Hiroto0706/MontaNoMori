package handler

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"todo/app/models"
	"todo/config"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetImages(c echo.Context) error {
	images, err := models.GetImagePaths()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, images)
}

func SearchImages(c echo.Context) error {
	q := c.QueryParam("q")

	images, err := models.SearchImages(q)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, images)
}

func GetImage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	image, _ := models.GetImagePath(id)
	image_tagIDs, _ := models.GetImageTagIDFromImageID(image.ID)

	for _, image_tag := range image_tagIDs {
		tag, _ := models.GetTag(image_tag.TagID)
		image.Tags = append(image.Tags, tag)
	}

	return c.JSON(http.StatusOK, image)
}
func GetNextImage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	image, err := models.GetNextImagePath(id)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, image)
}
func GetPrevImage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	image, err := models.GetPrevImagePath(id)
	if err != nil {
		return nil
	}

	return c.JSON(http.StatusOK, image)
}

func GetImagesByTag(c echo.Context) error {
	name := c.Param("name")

	tag, err := models.GetTagByTagName(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	image_tags, _ := models.GetImageTagIDFromTagID(tag.ID)

	var images []models.Image
	for _, image_tag := range image_tags {
		image, _ := models.GetImagePath(image_tag.ImageID)
		images = append(images, image)
	}

	return c.JSON(http.StatusOK, images)
}

// 画像の新規作成する際に走らせるハンドラー
func UploadImage(c echo.Context) error {
	image := models.Image{}

	image.Title = c.FormValue("title")
	image.Description = c.FormValue("description")
	tags := c.Request().Form["tags[]"]

	// 画像ファイルを一時的なファイルとして保存
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	dst, src, ext, err := saveTempFile(file, c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// 認証されたGCSクライアントを作成
	ctx := context.Background()
	client, err := createGCSClient(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	// ファイルをGCSにアップロードする
	bucketName := config.Config.Bucket
	uniqueID := uuid.New().String()
	objectName := fmt.Sprintf("%s%s", uniqueID, ext)
	bucket := client.Bucket(bucketName)
	wc := bucket.Object(objectName).NewWriter(ctx)
	wc.ContentType = "image/png" // ContentTypeをimage/pngに設定
	defer os.Remove(dst.Name())
	defer dst.Close()
	if _, err = io.Copy(dst, src); err != nil {
		log.Fatalln(err)
	}

	// 一時ファイルの先頭に戻す
	dst.Seek(0, 0)

	defer wc.Close()
	if _, err := io.Copy(wc, dst); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	_, err = models.ConnectSqlite()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	url_path := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, objectName)
	err = models.CreateImage(image.Title, image.Description, url_path)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	image, _ = models.GetImageByImagePath(url_path)

	for _, id := range tags {
		id, _ := strconv.Atoi(id)
		tag, err := models.GetTag(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		image.Tags = append(image.Tags, tag)
	}

	for _, tag := range image.Tags {
		err = models.CreateImageTag(image.ID, tag.ID)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
	}

	return c.JSON(http.StatusOK, nil)
}

// 画像を編集するときに走らせるハンドラー
func EditImage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	title := c.FormValue("title")
	description := c.FormValue("description")
	tags := c.Request().Form["tags[]"]

	// 画像ファイルを一時的なファイルとして保存
	file, _ := c.FormFile("file")

	if file != nil {
		// 画像を受け取れば、既存の画像と書き換える処理を行う
		// 新たに取得した画像を一時保存する
		dst, src, ext, err := saveTempFile(file, c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		uniqueID := uuid.New().String()
		newObjectName := fmt.Sprintf("%s%s", uniqueID, ext)

		// 認証されたGCSクライアントを作成
		ctx := context.Background()
		client, err := createGCSClient(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// 既存画像の削除処理
		image, _ := models.GetImagePath(id)
		bucketName := config.Config.Bucket
		objectURL := image.URLPath
		objectName := path.Base(objectURL)
		err = deleteFileFromGCS(ctx, bucketName, objectName, client)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// ファイルをGCSにアップロードする
		bucket := client.Bucket(bucketName)
		wc := bucket.Object(newObjectName).NewWriter(ctx)
		wc.ContentType = "image/png" // ContentTypeをimage/pngに設定
		defer os.Remove(dst.Name())
		defer dst.Close()
		if _, err = io.Copy(dst, src); err != nil {
			log.Fatalln(err)
		}

		// 一時ファイルの先頭に戻す
		dst.Seek(0, 0)

		defer wc.Close()
		if _, err := io.Copy(wc, dst); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}

		// 編集した内容をDBに保存
		image.Title = title
		image.Description = description
		image.URLPath = fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, newObjectName)

		err = models.DeleteAllImageTagsByImageID(image.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		for _, id := range tags {
			id, _ := strconv.Atoi(id)
			_ = models.CreateImageTag(image.ID, id)
		}

		err = image.UpdateImage()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	} else {
		// 新たに画像を受け取らなかった場合の処理
		image, err := models.GetImagePath(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		image.Title = title
		err = models.DeleteAllImageTagsByImageID(image.ID)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
		image.Description = description
		for _, id := range tags {
			id, _ := strconv.Atoi(id)
			_ = models.CreateImageTag(image.ID, id)
		}

		err = image.UpdateImage()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}
	}

	image, _ := models.GetImagePath(id)
	image_tagIDs, _ := models.GetImageTagIDFromImageID(image.ID)

	for _, image_tag := range image_tagIDs {
		tag, _ := models.GetTag(image_tag.TagID)
		image.Tags = append(image.Tags, tag)
	}

	return c.JSON(http.StatusOK, image)
}

func DeleteImage(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	// 削除処理の準備
	image, _ := models.GetImagePath(id)
	bucketName := config.Config.Bucket
	objectURL := image.URLPath
	objectName := path.Base(objectURL)
	ctx := context.Background()
	client, err := createGCSClient(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}
	err = deleteFileFromGCS(ctx, bucketName, objectName, client)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	err = image.DeleteImage()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}
