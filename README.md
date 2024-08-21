# explain

If you are beginning your journey with [Senzing],
please start with [Senzing Quick Start guides].

You are in the [Senzing Garage] where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: explain is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

`explain` is a command in the [senzing-tools] suite of tools.
It is used to explain aspects of `senzing-tools`.

[![Go Reference Badge]][Package reference]
[![Go Report Card Badge]][Go Report Card]
[![License Badge]][License]
[![go-test-linux.yaml Badge]][go-test-linux.yaml]
[![go-test-darwin.yaml Badge]][go-test-darwin.yaml]
[![go-test-windows.yaml Badge]][go-test-windows.yaml]

[![golangci-lint.yaml Badge]][golangci-lint.yaml]

## Overview

`explain` performs the following:

1. Explains a message.
   Example:

    ```console
    senzing-tools explain --message-id SZSDK60010032
    ```

## Install

1. The `explain` command is installed with the [senzing-tools] suite of tools.
   See [senzing-tools install].

## Use

```console
export LD_LIBRARY_PATH=/opt/senzing/er/lib/
senzing-tools explain [flags]
```

1. For options and flags:
    1. [Online documentation]
    1. Runtime documentation:

        ```console
        export LD_LIBRARY_PATH=/opt/senzing/er/lib/
        senzing-tools explain --help
        ```

1. In addition to the following simple usage examples, there are additional [Examples].

### Using command line options

1. :pencil2: Specify database using command line option.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/er/lib/
    senzing-tools explain
    ```

1. See [Parameters] for additional parameters.

### Using environment variables

1. :pencil2: Specify database using environment variable.
   Example:

    ```console
    export SENZING_TOOLS_MESSAGE_ID=senzing-60010032
    export LD_LIBRARY_PATH=/opt/senzing/er/lib/
    senzing-tools explain
    ```

1. See [Parameters] for additional parameters.

### Parameters

- **[LD_LIBRARY_PATH]**
- **[SENZING_TOOLS_COMMAND]**
- **[SENZING_TOOLS_MESSAGE_ID]**
- **[SENZING_TOOLS_TTY_ONLY]**

## References

1. [API documentation]
1. [Development]
1. [Errors]
1. [Examples]
1. [Package reference]

[API documentation]: https://pkg.go.dev/github.com/senzing-garage/explain
[Development]: docs/development.md
[Errors]: docs/errors.md
[Examples]: docs/examples.md
[Go Reference Badge]: https://pkg.go.dev/badge/github.com/senzing-garage/explain.svg
[Go Report Card Badge]: https://goreportcard.com/badge/github.com/senzing-garage/explain
[Go Report Card]: https://goreportcard.com/report/github.com/senzing-garage/explain
[go-test-darwin.yaml Badge]: https://github.com/senzing-garage/explain/actions/workflows/go-test-darwin.yaml/badge.svg
[go-test-darwin.yaml]: https://github.com/senzing-garage/explain/actions/workflows/go-test-darwin.yaml
[go-test-linux.yaml Badge]: https://github.com/senzing-garage/explain/actions/workflows/go-test-linux.yaml/badge.svg
[go-test-linux.yaml]: https://github.com/senzing-garage/explain/actions/workflows/go-test-linux.yaml
[go-test-windows.yaml Badge]: https://github.com/senzing-garage/explain/actions/workflows/go-test-windows.yaml/badge.svg
[go-test-windows.yaml]: https://github.com/senzing-garage/explain/actions/workflows/go-test-windows.yaml
[golangci-lint.yaml Badge]: https://github.com/senzing-garage/explain/actions/workflows/golangci-lint.yaml/badge.svg
[golangci-lint.yaml]: https://github.com/senzing-garage/explain/actions/workflows/golangci-lint.yaml
[LD_LIBRARY_PATH]: https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#ld_library_path
[License Badge]: https://img.shields.io/badge/License-Apache2-brightgreen.svg
[License]: https://github.com/senzing-garage/explain/blob/main/LICENSE
[Online documentation]: https://hub.senzing.com/senzing-tools/senzing-tools_explain.html
[Package reference]: https://pkg.go.dev/github.com/senzing-garage/explain
[Parameters]: #parameters
[Senzing Garage]: https://github.com/senzing-garage
[Senzing Quick Start guides]: https://docs.senzing.com/quickstart/
[SENZING_TOOLS_COMMAND]: https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_command
[SENZING_TOOLS_MESSAGE_ID]: https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_message_id
[SENZING_TOOLS_TTY_ONLY]: https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_tty_only
[senzing-tools install]: https://github.com/senzing-garage/senzing-tools#install
[senzing-tools]: https://github.com/senzing-garage/senzing-tools
[Senzing]: https://senzing.com/
