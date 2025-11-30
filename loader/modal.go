package loader

type Modal struct {
	Parse parse
	Run   run
}

func (m *MFL) LoadModal(modal *Modal) {
	defer modalMutex.Unlock()
	modalMutex.Lock()
	m.Modals = append(m.Modals, modal)
}
