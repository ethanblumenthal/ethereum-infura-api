package app

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func benchmarkGetBlockNumber(numReqs int, b *testing.B) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "I am a super server")
	}))

	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "ping")
	}

	for i := 0; i <= numReqs; i++ {
		go func() {
			req := httptest.NewRequest("GET", "http://localhost:8080/blocks/number", nil)
			w := httptest.NewRecorder()
			handler(w, req)

			resp := w.Result()
			body, _ := ioutil.ReadAll(resp.Body)

			fmt.Println(resp.StatusCode)
			fmt.Println(string(body))

			
		} ()
	}

	defer server.Close()
}

func Test_GetBlockNumber(t *testing.T) {
	// controller.GetBlockNumber()
}

func Test_GetBlockByNumber(t *testing.T) {
	// controller.GetBlockByNumber()
}

func Test_GetBlockByHash(t *testing.T) {
	// controller.GetBlockByHash()
}

func Test_GetTransactionByHash(t *testing.T) {
	// controller.GetTransactionByHash()
}

func Benchmark_GetBlockNumber1(b *testing.B)  { benchmarkGetBlockNumber(1, b) }
func Benchmark_GetBlockNumber2(b *testing.B)  { benchmarkGetBlockNumber(2, b) }
func Benchmark_GetBlockNumber3(b *testing.B)  { benchmarkGetBlockNumber(300000, b) }
func Benchmark_GetBlockNumber10(b *testing.B) { benchmarkGetBlockNumber(10, b) }
func Benchmark_GetBlockNumber20(b *testing.B) { benchmarkGetBlockNumber(2, b) }
func Benchmark_GetBlockNumber40(b *testing.B) { benchmarkGetBlockNumber(40, b) }