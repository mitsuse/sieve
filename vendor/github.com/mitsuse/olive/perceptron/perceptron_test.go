package perceptron

import (
	"testing"

	"github.com/mitsuse/matrix-go/dense"
	"github.com/mitsuse/olive"
	"github.com/mitsuse/olive/classifier"
)

func TestPerceptronIsLearner(t *testing.T) {
	var _ olive.Learner = New(10)
}

func TestPerceptronLearnsClassification(t *testing.T) {
	classSize, dimensions := 2, 6
	iterations := 3

	feature := dense.New(1, dimensions)

	instances := []*olive.Instance{
		olive.NewInstance(feature(1, 1, 1, 0, 0, 0), 0),
		olive.NewInstance(feature(1, 1, 0, 0, 0, 0), 0),
		olive.NewInstance(feature(0, 0, 0, 1, 1, 1), 1),
		olive.NewInstance(feature(0, 0, 0, 0, 1, 1), 1),
	}

	c := New(iterations).Learn(classifier.New(classSize, dimensions), instances)

	for index, instance := range instances {
		class := c.Classify(instance.Feature())
		if class == instance.Class() {
			continue
		}

		t.Fatalf(
			"The instance %d should be classified into %d, but is classified into %d",
			index,
			instance.Class(),
			class,
		)
	}
}
