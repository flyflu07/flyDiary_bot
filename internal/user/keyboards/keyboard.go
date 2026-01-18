package keyboards

import (
	"github.com/go-telegram/bot/models"
)

func StartButton() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{
				Text:         "Приступим",
				CallbackData: "letsgo",
			},
		},
	}
}

func ConfirmationPasswordButtons() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{

				Text:         "Да",
				CallbackData: "yesconfirmationpassword",
			},
			{
				Text:         "Нет",
				CallbackData: "noconfirmationpassword",
			},
		},
	}
}

func MenuWithLock() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{
				Text:         "Найти свои записи",
				CallbackData: "finddate",
			},
		},
		{
			{
				Text:         "🔒Loсk",
				CallbackData: "lockdiary",
			},
		},
	}
	// +Удалить все записи из дневника
}
func MenuWithUnlock() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{
				Text:         "Найти свои записи",
				CallbackData: "finddate",
			},
		},
		{
			//{
			//
			//	Text:         "🔓Unloсk",
			//	CallbackData: "unlockdiary",
			//},
		},
	}
	// +Удалить все записи из дневника
}

func ChangeTimezoneButtons() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{
				Text:         "⬅️",
				CallbackData: "lefttimezone",
			},
			{
				Text:         "Сохранить",
				CallbackData: "savetimezone",
			},
			{
				Text:         "➡️",
				CallbackData: "righttimezone",
			},
		},

		{
			{
				Text:         "Вернуться назад ⬅️",
				CallbackData: "backtomenufromtimezone",
			},
		},
	}
}
