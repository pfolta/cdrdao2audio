package msf

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

const (
	MaxMinutes = 99
	MaxSeconds = 59
	MaxFrames  = 74

	FramesPerSecond = 75
	MaxTotalFrames  = MaxMinutes*60*FramesPerSecond + MaxSeconds*FramesPerSecond + MaxFrames

	SectorBytes = 2352
)

var msfRegex = regexp.MustCompile(`^(\d{2}):(0[0-9]|[1-5][0-9]):(0[0-9]|[1-6][0-9]|7[0-4])$`)

// ErrorInvalidMSF indicates an invalid [MSF] value.
var ErrInvalidMSF = errors.New("invalid MSF")

// MSF represents an audio CD timestamp in minutes:seconds:frames (MM:SS:FF).
type MSF struct {
	totalFrames uint32
}

// New returns a [MSF] for the specified total number of frames.
func New(frames uint32) (MSF, error) {
	if frames > MaxTotalFrames {
		err := fmt.Errorf("%w: %d > %d", ErrInvalidMSF, frames, MaxTotalFrames)
		return MSF{}, err
	}

	return MSF{frames}, nil
}

// Must is like [New] but panics if frames is invalid.
// It is intended for tests and trusted inputs.
func Must(frames uint32) MSF {
	msf, err := New(frames)
	if err != nil {
		panic(err)
	}

	return msf
}

// Parse parses a MM:SS:FF timestamp and returns the corresponding [MSF].
func Parse(str string) (MSF, error) {
	// Handle `00:00:00` formatted as `0`
	if str == "0" {
		return MSF{0}, nil
	}

	msfParts := msfRegex.FindStringSubmatch(str)
	if msfParts == nil {
		err := fmt.Errorf("%w: %s", ErrInvalidMSF, str)
		return MSF{}, err
	}

	m, _ := strconv.ParseUint(msfParts[1], 10, 32)
	s, _ := strconv.ParseUint(msfParts[2], 10, 32)
	f, _ := strconv.ParseUint(msfParts[3], 10, 32)

	return New(uint32(m*60*FramesPerSecond + s*FramesPerSecond + f))
}

// MustParse is like [Parse] but panics if the string cannot be parsed.
// It is intended for tests and trusted inputs.
func MustParse(str string) MSF {
	msf, err := Parse(str)
	if err != nil {
		panic(err)
	}

	return msf
}

// TotalFrames returns the total number of frames of this [MSF].
func (msf MSF) TotalFrames() uint32 {
	return msf.totalFrames
}

// SectorBytes returns the total size of this [MSF] in bytes.
func (msf MSF) SectorBytes() uint32 {
	return msf.totalFrames * SectorBytes
}

// String returns a MM:SS:FF timestamp representation of this [MSF].
func (msf MSF) String() string {
	m := msf.totalFrames / FramesPerSecond / 60
	s := msf.totalFrames / FramesPerSecond % 60
	f := msf.totalFrames % FramesPerSecond
	return fmt.Sprintf("%02d:%02d:%02d", m, s, f)
}
