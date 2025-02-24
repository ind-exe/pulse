package notification

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	envvar "oob-server/envVar"
	"oob-server/models"
)

func SendTelegram(data models.Notifier) {
	telegramBotToken, err := envvar.GetVar("TELEGRAM_BOT_TOKEN")
	if err != nil {
		log.Println("notifService | fail : could not read the TELEGRAM_BOT_TOKEN env var")
		return
	}
	
	telegramChatId, err := envvar.GetVar("TELEGRAM_CHAT_ID")
	if err != nil {
		log.Println("notifService | fail : could not read the TELEGRAM_CHAT_ID env var")
		return
	}
	
	apiURL := "https://api.telegram.org/bot" + telegramBotToken + "/sendMessage"
	payload := struct{
		ChatId string `json:"chat_id"`
		Text string `json:"text"`
		ParseMode string `json:"parse_mode"`
	}{
		ChatId: telegramChatId,
		Text: data.ToString(),
		ParseMode: "Markdown",
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("NotifSerivce | fail : could not marshal the payload")
		return
	}
	resp, err := http.Post(apiURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("NotifSerivce | fail : could not send the payload to telegram bot")
		return
	}
	defer resp.Body.Close()
	

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		log.Printf("NotifService | fail : failed to send message, status: %s\n", resp.Status)
		return
	}
	log.Println("NotifService | success : sent the log to telegram bot")
}

func SendDiscord(data models.Notifier) {
	discordWebhook, err := envvar.GetVar("DISCORD_WEBHOOK")
	if err != nil {
		log.Println("notifService | fail : could not read the DISCORD_WEBHOOK env var")
		return
	}
	payload := struct{
		Content string `json:"content"`
	}{Content: data.ToString()}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Println("NotifSerivce | fail : could not marshal the payload")
		return
	}
	resp, err := http.Post(discordWebhook, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Println("NotifSerivce | fail : could not send the payload to discord webhook")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent && resp.StatusCode != http.StatusOK {
		log.Printf("NotifService | fail : failed to send message, status: %s\n", resp.Status)
		return
	}
	log.Println("NotifService | success : sent the log to discord")
}