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
	"runtime/debug"
	"testing"

	"gotest.tools/v3/assert"
	is "gotest.tools/v3/assert/cmp"
)

func TestGetAppInfo(t *testing.T) {
	simulateInject(t, &version, "v1.4.2-test")
	simulateInject(t, &buildDate, "2026-07-15T22:29:10Z")

	appInfo := GetAppInfo()

	assert.Assert(t, is.Equal(appInfo.Name, programName))
	assert.Assert(t, is.Equal(appInfo.Version, "v1.4.2-test"))
	assert.Assert(t, is.Equal(appInfo.BuildDate, "2026-07-15T22:29:10Z"))
	assert.Assert(t, is.Contains(appInfo.License, "Copyright"))
	assert.Assert(t, is.Equal(appInfo.OS, runtime.GOOS))
	assert.Assert(t, is.Equal(appInfo.Arch, runtime.GOARCH))
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

	assert.Assert(t, is.Equal(appInfo.String(), "cdrdao2audio version v1.4.2-test-darwin-arm64 (2026-07-15T22:29:10Z)\n\nMIT"))
}

func TestAppInfoShortVersion(t *testing.T) {
	appInfo := AppInfo{Version: "v1.4.2-test"}
	assert.Assert(t, is.Equal(appInfo.ShortVersion(), ShortVersion{Version: "1.4.2-test"}))
}

func TestShortVersionString(t *testing.T) {
	version := ShortVersion{Version: "1.4.2-test"}
	assert.Assert(t, is.Equal(version.String(), "1.4.2-test"))
}

func TestDetermineVersion(t *testing.T) {
	tests := []struct {
		name        string
		version     string
		buildInfo   *debug.BuildInfo
		buildInfoOK bool
		want        string
	}{
		{
			name:    "injected version",
			version: "v1.4.2-test",
			want:    "v1.4.2-test",
		},
		{
			name:    "embedded Go build information",
			version: "",
			buildInfo: &debug.BuildInfo{
				Main: debug.Module{
					Version: "v1.4.2",
				},
			},
			buildInfoOK: true,
			want:        "v1.4.2",
		},
		{
			name:    "development build",
			version: "",
			buildInfo: &debug.BuildInfo{
				Main: debug.Module{
					Version: "(devel)",
				},
			},
			buildInfoOK: true,
			want:        "dev",
		},
		{
			name:        "fallback",
			version:     "",
			buildInfoOK: false,
			want:        "dev",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			simulateInject(t, &version, test.version)
			simulateInject(t, &readBuildInfo, func() (*debug.BuildInfo, bool) {
				return test.buildInfo, test.buildInfoOK
			})

			assert.Assert(t, is.Equal(determineVersion(), test.want))
		})
	}
}

func simulateInject[T any](t *testing.T, variable *T, value T) {
	t.Helper()

	original := *variable
	*variable = value

	t.Cleanup(func() {
		*variable = original
	})
}
