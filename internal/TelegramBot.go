package internal

import (
	"context"
	"fmt"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var tgBot *bot.Bot

func InitTelegramBot() {
	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	tgBot, _ = bot.New(Configuration.TelegramBotToken, opts...)
}

func NotifyInmueble(inm Inmueble) {
	tgBot.SendPhoto(Context, &bot.SendPhotoParams{
		ChatID:  Configuration.TelegramUserId,
		Photo:   &models.InputFileString{Data: inm.foto},
		Caption: fmt.Sprintf("Encontré este depto:\nID: %v\nDirección: %v\nPrecio: %v\nSuperficie total/cubierta: %vm2-%vm2\n%v habitaciones\n%v baños\nMiralo en https://inmoclick.com.ar%v", inm.kid, inm.direccion, inm.precio, inm.supTotal, inm.supCubierta, inm.habitaciones, inm.banos, inm.url),
	})
}

func SendMessage(text string) {
	tgBot.SendMessage(Context, &bot.SendMessageParams{
		ChatID: Configuration.TelegramUserId,
		Text:   text,
	})
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   "Sigo vivo",
	})
}
