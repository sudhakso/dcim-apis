package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/dcim-apis/pkg/dcim"
	"github.com/dcim-apis/pkg/models"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type DcimServer struct {
	H *DBHandler
}

func (d DcimServer) GetManagedDeployments(w http.ResponseWriter, r *http.Request, params dcim.GetManagedDeploymentsParams) {
	// Get all Deployments filtered by a location Id
	// (GET /api/v1alpha1/deployments/)

}

func (d DcimServer) GetManagedDeploymentById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// Get a DriftDeployment by Id
	// (GET /api/v1alpha1/deployments/{id}/)

}

func (d DcimServer) ListLocations(w http.ResponseWriter, r *http.Request, params dcim.ListLocationsParams) {
	// Get list of location for a specific provider
	// (GET /api/v1alpha1/locations/)

}
func (d DcimServer) GetManagedDevicesById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// Get managed devices in a location by Id
	// (GET /api/v1alpha1/locations/{id}/devices/)

}
func (d DcimServer) GetManagedRacksById(w http.ResponseWriter, r *http.Request, id openapi_types.UUID) {
	// Get managed Racks by a location Id
	// (GET /api/v1alpha1/locations/{id}/racks/)

}

func (d DcimServer) ListProviders(w http.ResponseWriter, r *http.Request, params dcim.ListProvidersParams) {
	// Get list of supported provider and regions
	// (GET /api/v1alpha1/providers/)

	var p []models.ColoProvider
	d.H.GetProvider(&p)

	pl := serializeProviderList(p)

	// Rest HTTP stuffs
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(pl)
}

func (d DcimServer) RegisterProvider(w http.ResponseWriter, r *http.Request) {
	// Configure a Colo provider
	// (POST /api/v1alpha1/providers/)

	var provInput dcim.ProviderRegistrationInput
	err := json.NewDecoder(r.Body).Decode(&provInput)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Create and serialize a provider
	var p models.ColoProvider
	p, err = d.H.CreateProvider(provInput)

	// Prepare the response
	var pe dcim.ColoProviderEntry
	pe = deserializeProvider(p)

	// Rest HTTP stuffs
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pe)
}

func (d DcimServer) GetAllInventory(w http.ResponseWriter, r *http.Request, params dcim.GetAllInventoryParams) {
	// Get all inventory
	// (GET /api/v1alpha1/services/inventory/actions/get-all/)

}
