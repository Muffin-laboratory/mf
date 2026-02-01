package builders

import (
	"github.com/bwmarrin/discordgo"
)

// A Container is builder for container
type Container struct {
	container *discordgo.Container
}

// ContainerBuilder creates a new container
func ContainerBuilder() *Container {
	return &Container{
		container: &discordgo.Container{},
	}
}

// SetAccentColor sets accent color
func (c *Container) SetAccentColor(color int) *Container {
	c.container.AccentColor = &color
	return c
}

// SetSpoiler sets spoiler
func (c *Container) SetSpoiler(spoiler bool) *Container {
	c.container.Spoiler = spoiler
	return c
}

// AddComponents adds components
func (c *Container) AddComponents(components ...ComponentBuilder) *Container {
	for _, cmp := range components {
		c.container.Components = append(c.container.Components, cmp.Build())
	}
	return c
}

// AddText adds a text
func (c *Container) AddText(format string, a ...any) *Container {
	c.AddComponents(TextDisplayBuilder(format, a...))
	return c
}

// Builds returns discordgo.Container(discordgo.MessageComponent)
func (c *Container) Build() discordgo.MessageComponent {
	return c.container
}

func MakeErrorContainer(format string, a ...any) *Container {
	return ContainerBuilder().
		AddComponents(
			TextDisplayBuilder("### ❌ Error"),
			TextDisplayBuilder(format, a...),
		)
}

func MakeDeclineContainer(format string, a ...any) *Container {
	return ContainerBuilder().
		AddComponents(
			TextDisplayBuilder("### ❌ Declined"),
			TextDisplayBuilder(format, a...),
		)
}

func MakeCanceledContainer(format string, a ...any) *Container {
	return ContainerBuilder().
		AddComponents(
			TextDisplayBuilder("### ❌ Canceled"),
			TextDisplayBuilder(format, a...),
		)
}

func MakeSuccessContainer(format string, a ...any) *Container {
	return ContainerBuilder().
		AddComponents(
			TextDisplayBuilder("### ✅ Success"),
			TextDisplayBuilder(format, a...),
		)
}
