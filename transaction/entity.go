package transaction

import (
	"time"

	"github.com/irfan-arrosid/startup-aid/campaign"
	"github.com/irfan-arrosid/startup-aid/user"
)

type Transaction struct {
	Id         int
	CampaignId int
	UserId     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
