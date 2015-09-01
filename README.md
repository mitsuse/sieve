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

### Dependencies

Still uses the following libraries:

- [`mitsuse/matrix-go (v0.1.3)`][repo-matrix-go]
- [`mitsuse/olive (v0.1.1)`][repo-olive]
- [`codegangsta/cli`][repo-cli]


[repo-matrix-go]: https://github.com/mitsuse/matrix-go/tree/v0.1.3
[repo-olive]: https://github.com/mitsuse/matrix-go/tree/v0.1.1
[repo-cli]: https://github.com/codegangsta/cli/tree/5149e2fc0c3ae4bdd330358bc405e614a07cb8c9


## License

Please read [LICENSE.txt](LICENSE.txt).
