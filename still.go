package still

import (
	"io"

	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/olive/classifier"
)

type Still struct {
	extractor func(string) matrix.Matrix
	c         *classifier.Classifier
}

func Deserialize(reader io.Reader) (*Still, error) {
	// TODO: Implement.
	return nil, nil
}

func (s *Still) Serialize(writer io.Writer) error {
	// TODO: Implement.
	return nil
}

func (s *Still) FilterAll(inputSeq []string) []string {
	outputSeq := make([]string, 0, len(inputSeq))

	for _, text := range inputSeq {
		if s.Filter(text) {
			continue
		}

		outputSeq = append(outputSeq, text)
	}

	return outputSeq
}

func (s *Still) Filter(text string) bool {
	return s.c.Classify(s.extractor(text)) == 0
}
