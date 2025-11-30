package handler

import (
	"fmt"

	"github.com/Muffin-laboratory/mf/builders"
	"github.com/Muffin-laboratory/mf/loader"
	"github.com/bwmarrin/discordgo"
)

func InteractionCreate(s *discordgo.Session, i *discordgo.InteractionCreate) {
	go func() {
		switch i.Type {
		case discordgo.InteractionApplicationCommand:
			if err := loader.GetMFL().ChatInputRun(i.ApplicationCommandData().Name, s, i); err != nil {
				returnErr(err, s, i)
			}
		case discordgo.InteractionMessageComponent:
			if err := loader.GetMFL().ComponentRun(s, i); err != nil {
				returnErr(err, s, i)
			}
		case discordgo.InteractionModalSubmit:
			if err := loader.GetMFL().ModalRun(s, i); err != nil {
				returnErr(err, s, i)
			}
		case discordgo.InteractionApplicationCommandAutocomplete:
			if err := loader.GetMFL().ChatInputAutocomplete(i.ApplicationCommandData().Name, s, i); err != nil {
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
