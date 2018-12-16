# hnparser
This projects aims at building a go service that can extract most frequent requests from an HN extract.


## Requirements

Have go installed on your machine https://golang.org/doc/install


## How-to build

```shell
make build
```

## How-to run

Two options:

```shell
HN_LOGS=<my_file.tsv> make run
```

Or:

```shell
./hnparser -file-path <my_file.tsv>
```

The port is also configurable:
```shell
./hnparser -file-path <my_file.tsv> -port 8888
```

## How-to launch tests
```shell
make test
```

## How-to test the API

See specs/api.yaml for a detailed description of the API, following OpenAPI specification

## Third-party libraries

The exercice doesn't advise to use third-party libraries,
so here everything is done using go standard libs, except:
- github.com/pkg/errors which allows to wrap errors in order to keep context
- github.com/stretchr/testify which allows to easily assert and require in tests.
  Basically this only provides helpers to reduce the amount of boilerplate code
  in go tests, which is crucial in my opinion

Otherwise, the following libraries could be helpful:
- a router like gorilla/mux or gin-gonic/gin - I personnally would recommend gin for its ease of use


## TODO list

- add more handlers tests, in particular some with real data
- provide dockerfile for users who don't want to install Go
- return JSON with error messages in case of user errors
- the handlers tests should mock the usecase to test only the handler code


## Open questions

### Should we consider the fact that the results are sorted ?

I think it's not something we should rely on. It could reduce computation times, both when starting
and when getting queries, but we probably want the product to be more resilient than this.


### How should we handle empty results ?

I tend to think that returning 404s on empty results is a bad practice, as it prevents the
monitoring systems to split between real user errors and missing data. So, my choice here
is to return 200s with empty body.
For dates in the future, this is more questionable; it is really a client - server
misalignment so it probably means the API or its implementation on client side is not clear
enough; for these ones I could return a specific error (and 404 seems the best to me).


### How should we handle size > nb results ?

I made the choice that the nb_returned_results = min(size, nb_results) rather than returning
an error, because the client shouldn't have to guess how many different results are available
in the dataset.
