// Copyright (c) 2026 Peter Folta
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cdrdao2audio

import (
	"runtime"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetAppInfo(t *testing.T) {
	originalVersion := version
	originalBuildDate := buildDate

	defer func() {
		version = originalVersion
		buildDate = originalBuildDate
	}()

	// Simulate build-time metadata injection.
	version = "v1.4.2-test"
	buildDate = "2026-07-15T22:29:10Z"

	appInfo := GetAppInfo()

	assert.Equal(t, appInfo.Name, programName)
	assert.Equal(t, appInfo.Version, "v1.4.2-test")
	assert.Equal(t, appInfo.BuildDate, "2026-07-15T22:29:10Z")
	assert.Assert(t, is.Contains(appInfo.License, "Copyright"))
	assert.Equal(t, appInfo.OS, runtime.GOOS)
	assert.Equal(t, appInfo.Arch, runtime.GOARCH)
}

func TestAppInfoString(t *testing.T) {
	appInfo := AppInfo{
		Name:      "cdrdao2audio",
		Version:   "v1.4.2-test",
		BuildDate: "2026-07-15T22:29:10Z",
		License:   "MIT",
		OS:        "darwin",
		Arch:      "arm64",
	}

	assert.Equal(t, appInfo.String(), "cdrdao2audio version v1.4.2-test-darwin-arm64 (2026-07-15T22:29:10Z)\n\nMIT")
}

func TestAppInfoShortVersion(t *testing.T) {
	appInfo := AppInfo{Version: "v1.4.2-test"}
	assert.Equal(t, appInfo.ShortVersion(), ShortVersion{Version: "1.4.2-test"})
}

func TestShortVersionString(t *testing.T) {
	version := ShortVersion{Version: "1.4.2-test"}
	assert.Equal(t, version.String(), "1.4.2-test")
}
