package main

import "testing"

func TestRun(t *testing.T) {
	testCases := []struct {
		src string
		ok  bool
	}{
		{"testdata/t1.json", true},
		{"testdata/t2.json", true},
		{"testdata/t3.json", false},
		{"testdata/t4.json", false},
	}

	for _, tc := range testCases {
		if err := run("testdata/example.proto", "foo.Bar", nil, tc.src); (err == nil) != tc.ok {
			t.Errorf("%q expected ok? %v, error: %+v", tc.src, tc.ok, err)
		}
	}
}
