package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"strings"
)

//CORS cross site resource sharing
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		var originHostName string
		//Domain name to run the request (can be run using localhost)
		var hosts = []string{"http://localhost", "https://localhost", "localhost"}
		origin := c.GetHeader("Origin")
		allowed := origin == ""
		if !allowed {
			originURL, err := url.Parse(origin)
			if err != nil {
				c.AbortWithStatus(http.StatusNoContent)
				return
			}
			originHostName = originURL.Hostname()

		}
		if !allowed {
			for _, host := range hosts {
				if strings.HasSuffix(origin, host) || strings.HasSuffix(originHostName, host) {
					allowed = true
					break
				}
			}
		}
		if !allowed {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, Authorization, identity-token, accept, origin, Cache-Control, X-Forwarded-For")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		c.Next()
	}
}
