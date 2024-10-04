package listeners

import "github.com/bwmarrin/discordgo"

func RegisterListeners(bot *discordgo.Session) {
	bot.AddHandler(MessageCreate)
}