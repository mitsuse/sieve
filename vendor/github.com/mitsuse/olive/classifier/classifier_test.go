package classifier

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/mitsuse/matrix-go"
	"github.com/mitsuse/matrix-go/dense"
)

type constructionTest struct {
	classSize  int
	dimensions int
}

type classificationTest struct {
	feature matrix.Matrix
	class   int
}

func TestNewSucceeds(t *testing.T) {
	test := constructionTest{
		classSize:  4,
		dimensions: 8,
	}

	defer func() {
		if p := recover(); p == nil {
			return
		}

		t.Fatal("New should not panic for one or more classes and positive dimensions.")
	}()
	New(test.classSize, test.dimensions)
}

func TestNewPanicsForNonPositiveClassSize(t *testing.T) {
	test := constructionTest{
		classSize:  0,
		dimensions: 8,
	}

	defer func() {
		if p := recover(); p == INVALID_CLASS_SIZE {
			return
		}

		t.Fatal("Update should use \"validate.ShouldBeOneOrMoreClasses\".")
	}()
	New(test.classSize, test.dimensions)
}

func TestWeightsReturnsTheWeightsMatrix(t *testing.T) {
	test := constructionTest{
		classSize:  4,
		dimensions: 8,
	}

	c := New(test.classSize, test.dimensions)
	zeroMatrix := dense.Zeros(test.classSize, test.dimensions)

	if weights := c.Weights(); !weights.Equal(zeroMatrix) {
		t.Fatalf("A %d x %d matrix should be returned as the weights matrix.", 4, 8)
	}
}

func TestClassSizeReturnsTheSizeOfClasses(t *testing.T) {
	test := constructionTest{
		classSize:  4,
		dimensions: 8,
	}

	c := New(test.classSize, test.dimensions)

	if classSize := c.ClassSize(); classSize != test.classSize {
		t.Fatalf(
			"The size of classes should be %d, but %d is returned.",
			test.classSize,
			classSize,
		)
	}
}

func TestDimensionsReturnsTheDimensionsOfFeatures(t *testing.T) {
	test := constructionTest{
		classSize:  4,
		dimensions: 8,
	}

	c := New(test.classSize, test.dimensions)

	if dimensions := c.Dimensions(); dimensions != test.dimensions {
		t.Fatalf(
			"The dimensions of features should be %d, but %d is returned.",
			test.dimensions,
			dimensions,
		)
	}
}

func TestUpdateSucceedsForMatricesWithSameShape(t *testing.T) {
	classSize, dimensions := 4, 8
	test := dense.Zeros(classSize, dimensions)

	c := New(classSize, dimensions)

	defer func() {
		if p := recover(); p == nil {
			return
		}

		t.Fatal("Update should not panic for the weights as same shape as the old.")
	}()
	c.Update(test)
}

func TestUpdatePanicsByIncompatibleWeights(t *testing.T) {
	classSize, dimensions := 4, 8
	test := dense.Zeros(classSize+1, dimensions)

	c := New(classSize, dimensions)

	defer func() {
		if p := recover(); p == INCOMPATIBLE_WEIGHTS_PANIC {
			return
		}

		t.Fatal("Update should use \"validate.ShouldBeCompatibleWeights\".")
	}()
	c.Update(test)
}

func TestClassifyAssignsHighestScoredClassToFeature(t *testing.T) {
	classSize, dimensions := 4, 8

	test := classificationTest{
		feature: dense.New(1, 8)(0, 1, 0.5, -1, 0, 0, 2, 0),
		class:   2,
	}

	weights := dense.New(classSize, dimensions)(
		0, 0, 0, 0, 0, 0, 0, 0,
		0, -1, -1, 1, 0, 0, -1, 0,
		0, 1, 1, 0, 0, 0, 1, 0,
		0, 1, 1, 1, 0, 0, 1, 0,
	)

	c := New(classSize, dimensions).Update(weights)

	if class := c.Classify(test.feature); class != test.class {
		t.Fatalf(
			"Classifier should assign %d to the feature, but %d is assigned.",
			test.class,
			class,
		)
	}
}

func TestClassifyPanicsByNonFeatureMatrix(t *testing.T) {
	classSize, dimensions := 4, 8

	test := dense.Zeros(classSize, dimensions+1)

	defer func() {
		if p := recover(); p == NON_FEATURE_MATRIX_PANIC {
			return
		}

		t.Fatal("Update should use \"validate.ShouldBeFeature\".")
	}()
	New(classSize, dimensions).Classify(test)
}

func TestSerialize(t *testing.T) {
	c := New(4, 8)

	writer := bytes.NewBuffer([]byte{})

	if err := c.Serialize(writer); err != nil {
		t.Fatalf("An expected error occured on serialization: %s", err)
	}

	reader := bytes.NewReader(writer.Bytes())

	d, err := Deserialize(reader)

	if err != nil {
		t.Fatalf("An expected error occured on deserialization: %s", err)
	}

	if !c.Weights().Equal(d.Weights()) {
		t.Fatal("The origianl classifier is not equivalent to a deserialized classifier.")
	}
}

func TestUnmarshalJSONFailsWithAlreadyInitializedMatrix(t *testing.T) {
	c := New(4, 8)
	d := New(4, 8)

	b, _ := json.Marshal(c)

	if err := json.Unmarshal(b, d); err == nil || err.Error() != AlreadyInitializedError {
		t.Fatalf("Unmarshal can be applied to uninitialized classifier.")
	}
}

func TestUnmarshalJSONFailsWithIncompatibleVersion(t *testing.T) {
	c := &classifierJson{
		Version: 99999,
		Weights: dense.Zeros(4, 8),
	}

	d := &Classifier{}

	b, _ := json.Marshal(c)

	if err := json.Unmarshal(b, d); err == nil || err.Error() != IncompatibleVersionError {
		t.Fatalf("Unmarshal can be applied to compatible-version classifier.")
	}
}
