package handler

import (
	"context"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"todo/config"

	"cloud.google.com/go/storage"
	"github.com/labstack/echo/v4"
	"google.golang.org/api/option"
)

// 一時ファイル、fileから読み取ったファイル、拡張子を作成して返す
func saveTempFile(file *multipart.FileHeader, c echo.Context) (*os.File, multipart.File, string, error) {
	// ファイルのContent-Typeが"image/png"以外の場合、適切なMIMEタイプに変更する
	if file.Header.Get("Content-Type") != "image/png" {
		file.Header.Set("Content-Type", "image/png")
	}

	src, err := file.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)
	dst, err := os.CreateTemp("", "temp-image-*"+ext)
	if err != nil {
		log.Fatalln(err)
	}

	return dst, src, ext, err
}

// GCSクライアントとの接続
func createGCSClient(ctx context.Context) (*storage.Client, error) {
	credentialFilePath := config.Config.Json
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		log.Fatalln(err)
	}

	defer client.Close()

	return client, err
}

// GCS上の画像を削除する
func deleteFileFromGCS(ctx context.Context, bucketName string, fileName string, client *storage.Client) error {
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(fileName)

	err := obj.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
