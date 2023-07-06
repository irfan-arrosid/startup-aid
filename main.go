package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/irfan-arrosid/startup-aid/auth"
	"github.com/irfan-arrosid/startup-aid/campaign"
	"github.com/irfan-arrosid/startup-aid/handler"
	"github.com/irfan-arrosid/startup-aid/helper"
	"github.com/irfan-arrosid/startup-aid/payment"
	"github.com/irfan-arrosid/startup-aid/transaction"
	"github.com/irfan-arrosid/startup-aid/user"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load godotenv
	godotenv.Load()

	// Connecting to database
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	} else {
		fmt.Println("Database connected....")
	}

	// Import Repository
	userRepository := user.NewRepository(db)
	campaignRepository := campaign.NewRepository(db)
	transactionRepository := transaction.NewRepository(db)

	// Import Service
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	paymentService := payment.NewService()
	campaignService := campaign.NewService(campaignRepository)
	transactionService := transaction.NewService(transactionRepository, campaignRepository, paymentService)

	// Import Handler
	userHandler := handler.NewUserHandler(userService, authService)
	campaignHandler := handler.NewCampaignHandler(campaignService)
	transactionHandler := handler.NewTransactionHandler(transactionService)

	// Init Route
	r := gin.Default()
	r.Static("/images", "./images")
	api := r.Group("/api/v1")

	// List of USER endpoints
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddelware(authService, userService), userHandler.UploadAvatar)

	// List of CAMPAIGN endpoints
	api.GET("/campaigns", campaignHandler.GetCampaigns)
	api.GET("/campaigns/:id", campaignHandler.GetCampaign)
	api.POST("/campaigns", authMiddelware(authService, userService), campaignHandler.CreateCampaign)
	api.PUT("/campaigns/:id", authMiddelware(authService, userService), campaignHandler.UpdateCampaign)
	api.POST("/campaign-images", authMiddelware(authService, userService), campaignHandler.UploadImage)

	// List of TRANSACTION endpoints
	api.GET("/campaigns/:id/transactions", authMiddelware(authService, userService), transactionHandler.GetCampaignTransactions)
	api.GET("/transactions", authMiddelware(authService, userService), transactionHandler.GetUserTransactions)
	api.POST("/transactions", authMiddelware(authService, userService), transactionHandler.CreateTransaction)

	r.Run()
}

func authMiddelware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claim, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userId := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userId)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
