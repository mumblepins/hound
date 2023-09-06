package vcs

import (
	"fmt"
	"testing"
)

type testRefDetector struct {
	result string
}

func (d *testRefDetector) detectRef(dir string) string {
	return d.result
}

func TestTargetRef(t *testing.T) {
	testCases := []struct {
		explicitRef      string
		detectRefEnabled bool
		detectRefResult  string
		expectedResult   string
	}{
		{
			explicitRef:      "",
			detectRefEnabled: true,
			detectRefResult:  "detected-ref",
			expectedResult:   "detected-ref",
		},
		{
			explicitRef:      "",
			detectRefEnabled: true,
			detectRefResult:  "",
			expectedResult:   defaultRef,
		},
		{
			explicitRef:      "explicit-ref",
			detectRefEnabled: true,
			detectRefResult:  "detected-ref",
			expectedResult:   "explicit-ref",
		},
		{
			explicitRef:      "explicit-ref",
			detectRefEnabled: true,
			detectRefResult:  "",
			expectedResult:   "explicit-ref",
		},
		{
			explicitRef:      "explicit-ref",
			detectRefEnabled: false,
			detectRefResult:  "foo",
			expectedResult:   "explicit-ref",
		},
		{
			explicitRef:      "",
			detectRefEnabled: false,
			detectRefResult:  "",
			expectedResult:   defaultRef,
		},
		{
			explicitRef:      "explicit-ref",
			detectRefEnabled: false,
			detectRefResult:  "",
			expectedResult:   "explicit-ref",
		},
		{
			explicitRef:      "explicit-ref",
			detectRefEnabled: false,
			detectRefResult:  "detected-ref",
			expectedResult:   "explicit-ref",
		},
	}
	for idx, testCase := range testCases {
		t.Run(fmt.Sprintf("test case %d", idx), func(t *testing.T) {
			driver := &GitDriver{
				Ref:           testCase.explicitRef,
				DetectRef:     testCase.detectRefEnabled,
				refDetetector: &testRefDetector{result: testCase.detectRefResult},
			}
			actualResult := driver.targetRef("dir")
			if actualResult != testCase.expectedResult {
				t.Errorf("expected target ref: %q, got: %q", testCase.expectedResult, actualResult)
			}
		})
	}
}

// Tests that the git driver is able to parse its config.
func TestGitConfig(t *testing.T) {
	cfg := `{"username" : "git_user", "password" : "git_pass"}`

	d, err := New("git", []byte(cfg))
	if err != nil {
		t.Fatal(err)
	}

	git := d.Driver.(*GitDriver)
	if git.Username != "git_user" {
		t.Fatalf("expected username of \"git_user\", got %s", git.Username)
	}

	if git.Password != "git_pass" {
		t.Fatalf("expected password of \"git_pass\", got %s", git.Password)
	}
}
