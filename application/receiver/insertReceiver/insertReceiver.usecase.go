package applicationreceiver

import (
	"github.com/RenzoReccio/project-management.worker/domain/model"
	"github.com/RenzoReccio/project-management.worker/domain/repository"
)

type InsertReceiverUseCase struct {
	receiverRepository repository.ReceiverRepository
}

func NewInsertReceiverUseCase(
	receiverRepository repository.ReceiverRepository,
) *InsertReceiverUseCase {
	return &InsertReceiverUseCase{receiverRepository: receiverRepository}
}

func (u *InsertReceiverUseCase) Process(in *InsertReceiverDto) (int, error) {
	InsertReceiver := model.NewReceiver(in.ID, in.SubscriptionID, in.EventType, in.CreatedDate, in.Resource.ID, in.Resource.URL)

	if _, err := u.receiverRepository.InsertReceiver(InsertReceiver); err != nil {
		return -1, err
	}

	return 1, nil
}
