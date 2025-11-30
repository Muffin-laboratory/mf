package components

import (
	"strings"

	"github.com/Muffin-laboratory/mf/builders"
	"github.com/Muffin-laboratory/mf/loader"
	"github.com/Muffin-laboratory/mf/utils"
	"github.com/bwmarrin/discordgo"
)

func init() {
	loader.GetMFL().LoadComponent(&loader.Component{
		Parse: func(inter *builders.InteractionCreate) bool {
			if inter.MessageComponentData().ComponentType == discordgo.ButtonComponent {
				customID := inter.MessageComponentData().CustomID

				isPrev := strings.HasPrefix(customID, utils.PaginationEmbedPrev)
				isNext := strings.HasPrefix(customID, utils.PaginationEmbedNext)
				isSetPage := strings.HasPrefix(customID, utils.PaginationEmbedPages)
				if !isPrev && !isNext && !isSetPage {
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
			} else {
				return false
			}
			return true
		},
		Run: func(inter *builders.InteractionCreate) error {
			customID := inter.MessageComponentData().CustomID
			id := utils.GetPaginationEmbedID(customID)
			p := builders.GetPaginationContainer(id)

			if strings.HasPrefix(customID, utils.PaginationEmbedPrev) {
				return p.Prev(inter)
			} else if strings.HasPrefix(customID, utils.PaginationEmbedNext) {
				return p.Next(inter)
			} else {
				return p.ShowModal(inter)
			}
		},
	})
}
