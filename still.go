package still

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/mitsuse/olive"
	"github.com/mitsuse/olive/classifier"
	"github.com/mitsuse/olive/perceptron"
)

const (
	version    int = 1
	minVersion     = 0
	maxVersion     = 0
)

const (
	AlreadyInitializedError  = "AlreadyInitializedError"
	IncompatibleVersionError = "IncompatibleVersionError"
)

type Still struct {
	intialized bool
	extractor  *Extractor
	c          *classifier.Classifier
}

func Learn(exampleSeq []*Example) *Still {
	textSeq := make([]string, 0, len(exampleSeq))
	for _, example := range exampleSeq {
		textSeq = append(textSeq, example.Text)
	}
	extractor := newExtractor(3, textSeq)

	instanceSeq := make([]*olive.Instance, 0, len(exampleSeq))
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
		intialized: true,
		extractor:  extractor,
		c:          c,
	}

	return s
}

func Deserialize(reader io.Reader) (*Still, error) {
	s := &Still{}

	if err := json.NewDecoder(reader).Decode(s); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Still) Serialize(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(s)
}

func (s *Still) UnmarshalJSON(b []byte) error {
	if s.intialized {
		return errors.New(AlreadyInitializedError)
	}

	jsonObject := &stillJson{}

	if err := json.Unmarshal(b, jsonObject); err != nil {
		return err
	}

	if jsonObject.Version < minVersion || maxVersion < jsonObject.Version {
		return errors.New(IncompatibleVersionError)
	}

	s.extractor = jsonObject.Extractor
	s.c = jsonObject.Classifier

	return nil
}

func (s *Still) MarshalJSON() ([]byte, error) {
	jsonObject := &stillJson{
		Version:    version,
		Extractor:  s.extractor,
		Classifier: s.c,
	}

	return json.Marshal(jsonObject)
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

type stillJson struct {
	Version    int                    `json:"version"`
	Extractor  *Extractor             `json:"extractor"`
	Classifier *classifier.Classifier `json:"classifier"`
}
