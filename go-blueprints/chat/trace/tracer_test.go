package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("value from New is nil")
	} else {
		tracer.Trace("trace package!!")
		if buf.String() != "trace package!!\n" {
			t.Errorf(" '%s' had a wrong string output", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	var silentTracer Tracer = Off()
	silentTracer.Trace("Date")
}
