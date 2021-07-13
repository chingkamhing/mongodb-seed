package model

type Receipt struct {
	Id             string `json:"id,omitempty" bson:"_id,omitempty"`
	TxnTypeDisplay string `json:"txnTypeDisplay,omitempty"`
	Tid            string `json:"tid,omitempty"`
	Currency       string `json:"currency,omitempty"`
	Icc            ICC
	ItemId         string `json:"itemId,omitempty"`
	Merchant       Merchant
	Prefix         string   `json:"prefix,omitempty"`
	Amount         int      `json:"amount,omitempty"`
	Trace          string   `json:"trace,omitempty"`
	Batch          string   `json:"batch,omitempty"`
	UpiRrn         string   `json:"UpiRrn,omitempty"`
	AcctId         string   `json:"acctId,omitempty"`
	Logo           string   `json:"logo,omitempty"`
	Mid            string   `json:"mid,omitempty"`
	ProcessorName  string   `json:"processorName,omitempty"`
	ReceiptId      string   `json:"receiptId,omitempty"`
	Outlet         string   `json:"outlet,omitempty"`
	Address        string   `json:"address,omitempty"`
	ExpiryDate     int64    `json:"expiryDate,omitempty"`
	TermId         string   `json:"termId,omitempty"`
	Scheme         string   `json:"scheme,omitempty"`
	Last4Digit     string   `json:"last4Digit,omitempty"`
	AppCode        string   `json:"appCode,omitempty"`
	UpiTrace       string   `json:"UpiTrace,omitempty"`
	CuponInfo      string   `json:"CuponInfo,omitempty"`
	MertId         string   `json:"mertId,omitempty"`
	TxnTime        int64    `json:"txnTime,omitempty"`
	Method         string   `json:"method,omitempty"`
	Ref            string   `json:"ref,omitempty"`
	ReportId       string   `json:"reportId,omitempty"`
	TxnType        []string `json:"txnType,omitempty"`
	Signature      string   `json:"signature,omitempty"`
}

type ICC struct {
	App string `json:"app,omitempty"`
	Tc  string `json:"tc,omitempty"`
	Tvr string `json:"tvr,omitempty"`
}

type Merchant struct {
	Outlet  string `json:"outlet,omitempty"`
	Address string `json:"address,omitempty"`
}
