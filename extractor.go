package still

import (
	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
)

type Extractor struct {
	maxOrder int
	ngramMap map[string]int
}

func newExtractor(maxOrder int) *Extractor {
	e := &Extractor{
		maxOrder: maxOrder,
		ngramMap: make(map[string]int),
	}

	return e
}

func updateExtractor(e *Extractor, textSeq []string) *Extractor {
	m := make(map[string]int)

	for ngram, id := range e.ngramMap {
		m[ngram] = id
	}

	for _, text := range textSeq {
		for _, ngram := range extractNgrams(e.maxOrder, text) {
			m[ngram] = len(m)
		}
	}

	updatedExtractor := &Extractor{
		ngramMap: m,
	}

	return updatedExtractor
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
