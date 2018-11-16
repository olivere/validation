# What is it?

[![Build Status](https://travis-ci.org/olivere/validation.svg?branch=master)](https://travis-ci.org/olivere/validation)

The `validation` package is a very simple way to validate entities before they e.g. get saved to a data store. It merely collects all errors found and wraps them in the `error` interface that Go provides. It has some functions to format those errors.

The `validation` package isn't meant to be used for end-users, so it misses some crucial features like the ability to localize the error messages.

Example:

```go
import (
    "fmt"
	"time"

	"github.com/olivere/validation"
)

type Person struct {
	Name string
	DoB  time.Time
}

func (p Person) Validate() error {
	var errs validation.Errors

	if p.Name == "" {
		errs = errs.Add("Name", "is missing")
	}
	if p.DoB.IsZero() {
		errs = errs.Add("DoB", "is missing")
	}

	if errs.HasErrors() {
		return errs
	}
	return nil
}

func main() {
	p := &Person{}

	err := p.Validate()
	if err != nil {
		fmt.Printf("Validation failed: %v\n", err)
		if errs, ok := err.(validation.Errors); ok {
			fmt.Println(errs.Join(", "))
		}
	}
}
```

# License

MIT. See [LICENSE](https://github.com/olivere/validation/blob/master/LICENSE) file.
