package installer

import (
	"github.com/clnbs/hulotte/internal/app/helper"
)

func DoesHulotteExists() (bool, error) {
	path := GetHulottePath()
	return helper.DoesFileExsist(path)
}
