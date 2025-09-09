package registry

import gptbot "github.com/green-api/max-chatgpt-go"

var gptHelperInstance *gptbot.MaxGptBot

func RegisterGptHelper(instance *gptbot.MaxGptBot) {
	gptHelperInstance = instance
}

func GetGptHelper() *gptbot.MaxGptBot {
	return gptHelperInstance
}
