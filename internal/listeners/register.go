package listeners

import "github.com/bwmarrin/discordgo"

func RegisterListeners(bot *discordgo.Session) {
	bot.AddHandler(MessageCreate)
	bot.AddHandler(MessageDelete)
	bot.AddHandler(ChannelCreate)
	bot.AddHandler(ChannelDelete)
	bot.AddHandler(ChannelUpdate)
	bot.AddHandler(Ready)
}