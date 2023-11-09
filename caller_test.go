package runtimeutils_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	runtimeutils "gitea.linesip.com/libraries/runtimeutils"
)

const (
	callerSkip  = 2
	maxFrames   = 10
	packageName = "runtimeutils_test"
	fileName    = "caller_test.go"
)

func TestGetCallerFunctionName(t *testing.T) {
	t.Parallel()

	name := runtimeutils.GetCallerFunctionName()

	if name != "TestGetCallerFunctionName" {
		t.Errorf(
			"Expected function name only, but package name also returned: %s",
			name,
		)
	}
}

func TestPackageName(t *testing.T) {
	t.Parallel()

	info, err := runtimeutils.GetCallerInfo(callerSkip)
	if err != nil {
		t.Errorf("Expected no error, but %+v found", err)
	}

	infoPackageName := info.PackageName()
	if infoPackageName != packageName {
		t.Errorf("Expected %s, got %s", packageName, infoPackageName)
	}
}

func TestFileName(t *testing.T) {
	t.Parallel()

	info, err := runtimeutils.GetCallerInfo(callerSkip)
	if err != nil {
		t.Errorf("Expected no error, but %+v found", err)
	}

	infoFileName := info.FileName()
	if infoFileName != fileName {
		t.Errorf("Expected filename of %s got %s", fileName, infoFileName)
	}
}

func TestString(t *testing.T) {
	t.Parallel()

	info, err := runtimeutils.GetCallerInfo(callerSkip)
	if err != nil {
		t.Errorf("Expected no error, but %+v found", err)
	}

	infoString := info.String()
	expectedString := fmt.Sprintf("%s.TestString", packageName)

	if infoString != expectedString {
		t.Errorf("expected %s but have %s", expectedString, infoString)
	}
}

func TestDebugInfo(t *testing.T) {
	t.Parallel()

	info, err := runtimeutils.GetCallerInfo(callerSkip)
	if err != nil {
		t.Errorf("Expected no error, but %+v found", err)
	}

	infoDebugInfo := info.DebugInfo()
	// The reason for regex instead of structured string is that line number can
	// be changed when writing additional tests at the file, or adding more code
	// for existed tests, so the line number is dynamic.
	expectedDebugInfo := fmt.Sprintf(
		"%s:[0-9]+->%s.TestDebugInfo", fileName, packageName,
	)

	regex, err := regexp.Compile(expectedDebugInfo)
	if err != nil {
		t.Error(err)
	}

	if !regex.MatchString(infoDebugInfo) {
		t.Errorf(
			"Expected FileName():LineNo->PackageName().Func, got %s",
			infoDebugInfo,
		)
	}
}

func TestGestCallerStack(t *testing.T) {
	t.Parallel()

	const callerSkip = 0

	stackInfoList, err := runtimeutils.GetCallerStack(callerSkip, maxFrames)
	if err != nil {
		t.Errorf("Unable to get caller stack: %s", err)
	}

	if stackInfoList == nil {
		t.Error("stackInfoList is nil, expected at least one item")
	}

	itemLength := len(stackInfoList)
	if itemLength == 0 {
		t.Error("stackInfoList is empty slice")
	}
}

func TestGestCallerStackWithError(t *testing.T) {
	t.Parallel()

	const maxFrames = 0

	_, err := runtimeutils.GetCallerStack(callerSkip, maxFrames)
	if err == nil {
		t.Error("expected error, but nil provided")
	}

	if !errors.Is(err, runtimeutils.ErrMaxStackIsLowerThan1) {
		t.Errorf(
			"expected err \"%s\" but \"%s\" provided",
			runtimeutils.ErrMaxStackIsLowerThan1, err,
		)
	}
}
