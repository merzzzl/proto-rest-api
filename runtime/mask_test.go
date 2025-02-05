package runtime_test

import (
	"testing"

	pb "github.com/merzzzl/proto-rest-api/example/api"
	"github.com/merzzzl/proto-rest-api/runtime"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/encoding/protojson"
)

func TestMergeByMask(t *testing.T) {
	t.Parallel()

	t.Run("merge fields by mask", func(t *testing.T) {
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

		require.Equal(t, "hi!", msg.GetMessage(), "message should be updated")
		require.Empty(t, msg.GetAuthor().GetEmail(), "email should be cleared")
		require.Equal(t, "+79999999999", msg.GetAuthor().GetPhone(), "phone should be updated")
	})
}
