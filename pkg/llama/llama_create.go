package llama

// A new llama for the llama store.
type LlamaCreate struct {
	// The name of the llama. This must be unique across all llamas.
	Name *string `json:"name,omitempty" required:"true"`
	// The age of the llama in years.
	Age *int64 `json:"age,omitempty" required:"true"`
	// The color of a llama.
	Color *LlamaColor `json:"color,omitempty" required:"true"`
	// The rating of the llama from 1 to 5.
	Rating *int64 `json:"rating,omitempty" required:"true"`
}

func (l *LlamaCreate) SetName(name string) {
	l.Name = &name
}

func (l *LlamaCreate) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *LlamaCreate) SetAge(age int64) {
	l.Age = &age
}

func (l *LlamaCreate) GetAge() *int64 {
	if l == nil {
		return nil
	}
	return l.Age
}

func (l *LlamaCreate) SetColor(color LlamaColor) {
	l.Color = &color
}

func (l *LlamaCreate) GetColor() *LlamaColor {
	if l == nil {
		return nil
	}
	return l.Color
}

func (l *LlamaCreate) SetRating(rating int64) {
	l.Rating = &rating
}

func (l *LlamaCreate) GetRating() *int64 {
	if l == nil {
		return nil
	}
	return l.Rating
}
