# cdrdao2audio

**💿 Convert a cdrdao dump to individual audio tracks**

[![GitHub Release](https://img.shields.io/github/v/release/pfolta/cdrdao2audio?include_prereleases)](https://github.com/pfolta/cdrdao2audio/releases/latest)
[![Docker Image Version (tag)](https://img.shields.io/docker/v/pfolta/cdrdao2audio/latest?logo=docker&label=Docker%20image)](https://hub.docker.com/r/pfolta/cdrdao2audio)
[![Go version](https://img.shields.io/github/go-mod/go-version/pfolta/cdrdao2audio/master?logo=go&label=Go)](go.mod)
[![License](https://img.shields.io/github/license/pfolta/cdrdao2audio)](LICENSE)
[![Build status](https://img.shields.io/github/actions/workflow/status/pfolta/cdrdao2audio/ci.yaml?branch=master&logo=github)](https://github.com/pfolta/cdrdao2audio/actions/workflows/ci.yaml)
[![Test coverage](https://img.shields.io/codecov/c/github/pfolta/cdrdao2audio/master?logo=codecov)](https://codecov.io/gh/pfolta/cdrdao2audio)

cdrdao2audio is a command-line application for ripping audio tracks from a [cdrdao](https://cdrdao.sourceforge.net) backup of an audio CD.

## Installation

- Binaries are available for **Linux**, **macOS** and **Windows** on the [Releases](https://github.com/pfolta/cdrdao2audio/releases) page.

  **macOS:** Run the following command to remove the downloaded file from quarantine:

      xattr -d com.apple.quarantine cdrdao2audio-*-darwin-*

- Or, use **Docker** to run cdrdao2audio directly:

      docker run -ti --rm pfolta/cdrdao2audio

- Or, install it with **Go**:

      go install github.com/pfolta/cdrdao2audio/cmd/cdrdao2audio@latest

- Or, build and install it **from source** (requires [Go](https://go.dev)):

      git clone git@github.com:pfolta/cdrdao2audio.git && \
        cd cdrdao2audio && \
        make install

## License

cdrdao2audio is released under the MIT License. See [LICENSE](LICENSE).
