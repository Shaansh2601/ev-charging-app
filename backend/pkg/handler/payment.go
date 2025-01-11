package handler

import (
	"backend/pkg/model"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/paymentintent"
	"github.com/stripe/stripe-go/webhook"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	// For sample support and debugging, not required for production:
	stripe.SetAppInfo(&stripe.AppInfo{
		Name:    "stripe-samples/accept-a-payment/payment-element",
		Version: "0.0.1",
		URL:     "https://github.com/stripe-samples",
	})
}

func (h *Handler) handleConfig(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"publishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY"),
	})
}

/*
func (h *Handler) createPayment(c *gin.Context) {
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	// Process the JSON payload
	c.JSON(200, gin.H{"message": "JSON received", "data": jsonData})
}
*/

func (h *Handler) createPayment(c *gin.Context) {
	var input model.Transaction
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid body")
		return
	}

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(input.Cost)),
		Currency: stripe.String("GBP"),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		if stripeErr, ok := err.(*stripe.Error); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": stripeErr.Error()})
			newErrorResponse(c, http.StatusBadRequest, stripeErr.Error())
		} else {
			newErrorResponse(c, http.StatusInternalServerError, "Unknown server error")
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"clientSecret": pi.ClientSecret})
}

func (h *Handler) handleWebhook(c *gin.Context) {
	b, err := io.ReadAll(c.Request.Body)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to read request body")
	}

	_, err = webhook.ConstructEvent(b, c.GetHeader("Stripe-Signature"), os.Getenv("STRIPE_WEBHOOK_SECRET"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Failed to construct webhook event")
	}

	c.JSON(http.StatusOK, "Checkout Session completed!")
}
