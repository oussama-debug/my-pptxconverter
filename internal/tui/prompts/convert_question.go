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
