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

	// Cart
	_ = e.Group("/cart")
	// cart.POST("", cartHdl.AddToCart)
	// cart.DELETE("/cancel/:cart_id/:no", cartHdl.Cancel)
	// cart.GET(("/customer/:customer_id"), cartHdl.GetCartByCustomerId)
	// cart.GET("/:cart_id", cartHdl.GetCartById)

	// Order
	// order := e.Group("/order")
	// order.POST("", orderHdl.Order)
	// order.GET("/customer/:customer_id", orderHdl.GetOrderByCustomerId)
	// order.GET("/date/:date", orderHdl.GetOrderByDate)
	// order.GET("/active", orderHdl.GetActiveOrders)
	// order.POST("/proceed/:order_id", orderHdl.ProceedOrder)
	// order.DELETE("/cancel/:order_id", orderHdl.CancelOrder)

	// // Discount
	// discount := e.Group("/discount")
	// discount.GET("/:code", discountHdl.GetDiscountByCode)
}
