package notification

import "github.com/ind-exe/pulse/models"



type NotifMarker struct {
	Telegram bool `json:"telegram"`
	Discord bool	`json:"discord"`
}

func (nm *NotifMarker) Decider(data models.Notifier) {
	if nm.Discord {
		SendDiscord(data)
	}

	if nm.Telegram {
		SendTelegram(data)
	}
}