package service

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ServicePackage is the core interface that all service packages must implement.
type ServicePackage interface {
	// Configure is called during provider configuration.
	Configure(context.Context) error

	// DataSources returns a map of the data sources implemented in this service package.
	DataSources() map[string]*schema.Resource

	// ID returns the service package's ID.
	// This ID must be unique.
	ID() string

	// Resources returns a map of the resources implemented in this service package.
	Resources() map[string]*schema.Resource
}
