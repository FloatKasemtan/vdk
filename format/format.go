package format

import (
	"github.com/FloatKasemtan/vdk/av/avutil"
	"github.com/FloatKasemtan/vdk/format/aac"
	"github.com/FloatKasemtan/vdk/format/flv"
	"github.com/FloatKasemtan/vdk/format/mp4"
	"github.com/FloatKasemtan/vdk/format/rtmp"
	"github.com/FloatKasemtan/vdk/format/rtsp"
	"github.com/FloatKasemtan/vdk/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
