package envigo

import (
	"log"
	"os"
	"strconv"
	"time"
)

// Parse non negative integer from environment variable.
// Report error and crash when environment variable is in bad format.
func GetUintFromEnv(name string, def_value uint) uint {
	if s := os.Getenv(name); 0 < len(s) {
		i, err := strconv.Atoi(s)
		if err != nil || i < 0 {
			log.Fatalf("Bad value for %s: `%s`", name, s)
		}
		return uint(i)
	}
	return def_value
}

// Parse non uint64 integer from environment variable.
// Report error and crash when environment variable is in bad format.
func GetUint64FromEnv(name string, def_value uint64) uint64 {
	if s := os.Getenv(name); 0 < len(s) {
		i, err := strconv.ParseUint(s, 10, 64)
		if err != nil {
			log.Fatalf("Bad value for %s: `%s`", name, s)
		}
		return i
	}
	return def_value
}

// Parse positive integer from environment variable.
// Report error and crash when environment variable is in bad format.
func GetPosIntFromEnv(name string, def_value uint) uint {
	if s := os.Getenv(name); 0 < len(s) {
		i, err := strconv.Atoi(s)
		if err != nil || i < 1 {
			log.Fatalf("Bad value for %s: `%s`", name, s)
		}
		return uint(i)
	}
	return def_value
}

// Parse integer from environment variable.
// Report error and crash when environment variable is in bad format.
func GetIntFromEnv(name string, def_value int) int {
	if s := os.Getenv(name); 0 < len(s) {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Bad value for %s: `%s`", name, s)
		}
		return i
	}
	return def_value
}

// Parse integer from environment variable.
// Report error and crash when environment variable is in bad format.
func GetIntMinMaxFromEnv(name string, min, max, def int) int {
	if s := os.Getenv(name); 0 < len(s) {
		i, err := strconv.Atoi(s)
		if err != nil || i < min || max < i {
			log.Fatalf(
				"Bad value for %s: `%s`. Valid values are %d..%d",
				name, s, min, max)
		}
		return i
	}
	return def
}

// Parse TCP/UDP port number from environment variable.
// Report error and crash when environment variable is in bad format.
func GetPortNumberFromEnv(name string, def_value uint16) uint16 {
	if s := os.Getenv(name); 0 < len(s) {
		i, err := strconv.Atoi(s)
		if err != nil || i < 0 {
			log.Fatalf("Bad value for %s: `%s`", name, s)
		}
		if i < 1 || 0xffff < i {
			log.Fatalf("Bad value for %s: `%s`", name, s)
		}
		return uint16(i)
	}
	return def_value
}

// Parse boolean from environment variable.
// Report error and crash when environment variable is in bad format.
func GetBoolFromEnv(name string, def_value bool) bool {
	if s := os.Getenv(name); 0 < len(s) {
		b, err := strconv.ParseBool(s)
		if err != nil {
			log.Fatalf("Bad value for %s: `%s`", name, s)
		}
		return b
	}
	return def_value
}

// Return index of environment variable value within valid_values array.
func GetStringIndexFromEnv(name string, valid_values []string, def_value int) int {
	s := os.Getenv(name)
	if 0 < len(s) {
		for i, v := range valid_values {
			if v == s {
				return i
			}
		}
		log.Fatalf(
			"Bad value for %s: `%s`. Valid values are: %v",
			name, s, valid_values)
	}
	return def_value
}

// Return string from environment variable.
func GetValidStringFromEnv(name string, validValues []string, def string) string {
	if s := os.Getenv(name); 0 < len(s) {
		for _, v := range validValues {
			if v == s {
				return v
			}
		}
		log.Fatalf(
			"Bad value for %s: %s. Valid values are: %v",
			name, s, validValues)
	}
	return def
}

// Parse duration from environment variable.
// Report error and crash when environment variable is in bad format.
func GetDurationFromEnv(name string, def_value time.Duration) time.Duration {
	if s := os.Getenv(name); 0 < len(s) {
		duration, err := time.ParseDuration(s)
		if err != nil {
			log.Fatalf("Bad duration value for %s: `%s`", name, s)
		}
		return duration
	}
	return def_value
}

// Parse positive duration from environment variable.
// Report error and crash when environment variable is in bad format.
func GetUDurationFromEnv(name string, def_value time.Duration) time.Duration {
	if s := os.Getenv(name); 0 < len(s) {
		duration, err := time.ParseDuration(s)
		if err != nil || duration < 0 {
			log.Fatalf("Bad duration value for %s: `%s`", name, s)
		}
		return duration
	}
	return def_value
}

// Get string value from environment variable.
func GetStrFromEnv(name string, def string) string {
	if s := os.Getenv(name); s != "" {
		return s
	}
	return def
}
