package interceptors

import (
	"context"
	"zhihu/pkg/xcode"

	"google.golang.org/grpc"
)

// Xcode -> *status.Status -> error接口（供后续grpc框架调用 -> *status.Status协议元数据）
func ServerErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		resp, err = handler(ctx, req)
		return resp, xcode.FromError(err).Err()
	}
}

// 不封装：
// 原始的error由rpc返回bff层，bff层无法根据错误码分析错误原因，只能将错误信息原样返回给前端（或进行字符串匹配，不方便）
// 封装：
// rpc层返回的error被封装成xcode，bff层可以根据错误码分析错误原因，对错误进行进一步处理
