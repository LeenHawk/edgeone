package edgeone

import (
	"testing"

	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/libdns/edgeone"
)

func TestUnmarshalCaddyfile(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantErr      bool
		secretId     string
		secretKey    string
		sessionToken string
		region       string
	}{
		{
			name: "basic config",
			input: `edgeone {
				secret_id "test_id"
				secret_key "test_key"
			}`,
			wantErr:   false,
			secretId:  "test_id",
			secretKey: "test_key",
		},
		{
			name: "config with session token and region",
			input: `edgeone {
				secret_id "test_id"
				secret_key "test_key"
				session_token "test_token"
				region "ap-hongkong"
			}`,
			wantErr:      false,
			secretId:     "test_id",
			secretKey:    "test_key",
			sessionToken: "test_token",
			region:       "ap-hongkong",
		},
		{
			name: "missing secret_id",
			input: `edgeone {
				secret_key "test_key"
			}`,
			wantErr: true,
		},
		{
			name: "missing secret_key",
			input: `edgeone {
				secret_id "test_id"
			}`,
			wantErr: true,
		},
		{
			name: "unknown directive",
			input: `edgeone {
				secret_id "test_id"
				secret_key "test_key"
				unknown_field "value"
			}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Provider{new(edgeone.Provider)}
			d := caddyfile.NewTestDispenser(tt.input)

			err := p.UnmarshalCaddyfile(d)

			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalCaddyfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if p.Provider.SecretId != tt.secretId {
					t.Errorf("SecretId = %v, want %v", p.Provider.SecretId, tt.secretId)
				}
				if p.Provider.SecretKey != tt.secretKey {
					t.Errorf("SecretKey = %v, want %v", p.Provider.SecretKey, tt.secretKey)
				}
				if p.Provider.SessionToken != tt.sessionToken {
					t.Errorf("SessionToken = %v, want %v", p.Provider.SessionToken, tt.sessionToken)
				}
				if p.Provider.Region != tt.region {
					t.Errorf("Region = %v, want %v", p.Provider.Region, tt.region)
				}
			}
		})
	}
}

func TestCaddyModule(t *testing.T) {
	p := Provider{}
	info := p.CaddyModule()

	expectedID := "dns.providers.edgeone"
	if string(info.ID) != expectedID {
		t.Errorf("CaddyModule().ID = %v, want %v", info.ID, expectedID)
	}

	if info.New == nil {
		t.Error("CaddyModule().New is nil")
	}

	// Test creating a new instance
	module := info.New()
	if _, ok := module.(*Provider); !ok {
		t.Errorf("New() returned %T, want *Provider", module)
	}
}
