package scenes

import (
	"encoding/json"

	chatbot "github.com/green-api/max-chatbot-golang"
	"github.com/green-api/max-demo-chatbot-golang/util"
)

type CreateGroupScene struct {
}

func (s CreateGroupScene) Start(bot *chatbot.Bot) {

	bot.IncomingMessageHandler(func(message *chatbot.Notification) {
		if !util.IsSessionExpired(message) {
			lang := message.GetStateData()["lang"].(string)
			senderId, _ := message.Sender()

			groupResp, err := message.Groups().CreateGroup(
				util.GetString([]string{"group_name", lang}),
				[]string{senderId})
			if err != nil {
				*message.ErrorChannel <- err
				message.ActivateNextScene(EndpointsScene{})
			}
			var group map[string]interface{}
			_ = json.Unmarshal(groupResp.Body, &group)

			var groupId = group["chatId"].(string)
			message.StateManager.Create(groupId)
			message.StateManager.UpdateStateData(groupId, message.GetStateData())
			message.StateManager.ActivateNextScene(groupId, EndpointsScene{})

			resp, err := message.Groups().SetGroupPicture(
				"assets/Group_avatar.jpg",
				groupId)
			if err != nil {
				*message.ErrorChannel <- err
			} else {
				var picResp map[string]interface{}
				_ = json.Unmarshal(resp.Body, &picResp)

				if picResp["setGroupPicture"].(bool) {
					_, err := message.Sending().SendMessage(
						groupId,
						util.GetString([]string{"send_group_message", lang})+util.GetString([]string{"links", lang, "groups_documentation"}))
					if err != nil {
						*message.ErrorChannel <- err
					}
				} else {
					_, err := message.Sending().SendMessage(
						groupId,
						util.GetString([]string{"send_group_message_set_picture_false", lang})+util.GetString([]string{"links", lang, "groups_documentation"}))
					if err != nil {
						*message.ErrorChannel <- err
					}
				}
			}
			message.SendText(util.GetString([]string{"group_created_message", lang}) +
				group["groupInviteLink"].(string))
			message.SendText(util.GetString([]string{"add_to_contact", lang}), "true")
			message.ActivateNextScene(EndpointsScene{})
		} else {
			message.ActivateNextScene(MainMenuScene{})
			message.SendText(util.GetString([]string{"select_language"}))
		}
	})
}
