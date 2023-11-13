package handlers

import (
	"errors"

	"github.com/dcim-apis/pkg/models"
	"gorm.io/gorm"
)

type DBHandler struct {
	DB *gorm.DB
}

func (d *DBHandler) GetProvider(p *[]models.ColoProvider) error {
	// Dummy
	if result := d.DB.Find(&p); result.Error != nil {
		return errors.New("fetching providers from the database resulted in an error")
	}
	return nil
}
