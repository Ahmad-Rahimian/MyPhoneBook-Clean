package repository

import (
	"phonebook/internal/domain"
	pkg "phonebook/pkg/db"
)

// GetAllContacts fetches all contacts from DB.
func GetAllContacts() ([]domain.Contact, error) {
	rows, err := pkg.DB.Query("Select id,name,phone,email FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var contacts []domain.Contact

	for rows.Next() {
		var c domain.Contact
		err := rows.Scan(&c.ID, &c.Name, &c.Phone, &c.Email)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, c)
	}
	return contacts, nil
}
