package entity

type Attachment struct {
	ID          string `json:"id"`
	Filename    string `json:"filename"`
	Description string `json:"description,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Size        int    `json:"size"`
	URL         string `json:"url"`
	ProxyURL    string `json:"proxy_url"`
	Height      int    `json:"height,omitempty"`
	Width       int    `json:"width,omitempty"` // only for picture
	Ephemeral   bool   `json:"ephemeral,omitempty"`
	Data        []byte `json:"-"`
}

func NewAttachment(data []byte, filename string) *Attachment {
	return &Attachment{
		Data:     data,
		Filename: filename,
		Size:     len(data),
	}
}

func (a *Attachment) SetID(id string) *Attachment {
	a.ID = id
	return a
}

func (a *Attachment) SetFilename(filename string) *Attachment {
	a.Filename = filename
	return a
}

func (a *Attachment) SetDescription(desc string) *Attachment {
	a.Description = desc
	return a
}

func (a *Attachment) SetContentType(contentType string) *Attachment {
	a.ContentType = contentType
	return a
}

func (a *Attachment) SetSize(size int) *Attachment {
	a.Size = size
	return a
}

func (a *Attachment) SetURL(url string) *Attachment {
	a.URL = url
	return a
}

func (a *Attachment) SetProxyURL(proxyURL string) *Attachment {
	a.ProxyURL = proxyURL
	return a
}

func (a *Attachment) SetDimensions(width, height int) *Attachment {
	a.Width = width
	a.Height = height
	return a
}

func (a *Attachment) SetEphemeral(ephemeral bool) *Attachment {
	a.Ephemeral = ephemeral
	return a
}

func (a *Attachment) SetData(data []byte) *Attachment {
	a.Data = data
	a.Size = len(data)
	return a
}
