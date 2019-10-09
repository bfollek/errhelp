# errhelp

Simple helper functions to make golang error handling easier. They reduce the boilerplate. For example. you can replace

```
usr, err := user.Current()
if err != nil {
  log.Fatalf("Error getting current user: %v", err)
}
```

with

```
import eh "github.com/bfollek/errhelp"

usr, err := user.Current()
eh.FatalIfError("getting current user", err)
```

Better docs coming someday
