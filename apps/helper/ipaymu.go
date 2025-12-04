package helper

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
)

type RequestBody struct {
	Product     []string `json:"product"`
	Qty         []string `json:"qty"`
	Price       []string `json:"price"`
	ReturnUrl   string   `json:"returnUrl"`
	CancelUrl   string   `json:"cancelUrl"`
	NotifyUrl   string   `json:"notifyUrl"`
	ReferenceId string   `json:"referenceId"`
	BuyerName   string   `json:"buyerName"`
	BuyerEmail  string   `json:"buyerEmail"`
	BuyerPhone  string   `json:"buyerPhone"`
}

type ApiResponse struct {
	Status int `json:"Status"`
	Data   struct {
		SessionID string `json:"SessionID"`
		Url       string `json:"Url"`
	} `json:"Data"`
}

func GenerateSignature(va, apiKey, requestBody, method string) string {
	// Create the string to sign
	stringToSign := fmt.Sprintf("%s:%s:%s:%s", method, va, requestBody, apiKey)

	// Create HMAC SHA256 hash
	hmac := hmac.New(sha256.New, []byte(apiKey))
	hmac.Write([]byte(stringToSign))
	signature := hex.EncodeToString(hmac.Sum(nil))

	return signature
}

func Payment(c *fiber.Ctx, name string, phone string, email string, product string, price, link string, uniqid string) error {
	// Request body data
	va := os.Getenv("ipaymuVa")      // Get this from iPaymu dashboard
	apiKey := os.Getenv("ipaymuKey") // Get this from iPaymu dashboard
	url := os.Getenv("url")          // For development

	method := "POST"

	// Request body
	body := RequestBody{
		Product:     []string{product},
		Qty:         []string{"1"},
		Price:       []string{price},
		ReturnUrl:   link + "/tiketku/" + uniqid,
		CancelUrl:   link + "/tiketku/" + uniqid,
		NotifyUrl:   link + "/tiketku/" + uniqid,
		ReferenceId: uniqid,
		BuyerName:   name,  // optional
		BuyerEmail:  email, // optional
		BuyerPhone:  phone,
	}

	// Convert body to JSON
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error marshalling request body")
	}

	// Generate the SHA256 hash of the request body
	requestBody := fmt.Sprintf("%x", sha256.Sum256(jsonBody))

	// Generate the signature
	signature := GenerateSignature(va, apiKey, requestBody, method)

	// Timestamp (YYYYMMDDHHMMSS)
	timestamp := time.Now().Format("20060102150405")

	// Prepare the headers
	headers := map[string]string{
		"Accept":       "application/json",
		"Content-Type": "application/json",
		"va":           va,
		"signature":    signature,
		"timestamp":    timestamp,
	}

	// Make the HTTP request to iPaymu API
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error creating request")
	}

	// Set the headers
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Set the body as JSON
	req.Body = io.NopCloser(bytes.NewReader(jsonBody))

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error sending request")
	}
	defer resp.Body.Close()

	// Parse the response
	var apiResp ApiResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return c.Status(http.StatusInternalServerError).SendString("Error decoding response")
	}

	// Check the response status
	if apiResp.Status == 200 {
		// Redirect to the payment URL
		return c.Redirect(apiResp.Data.Url)
	} else {
		// Return error response from API
		return c.Status(http.StatusBadRequest).JSON(apiResp)
	}
}
