package still

import (
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
			ngramMap[ngram] = len(ngramMap)
		}
	}

	updatedExtractor := &Extractor{
		ngramMap: ngramMap,
	}

	return updatedExtractor
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
	ngramSeq := make([]string, len(text)*maxOrder)

	for order := 1; order <= maxOrder; order++ {
		for begin := 0; begin < len(text)-order; begin++ {
			end := begin + order
			ngramSeq = append(ngramSeq, text[begin:end])
		}
	}

	return ngramSeq
}
