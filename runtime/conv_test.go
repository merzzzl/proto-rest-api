package runtime_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	
	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestParseInt32(t *testing.T) {
	t.Run("valid int32", func(t *testing.T) {
		result, err := runtime.ParseInt32("123")
		require.NoError(t, err, "should not return an error for valid input")
		require.Equal(t, int32(123), result, "result should match expected value")
	})

	t.Run("empty string", func(t *testing.T) {
		result, err := runtime.ParseInt32("")
		require.NoError(t, err, "should not return an error for empty input")
		require.Equal(t, int32(0), result, "empty string should return zero")
	})

	t.Run("invalid input", func(t *testing.T) {
		_, err := runtime.ParseInt32("invalid")
		require.Error(t, err, "should return an error for invalid input")
	})
}

func TestParseInt64(t *testing.T) {
	t.Run("valid int64", func(t *testing.T) {
		result, err := runtime.ParseInt64("123456789")
		require.NoError(t, err, "should not return an error for valid input")
		require.Equal(t, int64(123456789), result, "result should match expected value")
	})

	t.Run("empty string", func(t *testing.T) {
		result, err := runtime.ParseInt64("")
		require.NoError(t, err, "should not return an error for empty input")
		require.Equal(t, int64(0), result, "empty string should return zero")
	})

	t.Run("invalid input", func(t *testing.T) {
		_, err := runtime.ParseInt64("invalid")
		require.Error(t, err, "should return an error for invalid input")
	})
}

func TestParseUint32(t *testing.T) {
	t.Run("valid uint32", func(t *testing.T) {
		result, err := runtime.ParseUint32("123")
		require.NoError(t, err, "should not return an error for valid input")
		require.Equal(t, uint32(123), result, "result should match expected value")
	})

	t.Run("empty string", func(t *testing.T) {
		result, err := runtime.ParseUint32("")
		require.NoError(t, err, "should not return an error for empty input")
		require.Equal(t, uint32(0), result, "empty string should return zero")
	})

	t.Run("invalid input", func(t *testing.T) {
		_, err := runtime.ParseUint32("invalid")
		require.Error(t, err, "should return an error for invalid input")
	})
}

func TestParseUint64(t *testing.T) {
	t.Run("valid uint64", func(t *testing.T) {
		result, err := runtime.ParseUint64("123456789")
		require.NoError(t, err, "should not return an error for valid input")
		require.Equal(t, uint64(123456789), result, "result should match expected value")
	})

	t.Run("empty string", func(t *testing.T) {
		result, err := runtime.ParseUint64("")
		require.NoError(t, err, "should not return an error for empty input")
		require.Equal(t, uint64(0), result, "empty string should return zero")
	})

	t.Run("invalid input", func(t *testing.T) {
		_, err := runtime.ParseUint64("invalid")
		require.Error(t, err, "should return an error for invalid input")
	})
}
