package handler

import (
	"net/http"
	"strconv"
	"todo/app/models"

	"github.com/labstack/echo/v4"
)

func GetTags(c echo.Context) error {
	tags, err := models.GetTags()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, tags)
}

func CreateTag(c echo.Context) error {
	name := c.FormValue("name")
	err := models.CreateTag(name)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

func UpdateTag(c echo.Context) error {
	var newTag models.Tag
	if err := c.Bind(&newTag); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err := newTag.UpdateTag()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}

func DeleteTag(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	tag, err := models.GetTag(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = tag.DeleteTag()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, nil)
}
