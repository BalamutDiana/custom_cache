# Go in-memory cache

---

Golang key:value map-based store with set/get/delete functions and error handling.


## Example:

---

```golang
package main

import (
	"fmt"

	"github.com/BalamutDiana/custom_cache"
)

func main() {
	cache := custom_cache.New()

	cache.Set("userId", 67)

	userId, err := cache.Get("userId")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(userId)
	}

	err = cache.Delete("userId")
	if err != nil {
		fmt.Println(err.Error())
	}

}

```
