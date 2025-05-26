package component

import (
	"context"
	"net/url"
	
	"github.com/spf13/cast"
	
	"common/pkg/config/assets_config"
	"common/pkg/facade"
)

func GetAssetsURL(c context.Context, extension string) string {
	entryName := facade.Config(c).GetString(assets_config.EntryName)
	name := cast.ToString(c.Value(entryName + "." + extension))
	result, err := url.JoinPath("/static/dist/", name)
	if err != nil {
		return ""
	}
	return result
}
