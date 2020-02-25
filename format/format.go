package format

import (
	"github.com/edgewothsecurty/vdk/av/avutil"
	"github.com/edgewothsecurty/vdk/format/aac"
	"github.com/edgewothsecurty/vdk/format/flv"
	"github.com/edgewothsecurty/vdk/format/mp4"
	"github.com/edgewothsecurty/vdk/format/rtmp"
	"github.com/edgewothsecurty/vdk/format/rtsp"
	"github.com/edgewothsecurty/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
