package main

import (
	"listing-service/internal/listing"
	"log"
	"os"
	"strconv"

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

	repo := listing.NewRepository(db)
	service := listing.NewService(repo)
	handler := listing.NewHandler(service)

	router := gin.Default()
	router.POST("/listings", handler.CreateListing)
	router.PUT("/listings/:id", handler.UpdateListing)
	router.DELETE("/listings/:id", handler.DeleteListing)
	router.GET("/listings/:id", handler.GetListing)
	router.GET("/listings", handler.GetAllListings)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
