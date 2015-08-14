package still

import (
	"io"

	"github.com/mitsuse/olive"
	"github.com/mitsuse/olive/classifier"
	"github.com/mitsuse/olive/perceptron"
)

type Still struct {
	extractor *Extractor
	c         *classifier.Classifier
}

func Learn(exampleSeq []*Example) *Still {
	textSeq := make([]string, len(exampleSeq))
	for _, example := range exampleSeq {
		textSeq = append(textSeq, example.Text)
	}
	extractor := newExtractor(3, textSeq)

	instanceSeq := make([]*olive.Instance, len(exampleSeq))
	for _, example := range exampleSeq {
		instance := olive.NewInstance(
			extractor.Extract(example.Text),
			example.Class,
		)
		instanceSeq = append(instanceSeq, instance)
	}

	c := perceptron.New(10).Learn(
		classifier.New(2, extractor.Dimensions()),
		instanceSeq,
	)

	s := &Still{
		extractor: extractor,
		c:         c,
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
