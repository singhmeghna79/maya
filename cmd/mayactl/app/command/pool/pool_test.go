// Copyright © 2018 The OpenEBS Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pool

import (
	"testing"

	"github.com/spf13/cobra"
)

func TestNewCmdPool(t *testing.T) {
	tests := map[string]*struct {
		expectedCmd *cobra.Command
	}{
		"NewCmdVolumeStats": {
			expectedCmd: &cobra.Command{
				Use:   "pool",
				Short: "Provides operations related to a storage pool",
				Long:  poolCommandHelpText,
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := NewCmdPool()
			if (got.Use != tt.expectedCmd.Use) || (got.Short != tt.expectedCmd.Short) || (got.Long != tt.expectedCmd.Long) || (got.Example != tt.expectedCmd.Example) {
				t.Fatalf("TestName: %v | NewCmdPool() => Got: %v | Want: %v \n", name, got, tt.expectedCmd)
			}
		})
	}
}

// returns true when both errors are true or else returns false
func checkErr(err1, err2 error) bool {
	if (err1 != nil && err2 == nil) || (err1 == nil && err2 != nil) || (err1 != nil && err2 != nil && err1.Error() != err2.Error()) {
		return false
	}
	return true
}
