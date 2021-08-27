// Copyright © 2021 Alibaba Group Holding Ltd.
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

package apply

import (
	"fmt"

	"github.com/alibaba/sealer/common"

	"github.com/alibaba/sealer/logger"

	"testing"
)

func TestNewCleanApplierFromArgs(t *testing.T) {
	tests := []struct {
		cFile   string
		cArgs   *common.RunArgs
		cOpts   *common.RunOpts
		name    string
		wantErr bool
	}{
		{
			"Clusterfile",
			&common.RunArgs{
				Masters: "10.110.101.1-10.110.101.5",
				Nodes:   "10.110.101.1-10.110.101.5",
			},
			&common.RunOpts{
				All:   false,
				Force: false,
			},
			"test1",
			false,
		},
		{
			"Clusterfile",
			&common.RunArgs{
				Masters: "10.110.101.1,10.110.101.2",
				Nodes:   "10.110.101.1,10.110.101.5",
			},
			&common.RunOpts{
				All:   false,
				Force: false,
			},
			"test2",
			false,
		},
		{
			"Clusterfile",
			&common.RunArgs{
				Masters: "2",
				Nodes:   "1",
			},
			&common.RunOpts{
				All:   false,
				Force: false,
			},
			"test3",
			false,
		},
		{
			"Clusterfile",
			&common.RunArgs{
				Masters: "-10.110.101.2",
				Nodes:   "10.110.101.2-",
			},
			&common.RunOpts{
				All:   false,
				Force: false,
			},
			"test4",
			true,
		},
		{
			"Clusterfile",
			&common.RunArgs{
				Masters: "-10.110.101.2",
				Nodes:   "10.110.101.2-",
			},
			&common.RunOpts{
				All:   false,
				Force: false,
			},
			"test4",
			true,
		},
		{
			"Clusterfile",
			&common.RunArgs{
				Masters: "b-a",
				Nodes:   "a-b",
			},
			&common.RunOpts{
				All:   false,
				Force: false,
			},
			"test4",
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := NewCleanApplierFromArgs(tt.cFile, tt.cArgs, tt.cOpts); (err != nil) != tt.wantErr {
				logger.Error("masters : %v , nodes : %v", &tt.cArgs.Masters, &tt.cArgs.Nodes)
			}
			logger.Info("masters : %v , nodes : %v", &tt.cArgs.Masters, &tt.cArgs.Nodes)
		})
	}
}

func Test_returnFilteredIPList(t *testing.T) {
	tests := []struct {
		name              string
		clusterIPList     []string
		toBeDeletedIPList []string
		wantErr           bool
	}{
		{
			"test",
			[]string{"10.10.10.1", "10.10.10.2", "10.10.10.3", "10.10.10.4"},
			[]string{"10.10.10.1", "10.10.10.2", "10.10.10.3", "10.10.10.4"},
			false,
		},
		{
			"test1",
			[]string{"10.10.10.1", "10.10.10.2", "10.10.10.3", "10.10.10.4"},
			[]string{},
			false,
		},
		{
			"test2",
			[]string{"10.10.10.1", "10.10.10.2", "10.10.10.3", "10.10.10.4"},
			[]string{"10.10.10.4"},
			false,
		},
		{
			"test3",
			[]string{},
			[]string{"10.10.10.1", "10.10.10.2", "10.10.10.3", "10.10.10.4"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := returnFilteredIPList(tt.clusterIPList, tt.toBeDeletedIPList); (res != nil) != tt.wantErr {
				fmt.Println(res)
			}
			logger.Error("is empty")
		})
	}
}
