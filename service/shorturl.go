package service

import (
	"github.com/danyouknowme/beshorter/helper"
	"github.com/danyouknowme/beshorter/repository"

	logger "github.com/sirupsen/logrus"
)

type shortUrlService struct {
	shortUrlRepository repository.ShortUrlRepository
}

type ShortUrlService interface {
	CreateShortenerUrl(fullUrl string) (string, error)
	GetShortenerUrl(shortUrl string) (string, error)
}

func NewShortUrlService(shortUrlRepository repository.ShortUrlRepository) ShortUrlService {
	return &shortUrlService{
		shortUrlRepository: shortUrlRepository,
	}
}

func (s *shortUrlService) CreateShortenerUrl(fullUrl string) (string, error) {
	logger.Info("Start create url shortener")
	defer logger.Info("End create url shortener")

	url, err := helper.GenerateShorterUrl()
	if err != nil {
		return "", err
	}

	err = s.shortUrlRepository.InsertShortenerUrl(fullUrl, url, 0)
	return url, err
}

func (s *shortUrlService) GetShortenerUrl(shortUrl string) (string, error) {
	shortUrlInfo, err := s.shortUrlRepository.SelectShortenerUrl(shortUrl)
	if err != nil {
		return "", err
	}

	if err := s.shortUrlRepository.UpdateShortenerUrlClicks(shortUrl); err != nil {
		return "", err
	}

	return shortUrlInfo.FullUrl, nil
}
