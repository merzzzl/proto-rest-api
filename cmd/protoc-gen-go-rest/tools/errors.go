package tools

import "errors"

var (
	ErrPathAlreadyExists    = errors.New("path already exists")
	ErrFieldNotFound        = errors.New("field not found")
	ErrUnsupportedFieldType = errors.New("unsupported field type")
	ErrInvalidPath          = errors.New("invalid path")
)
