package repository

import (
	"github.com/danyouknowme/beshorter/models"
	"github.com/jmoiron/sqlx"
)

type shortUrlRepository struct {
	db *sqlx.DB
}

type ShortUrlRepository interface {
	InsertShortenerUrl(fullUrl string, url string, click int) error
	SelectShortenerUrl(shortUrl string) (shortUrlInfo models.ShortUrl, err error)
	UpdateShortenerUrlClicks(shortUrl string) error
}

func NewShortUrlRepository(db *sqlx.DB) ShortUrlRepository {
	return &shortUrlRepository{
		db: db,
	}
}

func (s *shortUrlRepository) InsertShortenerUrl(fullUrl string, url string, click int) error {
	_, err := s.db.Query("INSERT INTO shorturl (full_url, url, click) VALUES ($1, $2, $3)", fullUrl, url, click)
	return err
}

func (s *shortUrlRepository) SelectShortenerUrl(shortUrl string) (shortUrlInfo models.ShortUrl, err error) {
	err = s.db.Get(&shortUrlInfo, "SELECT * from shorturl WHERE url=$1", shortUrl)
	if err != nil {
		return
	}

	return shortUrlInfo, nil
}

func (s *shortUrlRepository) UpdateShortenerUrlClicks(shortUrl string) error {
	_, err := s.db.Query("UPDATE shorturl SET click = click + 1 WHERE url=$1", shortUrl)
	return err
}
