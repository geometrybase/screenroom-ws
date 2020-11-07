package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strings"
	"sync"
)

var (
	ginMode    string
	ginPort    int
	wsUpGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	channels = make(map[string]*Channel)
	mu       sync.Mutex
)

func wsHandler(c *gin.Context) {
	w := c.Writer
	r := c.Request
	topicID := c.Param("topicID")
	mu.Lock()
	if _, ok := channels[topicID]; !ok {
		channels[topicID] = newChannel()
		go channels[topicID].run()
	}
	serveWs(channels[topicID], w, r)
	mu.Unlock()
}

func isWebsocket() gin.HandlerFunc {
	return func(c *gin.Context) {
		contains := func(key, val string) bool {
			vv := strings.Split(c.Request.Header.Get(key), ",")
			for _, v := range vv {
				if val == strings.ToLower(strings.TrimSpace(v)) {
					return true
				}
			}
			return false
		}
		if !contains("Connection", "upgrade") || !contains("Upgrade", "websocket") {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    http.StatusNotFound,
				"message": "NOT WEBSOCKET",
			})
			c.Abort()
			return
		}
	}
}

func init() {
	fmt.Printf("####  BUILD @ 20201107 04:21:47  ####\n")
}

func main() {
	flag.StringVar(&ginMode, "mode", "debug", "set gin mode")
	flag.IntVar(&ginPort, "port", 8000, "set gin listen port")
	flag.Parse()
	gin.SetMode(ginMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Any("/s/:topicID", isWebsocket(), wsHandler)
	r.Any("/m/:topicID", isWebsocket(), wsHandler)
	if err := http.ListenAndServe(
		fmt.Sprintf(":%d", ginPort),
		r,
	); err != nil {
		log.Fatal(err)
	}
}
