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
	models "github.com/oussama-debug/pptx/internal/prompts/models"
)

func (c *ConvertPromptModel) OnUpdateKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	if c.complete {
		return c, tea.Quit
	}

	current := &c.questions[c.index]
	var cmd tea.Cmd

	switch msg.String() {
	case "q", "esc", "ctrl+c":
		c.quitting = true
		return c, tea.Quit
	case "enter":
		if c.questions[c.index].GetGenre() == models.QuestionChoice {
			current.SetAnswer(c.questions[c.index].GetChoices()[c.cursor])
		} else {
			current.SetAnswer(c.answerField.Value())
			c.answerField.SetValue("")
		}
		c.Next()
		if c.complete {
			return c, tea.Quit
		}
		return c, nil
	case "up":
		if c.questions[c.index].GetGenre() == models.QuestionChoice {
			c.cursor--
			if c.cursor < 0 {
				c.cursor = len(c.questions[c.index].GetChoices()) - 1
			}
		}
		return c, nil
	case "down":
		if c.questions[c.index].GetGenre() == models.QuestionChoice {
			c.cursor++
			if c.cursor >= len(c.questions[c.index].GetChoices()) {
				c.cursor = 0
			}
		}
		if c.complete {
			return c, tea.Quit
		}
		return c, nil
	default:
		c.answerField, cmd = c.answerField.Update(msg)
		c.quitting = false
		return c, cmd
	}
}
