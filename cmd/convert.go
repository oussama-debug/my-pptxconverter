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

package cmd

import (
	tea "github.com/charmbracelet/bubbletea"
	convert_prompt "github.com/oussama-debug/pptx/internal/tui/prompts/convert"
	models "github.com/oussama-debug/pptx/internal/tui/prompts/models"
	"github.com/spf13/cobra"
)

func convert() *cobra.Command {
	convertModelQuestions := []models.Question{
		*models.NewQuestion("Enter the path to the pptx file", models.QuestionString, []string{}),
		*models.NewQuestion("Choose the output format", models.QuestionChoice, []string{"html", "pdf"}),
	}
	convertPromptModel := convert_prompt.NewConvertPromptModel(convertModelQuestions)

	init := &cobra.Command{
		Use:     "convert",
		Short:   "convert the input presentation file",
		Long:    "convert the input presentation pptx file to another format",
		Example: "pptx convert",
		Aliases: []string{"c", "convert"},
		RunE: func(cmd *cobra.Command, args []string) error {
			p := tea.NewProgram(convertPromptModel)
			if _, err := p.Run(); err != nil {
				return tea.ErrProgramKilled
			}
			return nil
		},
	}

	return init
}
