# timeoutd

HTTP server that will just hang forever, for testing timeout error handling.

## Usage

```
$ go install github.com/seeruk/timeoutd/cmd/...
$ timeoutd
```

There's no output unless there's an error. Make any HTTP request and it should just hang there
indefinitely (or near enough at least...).
