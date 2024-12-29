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

package models

type QuestionGenre int

const (
	QuestionString QuestionGenre = iota
	QuestionChoice
)

type Question struct {
	question string
	answer   string
	choices  []string
	genre    QuestionGenre
}

func NewQuestion(question string, genre QuestionGenre, choices []string) *Question {
	return &Question{
		question: question,
		genre:    genre,
		choices:  choices,
	}
}

func (q *Question) GetQuestion() string {
	return q.question
}

func (q *Question) GetAnswer() string {
	return q.answer
}

func (q *Question) GetChoices() []string {
	return q.choices
}

func (q *Question) GetGenre() QuestionGenre {
	return q.genre
}

func (q *Question) SetAnswer(answer string) {
	q.answer = answer
}

func (q *Question) SetChoices(choices []string) {
	q.choices = choices
}

func (q *Question) SetQuestion(question string) {
	q.question = question
}

func (q *Question) SetGenre(genre QuestionGenre) {
	q.genre = genre
}
