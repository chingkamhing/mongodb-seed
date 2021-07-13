package repository

import (
	"github.com/chingkamhing/mongodb-seed/model"
)

// Interfaces defines the repository interfaces
type Interface interface {
	// general database interfaces

	// Get database server url
	GetURL() string
	// Open open database connection
	Open() error
	// Ping verify if the database server is active
	Ping() error
	// Close close database connection
	Close() error
	// MigrateUp perform data migration up
	MigrateUp() (err error)
	// MigrateDown perform data migration down
	MigrateDown() (err error)

	// receipt interfaces

	// CreateReceipt create a new model.Receipt document in "receipts" collection
	CreateReceipt(receipt *model.Receipt) (id string, err error)

	// ReadAllReceipts read all model.Receipt document in "receipts" collection
	ReadAllReceipts() (receipts []*model.Receipt, err error)

	// ReadReceiptByID read model.Receipt document in "receipts" collection with the specified receipt ID
	ReadReceiptByID(id string) (receipt *model.Receipt, err error)

	// UpdateReceiptByID update model.Receipt from document in "receipts" collection with the specified receipt
	UpdateReceiptByID(receipt *model.Receipt) (count int64, err error)

	// DeleteReceiptByID delete model.Receipt from document in "receipts" collection with the specified receipt ID
	DeleteReceiptByID(id string) (count int64, err error)

	// DeleteAllReceipts delete all model.Receipt document in "receipts" collection
	DeleteAllReceipts() (count int64, err error)
}
