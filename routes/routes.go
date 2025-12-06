package routes

import (
	"backend-kurikulum-apps/config"
	"backend-kurikulum-apps/controller"
	"backend-kurikulum-apps/middleware"
	"time"

	"github.com/gin-gonic/gin"
)

type Controllers struct {
	AuthController          *controller.AuthController
	UserController          *controller.UserController
	CPLController           *controller.CPLController
	MataKuliahController    *controller.MataKuliahController
	CPLAssignmentController *controller.CPLAssignmentController
	RPSController           *controller.RPSController
	NotificationController  *controller.NotificationController
	DashboardController     *controller.DashboardController
	DocumentController      *controller.DocumentController
	FileController          *controller.FileController
	CPLMKMappingController  *controller.CPLMKMappingController
}

func SetupRouter(
	jwtConfig *config.JWTConfig,
	controllers *Controllers,
	allowedOrigins []string,
	uploadDir string,
) *gin.Engine {
	router := gin.New()

	// Global middlewares
	router.Use(middleware.RecoveryMiddleware())
	router.Use(middleware.LoggerMiddleware())
	router.Use(middleware.RequestIDMiddleware())
	router.Use(middleware.CORSMiddleware(allowedOrigins))
	router.Use(middleware.RateLimitMiddleware(100, time.Minute)) // 100 requests per minute

	// Serve static files
	router.Static("/uploads", uploadDir)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{

		// Auth routes (public)
		auth := v1.Group("/auth")
		{
			auth.POST("/login", controllers.AuthController.Login)
			auth.POST("/register", controllers.AuthController.Register)
			auth.POST("/refresh", controllers.AuthController.RefreshToken)
		}

		// Auth routes (protected)
		authProtected := v1.Group("/auth")
		authProtected.Use(middleware.AuthMiddleware(jwtConfig))
		{
			authProtected.POST("/logout", controllers.AuthController.Logout)
			authProtected.POST("/change-password", controllers.AuthController.ChangePassword)
			authProtected.GET("/profile", controllers.AuthController.GetProfile)
			authProtected.PUT("/profile", controllers.AuthController.UpdateProfile)
		}

		// Protected routes
		protected := v1.Group("")
		protected.Use(middleware.AuthMiddleware(jwtConfig))
		{

			dashboard := protected.Group("/dashboard")
			{
				dashboard.GET("", controllers.DashboardController.GetMyDashboard)
				dashboard.GET("/kaprodi", middleware.KaprodiOnly(), controllers.DashboardController.GetKaprodiDashboard)
				dashboard.GET("/dosen", controllers.DashboardController.GetDosenDashboard)
			}

			// User routes (Kaprodi only for CUD)
			users := protected.Group("/users")
			{
				users.GET("", controllers.UserController.GetAllUsers)
				users.GET("/dosen", controllers.UserController.GetDosen)
				users.GET("/:id", controllers.UserController.GetUserByID)
				users.POST("", middleware.KaprodiOnly(), controllers.UserController.CreateUser)
				users.PUT("/:id", middleware.KaprodiOnly(), controllers.UserController.UpdateUser)
				users.DELETE("/:id", middleware.KaprodiOnly(), controllers.UserController.DeleteUser)
				users.PATCH("/:id/toggle-status", middleware.KaprodiOnly(), controllers.UserController.ToggleUserStatus)
			}

			// CPL routes (Kaprodi only for CUD)
			cpl := protected.Group("/cpl")
			{
				cpl.GET("", controllers.CPLController.GetAllCPL)
				cpl.GET("/statistics", controllers.CPLController.GetCPLStatistics)
				cpl.GET("/active", controllers.CPLController.GetActiveCPL)
				cpl.GET("/:id", controllers.CPLController.GetCPLByID)
				cpl.POST("", middleware.KaprodiOnly(), controllers.CPLController.CreateCPL)
				cpl.PUT("/:id", middleware.KaprodiOnly(), controllers.CPLController.UpdateCPL)
				cpl.DELETE("/:id", middleware.KaprodiOnly(), controllers.CPLController.DeleteCPL)
				cpl.PATCH("/:id/status", middleware.KaprodiOnly(), controllers.CPLController.UpdateCPLStatus)
			}

			// Mata Kuliah routes
			mk := protected.Group("/mata-kuliah")
			{
				mk.GET("", controllers.MataKuliahController.GetAllMataKuliah)
				mk.GET("/my", controllers.MataKuliahController.GetMyMataKuliah)
				mk.GET("/semester/:semester", controllers.MataKuliahController.GetMataKuliahBySemester)
				mk.GET("/dosen/:dosen_id", controllers.MataKuliahController.GetMataKuliahByDosen)
				mk.GET("/:id", controllers.MataKuliahController.GetMataKuliahByID)
				mk.POST("", middleware.KaprodiOnly(), controllers.MataKuliahController.CreateMataKuliah)
				mk.PUT("/:id", middleware.KaprodiOnly(), controllers.MataKuliahController.UpdateMataKuliah)
				mk.DELETE("/:id", middleware.KaprodiOnly(), controllers.MataKuliahController.DeleteMataKuliah)
				mk.PATCH("/:id/toggle-status", middleware.KaprodiOnly(), controllers.MataKuliahController.ToggleMataKuliahStatus)
				mk.PATCH("/:id/assign-dosen", middleware.KaprodiOnly(), controllers.MataKuliahController.AssignDosen)
				mk.PATCH("/:id/unassign-dosen", middleware.KaprodiOnly(), controllers.MataKuliahController.UnassignDosen)
			}

			// CPL Assignment routes
			assignment := protected.Group("/cpl-assignments")
			{
				assignment.GET("", controllers.CPLAssignmentController.GetAllAssignments)
				assignment.GET("/my", controllers.CPLAssignmentController.GetMyAssignments)
				assignment.GET("/cpl/:cpl_id", controllers.CPLAssignmentController.GetAssignmentsByCPL)
				assignment.GET("/:id", controllers.CPLAssignmentController.GetAssignmentByID)
				assignment.POST("", middleware.KaprodiOnly(), controllers.CPLAssignmentController.CreateAssignment)
				assignment.DELETE("/:id", middleware.KaprodiOnly(), controllers.CPLAssignmentController.DeleteAssignment)
				assignment.PATCH("/:id/status", controllers.CPLAssignmentController.UpdateAssignmentStatus)
			}

			// CPL-MK Mapping routes
			cplMKMapping := protected.Group("/cpl-mk-mappings")
			{
				cplMKMapping.GET("", controllers.CPLMKMappingController.GetAllMappings)
				cplMKMapping.GET("/:id", controllers.CPLMKMappingController.GetMappingByID)
				cplMKMapping.POST("/upsert", middleware.KaprodiOnly(), controllers.CPLMKMappingController.UpsertMapping)
				cplMKMapping.DELETE("/:id", middleware.KaprodiOnly(), controllers.CPLMKMappingController.DeleteMapping)
			}

			// RPS routes
			rps := protected.Group("/rps")
			{
				rps.GET("", controllers.RPSController.GetAllRPS)
				rps.GET("/my", controllers.RPSController.GetMyRPS)
				rps.GET("/mata-kuliah/:mata_kuliah_id", controllers.RPSController.GetRPSByMataKuliah)
				rps.GET("/:id", controllers.RPSController.GetRPSByID)
				rps.POST("", controllers.RPSController.CreateRPS)
				rps.PUT("/:id", controllers.RPSController.UpdateRPS)
				rps.DELETE("/:id", controllers.RPSController.DeleteRPS)
				rps.PATCH("/:id/submit", controllers.RPSController.SubmitRPS)
				rps.PATCH("/:id/approve", middleware.KaprodiOnly(), controllers.RPSController.ApproveRPS)
				rps.PATCH("/:id/reject", middleware.KaprodiOnly(), controllers.RPSController.RejectRPS)
				rps.PATCH("/:id/request-revision", middleware.KaprodiOnly(), controllers.RPSController.RequestRevision)

				// CPMK sub-routes
				rps.POST("/:rps_id/cpmk", controllers.RPSController.AddCPMK)
				rps.PUT("/cpmk/:cpmk_id", controllers.RPSController.UpdateCPMK)
				rps.DELETE("/cpmk/:cpmk_id", controllers.RPSController.DeleteCPMK)

				// Rencana Pembelajaran sub-routes
				rps.POST("/:rps_id/rencana-pembelajaran", controllers.RPSController.AddRencanaPembelajaran)
				rps.PUT("/rencana-pembelajaran/:rencana_id", controllers.RPSController.UpdateRencanaPembelajaran)
				rps.DELETE("/rencana-pembelajaran/:rencana_id", controllers.RPSController.DeleteRencanaPembelajaran)

				// Bahan Bacaan sub-routes
				rps.POST("/:rps_id/bahan-bacaan", controllers.RPSController.AddBahanBacaan)
				rps.PUT("/bahan-bacaan/:bahan_id", controllers.RPSController.UpdateBahanBacaan)
				rps.DELETE("/bahan-bacaan/:bahan_id", controllers.RPSController.DeleteBahanBacaan)

				// Evaluasi sub-routes
				rps.POST("/:rps_id/evaluasi", controllers.RPSController.AddEvaluasi)
				rps.PUT("/evaluasi/:evaluasi_id", controllers.RPSController.UpdateEvaluasi)
				rps.DELETE("/evaluasi/:evaluasi_id", controllers.RPSController.DeleteEvaluasi)
			}

			// Notification routes
			notifications := protected.Group("/notifications")
			{
				notifications.GET("", controllers.NotificationController.GetMyNotifications)
				notifications.GET("/unread-count", controllers.NotificationController.GetUnreadCount)
				notifications.POST("", middleware.KaprodiOnly(), controllers.NotificationController.CreateNotification)
				notifications.PATCH("/:id/read", controllers.NotificationController.MarkAsRead)
				notifications.PATCH("/read-all", controllers.NotificationController.MarkAllAsRead)
				notifications.DELETE("/:id", controllers.NotificationController.DeleteNotification)
			}

			// Document routes
			documents := protected.Group("/documents")
			{
				documents.GET("", controllers.DocumentController.GetAllDocuments)
				documents.GET("/:id", controllers.DocumentController.GetDocumentByID)
				documents.POST("/generate", controllers.DocumentController.GenerateDocument)
				documents.DELETE("/:id", controllers.DocumentController.DeleteDocument)

				// Template routes (Kaprodi only for CUD)
				documents.GET("/templates", controllers.DocumentController.GetAllTemplates)
				documents.GET("/templates/:id", controllers.DocumentController.GetTemplateByID)
				documents.POST("/templates", middleware.KaprodiOnly(), controllers.DocumentController.CreateTemplate)
				documents.PUT("/templates/:id", middleware.KaprodiOnly(), controllers.DocumentController.UpdateTemplate)
				documents.DELETE("/templates/:id", middleware.KaprodiOnly(), controllers.DocumentController.DeleteTemplate)
			}

			// File routes
			files := protected.Group("/files")
			{
				files.POST("/upload", controllers.FileController.UploadFile)
				files.GET("/info", controllers.FileController.GetFileInfo)
				files.DELETE("", controllers.FileController.DeleteFile)
			}
		}
	}

	return router
}
