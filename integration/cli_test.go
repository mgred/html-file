package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var update = flag.Bool("update", false, "update golden files")
var binaryName = "html-file-coverage"
var binaryPath = ""

func fixturePath(t *testing.T, fixture string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("problems recovering caller information")
	}

	return filepath.Join(filepath.Dir(filename), fixture)
}

func writeFixture(t *testing.T, fixture string, content []byte) {
	err := os.WriteFile(fixturePath(t, fixture), content, 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func loadFixture(t *testing.T, fixture string) string {
	content, err := os.ReadFile(fixturePath(t, fixture))
	if err != nil {
		t.Fatal(err)
	}

	return string(content)
}

func TestCliArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"js files", []string{
			"foo.js", "bar.js", "baz.js",
		}, "js-files.golden"},
		{"js css files", []string{
			"foo.js", "bar.js", "baz.js", "foo.css", "bar.css", "baz.css",
		}, "js-css-files.golden"},
		{"mixed options", []string{
			"--scripts", "--async", "a.js", "b.js", "c.js",
			"--styles", "--preload", "a.css", "b.css",
			"--scripts", "--async", "--module", "d.js", "e.js", "f.js",
			"--favicon", "favicon.ico",
		}, "mixed-options.golden"},
		{"mixed options short", []string{
			"-s", "-a", "a.js", "b.js", "c.js",
			"-S", "-p", "a.css", "b.css",
			"-s", "-a", "-m", "d.js", "e.js", "f.js",
			"-f", "favicon.ico",
		}, "mixed-options-short.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			output, err := runBinary(tt.args)

			if err != nil {
				t.Fatal(err)
			}

			if *update {
				writeFixture(t, tt.fixture, output)
			}

			actual := string(output)

			expected := loadFixture(t, tt.fixture)

			if !reflect.DeepEqual(actual, expected) {
				t.Fatalf("actual = %s, expected = %s", actual, expected)
			}
		})
	}
}

func TestMain(m *testing.M) {
	err := os.Chdir("..")
	if err != nil {
		fmt.Printf("could not change dir: %v", err)
		os.Exit(1)
	}

	dir, err := os.Getwd()
	if err != nil {
		fmt.Printf("could not get current dir: %v", err)
	}

	binaryPath = filepath.Join(dir, binaryName)

	os.Exit(m.Run())
}

func runBinary(args []string) ([]byte, error) {
	cmd := exec.Command(binaryPath, args...)
	env := []string{"GOCOVERDIR=.coverage", "HTML_FILE_HASH=25032019"}
	cmd.Env = append(os.Environ(), env...)
	return cmd.CombinedOutput()
}
