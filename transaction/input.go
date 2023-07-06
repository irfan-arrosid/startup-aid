package transaction

type GetCampaignTransactionInput struct {
	Id int `uri:"id" binding:"required"`
}
