package handler

import (
	"canteen-prakerja/database"
	"canteen-prakerja/repository/barang_repository/barang_my"
	"canteen-prakerja/repository/report_repository/report_my"
	"canteen-prakerja/repository/transaksi_repository/transaksi_my"
	"canteen-prakerja/repository/user_repository/user_my"
	"canteen-prakerja/service"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	_ "github.com/tbxark/g4vercel"
)

func StartApp() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Panicf("error while getting .env file : %s", err.Error())
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	database.InitializeDatabase()

	db := database.GetDatabaseInstance()

	userRepo := user_my.NewUserMy(db)
	userService := service.NewUserService(userRepo)
	userHandler := NewUserHandler(userService)

	barangRepo := barang_my.NewBarangMy(db)
	barangService := service.NewBarangService(barangRepo)
	barangHandler := NewBarangHandler(barangService)

	transaksiRepo := transaksi_my.NewTransaksiMy(db)
	transaksiService := service.NewTransaksiService(transaksiRepo)
	transaksiHandler := NewTransaksiHandler(transaksiService)

	reportRepo := report_my.NewReportMy(db)
	reportService := service.NewReportService(reportRepo)
	reportHandler := NewReportHandler(reportService)

	gin.SetMode(gin.ReleaseMode)
	route := gin.New()
	route.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:    []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
	}))

	loginRoute := route.Group("/login")
	{
		loginRoute.POST("/", userHandler.Login)
	}

	barangRoute := route.Group("/barangs")
	{
		barangRoute.GET("/", barangHandler.GetAllBarang)
		barangRoute.GET("/:barangId", barangHandler.GetBarangById)
		barangRoute.POST("/", barangHandler.CreateNewBarang)
		barangRoute.PUT("/:barangId", barangHandler.UpdateBarangById)
		barangRoute.DELETE("/:barangId", barangHandler.DeleteBarangById)
	}

	transaksiRoute := route.Group("/txs")
	{
		transaksiRoute.GET("/", transaksiHandler.GetAllTransaksi)
		transaksiRoute.POST("/", transaksiHandler.CreateTransaksi)
		transaksiRoute.POST("/date", transaksiHandler.GetTransaksiDateBetween)
		transaksiRoute.GET("/:txId", transaksiHandler.GetTransaksiById)
		transaksiRoute.DELETE("/:txId", transaksiHandler.DeleteTransaksiById)
	}

	reportRoute := route.Group("/report")
	{
		reportRoute.POST("/date", reportHandler.GetReportDate)
		reportRoute.POST("/range", reportHandler.GetReportDateBetween)
	}

	route.POST("/", helloWorld)

	route.Run(":" + port)
}

func helloWorld(c *gin.Context) {
	c.JSON(200, "Hello World")
}
