package episode

type Data struct {
	Payload []Show `json:"payload,omitempty" bson:"name,omitempty" binding:"required"`
}

type Show struct {
	Drm          bool   `json:"drm,omitempty" bson:"name,omitempty" binding:"required"`
	Slug         string `json:"slug,omitempty" bson:"name,omitempty" binding:"required"`
	Image        Image  `json:"image,omitempty" bson:"name,omitempty" binding:"required"`
	Title        string `json:"title,omitempty" bson:"name,omitempty" binding:"required"`
	EpisodeCount int    `json:"episodeCount,omitempty" bson:"name,omitempty" binding:"required"`
}

type Image struct {
	ShowImage string `json:"showImage,omitempty" bson:"name,omitempty" binding:"required"`
}

func IsDRMEnabled(show Show) bool {
	return show.Drm && (show.EpisodeCount > 0)
}
