# jsonapi-errors [![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](http://godoc.org/github.com/AmuzaTkts/jsonapi-errors) [![Build Status](https://travis-ci.org/AmuzaTkts/jsonapi-errors.svg?branch=master)](https://travis-ci.org/AmuzaTkts/jsonapi-errors) [![Coverage Status](https://coveralls.io/repos/github/AmuzaTkts/jsonapi-errors/badge.svg?branch=master)](https://coveralls.io/github/AmuzaTkts/jsonapi-errors?branch=master)

This package provides error bindings based on the [JSON API](http://jsonapi.org/format/#errors)
reference.

The package provides two main structs that you can use on your application, the
`Error` and `Bag` structs. When returning errors from your API you should return
a `Bag` containing one or more errors.

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

## Multiple errors
This package adds the possibility to add multiple errors with different status
codes. The package will check for the range of the errors and will set the main
status key to the lower bound of the error class.

Eg: If add an error `501` and `502`, the main status key will be `500`.

```Go
bag := NewBag()
bag.AddError(501, "Server Error 1")
bag.AddError(502, "Server Error 2")

jsonStr, _ := json.Marshal(bag)
```

Will return:
```JSON
{
    "errors": [
        {
            "detail": "Server Error 1",
            "status": "501"
        },
        {
            "detail": "Server Error 2",
            "status": "502"
        }
    ],
    "status": "500"
}
```

It's also possible to have errors of different classes(`400` and `500` for example),
in this case the package will silently return `400` as the main status.

```Go
bag := NewBag()
bag.AddError(401, "Client Error 1")
bag.AddError(502, "Server Error 1")

jsonStr, _ := json.Marshal(bag)
```

```JSON
{
    "errors": [
        {
            "detail": "Client Error 1",
            "status": "401"
        },
        {
            "detail": "Server Error 1",
            "status": "502"
        }
    ],
    "status": "400"
}
```
