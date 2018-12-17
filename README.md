# hnparser
This projects aims at building a go service that can extract most frequent requests from an HN extract.


## Requirements

Have go and dep installed on your machine:
- https://golang.org/doc/install
- https://github.com/golang/dep


## How-to build

Locally:
```shell
make build
```

Or with Docker:
```shell
make docker-build
```

## How-to run

Three options! With environment variable:

```shell
HN_LOGS=<my_file.tsv> make run
```

Or by passing an arg in command-line:

```shell
./hnparser -file-path <my_file.tsv>
```

Or with Docker:
```shell
 docker run -v <my_file.tsv>:/usr/share/results.tsv -p 8080:8080 hnparser
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


## Code organization

When starting the program, the file is loaded and parsed (see parser.go).
Then, all HTTP requests are handled by handler.go, which gets all parameters and formats
the response. The calculation are done in counter.go, and our data structures are shown in
data_structures.go.
It could be time to have a bit more organization in the repo, for instance
```shell
main.go
http/
   handlers.go
   handlers_test.go
usecases/
   counter.go
   counters_test.go
   parser.go
   parser_test.go
model/
   data_structures.go
```
I personally think the repo can be flat until it gets really big, which is in my opinion not the
case yet.


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
- return JSON with error messages in case of user errors
- the handlers tests should mock the usecase to test only the handler code
- replace fmt.Println by a proper syslog logging


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
