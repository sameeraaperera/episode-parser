package server_test

import (
	"testing"

	"github.com/sap4001/episode-parser/internal/server"
	"github.com/stretchr/testify/assert"
)

func Test_GetListenPort(t *testing.T) {
	tests := []struct {
		name     string
		portStr  string
		wantPort int
		wantErr  bool
		errMsg   string
	}{
		{
			name:     "Missing port",
			portStr:  "",
			wantPort: -1,
			wantErr:  true,
			errMsg:   "listen port not defined",
		},
		{
			name:     "Invalid port",
			portStr:  "corrupt",
			wantPort: -1,
			wantErr:  true,
			errMsg:   "error parsing port",
		},
		{
			name:     "Valid port",
			portStr:  "8080",
			wantPort: 8080,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := server.GetListenPort(tt.portStr)
			if tt.wantErr {
				assert.NotNil(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.Nil(t, err)
			}
			assert.Equal(t, tt.wantPort, got)
		})
	}
}
