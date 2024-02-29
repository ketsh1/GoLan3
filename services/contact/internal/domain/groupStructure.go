package domain

import "fmt"

// Group represents a group of contacts.
type Group struct {
	ID int64 `db:"id"`

	Name string `db:"name"`
}

// Validate validates the group's data.
func (g *Group) Validate() error {
	if g.Name == "" {
		return fmt.Errorf("name is required")
	}

	if len(g.Name) > 250 {
		return fmt.Errorf("name must be less than 251 characters")
	}

	return nil
}
