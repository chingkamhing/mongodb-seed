package model

type Receipt struct {
	Id             string `json:"id,omitempty" bson:"_id,omitempty"`
	TxnTypeDisplay string `json:"txnTypeDisplay,omitempty" bson:"txnTypeDisplay,omitempty"`
	Tid            string `json:"tid,omitempty" bson:"tid,omitempty"`
	Currency       string `json:"currency,omitempty" bson:"currency,omitempty"`
	Icc            ICC
	ItemId         string `json:"itemId,omitempty" bson:"itemId,omitempty"`
	Merchant       Merchant
	Prefix         string   `json:"prefix,omitempty" bson:"prefix,omitempty"`
	Amount         int      `json:"amount,omitempty" bson:"amount,omitempty"`
	Trace          string   `json:"trace,omitempty" bson:"trace,omitempty"`
	Batch          string   `json:"batch,omitempty" bson:"batch,omitempty"`
	UpiRrn         string   `json:"UpiRrn,omitempty" bson:"UpiRrn,omitempty"`
	AcctId         string   `json:"acctId,omitempty" bson:"acctId,omitempty"`
	Logo           string   `json:"logo,omitempty" bson:"logo,omitempty"`
	Mid            string   `json:"mid,omitempty" bson:"mid,omitempty"`
	ProcessorName  string   `json:"processorName,omitempty" bson:"processorName,omitempty"`
	ReceiptId      string   `json:"receiptId,omitempty" bson:"receiptId,omitempty"`
	Outlet         string   `json:"outlet,omitempty" bson:"outlet,omitempty"`
	Address        string   `json:"address,omitempty" bson:"address,omitempty"`
	ExpiryDate     int64    `json:"expiryDate,omitempty" bson:"expiryDate,omitempty"`
	TermId         string   `json:"termId,omitempty" bson:"termId,omitempty"`
	Scheme         string   `json:"scheme,omitempty" bson:"scheme,omitempty"`
	Last4Digit     string   `json:"last4Digit,omitempty" bson:"last4Digit,omitempty"`
	AppCode        string   `json:"appCode,omitempty" bson:"appCode,omitempty"`
	UpiTrace       string   `json:"UpiTrace,omitempty" bson:"UpiTrace,omitempty"`
	CuponInfo      string   `json:"CuponInfo,omitempty" bson:"CuponInfo,omitempty"`
	MertId         string   `json:"mertId,omitempty" bson:"mertId,omitempty"`
	TxnTime        int64    `json:"txnTime,omitempty" bson:"txnTime,omitempty"`
	Method         string   `json:"method,omitempty" bson:"method,omitempty"`
	Ref            string   `json:"ref,omitempty" bson:"ref,omitempty"`
	ReportId       string   `json:"reportId,omitempty" bson:"reportId,omitempty"`
	TxnType        []string `json:"txnType,omitempty" bson:"txnType,omitempty"`
	Signature      string   `json:"signature,omitempty" bson:"signature,omitempty"`
}

type ICC struct {
	App string `json:"app,omitempty" bson:"app,omitempty"`
	Tc  string `json:"tc,omitempty" bson:"tc,omitempty"`
	Tvr string `json:"tvr,omitempty" bson:"tvr,omitempty"`
}

type Merchant struct {
	Outlet  string `json:"outlet,omitempty" bson:"outlet,omitempty"`
	Address string `json:"address,omitempty" bson:"address,omitempty"`
}
