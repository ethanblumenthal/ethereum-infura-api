package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/INFURA/go-libs/jsonrpc_client"
	"github.com/gorilla/mux"
)

type Controller struct {
    EthereumClient jsonrpc_client.EthereumClient
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

// GetBlockNumber GET - Gets the current block number
func (c *Controller) GetBlockNumber(w http.ResponseWriter, r *http.Request) {
    blockNum, err := c.EthereumClient.Eth_blockNumber()
	if err != nil {
        log.Fatalln("Error GetBlockNumber", err)
    }
	
    c.writeRes(w, blockNum)
    return
}

// GetBlockByNumber GET - Gets a single block by number
func (c *Controller) GetBlockByNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    blockNum := vars["num"]
    blockInt, err := strconv.Atoi(blockNum)

    block, err := c.EthereumClient.Eth_getBlockByNumber(blockInt, true)
	if err != nil {
        log.Fatalln("Error GetBlockByNumber", err)
    }
	
    c.writeRes(w, block)
    return
}

// GetBlockByHash GET - Gets a single block by hash
func (c *Controller) GetBlockByHash(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    blockHash := vars["hash"]

    block, err := c.EthereumClient.Eth_getBlockByHash(blockHash, true)
	if err != nil {
        log.Fatalln("Error GetBlockByHash", err)
    }
	
    c.writeRes(w, block)
    return
}

// GetTransactionByHash GET - Gets a single transaction by hash
func (c *Controller) GetTransactionByHash(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    txHash := vars["hash"]

    tx, err := c.EthereumClient.Eth_getTransactionByHash(txHash)
	if err != nil {
        log.Fatalln("Error GetTransactionByHash", err)
    }
	
    c.writeRes(w, tx)
    return
}
