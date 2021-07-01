package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(FirstMiddleware(),SecondMiddleware()) // (1)
	router.Run()
}

func FirstMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("first middleware before next()")
		isAbort := c.Query("isAbort")
		bAbort, err := strconv.ParseBool(isAbort)
		if err != nil {
			fmt.Printf("is abort value err, value %s\n", isAbort)
			c.Next() //
		}
		if bAbort {
			fmt.Println("first middleware abort") //(3)
			c.Abort()
			c.Next()
			//c.AbortWithStatusJSON(http.StatusOK, "abort is true")
			// return
		}

		fmt.Println("first middleware after next()")
	}
}

func SecondMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("current inside of second middleware")
	}
}

// first middleware before next()
// is abort value err, value truexx
// current inside of second middleware
// first middleware after next()
// [GIN] 2021/07/01 - 14:44:44 | 404 |      33.106µs |       127.0.0.1 | GET      "/?isAbort=truexx"
// first middleware before next()
// first middleware abort
// first middleware after next()
// [GIN] 2021/07/01 - 14:44:51 | 404 |      19.461µs |       127.0.0.1 | GET      "/?isAbort=true"