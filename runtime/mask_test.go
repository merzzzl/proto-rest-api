package runtime_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/merzzzl/proto-rest-api/example/gen/go/example"
	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestMergeByMask_0(t *testing.T) {
	t.Parallel()

	msg := &example.Message{
		Message: "hello",
		Author: &example.Author{
			Name: "Alex",
			Contact: &example.Author_Email{
				Email: "alex@example.org",
			},
		},
	}
	js := `{"message":"hi!","author":{"phone":"+79999999999"}}`

	fm, err := runtime.GetFieldMaskJS([]byte(js))
	require.NoError(t, err)

	var in example.Message

	err = protojson.Unmarshal([]byte(js), &in)
	require.NoError(t, err)

	runtime.MergeByMask(&in, msg, fm)
	require.Equal(t, msg.GetMessage(), in.GetMessage())
	require.Empty(t, msg.GetAuthor().GetEmail())
	require.Equal(t, msg.GetAuthor().GetPhone(), in.GetAuthor().GetPhone())
}
