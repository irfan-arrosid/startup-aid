package transaction

type Service interface {
	GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error) {
	transactions, err := s.repository.FindByCampaignId(input.Id)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
