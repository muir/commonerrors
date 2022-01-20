package commonerrors

import (
	"fmt"
	"testing"

	"github.com/pkg/errors"
)

func TestIs(t *testing.T) {
	if !IsUsageError(errors.Wrap(UsageError(fmt.Errorf("inner")), "outer")) {
		t.Fail()
	}
	if !IsUsageError(UsageError(fmt.Errorf("inner"))) {
		t.Fail()
	}
	if IsUsageError(fmt.Errorf("inner")) {
		t.Fail()
	}
}
