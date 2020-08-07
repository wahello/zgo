package api

import (
	"github.com/BurntSushi/toml"
	"github.com/casbin/casbin/v2"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/suisrc/zgo/app/service"
	"github.com/suisrc/zgo/middleware"
	"github.com/suisrc/zgo/middlewire"
	"github.com/suisrc/zgo/modules/auth"
	"github.com/suisrc/zgo/modules/auth/jwt"
	"github.com/suisrc/zgo/modules/auth/jwt/store/buntdb"
	casbinjson "github.com/suisrc/zgo/modules/casbin/adapter/json"
	"github.com/suisrc/zgo/modules/config"
	"github.com/suisrc/zgo/modules/logger"
	"golang.org/x/text/language"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// EndpointSet wire注入声明
var EndpointSet = wire.NewSet(
	service.ServiceSet,             // 系统提供的服务列表
	wire.Struct(new(Options), "*"), // 初始化接口参数
	InitEndpoints,                  // 初始化接口方法
	casbinjson.CasbinAdapterSet,    // Casbin依赖
	NewAuther,                      // Auther注册
	NewBundle,                      // 国际化注册

	// 接口注册
	wire.Struct(new(Auth), "*"),
	wire.Struct(new(Signin), "*"),
	wire.Struct(new(User), "*"),
)

//=====================================
// Endpoint
//=====================================

// Options options
type Options struct {
	Engine   *gin.Engine
	Enforcer *casbin.SyncedEnforcer
	Auther   auth.Auther
	Bundle   *i18n.Bundle
	Router   middlewire.Router

	// 接口注入
	Auth   *Auth
	Signin *Signin
	User   *User
}

// Endpoints result
type Endpoints struct {
}

// InitEndpoints init
func InitEndpoints(o *Options) *Endpoints {
	// 国际化
	o.Engine.Use(middleware.I18nMiddleware(o.Bundle))

	// 在根路由注册通用授权接口, (没有ContextPath限定,一般是给nginx使用)
	// 在nginx注册认证接口时候,请放行zgo服务器其他接口,防止重复认证
	o.Auth.RegisterWithUAC(o.Engine)

	// ContextPath路由
	r := o.Router

	// 服务器授权控制器
	uac := middleware.UserAuthCasbinMiddleware(
		o.Auther,
		o.Enforcer,
		middleware.AllowPathPrefixSkipper(
			// sign 登陆接口需要排除
			// 注意[/api/sign,都会被排除]
			middleware.JoinPath(config.C.HTTP.ContextPath, "sign"),
			// pub => public 为系统公共信息
			// 注意[/api/pub开头的,都会被排除]
			middleware.JoinPath(config.C.HTTP.ContextPath, "pub"),
		),
	)
	// 增加权限认证
	r.Use(uac)

	// 注册登陆接口
	o.Auth.Register(r)
	o.Signin.Register(r)
	o.User.Register(r)

	return &Endpoints{}
}

// NewAuther of auth.Auther
// 注册认证使用的auther内容
func NewAuther() auth.Auther {
	store, err := buntdb.NewStore(":memory:") // 使用内存缓存
	if err != nil {
		panic(err)
	}
	secret := config.C.JWTAuth.SigningSecret
	if secret == "" {
		secret = auth.UUID(128)
		logger.Infof(nil, "jwt secret: %s", secret)
	}
	auther := jwt.New(store,
		jwt.SetSigningSecret(secret), // 注册令牌签名密钥
	)

	return auther
}

// NewBundle 国际化
func NewBundle() *i18n.Bundle {
	//bundle := i18n.NewBundle(language.Chinese)
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("locales/active.zh-CN.toml")
	bundle.LoadMessageFile("locales/active.en-US.toml")
	return bundle
}
