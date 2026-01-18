package main

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"os"
	"os/signal"
	"tg_bot/config"
	"tg_bot/internal/bot"
	handlers2 "tg_bot/internal/user/handlers"
)

func main() {
	fmt.Println("Starting bot...")
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handlers2.Handler),
		bot.WithCallbackQueryDataHandler("letsgo", bot.MatchTypeExact, handlers2.LetsGoButtonHandler),
		bot.WithCallbackQueryDataHandler("yesconfirmationpassword", bot.MatchTypeExact, handlers2.YesButtonHandler),
		bot.WithCallbackQueryDataHandler("noconfirmationpassword", bot.MatchTypeExact, handlers2.NoButtonHandler),
		bot.WithCallbackQueryDataHandler("lockdiary", bot.MatchTypeExact, handlers2.LockDiaryButtonHandler),
		bot.WithCallbackQueryDataHandler("lefttimezone", bot.MatchTypeExact, handlers2.LeftTimeZoneButtonHandler),
		bot.WithCallbackQueryDataHandler("savetimezone", bot.MatchTypeExact, handlers2.SaveTimeZoneButtonHandler),
		bot.WithCallbackQueryDataHandler("righttimezone", bot.MatchTypeExact, handlers2.RightTimeZoneButtonHandler),
		bot.WithCallbackQueryDataHandler("backtomenufromtimezone", bot.MatchTypeExact, handlers2.MenuDiaryButtonHandler),
		bot.WithCallbackQueryDataHandler("finddate", bot.MatchTypeExact, handlers2.MenuFindDateWithYearsHandler),

		bot.WithAllowedUpdates(internal.AllowedUpdates()),

		bot.WithCallbackQueryDataHandler("years_", bot.MatchTypePrefix, handlers2.YearsButtonHandler),
		bot.WithCallbackQueryDataHandler("months_", bot.MatchTypePrefix, handlers2.MonthsButtonHandler),
		bot.WithCallbackQueryDataHandler("days_", bot.MatchTypePrefix, handlers2.DaysButtonHandler),
		bot.WithCallbackQueryDataHandler("closeFD", bot.MatchTypePrefix, handlers2.CloseFDButtonHandler),
		bot.WithCallbackQueryDataHandler("clearall", bot.MatchTypeExact, handlers2.ClearAllHandler),
	}

	b, err := bot.New(config.GetEntranceToken(), opts...)
	if nil != err {
		// panics for the sake of simplicity.
		// you should handle this error properly in your code.
		panic(err)
	}
	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeExact, handlers2.StartHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/menu", bot.MatchTypeExact, handlers2.MenuDiaryCommandHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/lock", bot.MatchTypeExact, handlers2.LockDiaryCommandHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "/changetimezone", bot.MatchTypeExact, handlers2.MenuTineZoneButtonHandler)
	b.Start(ctx)
}
