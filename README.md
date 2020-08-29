# Persian Text PreProcessor

![Go Version](https://img.shields.io/github/go-mod/go-version/go-ptpp/ptpp?style=flat-square)
![License](https://img.shields.io/github/license/go-ptpp/ptpp?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/gopkg.in/ptpp.v1?tab=doc)](https://pkg.go.dev/gopkg.in/ptpp.v1?tab=doc)

Persian Text PreProcessor is a tool to help search engines to improve their
results. Although this library is especially optimized for Persian language, it
can also be used for English or mixed-language texts.

## Getting Started

Use `go get` to install PTPP library into your awesome project:

```sh
$ go get -u gopkg.in/ptpp.v1
```

PTPP uses a simple memory model which must be trained before being used. In
order to train the model, use simple correctly-spelled words or phrases:

```go
var processor ptpp.Processor
processor.Train([]string{
    "bass guitar",
    "garbage collector",
    ...
})
```

Now, the trained `processor` is able to auto-correct misspelled words and form
composite keywords from search phrases:

```go
processor.Process(strings.NewReader("electric base guitarr"))
// Returns: []string{"electric", "bass guitar"}
```

## License

PTPP is published under MIT license.
