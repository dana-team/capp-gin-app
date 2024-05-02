package server

import (
	"fmt"
	"net/http"
	"os"
	"syscall"

	"github.com/gin-gonic/gin"
)

const (
	envKey          = "ENV-TEST"
	envMessage      = " | env var: '%s'"
	echoMessage     = "echoing: %s%s"
	healthMessage   = "OK%s"
	errNoQueryValue = "No value provided to echo"
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
	env, ok := os.LookupEnv(envKey)
	message := ""
	if ok {
		message = fmt.Sprintf(envMessage, env)
	}

	if echoVar != "" {
		c.String(http.StatusOK, echoMessage, echoVar, message)
	} else {
		c.String(http.StatusBadRequest, errNoQueryValue)
	}
}

// handleHealth handles requests to the /health endpoint.
// It returns an OK response to indicate that the server is healthy.
func handleHealth(c *gin.Context) {
	env, ok := os.LookupEnv(envKey)
	message := ""
	if ok {
		message = fmt.Sprintf(envMessage, env)
	}
	c.String(http.StatusOK, healthMessage, message)
}
