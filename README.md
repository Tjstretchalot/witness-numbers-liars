# Witness Numbers - Strongest Liars

This repository explores the question proposed by Brady Haran to Matt Parker in
his youtube video
[Witness Numbers (and the truthful 1,662,803)](https://www.youtube.com/watch?v=_MscGSN5J6o)
on his [Numberphile channel](https://www.youtube.com/channel/UCoxcjq-8xIDTYp3uz647V5A). The
question is reiterated around 13:20.

The video is assumed as context, but in short summary it discusses the Miller-Rabin
primality test. In the test an unknown number `n` can be determined to be composite
reliably using a witness number `a` and an easy-to-compute test. On the other hand,
the other outcome of the test is indeterminate - the number may be composite or it may
be prime. If the witness indicates the number may be prime but the number is not in
fact prime, then the witness can be said to be a liar.

A lot of research has already been done to determine what the most reliable witnesses
are - but which numbers are the _least_ reliable?

## Implementation

This is written in Golang due to its ease of concurrency. The source code is in
the `src/milrabliars` directory. The main algorithm is in `algorithm.go` and some
testable examples are available in `algorithm_test.go`. These can be
run using `go test`.

The logic for keeping track of the running totals is in `data.go`, and how they
are exported is handled in `output.go`

The entrypoint is `main.go`.

### Notes

- I originally wrote this with Montgomery multiplication, but found it to be slower
  than a naive modulo when benchmarking, so I reverted to the naive method.

- I originally wrote this with everything supporting uint64s, but it was very quickly
  clear that was not going to be relevant as this is not particularly fast. So I
  switched to regular ints to save memory and time.

## Output

The numbers who lie the most at somewhat arbitrary points in time are available
in the `out/snapshot-{num}.log` files.

A log for whenever an entry changes in the top 10 liars is available in the
`out/running.log` file.

## Running this yourself

This expects to be executed within the exact folder structure that this repository
is laid out. `main.go` will delete the `out` folder (if it already exists) and then
the program will regenerate it. To ensure everything is trivially reproducible there
is minimal configuration available.

This will use up to all the hardware cores available using `runtime.NumCPU()`,
which returns the number of logical CPUs.

- [Install the latest version of Go](https://go.dev/learn/).
- Clone this repository
- Open a terminal in the `src` directory.
- `go run`
