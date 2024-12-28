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

package prompts

import (
	"fmt"

	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ConvertPromptModel struct {
	spinner   spinner.Model
	quitting  bool
	questions []string
	width     int
	height    int
	err       error
}

func NewConvertPromptModel(questions []string) *ConvertPromptModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return &ConvertPromptModel{
		questions: questions,
		spinner:   s,
	}
}

func (c ConvertPromptModel) Init() tea.Cmd {
	return c.spinner.Tick
}

func (c ConvertPromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.width = msg.Width
		c.height = msg.Height
		return c, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			c.quitting = true
			return c, tea.Quit
		default:
			return c, nil
		}
	default:
		var cmd tea.Cmd
		c.spinner, cmd = c.spinner.Update(msg)
		return c, cmd
	}
}

func (c ConvertPromptModel) View() string {
	if c.err != nil {
		return c.err.Error()
	}
	str := fmt.Sprintf("%s Initializing the powerpoint converter CLI\n\n", c.spinner.View())
	if c.quitting {
		return str + "\n"
	}
	return str
}
