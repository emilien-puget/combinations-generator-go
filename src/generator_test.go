package combinations

import (
	"reflect"
	"testing"
)

func Test_generator(t *testing.T) {
	type args struct {
		characters []string
		length     uint
		separator  string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				characters: []string{"a", "b"},
				length:     2,
				separator:  "",
			},
			want:    []string{"aa", "ba", "ab", "bb"},
			wantErr: false,
		},
		{
			name: "test",
			args: args{
				characters: []string{"a", "b", "c"},
				length:     2,
				separator:  "",
			},
			want:    []string{"aa", "ba", "ca", "ab", "bb", "cb", "ac", "bc", "cc"},
			wantErr: false,
		},
		{
			name: "test",
			args: args{
				characters: []string{"a", "b"},
				length:     3,
				separator:  ".",
			},
			want:    []string{"a.a.a", "b.a.a", "a.b.a", "b.b.a", "a.a.b", "b.a.b", "a.b.b", "b.b.b"},
			wantErr: false,
		},
		{
			name: "test",
			args: args{
				characters: []string{"jean", "pierre"},
				length:     2,
				separator:  ".",
			},
			want:    []string{"jean.jean", "pierre.jean", "jean.pierre", "pierre.pierre"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			returnChan := MakeChan(tt.args.characters, tt.args.length)
			go Generator(tt.args.characters, tt.args.length, tt.args.separator, &returnChan)

			var results []string
			for result := range returnChan {
				results = append(results, result)

			}
			if !reflect.DeepEqual(results, tt.want) {
				t.Errorf("generator() = %v, want %v", results, tt.want)
			}
		})
	}
}
