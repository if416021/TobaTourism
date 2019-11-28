package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	_ "github.com/lib/pq"

	"github.com/TobaTourism/middleware"
	"github.com/TobaTourism/pkg/common/config"
	pariwisataDeliver "github.com/TobaTourism/pkg/delivery/pariwisata/http"
	"github.com/TobaTourism/pkg/models"
	pariwisataRepo "github.com/TobaTourism/pkg/repository/pariwisata/postgres"
	pariwisataUseCase "github.com/TobaTourism/pkg/usecase/pariwisata/module"

	experienceDeliver "github.com/TobaTourism/pkg/delivery/experience/http"
	experienceRepo "github.com/TobaTourism/pkg/repository/experience/postgres"
	experienceUseCase "github.com/TobaTourism/pkg/usecase/experience/module"

	restoDeliver "github.com/TobaTourism/pkg/delivery/resto/http"
	restoRepo "github.com/TobaTourism/pkg/repository/resto/postgres"
	restoUseCase "github.com/TobaTourism/pkg/usecase/resto/module"

	attachmentDeliver "github.com/TobaTourism/pkg/delivery/attachment/http"
	attachmentRepo "github.com/TobaTourism/pkg/repository/attachment/postgres"
	attachmentUseCase "github.com/TobaTourism/pkg/usecase/attachment/module"
)

var Conf *models.Config

func main() {
	Conf = config.InitConfig()
	//http
	e := echo.New()
	middL := middleware.InitMiddleware()
	e.Use(middL.CORS)

	//DB
	// db := conn.InitDB(Conf.Db.Conn)
	db, err := sql.Open("postgres", Conf.Db.Conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	// Start all services
	startService(e, db)

	restoran(e, db)
	attachment(e, db)

	log.Fatal(e.Start(":9090"))
}

func startService(e *echo.Echo, db *sql.DB) {
	pariwisataRepo := pariwisataRepo.InitPariwisataRepo(db)
	pariwisataUsecase := pariwisataUseCase.InitPariwisataUsecase(pariwisataRepo)
	pariwisataDeliver.InitPariwisataHandler(e, pariwisataUsecase)

	experienceRepo := experienceRepo.InitExperienceRepo(db)
	experienceUsecase := experienceUseCase.InitExperienceUsecase(experienceRepo)
	experienceDeliver.InitExperienceHandler(e, experienceUsecase)
}

func restoran(e *echo.Echo, db *sql.DB) {
	restoRepo := restoRepo.InitRestoRepo(db)
	attachmentRepo := attachmentRepo.InitAttachmentRepo(db)
	restoUsecase := restoUseCase.InitRestoUsecase(restoRepo)
	attachmentUseCase := attachmentUseCase.InitAttachmentUsecase(attachmentRepo)
	restoDeliver.InitRestoHandler(e, restoUsecase, attachmentUseCase)
}

func attachment(e *echo.Echo, db *sql.DB) {
	attachmentRepo := attachmentRepo.InitAttachmentRepo(db)
	attachmentUseCase := attachmentUseCase.InitAttachmentUsecase(attachmentRepo)
	attachmentDeliver.InitAttachmentHandler(e, attachmentUseCase)
}
