package factory_workitem

import (
	"context"
)

type IWorkItem interface {
	ExecuteWorkItem(context context.Context, resourceURL string)
}
