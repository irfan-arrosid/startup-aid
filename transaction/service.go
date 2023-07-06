package transaction

import (
	"errors"
	"strconv"

	"github.com/irfan-arrosid/startup-aid/campaign"
	"github.com/irfan-arrosid/startup-aid/payment"
)

type Service interface {
	GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

type service struct {
	repository         Repository
	campaignRepository campaign.Repository
	paymentService     payment.Service
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionByCampaignId(input GetCampaignTransactionInput) ([]Transaction, error) {
	campaign, err := s.campaignRepository.FindById(input.Id)
	if err != nil {
		return []Transaction{}, err
	}

	if campaign.UserId != input.User.Id {
		return []Transaction{}, errors.New("not an owner of the campaign")
	}

	transactions, err := s.repository.FindByCampaignId(input.Id)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) GetTransactionByUserId(userId int) ([]Transaction, error) {
	transactions, err := s.repository.FindByUserId(userId)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{}
	transaction.CampaignId = input.CampaignId
	transaction.Amount = input.Amount
	transaction.UserId = input.User.Id
	transaction.Status = "pending"

	newTransaction, err := s.repository.Save(transaction)
	if err != nil {
		return newTransaction, err
	}

	paymentTransacation := payment.Transaction{
		Id:     newTransaction.Id,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransacation, input.User)
	if err != nil {
		return newTransaction, err
	}

	newTransaction.PaymentURL = paymentURL

	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil {
		return newTransaction, err
	}

	return newTransaction, nil
}

func (s *service) ProcessPayment(input TransactionNotificationInput) error {
	transactionId, _ := strconv.Atoi(input.OrderId)

	transaction, err := s.repository.FindById(transactionId)
	if err != nil {
		return err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaign, err := s.campaignRepository.FindById(updatedTransaction.CampaignId)
	if err != nil {
		return err
	}

	if updatedTransaction.Status == "paid" {
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil {
			return err
		}
	}

	return nil
}
