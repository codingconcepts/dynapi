# dynoapi
A dynamic API allowing for easy web client testing.  A collection of example routes are provided with routes.yaml and this can be tailored (or removed) to taste.

## Installation

``` bash
$ go get -u github.com/codingconcepts/dynoapi
```

## Usage

``` bash
$ HOST=localhost PORT=1234 SSL=false dynoapi -c routes.yaml
```