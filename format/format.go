package format

import (
	"github.com/edgeworthsecurity/vdk/av/avutil"
	"github.com/edgeworthsecurity/vdk/format/aac"
	"github.com/edgeworthsecurity/vdk/format/flv"
	"github.com/edgeworthsecurity/vdk/format/mp4"
	"github.com/edgeworthsecurity/vdk/format/rtmp"
	"github.com/edgeworthsecurity/vdk/format/rtsp"
	"github.com/edgeworthsecurity/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
