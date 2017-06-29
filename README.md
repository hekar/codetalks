# CodeTalks

## Routes

```
pip install httpie
```

```
http --json POST http://localhost:5000/api/v1/talk name=example emails:='["example@gmail.com"]'
http --json GET http://localhost:5000/api/v1/user/1
http --json POST http://localhost:5000/api/v1/user name=example emails:='["example@gmail.com"]'
```

## Install

Clone the repo:

```
$ git clone git@github.com:olebedev/go-starter-kit.git $GOPATH/src/github.com/<username>/<project>
$ cd $GOPATH/src/github.com/<username>/<project>
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

that's it. Open [http://localhost:5001/](http://localhost:5001/)(if you use default port) at your browser. Now you ready to start coding your awesome project.

## Build

Install dependencies and type `NODE_ENV=production make build`. This rule is producing webpack build and regular golang build after that. Result you can find at `$GOPATH/bin`. Note that the binary will be named **as the current project directory**.

## License
MIT
