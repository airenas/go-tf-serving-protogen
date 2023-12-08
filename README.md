# go-tf-serving-protogen

Welcome to the `go-tf-serving-protogen` repository, where I make TensorFlow Serving and Protobufs in Go a little less intimidating.

## Versioning

The versioning of this library aligns with TensorFlow versions to provide clarity on compatibility:

## Compatibility Matrix

Here's a simple compatibility matrix to help you choose the right version based on your TensorFlow Serving version:

| Library Version | TensorFlow/Serving Compatibility  |
| --------------- | ------------------------ |
| v2.14.1         | 2.14.1                   |
| v2.14.0         | 2.14.0                   |
| v2.13.0         | 2.13.0                   |
| v2.12.0         | 2.12.0                   |
| v2.11.0         | 2.11.0                   |
| v2.10.0         | 2.10.0                   |

## Installation

```bash
go get github.com/airenas/go-tf-serving-protogen/v2
## or better 
go get github.com/airenas/go-tf-serving-protogen/v2@<tf-serving-version>
```

## Generating protoc code for new version

### Requirements
Just:

- git
- make
- docker
- go

When preparing for a new TensorFlow version, follow these steps:

```bash
## select version
export tfserving_version=2.14.0

## generate code
make clean generate tf=${tfserving_version}

## comit code and do 
git tag v${tfserving_version}
git push origin v${tfserving_version}

```

## Examples

Explore the library usage by checking out the [examples](examples) folder. 

## Contributing

I welcome contributions! If you encounter issues or have ideas for improvements, please feel free to open an issue or submit a pull request.
