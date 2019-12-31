package GameInfo

type GmInfo struct {
	Title string
}

type GUI struct {
	Width  int32
	Height int32
}

var GInfo = GmInfo{"Game"}
var Screen = GUI{800, 450}
