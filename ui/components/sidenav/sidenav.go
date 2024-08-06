package sidenav

func isActiveLink(currentPath, href string) string {
	if currentPath == href {
		return "page"
	}
	return ""
}
