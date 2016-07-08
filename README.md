# jsonapi-errors [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/AmuzaTkts/jsonapi-errors) [![Build Status](https://travis-ci.org/AmuzaTkts/jsonapi-errors.svg?branch=master)](https://travis-ci.org/AmuzaTkts/jsonapi-errors) [![Coverage Status](https://coveralls.io/repos/github/AmuzaTkts/jsonapi-errors/badge.svg?branch=master)](https://coveralls.io/github/AmuzaTkts/jsonapi-errors?branch=master)

This package provides error bindings based on the [JSON API](http://jsonapi.org/format/#errors) reference.

The package provides two main structs that you can use on your application, the `Error` and `Bag` structs. When returning errors from your API you should return a `Bag` containing one or more errors.

```Go
bag := NewBagWithError(502, "Oops =(")
jsonStr, _ := json.Marshal(bag)
```

The above code will return the following JSON structure:

```JSON
{
  "errors": [
    {
      "detail": "Oops =(",
      "status": "502"
    }
  ],
  "status": "502"
}
```
