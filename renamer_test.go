package renamer_test

import (
	"testing"

	"github.com/go-generalize/renamer"
)

func Test_Renamer(t *testing.T) {
	tests := []struct {
		name        string
		defaultOpts []renamer.Option
		caller      func(r *renamer.Renamer, ch chan<- string)
		want        []string
	}{
		{
			name: "success",
			defaultOpts: []renamer.Option{
				renamer.WithPackageName,
			},
			caller: func(r *renamer.Renamer, ch chan<- string) {
				ch <- r.Renamed("example.com/foo/bar.Status", "Status", renamer.FromFullPath("example.com/foo/bar.Status"))
				ch <- r.Renamed("example.com/foo/bar/hoge.Status", "Status", renamer.FromFullPath("example.com/foo/bar/hoge.Status"))
				ch <- r.Renamed("example.com/foo/bar.Status", "Status", renamer.FromFullPath("example.com/foo/bar.Status"))
				ch <- r.Renamed("example.com/foo/bar/hoge.Status", "Status", renamer.FromFullPath("example.com/foo/bar/hoge.Status"))
				ch <- r.Renamed("example.com/foo/bar/hoge/hoge.Status", "Status", renamer.FromFullPath("example.com/foo/bar/hoge/hoge.Status"))
			},
			want: []string{
				"Status",
				"HogeStatus",
				"Status",
				"HogeStatus",
				"Status_c79",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := renamer.New(tt.defaultOpts...)
			ch := make(chan string, 100)
			go func() {
				tt.caller(r, ch)
				close(ch)
			}()
			for _, want := range tt.want {
				got, ok := <-ch
				if !ok {
					t.Fatalf("the number of cases is <%d", len(tt.want))
				}
				if got != want {
					t.Errorf("got %q, want %q", got, want)
				}
			}

			got, ok := <-ch
			if ok {
				t.Fatalf("the number of cases is >%d(got %s)", len(tt.want), got)

				for range ch {
				}
			}
		})
	}
}
