package main

import (
	"fmt"
	"net/http"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/config"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/helper/logger"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}

	logger.InitLog(cfg.Log)
	defer logger.CloseLogger()

	_, err = initDB(cfg.Database)
	if err != nil {
		panic(err)
	}

	// Init repositories
	// cartDB := cartdb.New(db)
	// orderDB := orderdb.New(db)
	// discountDB := discountdb.New(db)

	// Init gateway

	// Init services
	// cartSvc := cartservice.New(cartDB, productServiceGw)

	// Init handlers
	// cartHdl := carthandler.NewHTTPHandler(cartSvc)
	// orderHdl := orderhandler.NewHTTPHandler(orderSrv)
	// discountHdl := discounthandler.NewHTTPHandler(discountSrv)

	// Starting server
	e := echo.New()
	handler.InitRoute(e)

	logger.Infof("Starting server on port %v...\n", cfg.App.Port)
	if err := e.Start(":" + cfg.App.Port); err != http.ErrServerClosed {
		logger.Fatal(err.Error())
	}
}

func initDB(dbCfg config.Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.Username,
		dbCfg.Password,
		dbCfg.Host,
		dbCfg.Port,
		dbCfg.DBName,
	)

	logger.Infof("Connecting database... %v", dsn)

	gormDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	logger.Info("Connect database successfully!!")

	return gormDB, err
}
