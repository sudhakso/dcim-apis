package handlers

import (
	"errors"

	"github.com/dcim-apis/pkg/dcim"
	"github.com/dcim-apis/pkg/models"
	"github.com/google/uuid"
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

func (d *DBHandler) CreateProvider(p dcim.ProviderRegistrationInput) (models.ColoProvider, error) {

	var mp models.ColoProvider
	mp.Apiurl = &p.ProviderUrl
	mp.DisplayName = &p.ProviderName

	//(fake Id, TODO(Note:Sonu) replace with a secret ID from vault)
	_id := uuid.New()
	mp.CredentialsId = &_id

	result := d.DB.Create(&mp)
	return mp, result.Error
}
