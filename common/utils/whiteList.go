package utils

func DefaultIsWhiteListFunc(pathList []string) func(path string) bool {
	return func(path string) bool {
		for _, p := range pathList {
			if path == "/api/v1/auth/refresh" || p == path {
				return true
			}

		}
		return false
	}

}
