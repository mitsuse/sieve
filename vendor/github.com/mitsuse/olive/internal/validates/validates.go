package validates

import (
	. "github.com/mitsuse/matrix-go"
)

const (
	INVALID_CLASS_SIZE         = "INVALID_CLASS_SIZE"
	NON_FEATURE_MATRIX_PANIC   = "NON_FEATURE_MATRIX_PANIC"
	INCOMPATIBLE_WEIGHTS_PANIC = "INCOMPATIBLE_WEIGHTS_PANIC"
)

func ShouldBeOneOrMoreClasses(classSize int) {
	if classSize > 0 {
		return
	}

	panic(INVALID_CLASS_SIZE)
}

func ShouldBeFeature(matrix Matrix) {
	if matrix.Rows() == 1 {
		return
	}

	panic(NON_FEATURE_MATRIX_PANIC)
}

func ShouldBeCompatibleWeights(old, update Matrix) {
	oldClassSize, oldDimensions := old.Shape()
	updateClassSize, updateDimensions := update.Shape()

	if oldClassSize == updateClassSize && oldDimensions == updateDimensions {
		return
	}

	panic(INCOMPATIBLE_WEIGHTS_PANIC)
}
