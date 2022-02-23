package main

import "testing"

func TestAddHTTPPrefix(t *testing.T) {
	tables := []struct {
		url      string
		expected string
	}{
		{
			url:      "http://google.com",
			expected: "http://google.com",
		},
		{
			url:      "google.com",
			expected: "http://google.com",
		},
		{
			url:      "random.com",
			expected: "http://random.com",
		},
	}

	for i := range tables {
		actual := addHTTPPrefix(tables[i].url)
		if actual != tables[i].expected {
			t.Errorf("expected = %s, got = %s", tables[i].expected, actual)
			t.FailNow()
		}
	}
}
