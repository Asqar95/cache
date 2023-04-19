package cach

import (
	"github.com/Asqar95/cache.git/internal"
)

func Version() string {
	return internal.Version
}

func New() *internal.Cache {
	return internal.New()
}
