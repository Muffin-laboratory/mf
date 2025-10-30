package commands

import (
	"fmt"

	"github.com/Muffin-laboratory/mf/builders"
	"github.com/bwmarrin/discordgo"
)

var PingCommand = &Command{
	ApplicationCommand: &discordgo.ApplicationCommand{
		Name:        "ping",
		Description: "Check the bot's latency",
	},
	Run: func(inter *builders.InteractionCreate) error {
		title := "### 🏓 Pong!"

		if err := builders.NewMessageSender(inter).
			AddComponents(builders.ContainerBuilder().AddText(title).AddText("- Calculating latency...")).
			SetComponentsV2(true).
			Send(); err != nil {
			return err
		}

		message, err := inter.FetchReply()
		if err != nil {
			return err
		}

		createdTimestamp, _ := discordgo.SnowflakeTimestamp(inter.ID)
		discordPing := message.Timestamp.Sub(createdTimestamp).Milliseconds()

		return builders.NewMessageSender(inter).
			AddComponents(
				builders.ContainerBuilder().
					AddText(title).
					AddText(fmt.Sprintf("- **Discord latency:** `%d`ms", discordPing)),
			).
			SetComponentsV2(true).
			Send()
	},
}

func init() {
	GetDiscommand().LoadCommand(PingCommand)
}
