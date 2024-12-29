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
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ConvertPromptModel struct {
	spinner     spinner.Model
	quitting    bool
	questions   []Question
	answerField textinput.Model
	index       int
	width       int
	height      int
	cursor      int
	styles      *ConvertPromptModelStyles
	err         error
}

type ConvertPromptModelStyles struct {
	BorderColor lipgloss.Color
	InputField  lipgloss.Style
}

var (
	arrowStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6"))
	inputStyle   = lipgloss.NewStyle().BorderForeground(lipgloss.Color("#01FAC6")).Width(80)
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6"))
	borderStyle  = lipgloss.Color("#01FAC6")
)

func DefaultStyles() *ConvertPromptModelStyles {
	s := new(ConvertPromptModelStyles)
	s.BorderColor = borderStyle
	s.InputField = inputStyle

	return s
}

func NewConvertPromptModel(questions []Question) *ConvertPromptModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = spinnerStyle

	answerField := textinput.New()
	answerField.Placeholder = "Type your powerpoint file here"
	answerField.Focus()

	styles := DefaultStyles()

	return &ConvertPromptModel{
		questions:   questions,
		spinner:     s,
		answerField: answerField,
		styles:      styles,
	}
}

func (c ConvertPromptModel) Init() tea.Cmd {
	return c.spinner.Tick
}

func (c *ConvertPromptModel) Next() {
	if c.index < len(c.questions)-1 {
		c.index++
	} else {
		c.index = 0
	}
}

func (c ConvertPromptModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	current := &c.questions[c.index]

	switch msg := msg.(type) {

	// on resize, update the width and height
	// it triggers also in the first init
	case tea.WindowSizeMsg:
		c.width = msg.Width
		c.height = msg.Height
		c.quitting = false
		return c, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			c.quitting = true
			return c, tea.Quit
		case "enter":
			if c.questions[c.index].genre == QuestionChoice {
				current.SetAnswer(c.questions[c.index].GetChoices()[c.cursor])
				c.Next()
			} else {
				current.SetAnswer(c.answerField.Value())
				c.answerField.SetValue("")
				c.Next()
			}
			return c, nil
		case "up":
			if c.questions[c.index].GetGenre() == QuestionChoice {
				c.cursor--
				if c.cursor < 0 {
					c.cursor = len(c.questions[c.index].GetChoices()) - 1
				}
			}
			return c, nil
		case "down":
			if c.questions[c.index].GetGenre() == QuestionChoice {
				c.cursor++
				if c.cursor >= len(c.questions[c.index].GetChoices()) {
					c.cursor = 0
				}
			}
			return c, nil
		default:
			var cmd tea.Cmd
			c.answerField, cmd = c.answerField.Update(msg)
			c.quitting = false
			return c, cmd
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
	if c.width == 0 && !c.quitting {
		return str
	}

	if c.questions[c.index].genre == QuestionChoice {
		str = fmt.Sprintf("%s\n", c.questions[c.index].question)
		for i, choice := range c.questions[c.index].choices {
			if i == c.cursor {
				str += fmt.Sprintf("  %s %s\n", arrowStyle.Render("➜"), choice)
			} else {
				str += fmt.Sprintf("    %s\n", choice)
			}
		}
		return str
	} else {
		return lipgloss.Place(
			c.width,
			c.height,
			lipgloss.Left,
			lipgloss.Left,
			lipgloss.JoinVertical(lipgloss.Left, c.questions[c.index].question, c.styles.InputField.Render(c.answerField.View())),
		)
	}

}
