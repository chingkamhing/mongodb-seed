package mongodb

import (
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mongodb"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go.mongodb.org/mongo-driver/mongo"
)

//
// MongoDB database naming conventions
// * database
// 	 + camelCase
// 	 + append DB at the endo of the database name (i.e. xxxDB)
// 	 + singular
// * collection
// 	 + camelCase
//   + plural
// * field
// 	 + camelCase
//

// StorerMongodb define the storer structure
type StorerMongodb struct {
	options  driverOptions
	client   *mongo.Client
	database *mongo.Database
}

// New return pointer of repository interface
func New(optionFunctions ...MongodbOption) (storer *StorerMongodb) {
	storer = &StorerMongodb{}
	storer.new(optionFunctions...)
	return storer
}

//
// implement repository Interface interface
//

// Open open database connection
func (storer *StorerMongodb) Open() (err error) {
	// retry connecting to mongodb
	for i := 1; i <= storer.options.retryCount; i++ {
		// connect to database
		err = storer.open()
		if err == nil {
			// connect success, ping the connection
			err = storer.Ping()
			if err == nil {
				// ping success, break the retry
				break
			}
		}
		// connect fail
		log.Printf("fail to connect to db: try %d err %v", i, err)
		wait := time.Duration(i) * storer.options.retryInterval
		time.Sleep(wait)
	}
	if err != nil {
		return fmt.Errorf("fail to connect db after retrying %v times: %w", storer.options.retryCount, err)
	}
	// success, perform database migration up
	if storer.options.migrateUp {
		err = storer.MigrateUp()
	}
	return err
}

//
// database migration
// * use migration files in storer.options.migrationPath to perform migration up or down
//   - 0001_init_users.up.json create users indexes
//   - 0001_init_users.down.json drop users indexes
//   - 0002_init_privilegeProfiles.up.json create privilegeProfiles indexes and default privilege profiles
//   - 0002_init_privilegeProfiles.down.json drop privilegeProfiles indexes and remove default privilege profiles
// * may install [migrate CLI](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate) to perform migration up or down
//   - install the migrate cli program
//   - invoke "migrate -path user/deploy/migrations/ -database "mongodb://$IMS_SECURITY_SERVICE_USER_DATABASE_USER:$IMS_SECURITY_SERVICE_USER_DATABASE_PASSWORD@localhost:27017/imsSecurityUserDB" up" to migrate up
//   - invoke "migrate -path user/deploy/migrations/ -database "mongodb://$IMS_SECURITY_SERVICE_USER_DATABASE_USER:$IMS_SECURITY_SERVICE_USER_DATABASE_PASSWORD@localhost:27017/imsSecurityUserDB" down" to migrate down
//   - invoke "migrate -path user/deploy/migrations/ -database "mongodb://$IMS_SECURITY_SERVICE_USER_DATABASE_USER:$IMS_SECURITY_SERVICE_USER_DATABASE_PASSWORD@localhost:27017/imsSecurityUserDB" drop" to drop everything inside the database
// * may connect to mongodb server to check the database
//   - invoke "sudo apt install mongodb-clients" to install mongo client
//   - invoke "mongo localhost:27017/imsSecurityUserDB -u $IMS_SECURITY_SERVICE_USER_DATABASE_USER -p $IMS_SECURITY_SERVICE_USER_DATABASE_PASSWORD" to connect to server
//   - invoke "mongo localhost:27017/imsSecurityUserDB -u $IMS_SECURITY_SERVICE_USER_DATABASE_USER -p $IMS_SECURITY_SERVICE_USER_DATABASE_PASSWORD --eval "db.getCollectionNames()"" to show imsSecurityUserDB's all collections
//   - invoke "mongo localhost:27017/imsSecurityUserDB -u $IMS_SECURITY_SERVICE_USER_DATABASE_USER -p $IMS_SECURITY_SERVICE_USER_DATABASE_PASSWORD --eval "db.privilegeProfiles.find()"" to list all documents in collection privilegeProfiles
//

// MigrateUp perform database migration up
func (storer *StorerMongodb) MigrateUp() (err error) {
	config := &mongodb.Config{
		DatabaseName:         storer.options.dbname,
		MigrationsCollection: storer.options.migrationCollectionName,
	}
	driver, err := mongodb.WithInstance(storer.client, config)
	if err != nil {
		return fmt.Errorf("fail to mongodb.WithInstance: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+storer.options.migrationPath, storer.options.dbname, driver)
	if err != nil {
		return fmt.Errorf("fail to migrate.NewWithDatabaseInstance: %w", err)
	}
	err = m.Up()
	if (err != nil) && (err != migrate.ErrNoChange) {
		return fmt.Errorf("fail to migrate.Up: %w", err)
	}
	return nil
}

// MigrateDown perform database migration down
func (storer *StorerMongodb) MigrateDown() (err error) {
	config := &mongodb.Config{
		DatabaseName:         storer.options.dbname,
		MigrationsCollection: storer.options.migrationCollectionName,
	}
	driver, err := mongodb.WithInstance(storer.client, config)
	if err != nil {
		return fmt.Errorf("fail to mongodb.WithInstance: %w", err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://"+storer.options.migrationPath, storer.options.dbname, driver)
	if err != nil {
		return fmt.Errorf("fail to migrate.NewWithDatabaseInstance: %w", err)
	}
	err = m.Down()
	if (err != nil) && (err != migrate.ErrNoChange) {
		return fmt.Errorf("fail to migrate.Down: %w", err)
	}
	return nil
}
