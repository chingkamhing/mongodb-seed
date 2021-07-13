package mongodb

import (
	"strconv"
	"time"
)

// driverOptions is a structure that hold all the options to new a mongodb driver
type driverOptions struct {
	host                    string        // database url string
	port                    string        // database port number in string
	dbname                  string        // database name
	user                    string        // database user login name
	password                string        // database password
	timeout                 time.Duration // database connection timeout time
	retryCount              int           // database connection retry number of times
	retryInterval           time.Duration // database connection retry interval: (retryIndex + 1) * retryInterval
	migrationPath           string        // where the migration scheme files locate
	migrationCollectionName string        // migration collection name
	migrateUp               bool          // database migrate up
}

// mongodb default settings
var defaultOptions = driverOptions{
	host:                    "127.0.0.1",
	port:                    "27017",
	timeout:                 10 * time.Second,
	retryCount:              20,
	retryInterval:           3 * time.Second,
	migrationPath:           "deploy/migrations",
	migrationCollectionName: "schema_migrations", // note: migrate's default collection name is "schema_migrations", better follow it
}

// MongodbOption control driverOptions behavior
type MongodbOption func(*driverOptions)

// New create new mongodb driver
func (storer *StorerMongodb) new(optionFunctions ...MongodbOption) {
	storer.options = defaultOptions
	for _, function := range optionFunctions {
		function(&storer.options)
	}
}

// OptionSetHost set database host url to options
func OptionSetHost(host string) MongodbOption {
	return func(options *driverOptions) {
		options.host = host
	}
}

// OptionSetPort set database port number to options
func OptionSetPort(port int) MongodbOption {
	return func(options *driverOptions) {
		options.port = strconv.Itoa(port)
	}
}

// OptionSetDatabaseName set database name to options
func OptionSetDatabaseName(databaseName string) MongodbOption {
	return func(options *driverOptions) {
		options.dbname = databaseName
	}
}

// OptionSetUser set database user name to options
func OptionSetUser(user string) MongodbOption {
	return func(options *driverOptions) {
		options.user = user
	}
}

// OptionSetPassword set database password to options
func OptionSetPassword(password string) MongodbOption {
	return func(options *driverOptions) {
		options.password = password
	}
}

// OptionSetTimeout set database connection timeout in seconds to options
func OptionSetTimeout(seconds int) MongodbOption {
	return func(options *driverOptions) {
		options.timeout = time.Duration(seconds) * time.Second
	}
}

// OptionSetRetryCount set database connection retry number of times
func OptionSetRetryCount(count int) MongodbOption {
	return func(options *driverOptions) {
		options.retryCount = count
	}
}

// OptionSetSetRetryInterval set database connection retry interval: (RetryIndex + 1) * RetryInterval
func OptionSetSetRetryInterval(seconds int) MongodbOption {
	return func(options *driverOptions) {
		options.retryInterval = time.Duration(seconds) * time.Second
	}
}

// OptionSetSetMigrationPath set migration relative path
func OptionSetSetMigrationPath(path string) MongodbOption {
	return func(options *driverOptions) {
		options.migrationPath = path
	}
}

// OptionSetSetMigrationCollectionName set migration collection name
func OptionSetSetMigrationCollectionName(name string) MongodbOption {
	return func(options *driverOptions) {
		options.migrationCollectionName = name
	}
}

// OptionSetMigrateUp perform database migrate up
func OptionSetMigrateUp() MongodbOption {
	return func(options *driverOptions) {
		options.migrateUp = true
	}
}
