package handler

import (
	chathandler "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/chatHandler"
	househandler "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/houseHandler"
	roommatehandler "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/roommateHandler"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/userHandler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(e *echo.Echo, userHandler userHandler.HTTPHandler, roommateHandler roommatehandler.HTTPHandler,
	houseHandler househandler.HTTPHandler, chatHandler chathandler.HTTPHandler) {

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
	user.GET("/users", userHandler.GetAll)              // get all user
	user.GET("/user/:user_id", userHandler.GetUserById) // get user by id
	user.POST("/user/register", userHandler.Register)   // create user
	user.POST("/user/login", userHandler.Login)         // login
	user.GET("/user/profile", userHandler.GetProfile)   // get user profile

	user.POST("/user/preferences", userHandler.StorePreferences) // store preference

	// roommate
	roommate := e.Group("/roommate")
	roommate.POST("/roommate/accept", roommateHandler.Accept)
	roommate.POST("/roommate/decline", roommateHandler.Decine)

	// chat
	chat := e.Group("/chat")
	chat.POST("/chat/store", chatHandler.StoreChatHistory)
	chat.GET("/chat/:conversation_id", chatHandler.GetChatByConversationId)        // get conversation room by id
	chat.GET("/chat/:conversation_id/:message_id", chatHandler.GetChatByMessageId) // get message id

	// house
	house := e.Group("/house")
	house.GET("/house/:house_id", houseHandler.GetHouseById) // get house detail by id
}
