// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package injector

import (
	"github.com/suisrc/zgo/app/api"
	"github.com/suisrc/zgo/app/model/entc"
	"github.com/suisrc/zgo/app/model/sqlxc"
	"github.com/suisrc/zgo/app/service"
	"github.com/suisrc/zgo/middlewire"
	"github.com/suisrc/zgo/modules/casbin"
	"github.com/suisrc/zgo/modules/casbin/watcher/mem"
	"github.com/suisrc/zgo/modules/passwd"
)

// Injectors from wire.go:

func BuildInjector() (*Injector, func(), error) {
	bundle := NewBundle()
	useEngine := api.NewUseEngine(bundle)
	engine := middlewire.InitGinEngine(useEngine)
	router := middlewire.NewRouter(engine)
	client, cleanup, err := entc.NewClient()
	if err != nil {
		return nil, nil, err
	}
	db, cleanup2, err := sqlxc.NewClient()
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	gpa := service.GPA{
		Entc: client,
		Sqlx: db,
	}
	casbinAdapter := service.CasbinAdapter{
		GPA: gpa,
	}
	syncedEnforcer, cleanup3, err := casbin.NewCasbinEnforcer(casbinAdapter)
	if err != nil {
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	authOpts := &service.AuthOpts{
		GPA: gpa,
	}
	auther := service.NewAuther(authOpts)
	demo := &api.Demo{
		GPA: gpa,
	}
	auth := &api.Auth{
		Enforcer: syncedEnforcer,
		Auther:   auther,
	}
	validator := &passwd.Validator{}
	signin := service.Signin{
		GPA:    gpa,
		Passwd: validator,
	}
	apiSignin := &api.Signin{
		GPA:           gpa,
		Auther:        auther,
		SigninService: signin,
	}
	user := &api.User{
		GPA: gpa,
	}
	options := &api.Options{
		Engine:   engine,
		Router:   router,
		Enforcer: syncedEnforcer,
		Auther:   auther,
		Demo:     demo,
		Auth:     auth,
		Signin:   apiSignin,
		User:     user,
	}
	endpoints := api.InitEndpoints(options)
	watcher, cleanup4, err := casbinmem.NewCasbinWatcher(casbinAdapter, syncedEnforcer)
	if err != nil {
		cleanup3()
		cleanup2()
		cleanup()
		return nil, nil, err
	}
	i18n := &service.I18n{
		GPA:    gpa,
		Bundle: bundle,
	}
	i18nLoader := service.InitI18nLoader(i18n)
	swagger := middlewire.NewSwagger(engine)
	healthz := middlewire.NewHealthz(engine)
	injector := &Injector{
		Engine:     engine,
		Endpoints:  endpoints,
		Bundle:     bundle,
		Enforcer:   syncedEnforcer,
		Auther:     auther,
		Watcher:    watcher,
		I18nLoader: i18nLoader,
		Swagger:    swagger,
		Healthz:    healthz,
	}
	return injector, func() {
		cleanup4()
		cleanup3()
		cleanup2()
		cleanup()
	}, nil
}
