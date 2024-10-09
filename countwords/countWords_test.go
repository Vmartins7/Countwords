package countwords

import (
	"reflect"
	"testing"
)

type scenario struct {
	input          string
	expectedReturn Frequencies
}

func TestCountWords(t *testing.T) {
	scenarios := initScenarios()

	for _, scenario := range scenarios {
		receivedReturn := CountWords(scenario.input)
		sortByWordAsc(scenario.expectedReturn)
		if !reflect.DeepEqual(receivedReturn, scenario.expectedReturn) {
			t.Errorf("%s\n difere do esperado\n %s\n\n", receivedReturn, scenario.expectedReturn)
		}
	}
}

func initScenarios() []scenario {
	return []scenario{
		{"Tive sonhos em que comia sonhos",
			Frequencies{
				{"sonhos", 2},
				{"Tive", 1},
				{"em", 1},
				{"que", 1},
				{"comia", 1},
			},
		},
		{"Na rota contava lorotas sobre a rota",
			Frequencies{
				{"rota", 2},
				{"Na", 1},
				{"contava", 1},
				{"lorotas", 1},
				{"sobre", 1},
				{"a", 1},
			},
		},
		{
			"111t2t3Palavras111r444 texto",
			Frequencies{
				{"Palavras", 1},
				{"texto", 1},
				{"t", 2},
				{"r", 1},
			},
		},
		{"小蚱蜢: Старый друг лучше новых двух.",
			Frequencies{
				{"Старый", 1},
				{"друг", 1},
				{"лучше", 1},
				{"новых", 1},
				{"двух", 1},
				{"小蚱蜢", 1},
			},
		},
		{
			"", Frequencies{},
		},
	}
}
