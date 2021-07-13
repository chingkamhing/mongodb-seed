package mongodb

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/chingkamhing/mongodb-seed/model"
)

// mongodb receipts collection name
const receiptsCollectionName = "Receipts"

//
// implement repository receipts Interface interface
//

// CreateReceipt create a new model.Receipt document in "receipt" collection
func (storer *StorerMongodb) CreateReceipt(receipt *model.Receipt) (id string, err error) {
	id, err = storer.createOne(receiptsCollectionName, receipt)
	if err != nil {
		return primitive.NilObjectID.Hex(), fmt.Errorf("fail to CreateReceipt(%v): %w", receipt.ReceiptId, err)
	}
	return id, nil
}

// ReadAllReceipts read all model.Receipt document in "receipt" collection
func (storer *StorerMongodb) ReadAllReceipts() (receipts []*model.Receipt, err error) {
	var cursor *driverCursor
	cursor, err = storer.findMany(receiptsCollectionName, filterAll())
	if err != nil {
		return nil, fmt.Errorf("fail to ReadAllReceipts.findMany(): %w", err)
	}
	for cursor.next() {
		var receipt model.Receipt
		err = cursor.decode(&receipt)
		if err != nil {
			return nil, fmt.Errorf("fail to ReadAllReceipts.decode(): %w", err)
		}
		receipts = append(receipts, &receipt)
	}
	return receipts, nil
}

// ReadReceiptByID read model.Receipt document in "receipts" collection with the specified receipt ID
func (storer *StorerMongodb) ReadReceiptByID(idString string) (receipt *model.Receipt, err error) {
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return nil, err
	}
	receipt = &model.Receipt{}
	if primitive.ObjectID(id).IsZero() {
		return receipt, nil
	}
	result := storer.findOne(receiptsCollectionName, filterID(primitive.ObjectID(id)))
	err = result.Decode(receipt)
	if err != nil {
		return nil, fmt.Errorf("fail to ReadReceiptByID.Decode(): %w", err)
	}
	return receipt, nil
}

// UpdateReceiptByReceiptId update model.Receipt from document in "receipt" collection with the specified receipt
func (storer *StorerMongodb) UpdateReceiptByReceiptId(receipt *model.Receipt) (count int64, err error) {
	filter := filterReceiptId(receipt.ReceiptId)
	update := updateSetDocument(receipt)
	return storer.updateOne(receiptsCollectionName, filter, update)
}

// UpdateReceiptByID update model.Receipt from document in "receipt" collection with the specified receipt
func (storer *StorerMongodb) UpdateReceiptByID(receipt *model.Receipt) (count int64, err error) {
	id, err := primitive.ObjectIDFromHex(receipt.Id)
	if err != nil {
		return 0, err
	}
	filter := filterID(id)
	update := updateSetDocument(receipt)
	count, err = storer.updateOne(receiptsCollectionName, filter, update)
	return count, err
}

// DeleteReceiptByID delete model.Receipt from document in "receipt" collection with the specified receipt ID
func (storer *StorerMongodb) DeleteReceiptByID(idString string) (count int64, err error) {
	id, err := primitive.ObjectIDFromHex(idString)
	if err != nil {
		return 0, err
	}
	filter := filterID(id)
	return storer.deleteOne(receiptsCollectionName, filter)
}

// DeleteAllReceipts delete model.Receipt document in "receipt" collection with specified search bson pattern
func (storer *StorerMongodb) DeleteAllReceipts() (count int64, err error) {
	return storer.deleteMany(receiptsCollectionName, filterAll())
}
