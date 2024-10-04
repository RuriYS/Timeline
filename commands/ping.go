package commands

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
)

func Ping(s *discordgo.Session, m *discordgo.MessageCreate) {
    var totalLatency time.Duration
    iterations := 5

    msg, _ := s.ChannelMessageSend(m.ChannelID, "Pinging...")

    for i := 0; i < iterations; i++ {
        start := time.Now()
		s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf(
			"Pinging... (Avg: `%.2f` ms)",
			safeArg(totalLatency, i)))
        latency := time.Since(start)
        totalLatency += latency
    }

    avgLatency := float64(totalLatency) / float64(iterations) / float64(time.Millisecond)
    s.ChannelMessageEdit(m.ChannelID, msg.ID, fmt.Sprintf("Pong! (Avg: `%.2f` ms)", avgLatency))
}

func safeArg(totalLatency time.Duration, i int) float64 {
	if i == 0 {
		return 0
	}
	return float64(totalLatency) / float64(i) / float64(time.Millisecond)
}
