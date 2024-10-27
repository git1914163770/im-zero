package test

import (
	"context"
	"im-zero/apps/user/rpc/internal/config"
	"im-zero/apps/user/rpc/internal/logic"
	"im-zero/apps/user/rpc/internal/svc"
	"im-zero/apps/user/rpc/user"
	"path/filepath"
	"testing"

	"github.com/zeromicro/go-zero/core/conf"
)

var svcCtx *svc.ServiceContext

func init() {
	var c config.Config
	conf.MustLoad(filepath.Join("../../etc/dev/user.yaml"), &c)
	svcCtx = svc.NewServiceContext(c)
}

func TestRegisterLogic_Register(t *testing.T) {
	type args struct {
		in *user.RegisterReq
	}
	tests := []struct {
		name      string
		args      args
		wantPrint bool
		wantErr   bool
	}{
		{
			"1", args{in: &user.RegisterReq{
			Phone:    "13000000000",
			Nickname: "123456",
			Password: "123456",
			Avatar:   "png.jpg",
			Sex:      1,
		}}, true, false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := logic.NewRegisterLogic(context.Background(), svcCtx)
			got, err := l.Register(tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("Register() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantPrint {
				t.Log(tt.name, got)
			}
		})
	}
}
