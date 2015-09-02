# Still

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![Version](https://img.shields.io/github/tag/mitsuse/still.svg?style=flat-square)][release]
[![Wercker](http://img.shields.io/wercker/ci/55c81ed0ec6f88065000b336.svg?style=flat-square)][wercker]
[![Coverage](https://img.shields.io/codecov/c/github/mitsuse/still/develop.svg?style=flat-square)][coverage]

[license]: LICENSE.txt
[release]: https://github.com/mitsuse/still/releases
[wercker]: https://app.wercker.com/project/bykey/ede506215c68441b2129ea4b5f9e44ee
[coverage]: https://codecov.io/github/mitsuse/still

A command-line tool to filter out needless text by using statistical classifier.


## Installation

For installation, execute the following command:

```
$ go get github.com/mitsuse/still/cmd/still
```

## Dependencies

Still includes the following libraries as vendored packages:

- [`mitsuse/matrix-go (v0.1.3)`][repo-matrix-go]
- [`mitsuse/olive (v0.1.1)`][repo-olive]
- [`codegangsta/cli`][repo-cli]


[repo-matrix-go]: https://github.com/mitsuse/matrix-go/tree/v0.1.3
[repo-olive]: https://github.com/mitsuse/matrix-go/tree/v0.1.1
[repo-cli]: https://github.com/codegangsta/cli/tree/5149e2fc0c3ae4bdd330358bc405e614a07cb8c9


## Usage

### Build a model

Still requires the model file to filter out text,
which consists of weights for the binary linear classifier.

To build the model, use `still build`:

```
$ still build -m model.still -e examples.json -i 3
```

`-m` represents the output path of a built model.
`-e` is used to specify the path of training data.
The JSON of training data should be a single array of objects which consists of "text" and "class" as follow:

```json
[
  {
    "text": "Go 1.5 is released https://blog.golang.org/go1.5 #go_blog",
    "class": 1
  },
  {
    "text": "OnHub – Google https://on.google.com/hub/",
    "class": 0
  }
]
```

The "text" field is used for example of classification.
The "class" field represents the correct label of classification result.

To set the number of iterations, use `-i`.
The training data are read N times when N is given as the value for `-i`.


### Test a model

Still can test the trained model on test data with the following command:

```
$ still test -m model.still -e examples.json
```

`-m` is used for the path of a training model.
`-e` represents the path of test data.
The test data has the same format as the training data.

Test command show [precision and recall][wikipedia-precision-recall].

[wikipedia-precision-recall]: https://en.wikipedia.org/wiki/Precision_and_recall


### Filter out text


Still is used as a filter for the standard IO like `grep`:

```
$ cat input.txt | still filter -m model.still
```

In the above command, The classification examples are lines of `input.txt`.
The option `-f` can be used to print filtered-out text instead.


## License

Please read [LICENSE.txt](LICENSE.txt).
