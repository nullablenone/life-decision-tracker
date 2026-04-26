package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// Bikin map di memori buat nyimpen IP dan waktu terakhir mereka klik
var ipCache sync.Map

// IPRateLimit membatasi request per IP (1 request per 10 menit)
func IPRateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()

		lastSeen, exists := ipCache.Load(clientIP)
		if exists {
			// Kalau dia ngeklik lagi padahal belum lewat 10 menit, blokir!
			if time.Since(lastSeen.(time.Time)) < 60*time.Minute {
				c.String(http.StatusTooManyRequests, "Dont spam bruh! tunggu 1 jam")
				c.Abort()
				return
			}
		}

		// Kalau aman, catat waktu sekarang untuk IP ini
		ipCache.Store(clientIP, time.Now())
		c.Next()
	}
}
