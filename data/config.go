package data

type ConfigScheme struct{
	TelegramBotToken string
	TelegramChatId string
	Domain string
	DiscordWebhook string
}

var Config ConfigScheme