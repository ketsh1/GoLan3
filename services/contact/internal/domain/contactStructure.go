package domain

import (
	"fmt"
	"strings"
)

// Contact represents a contact.
type Contact struct {
	ID int64 `db:"id"`

	// FullName is the read-only field for the contact's full name.
	FullName string `db:"-"`

	// FirstName is the first name of the contact.
	FirstName string `db:"first_name"`

	// LastName is the last name of the contact.
	LastName string `db:"last_name"`

	// MiddleName is the middle name of the contact.
	MiddleName string `db:"middle_name"`

	// PhoneNumber is the phone number of the contact.
	// It should only contain digits.
	PhoneNumber string `db:"phone_number"`
}

// GetFullName returns the full name of the contact.
func (c *Contact) GetFullName() string {
	if c.FullName != "" {
		return c.FullName
	}

	parts := []string{c.FirstName, c.MiddleName, c.LastName}
	trimmedParts := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			trimmedParts = append(trimmedParts, trimmed)
		}
	}

	return strings.Join(trimmedParts, " ")
}

// SetPhoneNumber sets the phone number of the contact,
// ensuring it only contains digits.
func (c *Contact) SetPhoneNumber(phoneNumber string) {
	c.PhoneNumber = ""
	for _, char := range phoneNumber {
		if char >= '0' && char <= '9' {
			c.PhoneNumber += string(char)
		}
	}
}

// Validate validates the contact's data.
func (c *Contact) Validate() error {
	if c.FirstName == "" {
		return fmt.Errorf("first name is required")
	}

	if c.LastName == "" {
		return fmt.Errorf("last name is required")
	}

	if c.PhoneNumber == "" {
		return fmt.Errorf("phone number is required")
	}

	return nil
}
