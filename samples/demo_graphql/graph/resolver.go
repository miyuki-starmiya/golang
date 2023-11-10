package graph

import (
	"demo_graphql/graph/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
    EntryService services.EntryService
}

func NewResolver() *Resolver {
    return &Resolver{
        EntryService: services.NewEntryService(),
    }
}
