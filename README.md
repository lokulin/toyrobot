Toy Robot
=========

[![Join the chat at https://gitter.im/lokulin/toyrobot](https://badges.gitter.im/lokulin/toyrobot.svg)](https://gitter.im/lokulin/toyrobot?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

[![Build Status](https://travis-ci.org/lokulin/toyrobot.svg)](https://travis-ci.org/lokulin/toyrobot)

Toy Robot is a minimalistic toy robot golang implementation conforming to the specifications from the [problem description](PROBLEM.md).

Install
-------

```
go get bitbucket.org/lokulin/toyrobot
go install bitbucket.org/lokulin/toyrobot
```


Testing
-------

Tests are supplied as go tests and a set of example input files.

To run tests:

```
go test bitbucket.org/lokulin/toyrobot/toyrobot
```

Running
-------

To run the toy robot against an input file:

```
toyrobot examples/example1.txt
```
