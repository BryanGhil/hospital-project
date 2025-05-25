package app

import (
	"backend/constant"
	"backend/db"
	"backend/handler"
	"backend/middleware"
	"backend/repository"
	"backend/usecase"
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type App struct {
	Router *gin.Engine
	DB     *sql.DB
}

func (a *App) Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("error loading .env file", err)
	}

	a.DB, err = db.ConnectDB()
	if err != nil {
		log.Fatalf("error connect DB: %s\n", err)
	}

	a.Router = gin.Default()
	a.Router.ContextWithFallback = true
	a.Router.Use(middleware.ErrorMiddleware())

	a.initRoutes()
}

func (a *App) initRoutes() {
	trx := repository.NewTransactor(a.DB)

	ur := repository.NewUserRepo(a.DB)
	uuc := usecase.NewUserUsecaseImpl(ur, trx)
	uh := handler.NewUserHandler(uuc)

	a.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	a.Router.POST("/register", uh.RegisterUser)
	a.Router.POST("/login", uh.LoginUser)
}

func (a *App) Run() {
	srv := &http.Server{
		Addr:    os.Getenv(constant.ServerPort),
		Handler: a.Router.Handler(),
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server shutdown: ", err)
	}

	<-ctx.Done()
	log.Println("Timeout of 5 seconds")
	log.Println("Server exiting")
}
