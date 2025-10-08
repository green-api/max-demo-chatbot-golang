package scenes

import (
	chatbot "github.com/green-api/max-chatbot-golang"
	"github.com/green-api/max-demo-chatbot-golang/util"
)

type MainMenuScene struct {
}

func (s MainMenuScene) Start(bot *chatbot.Bot) {
	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		if !util.IsSessionExpired(message) {
			s.SendMainMenu(message)
		} else {
			message.ActivateNextScene(StartScene{})
			message.SendText(util.GetString([]string{"select_language"}))
		}
	})
}

func (s MainMenuScene) SendMainMenu(message *chatbot.Notification) {
	text, _ := message.Text()
	switch text {
	case "1":
		s.sendMainMenu(message, "en")
	case "2":
		s.sendMainMenu(message, "kz")
	case "3":
		s.sendMainMenu(message, "ru")
	case "4":
		s.sendMainMenu(message, "es")
	case "5":
		s.sendMainMenu(message, "he")
	// enable only and only when ar language will be released and ready
	// otherwise, crash happens
	//case "6":
	//	s.sendMainMenu(message, "ar")
	default:
		message.SendText(util.GetString([]string{"specify_language"}))
	}
}

func (s MainMenuScene) sendMainMenu(message *chatbot.Notification, lang string) {
	message.UpdateStateData(map[string]interface{}{"lang": lang})

	var welcomeFileURL string
	if lang == "en" {
		welcomeFileURL = "https://raw.githubusercontent.com/green-api/max-demo-chatbot-golang/refs/heads/master/assets/welcome_en.jpg"
	} else {
		welcomeFileURL = "https://raw.githubusercontent.com/green-api/max-demo-chatbot-golang/refs/heads/master/assets/welcome_ru.jpg"
	}

	var name string
	if message.Body["senderData"].(map[string]interface{})["senderName"] != nil && len(message.Body["senderData"].(map[string]interface{})["senderName"].(string)) > 0 {
		name = ", " + message.Body["senderData"].(map[string]interface{})["senderName"].(string)
	}
	message.SendUrlFile(welcomeFileURL,
		"welcome.jpg",
		util.GetString([]string{"welcome_message", lang})+name+"!"+"\n"+util.GetString([]string{"menu", lang}))
	message.ActivateNextScene(EndpointsScene{})

}
