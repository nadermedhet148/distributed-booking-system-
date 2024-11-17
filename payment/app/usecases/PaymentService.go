package usecases

import (
	"github.com/coroo/go-starter/app/entity"
	"github.com/coroo/go-starter/app/repositories"
)

type PaymentService interface {
	SavePayment(Payment entity.Payment) (int, error)
}

type paymentService struct {
	repositories repositories.PaymentRepository
}

func NewTicketService(repository repositories.PaymentRepository) PaymentService {
	return &paymentService{
		repositories: repository,
	}
}

func (usecases *paymentService) SavePayment(Payment entity.Payment) (int, error) {
	return usecases.repositories.Save(Payment)
}
