Go in-memory Cache.
=======================

See it in action:

```go
package main

import (
	"fmt"
	"github.com/Asqar95/cache"
)

func main() {
	cache := cache.New("10h00m")

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)
}
```