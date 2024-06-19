package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func SSEDemoView(c *gin.Context) {

	var msgChan = make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			msgChan <- i
			time.Sleep(time.Second)
		}
		close(msgChan)
	}()

	c.Stream(func(w io.Writer) bool {
		if s, ok := <-msgChan; ok {
			c.SSEvent("", s)
			return true
		}
		return false
	})

}

func main() {
	r := gin.Default()
	r.GET("/sse", SSEDemoView)
	r.Run(":8081")
}
