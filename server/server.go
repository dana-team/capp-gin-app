package server

import (
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
)

// setupRouter returns a gin router.
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/echo", handleEcho)
	router.GET("/health", handleHealth)

	return router
}

// RunServer starts the HTTP server on port 8080 and handles incoming requests.
func RunServer(quit chan os.Signal) {
	router := setupRouter()

	if err := router.Run(":8080"); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		quit <- syscall.SIGTERM
	}
}

// handleEcho handles requests to the /echo endpoint.
// It echoes back the value of the query parameter "var" if provided,
// or returns a Bad Request response if no value is provided.
func handleEcho(c *gin.Context) {
	echoVar := c.Query("var")
	if echoVar != "" {
		c.String(http.StatusOK, "echoing: %s", echoVar)
	} else {
		c.String(http.StatusBadRequest, "No value provided to echo")
	}
}

// handleHealth handles requests to the /health endpoint.
// It returns an OK response to indicate that the server is healthy.
func handleHealth(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
