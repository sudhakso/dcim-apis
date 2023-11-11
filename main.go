package main

import (
	"net/http"

	"github.com/dcim-apis/pkg/dcim"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

type DcimServer struct {
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

}

func (d DcimServer) RegisterProvider(w http.ResponseWriter, r *http.Request) {
	// Configure a Colo provider
	// (POST /api/v1alpha1/providers/)

}

func (d DcimServer) GetAllInventory(w http.ResponseWriter, r *http.Request, params dcim.GetAllInventoryParams) {
	// Get all inventory
	// (GET /api/v1alpha1/services/inventory/actions/get-all/)

}

func main() {
	s := DcimServer{}
	h := dcim.Handler(s)

	http.ListenAndServe(":3000", h)
}