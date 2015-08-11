package validates

import (
	"testing"

	. "github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
)

type weightUpdateTest struct {
	old    Matrix
	update Matrix
}

func TestShouldBeOneOrMoreClassesSucceedsForPositiveSizedClasses(t *testing.T) {
	test := 4

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf(
				"classes-size validation should not cause any panic for positve size: %s",
				p,
			)
		}
	}()
	ShouldBeOneOrMoreClasses(test)
}

func TestShouldBeOneOrMoreClassesPanicsForNonPositiveSizedClasses(t *testing.T) {
	test := 0

	defer func() {
		if p := recover(); p != INVALID_CLASS_SIZE {
			t.Fatalf("%s should be caused, but %s causes", INVALID_CLASS_SIZE, p)
		}
	}()
	ShouldBeOneOrMoreClasses(test)
}

func TestShouldBeFeatureSucceedsForFeature(t *testing.T) {
	test := dense.Zeros(1, 1)

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf(
				"Feature validation should not cause any panic for feature matirx: %s",
				p,
			)
		}
	}()
	ShouldBeFeature(test)
}

func TestShouldBeFeaturePanicsForNonFeature(t *testing.T) {
	test := dense.Zeros(4, 1)

	defer func() {
		if p := recover(); p != NON_FEATURE_MATRIX_PANIC {
			t.Fatalf("%s should be caused, but %s causes", NON_FEATURE_MATRIX_PANIC, p)
		}
	}()
	ShouldBeFeature(test)
}

func TestShouldBeCompatibleWeightSucceedsForSameShapeMatrices(t *testing.T) {
	test := weightUpdateTest{
		old:    dense.Zeros(4, 8),
		update: dense.Zeros(4, 8),
	}

	defer func() {
		if p := recover(); p != nil {
			t.Fatalf(
				"Weights validation should not cause any panic for the same shape one: %s",
				p,
			)
		}
	}()
	ShouldBeCompatibleWeights(test.old, test.update)
}

func TestShouldBeCompatibleWeightPanicsForMatricsWithDifferentClassSize(t *testing.T) {
	test := weightUpdateTest{
		old:    dense.Zeros(4, 8),
		update: dense.Zeros(2, 8),
	}

	defer func() {
		if p := recover(); p != INCOMPATIBLE_WEIGHTS_PANIC {
			t.Fatalf("%s should be caused, but %s causes", INCOMPATIBLE_WEIGHTS_PANIC, p)
		}
	}()
	ShouldBeCompatibleWeights(test.old, test.update)
}

func TestShouldBeCompatibleWeightPanicsForMatricsWithDifferentDimensions(t *testing.T) {
	test := weightUpdateTest{
		old:    dense.Zeros(4, 8),
		update: dense.Zeros(4, 10),
	}

	defer func() {
		if p := recover(); p != INCOMPATIBLE_WEIGHTS_PANIC {
			t.Fatalf("%s should be caused, but %s causes", INCOMPATIBLE_WEIGHTS_PANIC, p)
		}
	}()
	ShouldBeCompatibleWeights(test.old, test.update)
}
