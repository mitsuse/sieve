package still

import (
	"encoding/json"

	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
)

type Extractor struct {
	maxOrder int
	ngramMap map[string]int
}

func newExtractor(maxOrder int, textSeq []string) *Extractor {
	ngramMap := make(map[string]int)

	for _, text := range textSeq {
		for _, ngram := range extractNgrams(maxOrder, text) {
			if _, exist := ngramMap[ngram]; exist {
				continue
			}
			ngramMap[ngram] = len(ngramMap)
		}
	}

	updatedExtractor := &Extractor{
		maxOrder: maxOrder,
		ngramMap: ngramMap,
	}

	return updatedExtractor
}

func (e *Extractor) UnmarshalJSON(b []byte) error {
	jsonObject := &extractorJson{}

	if err := json.Unmarshal(b, jsonObject); err != nil {
		return err
	}

	e.maxOrder = jsonObject.MaxOrder
	e.ngramMap = jsonObject.NgramMap

	return nil
}

func (e *Extractor) MarshalJSON() ([]byte, error) {
	jsonObject := &extractorJson{
		MaxOrder: e.maxOrder,
		NgramMap: e.ngramMap,
	}

	return json.Marshal(jsonObject)
}

func (e *Extractor) Dimensions() int {
	return len(e.ngramMap)
}

func (e *Extractor) Extract(text string) matrix.Matrix {
	feature := dense.Zeros(1, len(e.ngramMap))

	for _, ngram := range extractNgrams(e.maxOrder, text) {
		index, exist := e.ngramMap[ngram]
		if !exist {
			continue
		}

		feature.Update(0, index, 1)
	}

	return feature
}

func extractNgrams(maxOrder int, text string) []string {
	charSeq := []rune(text)

	ngramSeq := make([]string, 0, len(charSeq)*maxOrder)

	for order := 1; order <= maxOrder; order++ {
		for begin := 0; begin < len(charSeq)-order; begin++ {
			end := begin + order
			ngramSeq = append(ngramSeq, string(charSeq[begin:end]))
		}
	}

	return ngramSeq
}

type extractorJson struct {
	MaxOrder int            `json:"max_order"`
	NgramMap map[string]int `json:"ngram_map"`
}
