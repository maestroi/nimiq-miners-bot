package main

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	FAQText := `FAQ:
	Question: When PoS!?
	Answer: There is no official date for Nimiq 2.0!

	Question: Will Nimiq always allow mining?
	Answer: No, Nimiq will go full proof of stake(PoS) after that you will not be able to mine anymore!
	
	Question: Can i mine on a phone?
	Answer: Yes that is possible, BUT Don't do it! it will give you almost no reward and is bad for your phone.
	
	Question: Should i mine with my laptop?
	Answer: IF your laptop has really good cooling you could mine on it. But overall we do not recommend it.
	
	Question: Can i mine with GPU?
	Answer: Yes, GPU and FPGA are most efficient right now! mining on CPU will probably cost you money.
	
	Question: What is the easiest miner to use?
	Answer: https://hub.shortnim.me/setupMiner This is a easy and fast wat to setup a gpu/cpu miner for new users
	
	Question: At what pools can i mine?
	Answer: An easy to use interface for pools can be found here: https://hub.shortnim.me and if you want to see all pools: https://miningpoolstats.stream/nimiq
	Remember try to mine on lower hashrate pools to secure the network!
	
	Question: How much will i earn by mining?
	Answer: You can calculate your mining reward with the following calculator: https://calc.nimiqx.com/`

	WelcomeText := `Welcome Nimiq Miner!
	Please use /faq First before asking any questions!

	Remember: Mining with Nimiq stops when the NIMIQ 2.0 is launched this will be full POS, but no date yet!`

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Printf("Did you add a bot token!?")
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

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
				msg.Text = FAQText
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
