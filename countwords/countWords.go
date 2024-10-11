package countwords

import (
	"fmt"
	"sort"
	"strings"
	"unicode"
	"unicode/utf8"
)

type Frequency struct {
	Word              string
	AbsoluteFrequency uint32
	RelativeFrequency float64
}

type Frequencies []Frequency

// CountWords registra a frequência de cada palavra
func CountWords(s string) Frequencies {
	counts := make(map[string]uint32)
	size := len(s)
	var acm uint32
	for i := 0; i < size; {
		word, sw := nextWord(s, i, size)
		if len(word) > 0 {
			counts[word]++
			acm++
		}
		i = sw
	}

	return generateStats(counts, acm)
}

// generateStats preenche vetor de frequências
func generateStats(counts map[string]uint32, totalWords uint32) []Frequency {
	hertzes := make([]Frequency, 0, len(counts))
	stat := Frequency{}
	for word, count := range counts {
		stat.Word = word
		stat.AbsoluteFrequency = count
		stat.RelativeFrequency = float64(count) / float64(totalWords)
		hertzes = append(hertzes, stat)
	}
	sortByWordAsc(hertzes)
	return hertzes
}

func sortByWordAsc(hertzes Frequencies) {
	sort.Slice(hertzes, func(i, j int) bool {
		return hertzes[i].Word < hertzes[j].Word
	})
}

// nextWord retorna uma sequência de letras encontradas no texto
func nextWord(str string, startIndex, limit int) (string, int) {

	var i, sizeWord, startWord int

	i = startIndex
	for i < limit {
		r, n := utf8.DecodeRuneInString(str[i:])

		if unicode.IsLetter(r) {
			if sizeWord == 0 {
				startWord = i
			}
			sizeWord += n
		} else {
			if sizeWord > 0 {
				break
			}
		}
		i += n
	}

	return str[startWord : startWord+sizeWord], i
}

func (st Frequencies) String() string {
	builderStr := strings.Builder{}

	for _, w := range st {
		builderStr.WriteString(
			fmt.Sprintf("{%v, %v, %v}\n", w.Word, w.AbsoluteFrequency, w.RelativeFrequency))
	}

	return builderStr.String()
}
