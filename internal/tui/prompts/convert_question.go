package prompts

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
