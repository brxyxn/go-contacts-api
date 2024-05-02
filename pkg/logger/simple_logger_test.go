package logger

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"io"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"testing"
	"time"
)

var sl SimpleLogger

func init() {
	cfg := LogConfig{
		os.Stdout,
		LogOptions{
			false,
			zerolog.LevelTraceValue,
			CallerOptions{
				WithCaller:       true,
				WithCustomCaller: true,
				BasePath:         "go-phonebook-api",
			},
			ConsoleOptions{
				IsConsole:     false,
				WithoutColor:  false,
				FormatMessage: "",
			},
			DateTimeOptions{
				WithTime:   true,
				TimeFormat: time.RFC3339,
			},
		},
	}

	sl = NewLogger(cfg)
}

func TestLogger_Panic(t *testing.T) {
	// {"level":"panic","time":"2023-03-21T20:30:38-06:00","caller":"pkg/logger/simple_logger_test.go:43","msg":"panic"}
	defer func() {
		if r := recover(); r != "panic" {
			t.Errorf("The code did not panic")
		}
	}()

	sl.Panic().Msg("panic")
}

func Test_Logger_Fatal(t *testing.T) {
	var now = func() time.Time {
		return time.Date(2008, 1, 8, 17, 5, 5, 0, time.Local)
	}
	var ti = now()
	var ct = ti.Format(time.RFC3339)
	err := errors.New("errorMsg")
	service := "myService"
	msg := fmt.Sprintf("Cannot start %s", service)
	file := "go-phonebook-api/pkg/logger/simple_logger_test.go"
	_, _, line, _ := runtime.Caller(0)

	if os.Getenv("BE_CRASHER") == "1" {
		zerolog.TimestampFunc = func() time.Time {
			return time.Date(2008, 1, 8, 17, 5, 5, 0, time.Local)
		}
		sl.Fatal().
			Err(err).
			Str("service", service).
			Msgf("%s", msg)
		return
	}

	// Start the actual test in a different subprocess
	cmd := exec.Command(os.Args[0], "-test.run=Test_Logger_Fatal")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	stdout, _ := cmd.StdoutPipe()
	if err = cmd.Start(); err != nil {
		t.Fatal(err)
	}

	// Check that the log fatal message is what we want
	gotBytes, _ := io.ReadAll(stdout)
	got := string(gotBytes)

	caller := fmt.Sprintf("%s:%d", file, line+9)
	want := `{"level":"fatal","error":"errorMsg","service":"myService","time":"` + ct + `","caller":"` + caller + `","msg":"` + msg + `"}` + "\n"
	if !reflect.DeepEqual(got, want) {
		t.Errorf("invalid log output:\ngot:  %v\nwant: %v", got, want)
	}

	// Check that the program exited
	err = cmd.Wait()
	if e, ok := err.(*exec.ExitError); !ok || e.Success() {
		t.Fatalf("Process ran with err %v, want exit status 1", err)
	}
}

// Examples

func ExampleLogger_With() {
	/*
		{"level":"error","str":"value","int":0,"int":0,"int":[0,1,2,3],"error":"error value","time":"2023-03-21T20:39:26-06:00","caller":"pkg/logger/simple_logger_test.go:57","msg":"msg"}
	*/
	err := errors.New("error value")
	sl.Error().
		Str("str", "value").
		Int("int", 0).
		Int64("int", 0).
		Ints("int", []int{0, 1, 2, 3}).
		Err(err).
		Msg("msg")
	// Output:
}

func ExampleLogger_Debug() {
	// {"level":"debug","time":"2023-03-21T20:25:36-06:00","caller":"pkg/logger/simple_logger_test.go:52","msg":"debug message"}
	sl.Debug().Msg("debug message")
	// Output:
}

func ExampleLogger_Info() {
	// {"level":"info","time":"2023-03-21T20:25:06-06:00","caller":"pkg/logger/simple_logger_test.go:56","msg":"info message"}
	sl.Info().Msg("info message")
	// Output:
}

func ExampleLogger_Warn() {
	// {"level":"warn","time":"2023-03-21T20:25:57-06:00","caller":"pkg/logger/simple_logger_test.go:63","msg":"warn message"}
	sl.Warn().Msg("warn message")
	// Output:
}

func ExampleLogger_Error() {
	// {"level":"error","time":"2023-03-21T20:26:17-06:00","caller":"pkg/logger/simple_logger_test.go:69","msg":"error message"}
	sl.Error().Msg("error message")
	// Output:
}

func ExampleLogger_Print() {
	// {"level":"debug","time":"2023-03-21T20:26:29-06:00","caller":"pkg/logger/simple_logger.go:33","msg":"print message"}
	sl.Print("print message")
	// Output:
}

func ExampleLogger_Printf() {
	// {"level":"debug","time":"2023-03-21T20:26:42-06:00","caller":"pkg/logger/simple_logger.go:37","msg":"int: 4, string: text"}
	number := 4
	text := "text"
	sl.Printf("int: %v, string: %v", number, text)
	// Output:
}
