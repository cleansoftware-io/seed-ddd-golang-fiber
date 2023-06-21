package entities

type Image struct {
	url     string
	visible bool
}

func (i *Image) SetUrl(url string) {
	i.url = url
}

func (i *Image) SetVisible(visible bool) {
	i.visible = visible
}

func (i *Image) GetUrl() string {
	return i.url
}

func (i *Image) GetVisible() bool {
	return i.visible
}

func NewImage() *Image {
	return &Image{}
}
