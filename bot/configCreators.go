package bot

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type ChatAdministratorsConfig struct {
	ChatConfig tgbotapi.ChatConfig
}

func restrictConfigToMuteCreator(user *tgbotapi.User, chat *tgbotapi.Chat) tgbotapi.RestrictChatMemberConfig {
	return tgbotapi.RestrictChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID:          chat.ID,
			ChannelUsername: chat.UserName,
			UserID:          user.ID,
		},
		UntilDate: time.Now().Unix(),
	}
}

func restrictConfigToUnMuteCreator(user *tgbotapi.User, chat *tgbotapi.Chat) tgbotapi.RestrictChatMemberConfig {
	trus := true
	return tgbotapi.RestrictChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID:          chat.ID,
			ChannelUsername: chat.UserName,
			UserID:          user.ID,
		},
		UntilDate:             time.Now().Unix(),
		CanSendMessages:       &trus,
		CanSendMediaMessages:  &trus,
		CanSendOtherMessages:  &trus,
		CanAddWebPagePreviews: &trus,
	}
}

func kickConfigCreator(user *tgbotapi.User, chat *tgbotapi.Chat) tgbotapi.KickChatMemberConfig {
	return tgbotapi.KickChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID:          chat.ID,
			ChannelUsername: chat.UserName,
			UserID:          user.ID,
		},
		UntilDate: time.Now().Unix(),
	}
}

func promoteConfigCreator(user *tgbotapi.User, chat *tgbotapi.Chat, promoteOrLower bool) tgbotapi.PromoteChatMemberConfig {
	return tgbotapi.PromoteChatMemberConfig{
		ChatMemberConfig: tgbotapi.ChatMemberConfig{
			ChatID: chat.ID,
			UserID: user.ID,
		},
		CanChangeInfo:      &promoteOrLower,
		CanPostMessages:    &promoteOrLower,
		CanEditMessages:    &promoteOrLower,
		CanDeleteMessages:  &promoteOrLower,
		CanInviteUsers:     &promoteOrLower,
		CanRestrictMembers: &promoteOrLower,
		CanPinMessages:     &promoteOrLower,
		CanPromoteMembers:  &promoteOrLower,
	}
}
