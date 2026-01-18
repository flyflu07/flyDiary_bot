package keyboards

import (
	"github.com/go-telegram/bot/models"
	"tg_bot/internal/storage"
)

func FindYears(userID int64) [][]models.InlineKeyboardButton {
	var keyboardsYears [][]models.InlineKeyboardButton
	years := storage.GetUniqueYearsOfDiaryEntries(userID)

	var currentRow []models.InlineKeyboardButton

	for _, year := range years {
		btn := models.InlineKeyboardButton{
			Text:         year,
			CallbackData: "years_" + year,
		}

		currentRow = append(currentRow, btn)

		if len(currentRow) == 8 {
			keyboardsYears = append(keyboardsYears, currentRow)
			currentRow = []models.InlineKeyboardButton{} // Сброс для новой строки
		}
	}

	if len(currentRow) > 0 {
		keyboardsYears = append(keyboardsYears, currentRow)
	}
	keyboardsYears = append(keyboardsYears, []models.InlineKeyboardButton{
		{
			Text:         "Закрыть",
			CallbackData: "closeFD_years",
		},
	})

	return keyboardsYears
}

func FindMonths(userID int64) [][]models.InlineKeyboardButton {
	var keyboardsMonths [][]models.InlineKeyboardButton
	months := storage.GetUniqueMonthsOfDiaryEntries(userID)

	var currentRow []models.InlineKeyboardButton

	for _, month := range months {
		btn := models.InlineKeyboardButton{
			Text:         month,
			CallbackData: "months_" + month,
		}

		currentRow = append(currentRow, btn)

		if len(currentRow) == 8 {
			keyboardsMonths = append(keyboardsMonths, currentRow)
			currentRow = []models.InlineKeyboardButton{} // Сброс для новой строки
		}
	}

	if len(currentRow) > 0 {
		keyboardsMonths = append(keyboardsMonths, currentRow)
	}

	keyboardsMonths = append(keyboardsMonths, []models.InlineKeyboardButton{
		{
			Text:         "Назад",
			CallbackData: "closeFD_month",
		},
	})

	return keyboardsMonths
}

func FindDays(userID int64) [][]models.InlineKeyboardButton {
	var keyboardsDays [][]models.InlineKeyboardButton
	days := storage.GetUniqueDaysOfDiaryEntries(userID)

	var currentRow []models.InlineKeyboardButton

	for _, day := range days {
		btn := models.InlineKeyboardButton{
			Text:         day,
			CallbackData: "days_" + day,
		}

		currentRow = append(currentRow, btn)

		if len(currentRow) == 8 {
			keyboardsDays = append(keyboardsDays, currentRow)
			currentRow = []models.InlineKeyboardButton{} // Сброс для новой строки
		}
	}

	if len(currentRow) > 0 {
		keyboardsDays = append(keyboardsDays, currentRow)
	}

	keyboardsDays = append(keyboardsDays, []models.InlineKeyboardButton{
		{
			Text:         "Назад",
			CallbackData: "closeFD_days",
		},
	})

	return keyboardsDays
}

func ClearAll() [][]models.InlineKeyboardButton {
	return [][]models.InlineKeyboardButton{
		{
			{
				Text:         "Скрыть",
				CallbackData: "clearall",
			},
		},
	}
}
