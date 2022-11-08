package helper

import "github.com/teris-io/shortid"

func GenerateShorterUrl() (string, error) {
	sid, err := shortid.New(1, shortid.DefaultABC, 2342)
	if err != nil {
		return "", err
	}

	return sid.Generate()
}
