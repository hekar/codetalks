# CodeTalks

[![work in progress](https://img.shields.io/badge/work%20in-progress-brightred.svg?style=flat)](https://en.wikipedia.org/wiki/Work_in_process)

## Install

Clone the repo:

```
$ git clone git@github.com:hekar/codetalks.git $GOPATH/src/github.com/hekar/codetalks
$ cd $GOPATH/src/github.com/hekar/codetalks
```

Install dependencies:

```
$ make install
```

## Run development

Start dev server:

```
$ make serve
```

[http://localhost:5001/](http://localhost:5001/) 

## Build

Install dependencies and type `NODE_ENV=production make build`. This rule is producing webpack build and regular golang build after that. Result you can find at `$GOPATH/bin`. Note that the binary will be named **as the current project directory**.

## License
[MIT](./LICENSE)
