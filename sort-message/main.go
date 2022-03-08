package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := NewSatellite([]string{"", "este", "es", "un", "mensaje"})
	s2 := NewSatellite([]string{"este", "", "un", "mensaje"})
	s3 := NewSatellite([]string{"", "", "es", "", "mensaje"})

	frequencies := CountFrequenciesInPosition(s1.message, s2.message, s3.message)
	message := buildMessage(frequencies)
	fmt.Println(message)
}

func buildMessage(frequencies []map[string]int) string {
	var message []string
	for i := 0; i < len(frequencies); i++ {
		word := GetBestWord(frequencies[i])
		for IsSameWordAsBefore(i, word, message) && len(frequencies[i]) > 0 {
			delete(frequencies[i], word)
			word = GetBestWord(frequencies[i])
		}
		if word != "" {
			message = append(message, word)
		}
	}
	return strings.Join(message, " ")
}

func IsSameWordAsBefore(i int, word string, message []string) bool {
	return i >= 1 && i <= len(message) && word == message[i-1]
}

func GetBestWord(f map[string]int) string {
	max_k := ""
	max_v := -1
	for k, v := range f {
		if v > max_v {
			max_v = v
			max_k = k
		}
	}
	return max_k
}

func CountFrequenciesInPosition(msgs ...[]string) []map[string]int {
	maxLength := GetMaxMessageLenght(msgs...)
	words := []map[string]int{}
	for i := 0; i < maxLength; i++ {
		ranking := make(map[string]int, maxLength)
		for _, msg := range msgs {
			if i >= len(msg) {
				continue
			}
			if msg[i] == "" {
				continue
			}
			ranking[msg[i]]++
		}
		if len(ranking) != 0 {
			words = append(words, ranking)
		}

	}
	fmt.Println(words)
	return words
}

func GetMaxMessageLenght(msgs ...[]string) int {
	max := len(msgs[0])
	for _, msg := range msgs {
		if len(msg) > max {
			max = len(msg)
		}
	}
	return max
}

type Satellite struct {
	message []string
}

func NewSatellite(message []string) *Satellite {
	return &Satellite{
		message: message,
	}
}

func (s *Satellite) HasMoreWords(i int) bool {
	return i < len(s.message)-1
}
