package loader

import (
	"github.com/Muffin-laboratory/mf/builders"
	"github.com/bwmarrin/discordgo"
)

type Component struct {
	Parse parse
	Run   run
}

func (m *MFL) LoadComponent(component *Component) {
	defer componentMutex.Unlock()
	componentMutex.Lock()
	m.Components = append(m.Components, component)
}

func (m *MFL) ComponentRun(s *discordgo.Session, inter *discordgo.InteractionCreate) error {
	var err error

	i := &builders.InteractionCreate{
		InteractionCreate: inter,
		Session:           s,
	}

	i.InteractionCreate.User = builders.GetInteractionUser(inter)

	for _, c := range m.Components {
		if !c.Parse(i) {
			continue
		}

		err = c.Run(i)
		break
	}
	return err
}
