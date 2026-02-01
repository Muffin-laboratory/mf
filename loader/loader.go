package loader

import (
	"sync"

	"github.com/Muffin-laboratory/mf/builders"
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
var once sync.Once

func GetMFL() *MFL {
	once.Do(func() {
		instance = &MFL{
			Commands: make(map[string]*Command),
		}
	})

	return instance
}
