package repository

import (
	"context"

	"architecture_go/services/contact/internal/domain"
)

// ContactRepository defines the repository for the contact domain.
type ContactRepository interface {
	// CRUD operations for contacts
	Create(ctx context.Context, contact *domain.Contact) error
	Get(ctx context.Context, id int64) (*domain.Contact, error)
	Update(ctx context.Context, contact *domain.Contact) error
	Delete(ctx context.Context, id int64) error

	// Additional methods based on Use Case requirements
	// ...
}

// GroupRepository defines the repository for the group domain.
type GroupRepository interface {
	// ...
}
