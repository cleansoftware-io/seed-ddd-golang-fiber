package entities

type Description struct {
	value string
}

func (d *Description) SetValue(value string) {
	d.value = value
}

func (d *Description) GetValue() string {
	return d.value
}

func NewDescription() *Description {
	return &Description{}
}
