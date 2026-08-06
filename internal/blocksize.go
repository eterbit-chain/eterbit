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
    "os"
    "path/filepath"
)

// GetBlockChainStorageSize calculates the total physical storage size of the blockchain 
// database directory recursively in bytes by walking through all files.
func GetBlockChainStorageSize(dbPath string) (int64, error) {
    var totalSize int64 = 0

    err := filepath.Walk(dbPath, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.IsDir() {
            totalSize += info.Size()
        }
        return nil
    })

    if err != nil {
        return 0, err
    }

    return totalSize, nil
}
