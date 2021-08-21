package installer

import (
	"github.com/clnbs/hulotte/internal/app/helper"
)

func DoesHulotteExists() (bool, error) {
	path := GetHullotePath()
	return helper.DoesFileExsist(path)
}
