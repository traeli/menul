package logic

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"menul-service/service/api/internal/svc"
	"menul-service/service/api/internal/types"
	"menul-service/service/cache"
	"menul-service/service/model"
	"net/http"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxLoginLogic {
	return &WxLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxLoginLogic) WxLogin(req *types.WxLoginReq) (resp *types.WxLoginReply, err error) {
	// 1. 根据 code 去微信服务器换取 openid 和 session_key
	appId := l.svcCtx.Config.WxMiniApp.AppId      // 配置里读取
	secret := l.svcCtx.Config.WxMiniApp.AppSecret // 配置里读取
	fmt.Println("AppId", l.svcCtx.Config.WxMiniApp.AppId)
	fmt.Println("AppSecret", l.svcCtx.Config.WxMiniApp.AppSecret)
	fmt.Println("req.Code", req.Code)

	wxApi := fmt.Sprintf(
		"https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appId, secret, req.Code,
	)

	httpResp, err := http.Get(wxApi)
	if err != nil {
		return nil, fmt.Errorf("请求微信接口失败: %v", err)
	}
	defer httpResp.Body.Close()

	var wxResp struct {
		Openid     string `json:"openid"`
		SessionKey string `json:"session_key"`
		ErrCode    int    `json:"errcode"`
		ErrMsg     string `json:"errmsg"`
	}
	err = json.NewDecoder(httpResp.Body).Decode(&wxResp)
	if err != nil {
		return nil, fmt.Errorf("解析微信接口返回失败: %v", err)
	}
	if wxResp.ErrCode != 0 {
		return nil, fmt.Errorf("微信接口错误: %v", wxResp.ErrMsg)
	}
	openid := wxResp.Openid
	if openid == "" {
		return nil, fmt.Errorf("微信登录失败: openid为空")
	}

	userTable := l.svcCtx.UserModel.User
	user, selectErr := userTable.WithContext(l.ctx).Where(userTable.OpenID.Eq(openid)).First()
	fmt.Println("selectErr", selectErr)
	if selectErr != nil && !errors.Is(selectErr, gorm.ErrRecordNotFound) {
		l.Logger.Errorf("查询错误", selectErr.Error())
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}
	isNew := false
	if user == nil {
		// 用户不存在，注册一个新用户
		newUser := &model.User{
			OpenID:    openid,
			Nickname:  req.Nickname,
			AvatarURL: req.AvatarUrl,
			Gender:    int32(req.Gender),
			CreateAt:  time.Now(),
			UpdateAt:  time.Now(),
			ID:        uuid.New().String(),
		}
		err = l.svcCtx.UserModel.WithContext(l.ctx).User.Create(newUser)
		if err != nil {
			return nil, fmt.Errorf("新用户注册失败: %v", err)
		}
		user = newUser
		isNew = true
	}

	token, err := l.generateToken(user.ID)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %v", err)
	}

	return &types.WxLoginReply{
		Token:  token,
		UserId: user.ID,
		IsNew:  isNew,
	}, nil
}

func (l *WxLoginLogic) generateToken(userId string) (string, error) {
	token, err := generateRandomToken()
	if err != nil {
		return "", err
	}
	l.svcCtx.Redis.Set(context.Background(), fmt.Sprintf(cache.UserTokenKey, token), userId, time.Hour*24)
	return fmt.Sprintf(token), nil
}

func generateRandomToken() (string, error) {
	b := make([]byte, 32) // 256位
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
