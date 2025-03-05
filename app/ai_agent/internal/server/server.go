// Description: This file is used to create a new echo server instance.

package server

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"golang.org/x/time/rate"
)

// type Server struct {
// 	port int

// 	// db database.Service
// }

// func NewServer() *http.Server {
// 	port, _ := strconv.Atoi(os.Getenv("PORT"))
// 	NewServer := &Server{
// 		port: port,

// 		// db: database.New(),
// 	}

// 	// Declare Server config
// 	server := &http.Server{
// 		Addr:         fmt.Sprintf(":%d", NewServer.port),
// 		Handler:      NewServer.RegisterRoutes(),
// 		IdleTimeout:  time.Minute,
// 		ReadTimeout:  10 * time.Second,
// 		WriteTimeout: 30 * time.Second,
// 	}

// 	return server
// }

func NewEchoServer() *echo.Echo{
	e := echo.New()
	e.Logger.SetLevel(log.INFO)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Configure the rate limiter middleware
	// 在配置的到期时间ExpiresIn过后, 通过删除未再次访问的用户ip的陈旧记录，帮助管理访客map的大小
	e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper:          middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: rate.Limit(10000), Burst: 30000, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(c echo.Context) (string, error) {
			id := c.RealIP()
			return id, nil
		},
	
		ErrorHandler: func(c echo.Context, err error) error {
			return middleware.ErrExtractorError
		},
		DenyHandler: func(c echo.Context, identifier string, err error) error {
			return middleware.ErrRateLimitExceeded
		},
	}))
	return e
}
