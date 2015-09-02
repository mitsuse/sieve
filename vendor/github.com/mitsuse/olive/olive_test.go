package olive

import (
	"testing"

	"github.com/mitsuse/matrix-go/dense"
	"github.com/mitsuse/olive/classifier"
)

func TestInstanceFailsForNonFeatureMatirx(t *testing.T) {
	m := dense.New(2, 3)(
		0, 1, 2,
		1, 0, -1,
	)

	class := 4

	defer func() {
		if p := recover(); p == classifier.NON_FEATURE_MATRIX_PANIC {
			return
		}

		t.Fatal("NewInstance should use \"validate.ShouldBeFeature\".")
	}()
	NewInstance(m, class)
}

func TestInstanceFeature(t *testing.T) {
	feature := dense.New(1, 6)(0, 1, 2, 1, 0, -1)
	class := 4
	instance := NewInstance(feature, class)

	m := dense.New(1, 6)(0, 1, 2, 1, 0, -1)

	if instance.Feature().Equal(m) {
		return
	}

	t.Fatal("The feature vector should not be modified on creating instance.")
}

func TestInstanceclass(t *testing.T) {
	feature := dense.New(1, 6)(0, 1, 2, 1, 0, -1)
	class := 4
	instance := NewInstance(feature, class)

	if instance.Class() == class {
		return
	}

	t.Fatal("The class should not be modified on creating instance.")
}
