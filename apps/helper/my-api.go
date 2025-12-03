package helper

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

// send whatsapp
func SendWa(c *fiber.Ctx, uniqid string, full_name string, phone string) error {
	// Payload langsung diisi di dalam function
	link := c.BaseURL() + "/my-qr/" + uniqid
	payload := map[string]interface{}{
		"to_number":              "62" + phone,
		"to_name":                full_name,
		"message_template_id":    "76293087-b4a2-4cda-86b3-7c74626fbf17",
		"channel_integration_id": "488e5ce4-90c5-498d-9813-c8c90c80072f",
		"language": map[string]string{
			"code": "id",
		},
		"parameters": map[string]interface{}{
			"body": []map[string]string{
				{"key": "1", "value": "nama", "value_text": full_name},
				{"key": "2", "value": "link", "value_text": link},
			},
		},
	}

	// Convert ke JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Setup request
	url := "https://service-chat.qontak.com/api/open/v1/broadcasts/whatsapp/direct"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Authorization token (sebaiknya dari env variable)
	req.Header.Set("Authorization", "Bearer NX5m3etJv_eTZKMhhC-aI5OGYxGS5B5_o3i4w5uDAHs")
	req.Header.Set("Content-Type", "application/json")

	// Kirim request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	defer resp.Body.Close()

	return nil
}

// send email
func SendEmail(c *fiber.Ctx, uniqid string, full_name string, email string) error {
	link := c.BaseURL() + "/my-qr/" + uniqid
	params := url.Values{}
	params.Add("from_name", "[Reminder] Pengambilan Racepack AliRun Cinta Bumi 2025")
	params.Add("from_email", "no-reply@goopps.id")
	params.Add("recipient", email)
	params.Add("subject", "[Reminder] Pengambilan Racepack AliRun Cinta Bumi 2025")
	params.Add("content", "<p>Halo AliRunners!<br>Dear <strong>"+full_name+"</strong>&nbsp;<br>Terima kasih telah mendaftar sebagai peserta AliRun Cinta Bumi 2025 🌱<br><br>Berikut informasi penting terkait pengambilan racepack:<br><br>📅 Tanggal:<strong>Jumat, 31 Oktober 2025 (satu hari saja)</strong><br>🕓 Waktu:<strong>10.00 &ndash; 20.00 WIB</strong><br>📍 Tempat:<strong>BXC MALL Area WICO, Lt. GF (Belakang Remboelan)</strong><br><br>Wajib membawa:<br>KTP asli<br>Jika diwakilkan, wajib membawa fotokopi KTP peserta yang diwakilkan<br><br>Harap datang sesuai jadwal dan tunjukan e-Ticket Anda agar proses pengambilan berjalan lancar.<br>Sampai jumpa di hari acara, dan semangat berlari untuk bumi yang lebih hijau! 💚<br><br><strong>Link e-Ticket</strong></p><p>"+link+"<br><br>Salam,<br>Panitia AliRun Cinta Bumi 2025</p>")
	params.Add("api_token", "45ab6fa285888b841b2a394aceff021e")

	url := "https://api.mailketing.co.id/api/v1/send?" + params.Encode()

	req, _ := http.NewRequest("POST", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer resp.Body.Close()
	return nil
}
