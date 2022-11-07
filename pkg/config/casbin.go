package config

import (
	"errors"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	entadapter "github.com/casbin/ent-adapter"
	"github.com/zeromicro/go-zero/core/logx"
)

type CasbinConf struct {
	ModelText string `json:"ModelText,optional"`
}

func (l CasbinConf) NewCasbin(c DatabaseConf) (*casbin.Enforcer, error) {
	var dsn string
	switch c.Type {
	case "mysql":
		dsn = c.MysqlDSN()
	case "postgres":
		dsn = c.PostgresDSN()
	default:
		return nil, errors.New("unsupported database type")
	}

	adapter, err := entadapter.NewAdapter(c.Type, dsn)
	logx.Must(err)

	var text string
	if l.ModelText == "" {
		text = `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
	} else {
		text = l.ModelText
	}

	m, err := model.NewModelFromString(text)
	logx.Must(err)

	enforcer, err := casbin.NewEnforcer(m, adapter)
	logx.Must(err)

	err = enforcer.LoadPolicy()
	logx.Must(err)

	return enforcer, nil
}
