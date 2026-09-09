package groups

import (
	"maps"

	"github.com/spf13/viper"
)

type GroupMap map[string][]string

func BuildGroupMap() GroupMap {
	gm := make(GroupMap)
	groups := viper.GetStringMapStringSlice("groups")
	maps.Copy(gm, groups)
	return gm
}
