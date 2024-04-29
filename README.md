# Capp Gin App

This code provides functionality to run a simple HTTP server using the Gin framework. The purpose is to have a simple app to be used for [`Capp`](https://github.com/dana-team/container-app-operator), which also supports `SIGINT` and `SIGKILL`.

## Getting Started

To run this code, use:

```bash
$ make run
```

To test this case, use:

```bash
$ make test
```

To build an image from this code, use:

```bash
$ make docker-build docker-push IMG=<registry>/test-gin-app:<version>
```