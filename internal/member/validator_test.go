package member

import "testing"

func TestValidateName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "+Name is valid",
			args: args{
				name: "John Doe",
			},
			wantErr: false,
		},
		{
			name: "-Name is empty",
			args: args{
				name: "",
			},
			wantErr: true,
		},
		{
			name: "-Name too short",
			args: args{
				name: "1",
			},
			wantErr: true,
		},
		{
			name: "-Name too long",
			args: args{
				name: "111111111111111111111111111111111111111111111111111",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateName(tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("ValidateName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
