# rest-api-go

### Usage Examples

##### GET `/`
**Response Code**: 200
```
{
    "version": "0.11.1",
    "plugins": {
        "enabled_in_cluster": [
            "cors"
        ],
        "available_on_server": {
            "response-transformer": true,
            "correlation-id": true,

... etc

}
```
##### GET `/quote/{id}`
**Response Code**: 200
```
{
    "ID": 2,
    "CreatedAt": "2018-01-20T22:36:14.840646-05:00",
    "UpdatedAt": "2018-01-20T22:36:14.840646-05:00",
    "DeletedAt": null,
    "quote": "Test quote2.",
    "AuthorID": 1,
    "author": {
        "ID": 1,
        "CreatedAt": "2018-01-20T22:35:58.192347-05:00",
        "UpdatedAt": "2018-01-20T22:35:58.192347-05:00",
        "DeletedAt": null,
        "first": "First",
        "last": "Last",
        "born": "2000-01-01T19:00:00-05:00",
        "died": "2010-01-02T19:00:00-05:00",
        "description": "Test description.",
        "biolink": "http://somesite.com"
    }
}
```
**Resonse Code**: 400
```
{
    "error": "Quote ID: 22 not found."
}
```
##### GET `/quotes/`
**Response Code**: 200
```
[
       {
           "ID": 2,
           "CreatedAt": "2018-01-20T22:36:14.840646-05:00",
           "UpdatedAt": "2018-01-20T22:36:14.840646-05:00",
           "DeletedAt": null,
           "quote": "Test quote2.",
           "AuthorID": 1,
           "author": {
               "ID": 1,
               "CreatedAt": "2018-01-20T22:35:58.192347-05:00",
               "UpdatedAt": "2018-01-20T22:35:58.192347-05:00",
               "DeletedAt": null,
               "first": "First",
               "last": "Last",
               "born": "2000-01-01T19:00:00-05:00",
               "died": "2010-01-02T19:00:00-05:00",
               "description": "Test description.",
               "biolink": "http://somesite.com"
           }
       }
   ]

```
##### GET `/author/{id}`
**Response Code**: 200
```
{
    "ID": 1,
    "CreatedAt": "2018-01-20T22:35:58.192347-05:00",
    "UpdatedAt": "2018-01-20T22:35:58.192347-05:00",
    "DeletedAt": null,
    "first": "First",
    "last": "Last",
    "born": "2000-01-01T19:00:00-05:00",
    "died": "2010-01-02T19:00:00-05:00",
    "description": "Test description.",
    "biolink": "http://somesite.com",
    "quotes": [
        {
            "ID": 1,
            "CreatedAt": "2018-01-20T22:36:10.788467-05:00",
            "UpdatedAt": "2018-01-20T22:36:10.788467-05:00",
            "DeletedAt": null,
            "quote": "Test quote1.",
            "AuthorID": 1
        },
        {
            "ID": 2,
            "CreatedAt": "2018-01-20T22:36:14.840646-05:00",
            "UpdatedAt": "2018-01-20T22:36:14.840646-05:00",
            "DeletedAt": null,
            "quote": "Test quote2.",
            "AuthorID": 1
        }
    ]
}
```
**Response Code**: 400
```
{
    "error": "Author ID: 11 not found."
}
```
##### GET `/authors/`
**Response Code**: 200
```
[
    {
        "ID": 1,
        "CreatedAt": "2018-01-20T22:35:58.192347-05:00",
        "UpdatedAt": "2018-01-20T22:35:58.192347-05:00",
        "DeletedAt": null,
        "first": "First",
        "last": "Last",
        "born": "2000-01-01T19:00:00-05:00",
        "died": "2010-01-02T19:00:00-05:00",
        "description": "Test description.",
        "biolink": "http://somesite.com",
        "quotes": [
            {
                "ID": 2,
                "CreatedAt": "2018-01-20T22:36:14.840646-05:00",
                "UpdatedAt": "2018-01-20T22:36:14.840646-05:00",
                "DeletedAt": null,
                "quote": "Test quote2.",
                "AuthorID": 1
            }
        ]
    }
]
```

##### POST `/quote/`
**Response Code**: 201
**BODY**
```
{
	"quote": "Test quote3.",
	"authorid": 1
}
```
```
{
    "status": "Quote ID: 3 created for authorID: 1."
}
```
##### POST `/author/`
**Response Code**: 201
**BODY**
```
{
	"first": "First",
	"last": "Last2",
	"born": "2000-01-02T00:00:00Z",
	"died": "2010-01-03T00:00:00Z",
	"description": "Test description.",
	"biolink": "http://somesite.com"
}
```
```
{
    "status": "Author ID: 2 created."
}
```

##### DELETE `/quote/{id}`
**Response Code**: 200
```
{
    "status": "Quote ID: 1 deleted."
}
```
##### DELETE `/author/{id}`
**Response Code**: 200
```
{
    "status": "Author ID: 1 deleted."
}
```

##### GET `/health/`
```
{
    "reference": "https://golang.org/pkg/runtime/#MemStats",
    "alloc": 3495,
    "total-alloc": 4304,
    "sys": 9030,
    "numgc": 8
}
```
##### GET `/ready/`
```
{
    "ready": "OK"
}
```
##### GET `/version/`
```
{
    "version": "v0.01",
    "release-date": "2018-01-06T18:00:00"
}
```
