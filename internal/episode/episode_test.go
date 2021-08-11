package episode_test

import (
	"testing"

	"github.com/sap4001/episode-parser/internal/episode"
	"github.com/stretchr/testify/assert"
)

func Test_IsDRMEnabledTrue(t *testing.T) {
	show := episode.Show{
		Drm:          true,
		EpisodeCount: 1,
	}
	got := episode.IsDRMEnabled(show)
	assert.True(t, got)
}

func Test_IsDRMEnabledFalse(t *testing.T) {
	show1 := episode.Show{
		Drm:          false,
		EpisodeCount: 1,
	}
	got := episode.IsDRMEnabled(show1)
	assert.False(t, got)

	show2 := episode.Show{
		Drm:          true,
		EpisodeCount: 0,
	}
	got = episode.IsDRMEnabled(show2)
	assert.False(t, got)
}
