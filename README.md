# Matrix

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)][godoc]
[![Version](https://img.shields.io/github/tag/mitsuse/matrix-go.svg?style=flat-square)][release]
[![Wercker](http://img.shields.io/wercker/ci/55672222ee357fac39001a2a.svg?style=flat-square)][wercker]
[![Coverage](https://img.shields.io/codecov/c/github/mitsuse/matrix-go.svg?style=flat-square)][coverage]

[license]: LICENSE.txt
[godoc]: http://godoc.org/github.com/mitsuse/matrix-go
[release]: https://github.com/mitsuse/matrix-go/releases
[wercker]: https://app.wercker.com/project/bykey/093a5cff0964f0f4ba5fcf9117e940e4
[coverage]: https://codecov.io/github/mitsuse/matrix-go

An experimental library for matrix manipulation implemented in [Golang][golang].

[golang]: http://golang.org/


## Motivations

1. Portability - Implement in pure Golang to achieve cgo-free.
1. Efficiency - Pursue performance as possible without highly optimized back-ends like blas.
1. Simplicity - Provide clean API.


## Installation

For installation, execute the following command:

```
$ go get github.com/mitsuse/matrix-go
```


## Features

### Matrix Types

Currently, the following types are implemented:

- mutable dense matrix


### Creation

Use `dense.New` to create a new dense matrix with given elements.

```go
// Create a 2 x 3 matrix.
m := dense.New(2, 3)(
    0, 1, 2,
    3, 4, 5,
)
```

To create zero matrix, call `dense.Zeros` instead.

```go
// Create a 2 x 3 zero matrix.
m := dense.Zeros(2, 3)
```


### Operations

#### Addition & Subtraction

Add a matrix to other with `(Matrix).Add`:

```go
m := dense.New(2, 3)(
    0, 1, 2,
    3, 4, 5,
)

n := dense.New(2, 3)(
    5, 4, 3,
    2, 1, 0,
)

r := dense.New(2, 3)(
    5, 5, 5,
    5, 5, 5,
)

// true
m.Add(n).Equal(r)
```

Similarly, `(Matrix).Subtract` is used for subtraction on two matrix.

When the receiver is mutable,
`(Matrix).Add` and `(Matrix).Subtract` return the receiver itself,
the elements of which is rewritten.


#### Matrix Multiplication

The product of two matrices can be calculated by `(Matrix).Multiply`.

```go
m := dense.New(3, 2)(
    0, 1,
    2, 3,
    4, 5,
)

n := dense.New(2, 1)(
    0,
    -1,
)

r := dense.New(3, 1)(
    -1,
    -3,
    -5,
)

m.Multiply(n).Equal(r)
```

Matrix multiplication always create a new matrix.
The type of the result matrix is same as the type of the receiver.


#### Scalar Multiplication

`(Matrix).Scalar` is available for Scalar multiplication (scalar-left multiplication).

```go
m := dense.New(2, 2)(
    0, 1,
    2, 3,
)

r := dense.New(2, 2)(
    0, -1,
    -2, -3,
)

// true
m.Scalar(-1).Equal(r)
```

For scalar-right multiplication, use `(Scalar).Multiply`.

```go
m := dense.New(2, 2)(
    0, 1,
    2, 3,
)

r := dense.New(2, 2)(
    0, -1,
    -2, -3,
)

// true
Scalar(-1).Multiply(m).Equal(r)
```

When the matrix used for scalar multiplication is mutable,
`(Matrix).Scalar` and `(Scalar).Multiply` rewrite elements of the matrix.


### Cursor

`Matrix` has several methods to iterate elements.
They return a value typed as `Cursor` which is a reference to the element to visit.

```go
m := dense.New(2, 3)(
    0, 1, 2,
    3, 4, 5,
)

// Create a cursor to iterate all elements of matrix m.
c := m.All()

// Check whether the element to visit exists or not.
for c.HasNext() {
    element, row, column := c.Get()

    fmt.Printf(
        "element = %d, row = %d, column = %d\n",
        element,
        row,
        column,
    )
}
```

Currently, three methods are implemented which return a cursor:

- `(Matrix).All`
- `(Matrix).NonZeros`
- `(Matrix).Diagonal`

For details, please read the documentation of
[`types.Matrix`](http://godoc.org/github.com/mitsuse/matrix-go/internal/types/#Matrix).


### Find the Maximum/Minimum Element

`Maxtrix` provides methods to find the maximum or minimum elements.
`(Matrix).Max` returns the one of maximum elements and its index (the row and column).
`(Matrix).Min` also does similarly.

```go
m := dense.New(3, 3)(
    0, 1, 2,
    3, 4, 5,
    4, 3, 2,
)

// Find the one of maximum elements.
element, row, column := m.Max()

// true
element == 5

// true
row == 1

// true
column == 2
```


### Create View of Matrix


`(Matrix).View(ro, co, rs, cs)` creates a view of matrix,
which can be used as a `(rs * cs)` matrix.
The access to the element of view at index `(i, j)` is mapped to
the element of the base matrix at index `(ro + i, co + j)`.

```go
m := dense.New(4, 4)(
    9, 9, 9, 9,
    9, 1, 2, 9,
    9, 3, 4, 9,
    9, 9, 9, 9,
).View(1, 1, 2, 2)

n := dense.New(2, 2)(
    1, 2,
    3, 4,
)

// true
m.Equal(n)
```

Operations between view and matrix is defined
on the same condition as the case of operations between matrices.

```go
m := dense.New(4, 4)(
    9, 9, 9, 9,
    9, 1, 2, 9,
    9, 3, 4, 9,
    9, 9, 9, 9,
).View(1, 1, 2, 2)

n := dense.New(2, 2)(
    8, 7,
    6, 5,
)

m.Add(n)

r := dense.new(2, 2)(
    9, 9,
    9, 9,
)

// true
m.Equal(r)
```

A view of mutable matrix updates the receiver view by addition.
The view refers to elements of the base matrix,
therefore the base matrix is also updated.

```go
m := dense.New(4, 4)(
    9, 9, 9, 9,
    9, 1, 2, 9,
    9, 3, 4, 9,
    9, 9, 9, 9,
)

n := dense.New(2, 2)(
    8, 7,
    6, 5,
)

m.View(1, 1, 2, 2).Add(n)

r := dense.new(4, 4)(
    9, 9, 9, 9,
    9, 9, 9, 9,
    9, 9, 9, 9,
    9, 9, 9, 9,
)

// true
m.Equal(r)
```

`(Matrix).Row` is an alias for `(Matrix).View` to create a row-vector view.
`(Matrix).Column` is also available for column-vector view.

```go
m := dense.New(3, 3)(
    0, 1, 2,
    3, 4, 5,
    6, 7, 8,
)

r := dense.New(1, 3)(3, 4, 5)
c := dense.New(3, 1)(1, 4, 7)

// true
m.Row(1).Equal(r)

// true
m.Column(1).Equal(c)
```


## More Details

Please read the [documentation][godoc].


## Related Projects

This is a list of projects using `mitsuse/matrix-go`.

- [Olive][repo-olive] - Online algorithms for machine learning implemented in Golang.

[repo-olive]: https://github.com/mitsuse/olive


## License

Please read [LICENSE.txt](LICENSE.txt).
