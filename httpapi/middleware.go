package httpapi

import (
	"git.containerum.net/ch/resource-service/server"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"
)

var logger = logrus.New()

func initializeContext(srv server.ResourceSvcInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := uuid.NewV4().String()
		c.Header("x-request-id", requestID)
		c.Set("request-id", requestID)
		c.Set("logger", logrus.NewEntry(logger).
			WithField("client-ip", c.ClientIP()).
			WithField("request-id", requestID).
			WithField("http-method", c.Request.Method).
			WithField("http-uri", c.Request.RequestURI))
		c.Set("server", srv)
	}
}

// depends on initializeContext
func parseHeaders(c *gin.Context) {
	userID := c.Request.Header.Get("x-user-id")
	userRole := c.Request.Header.Get("x-user-role")
	tokenID := c.Request.Header.Get("x-user-token-id")

	c.Set("user-id", userID)
	c.Set("user-role", userRole)
	c.Set("token-id", tokenID)

	logger := c.MustGet("logger").(*logrus.Entry)
	logger = logger.
		WithField("user-id", userID).
		WithField("actor-user-id", userID).
		WithField("user-role", userRole).
		WithField("token-id", tokenID)
	c.Set("logger", logger)
}

// adminAction checks whether the request is performed by the
// ‘admin’ account. If so, it substitutes user ID from query
// parameters, if present, and sets "admin-action" context
// field.
//
// depends on parseHeaders
func adminAction(c *gin.Context) {
	logger := c.MustGet("logger").(*logrus.Entry)

	if c.MustGet("user-role").(string) == "admin" {
		if qpUserID, exists := c.GetQuery("user-id"); exists {
			c.Set("user-id", qpUserID)
			logger = logger.WithField("user-id", qpUserID)
		}
		c.Set("admin-action", true)
		logger = logger.WithField("admin-action", true)
	} else {
		c.Set("admin-action", false)
		logger = logger.WithField("admin-action", false)
	}

	c.Set("logger", logger)
}
