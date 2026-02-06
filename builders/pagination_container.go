package builders

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Muffin-laboratory/mf/utils"
	"github.com/bwmarrin/discordgo"
)

// PaginationContainer is container with page
type PaginationContainer struct {
	Containers []*Container
	Current    int
	Total      int
	ID         string
	m          any
	timer      *time.Timer
}

var paginationContainers = make(map[string]*PaginationContainer)

const endDuration = 10 * time.Minute

// PaginationContainerBuilder creates a new PaginationContainer
func PaginationContainerBuilder(m any) *PaginationContainer {
	var userID string

	switch m := m.(type) {
	case *MessageCreate:
		userID = m.Author.ID
	case *InteractionCreate:
		userID = m.Member.User.ID
	}

	id := fmt.Sprintf("%s/%d", userID, rand.Intn(100))
	return &PaginationContainer{
		Current: 1,
		ID:      id,
		m:       m,
		timer:   time.NewTimer(endDuration),
	}
}

func (p *PaginationContainer) waitTimerEnd() {
	<-p.timer.C
	delete(paginationContainers, p.ID)
}

func (p *PaginationContainer) resetTimer() {
	p.timer.Reset(endDuration)
}

// AddContainers adds containers
func (p *PaginationContainer) AddContainers(containers ...*Container) *PaginationContainer {
	p.Total += len(containers)
	p.Containers = append(p.Containers, containers...)
	return p
}

// Start starts the paginated-container
func (p *PaginationContainer) Start() error {
	container := *p.Containers[0]
	container.AddComponents(makeComponents(p.ID, p.Current, p.Total))
	paginationContainers[p.ID] = p

	go p.waitTimerEnd()

	return NewMessageSender(p.m).
		AddComponents(&container).
		SetReply(true).
		SetEphemeral(true).
		SetComponentsV2(true).
		Send()
}

func makeComponents(id string, current, total int) *ActionsRow {
	disabled := false

	if total == 1 {
		disabled = true
	}

	return ActionsRowBuilder(
		ButtonBuilder().
			SetStyle(discordgo.PrimaryButton).
			SetLabel("Previous").
			SetCustomID(utils.MakePaginationEmbedPrev(id)).
			SetDisabled(disabled),
		ButtonBuilder().
			SetStyle(discordgo.SecondaryButton).
			SetLabel(fmt.Sprintf("(%d/%d)", current, total)).
			SetCustomID(utils.MakePaginationEmbedPages(id)).
			SetDisabled(disabled),
		ButtonBuilder().
			SetStyle(discordgo.PrimaryButton).
			SetLabel("Next").
			SetCustomID(utils.MakePaginationEmbedNext(id)).
			SetDisabled(disabled),
	)
}

// GetPaginationContainer gets PaginationContainer
func GetPaginationContainer(id string) *PaginationContainer {
	if p, ok := paginationContainers[id]; ok {
		return p
	}
	return nil
}

// Prev move to previous page
func (p *PaginationContainer) Prev(i *InteractionCreate) error {
	if p.Current == 1 {
		p.Current = p.Total
	} else {
		p.Current -= 1
	}

	return p.Set(i, p.Current)
}

// Next moves to next page
func (p *PaginationContainer) Next(i *InteractionCreate) error {
	if p.Current >= p.Total {
		p.Current = 1
	} else {
		p.Current += 1
	}

	return p.Set(i, p.Current)
}

// Set sets to page
func (p *PaginationContainer) Set(i *InteractionCreate, page int) error {
	p.resetTimer()

	if page <= 0 {
		p.Current = 1
	} else if page > p.Total {
		p.Current = p.Total
	} else {
		p.Current = page
	}

	container := *p.Containers[p.Current-1]
	container.AddComponents(makeComponents(p.ID, p.Current, p.Total))

	return i.Update(&discordgo.InteractionResponseData{
		Flags:      discordgo.MessageFlagsIsComponentsV2,
		Components: []discordgo.MessageComponent{container.Build()},
	})
}

// ShowModal show discord's modal
func (p *PaginationContainer) ShowModal(i *InteractionCreate) error {
	return i.ShowModal(
		ModalBuilder().
			SetCustomID(utils.MakePaginationEmbedModal(p.ID)).
			SetTitle("Set Page").
			AddComponents(
				LabelBuilder().
					SetLabel("Page number").
					SetDescription("Put the page number you want to move.").
					SetComponent(
						TextInputBuilder().
							SetCustomID(utils.MakePaginationEmbedSetPage(p.ID)).
							SetStyle(discordgo.TextInputShort).
							SetPlaceholder("Set page number...").
							SetValue(fmt.Sprint(p.Current)).
							SetRequired(true),
					),
			),
	)
}
