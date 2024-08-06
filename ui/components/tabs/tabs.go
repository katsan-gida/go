package tabs

func isTabActive(props TabItemProps) string {
	if props.ActiveHref == props.Href {
		return "page"
	}
	return ""
}
