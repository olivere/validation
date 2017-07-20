// Copyright 2017 Oliver Eilhard. All rights reserved.
// Use of this source code is governed by a MIT-license.
// See http://olivere.mit-license.org/license.txt for details.

package validation

import (
	"testing"
)

func TestErrors(t *testing.T) {
	var v Errors

	if want, got := false, v.HasErrors(); want != got {
		t.Errorf("expected HasErrors = %v, got %v", want, got)
	}
	if want, got := "", v.Error(); want != got {
		t.Errorf("expected Error = %q, got %q", want, got)
	}
	if want, got := "", v.String(); want != got {
		t.Errorf("expected String = %q, got %q", want, got)
	}

	v = v.Add("Field", "has a problem")

	if want, got := true, v.HasErrors(); want != got {
		t.Errorf("expected HasErrors = %v, got %v", want, got)
	}
	if want, got := "Invalid record with 1 error", v.Error(); want != got {
		t.Errorf("expected Error = %q, got %q", want, got)
	}
	if want, got := "Field has a problem", v.String(); want != got {
		t.Errorf("expected String = %q, got %q", want, got)
	}
	if want, got := "Field has a problem", v.Join(", "); want != got {
		t.Errorf("expected String = %q, got %q", want, got)
	}

	v = v.Add("Field2", "has a different problem")

	if want, got := true, v.HasErrors(); want != got {
		t.Errorf("expected HasErrors = %v, got %v", want, got)
	}
	if want, got := "Invalid record with 2 errors", v.Error(); want != got {
		t.Errorf("expected Error = %q, got %q", want, got)
	}
	if want, got := "Field has a problem\nField2 has a different problem", v.String(); want != got {
		t.Errorf("expected String = %q, got %q", want, got)
	}
	if want, got := "Field has a problem, Field2 has a different problem", v.Join(", "); want != got {
		t.Errorf("expected String = %q, got %q", want, got)
	}
}
