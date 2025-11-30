package loader

import (
	"sync"

	"github.com/Muffin-laboratory/mf/builders"
	"github.com/bwmarrin/discordgo"
)

type run func(inter *builders.InteractionCreate) error
type parse func(inter *builders.InteractionCreate) bool

type MFL struct {
	Commands   map[string]*Command
	Components []*Component
	Modals     []*Modal
}

var (
	commandMutex   sync.Mutex
	componentMutex sync.Mutex
	modalMutex     sync.Mutex
)

var instance *MFL

func GetMFL() *MFL {
	if instance == nil {
		instance = &MFL{
			Commands:   map[string]*Command{},
			Components: []*Component{},
			Modals:     []*Modal{},
		}
	}

	return instance
}

func (m *MFL) ModalRun(s *discordgo.Session, inter *discordgo.InteractionCreate) error {
	var err error

	i := &builders.InteractionCreate{
		InteractionCreate: inter,
		Session:           s,
	}

	for _, m := range m.Modals {
		if !m.Parse(i) {
			continue
		}

		err = m.Run(i)
		break
	}
	return err
}
