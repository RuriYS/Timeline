package listeners

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func Ready(s *discordgo.Session, c *discordgo.Ready)  {
	log.Printf("INFO: Logged in as %s#%s", c.User.Username, c.User.Discriminator)
}