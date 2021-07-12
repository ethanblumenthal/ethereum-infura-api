package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/INFURA/go-libs/jsonrpc_client"
	"github.com/coocood/freecache"
	"github.com/gorilla/mux"
)

type Handler struct {
    Cache freecache.Cache
    EthereumClient jsonrpc_client.EthereumClient
}

type ErrRes struct {
	Error string `json:"error"`
}

func (h *Handler) writeErrRes(w http.ResponseWriter, err error) {
	jsonErrRes, _ := json.Marshal(ErrRes{err.Error()})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write(jsonErrRes)
}

func (h *Handler) writeRes(w http.ResponseWriter, contentJson []byte) {
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(contentJson)
}

func (h *Handler) setCache(key []byte, val []byte) {
    expire := 600 // expire in 600 seconds (10 min)
    err := h.Cache.Set(key, val, expire)
    if err != nil {
        fmt.Println(err)
    }
}

func (h *Handler) getCache(key []byte) (val []byte, err error) {
    val, err = h.Cache.Get(key)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(string(val))
    return val, err
}

// GetBlockNumber GET - Gets the current block number
func (h *Handler) GetBlockNumber(w http.ResponseWriter, r *http.Request) {
    blockNum, err := h.EthereumClient.Eth_blockNumber()
	if err != nil {
        log.Fatalln("Error GetBlockNumber", err)
    }

    contentJson, err := json.Marshal(blockNum)
	if err != nil {
		h.writeErrRes(w, err)
		return
	}
	
    h.writeRes(w, contentJson)
    return
}

// GetBlockByNumber GET - Gets a single block by number
func (h *Handler) GetBlockByNumber(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    blockNum := vars["num"]
    blockInt, err := strconv.Atoi(blockNum)

    cachedBlock, err := h.getCache([]byte(blockNum))
    if cachedBlock != nil {
        log.Println("Found cachedBlock", cachedBlock)
        h.writeRes(w, cachedBlock)
        return
    }

    block, err := h.EthereumClient.Eth_getBlockByNumber(blockInt, true)
	if err != nil {
        log.Fatalln("Error GetBlockByNumber", err)
    }

    contentJson, err := json.Marshal(block)
	if err != nil {
		h.writeErrRes(w, err)
		return
	}
	
    h.setCache([]byte(blockNum), contentJson)
    h.writeRes(w, contentJson)
    return
} 

// GetBlockByHash GET - Gets a single block by hash
func (h *Handler) GetBlockByHash(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    blockHash := vars["hash"]

    cachedBlock, err := h.getCache([]byte(blockHash))
    if cachedBlock != nil {
        log.Println("Found cachedBlock", cachedBlock)
        h.writeRes(w, cachedBlock)
        return
    }

    block, err := h.EthereumClient.Eth_getBlockByHash(blockHash, true)
	if err != nil {
        log.Fatalln("Error GetBlockByHash", err)
    }

    contentJson, err := json.Marshal(block)
	if err != nil {
		h.writeErrRes(w, err)
		return
	}
	
    h.setCache([]byte(blockHash), contentJson)
    h.writeRes(w, contentJson)
    return
}

// GetTransactionByHash GET - Gets a single transaction by hash
func (h *Handler) GetTransactionByHash(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    txHash := vars["hash"]

    cachedTx, err := h.getCache([]byte(txHash))
    if cachedTx != nil {
        log.Println("Found cachedTx", cachedTx)
        h.writeRes(w, cachedTx)
        return
    }
    
    tx, err := h.EthereumClient.Eth_getTransactionByHash(txHash)     
	if err != nil {
        log.Fatalln("Error GetTransactionByHash", err)
    }
	
    contentJson, err := json.Marshal(tx)
	if err != nil {
		h.writeErrRes(w, err)
		return
	}

    h.setCache([]byte(txHash), contentJson)
    h.writeRes(w, contentJson)
    return
}
