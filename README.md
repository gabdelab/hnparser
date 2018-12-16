# hnparser
This projects aims at building a go service that can extract most frequent requests from an HN extract.


## Requirements

Have go installed on your machine https://golang.org/doc/install


## How-to build

```shell
make build
```

## How-to run

```shell
make run
```

## How-to test the API

See specs/api.yaml for a detailed description of the API, following OpenAPI specification


## Open questions

### How should we handle empty results ?

I tend to think that returning 404s on empty results is a bad practice, as it prevents the
monitoring systems to split between real user errors and missing data. So, my choice here
is to return 200s with empty body.
For dates in the future, this is more questionable; it is really a client - server
misalignment so it probably means the API or its implementation on client side is not clear
enough; for these ones I'll tend to return a specific error (and 404 seems the best to me).


### How should we handle size > nb results ?

I made the choice that the nb_returned_results = min(size, nb_results) rather than returning
an error, because the client shouldn't have to guess how many different results are available
in the dataset.
