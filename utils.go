package router

func normalizePath(path string) string {
	if path[0] != '/' {
		path = "/" + path
	}
	if path[len(path)-1] == '/' {
		path = path[:len(path)-1]
	}
	return path
}