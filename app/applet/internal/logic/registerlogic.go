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

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.MobileEmpty
	}
	req.Name = strings.TrimSpace(req.Name)
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		return nil, code.PasswordEmpty
	} else {
		req.Password, err = encrypt.EnPassword(req.Password)
		if err != nil {
			logx.Errorf("encrypt.EnPassword error: password: %s err: %v", req.Password, err)
			return nil, err
		}
	}
	mobile, err := encrypt.EnMobile(req.Mobile)
	if err != nil {
		logx.Errorf("encrypt.EnMobile error: mobile: %s err: %v", req.Mobile, err)
		return nil, err
	}
	// 从rpc服务中根据手机号查找用户
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &userclient.FindByMobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u != nil && u.UserId > 0 {
		return nil, code.MobileHasRegister
	}
	// 用户注册得到uid
	regRes, err := l.svcCtx.UserRPC.Register(l.ctx, &userclient.RegisterRequest{
		Username: req.Name,
		Mobile:   mobile,
	})
	if err != nil {
		logx.Errorf("Register error: %v", err)
		return nil, err
	}

	token, err := token.BuildToken(token.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		UserId:       regRes.UserId,
	})
	if err != nil {
		logx.Errorf("BuildToken error: %v", err)
		return nil, err
	}
	// TODO: 密码处理
	return &types.RegisterResponse{
		UserID: regRes.UserId,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
	}, nil
}
