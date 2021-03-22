package token

import (
	"github.com/kaanaktas/openbanking-paymentinitiation/api"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/cache"
	"github.com/kaanaktas/openbanking-paymentinitiation/internal/store"
	cfg "github.com/kaanaktas/openbanking-paymentinitiation/pkg/configmanager"
	"testing"
)

func Test_service_GetAccessToken(t *testing.T) {
	type fields struct {
		cfg cfg.Service
	}
	type args struct {
		aspspId   string
		scopeType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"get_access_token_wrong_aspsp_id",
			fields{cfg: cfg.NewService(cfg.NewRepository(store.LoadDBConnection()), cache.LoadInMemory())},
			args{aspspId: "hsbc", scopeType: api.ScopePayments},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				cfg: tt.fields.cfg,
			}
			got, err := s.GetAccessToken(tt.args.aspspId, tt.args.scopeType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr == false && got != "" && len(got) <= 0 {
				t.Errorf("GetAccessToken() got = %v", got)
			}
		})
	}
}

func Test_service_RefreshAccessToken(t *testing.T) {
	type fields struct {
		cfg cfg.Service
	}
	type args struct {
		aspspId          string
		scopeType        string
		refreshTokenData string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"refresh_access_token_expect_invalid_grant",
			fields{cfg: cfg.NewService(cfg.NewRepository(store.LoadDBConnection()), cache.LoadInMemory())},
			args{aspspId: "danske", scopeType: api.ScopePayments, refreshTokenData: "dummy_token"},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				cfg: tt.fields.cfg,
			}
			got, err := s.RefreshAccessToken(tt.args.aspspId, tt.args.scopeType, tt.args.refreshTokenData)
			if (err != nil) != tt.wantErr {
				t.Errorf("RefreshAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil {
				return
			}
		})
	}
}

func Test_service_GetResourceAccessRefreshToken(t *testing.T) {
	type fields struct {
		cfg cfg.Service
	}
	type args struct {
		aspspId  string
		authCode string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			"get_resource_access_refresh_token_success",
			fields{cfg: cfg.NewService(cfg.NewRepository(store.LoadDBConnection()), cache.LoadInMemory())},
			args{
				aspspId:  "danske",
				authCode: "dummy_auth_code",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := service{
				cfg: tt.fields.cfg,
			}
			got, err := s.GetResourceAccessRefreshToken(tt.args.aspspId, tt.args.authCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetResourceAccessRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil {
				return
			}
		})
	}
}
