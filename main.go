package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/chingkamhing/mongodb-seed/config"
	"github.com/chingkamhing/mongodb-seed/model"
	"github.com/chingkamhing/mongodb-seed/repository"
)

var receiptAccount = map[string]int{
	"MI MING MART - TAIKOO":    7354,
	"MI MING MART - CWB":       20285,
	"MI MING MART - TST":       14305,
	"MI MING MART - YOHO MALL": 8386,
	"MI MING MART - TUEN MUN":  7580,
	"MI MING MART - TSUEN WAN": 8848,
	"MI MING MART - MK":        12417,
	"MI MING MART - TKO":       7206,
	"MI MING MART - SHATIN":    11070,
	"MI MING MART - APM":       9689,
	"MI MING MART - Office":    46,
}

const myLogo = "iVBORw0KGgoAAAANSUhEUgAAASIAAABHCAYAAACj8ErrAAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAA3NpVFh0WE1MOmNvbS5hZG9iZS54bXAAAAAAADw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNS1jMDE0IDc5LjE1MTQ4MSwgMjAxMy8wMy8xMy0xMjowOToxNSAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD0ieG1wLmRpZDo1YzQzMDg2MC0yZDBkLTRmNTYtOGJkOC0zY2VhYmQwNWZmMWUiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6MDc2NDJEN0FEMzExMTFFN0JDQjdGRjdGMzI0MjRGN0QiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6MDc2NDJENzlEMzExMTFFN0JDQjdGRjdGMzI0MjRGN0QiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIChNYWNpbnRvc2gpIj4gPHhtcE1NOkRlcml2ZWRGcm9tIHN0UmVmOmluc3RhbmNlSUQ9InhtcC5paWQ6ZWNhMmY3MTYtYTliMS00NzI3LTlmODUtMjk1NzE5OGJjYTdmIiBzdFJlZjpkb2N1bWVudElEPSJ4bXAuZGlkOjVjNDMwODYwLTJkMGQtNGY1Ni04YmQ4LTNjZWFiZDA1ZmYxZSIvPiA8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"

var numAccount int

func main() {
	// parse flags
	flag.StringVar(&config.Config.Database.Host, "host", "localhost", "mongodb host address")
	flag.IntVar(&config.Config.Database.Port, "port", 27017, "mongodb port number")
	flag.StringVar(&config.Config.Database.Dbname, "dbname", "db_test", "mongodb database name")
	flag.StringVar(&config.Config.Database.Username, "username", "db_test_user", "mongodb username")
	flag.StringVar(&config.Config.Database.Password, "password", "", "mongodb password")
	flag.IntVar(&numAccount, "num", 10, "Number of accounts will be created")
	flag.Parse()

	// connect to database
	repo := repository.New()
	err := repo.Open()
	if err != nil {
		log.Fatalf("Fail to open repository: %v\n", err)
	}
	defer repo.Close()

	// create receipt records for different outlets
	for i := 1; i <= numAccount; i++ {
		accountId := fmt.Sprintf("1%011d", i)
		fmt.Printf("acctId[%s]...\n", accountId)
		for outletName, count := range receiptAccount {
			receipt := newOutletReceipt(repo, accountId, outletName, 1609459200000)
			receipt.createReceipts(count)
		}
	}
}

// receiptCreator

type receiptCreator struct {
	repo         repository.Interface
	acctId       string
	outlet       string
	startTxnTime int64
	amount       int
}

func newOutletReceipt(repo repository.Interface, acctId string, outlet string, txnTime int64) *receiptCreator {
	return &receiptCreator{
		repo:         repo,
		acctId:       acctId,
		outlet:       outlet,
		startTxnTime: txnTime,
		amount:       1000000,
	}
}

func (r *receiptCreator) createReceipts(count int) (err error) {
	fmt.Printf("createReceipts[%s %s]...\n", r.acctId, r.outlet)
	for i := 0; i < count; i++ {
		_, err = r.createReceipt(i)
		if err != nil {
			return fmt.Errorf("createReceipts %s %s %d error: %w", r.acctId, r.outlet, i, err)
		}
	}
	return nil
}

func (r *receiptCreator) createReceipt(index int) (id string, err error) {
	indexString := strconv.Itoa(index)
	index12String := fmt.Sprintf("%012d", index)
	receipt := &model.Receipt{
		TxnTypeDisplay: " 銷售 Sales #" + indexString,
		Tid:            "",
		Currency:       "HKD",
		Icc:            model.ICC{App: "Credit Application", Tc: "56fe3daea2afb234", Tvr: "0000000800"},
		ItemId:         "landi-aposa8-" + index12String,
		Merchant:       model.Merchant{Outlet: "Merchant's Outlet Name " + indexString, Address: "Outlet address, xxx distrinct, xxx road, xxx building, xxx shop, number " + indexString},
		Prefix:         "China Union",
		Amount:         r.amount + index,
		Trace:          "000351",
		Batch:          "000029",
		UpiRrn:         "",
		AcctId:         r.acctId,
		Mid:            "",
		ProcessorName:  "cup",
		ReceiptId:      "f369c59e-397b-45a0-bbd8-" + index12String,
		Outlet:         r.outlet,
		Address:        "",
		ExpiryDate:     2556115199999,
		TermId:         "",
		Scheme:         "UnionPay",
		Last4Digit:     "0141",
		AppCode:        "053344",
		UpiTrace:       "000000",
		CuponInfo:      "",
		MertId:         "",
		TxnTime:        r.startTxnTime + int64(index*1000),
		Method:         "Q",
		Ref:            index12String,
		ReportId:       "45ecb758-4f13-43ae-b856-" + index12String,
		TxnType:        []string{"sale"},
		Signature:      "",
		Logo:           myLogo,
	}
	id, err = r.repo.CreateReceipt(receipt)
	if err != nil {
		return "", fmt.Errorf("createReceipt repo.CreateReceipt error: %w", err)
	}
	return id, nil
}
