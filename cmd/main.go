package main

import (
	"fmt"
	"net/http"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/config"
	conversationdb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/conversationDB"
	matchhistdb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/matchHistDB"
	roledb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/roleDB"
	roomdb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/roomDB"
	userdb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/userDB"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/userHandler"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/helper/logger"
	userservice "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service/userService"
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

	db, err := initDB(cfg.Database)
	if err != nil {
		panic(err)
	}

	// Init repositories
	userDB := userdb.New(db)
	roomDB := roomdb.New(db)
	discountDB := roledb.New(db)
	matchHistDB := matchhistdb.New(db)
	conversationDB := conversationdb.New(db)

	// Init services
	userSvc := userservice.New(userDB)

	// Init handlers
	userHdl := userHandler.NewHTTPHandler(userSvc)
	// orderHdl := orderhandler.NewHTTPHandler(orderSrv)
	// discountHdl := discounthandler.NewHTTPHandler(discountSrv)

	// Starting server
	e := echo.New()
	handler.InitRoute(e, userHdl)

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
