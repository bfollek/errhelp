# errhelp

Simple helper functions for golang error handling. They reduce the boilerplate, e.g. you can replace

```
usr, err := user.Current()
if err != nil {
  log.Fatalf("Error getting current user: %v", err)
}
```

with

```
import blah blah

usr, err := user.Current()
eh.FatalIfError("getting current user", err)
```
