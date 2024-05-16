package mongoInfraestructure

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type ReceiverService struct{}

func NewReceiverService() repository.ReceiverRepository {
	return &ReceiverService{}
}

func (u *ReceiverService) InsertReceiver(in *model.Receiver) (*model.Receiver, error) {
	return in, nil
}
