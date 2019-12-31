package GameInfo

type GmInfo struct {
	Title string
}

type GUI struct {
	Height int32
	Width  int32
}

var GInfo = GmInfo{"Game"}
var Screen = GUI{800, 600}
