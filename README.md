# Banzai Linter

[![CircleCI](https://circleci.com/gh/banzaicloud/banzailint.svg?style=svg)](https://circleci.com/gh/banzaicloud/banzailint)
[![Go Report Card](https://goreportcard.com/badge/github.com/banzaicloud/banzailint?style=flat-square)](https://goreportcard.com/report/github.com/banzaicloud/banzailint)

**Custom lint rules for Banzai Cloud code.**


## Installation

Download prebuilt binaries from the [Releases](https://github.com/banzaicloud/banzailint/releases/latest) page.

A common use case for this linter is to be used in a *Makefile*:

```makefile
OS = $(shell uname | tr A-Z a-z)

# ...

BANZAILINT_VERSION = 0.0.1

# ...

bin/banzailint: bin/banzailint-${BANZAILINT_VERSION}
	@ln -sf banzailint-${BANZAILINT_VERSION} bin/banzailint
bin/banzailint-${BANZAILINT_VERSION}:
	@mkdir -p bin
	curl -L https://github.com/banzaicloud/banzailint/releases/download/v${BANZAILINT_VERSION}/banzailint_${BANZAILINT_VERSION}_${OS}_amd64.tar.gz | tar -zOxf - banzailint > ./bin/banzailint-${BANZAILINT_VERSION} && chmod +x ./bin/banzailint-${BANZAILINT_VERSION}

.PHONY: lint
lint: bin/banzailint
	bin/banzailint
```

Alternatively, you can install it from source:

```bash
go get github.com/banzaicloud/banzailint
```


## License

The Apache 2.0 License. Please see [License File](LICENSE) for more information.
