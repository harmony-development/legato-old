package v1

import (
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/dyatlov/go-opengraph/opengraph"
	"github.com/labstack/echo/v4"

	mediaproxyv1 "github.com/harmony-development/legato/gen/mediaproxy/v1"
	"github.com/harmony-development/legato/server/api/middleware"
	lru "github.com/hashicorp/golang-lru"
)

type instantViewData struct {
	instantLRU *lru.ARCCache
}

type instantViewResult struct {
	ok   bool
	data string
}

var converter = md.NewConverter("", true, nil)

func init() {
	middleware.RegisterRPCConfig(middleware.RPCConfig{
		RateLimit: middleware.RateLimit{
			Duration: 5 * time.Second,
			Burst:    1,
		},
	}, "/protocol.mediaproxy.v1.MediaProxyService/InstantView")
}

func (v1 *V1) CanInstantView(c echo.Context, r *mediaproxyv1.InstantViewRequest) (resp *mediaproxyv1.CanInstantViewResponse, err error) {
	data, err := v1.fetch(r.Url)
	if err != nil {
		return nil, err
	}

	return &mediaproxyv1.CanInstantViewResponse{
		CanInstantView: data.RD != nil,
	}, nil
}

// InstantView implements the InstantView RPC
func (v1 *V1) InstantView(c echo.Context, r *mediaproxyv1.InstantViewRequest) (resp *mediaproxyv1.InstantViewResponse, err error) {
	data, err := v1.fetch(r.Url)
	if err != nil {
		return nil, err
	}

	content := ""
	if data.RD != nil {
		content = *data.RD
	}

	sm := &mediaproxyv1.InstantViewResponse{
		Metadata: &mediaproxyv1.SiteMetadata{},
		Content:  content,
		IsValid:  data.RD != nil,
	}

	copyOGIntoProtobuf((*opengraph.OpenGraph)(data.OG), sm.Metadata)

	return sm, nil
}
