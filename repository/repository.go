package repository

import (
	"github.com/chingkamhing/mongodb-seed/config"
	"github.com/chingkamhing/mongodb-seed/repository/mongodb"
)

// defaultOptions get the default mongodb options
func defaultOptions() []mongodb.MongodbOption {
	return []mongodb.MongodbOption{
		mongodb.OptionSetHost(config.Config.Database.Host),
		mongodb.OptionSetPort(config.Config.Database.Port),
		mongodb.OptionSetDatabaseName(config.Config.Database.Dbname),
		mongodb.OptionSetUser(config.Config.Database.Username),
		mongodb.OptionSetPassword(config.Config.Database.Password),
	}
}

// New return pointer of repository interface
func New() (repoInterface Interface) {
	return mongodb.New(defaultOptions()...)
}

// NewWithMigrateUp return pointer of repository interface
func NewWithMigrateUp() (repoInterface Interface) {
	options := append(defaultOptions(), mongodb.OptionSetMigrateUp())
	return mongodb.New(options...)
}
