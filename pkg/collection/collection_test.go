package collection

import (
	"reflect"
	"testing"
)

func TestGetCollection(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       func(t *testing.T) args
		want1      *Collection
		wantErr    bool
		inspectErr func(err error, t *testing.T) //use for more precise error evaluation after test
	}{

		// BUG: Collection should check if directory exists, create if it doesn't.
		{
			name: "happy path get collection",
			args: func(t *testing.T) args {
				return args{
					name: "test_collection",
				}
			},
			want1:   &Collection{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tArgs := tt.args(t)

			got1, err := GetCollection(tArgs.name)

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetCollection got1 = %v, want1: %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Fatalf("GetCollection error = %v, wantErr: %t", err, tt.wantErr)
			}

			if tt.inspectErr != nil {
				tt.inspectErr(err, t)
			}
		})
	}
}
