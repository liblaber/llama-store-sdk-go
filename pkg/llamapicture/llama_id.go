package llamapicture

// A llama id.
type LlamaId struct {
	// The ID of the llama.
	Id *int64 `json:"id,omitempty" required:"true"`
}

func (l *LlamaId) SetId(id int64) {
	l.Id = &id
}

func (l *LlamaId) GetId() *int64 {
	if l == nil {
		return nil
	}
	return l.Id
}
