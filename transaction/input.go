package transaction

import "github.com/irfan-arrosid/startup-aid/user"

type GetCampaignTransactionInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}
