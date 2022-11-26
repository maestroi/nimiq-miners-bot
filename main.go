package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/mitchellh/mapstructure"
	"github.com/tkanos/gonfig"
)

const (
	configFile   = "./config.json"
	questionFile = "./questions.json"
)

func GetConfig(params ...string) Configuration {
	configuration := Configuration{}
	fileName := configFile
	err := gonfig.GetConf(fileName, &configuration)
	if err != nil {
		log.Fatal("Error Loading config: ", err)
	}
	log.Println("Config loaded succesfully")
	return configuration
}

func ReadQuestions() string {
	content, err := ioutil.ReadFile(questionFile)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	m := map[string]interface{}{}
	err = json.Unmarshal([]byte(content), &m)
	if err != nil {
		log.Fatal("We can't unmarshal data!")
	}

	question := Questions{}

	questionString := ""

	for _, name := range m {
		mapstructure.Decode(name, &question)

		qs := fmt.Sprintf("<b>Question</b>: %s \n<b>Answer</b>: %s\n\n", question.Question, question.Answer)

		questionString += qs
	}

	return questionString
}

func main() {
	log.Println("Bot succesfully started!")
	WelcomeText := `Welcome Nimiq Miner!
	Please use /faq First before asking any questions!

	Remember: Mining with Nimiq stops when the NIMIQ 2.0 is launched this will be full POS, but no date yet!`

	config := GetConfig()

	bot, err := tgbotapi.NewBotAPI(config.BOT_TOKEN)
	if err != nil {
		log.Printf("Did you add a bot token!?")
		log.Panic(err)
	}
	bot.Debug = bool(config.BOT_DEBUG)

	log.Printf("Authorized on account %s ID: %v", bot.Self.UserName, bot.Self.ID)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("User: %s | command: %s | chatID: %v |", update.Message.From.UserName, update.Message.Text, update.Message.MessageID)

		if update.Message.NewChatMembers != nil {
			log.Println(update.Message.NewChatMembers)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = WelcomeText
			bot.Send(msg)
		}

		if update.Message.IsCommand() {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			switch update.Message.Command() {
			case "help":
				msg.Text = "type /faq or /status."
			case "faq":
				msg.ParseMode = "html"
				msg.Text = ReadQuestions()
			case "welcome":
				msg.ParseMode = "html"
				msg.Text = WelcomeText
			case "status":
				msg.Text = "I'm Healthy!."
			default:
				msg.Text = "Unknow command try: /help"
			}
			bot.Send(msg)
		}

	}
}
