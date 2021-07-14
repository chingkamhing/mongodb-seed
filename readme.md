# Mongodb Seed Program for e-receipt

* this program is created issue [Searching Issues: Use filter return 0](https://tess.hk-tess.com:7080/cloud/receipt/issues/1)
* it seed the mongodb with simulated data to evaluate how to improve database query performance

## Measurement Result

* test pc info
    + Acer Aspire E5-571G notebook
    + Intel i5 - 4210U, 1.70 GHz, 3 MB
    + 12GB RAM
    + Kingston A400 SATA 240G SSD
* before create any indexes (except the default "_id")
    ```
    $ time mongo mongodb://localhost:27017/database_name --username database_username --password database_password --quiet --eval 'db.Receipts.find({acctId:"954815416"}).count()'
    107186
    real    3m50.448s
    user    0m0.089s
    sys     0m0.028s
    ```
* create invidual index
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
* create compound index
    ```
    db.Receipts.createIndex({"acctId": 1,"outlet": 1,"txnTime": 1})
    $ time mongo mongodb://localhost:27017/database_name --username database_username --password database_password --quiet --eval 'db.Receipts.find({acctId:"954815416",outlet:"MI MING MART 954815416 - MK"}).sort({txnTime:-1})' > /dev/null
    real    0m0.277s
    user    0m0.103s
    sys     0m0.033s
    ```
* drop individual indexes which are covered by the compound index while all the previous queries still takes less than 1s
    ```
    db.Receipts.dropIndexes("acctId_1")
    db.Receipts.dropIndexes("outlet_1")
    db.Receipts.dropIndexes("txnTime_1")
    ```

## TLDR

* compound index increase compound query dramatically from >25s to <1s for 857488 records
