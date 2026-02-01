package loader

import (
	"github.com/Muffin-laboratory/mf/builders"
	"github.com/bwmarrin/discordgo"
)

type Modal struct {
	Parse parse
	Run   run
}

func (m *MFL) LoadModal(modal *Modal) {
	defer modalMutex.Unlock()
	modalMutex.Lock()
	m.Modals = append(m.Modals, modal)
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
