package llama

// A llama, with details of its name, age, color, and rating from 1 to 5.
type Llama struct {
	// The name of the llama. This must be unique across all llamas.
	Name *string `json:"name,omitempty" required:"true"`
	// The age of the llama in years.
	Age *int64 `json:"age,omitempty" required:"true"`
	// The color of a llama.
	Color *LlamaColor `json:"color,omitempty" required:"true"`
	// The rating of the llama from 1 to 5.
	Rating *int64 `json:"rating,omitempty" required:"true"`
	// The ID of the llama.
	Id *int64 `json:"id,omitempty" required:"true"`
}

func (l *Llama) SetName(name string) {
	l.Name = &name
}

func (l *Llama) GetName() *string {
	if l == nil {
		return nil
	}
	return l.Name
}

func (l *Llama) SetAge(age int64) {
	l.Age = &age
}

func (l *Llama) GetAge() *int64 {
	if l == nil {
		return nil
	}
	return l.Age
}

func (l *Llama) SetColor(color LlamaColor) {
	l.Color = &color
}

func (l *Llama) GetColor() *LlamaColor {
	if l == nil {
		return nil
	}
	return l.Color
}

func (l *Llama) SetRating(rating int64) {
	l.Rating = &rating
}

func (l *Llama) GetRating() *int64 {
	if l == nil {
		return nil
	}
	return l.Rating
}

func (l *Llama) SetId(id int64) {
	l.Id = &id
}

func (l *Llama) GetId() *int64 {
	if l == nil {
		return nil
	}
	return l.Id
}
