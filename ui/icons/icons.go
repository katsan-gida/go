package icons

import "strings"

type iconSize int

const (
	SizeXSmall iconSize = iota
	SizeSmall
	SizeMedium
	SizeLarge
)

type IconProps struct {
	Size iconSize
}

func getIconClass(props IconProps) string {
	base := strings.Builder{}

	switch props.Size {
	case SizeXSmall:
		base.WriteString("size-4")
	case SizeSmall:
		base.WriteString("size-5")
	case SizeMedium:
		base.WriteString("size-6")
	case SizeLarge:
		base.WriteString("size-7")
	default:
		base.WriteString("size-6")
	}
	return base.String()
}
