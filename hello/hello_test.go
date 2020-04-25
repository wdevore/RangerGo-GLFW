package main

import (
	"testing"

	rmath "github.com/wdevore/ranger/math"
)

func Test_Vector3(t *testing.T) {
	v := rmath.NewVector3()
	t.Log("\n" + v.String())
}
