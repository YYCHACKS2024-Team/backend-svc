package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo) {

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowOrigins: []string{"*"}, // Replace with your frontend origin(s)
			AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		}),
	)

	// User
	user := e.Group("/user")
	user.GET("/users", nil)          // get all user
	user.GET("/user/:user_id", nil)  // get user by id
	user.POST("/user/register", nil) // create user
	user.POST("/user/login", nil)    // login
	user.GET("/user/profile", nil)   // get user profile

	user.POST("/user/preferences", nil) // store preference

	// roommate
	room := e.Group("/roommate")
	room.GET("/roommate/:room_id", nil) // get room list
	room.POST("/roommate/accept", nil)
	room.POST("/roommate/decline", nil)

	// chat
	chat := e.Group("/chat")
	chat.POST("/chat/store", nil)
	chat.GET("/chat/:conversation_id", nil)             // get conversation room by id
	chat.GET("/chat/:conversation_id/:message_id", nil) // get message id

	// house
	house := e.Group("/house")
	house.GET("/house/:house_id", nil) // get house detail by id
}
