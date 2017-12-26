package excelize

var DefaultFormatChart *formatChart  // Todo - WIP

func init() {
	DefaultFormatChart =  &formatChart{
		Format: formatPicture{
			FPrintsWithSheet: true,
			FLocksWithSheet:  false,
			NoChangeAspect:   false,
			OffsetX:          0,
			OffsetY:          0,
			XScale:           1.0,
			YScale:           1.0,
		},
		Legend: formatChartLegend{
			Position:      "bottom",
			ShowLegendKey: false,
		},
		Title: formatChartTitle{
			Name: " ",
		},
		ShowBlanksAs: "gap",
	}
}
