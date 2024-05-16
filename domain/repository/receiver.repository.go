package repository

import "github.com/RenzoReccio/project-management.worker/domain/model"

type ReceiverRepository interface {
	InsertReceiver(in *model.Receiver) (*model.Receiver, error)
}
