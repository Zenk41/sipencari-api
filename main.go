package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	_dbDriver "sipencari-api/drivers/mysql"
	_middleware "sipencari-api/app/middlewares"
	_driverFactory "sipencari-api/drivers"
	_userUsecase "sipencari-api/businesses/users"
	_userController "sipencari-api/controllers/users"
	_routes "sipencari-api/app/routes"
	util "sipencari-api/utils"
	"sync"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
)

// operation is a clean up function on shutting down
type operation func(ctx context.Context) error

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_USERNAME: util.GetEnv("DB_USERNAME"),
		DB_PASSWORD: util.GetEnv("DB_PASSWORD"),
		DB_HOST:     util.GetEnv("DB_HOST"),
		DB_PORT:     util.GetEnv("DB_PORT"),
		DB_NAME:     util.GetEnv("DB_NAME"),
	}

	db := configDB.InitDB()
	_dbDriver.DBMigrate(db)

	configJWT := _middleware.ConfigJWT{
		SecretJWT:      util.GetEnv("JWT_SECRET_KEY"),
		ExpireDuration: 24,
	}

	configLogger := _middleware.ConfigLogger{
		Format: "[${time_rfc3339}] method=${method}, uri=${uri}, status=${status}, latency_human=${latency_human}\n",
	}

	e := echo.New()

	userRepo := _driverFactory.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo, &configJWT)
	userCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.ControllerList{
		LoggerMiddleware: configLogger.Init(),
		JWTMIddleware: configJWT.Init(),
		AuthController: *userCtrl,
	}
	routesInit.RouteRegister(e)

	go func() {
		if err := e.Start(":" + util.GetEnv("PORT")); err != nil && err != http.ErrServerClosed{
			e.Logger.Fatal("shutting down server")
		}
	}()
	wait := gracefulShutdown(context.Background(), 2 * time.Second, map[string]operation {
		"database" : func(ctx context.Context) error {
			return _dbDriver.CloseDB(db)
		},
		"http-server" : func(ctx context.Context) error {
			return e.Shutdown(context.Background())
		},
	})

	<-wait

}

// gracefulShutdown performs gracefully shutdown 
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("shutting down")

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("timeout %d ms has been elased, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("cleaning up: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: clean up failed : %s", innerKey, err.Error())
					return
				}

				log.Printf("%s was shutdown gracefully", innerKey)
			}()
		}
		wg.Wait()

		close(wait)
	}()

	return wait
}
