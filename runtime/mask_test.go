package runtime_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"

	pb "github.com/merzzzl/proto-rest-api/example/api"
	"github.com/merzzzl/proto-rest-api/runtime"
)

func TestMergeByMask_0(t *testing.T) {
	t.Parallel()

	msg := &pb.Message{
		Message: "hello",
		Author: &pb.Author{
			Name: "Alex",
			Contact: &pb.Author_Email{
				Email: "alex@example.org",
			},
		},
	}
	js := `{"message":"hi!","author":{"phone":"+79999999999"}}`

	fm, err := runtime.GetFieldMaskJS([]byte(js))
	require.NoError(t, err)

	var in pb.Message

	err = protojson.Unmarshal([]byte(js), &in)
	require.NoError(t, err)

	runtime.MergeByMask(&in, msg, fm)
	require.Equal(t, msg.GetMessage(), in.GetMessage())
	require.Empty(t, msg.GetAuthor().GetEmail())
	require.Equal(t, msg.GetAuthor().GetPhone(), in.GetAuthor().GetPhone())
}
