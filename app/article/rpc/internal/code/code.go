package code

import "zhihu/pkg/xcode"

var (
	SortTypeInvalid         = xcode.Define(60001, "排序类型无效")
	UserIdInvalid           = xcode.Define(60002, "用户ID无效")
	ArticleTitleCantEmpty   = xcode.Define(60003, "文章标题不能为空")
	ArticleContentCantEmpty = xcode.Define(60004, "文章内容不能为空")
	ArticleIdInvalid        = xcode.Define(60005, "文章ID无效")
)
