package ImageStatus

type ImageStatus string

const (
	Missing     ImageStatus = "missing"
	Placeholder             = "placeholder"
	LowRes                  = "lowres"
	HighresScan             = "highres_scan"
)
