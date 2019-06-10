package multiflag

import (
	"flag"
	"time"
)

// Bool defines a bool flag with specified default value, usage string and, names.
// The return value is the address of a bool variable that stores the value of the flag.
func Bool(value bool, usage string, names ...string) *bool {
	var v bool
	for _, name := range names {
		flag.BoolVar(&v, name, value, usage)
	}
	return &v
}

// Duration defines a time.Duration flag with specified default value, usage string, and names.
// The return value is the address of a time.Duration variable that stores the value of the flag.
// The flag accepts a value acceptable to time.ParseDuration.
func Duration(value time.Duration, usage string, names ...string) *time.Duration {
	var v time.Duration
	for _, name := range names {
		flag.DurationVar(&v, name, value, usage)
	}
	return &v
}

// Float64 defines a float64 flag with specified default value, usage string, and names.
// The return value is the address of a float64 variable that stores the value of the flag.
func Float64(value float64, usage string, names ...string) *float64 {
	var v float64
	for _, name := range names {
		flag.Float64Var(&v, name, value, usage)
	}
	return &v
}

// Int defines an int flag with specified default value, usage string, and names.
// The return value is the address of an int variable that stores the value of the flag.
func Int(value int, usage string, names ...string) *int {
	var v int
	for _, name := range names {
		flag.IntVar(&v, name, value, usage)
	}
	return &v
}

// Int64 defines an int flag with specified default value, usage string, and names.
// The return value is the address of an int64 variable that stores the value of the flag.
func Int64(value int64, usage string, names ...string) *int64 {
	var v int64
	for _, name := range names {
		flag.Int64Var(&v, name, value, usage)
	}
	return &v
}

// String defines an int flag with specified default value, usage string, and names.
// The return value is the address of an string variable that stores the value of the flag.
func String(value string, usage string, names ...string) *string {
	var v string
	for _, name := range names {
		flag.StringVar(&v, name, value, usage)
	}
	return &v
}

// Uint defines an int flag with specified default value, usage string, and names.
// The return value is the address of an uint variable that stores the value of the flag.
func Uint(value uint, usage string, names ...string) *uint {
	var v uint
	for _, name := range names {
		flag.UintVar(&v, name, value, usage)
	}
	return &v
}

// Uint64 defines an int flag with specified default value, usage string, and names.
// The return value is the address of an uint64 variable that stores the value of the flag.
func Uint64(value uint64, usage string, names ...string) *uint64 {
	var v uint64
	for _, name := range names {
		flag.Uint64Var(&v, name, value, usage)
	}
	return &v
}
