package envigo

import (
	"os"
	"testing"
	"time"
)

func TestGetUintFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Default  uint
		Expected uint
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"", 2, 2},
		{"100", 1, 100},
		{"10000", 1, 10000},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if res := GetUintFromEnv(name, test.Default); res != test.Expected {
			t.Errorf(
				"#%d %v: expected %d but %d returned",
				n, test, test.Expected, res)
		}
	}
}

func TestGetPosIntFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Default  uint
		Expected uint
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"1", 2, 1},
		{"100", 1, 100},
		{"10000", 1, 10000},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if res := GetPosIntFromEnv(name, test.Default); res != test.Expected {
			t.Errorf(
				"#%d %v: expected %d but %d returned",
				n, test, test.Expected, res)
		}
	}
}

func TestGetIntFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Default  int
		Expected int
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"", 2, 2},
		{"", -2, -2},
		{"100", 1, 100},
		{"-100", 1, -100},
		{"10000", 1, 10000},
		{"-10000", 1, -10000},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if res := GetIntFromEnv(name, test.Default); res != test.Expected {
			t.Errorf(
				"#%d %v: expected %d but %d returned",
				n, test, test.Expected, res)
		}
	}
}

func TestGetIntMinMaxFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Min      int
		Max      int
		Default  int
		Expected int
	}{
		{"", 1, 5, 0, 0},
		{"", 1, 5, 1, 1},
		{"", 1, 5, 2, 2},
		{"", 1, 5, -2, -2},
		{"-1", -1, 5, 1, -1},
		{"0", -1, 5, 1, 0},
		{"5", 1, 5, 1, 5},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if res := GetIntMinMaxFromEnv(name, test.Min, test.Max, test.Default); res != test.Expected {
			t.Errorf(
				"#%d %v: expected %d but %d returned",
				n, test, test.Expected, res)
		}
	}
}

func TestGetPortNumberFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Default  uint16
		Expected uint16
	}{
		{"", 0, 0},
		{"", 1, 1},
		{"", 2, 2},
		{"1", 2, 1},
		{"2", 1, 2},
		{"10000", 1, 10000},
		{"65535", 1, 0xffff},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if res := GetPortNumberFromEnv(name, test.Default); res != test.Expected {
			t.Errorf(
				"#%d %v: expected %d but %d returned",
				n, test, test.Expected, res)
		}
	}
}

func TestGetBoolFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Default  bool
		Expected bool
	}{
		{"", true, true},
		{"", false, false},
		{"1", true, true},
		{"1", false, true},
		{"t", true, true},
		{"t", false, true},
		{"T", true, true},
		{"T", false, true},
		{"true", true, true},
		{"true", false, true},
		{"TRUE", true, true},
		{"TRUE", false, true},
		{"0", true, false},
		{"0", false, false},
		{"f", true, false},
		{"f", false, false},
		{"F", true, false},
		{"F", false, false},
		{"false", true, false},
		{"false", false, false},
		{"FALSE", true, false},
		{"FALSE", false, false},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if res := GetBoolFromEnv(name, test.Default); res != test.Expected {
			t.Errorf(
				"#%d %v: expected %v but %v returned",
				n, test, test.Expected, res)
		}
	}
}

func TestGetStringIndexFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Valid    []string
		Default  int
		Expected int
	}{
		{"", []string{}, 0, 0},
		{"", []string{}, 1, 1},
		{"a", []string{"a", "b", "c"}, 0, 0},
		{"a", []string{"a", "b", "c"}, 1, 0},
		{"a", []string{"a", "b", "c"}, 2, 0},
		{"b", []string{"a", "b", "c"}, 0, 1},
		{"c", []string{"a", "b", "c"}, 1, 2},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if i := GetStringIndexFromEnv(name, test.Valid, test.Default); i != test.Expected {
			t.Errorf(
				"#%d %v: expected %d but %d returned",
				n, test, test.Expected, i)
		}
	}
}

func TestGetValidStringFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Valid    []string
		Default  string
		Expected string
	}{
		{"", []string{}, "", ""},
		{"", []string{}, "b", "b"},
		{"a", []string{"a", "b", "c"}, "", "a"},
		{"a", []string{"a", "b", "c"}, "x", "a"},
		{"a", []string{"a", "b", "c"}, "y", "a"},
		{"b", []string{"a", "b", "c"}, "", "b"},
		{"c", []string{"a", "b", "c"}, "x", "c"},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if i := GetValidStringFromEnv(name, test.Valid, test.Default); i != test.Expected {
			t.Errorf(
				"#%d %v: expected `%s` but `%s` returned",
				n, test, test.Expected, i)
		}
	}
}

func TestGetDurationFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Default  time.Duration
		Expected time.Duration
	}{
		{"", time.Second, time.Second},
		{"", time.Minute, time.Minute},
		{"-1s", 2, -time.Second},
		{"1s", 3, time.Second},
		{"1.5s", 4, time.Second + time.Millisecond*500},
		{"2h", 5, time.Hour * 2},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if res := GetDurationFromEnv(name, test.Default); res != test.Expected {
			t.Errorf(
				"#%d %v: expected %s but %s returned",
				n, test, test.Expected.String(), res.String())
		}
	}
}

func TestGetUDurationFromEnv(t *testing.T) {
	testset := []struct {
		EnvValue string
		Default  time.Duration
		Expected time.Duration
	}{
		{"", time.Second, time.Second},
		{"", time.Minute, time.Minute},
		{"1s", 3, time.Second},
		{"1.5s", 4, time.Second + time.Millisecond*500},
		{"2h", 5, time.Hour * 2},
	}
	name := "ENVX_TEST_VAR"
	for n, test := range testset {
		err := os.Setenv(name, test.EnvValue)
		if err != nil {
			panic(err)
		}
		if res := GetUDurationFromEnv(name, test.Default); res != test.Expected {
			t.Errorf(
				"#%d %v: expected %s but %s returned",
				n, test, test.Expected.String(), res.String())
		}
	}
}
