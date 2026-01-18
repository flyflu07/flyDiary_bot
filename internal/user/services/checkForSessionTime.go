package services

import (
	"fmt"
	"strconv"
	"tg_bot/internal/user/model"
	"time"
)

func CheckForSessionTime(userID int64) {
	for {
		time.Sleep(5 * time.Second)
		if model.InfoAbUser[userID].TimeOfStartSession.Add(model.InfoAbUser[userID].TimeOfAutoLogOut).After(time.Now()) {
			delete(model.InfoAbUser, userID)
			fmt.Println("Сессиу у " + strconv.FormatInt(userID, 10) + "закончилась")

		} else {
			fmt.Printf("Сессиу у %s еще длится", strconv.FormatInt(userID, 10))
		}
	}
}
