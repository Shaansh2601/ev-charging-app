package handler

import (
	"backend/pkg/service"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes(ginMode string) *gin.Engine {
	gin.SetMode(ginMode)
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8000"}, // Adjust to your frontend domains
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.Static("/static", os.Getenv("STATIC_DIR"))

	router.POST("/static", func(c *gin.Context) {
		var jsonData map[string]interface{}
		if err := c.ShouldBindJSON(&jsonData); err != nil {
			c.JSON(400, gin.H{"error": "Invalid JSON"})
			return
		}
		// Process the JSON payload
		c.Set("cost", jsonData["cost"])
		c.JSON(200, gin.H{"message": "JSON received", "data": jsonData})
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.GET("/:provider/login", h.signInWithProvider)
		auth.GET("/:provider/callback", h.callbackHandler)
	}

	myPage := router.Group("/home")
	{
		myPage.GET("/:username", h.userIdentity)
	}

	payPage := router.Group("/payment")
	{
		payPage.GET("/config", h.handleConfig)
		payPage.POST("/create-payment-intent", h.createPayment)
		payPage.POST("/webhook", h.handleWebhook)
	}

	chatPage := router.Group("/posts", h.userIdentity)
	{
		chatPage.POST("/create", h.createPost)
		chatPage.GET("/", h.getAllPosts)
		chatPage.POST("/:id/create", h.createComment)
		chatPage.GET("/:id", h.GetAllComments)
	}

	return router
}
