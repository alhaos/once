package facory

import (
	"fmt"
	"once/internal/model"
)

type Factory struct{}

func NewFactory() *Factory {
	return &Factory{}
}

func (f Factory) NewDocument(docType string) (model.Doc, error) {
	switch docType {
	case "One":
		return model.DocOne{}, nil
	case "Two":
		return model.DocTwo{}, nil
	default:
		return nil, fmt.Errorf("invalid doc type: %s", docType)
	}
}
