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
				{"sonhos", 2, 2.0 / 6.0},
				{"Tive", 1, 1.0 / 6.0},
				{"em", 1, 1.0 / 6.0},
				{"que", 1, 1.0 / 6.0},
				{"comia", 1, 1.0 / 6.0},
			},
		},
		{"Na rota contava lorotas sobre a rota",
			Frequencies{
				{"rota", 2, 2.0 / 7.0},
				{"Na", 1, 1.0 / 7.0},
				{"contava", 1, 1.0 / 7.0},
				{"lorotas", 1, 1.0 / 7.0},
				{"sobre", 1, 1.0 / 7.0},
				{"a", 1, 1.0 / 7.0},
			},
		},
		{
			"111t2t3Palavras111r444 texto",
			Frequencies{
				{"Palavras", 1, 1.0 / 5.0},
				{"texto", 1, 1.0 / 5.0},
				{"t", 2, 2.0 / 5.0},
				{"r", 1, 1.0 / 5.0},
			},
		},
		{"小蚱蜢: Старый друг лучше новых двух.",
			Frequencies{
				{"Старый", 1, 1.0 / 6.0},
				{"друг", 1, 1.0 / 6.0},
				{"лучше", 1, 1.0 / 6.0},
				{"новых", 1, 1.0 / 6.0},
				{"двух", 1, 1.0 / 6.0},
				{"小蚱蜢", 1, 1.0 / 6.0},
			},
		},
		{
			"", Frequencies{},
		},
	}
}
