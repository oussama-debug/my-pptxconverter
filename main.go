/*
 Copyright 2024 Oussama Jaaouani <oussama@jaaouani.com>

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

      https://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/oussama-debug/pptx/cmd"
)

var logoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)

const cmdLogo = `
     ____.                                        .__ 
    |    |____  _____    ____  __ _______    ____ |__|
    |    \__  \ \__  \  /  _ \|  |  \__  \  /    \|  |
/\__|    |/ __ \_/ __ \(  <_> )  |  // __ \|   |  \  |
\________(____  (____  /\____/|____/(____  /___|  /__|
              \/     \/                  \/     \/    
`

func main() {
	fmt.Printf("%s\n", logoStyle.Render(cmdLogo))
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
