// Copyright (c) 2026 AldianOkto. All rights reserved.
// Copyright (c) 2026 Eterbit Core.
// Use of this source code is governed by the Apache License.
// that can be found in the root directory of this repository.
// Project: Eterbit / Blockchain Core
//
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at. <http://www.apache.org/licenses/LICENSE-2.0>
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"eterbit/node"
)

// RPCRequest represents the incoming JSON-RPC request structure.
type RPCRequest struct {
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
	ID      interface{}   `json:"id"`
}

// RPCResponse represents the standard JSON-RPC response structure.
type RPCResponse struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	ID     interface{} `json:"id"`
}

// StartRPCServer starts the JSON-RPC HTTP server on the specified port.
func StartRPCServer(rpcPort string, ledger *node.Ledger) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req RPCRequest
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&req); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response := RPCResponse{ID: req.ID}

		switch req.Method {
		case "getblockcount":
			ledger.Mu.Lock()
			count := len(ledger.Chain)
			ledger.Mu.Unlock()
			response.Result = count

		case "getconnectioncount":
			// Placeholder for peer connection count
			response.Result = 1

		case "getinfo":
			ledger.Mu.Lock()
			tip := ledger.GetLatestBlock()
			height := 0
			if tip != nil {
				height = tip.Index
			}
			ledger.Mu.Unlock()

			response.Result = map[string]interface{}{
				"blocks":      height,
				"connections": 1,
				"version":     10000,
			}

		default:
			response.Error = map[string]interface{}{
				"code":    -32601,
				"message": "Method not found",
			}
		}

		json.NewEncoder(w).Encode(response)
	})

	addr := fmt.Sprintf(":%s", rpcPort)
	fmt.Printf("[RPC] JSON-RPC server listening on port %s\n", rpcPort)
	go func() {
		if err := http.ListenAndServe(addr, mux); err != nil {
			fmt.Printf("[RPC] Server error: %v\n", err)
		}
	}()
}
