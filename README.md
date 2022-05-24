# errhelp

Simple helper functions to make golang error handling easier. They reduce the boilerplate. For example. you can replace

```
user, err := createUser(data)
if err != nil {
  log.Printf("Error creating user: %v", err)
}
```

with

```
import eh "github.com/bfollek/errhelp"

user, err := createUser(data)
eh.LogIfError("creating user", err)
```

### Warning! Lightly used and tested.

### Available Helpers

* LogIfError - Log to the standard golang log.
* LogIfErrorNS - Log to a non-standard log you pass in as an io.Writer.
* FatalIfError - Log to the standard golang log, then stop the program.
* FatalIfErrorNS - Log to a non-standard log you pass in as an io.Writer, then stop the program.

