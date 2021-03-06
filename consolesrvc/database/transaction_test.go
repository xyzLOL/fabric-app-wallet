package database

import (
	"testing"
	"database/sql"
	util "baas/app-wallet/consolesrvc/common"
)

func TestAddTransaction(t *testing.T) {
	var db *sql.DB = GetDB()
	defer db.Close()

	var testNum = 2

	var trans = make([]Transaction, testNum)
	for i, _ := range trans {
		trans[i].TxUUID = util.GenerateUUID()
		trans[i].PayerUUID = util.GenerateUUID()
		trans[i].PayeeUUID = util.GenerateUUID()
		trans[i].PayerAccountID = util.GenerateUUID()
		trans[i].PayeeAccountID = util.GenerateUUID()
		trans[i].Amount = 30
		trans[i].Status = "pending"
		dbLogger.Debugf("transaction: %#v", trans[i])
	}
	trans[0].PayerUUID = "5cdb617c-2712-480a-a02b-facd8c86e579"
	trans[1].TxUUID = trans[0].TxUUID

	var tests = []struct{
		newline bool
		sep string
		arg *Transaction
		want int64
	}{
		{false, " ", &trans[0], 1},
		{true, " ", &trans[1], 0},
	}

	for i, testitem := range tests {
		rowsAff, _ := AddTransaction(db, testitem.arg)
		if rowsAff != testitem.want {
			t.Errorf("Test #%d: Add transaction %#v, affected rows = %d, but want %d", i, testitem.arg, rowsAff, testitem.want)
		}
	}
}

func TestGetTransactionByTransUUID(t *testing.T) {
	var db *sql.DB = GetDB()
	defer db.Close()

	var err error
	var us *Transaction
	var txuuid string
	//useruuid = "5cdb617c-2712-480a-a02b-facd8c86e579"
	txuuid = "4b264cf2-deaf-496d-8eb9-49ee39044e6e"
	us, err = GetTransaction(db, txuuid)
	if us == nil || err != nil{
		t.Error("Failed retrieving transaction")
	}
	dbLogger.Debugf("Get transaction: %#v", *us)
}

func TestGetTransactionsByPayeruuid(t *testing.T) {
	var db *sql.DB = GetDB()
	defer db.Close()

	var err error
	var txs []*Transaction

	var useruuid string = "5cdb617c-2712-480a-a02b-facd8c86e579"
	txs, err = GetTransactionsByPayeruuid(db, useruuid)
	if err != nil {
		t.Errorf("Failed retrieving user accounts by useruuid %s: %v", useruuid, err)
	}
	for i, txitem := range txs {
		dbLogger.Debugf("Accounts #%d: %v", i, *txitem)
	}
}

func TestUpdateTransaction(t *testing.T) {
	var db *sql.DB = GetDB()
	defer db.Close()

	var err error
	var us *Transaction

	var txuuid string
	//useruuid = "5cdb617c-2712-480a-a02b-facd8c86e579"
	txuuid = "4b264cf2-deaf-496d-8eb9-49ee39044e6e"
	us, err = GetTransaction(db, txuuid)
	if us == nil || err != nil{
		t.Error("Failed retrieving transaction")
	}
	dbLogger.Debugf("Get transaction: %#v", *us)

	us.BC_txuuid = util.GenerateUUID()
	us.BC_blocknum = 12
	us.Status = "fin"
	var affectedrows int64 = 0
	affectedrows, err = UpdateTransaction(db, us)
	if affectedrows != 1 {
		t.Errorf("failed updating transaction %v\n err: %v", *us, err)
	}
	dbLogger.Debugf("Updated transaction: %#v", *us)

}


