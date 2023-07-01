package campaign

import (
	"time"

	"github.com/irfan-arrosid/startup-aid/user"
)

type Campaign struct {
	Id               int
	UserId           int
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalMount        int
	CurrentMount     int
	Slug             string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	Id         int
	CampaignId int
	Filename   string
	IsPrimary  int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
