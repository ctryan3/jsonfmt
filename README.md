jsonfmt
=======

A formatting utility for JSON
----

`jsonfmt` is a formatting utility for files containing JSON data.

Given a file, `jsonfmt` indents the data and writes it to stdout, where it can be directed into a file or piped into a pager.

```
$ cat example.json
{"veggies":["lettuce","carrots","celery"],"fruits":["apple","orange","banana"]}
$ jsonfmt example.json
{
    "veggies": [
        "lettuce",
        "carrots",
        "celery"
    ],
    "fruits": [
        "apple",
        "orange",
        "banana"
    ]
}
```

`jsonfmt` automatically detects and handles JSON data wrapped in a callback (JSONP). This is useful if you've saved off a large JSONP response from a server into a file:

```
$ cat example.json
SOME.CALLBACK({"apples":true,"oranges":true,"pineapples":false})
$ jsonfmt example.json
SOME.CALLBACK({
    "apples": true,
    "oranges": true,
    "pineapples": false
})
```

#### Options

To format the file in-place by replacing it with its contents, use the `--replace`/`-r` option. Additionally, the `--sort`/`-s` flag can be used to recursively sort all keys alphabetically:

```
$ jsonfmt --sort example.json
{
    "fruits": [
        "apple",
        "banana",
        "orange"
    ],
    "veggies": [
        "carrots",
        "celery",
        "lettuce"
    ]
}
```
