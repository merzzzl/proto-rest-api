package runtime_test

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestGetFieldMaskJS_0(t *testing.T) {
	t.Parallel()

	reader := strings.NewReader(`{
        "quiz": {
            "sport": {{
                "q1": {
                    "question": "Which one is correct team name in NBA?",
                    "options": [[
                        "New York Bulls",
                        "Los Angeles Kings",
                        "Golden State Warriors",
                        "Huston Rocket"
                    ]],
                    "answer": "Huston Rocket"
                }
            }
        }
    }`)

	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	_, err = runtime.GetFieldMaskJS(data)
	require.ErrorIs(t, err, runtime.ErrInvalidJSON)
}

func TestGetFieldMaskJS_1(t *testing.T) {
	t.Parallel()

	reader := strings.NewReader(`{
        "quiz": {
            "sport": {
                "q1": {
                    "question": "\"Which one is correct team name in NBA?"",
                    "options": [[
                        "New York Bulls",
                        "Los Angeles Kings",
                        "Golden State Warriors",
                        "Huston Rocket"
                    ]],
                    "answer": "Huston Rocket"
                }
            }
        }
    }`)

	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	_, err = runtime.GetFieldMaskJS(data)
	require.ErrorIs(t, err, runtime.ErrInvalidJSON)
}

func TestGetFieldMaskJS_2(t *testing.T) {
	t.Parallel()

	reader := strings.NewReader(`{
        "quiz": {
            "sport": {
                "q1": {
                    "question": "Which one is correct team name in NBA?",
                    "options": [[
                        "New York Bulls",
                        "Los Angeles Kings",
                        "Golden State Warriors",
                        "Huston Rocket"
                    ],
                    "answer": "Huston Rocket"
                }
            }
        }
    }`)

	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	_, err = runtime.GetFieldMaskJS(data)
	require.ErrorIs(t, err, runtime.ErrInvalidJSON)
}

func TestGetFieldMaskJS_3(t *testing.T) {
	t.Parallel()

	reader := strings.NewReader(`{
        "quiz": {
            "sport": {
                "q1": {
                    "question": "Which one is correct team name in NBA?",
                    [
                        "New York Bulls",
                        "Los Angeles Kings",
                        "Golden State Warriors",
                        "Huston Rocket"
                    ],
                    "answer": "Huston Rocket"
                }
            }
        }
      }`)

	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	_, err = runtime.GetFieldMaskJS(data)
	require.ErrorIs(t, err, runtime.ErrInvalidJSON)
}

func TestGetFieldMaskJS_4(t *testing.T) {
	t.Parallel()

	reader := strings.NewReader(`{
        "quiz": {
            "sport": {
                "q1": {
                    "question": "Which one is correct team name in NBA?",
                    "options": [
                        "New York Bulls",
                        "Los Angeles Kings",
                        "Golden State Warriors",
                        "Huston Rocket"
                    ],
                    "answer": "Huston Rocket"
                }
            },
            "maths": {
                "q1": {
                    "question": "5 + 7 = ?",
                    "options": [
                        "10",
                        "11",
                        "12",
                        "13"
                    ],
                    "answer": "12"
                },
                "q2": {
                    "question": "12 - 8 = ?",
                    "options": [
                        "1",
                        "2",
                        "3",
                        "4"
                    ],
                    "answer": "4"
                }
            }
        }
    }`)

	data, err := io.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	paths, err := runtime.GetFieldMaskJS(data)
	require.NoError(t, err)
	require.Len(t, paths, 15)
}
