package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := NewSatellite([]string{"", "este", "es", "un", "mensaje"})
	s2 := NewSatellite([]string{"este", "", "un", "mensaje"})
	s3 := NewSatellite([]string{"", "", "es", "", "mensaje"})
	frequencies := CountFrequencies(s1.message, s2.message, s3.message)
	message := buildMessage(frequencies)
	fmt.Println(message)
}

func buildMessage(frequencies []map[string]int) string {
	var message []string
	for i, f := range frequencies {
		max_k := ""
		max_v := -1
		for k, v := range f {
			if k == "" {
				continue
			}

			if i-1 >= 0 && i-1 < len(message) && k == message[i-1] {
				continue
			}

			if v > max_v {
				max_v = v
				max_k = k
			}
		}
		if max_k != "" {
			message = append(message, max_k)
		}
	}
	return strings.Join(message, " ")
}

func CountFrequencies(msgs ...[]string) []map[string]int {
	maxLength := GetMaxMessageLenght(msgs...)
	words := []map[string]int{}
	for i := 0; i < maxLength; i++ {
		ranking := make(map[string]int, maxLength)
		for _, msg := range msgs {
			if i >= len(msg) {
				break
			}
			ranking[msg[i]]++
		}
		words = append(words, ranking)
	}
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
