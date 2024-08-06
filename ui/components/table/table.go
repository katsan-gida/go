package table

func getTableHeaderClass(props TableHeaderCellProps) string {
	base := "h-10 bg-slate-200 px-2 text-left font-medium text-slate-500 shadow first-of-type:rounded-l-lg last-of-type:rounded-r-lg"

	switch props.Align {
	case TableHeaderAlignCenter:
		base += " text-center"
	case TableHeaderAlignRight:
		base += " text-right"
	default:
		base += " text-left"
	}

	return base
}

func getTableCellClass(props TableCellProps) string {
	base := "h-12 px-2"

	switch props.Align {
	case TableCellAlignCenter:
		base += " text-center"
	case TableCellAlignRight:
		base += " text-right"
	default:
		base += " text-left"
	}

	return base
}
