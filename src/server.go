// // package main

// // import (
// // 	"fmt"
// // 	"log"
// // 	"net/http"
// // 	"time"

// // 	"github.com/gin-gonic/gin"
// // )

// // func main() {
// // 	router := gin.Default()
// // 	router.GET("/ping", func(c *gin.Context) {
// // 		ip := c.ClientIP()
// // 		fmt.Printf(ip + "\n")
// // 		c.JSON(200, gin.H{
// // 			"message": "pong",
// // 		})
// // 	})

// // 	router.GET("/user/:name", func(c *gin.Context) {
// // 		name := c.Param("name")
// // 		c.String(http.StatusOK, "Hello %s", name)
// // 	})

// // 	router.POST("/form_post", func(c *gin.Context) {
// // 		message := c.PostForm("message")
// // 		nick := c.DefaultPostForm("nick", "anonymous")

// // 		c.JSON(200, gin.H{
// // 			"status":  "posted",
// // 			"message": message,
// // 			"nick":    nick,
// // 		})
// // 	})

// // 	router.GET("/welcome", func(c *gin.Context) {
// // 		firstname := c.DefaultQuery("firstname", "Guest")
// // 		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

// // 		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
// // 	})

// // 	router.POST("/post", func(c *gin.Context) {

// // 		id := c.Query("id")
// // 		page := c.DefaultQuery("page", "0")
// // 		name := c.PostForm("name")
// // 		message := c.PostForm("message")

// // 		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
// // 		c.String(http.StatusOK, "Hello %s %s  %s %s", id, page, name, message)
// // 	})

// // 	router.POST("/upload", func(c *gin.Context) {
// // 		// single file
// // 		file, _ := c.FormFile("file")
// // 		log.Println(file.Filename)

// // 		// Upload the file to specific dst.
// // 		dst := "/"
// // 		c.SaveUploadedFile(file, dst)

// // 		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
// // 	})

// // 	router.GET("/someJSON", func(c *gin.Context) {
// // 		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
// // 	})

// // 	router.GET("/someJSONs", func(c *gin.Context) {
// // 		names := []string{"lena", "austin", "foo"}

// // 		// Will output  :   while(1);["lena","austin","foo"]
// // 		c.SecureJSON(http.StatusOK, names)
// // 	})

// // 	router.GET("/someDataFromReader", func(c *gin.Context) {
// // 		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
// // 		if err != nil || response.StatusCode != http.StatusOK {
// // 			c.Status(http.StatusServiceUnavailable)
// // 			return
// // 		}

// // 		reader := response.Body
// // 		contentLength := response.ContentLength
// // 		contentType := response.Header.Get("Content-Type")

// // 		extraHeaders := map[string]string{
// // 			"Content-Disposition": `attachment; filename="gopher.png"`,
// // 		}

// // 		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
// // 	})

// // 	router.GET("/test", func(c *gin.Context) {
// // 		c.Redirect(http.StatusMovedPermanently, "http://www.google.com/")
// // 	})

// // 	router.GET("/long_async", func(c *gin.Context) {
// // 		// create copy to be used inside the goroutine
// // 		cCp := c.Copy()
// // 		go func() {
// // 			// simulate a long task with time.Sleep(). 5 seconds
// // 			time.Sleep(5 * time.Second)

// // 			// note that you are using the copied context "cCp", IMPORTANT
// // 			log.Println("Done! in path " + cCp.Request.URL.Path)
// // 		}()
// // 		c.String(http.StatusOK, "Hello %s %s")

// // 	})

// // 	router.GET("/long_sync", func(c *gin.Context) {
// // 		// simulate a long task with time.Sleep(). 5 seconds
// // 		time.Sleep(5 * time.Second)

// // 		// since we are NOT using a goroutine, we do not have to copy the context
// // 		log.Println("Done! in path " + c.Request.URL.Path)
// // 	})

