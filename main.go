package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	m := melody.New()

	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		msg = []byte("this is test33333" +s.Request.URL.Query().Get("key") )
		fmt.Println(msg)
		m.Broadcast(msg)
		fmt.Println("this is test")

		time.Sleep(1 * time.Second)
		m.Broadcast(msg)
		fmt.Println("this is test1")
		time.Sleep(2 * time.Second)
		m.Broadcast(msg)
		fmt.Println("this is test2")
	})

	r.Run(":5000")
}
