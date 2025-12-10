package main

import (
	"backend-kurikulum-apps/config"
	"backend-kurikulum-apps/controller"
	"backend-kurikulum-apps/repository"
	"backend-kurikulum-apps/routes"
	"backend-kurikulum-apps/service"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found")
	}

	// Initialize database
	db := config.InitDB()

	// Initialize JWT config
	jwtConfig := config.LoadJWTConfig()

	// Initialize repositories
	userRepo := repository.NewUserRepository(db)
	refreshTokenRepo := repository.NewRefreshTokenRepository(db)
	cplRepo := repository.NewCPLRepository(db)
	mataKuliahRepo := repository.NewMataKuliahRepository(db)
	cplAssignmentRepo := repository.NewCPLAssignmentRepository(db)
	rpsRepo := repository.NewRPSRepository(db)
	rpsCPMKRepo := repository.NewRPSCPMKRepository(db)
	rpsRencanaRepo := repository.NewRPSRencanaPembelajaranRepository(db)
	rpsBahanRepo := repository.NewRPSBahanBacaanRepository(db)
	notificationRepo := repository.NewNotificationRepository(db)
	documentTemplateRepo := repository.NewDocumentTemplateRepository(db)
	generatedDocumentRepo := repository.NewGeneratedDocumentRepository(db)
	cplMKMappingRepo := repository.NewCPLMKMappingRepository(db)

	// New RPS Extended Repositories
	subCPMKRepo := repository.NewSubCPMKRepository(db)
	rencanaTugasRepo := repository.NewRPSRencanaTugasRepository(db)
	analisisKetercapaianRepo := repository.NewRPSAnalisisKetercapaianCPLRepository(db)

	// Initialize services
	authService := service.NewAuthService(userRepo, refreshTokenRepo)
	userService := service.NewUserService(userRepo)
	cplService := service.NewCPLService(cplRepo)
	mataKuliahService := service.NewMataKuliahService(mataKuliahRepo)
	notificationService := service.NewNotificationService(notificationRepo)
	cplAssignmentService := service.NewCPLAssignmentService(cplAssignmentRepo, cplRepo, cplMKMappingRepo, notificationService)
	rpsService := service.NewRPSService(rpsRepo, mataKuliahRepo, rpsCPMKRepo, rpsRencanaRepo, rpsBahanRepo, subCPMKRepo, notificationService)
	dashboardService := service.NewDashboardService(cplRepo, rpsRepo, cplAssignmentRepo, userRepo, generatedDocumentRepo)
	documentService := service.NewDocumentService(documentTemplateRepo, generatedDocumentRepo)
	cplMKMappingService := service.NewCPLMKMappingService(cplMKMappingRepo, cplRepo, mataKuliahRepo)

	// New RPS Extended Services
	subCPMKService := service.NewSubCPMKService(subCPMKRepo, rpsCPMKRepo)
	rencanaTugasService := service.NewRPSRencanaTugasService(rencanaTugasRepo, rpsRepo, subCPMKRepo)
	analisisKetercapaianService := service.NewRPSAnalisisKetercapaianCPLService(analisisKetercapaianRepo, rpsRepo)
	rpsExtendedService := service.NewRPSExtendedService(subCPMKService, rencanaTugasService, analisisKetercapaianService)

	// Initialize controllers
	uploadDir := getEnv("UPLOAD_DIR", "./uploads")
	authController := controller.NewAuthController(authService)
	userController := controller.NewUserController(userService)
	cplController := controller.NewCPLController(cplService)
	mataKuliahController := controller.NewMataKuliahController(mataKuliahService)
	cplAssignmentController := controller.NewCPLAssignmentController(cplAssignmentService)
	rpsController := controller.NewRPSController(rpsService)
	rpsExtendedController := controller.NewRPSExtendedController(rpsExtendedService)
	notificationController := controller.NewNotificationController(notificationService)
	dashboardController := controller.NewDashboardController(dashboardService)
	documentController := controller.NewDocumentController(documentService)
	fileController := controller.NewFileController(uploadDir)
	cplMKMappingController := controller.NewCPLMKMappingController(cplMKMappingService)

	// Bundle controllers
	controllers := &routes.Controllers{
		AuthController:          authController,
		UserController:          userController,
		CPLController:           cplController,
		MataKuliahController:    mataKuliahController,
		CPLAssignmentController: cplAssignmentController,
		RPSController:           rpsController,
		RPSExtendedController:   rpsExtendedController,
		NotificationController:  notificationController,
		DashboardController:     dashboardController,
		DocumentController:      documentController,
		FileController:          fileController,
		CPLMKMappingController:  cplMKMappingController,
	}

	allowedOrigins := strings.Split(getEnv("CORS_ALLOWED_ORIGINS", "*"), ",")

	// Setup router
	router := routes.SetupRouter(jwtConfig, controllers, allowedOrigins, uploadDir)

	// Get port from env
	port := getEnv("PORT", "8080")

	log.Printf("Server starting on port %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
