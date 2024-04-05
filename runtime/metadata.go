package runtime

import (
	"errors"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

var (
	ErrEmptyHeaderKey        = errors.New("empty header key")
	ErrContainsNonPrintables = errors.New("value contains non-printable characters")
	ErrContainsIllegal       = errors.New("header key contains illegal characters")
)

func ValidateMD(md metadata.MD) error {
	for k, vals := range md {
		if err := validatePair(k, vals...); err != nil {
			return err
		}
	}

	return nil
}

func hasNotPrintable(msg string) bool {
	for i := 0; i < len(msg); i++ {
		if msg[i] < 0x20 || msg[i] > 0x7E {
			return true
		}
	}

	return false
}

func validatePair(key string, vals ...string) error {
	if key == "" {
		return ErrEmptyHeaderKey
	}

	if key[0] == ':' {
		return nil
	}

	for i := 0; i < len(key); i++ {
		r := key[i]
		if !(r >= 'a' && r <= 'z') && !(r >= '0' && r <= '9') && r != '.' && r != '-' && r != '_' {
			return fmt.Errorf("%w in %q", ErrContainsIllegal, key)
		}
	}

	if strings.HasSuffix(key, "-bin") {
		return nil
	}

	for _, val := range vals {
		if hasNotPrintable(val) {
			return fmt.Errorf("%w in %q", ErrContainsNonPrintables, key)
		}
	}

	return nil
}
