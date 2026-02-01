package modals

import (
	"strconv"
	"strings"

	"github.com/Muffin-laboratory/mf/builders"
	"github.com/Muffin-laboratory/mf/loader"
	"github.com/Muffin-laboratory/mf/utils"
	"github.com/bwmarrin/discordgo"
)

func init() {
	loader.GetMFL().LoadModal(&loader.Modal{
		Parse: func(inter *builders.InteractionCreate) bool {
			data := inter.ModalSubmitData()
			customID := data.CustomID

			if data.Components[0].Type() != discordgo.ActionsRowComponent {
				return false
			}

			if !strings.HasPrefix(customID, utils.PaginationEmbedModal) {
				return false
			}

			id := utils.GetPaginationEmbedID(customID)
			userID := utils.GetPaginationEmbedUserID(id)

			if inter.Member.User.ID != userID {
				return false
			}

			if builders.GetPaginationContainer(id) == nil {
				return false
			}

			cmp := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)

			if _, err := strconv.Atoi(cmp.Value); err != nil {
				inter.Reply(&discordgo.InteractionResponseData{
					Components: []discordgo.MessageComponent{
						builders.MakeErrorContainer("The value must be number.").Build(),
					},
					Flags: discordgo.MessageFlagsEphemeral | discordgo.MessageFlagsIsComponentsV2,
				})
				return false
			}

			return true
		},
		Run: func(inter *builders.InteractionCreate) error {
			data := inter.ModalSubmitData()
			customID := data.CustomID
			id := utils.GetPaginationEmbedID(customID)
			p := builders.GetPaginationContainer(id)
			cmp := data.Components[0].(*discordgo.ActionsRow).Components[0].(*discordgo.TextInput)

			page, _ := strconv.Atoi(cmp.Value)

			return p.Set(inter, page)
		},
	})
}
