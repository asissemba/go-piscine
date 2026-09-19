package handler

import (
	"path/filepath"
	"runtime"
)

func templatePath(name string) string {
	_, filename, _, _ := runtime.Caller(0)

	return filepath.Join(
		filepath.Dir(filename),
		"../../templates",
		name,
	)
}
