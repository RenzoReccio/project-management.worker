package factory_workitem

import (
	"context"
	"encoding/json"
	"fmt"

	application_sendfeature "github.com/RenzoReccio/project-management.worker/application/feature/event/send-feature"
	application_getfeature "github.com/RenzoReccio/project-management.worker/application/feature/query/get-feature"
	"github.com/RenzoReccio/project-management.worker/domain/model"
	model_shared "github.com/RenzoReccio/project-management.worker/domain/model/shared"
	"github.com/mehdihadeli/go-mediatr"
)

type FeatureFactory struct {
}

func NewFeatureFactory() IWorkItem {
	return &FeatureFactory{}
}

func (c FeatureFactory) ExecuteWorkItem(context context.Context, resourceURL string) {
	resultFeature, _ := mediatr.Send[*application_getfeature.GetFeatureQuery, *model_shared.ResultWithValue[model.Feature]](context, &application_getfeature.GetFeatureQuery{ResourceURL: resourceURL})
	if !resultFeature.IsSuccess {
		fmt.Print(resultFeature.Error)
		return
	}
	foo_marshalled, _ := json.Marshal(resultFeature.Result())
	fmt.Print(string(foo_marshalled))
	featureSendEvent := &application_sendfeature.SendFeatureEvent{Data: resultFeature.Result()}
	mediatr.Publish(context, featureSendEvent)
}
