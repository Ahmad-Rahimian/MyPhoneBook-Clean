package service

import (
	"phonebook/internal/domain"
	"phonebook/internal/repository"
)

// GetContacts is the use-case layer entry for listing contacts.
func GetContacts() ([]domain.Contact, error) {
	return repository.GetAllContacts()
}
