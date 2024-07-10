package factory_workitem

import (
	"context"
	"encoding/json"
	"fmt"

	application_sendepic "github.com/RenzoReccio/project-management.worker/application/epic/event/send-epic"
	application_getepic "github.com/RenzoReccio/project-management.worker/application/epic/query/get-epic"
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/mehdihadeli/go-mediatr"
)

type EpicFactory struct {
}

func NewEpicFactory() IWorkItem {
	return &EpicFactory{}
}

func (c EpicFactory) ExecuteWorkItem(context context.Context, resourceURL string) {
	resultEpic, _ := mediatr.Send[*application_getepic.GetEpicQuery, *model_shared.ResultWithValue[model.Epic]](context, &application_getepic.GetEpicQuery{ResourceURL: resourceURL})
	if !resultEpic.IsSuccess {
		return
	}
	user := resultEpic.Result()
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	epicSendEvent := &application_sendepic.SendEpicEvent{Data: resultEpic.Result()}
	mediatr.Publish(context, epicSendEvent)
}
