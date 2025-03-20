package tools

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func LineComments(comm protogen.CommentSet) string {
	var comment string

	if comm.Leading != "" {
		comment = strings.TrimPrefix(comm.Leading.String(), "//")
		comment = strings.TrimSpace(comment)
	}

	if comm.Trailing != "" {
		if comment != "" {
			comment += "\\n"
		}

		comment = strings.TrimPrefix(comm.Trailing.String(), "//")
		comment = strings.TrimSpace(comment)
	}

	return comment
}

func LineComment(comm protogen.Comments) string {
	var comment string

	if comm != "" {
		comment = strings.TrimPrefix(comm.String(), "//")
		comment = strings.TrimSpace(comment)
	}

	return comment
}
