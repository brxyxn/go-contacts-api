package logger

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_parseMessage(t *testing.T) {
	type args struct {
		message []interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "default",
			args: args{
				message: []interface{}{"some", "message"},
			},
			want: "some message"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMessage(tt.args.message...); got != tt.want {
				t.Errorf("parseMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_format(t *testing.T) {
	value := interface{}("[value]")
	f1 := format(value)
	got := f1("")
	f2 := func(i interface{}) string {
		var str string
		if s, ok := value.(string); ok {
			str = s
		}
		return str + " " + fmt.Sprint(i)
	}
	want := f2("")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("unexpected output\ngot:  %v\nwant: %v", got, want)
	}
}

func Test_getLogLevel(t *testing.T) {
	type args struct {
		loglevel string
	}
	tests := []struct {
		name string
		args args
		want zerolog.Level
	}{
		{
			name: "trace",
			args: args{
				loglevel: zerolog.LevelTraceValue,
			},
			want: zerolog.TraceLevel,
		},
		{
			name: "debug",
			args: args{
				loglevel: zerolog.LevelDebugValue,
			},
			want: zerolog.DebugLevel,
		},
		{
			name: "info",
			args: args{
				loglevel: zerolog.LevelInfoValue,
			},
			want: zerolog.InfoLevel,
		},
		{
			name: "warn",
			args: args{
				loglevel: zerolog.LevelWarnValue,
			},
			want: zerolog.WarnLevel,
		},
		{
			name: "error",
			args: args{
				loglevel: zerolog.LevelErrorValue,
			},
			want: zerolog.ErrorLevel,
		},
		{
			name: "fatal",
			args: args{
				loglevel: zerolog.LevelFatalValue,
			},
			want: zerolog.FatalLevel,
		},
		{
			name: "panic",
			args: args{
				loglevel: zerolog.LevelPanicValue,
			},
			want: zerolog.PanicLevel,
		},
		{
			name: "disabled",
			args: args{
				loglevel: "disabled",
			},
			want: zerolog.Disabled,
		},
		{
			name: "nolevel",
			args: args{
				loglevel: "nolevel",
			},
			want: zerolog.NoLevel,
		},
		{
			name: "default",
			args: args{
				loglevel: "",
			},
			want: zerolog.InfoLevel,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLogLevel(tt.args.loglevel); got != tt.want {
				t.Errorf("getLogLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLogger(t *testing.T) {
	c1 := LogConfig{
		os.Stdout,
		LogOptions{
			true,
			zerolog.LevelDebugValue,
			CallerOptions{
				WithCaller:       true,
				WithCustomCaller: true,
				BasePath:         "common",
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

	c2 := LogConfig{
		os.Stdout,
		LogOptions{
			false,
			zerolog.LevelDebugValue,
			CallerOptions{
				WithCaller:       false,
				WithCustomCaller: false,
				BasePath:         "",
			},
			ConsoleOptions{
				IsConsole:     true,
				WithoutColor:  false,
				FormatMessage: "",
			},
			DateTimeOptions{
				WithTime:   true,
				TimeFormat: time.RFC3339,
			},
		},
	}

	type args struct {
		cfg LogConfig
	}
	tests := []struct {
		name string
		args args
		want SimpleLogger
	}{
		{
			name: "IsDebug",
			args: args{
				cfg: c1,
			},
			want: &Logger{logger: buildLogger(c1)},
		},
		{
			name: "IsConsole",
			args: args{
				cfg: c2,
			},
			want: &Logger{logger: buildLogger(c2)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewLogger(tt.args.cfg)
			tt.want = got
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLogger() = got %v, want %v", got, tt.want)
			}
		})
	}
}
