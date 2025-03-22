package logic

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"zhihu/app/article/api/internal/code"
	"zhihu/app/article/api/internal/svc"
	"zhihu/app/article/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10 * 2^20(byte), 即10MB

type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadCoverLogic) UploadCover(req *http.Request) (resp *types.UploadCoverResponse, err error) {
	_ = req.ParseMultipartForm(maxFileSize)
	file, handler, err := req.FormFile("cover")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bucket, err := l.svcCtx.OssClient.Bucket(l.svcCtx.Config.Oss.BucketName)
	if err != nil {
		logx.Errorf("get bucket failed, err: %v", err)
		return nil, code.GetBuckErr
	}
	objectKey := genFileName(handler.Filename)
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		logx.Errorf("put object failed, err: %v", err)
		return nil, code.PutBuckErr
	}

	return &types.UploadCoverResponse{CoverUrl: genFileURL(objectKey)}, nil
}

func genFileName(filename string) string {
	return fmt.Sprintf("%d_%s", time.Now().Unix(), filename)
}

func genFileURL(objectKey string) string {
	return fmt.Sprintf("https://xxx/%s", objectKey)
}
