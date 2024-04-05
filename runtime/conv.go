package runtime

import (
	"fmt"
	"strconv"
)

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
