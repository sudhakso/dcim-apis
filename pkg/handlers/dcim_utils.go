package handlers

import (
	"github.com/dcim-apis/pkg/dcim"
	"github.com/dcim-apis/pkg/models"
)

func serializeProviderList(pm []models.ColoProvider) dcim.ProviderList {
	var pl dcim.ProviderList

	for _, j := range pm {
		pe := j.ColoProviderEntry
		pe.CreatedAt = &j.Model.CreatedAt

		//append
		pl.Items = append(pl.Items, pe)
	}
	return pl
}

func deserializeProvider(p models.ColoProvider) dcim.ColoProviderEntry {
	pe := p.ColoProviderEntry
	return pe
}
