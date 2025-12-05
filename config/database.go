package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

type Config struct {
	AppName     string
	AppEnv      string
	AppPort     string
	AppDebug    bool
	DBHost      string
	DBPort      string
	DBName      string
	DBUser      string
	DBPassword  string
	DBCharset   string
	JWTSecret   string
	JWTExpires  time.Duration
	JWTRefresh  time.Duration
	UploadPath  string
	MaxFileSize int64
	CORSOrigins string
}

var AppConfig *Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}

	jwtExpires, _ := time.ParseDuration(getEnv("JWT_EXPIRES_IN", "24h"))
	jwtRefresh, _ := time.ParseDuration(getEnv("JWT_REFRESH_EXPIRES_IN", "168h"))

	AppConfig = &Config{
		AppName:     getEnv("APP_NAME", "Kurikulum Management System"),
		AppEnv:      getEnv("APP_ENV", "development"),
		AppPort:     getEnv("APP_PORT", "8080"),
		AppDebug:    getEnv("APP_DEBUG", "true") == "true",
		DBHost:      getEnv("DB_HOST", "localhost"),
		DBPort:      getEnv("DB_PORT", "3306"),
		DBName:      getEnv("DB_NAME", "kurikulum_db"),
		DBUser:      getEnv("DB_USER", "root"),
		DBPassword:  getEnv("DB_PASSWORD", ""),
		DBCharset:   getEnv("DB_CHARSET", "utf8mb4"),
		JWTSecret:   getEnv("JWT_SECRET", "your-secret-key"),
		JWTExpires:  jwtExpires,
		JWTRefresh:  jwtRefresh,
		UploadPath:  getEnv("UPLOAD_PATH", "./uploads"),
		MaxFileSize: 10485760,
		CORSOrigins: getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:3000"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func ConnectDatabase() {
	var err error

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		AppConfig.DBUser,
		AppConfig.DBPassword,
		AppConfig.DBHost,
		AppConfig.DBPort,
		AppConfig.DBName,
		AppConfig.DBCharset,
	)

	logMode := logger.Silent
	if AppConfig.AppDebug {
		logMode = logger.Info
	}

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logMode),
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	// Connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("Database connected successfully")
}

// InitDB initializes and returns database connection
func InitDB() *gorm.DB {
	LoadConfig()
	ConnectDatabase()
	return DB
}
