package runtime

import (
	"errors"
)

var ErrInvalidJSON = errors.New("invalid json")

func GetFieldMaskJS(b []byte) (FieldMask, error) {
	prefix := []string{}
	waitKey := false
	paths := make([][]string, 0)
	opensObj := 0
	opensArr := 0

	for i := 0; i < len(b); i++ {
		if b[i] == '"' {
			for j := i + 1; j < len(b); j++ {
				if b[j] == '\\' {
					j++

					continue
				}

				if b[j] == '"' {
					if waitKey {
						paths = append(paths, append(prefix, string(b[i+1:j])))
						waitKey = false
					}

					i = j

					break
				}

				if j == len(b)-1 {
					return nil, ErrInvalidJSON
				}
			}

			continue
		}

		if b[i] == '[' {
			if waitKey {
				return nil, ErrInvalidJSON
			}

			opensArr++

			continue
		}

		if b[i] == ']' {
			opensArr--

			continue
		}

		if b[i] == '{' {
			if waitKey {
				return nil, ErrInvalidJSON
			}

			if opensArr == 0 {
				waitKey = true
			}

			opensObj++

			if len(paths) > 0 {
				prefix = paths[len(paths)-1]
			}

			continue
		}

		if b[i] == '}' {
			opensObj--

			if len(prefix) > 0 {
				prefix = prefix[:len(prefix)-1]
			}

			continue
		}

		if b[i] == ',' {
			if opensArr == 0 {
				waitKey = true
			}

			continue
		}
	}

	if opensObj != 0 || opensArr != 0 {
		return nil, ErrInvalidJSON
	}

	return paths, nil
}
