package still

import (
	"io"

	"github.com/mitsuse/olive/classifier"
)

type Still struct {
	extractor *Extractor
	c         *classifier.Classifier
}

func New() *Still {
	s := &Still{
		extractor: newExtractor(3),
		c:         classifier.New(2, 8),
	}

	return s
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
	feature := s.extractor.Extract(text)
	return s.c.Classify(feature) == 0
}
