package utils

import (
	"errors"

	"github.com/tesla59/shortify/database"
	"github.com/tesla59/shortify/models"
)

func ValidateURL(URL models.URL) error {
	if URL.TargetURL == "" {
		return errors.New("Target URL not specified")
	}
	url, _ := database.FindURLbyShortURL(URL.ShortURL)
	if url.ID != "" {
		return errors.New("Cannot create URL with short URL: " + URL.ShortURL)
	}
	return nil
}