// // 	router.Run() // listen and serve on 0.0.0.0:8080
// // }

// package main

// import (
// 	"github.com/gin-gonic/gin"
// )

// var DB = make(map[string]string)

// func setupRouter() *gin.Engine {
// 	// Disable Console Color
// 	// gin.DisableConsoleColor()
// 	r := gin.Default()

// 	// Ping test
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.String(200, "pong")
// 	})

// 	// Get user value
// 	r.GET("/user/:name", func(c *gin.Context) {
// 		user := c.Params.ByName("name")
// 		value, ok := DB[user]
// 		if ok {
// 			c.JSON(200, gin.H{"user": user, "value": value})
// 		} else {
// 			c.JSON(200, gin.H{"user": user, "status": "no value"})
// 		}
// 	})

// 	// Authorized group (uses gin.BasicAuth() middleware)
// 	// Same than:
// 	// authorized := r.Group("/")
// 	// authorized.Use(gin.BasicAuth(gin.Credentials{
// 	//	  "foo":  "bar",
// 	//	  "manu": "123",
// 	//}))
// 	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
// 		"foo":  "bar", // user:foo password:bar
// 		"manu": "123", // user:manu password:123
// 	}))

// 	authorized.POST("admin", func(c *gin.Context) {
// 		user := c.MustGet(gin.AuthUserKey).(string)

// 		// Parse JSON
// 		var json struct {
// 			Value string `json:"value" binding:"required"`
// 		}

// 		if c.Bind(&json) == nil {
// 			DB[user] = json.Value
// 			c.JSON(200, gin.H{"status": "ok"})
// 		}
// 	})

// 	return r
// }

// func main() {
// 	r := setupRouter()
// 	// Listen and Server in 0.0.0.0:8080
// 	r.Run(":8080")
// }

package main

import (
	"fmt"
	"os"
)

var (
	HOME   = os.Getenv("HOME")
	USER   = os.Getenv("USER")
	GOROOT = os.Getenv("GOROOT")
	GOPATH = os.Getenv("GOPATH")
)

func main() {

	// fmt.Printf("HOME %s\n", HOME)
	// fmt.Printf("USER %s\n", USER)
	// fmt.Printf("GOROOT %s\n", GOROOT)
	// fmt.Printf("GOPATH %s\n", GOPATH)

	// var goos string = runtime.GOOS
	// fmt.Printf("The operating system is: %s\n", goos)
	// path := os.Getenv("PATH")
	// fmt.Printf("Path is %s\n", path)
	// var a = 2
	// var b int
	// b = a
	// var c *int

	// c = &b
	// fmt.Println(&a)
	// fmt.Println(&b)
	// fmt.Println(c)
	// var d int
	// d = *c
	// fmt.Println(d)
	// b = 1
	// fmt.Println(b)
	// fmt.Println(*c)
	// fmt.Println(d)

	// var a int
	// a = 1
	// var b *int
	// var c int
	// b = &a
	// c = a
	// a = 2
	// fmt.Println(a)
	// fmt.Println(*b)

	// fmt.Println(b)
	// fmt.Println(&a)
	// fmt.Println(&a == b)

	// fmt.Println(c)
	// var str string
	// str = "addfdsf"
	// var str1 *string
	// var str2 string
	// str1 = &str
	// str2 = *str1
	// str2 = "vvvv"
	// fmt.Println(str)
	// fmt.Println(str1)
	// fmt.Println(str2)
	// fmt.Println(*str1)
	// b := strings.Repeat(str, 8)
	// fmt.Println(b)
	// fmt.Println(str)
	// fmt.Println(*str1)
	// fmt.Println(str2)

	// var a = rand.Float64()
	// var b = int64(time.Now().Nanosecond())
	// fmt.Println(a)
	// fmt.Println(b)

	// var t string
	// t := time.Zone()
	// fmt.Println(t)
	var p *int

	fmt.Println(p)
}
