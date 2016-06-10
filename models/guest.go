package models

// Guester is a collection of methods that can be used for the Guest struct
type Guester interface {
	SetGuest(g Guest) error
	GetGuestByID(id string) (*Guest, error)
	GetGuestByName(id string) (*Guest, error)
	GetGuestAll() ([]Guest, error)
	RmGuestAll() error
	RmGuestByID(id string) error
}

// Guest represents a single guest
type Guest struct {
	Name    string
	Wedding bool
	Dinner  bool
}

// SetGuest will save a single Guest in the database
func (db DB) SetGuest(g Guest) error {
	return nil
}

// GetGuestByID will get a single Guest from the database
func (db DB) GetGuestByID(id string) (*Guest, error) {
	return &Guest{}, nil
}

// GetGuestByName will get a single Guest from the database
func (db DB) GetGuestByName(id string) (*Guest, error) {
	return &Guest{}, nil
}

// GetGuestAll returns a slice of all Guests from the database
func (db DB) GetGuestAll() ([]Guest, error) {
	return []Guest{}, nil
}

// RmGuestAll clears all Guests from the database
func (db DB) RmGuestAll() error {
	return nil
}

// RmGuestByID clears a single Guest from the database
func (db DB) RmGuestByID(id string) error {
	return nil
}
