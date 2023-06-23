package entities

type ImageList struct {
	images []Image
}

func NewImageList() *ImageList {
	return &ImageList{}
}

func (il *ImageList) AddImage(image Image) {
	il.images = append(il.images, image)
}

func (il *ImageList) GetImages() []Image {
	return il.images
}

func (il *ImageList) GetVisibleImages() []Image {
	var visibleImages []Image
	for _, image := range il.images {
		if image.visible {
			visibleImages = append(visibleImages, image)
		}
	}
	return visibleImages
}

func (il *ImageList) RemoveImage(image Image) {
	var newImages []Image
	for _, i := range il.images {
		if i != image {
			newImages = append(newImages, i)
		}
	}
	il.images = newImages
}

func (il *ImageList) ActivateImage(image Image) {
	for i, img := range il.images {
		if img == image {
			il.images[i].visible = true
		}
	}
}

func (il *ImageList) DeactivateImage(image Image) {
	for i, img := range il.images {
		if img == image {
			il.images[i].visible = false
		}
	}
}

func (il *ImageList) GetActiveImages() []Image {
	var activeImages []Image
	for _, image := range il.images {
		if image.visible {
			activeImages = append(activeImages, image)
		}
	}
	return activeImages
}

func (il *ImageList) GetInactiveImages() []Image {
	var inactiveImages []Image
	for _, image := range il.images {
		if !image.visible {
			inactiveImages = append(inactiveImages, image)
		}
	}
	return inactiveImages
}
