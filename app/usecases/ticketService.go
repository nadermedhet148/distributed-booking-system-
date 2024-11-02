package usecases

import (
	"github.com/coroo/go-starter/app/entity"
	"github.com/coroo/go-starter/app/repositories"
	_ "github.com/go-sql-driver/mysql"
)

type TicketService interface {
	SaveTicket(entity.Ticket) (int, error)
	GetAllTickets() []entity.Ticket
}

type ticketService struct {
	repositories repositories.TicketRepository
}

func NewticketService(repository repositories.TicketRepository) TicketService {
	return &ticketService{
		repositories: repository,
	}
}

func (usecases *ticketService) GetAllTickets() []entity.Ticket {
	return usecases.repositories.GetAllTickets()
}

func (usecases *ticketService) SaveTicket(ticket entity.Ticket) (int, error) {
	return usecases.repositories.SaveTicket(ticket)
}
