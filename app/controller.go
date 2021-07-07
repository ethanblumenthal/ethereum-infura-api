package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Controller ...
type Controller struct {
    Client Client
}

type ErrRes struct {
	Error string `json:"error"`
}

func (c *Controller) writeErrRes(w http.ResponseWriter, err error) {
	jsonErrRes, _ := json.Marshal(ErrRes{err.Error()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonErrRes)
}

func (c *Controller) writeRes(w http.ResponseWriter, content interface{}) {
	contentJson, err := json.Marshal(content)
	if err != nil {
		c.writeErrRes(w, err)
		return
	}

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(contentJson)
}

// GetAccount GET - Gets a single account by hash
func (c *Controller) GetAccount(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    log.Println(vars)

    accountHash := vars["hash"]
    log.Println(accountHash);

    account, err := c.Client.GetAccountByHash(accountHash)
	if err != nil {
        log.Fatalln("Error GetAccount", err)
    }
	
    c.writeRes(w, account)
    return
}

// GetBlock GET - Gets a single block by hash
func (c *Controller) GetBlock(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    log.Println(vars)

    blockHash := vars["hash"]
    log.Println(blockHash);

    block, err := c.Client.GetBlockByHash(blockHash)
	if err != nil {
        log.Fatalln("Error GetBlock", err)
    }
	
    c.writeRes(w, block)
    return
}

// GetTransaction GET - Gets a single tx by hash
func (c *Controller) GetTransaction(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    log.Println(vars)

    txHash := vars["hash"]
    log.Println(txHash);

    tx, err := c.Client.GetTransactionByHash(txHash)
	if err != nil {
        log.Fatalln("Error GetTransaction", err)
    }
	
    c.writeRes(w, tx)
    return
}
