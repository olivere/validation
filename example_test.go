// Copyright 2017 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package validation_test

import (
	"fmt"
	"time"

	"github.com/olivere/validation"
)

func Example() {
	p := struct {
		Name string
		DoB  time.Time
	}{
		Name: "",
	}

	// Validation, typically done in e.g. a Validate method on the model
	var err validation.Errors

	if p.Name == "" {
		err = err.Add("Name", "is missing")
	}
	if p.DoB.IsZero() {
		err = err.Add("DoB", "is missing")
	}

	if err.HasErrors() {
		fmt.Println("Person has issues.")
		fmt.Printf("There are %d issues:\n", len(err))
		fmt.Printf("%s\n", err.Join("\n"))
	}

	// Output:
	// Person has issues.
	// There are 2 issues:
	// Name is missing
	// DoB is missing
}
