package utils

import (
	"github.com/importcjj/sensitive"
	"wordma/config"
)

var CommentFilter *sensitive.Filter

func InitCommentFilter() {
	CommentFilter = sensitive.New()
	err := CommentFilter.LoadWordDict(config.FilterPath)
	if err != nil {
		panic("加载评论敏感词数据库出错" + err.Error())
	}
}
