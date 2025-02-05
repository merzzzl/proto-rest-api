package tools

import "errors"

var ErrPathAlreadyExists = errors.New("path already exists")
var ErrFieldNotFound = errors.New("field not found")
var ErrUnsupportedFieldType = errors.New("unsupported field type")
var ErrInvalidPath = errors.New("invalid path")
