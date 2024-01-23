package model

import mon "go-zero-yun/pkg/monkit"

const DemoCollectionName = "demo"

var _ DemoModel = (*customDemoModel)(nil)

type (
	// DemoModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDemoModel.
	DemoModel interface {
		demoModel
	}

	customDemoModel struct {
		*defaultDemoModel
	}
)

// NewDemoModel returns a model for the mongo.
func NewDemoModel() DemoModel {
	db := MonDb()
	conn := mon.MustNewModel(db, DemoCollectionName)
	return &customDemoModel{
		defaultDemoModel: newDefaultDemoModel(conn),
	}
}
