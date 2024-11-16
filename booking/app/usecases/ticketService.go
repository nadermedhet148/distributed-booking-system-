package usecases

import (
	"errors"
	"sync"

	"github.com/coroo/go-starter/app/entity"
	"github.com/coroo/go-starter/app/repositories"
	"github.com/coroo/go-starter/config"

	"github.com/go-redsync/redsync/v4"
)

type TicketService interface {
	SaveTicket(Ticket entity.Ticket) (int, error)
	SaveTicketWithLock(Ticket entity.Ticket) (int, error)
	SaveTicketWithDLock(Ticket entity.Ticket) (int, error)
	GetAllTickets() []entity.Ticket
}

type ticketService struct {
	repositories repositories.TicketRepository
	mutexes      sync.Map
}

func NewTicketService(repository repositories.TicketRepository) TicketService {
	return &ticketService{
		repositories: repository,
		mutexes:      sync.Map{},
	}
}

func (usecases *ticketService) GetAllTickets() []entity.Ticket {
	return usecases.repositories.GetAllTickets()
}

func (usecases *ticketService) SaveTicket(ticket entity.Ticket) (int, error) {
	var createdTicket = usecases.repositories.GetTicketBySlotId(ticket.SLOT_ID)
	if createdTicket != (entity.Ticket{}) {
		return 0, errors.New("slot already taken")
	}
	return usecases.repositories.Save(ticket)
}

func (usecases *ticketService) SaveTicketWithLock(ticket entity.Ticket) (int, error) {

	key := ticket.REF_NUMBER
	mutexInterface, _ := usecases.mutexes.LoadOrStore(key, &sync.Mutex{})
	mutex := mutexInterface.(*sync.Mutex)

	mutex.Lock()
	id, err := usecases.SaveTicket(ticket)
	defer mutex.Unlock()
	return id, err
}

func (usecases *ticketService) SaveTicketWithDLock(ticket entity.Ticket) (int, error) {

	rs := redsync.New(config.ConnectRedis())

	mutexname := "t-" + ticket.REF_NUMBER + "-mutex"
	mutex := rs.NewMutex(mutexname)

	if err := mutex.Lock(); err != nil {
		return 0, errors.New("lock failed")
	}
	id, err := usecases.SaveTicket(ticket)
	if ok, err := mutex.Unlock(); !ok || err != nil {
		return 0, errors.New("unlock failed")
	}
	return id, err
}
