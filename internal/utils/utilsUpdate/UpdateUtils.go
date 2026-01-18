package utilsUpdate

import (
	"github.com/go-telegram/bot/models"
	"strings"
)

// ExtractOriginSender extracts user that triggered an update
func ExtractOriginSender(update *models.Update) *models.User {
	if update.Message != nil {
		return update.Message.From
	}

	if update.MessageReaction != nil {
		return update.MessageReaction.User
	}

	return &update.CallbackQuery.From
}

// ExtractChatID extracts chat id where update was triggered
func ExtractChatID(update *models.Update) int64 {
	if update.Message != nil {
		return update.Message.Chat.ID
	}

	if update.MessageReaction != nil {
		return update.MessageReaction.Chat.ID
	}

	return update.CallbackQuery.Message.Message.Chat.ID
}

// ExtractMessageID extracts message id of updated message
func ExtractMessageID(update *models.Update) int {
	if update.Message != nil {
		return update.Message.ID
	}

	if update.MessageReaction != nil {
		return update.MessageReaction.MessageID
	}

	return update.CallbackQuery.Message.Message.ID
}

// ExtractUserID extracts user id of updated message
func ExtractUserID(update *models.Update) int64 {
	if update.Message != nil {
		return update.Message.From.ID
	}
	if update.MessageReaction != nil {
		return update.MessageReaction.User.ID
	}
	return update.CallbackQuery.From.ID
}

func ExtractCallbackQueryID(update *models.Update) string {
	if update.Message != nil {
		return update.CallbackQuery.ID
	}
	return "Unknown"
}

// ExtractUsername extracts username of the user who triggered an update
func ExtractUsername(update *models.Update) string {
	username := ExtractOriginSender(update).Username
	return username
}

// ExtractName extracts name of the user who triggered an update
func ExtractName(update *models.Update) string {
	user := ExtractOriginSender(update)
	return strings.TrimSpace(user.FirstName + " " + user.LastName)
}

// ExtractEntityContentFromMessage extracts content of MessageEntity (link, tag, etc.) from message
func ExtractEntityContentFromMessage(message *models.Message, entity *models.MessageEntity) (text string) {
	from, to := adjustOffsets(message.Text, entity.Offset, entity.Length)
	return strings.TrimSpace(message.Text[from:to])
}

// IsMessageReactionUpdated returns true if update is "add/remove message reaction"
func IsMessageReactionUpdated(update *models.Update) bool {
	return update.MessageReaction != nil
}

// IsMessageForwarded returns true if update is "forwarding message"
func IsMessageForwarded(update *models.Update) bool {
	return update.Message != nil && update.Message.ForwardOrigin != nil
}

// adjustOffsets UTF-8 and UTF-16 compatibility (black magic). Returns new "from" and "to" indexes
func adjustOffsets(s string, offset int, length int) (int, int) {
	// the problem is:
	// UTF-16 symbols take 2 runes in golang
	// if we have a string that contains both of latin and cyrillic
	// letters, we have to recalculate offsets this ugly way
	currentOffset := 0
	from := -1
	for pos := range s {
		if currentOffset >= offset && from == -1 {
			from = pos
		}
		if currentOffset >= offset+length {
			return from, pos
		}
		currentOffset++
	}

	return from, len(s)
}
