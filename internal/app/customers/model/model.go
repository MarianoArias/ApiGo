package model

import (
	"encoding/json"
	"time"

	"github.com/MarianoArias/ApiGo/pkg/elastic-search"
	"github.com/MarianoArias/ApiGo/pkg/entity-manager"
	"github.com/olivere/elastic/v7"
)

type Customer struct {
	ID          int        `gorm:"column:id;AUTO_INCREMENT;primary_key" json:"id"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"-"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" json:"-"`
	FirstName   string     `gorm:"column:first_name" json:"firstName"`
	LastName    string     `gorm:"column:last_name" json:"lastName"`
	PhoneNumber string     `gorm:"column:phone_number" json:"phoneNumber"`
	Email       string     `gorm:"column:email" json:"email"`
}

var index string = "customers"

func init() {
	entitymanager.GetClient().AutoMigrate(&Customer{})
}

func GetCustomers() ([]Customer, int64, error) {
	fullFilter := elastic.
		NewMultiMatchQuery("test", "firstName", "lastName", "phoneNumber", "email").
		Type("phrase_prefix")

	firstNameFilter := elastic.
		NewMatchQuery("firstName", "test")

	query := elastic.NewBoolQuery().
		Should().
		Filter(fullFilter).
		Filter(firstNameFilter)

	hits, total, err := elasticsearch.GetResults(index, query, 0, 10)
	if err != nil {
		return nil, 0, err
	}

	var customers []Customer
	for _, hit := range hits {
		var customer Customer
		if err := json.Unmarshal(hit.Source, &customer); err != nil {
			return nil, 0, err
		}
		customers = append(customers, customer)
	}

	return customers, total, nil
}
