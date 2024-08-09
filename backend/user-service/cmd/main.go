package main

import (
	"log"
	"os"
	"strconv"
	"user-service/internal/user"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	JWT struct {
		Secret     string `yaml:"secret"`
		Expiration string `yaml:"expiration"`
	} `yaml:"jwt"`
}

func loadConfig() (*Config, error) {
	var config Config
	data, err := os.ReadFile("configs/config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(data, &config)
	return &config, err
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Could not load config: %v", err)
	}

	dsn := "host=" + config.Database.Host + " user=" + config.Database.User + " password=" + config.Database.Password + " dbname=" + config.Database.Name + " port=" + strconv.Itoa(config.Database.Port) + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo, config.JWT.Secret)
	userHandler := user.NewHandler(userService)

	r := gin.Default()
	r.POST("/register", userHandler.Register)
	r.POST("/login", userHandler.Authenticate)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
