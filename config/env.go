package env

import (
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	AppPort           string
	JWTSecret         string
	JWTExpiresIn      time.Duration
	BasicAuthUsername string
	BasicAuthPassword string
	MongoDbUrl        string
	MongoDbName       string
	// Tambah config lain sesuai kebutuhan
}

var (
	Config *AppConfig
	once   sync.Once
)

// ConfigInit wajib dipanggil sekali (biasanya di main.go)
func ConfigInit() {
	once.Do(func() {
		// Load .env file
		err := godotenv.Load()
		if err != nil {
			log.Println("⚠️  No .env file found, using system env variables")
		}

		// Get and parse JWT_EXPIRES_IN
		jwtExpiresStr := getEnv("JWT_EXPIRES_IN", "15m")
		if !strings.HasSuffix(jwtExpiresStr, "m") &&
			!strings.HasSuffix(jwtExpiresStr, "h") &&
			!strings.HasSuffix(jwtExpiresStr, "s") {
			jwtExpiresStr += "m" // default ke menit
		}
		jwtExpiresDuration, err := time.ParseDuration(jwtExpiresStr)
		if err != nil {
			log.Printf("❌ Invalid JWT_EXPIRES_IN format (%s), fallback to 15m\n", jwtExpiresStr)
			jwtExpiresDuration = 15 * time.Minute
		}

		Config = &AppConfig{
			AppPort:           getEnv("PORT", "3000"),
			JWTSecret:         getEnv("JWT_SECRET", ""),
			JWTExpiresIn:      jwtExpiresDuration,
			BasicAuthUsername: getEnv("BASIC_AUTH_USERNAME", ""),
			BasicAuthPassword: getEnv("BASIC_AUTH_PASSWORD", ""),
			MongoDbUrl:        getEnv("MONGODB_URL", ""),
			MongoDbName:       getEnv("DB_NAME", "sighapp"),
		}

		if Config.JWTSecret == "" || Config.BasicAuthUsername == "" || Config.BasicAuthPassword == "" {
			log.Fatal("❌ Missing required environment variables (JWT_SECRET, BASIC_AUTH_USERNAME, BASIC_AUTH_PASSWORD)")
		}
	})
}

// Helper: ambil env, kalau kosong pakai default
func getEnv(key, fallback string) string {
	val := os.Getenv(key)
	if val == "" {
		return fallback
	}
	return val
}
