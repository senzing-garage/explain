# explain

## :warning: WARNING: explain is still in development :warning: _

At the moment, this is "work-in-progress" with Semantic Versions of `0.n.x`.
Although it can be reviewed and commented on,
the recommendation is not to use it yet.

## Synopsis

`explain` is a command in the
[senzing-tools](https://github.com/Senzing/senzing-tools)
suite of tools.
It is used to explain aspects of `senzing-tools`.

[![Go Reference](https://pkg.go.dev/badge/github.com/senzing/explain.svg)](https://pkg.go.dev/github.com/senzing/explain)
[![Go Report Card](https://goreportcard.com/badge/github.com/senzing/explain)](https://goreportcard.com/report/github.com/senzing/explain)
[![go-test.yaml](https://github.com/Senzing/explain/actions/workflows/go-test.yaml/badge.svg)](https://github.com/Senzing/explain/actions/workflows/go-test.yaml)
[![License](https://img.shields.io/badge/License-Apache2-brightgreen.svg)](https://github.com/Senzing/explain/blob/main/LICENSE)

## Overview

`explain` performs the following:

1. Explains an error.
   Example:

    ```console
    senzing-tools explain --error-id senzing-60010032
    ```

## Install

1. The `explain` command is installed with the
   [senzing-tools](https://github.com/Senzing/senzing-tools)
   suite of tools.
   See senzing-tools [install](https://github.com/Senzing/senzing-tools#install).

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

### Using Docker

This usage shows how to initialze a database with a Docker container.

1. :pencil2: Run `senzing/senzing-tools`.
   Example:

    ```console
    docker run \
        --env SENZING_TOOLS_COMMAND=explain \
        --env SENZING_TOOLS_ERROR_ID=senzing-60010032 \
        --rm \
        senzing/senzing-tools
    ```

1. See [Parameters](#parameters) for additional parameters.

### Parameters

- **[LD_LIBRARY_PATH](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#ld_library_path)**
- **[SENZING_TOOLS_COMMAND](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_command)**
- **[SENZING_TOOLS_ERROR_ID](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_error_id)**
- **[SENZING_TOOLS_TTY_ONLY](https://github.com/Senzing/knowledge-base/blob/main/lists/environment-variables.md#senzing_tools_tty_only)**

## References

- [Command reference](https://hub.senzing.com/senzing-tools/senzing-tools_explain.html)
- [Development](docs/development.md)
- [Errors](docs/errors.md)
- [Examples](docs/examples.md)
