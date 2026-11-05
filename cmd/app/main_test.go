package main

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"
)

func TestRunLogsGreeting(t *testing.T) {
	t.Parallel()

	var out bytes.Buffer
	lookup := func(name string) (string, bool) {
		value, ok := map[string]string{"APP_ENV": "production"}[name]
		return value, ok
	}

	if err := run(&out, lookup); err != nil {
		t.Fatalf("run() error = %v", err)
	}
	var event map[string]any
	if err := json.Unmarshal(out.Bytes(), &event); err != nil {
		t.Fatalf("run() output = %q; want JSON: %v", out.String(), err)
	}
	if event["greeting"] != "Hello, world!" || event["app_env"] != "production" {
		t.Fatalf("run() log = %v; want greeting and environment", event)
	}
}

func TestRunFailsFastOnInvalidConfig(t *testing.T) {
	t.Parallel()

	var out bytes.Buffer
	lookup := func(name string) (string, bool) {
		if name == "APP_ENV" {
			return "invalid", true
		}
		return "", false
	}

	err := run(&out, lookup)
	if err == nil || !strings.Contains(err.Error(), "APP_ENV") {
		t.Fatalf("run() error = %v; want APP_ENV error", err)
	}
	if out.Len() != 0 {
		t.Fatalf("run() output = %q; want no logs on invalid config", out.String())
	}
}
