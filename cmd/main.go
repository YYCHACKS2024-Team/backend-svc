package main

import (
	"fmt"
	"net/http"

	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/config"
	conversationdb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/conversationDB"
	matchhistdb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/matchHistDB"
	messagedb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/messageDB"
	roledb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/roleDB"
	roomdb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/roomDB"
	userdb "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/adaptor/repositories/database/userDB"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler"
	chathandler "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/chatHandler"
	househandler "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/houseHandler"
	roommatehandler "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/roommateHandler"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/handler/userHandler"
	"github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/helper/logger"
	chatservice "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service/chatService"
	houseservice "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service/houseService"
	roommateservice "github.com/CLCM3102-Ice-Cream-Shop/backend-payment-service/internal/service/roommateService"
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
	roleDB := roledb.New(db)
	matchHistDB := matchhistdb.New(db)
	conversationDB := conversationdb.New(db)
	messageDB := messagedb.New(db)

	// Init services
	userSvc := userservice.New(userDB, roleDB)
	roommateSvc := roommateservice.New(userDB, roomDB, matchHistDB)
	houseSvc := houseservice.New(roomDB)
	chatSvc := chatservice.New(conversationDB, messageDB)

	// Init handlers
	userHdl := userHandler.NewHTTPHandler(userSvc)
	roommateHdl := roommatehandler.NewHTTPHandler(roommateSvc)
	houseHdl := househandler.NewHTTPHandler(houseSvc)
	chatHdl := chathandler.NewHTTPHandler(chatSvc)

	// Starting server
	e := echo.New()
	handler.InitRoute(e, userHdl, roommateHdl, houseHdl, chatHdl)

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
