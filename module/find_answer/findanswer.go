package find_answer

import (
	"sort"
	"strings"

	"github.com/agnivade/levenshtein"
)

type Question struct {
	ID     string
	Title  string
	Option []string
	Answer string
	Type   string
}

type Answer struct {
	ID     string
	Option []string
	Answer string
}

func MatchAnswer(answerList []Answer, questionList []Question, questionRandomOption string) []Question {
	var matchedQuestions []Question
	for _, question := range questionList {
		for _, item := range answerList {
			if question.ID == item.ID {
				var answerStr string
				if question.Type == "单选题" || question.Type == "多选题" {
					if questionRandomOption == "false" {
						answerStr = item.Answer
					} else {
						answerStr = FindAnswer(question.Option, item.Option, item.Answer, question.Type)
					}
				} else {
					answerStr = item.Answer
				}
				matchedQuestions = append(matchedQuestions, Question{
					ID:     question.ID,
					Title:  question.Title,
					Option: question.Option,
					Answer: answerStr,
				})
			}
		}
	}
	return matchedQuestions
}

func FindAnswer(questionOption []string, answerOption []string, answer string, answerType string) string {
	var myAnswer string
	if answerType == "多选题" {
		var temp []string
		for _, item := range answerOption {
			if strings.Contains(answer, strings.Split(item, ".")[0]) {
				temp = append(temp, item)
			}
		}
		for _, item := range temp {
			index := DiffOption(item, questionOption)
			myAnswer += strings.Split(questionOption[index], " ")[0]
		}
		sortedAnswer := strings.Split(myAnswer, "")
		sort.Strings(sortedAnswer)
		myAnswer = strings.Join(sortedAnswer, "")
	} else if answerType == "单选题" {
		var temp string
		for _, item := range answerOption {
			if strings.Split(item, ".")[0] == answer {
				temp = item
			}
		}
		index := DiffOption(temp, questionOption)
		myAnswer = string(questionOption[index][0])
	} else {
		myAnswer = answer
	}
	return myAnswer
}

func DiffOption(item string, options []string) int {
	var sample []int
	for _, i := range options {
		temp := strings.Split(i, " ")[0]
		i = strings.ReplaceAll(strings.ReplaceAll(i, temp, ""), " ", "")
		sample = append(sample, levenshtein.ComputeDistance(i, strings.Split(item, string(item[1]))[1]))
	}
	maxIndex := 0
	for i := range sample {
		if sample[i] > sample[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}
