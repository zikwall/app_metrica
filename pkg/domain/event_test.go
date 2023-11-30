package domain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zikwall/app_metrica/pkg/testkit"
)

func TestEventDatetime_UnmarshalJSON(t *testing.T) {
	type args struct {
		b []byte
	}
	type dated struct {
		DateTime EventDatetime `json:"date_time"`
	}

	tests := []struct {
		name    string
		args    args
		want    EventDatetime
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "successfully unmarshal time",
			args: args{
				b: []byte("{\"date_time\": \"2023-11-16 06:00:04\"}"),
			},
			want: EventDatetime{Time: testkit.MustTime("2023-11-16T06:00:04Z")},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.NoError(t, err)
			},
		},
		{
			name: "failed unmarshal time",
			args: args{
				b: []byte("{\"date_time\": \"2023-11-16T06:00:04\"}"),
			},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.Error(t, err)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var d dated
			err := json.Unmarshal(tt.args.b, &d)
			tt.wantErr(t, err)

			assert.Equal(t, d.DateTime, tt.want)
		})
	}
}
