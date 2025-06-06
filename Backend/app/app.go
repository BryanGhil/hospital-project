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
	a.Router.Use(middleware.CORSMiddleware())
	a.Router.Use(middleware.ErrorMiddleware())

	a.initRoutes()
}

func (a *App) initRoutes() {
	trx := repository.NewTransactor(a.DB)

	ur := repository.NewUserRepo()
	uuc := usecase.NewUserUsecaseImpl(ur, trx)
	uh := handler.NewUserHandler(uuc)

	pr := repository.NewPatientRepo()
	puc := usecase.NewPatientUsecaseImpl(pr, trx)
	ph := handler.NewPatientHandler(puc)

	mr := repository.NewMedicineRepo()
	muc := usecase.NewMedicineUsecaseImpl(mr, trx)
	mh := handler.NewMedicineHandler(muc)

	a.Router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := a.Router.Group("/api/v1")
	{
		v1.POST("/register", uh.RegisterUser)
		v1.POST("/login", uh.LoginUser)
		patients := v1.Group("/patients", middleware.Authenticate()) 
		{
			patients.POST("", ph.AddPatient)
			patients.GET("", ph.GetAllPatients)
			patients.GET("/:id", ph.GetPatientById)
			patients.PATCH("/:id", ph.UpdatePatients)
			patients.DELETE("/:id", ph.DeletePatient)
			patients.PATCH("/:id/restore", ph.RestoreDeletedPatient)
		}
		medicines := v1.Group("/medicines", middleware.Authenticate()) 
		{
			medicines.POST("", mh.AddMedicine)
		}
	}

}

func (a *App) Run() {
	a.Init()

	defer a.DB.Close()

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
