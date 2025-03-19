package logic

import (
	"context"
	"strings"

	"zhihu/app/applet/internal/code"
	"zhihu/app/applet/internal/svc"
	"zhihu/app/applet/internal/types"
	"zhihu/app/user_rpc/userclient"
	"zhihu/pkg/encrypt"
	"zhihu/pkg/token"
	"zhihu/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.LoginMobileEmpty
	}
	mobile, err := encrypt.EnMobile(req.Mobile)
	if err != nil {
		logx.Errorf("Encrypt mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &userclient.FindByMobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u == nil || u.UserId == 0 {
		return nil, xcode.AccessDenied
	}
	token, err := token.BuildToken(token.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		UserId: u.UserId,
	})
	if err != nil {
		logx.Errorf("BuildToken error: %v", err)
		return nil, err
	}
	return &types.LoginResponse{
		UserID: u.UserId,
		Token: types.Token{
			AccessToken: token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
