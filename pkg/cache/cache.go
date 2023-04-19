package cach

import (
	"github.com/Asqar95/cache.git/internal"
	"time"
)

func main() {
	Cache().Set("2", "userid")
}

func Cache() *internal.Cache {
	return internal.New(time.Second * 2)
}
