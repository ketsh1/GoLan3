package useCase

import (
	"context"

	"architecture_go/services/contact/internal/domain"
)

// ContactUseCase defines the use cases for the contact domain.
type ContactUseCase interface {
	// Create creates a new contact.
	Create(ctx context.Context, contact *domain.Contact) error

	// Get retrieves a contact by ID.
	Get(ctx context.Context, id int64) (*domain.Contact, error)

	// Update updates an existing contact.
	Update(ctx context.Context, contact *domain.Contact) error

	// Delete deletes a contact by ID.
	Delete(ctx context.Context, id int64) error
}

func (uc *contactUseCase) Create(ctx context.Context, contact *domain.Contact) error {
	// ... validate contact
	// ... generate ID
	// ... save to database
	return nil
}

func (uc *contactUseCase) Get(ctx context.Context, id int64) (*domain.Contact, error) {

	return nil, nil
}

func (uc *contactUseCase) Update(ctx context.Context, contact *domain.Contact) error {

	return nil
}

func (uc *contactUseCase) Delete(ctx context.Context, id int64) error {

	return nil
}
