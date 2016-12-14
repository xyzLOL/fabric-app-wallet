package main

import (

	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
	"github.com/op/go-logging"

	"baas/app-wallet/consolesrvc/common"
	"baas/app-wallet/consolesrvc/database"
	"github.com/robfig/cron"
	"baas/app-wallet/consolesrvc/wallet/cronjob"
	"baas/app-wallet/consolesrvc/auth"
	"baas/app-wallet/consolesrvc/wallet/account"
	"baas/app-wallet/consolesrvc/wallet/transaction"
	"baas/app-wallet/consolesrvc/blockchain"
)

var consLogger *logging.Logger = common.NewLogger("console") //logging.MustGetLogger("Console")


func main() {
	var db *sql.DB = database.GetDB()
	defer db.Close()
	var c = cron.New()
	c.AddJob("*/5 * * * * ?", &cronjob.JobCreateAccount{})
	c.AddJob("*/5 * * * * ?", &cronjob.JobAccountTransfer{})
	c.Start()
	defer c.Stop()

	router := httprouter.New()
	RegisterRouter(router)

	consLogger.Info("start to listen and serve for localhost:8765")
	consLogger.Fatal(http.ListenAndServe(":8765", router))
}


func RegisterRouter(router *httprouter.Router){
	router.Handle("POST", "/auth/login", authsrvc.LoginPost)
	router.Handle("POST", "/auth/signup", authsrvc.SignupPost)
	router.Handle("POST", "/auth/refresh", authsrvc.RefreshPost)
	router.Handle("POST", "/auth/logout", authsrvc.LogoutPost)

	router.Handle("POST", "/wallet/account/create", account.AccountCreatePost)
	router.Handle("POST", "/wallet/account/list", account.AccountListPost)
	router.Handle("POST", "/wallet/account/transfer", account.TransferPost)
	router.Handle("POST", "/wallet/transaction/list", transaction.TransactionListPost)
	router.Handle("POST", "/blockchain/transaction", blockchain.TransactionDetailPost)
}