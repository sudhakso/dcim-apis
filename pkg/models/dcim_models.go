package models

import (
	"github.com/dcim-apis/pkg/dcim"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ColoProvider struct {
	gorm.Model
	dcim.ColoProviderEntry
}

func (cp *ColoProvider) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	cp.ColoProviderEntry.Id = uuid.New()
	return
}

type Deployment struct {
	gorm.Model
	dcim.DeploymentEntry
}

type Device struct {
	gorm.Model
	dcim.DeviceEntry
}

type Rack struct {
	gorm.Model
	dcim.RackEntry
}
