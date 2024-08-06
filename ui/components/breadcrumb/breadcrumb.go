package breadcrumb

func isBreadcrumbActive(isActive bool) string {
	if isActive {
		return "page"
	}
	return ""
}
