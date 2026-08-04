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
	_ "embed"
	"fmt"
	"runtime"
	"strings"
)

const programName = "cdrdao2audio"

// Build information injected via -ldflags.
var (
	buildDate = "unknown"
	version   = "dev"
)

//go:embed LICENSE
var license string

// AppInfo contains build and runtime information about the application.
type AppInfo struct {
	// Name is the application name.
	Name string `json:"name" yaml:"name"`

	// Version is the application version injected at build time.
	Version string `json:"version" yaml:"version"`

	// BuildDate is the timestamp when the binary was built.
	BuildDate string `json:"buildDate" yaml:"buildDate"`

	// License contains the application license text.
	License string `json:"license" yaml:"license"`

	// OS is the target operating system.
	OS string `json:"os" yaml:"os"`

	// Arch is the target architecture.
	Arch string `json:"arch" yaml:"arch"`
}

// ShortVersion returns the raw application's version number.
func (appInfo AppInfo) ShortVersion() ShortVersion {
	return ShortVersion{Version: strings.TrimPrefix(appInfo.Version, "v")}
}

// String formats the application's build and runtime information.
func (appInfo AppInfo) String() string {
	return fmt.Sprintf("%s version %s-%s-%s (%s)\n\n%s",
		appInfo.Name,
		appInfo.Version,
		appInfo.OS,
		appInfo.Arch,
		appInfo.BuildDate,
		appInfo.License,
	)
}

// ShortVersion contains the raw application's version number.
type ShortVersion struct {
	Version string `json:"version" yaml:"version"`
}

// String formats the application's raw version number.
func (shortVersion ShortVersion) String() string {
	return shortVersion.Version
}

// GetAppInfo returns the application's build and runtime information.
func GetAppInfo() AppInfo {
	return AppInfo{
		Name:      programName,
		Version:   version,
		BuildDate: buildDate,
		License:   strings.TrimSpace(license),
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}
