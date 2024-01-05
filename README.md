# explain

If you are beginning your journey with
[Senzing](https://senzing.com/),
please start with
[Senzing Quick Start guides](https://docs.senzing.com/quickstart/).

You are in the
[Senzing Garage](https://github.com/senzing-garage)
where projects are "tinkered" on.
Although this GitHub repository may help you understand an approach to using Senzing,
it's not considered to be "production ready" and is not considered to be part of the Senzing product.
Heck, it may not even be appropriate for your application of Senzing!

## :warning: WARNING: explain is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

`explain` is a command in the
[senzing-tools](https://github.com/senzing-garage/senzing-tools)
suite of tools.
It is used to explain aspects of `senzing-tools`.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing-garage/explain.svg)](https://pkg.go.dev/github.com/senzing-garage/explain)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing-garage/explain)](https://goreportcard.com/report/github.com/senzing-garage/explain)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/senzing-garage/explain/blob/main/LICENSE)

[![gosec.yaml](https://github.com/senzing-garage/explain/actions/workflows/gosec.yaml/badge.svg)](https://github.com/senzing-garage/explain/actions/workflows/gosec.yaml)
[![go-test-linux.yaml](https://github.com/senzing-garage/explain/actions/workflows/go-test-linux.yaml/badge.svg)](https://github.com/senzing-garage/explain/actions/workflows/go-test-linux.yaml)
[![go-test-darwin.yaml](https://github.com/senzing-garage/explain/actions/workflows/go-test-darwin.yaml/badge.svg)](https://github.com/senzing-garage/explain/actions/workflows/go-test-darwin.yaml)
[![go-test-windows.yaml](https://github.com/senzing-garage/explain/actions/workflows/go-test-windows.yaml/badge.svg)](https://github.com/senzing-garage/explain/actions/workflows/go-test-windows.yaml)

## Overview

`explain` performs the following:

1. Explains an error.
   Example:

    ```console
    senzing-tools explain --error-id senzing-60010032
    ```

## Install

1. The `explain` command is installed with the
   [senzing-tools](https://github.com/senzing-garage/senzing-tools)
   suite of tools.
   See senzing-tools [install](https://github.com/senzing-garage/senzing-tools#install).

## Use

```console
export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
senzing-tools explain [flags]
```

1. For options and flags:
    1. [Online documentation](https://hub.senzing.com/senzing-tools/senzing-tools_explain.html)
    1. Runtime documentation:

        ```console
        export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
        senzing-tools explain --help
        ```

1. In addition to the following simple usage examples, there are additional [Examples](docs/examples.md).

### Using command line options

1. :pencil2: Specify database using command line option.
   Example:

    ```console
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools explain
    ```

1. See [Parameters](#parameters) for additional parameters.

### Using environment variables

1. :pencil2: Specify database using environment variable.
   Example:

    ```console
    export SENZING_TOOLS_ERROR_ID=senzing-60010032
    export LD_LIBRARY_PATH=/opt/senzing/g2/lib/
    senzing-tools explain
    ```

1. See [Parameters](#parameters) for additional parameters.

### Parameters

- **[LD_LIBRARY_PATH](https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#ld_library_path)**
- **[SENZING_TOOLS_COMMAND](https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_command)**
- **[SENZING_TOOLS_ERROR_ID](https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_error_id)**
- **[SENZING_TOOLS_TTY_ONLY](https://github.com/senzing-garage/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_tty_only)**

## References

- [Command reference](https://hub.senzing.com/senzing-tools/senzing-tools_explain.html)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
