# Mongodb Seed Program for e-receipt

* this program is created issue [Searching Issues: Use filter return 0](https://tess.hk-tess.com:7080/cloud/receipt/issues/1)
* it seed the mongodb with simulated data to evaluate how to improve database query performance

## Measurement Result

* test pc info
    + Acer Aspire E5-571G notebook
    + Intel i5 - 4210U, 1.70 GHz, 3 MB
    + 12GB RAM
    + Kingston A400 SATA 240G SSD
* seed the test database with 1 million documents
    + note: this seeding takes ~15 minutes
    ```
    $ make build && ./mongodb-seed --password database_password
    ```
* before create any indexes (except the default "_id"), query for "acctId" is very slow
    ```
    $ time mongo mongodb://localhost:27017/database_name --username database_username --password database_password --quiet --eval 'db.Receipts.find({acctId:"954815416"}).count()'
    107186
    real    3m50.448s
    user    0m0.089s
    sys     0m0.028s
    ```
* create following invidual indexes can help query for "acctId" but not for "acctId" AND "outlet"
    + note: each createIndex() will take >4 minutes
    ```
    db.Receipts.createIndex({"acctId": 1})
    db.Receipts.createIndex({"outlet": 1})
    db.Receipts.createIndex({"txnTime": 1})
    $ time mongo mongodb://localhost:27017/database_name --username database_username --password database_password --quiet --eval 'db.Receipts.find({acctId:"954815416"}).count()'
    107186
    real    0m0.244s
    user    0m0.092s
    sys     0m0.021s
    $ time mongo mongodb://localhost:27017/database_name --username database_username --password database_password --quiet --eval 'db.Receipts.find({acctId:"954815416",outlet:"MI MING MART 954815416 - MK"}).sort({txnTime:-1})' > /dev/null
    real    0m27.904s
    user    0m0.126s
    sys     0m0.026s
    ```
* create following compound index can help query for "acctId" AND "outlet" AND "txnType"
    ```
    db.Receipts.createIndex({"acctId": 1,"outlet": 1,"txnTime": 1})
    $ time mongo mongodb://localhost:27017/database_name --username database_username --password database_password --quiet --eval 'db.Receipts.find({acctId:"954815416",outlet:"MI MING MART 954815416 - MK"}).sort({txnTime:-1})' > /dev/null
    real    0m0.277s
    user    0m0.103s
    sys     0m0.033s
    ```
* drop redundant indexes
    + base on [Want MongoDB Performance? You Will Need to Add and Remove Indexes!](https://www.percona.com/blog/2021/03/22/want-mongodb-performance-you-will-need-to-add-and-remove-indexes/), left-most indexes are redundant and may be dropped
    ```
    db.Receipts.dropIndex("acctId_1")
    ```

## Knowledge Base
* [Index Build Operations on a Populated Collection](https://docs.mongodb.com/v4.0/core/index-creation/)
* [Want MongoDB Performance? You Will Need to Add and Remove Indexes!](https://www.percona.com/blog/2021/03/22/want-mongodb-performance-you-will-need-to-add-and-remove-indexes/)

## TLDR

* compound index increase compound query dramatically from >25s to <1s for 857488 records
