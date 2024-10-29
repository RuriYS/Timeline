package listeners

import (
	"Timeline/internal/logger"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, c *discordgo.Ready) {
	l := logger.GetLogger()
	l.Info("Logged in as %s#%s", c.User.Username, c.User.Discriminator)
}
