/*
Package "classifier" provides an implementation of multi-class linear classifier.
*/
package classifier

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
	"github.com/mitsuse/olive/internal/validates"
)

const (
	INVALID_CLASS_SIZE         = validates.INVALID_CLASS_SIZE
	NON_FEATURE_MATRIX_PANIC   = validates.NON_FEATURE_MATRIX_PANIC
	INCOMPATIBLE_WEIGHTS_PANIC = validates.INCOMPATIBLE_WEIGHTS_PANIC
)

const (
	version    int = 0
	minVersion     = 0
	maxVersion     = 0
)

// Classifier is an implementation of multi-class linear classifier.
type Classifier struct {
	initialized bool
	weights     matrix.Matrix
}

// Create a new classifier with the given number of classes and dimensions of features.
// The generate classifier has a weight matrix with zeros.
func New(classSize, dimensions int) *Classifier {
	validates.ShouldBeOneOrMoreClasses(classSize)

	c := &Classifier{
		initialized: true,
		// TODO: Replace the weights with the immutable version of dense matrix.
		weights: dense.Zeros(classSize, dimensions),
	}

	return c
}

// Deserialize a classifier from the given reader.
func Deserialize(reader io.Reader) (*Classifier, error) {
	c := &Classifier{}

	if err := json.NewDecoder(reader).Decode(c); err != nil {
		return nil, err
	}

	return c, nil
}

// Serialize this classifier and write it by using the given writer.
func (c *Classifier) Serialize(writer io.Writer) error {
	return json.NewEncoder(writer).Encode(c)
}

func (c *Classifier) UnmarshalJSON(b []byte) error {
	if c.initialized {
		return errors.New(AlreadyInitializedError)
	}

	jsonObject := &classifierJson{}

	if err := json.Unmarshal(b, jsonObject); err != nil {
		return err
	}

	if jsonObject.Version < minVersion || maxVersion < jsonObject.Version {
		return errors.New(IncompatibleVersionError)
	}

	c.weights = jsonObject.Weights
	c.initialized = true

	return nil
}

func (c *Classifier) MarshalJSON() ([]byte, error) {
	jsonObject := &classifierJson{
		Version: version,
		Weights: dense.Convert(c.weights),
	}

	return json.Marshal(jsonObject)
}

// Return the weights matrix.
func (c *Classifier) Weights() matrix.Matrix {
	return c.weights
}

// Return the size of classes.
func (c *Classifier) ClassSize() int {
	return c.weights.Rows()
}

// Return the dimensions of features.
func (c *Classifier) Dimensions() int {
	return c.weights.Columns()
}

// Update the weight matrix.
// The new weight matrix should have same shape as the old one.
func (c *Classifier) Update(weights matrix.Matrix) *Classifier {
	validates.ShouldBeCompatibleWeights(c.weights, weights)

	c.weights = weights

	return c
}

// Classify the given feature matrix into one of the classes.
func (c *Classifier) Classify(feature matrix.Matrix) (class int) {
	validates.ShouldBeFeature(feature)

	_, class, _ = c.weights.Multiply(feature.Transpose()).Max()

	return class
}
