package jreader

import "testing"

func Test_jSONMap_StringValue(t *testing.T) {
	tests := []struct {
		name  string
		j     jSONMap
		want  string
		want1 bool
	}{
		{
			name:  "Test_jSONMap_StringValue",
			j:     jSONMap(map[string]any{"a": "b"}),
			want:  "{\"a\":\"b\"}",
			want1: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := tt.j.StringValue()
			if got != tt.want {
				t.Errorf("jSONMap.StringValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("jSONMap.StringValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
