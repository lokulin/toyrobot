Toy Robot
=========

Toy Robot is a minimalistic toy robot golang implementation conforming to the specifications from the [problem description](PROBLEM.md).

Testing
-------

Tests are supplied as go tests and a set of example input files.

To run rspec tests:

```
bundle
bundle exec rspec -f d
```

Running
-------

To run the toy robot against an input file:

```
go install
toyrobot examples/example1.txt
```
