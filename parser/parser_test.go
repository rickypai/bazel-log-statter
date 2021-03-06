package parser

import (
	"testing"
	"time"

	"github.com/rickypai/bazel-log-statter/bazel"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    *bazel.TargetResult
		wantErr error
	}{
		{
			"cached passed line",
			args{
				"//admin/server:go_default_test                                  (cached) PASSED in 0.3s",
			},
			&bazel.TargetResult{
				Name:      "//admin/server:go_default_test",
				Cached:    true,
				Status:    bazel.StatusPassed,
				Duration:  300 * time.Millisecond,
				Attempts:  1,
				Successes: 1,
			},
			nil,
		},
		{
			"no status line",
			args{
				"//summons/integration:go_default_test                                 NO STATUS",
			},
			&bazel.TargetResult{
				Name:   "//summons/integration:go_default_test",
				Status: bazel.StatusNoStatus,
			},
			nil,
		},
		{
			"uncached line",
			args{
				"//social-graph/worker:go_default_test                                    PASSED in 53.8s",
			},
			&bazel.TargetResult{
				Name:      "//social-graph/worker:go_default_test",
				Status:    bazel.StatusPassed,
				Duration:  53800 * time.Millisecond,
				Attempts:  1,
				Successes: 1,
			},
			nil,
		},
		{
			"flaky line",
			args{
				"//autobahn/stream:go_default_test                                         FLAKY, failed in 1 out of 2 in 13.5s",
			},
			&bazel.TargetResult{
				Name:      "//autobahn/stream:go_default_test",
				Status:    bazel.StatusFlaky,
				Duration:  13500 * time.Millisecond,
				Successes: 1,
				Attempts:  2,
			},
			nil,
		},
		{
			"failed line",
			args{
				"//subscription/worker-notification/consumer:go_default_test              FAILED in 30.9s",
			},
			&bazel.TargetResult{
				Name:      "//subscription/worker-notification/consumer:go_default_test",
				Status:    bazel.StatusFailed,
				Duration:  30900 * time.Millisecond,
				Attempts:  1,
				Successes: 0,
			},
			nil,
		},
		{
			"flaky failed line",
			args{
				"//media/consumer:go_default_test                                         FAILED in 3 out of 3 in 10.5s",
			},
			&bazel.TargetResult{
				Name:        "//media/consumer:go_default_test",
				Status:      bazel.StatusFailed,
				Duration:    10500 * time.Millisecond,
				CachedTimes: 0,
				Attempts:    3,
				Successes:   0,
			},
			nil,
		},
		{
			"flaky line with cached",
			args{
				"//autobahn/stream:go_default_test                            (1/2 cached) FLAKY, failed in 1 out of 2 in 14.9s",
			},
			&bazel.TargetResult{
				Name:        "//autobahn/stream:go_default_test",
				Status:      bazel.StatusFlaky,
				Duration:    14900 * time.Millisecond,
				CachedTimes: 1,
				Attempts:    2,
				Successes:   1,
			},
			nil,
		},
		{
			"timeout",
			args{
				"//social-graph/repos:go_default_test                                    TIMEOUT in 315.0s",
			},
			&bazel.TargetResult{
				Name:     "//social-graph/repos:go_default_test",
				Status:   bazel.StatusTimeout,
				Duration: 315000 * time.Millisecond,
				Attempts: 1,
			},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseLine(tt.args.line)

			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}

func Test_parseDuration(t *testing.T) {
	type args struct {
		durationStr string
	}
	tests := []struct {
		name    string
		args    args
		want    time.Duration
		wantErr error
	}{
		{
			"4.1",
			args{
				"4.1",
			},
			4100 * time.Millisecond,
			nil,
		},
		{
			"0.2",
			args{
				"0.2",
			},
			200 * time.Millisecond,
			nil,
		},
		{
			"0.0",
			args{
				"0.0",
			},
			0,
			nil,
		},
		{
			"315.0",
			args{
				"315.0",
			},
			315000 * time.Millisecond,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseDuration(tt.args.durationStr)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
