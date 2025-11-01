package builders

import (
	"github.com/bwmarrin/discordgo"
)

type Container struct {
	*discordgo.Container
}

func ContainerBuilder() *Container {
	return &Container{
		Container: &discordgo.Container{},
	}
}

func (c *Container) SetAccentColor(color int) *Container {
	c.Container.AccentColor = &color
	return c
}

func (c *Container) SetSpoiler(spoiler bool) *Container {
	c.Container.Spoiler = spoiler
	return c
}

func (c *Container) AddComponents(components ...ComponentBuilder) *Container {
	for _, cmp := range components {
		c.Container.Components = append(c.Container.Components, cmp.Build())
	}
	return c
}

func (c *Container) AddText(text string) *Container {
	c.AddComponents(TextDisplayBuilder(text))
	return c
}

func (c *Container) Build() discordgo.MessageComponent {
	return c.Container
}

func MakeErrorContainer(text string) *Container {
	return ContainerBuilder().
		AddComponents(
			TextDisplayBuilder("### ❌ Error"),
			TextDisplayBuilder(text),
		)
}

func MakeDeclineContainer(text string) *Container {
	return ContainerBuilder().
		AddComponents(
			TextDisplayBuilder("### ❌ Declined"),
			TextDisplayBuilder(text),
		)
}

func MakeCanceledContainer(text string) *Container {
	return ContainerBuilder().
		AddComponents(
			TextDisplayBuilder("### ❌ Canceled"),
			TextDisplayBuilder(text),
		)
}

func MakeSuccessContainer(text string) *Container {
	return ContainerBuilder().
		AddComponents(
			TextDisplayBuilder("### ✅ Success"),
			TextDisplayBuilder(text),
		)
}
