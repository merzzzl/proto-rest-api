package openapi

import "strings"

func FormatedPath(path string) string {
	path = strings.Split(path, "?")[0]
	segs := strings.Split(path, "/")

	for i, seg := range segs {
		if strings.HasPrefix(seg, ":") {
			segs[i] = "{" + seg[1:] + "}"
		}
	}

	return strings.Join(segs, "/")
}
