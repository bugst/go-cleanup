## go.bug.st/cleanup

This library provides a `signal.Interrupt`/CTRL-C interruptable context for golang.

The intended usage is as follows:

```go
ctx, cancel := cleanup.InterruptableContext(wrappedContext)
defer cancel() // release the CTRL-C hook, see below

// Execute a long operation that may be interrupted
err := VeryLongOperation(ctx)

// If VeryLongOperation is interrupted with CTRL-C or if it gets a signal.Interrupt
// then an err "interrupted" is returned (instead of terminating the process).
if err != nil && err.Error() == "interrupted" {
    // We've been interrupted
    [...do something...]
}

// when defer cancel() is executed the signal.Interrupt is no more hooked to the
// context and the resource is freed
```

