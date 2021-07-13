package mongodb

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//
// This driver file is for mongo driver related codes
// * should idealy seperate driver and app-related logics (e.g. user, privilege profile, etc.)
// * so that this driver can easily be seperated to another package in the future
//

// internal driver cursor functions

type driverCursor struct {
	timeout time.Duration
	cursor  *mongo.Cursor
}

func (cursor *driverCursor) next() bool {
	ctx, _ := context.WithTimeout(context.Background(), cursor.timeout)
	return cursor.cursor.Next(ctx)
}

func (cursor *driverCursor) decode(value interface{}) (err error) {
	return cursor.cursor.Decode(value)
}

// mongodb filters

// mongodb search all pattern
func filterAll() bson.M {
	return bson.M{}
}

// get mongodb filter that specify _id
func filterID(id primitive.ObjectID) bson.M {
	return bson.M{"_id": id}
}

// get mongodb filter that specify username
func filterReceiptId(username string) bson.M {
	return bson.M{"receiptId": username}
}

// mongodb updates

// mongodb set specified whole document
func updateSetDocument(document interface{}) bson.D {
	return bson.D{primitive.E{Key: "$set", Value: document}}
}

// mongodb driver functions

// Ping verify if the remote mongodb server is active
func (storer *StorerMongodb) GetURL() string {
	host := storer.options.host
	port := storer.options.port
	u, _ := url.Parse("")
	u.Scheme = "mongodb"
	u.Host = net.JoinHostPort(host, port)
	url := u.String()
	return url
}

// open connect to the remote mongodb server
// * connection string URI format:
// 	 mongodb://[username:password@]host1[:port1][,...hostN[:portN]][/[defaultauthdb][?options]]
// 	 e.g. mongodb://username:password@localhost:27017/myDB
func (storer *StorerMongodb) open() error {
	url := storer.GetURL()
	clientOptions := options.Client().ApplyURI(url)
	// if database name and user name is specified, set authentication options
	if (storer.options.dbname != "") && (storer.options.user != "") {
		clientOptions = clientOptions.SetAuth(options.Credential{
			AuthSource: storer.options.dbname,
			Username:   storer.options.user,
			Password:   storer.options.password,
		})
	}
	// connect to the monogodb server
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return err
	}
	storer.client = client
	// if the database name is specified, open the database
	if storer.options.dbname != "" {
		storer.useDatabase(storer.options.dbname)
	}
	return nil
}

// useDatabase select specified database name
func (storer *StorerMongodb) useDatabase(database string) {
	storer.database = storer.client.Database(database)
}

// Ping verify if the remote mongodb server is active
func (storer *StorerMongodb) Ping() error {
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	return storer.client.Ping(ctx, nil)
}

// Close close mongodb connection
func (storer *StorerMongodb) Close() error {
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	return storer.client.Disconnect(ctx)
}

// create related driver functions

// createOne create one document with specified collection name and document
func (storer *StorerMongodb) createOne(collectionName string, document interface{}) (id string, err error) {
	collection := storer.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	result, err := collection.InsertOne(ctx, document)
	if err != nil {
		return "", fmt.Errorf("fail to createOne(): %w", err)
	}
	return result.InsertedID.(primitive.ObjectID).Hex(), nil
}

// find related driver functions

// findOne find one document with specified collection name and filter
func (storer *StorerMongodb) findOne(collectionName string, filter bson.M) (result *mongo.SingleResult) {
	collection := storer.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	result = collection.FindOne(ctx, filter)
	return result
}

// findMany find documents with specified collection name and filter
func (storer *StorerMongodb) findMany(collectionName string, filter bson.M) (cursor *driverCursor, err error) {
	collection := storer.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	c, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("fail to findMany(): %w", err)
	}
	return &driverCursor{cursor: c, timeout: storer.options.timeout}, nil
}

// update related driver functions

// updateOne update one document with specified collection name, filter and document
func (storer *StorerMongodb) updateOne(collectionName string, filter bson.M, update interface{}) (count int64, err error) {
	collection := storer.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return 0, fmt.Errorf("fail to updateOne.UpdateOne(): %w", err)
	}
	return result.ModifiedCount, nil
}

// upsertOne update or insert one document with specified collection name, filter and document
func (storer *StorerMongodb) upsertOne(collectionName string, filter bson.M, update interface{}) (count int64, err error) {
	collection := storer.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return 0, fmt.Errorf("fail to upsertOne.UpdateOne(): %w", err)
	}
	return result.ModifiedCount, nil
}

// delete related driver functions

// deleteOne delete one document with specified collection name and filter
func (storer *StorerMongodb) deleteOne(collectionName string, filter bson.M) (count int64, err error) {
	collection := storer.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("fail to deleteOne.DeleteOne(): %w", err)
	}
	return result.DeletedCount, nil
}

// deleteMany delete many documents with specified collection name and filter
func (storer *StorerMongodb) deleteMany(collectionName string, filter bson.M) (count int64, err error) {
	collection := storer.database.Collection(collectionName)
	ctx, _ := context.WithTimeout(context.Background(), storer.options.timeout)
	result, err := collection.DeleteMany(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("fail to deleteMany.DeleteMany(): %w", err)
	}
	return result.DeletedCount, nil
}
