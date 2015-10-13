package main

import (
	"log"
	"github.com/Syfaro/telegram-bot-api"
	"github.com/valiknet18/IntertechTelegramBot/config"
	"strings"
	// "fmt"
)

func main() {
	config.ParseConfig()

	bot, err := tgbotapi.NewBotAPI("123528822:AAH-OEyfyOjJjy9Jjmq5ZJEbiqwBF-ybd8Q")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	err = bot.UpdatesChan(u)
	if err != nil {
		log.Panic(err)
	}

	for update := range bot.Updates {
		resultString := strings.Split(update.Message.Text, "|")

		var message tgbotapi.MessageConfig

		if len(resultString) > 0 {
			switch resultString[0] {
			case "CreateTask":
				{
					result := CreateNewTask(resultString[3], resultString[1], resultString[2])

					if result {
						message = tgbotapi.NewMessage(update.Message.Chat.ID, "Новое задание для пользователя " + resultString[3] + " успешно добавлено")
					} else {
						message = tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла какае-то ошибка")		
					}
				}

			case "ViewTasks":
				{
					resultMessage := GetAllUserTasks(resultString[1])
					message = tgbotapi.NewMessage(update.Message.Chat.ID, resultMessage)
				}

			case "DropTask":
				{
					result := RemoveTaskByTaskName(resultString[2], resultString[1])
					if result {
						message = tgbotapi.NewMessage(update.Message.Chat.ID, "Выбраное задание успешно удалено")
					} else {
						message = tgbotapi.NewMessage(update.Message.Chat.ID, "Произошла какае-то ошибка")		
					}
				}
			}

			bot.SendMessage(message)
		}
	}	
}