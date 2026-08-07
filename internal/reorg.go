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

package internal

import (
	"errors"
	"fmt"
)

// BlockReorgManager handles the blockchain reorganization (reorg) logic.
type BlockReorgManager struct {
	// Reference to your blockchain database or storage state can be placed here.
}

// NewBlockReorgManager initializes and returns a new instance of BlockReorgManager.
func NewBlockReorgManager() *BlockReorgManager {
	return &BlockReorgManager{}
}

// HandleReorg checks whether an incoming block extends the main chain or triggers a fork reorganization.
func (m *BlockReorgManager) HandleReorg(newBlock *Block, currentTip *Block) error {
	// 1. If the new block directly extends the current active chain tip
	if newBlock.PrevHash == currentTip.Hash {
		return m.AppendToMainChain(newBlock)
	}

	// 2. If the new block has an equal or higher height, indicating a potential fork
	if newBlock.Height >= currentTip.Height {
		fmt.Println("[Reorg] Fork detected! Evaluating alternative chain...")

		ancestor, err := m.FindCommonAncestor(currentTip, newBlock)
		if err != nil {
			return fmt.Errorf("[Reorg] failed to find common ancestor: %v", err)
		}

		// Rollback the old chain from the current tip back to the common ancestor
		if err := m.RollbackChain(currentTip, ancestor); err != nil {
			return fmt.Errorf("[Reorg] failed to rollback chain: %v", err)
		}

		// Apply the new alternative chain from the common ancestor up to the new block
		if err := m.ApplyAlternativeChain(newBlock, ancestor); err != nil {
			return fmt.Errorf("[Reorg] failed to apply alternative chain: %v", err)
		}

		fmt.Println("[Reorg] Block reorganization completed successfully.")
		return nil
	}

	// Reject the block if the alternative chain is too short or invalid
	return errors.New("[Reorg] alternative chain is too short or invalid")
}

// AppendToMainChain adds a valid block normally to the tip of the primary active chain.
func (m *BlockReorgManager) AppendToMainChain(block *Block) error {
	// Implementation logic for saving the block to the active main chain database
	fmt.Printf("[Chain] Appending block %s at height %d to main chain\n", block.Hash[:8], block.Height)
	return nil
}

// FindCommonAncestor traces backward to locate the latest shared ancestor block between two chains.
func (m *BlockReorgManager) FindCommonAncestor(tipA *Block, tipB *Block) (*Block, error) {
	// Implementation logic to traverse parent hashes until a common block is found
	fmt.Println("[Reorg] Locating common ancestor...")
	return &Block{}, nil // Adjust this to match your project's Block struct
}

// RollbackChain disconnects blocks from the active main chain and returns their transactions to the mempool.
func (m *BlockReorgManager) RollbackChain(fromTip *Block, ancestor *Block) error {
	fmt.Printf("[Reorg] Rolling back from block %s down to ancestor %s\n", fromTip.Hash[:8], ancestor.Hash[:8])
	// Implementation logic for disconnecting blocks and restoring transactions
	return nil
}

// ApplyAlternativeChain connects and validates blocks from the new heavier branch to establish the new active chain.
func (m *BlockReorgManager) ApplyAlternativeChain(toTip *Block, ancestor *Block) error {
	fmt.Printf("[Reorg] Applying new chain branch up to block %s\n", toTip.Hash[:8])
	// Implementation logic for validating and attaching new branch blocks
	return nil
}
