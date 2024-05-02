package logger

import (
	"errors"
	"github.com/rs/zerolog"
	"io"
	"os"
	"os/exec"
	"reflect"
	"testing"
	"time"
)

func ExampleDebug() {
	// {"level":"debug","message":"debug message"}
	Debug("debug", "message")
	// Output:
}

func ExampleInfo() {
	// {"level":"debug","message":"debug message"}
	Info("info", "message")
	// Output:
}

func ExampleWarn() {
	// {"level":"debug","message":"debug message"}
	Warn("warn", "message")
	// Output:
}

func ExampleError() {
	// {"level":"debug","message":"debug message"}
	err := errors.New("message")
	Error("err:", err)
	// Output:
}

func ExampleFatal() {
	// {"level":"debug","message":"debug message"}
	Fatal("debug", "message")
	// Outputs:
}

func ExamplePanic() {
	// {"level":"debug","message":"debug message"}
	Panic("debug", "message")
	// Outputs:
}

func Test_PanicLogger(t *testing.T) {
	defer func() {
		r := recover()
		if r != "panic" {
			t.Errorf("the code did not panic")
		}
	}()

	Panic("panic")
}

func Test_FatalLogger(t *testing.T) {
	var err error

	if os.Getenv("BE_CRASHER") == "1" {
		zerolog.TimestampFunc = func() time.Time {
			return time.Date(2008, 1, 8, 17, 5, 5, 0, time.Local)
		}
		Fatal("Cannot start")
		return
	}

	// Start the actual test in a different subprocess
	cmd := exec.Command(os.Args[0], "-test.run=Test_FatalLogger")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	stdout, _ := cmd.StdoutPipe()
	if err = cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Check that the log fatal message is what we want
	gotBytes, _ := io.ReadAll(stdout)
	got := string(gotBytes)

	want := `{"level":"fatal","msg":"Cannot start"}` + "\n"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
	}

	// Check that the program exited
	err = cmd.Wait()
	if e, ok := err.(*exec.ExitError); !ok || e.Success() {
		t.Fatalf("Process ran with err %v, want exit status 1", err)
	}
}
