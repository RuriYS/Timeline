package listeners

import (
	"Timeline/internal/database"

	"github.com/bwmarrin/discordgo"
)

type Client interface {
	GetDatabase() *database.Database
}

func RegisterListeners(c Client, s *discordgo.Session) {
	s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		MessageCreate(s, m, c)
	})

	// s.AddHandler(func(s *discordgo.Session, m *discordgo.MessageDelete) {
	// 	MessageDelete(c, s, m)
	// })

	s.AddHandler(func(s *discordgo.Session, r *discordgo.Ready) {
		Ready(c, s, r)
	})
}
