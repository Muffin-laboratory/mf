package builders

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

type TextDisplay struct {
	*discordgo.TextDisplay
}

func TextDisplayBuilder(format string, a ...any) *TextDisplay {
	return &TextDisplay{
		TextDisplay: &discordgo.TextDisplay{
			Content: fmt.Sprintf(format, a...),
		},
	}
}

func (t *TextDisplay) SetText(format string, a ...any) *TextDisplay {
	t.TextDisplay.Content = fmt.Sprintf(format, a...)
	return t
}

func (t *TextDisplay) Build() discordgo.MessageComponent {
	return t.TextDisplay
}
