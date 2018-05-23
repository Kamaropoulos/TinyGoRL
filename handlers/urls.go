package handlers

import (
	"net/http"

	"github.com/Kamaropoulos/TinyGoRL/models"

	"github.com/labstack/echo"

	Log "github.com/sirupsen/logrus"
)

type H map[string]interface{}

// GetURLs endpoint
func GetURLs() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetURLs())
	}
}

func GetURL() echo.HandlerFunc {
	return func(c echo.Context) error {
		url := c.Param("url")
		Log.Debug(url)
		result, err := models.GetURL(url)
		// return c.Redirect(http.StatusTemporaryRedirect, result.URL)

		if err == nil {
			return c.Redirect(http.StatusTemporaryRedirect, result.URL)
			// Handle any errors
		} else {
			return err
		}
	}
}

// PutURL endpoint
func PutURL() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new URL
		var url models.URL
		// Map imcoming JSON body to the new URL
		c.Bind(&url)
		Log.Debug(url.URL)
		// Add a url using our new model
		id, err := models.PutURL(url.URL)
		// Return a JSON response if successful
		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
			// Handle any errors
		} else {
			return err
		}
	}
}

// DeleteURL endpoint
func DeleteURL() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("url")
		// Use our new model to delete a url
		_, err := models.DeleteURL(id)
		// Return a JSON response on success
		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
			// Handle errors
		} else {
			return err
		}
	}
}
