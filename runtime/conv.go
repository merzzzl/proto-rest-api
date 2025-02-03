package runtime

import (
	"fmt"
	"strconv"
)

func ParseFloat32(s string) (float32, error) {
	if s == "" {
		return 0, nil
	}

	f64, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return 0, fmt.Errorf("parse float32: %w", err)
	}

	return float32(f64), nil
}

func ParseFloat64(s string) (float64, error) {
	if s == "" {
		return 0, nil
	}

	f64, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, fmt.Errorf("parse float64: %w", err)
	}

	return f64, nil
}

func ParseBool(s string) (bool, error) {
	if s == "" {
		return false, nil
	}

	b, err := strconv.ParseBool(s)
	if err != nil {
		return false, fmt.Errorf("parse bool: %w", err)
	}

	return b, nil
}

func ParseInt32(s string) (int32, error) {
	if s == "" {
		return 0, nil
	}

	s64, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse int32: %w", err)
	}

	return int32(s64), nil
}

func ParseInt64(s string) (int64, error) {
	if s == "" {
		return 0, nil
	}

	s64, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("parse int64: %w", err)
	}

	return s64, nil
}

func ParseUint32(s string) (uint32, error) {
	if s == "" {
		return 0, nil
	}

	s64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse uint32: %w", err)
	}

	return uint32(s64), nil
}

func ParseUint64(s string) (uint64, error) {
	if s == "" {
		return 0, nil
	}

	s64, err := strconv.ParseUint(s, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("parse uint64: %w", err)
	}

	return s64, nil
}
