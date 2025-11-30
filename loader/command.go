package loader

import (
	"github.com/Muffin-laboratory/mf/builders"
	"github.com/Muffin-laboratory/mf/configs"
	"github.com/bwmarrin/discordgo"
)

type CommandFlags uint8

type Command struct {
	*discordgo.ApplicationCommand
	Flags        CommandFlags
	Run          run
	Autocomplete run
}

const (
	CommandFlagsIsDeveloperOnlyCommand CommandFlags = 1 << iota
)

func (m *MFL) LoadCommand(command *Command) {
	defer commandMutex.Unlock()
	commandMutex.Lock()
	m.Commands[command.Name] = command
}

func (m *MFL) ChatInputRun(name string, s *discordgo.Session, inter *discordgo.InteractionCreate) error {
	i := &builders.InteractionCreate{
		InteractionCreate: inter,
		Session:           s,
		Options:           builders.MakeCommandInteractionOptionsMap(inter.ApplicationCommandData().Options),
	}

	i.InteractionCreate.User = builders.GetInteractionUser(inter)

	if command, ok := m.Commands[name]; ok {
		if command.Flags&CommandFlagsIsDeveloperOnlyCommand != 0 && i.User.ID != configs.GetConfig().Bot.OwnerID {
			return builders.NewMessageSender(i).
				AddComponents(builders.MakeDeclineContainer("This command is developer only command.")).
				SetComponentsV2(true).
				SetEphemeral(true).
				Send()
		}

		return command.Run(i)

	}
	return nil
}

func (m *MFL) ChatInputAutocomplete(name string, s *discordgo.Session, inter *discordgo.InteractionCreate) error {
	i := &builders.InteractionCreate{
		InteractionCreate: inter,
		Session:           s,
	}

	i.InteractionCreate.User = builders.GetInteractionUser(inter)

	if command, ok := m.Commands[name]; ok {
		return command.Autocomplete(i)
	}

	return nil
}
