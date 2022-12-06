package main

import (
	"os"
	"testing"

	"github.com/go-logr/logr"
	"k8s.io/klog/v2"
)

// EmptyLogger implements logr.Logging
type EmptyLogger struct{}

// Enabled always returns false
func (e *EmptyLogger) Enabled() bool {
	return false
}

// Info does nothing
func (e *EmptyLogger) Info(msg string, keysAndValues ...interface{}) {}

// Error does nothing
func (e *EmptyLogger) Error(err error, msg string, keysAndValues ...interface{}) {}

// V returns itself
func (e *EmptyLogger) V(level int) logr.Logger {
	return e
}

// WithValues returns itself
func (e *EmptyLogger) WithValues(keysAndValues ...interface{}) logr.Logger {
	return e
}

// WithName returns itself
func (e *EmptyLogger) WithName(name string) logr.Logger {
	return e
}

// disable logging in all tests
func TestMain(m *testing.M) {
	klog.SetLogger(&EmptyLogger{})
	os.Exit(m.Run())
}

func Test_needBypass(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{"bypass aec3_eae", []string{"...", "-codec:1", "eac3_eae", "-eaeprefix:1", "..."}, true},
		{"bypass ac3_eae", []string{"...", "-codec:1", "ac3_eae", "-eaeprefix:1", "..."}, true},
		{"bypass truehd_eae", []string{"...", "-codec:1", "truehd_eae", "-eaeprefix:1", "..."}, true},
		{"bypass mlp_eae", []string{"...", "-codec:1", "mlp_eae", "-eaeprefix:1", "..."}, true},
		{"don't bypass with ac3", []string{"...", "-codec:1", "ac3", "-prefix:1", "..."}, false},
		{"bypass http", []string{"...", "-i", "http://192.168.1.8:5004/auto/v665", "..."}, true},
		{"bypass live http", []string{"...", "-i", "http://127.0.0.1:32400/livetv/sessions/9fa55b21-381d-47b6-821b-9d80f41756e6/e10n1b33sku2unl3c5lnwlqq/index.m3u8?offset=0.000000&X-Plex-Incomplete-Segments=1&X-Plex-Token=xxxxxxxxxxxxxxxxxxxx", "..."}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := needBypass(tt.args); got != tt.want {
				t.Errorf("needBypass() = %v, want %v", got, tt.want)
			}
		})
	}
}
