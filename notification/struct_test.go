package notification

import (
	"github.com/ismdeep/justoj-monitor/config"
	"testing"
)

func TestSend(t *testing.T) {
	config.LoadConfig()

	type args struct {
		token string
		pack  *Pack
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSend-001",
			args: args{
				token: config.Notification.Token,
				pack: &Pack{
					SenderName: "Test Sender Name",
					Subject:    "Test Subject",
					Type:       "text/plain",
					Content:    "Hello",
					ToMailList: []string{"l.jiang.1024@gmail.com"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Send(tt.args.pack); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
