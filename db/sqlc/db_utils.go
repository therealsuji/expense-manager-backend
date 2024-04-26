package db

import "github.com/nrednav/cuid2"

func GenerateId() string {
	return cuid2.Generate()
}
