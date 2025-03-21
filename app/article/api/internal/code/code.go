package code

import "zhihu/pkg/xcode"

var (
	GetBuckErr                 = xcode.Define(30001, "获取bucket实例失败")
	PutBuckErr                 = xcode.Define(30002, "上传bucket失败")
	GetObjDetailErr            = xcode.Define(30003, "获取对象详细信息失败")
	ArticleTitleEmpty          = xcode.Define(30004, "文章标题为空")
	ArticleContentTooFewWords = xcode.Define(30005, "文章内容字数太少")
	ArticleCoverEmpty          = xcode.Define(30006, "文章封面为空")
)
