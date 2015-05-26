# chello - command line interface to *hello* services

*chello* is a small command line tool which can interact with services which
expose it's api via [hello][]. a sample service to toy around with is
[hello_pingpong][].

## usage

    $> chello <Method> [<arg1>, ... <argN>]

the argument can be given in this forms:

* key-value pairs like key1=val1 key2=val2 yield a JSON-object {"key1":
  "val1", "key2": "val2"}.  if a value can be parsed as a float, it will be
  encoded as a JSON-number.

* a string starting with a '{' yields a JSON-object. example:

    '{"key": "value"}'

* a list of strings yields a JSON-array of strings


## building

dependencies: 

* a golang compiler
* libzmq-dev (something that satisfies `pkg-config --libs libzmq`)

### make

if *make* is installed, you might just:

    $> make

### go get

    $> go get -v github.com/mgumz/chello



[hello]: https://github.com/travelping/hello
[hello_pingpong]: https://github.com/liveforeverx/hello_pingpong