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

package convert

import (
	tea "github.com/charmbracelet/bubbletea"
	convert_prompt "github.com/oussama-debug/pptx/internal/tui/prompts/convert"
)

func (c convert_prompt.ConvertPromptModel) OnUpdateDefault(msg *tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	c.spinner, cmd = c.spinner.Update(*msg)
	return c, cmd
}
