package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
	"tg_bot/config"
	"tg_bot/internal/bot"
	"tg_bot/internal/user/handlers"
)

func main() {
	fmt.Println("Starting bot...")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers.Handler),
		bot.WithCallbackQueryDataHandler("letsgo", bot.MatchTypeExact, handlers.LetsGoButtonHandler),
		bot.WithCallbackQueryDataHandler("yesconfirmationpassword", bot.MatchTypeExact, handlers.YesButtonHandler),
		bot.WithCallbackQueryDataHandler("noconfirmationpassword", bot.MatchTypeExact, handlers.NoButtonHandler),
		bot.WithCallbackQueryDataHandler("lockdiary", bot.MatchTypeExact, handlers.LockDiaryButtonHandler),
		bot.WithCallbackQueryDataHandler("lefttimezone", bot.MatchTypeExact, handlers.LeftTimeZoneButtonHandler),
		bot.WithCallbackQueryDataHandler("savetimezone", bot.MatchTypeExact, handlers.SaveTimeZoneButtonHandler),
		bot.WithCallbackQueryDataHandler("righttimezone", bot.MatchTypeExact, handlers.RightTimeZoneButtonHandler),
		bot.WithCallbackQueryDataHandler("backtomenufromtimezone", bot.MatchTypeExact, handlers.MenuDiaryButtonHandler),
		bot.WithCallbackQueryDataHandler("finddate", bot.MatchTypeExact, handlers.MenuFindDateWithYearsHandler),

		bot.WithAllowedUpdates(internal.AllowedUpdates()),

		bot.WithCallbackQueryDataHandler("years_", bot.MatchTypePrefix, handlers.YearsButtonHandler),
		bot.WithCallbackQueryDataHandler("months_", bot.MatchTypePrefix, handlers.MonthsButtonHandler),
		bot.WithCallbackQueryDataHandler("days_", bot.MatchTypePrefix, handlers.DaysButtonHandler),
		bot.WithCallbackQueryDataHandler("closeFD", bot.MatchTypePrefix, handlers.CloseFDButtonHandler),
		bot.WithCallbackQueryDataHandler("clearall", bot.MatchTypeExact, handlers.ClearAllHandler),
	}

	b, err := bot.New(config.GetEntranceToken(), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/menu", bot.MatchTypeExact, handlers.MenuDiaryCommandHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/lock", bot.MatchTypeExact, handlers.LockDiaryCommandHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/changetimezone", bot.MatchTypeExact, handlers.MenuTineZoneButtonHandler)
	b.Start(ctx)
}
