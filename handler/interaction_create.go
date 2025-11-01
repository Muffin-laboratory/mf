package handler

import (
	"fmt"

	"github.com/Muffin-laboratory/mf/builders"
	"github.com/Muffin-laboratory/mf/commands"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if err := commands.GetDiscommand().ChatInputRun(i.ApplicationCommandData().Name, s, i); err != nil {
				returnErr(err, s, i)
			}
		case discordgo.InteractionMessageComponent:
			if err := commands.GetDiscommand().ComponentRun(s, i); err != nil {
				returnErr(err, s, i)
			}
		case discordgo.InteractionModalSubmit:
			if err := commands.GetDiscommand().ModalRun(s, i); err != nil {
				returnErr(err, s, i)
			}
		case discordgo.InteractionApplicationCommandAutocomplete:
			if err := commands.GetDiscommand().ChatInputAutocomplete(i.ApplicationCommandData().Name, s, i); err != nil {
				returnErr(err, s, i)
			}
		}
	}()
}

func returnErr(err error, s *discordgo.Session, i *discordgo.InteractionCreate) {
	fmt.Println(err)
	builders.NewMessageSender(&builders.InteractionCreate{InteractionCreate: i, Session: s}).
		AddComponents(builders.MakeErrorContainer("An error occurred.")).
		SetComponentsV2(true).
		SetEphemeral(true).
		Send()
}
