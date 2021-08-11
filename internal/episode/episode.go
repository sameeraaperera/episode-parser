package episode

type Data struct {
	Payload []Show `json:"payload,omitempty" binding:"required"`
}

type Show struct {
	Drm          bool   `json:"drm,omitempty" binding:"required"`
	Slug         string `json:"slug,omitempty" binding:"required"`
	Image        Image  `json:"image,omitempty" binding:"required"`
	Title        string `json:"title,omitempty"  binding:"required"`
	EpisodeCount int    `json:"episodeCount,omitempty" binding:"required"`
}

type Image struct {
	ShowImage string `json:"showImage,omitempty" binding:"required"`
}

func IsDRMEnabled(show Show) bool {
	return show.Drm && (show.EpisodeCount > 0)
}
