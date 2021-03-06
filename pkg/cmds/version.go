// Copyright 2018 John Deng (hi.devops.io@gmail.com).
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

package cmds

import (
	"github.com/spf13/cobra"
	"fmt"
)

// NewCmdVersion creates a command for displaying the version of this binary
func NewCmdVersion(fullName string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Display client and server versions",
		Long:  "Display client and server versions.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s v2.0.0\n", fullName)
		},
	}

	return cmd
}
