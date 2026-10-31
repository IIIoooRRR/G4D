package entity

type Attachment struct {
	Id     string `json:"id"`
	Fn     string `json:"filename"`
	Desc   string `json:"description,omitempty"`
	CT     string `json:"content_type,omitempty"`
	Sz     int    `json:"size"`
	Url    string `json:"url"`
	PURL   string `json:"proxy_url"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"` // only for picture
	Eph    bool   `json:"ephemeral,omitempty"`
	Dt     []byte `json:"-"`
}

func Attach(data []byte, filename string) *Attachment {
	return &Attachment{
		Dt: data,
		Fn: filename,
		Sz: len(data),
	}
}

func (a *Attachment) ID(id string) *Attachment          { a.Id = id; return a }
func (a *Attachment) Filename(name string) *Attachment  { a.Fn = name; return a }
func (a *Attachment) Description(d string) *Attachment  { a.Desc = d; return a }
func (a *Attachment) ContentType(ct string) *Attachment { a.CT = ct; return a }
func (a *Attachment) Size(s int) *Attachment            { a.Sz = s; return a }
func (a *Attachment) URL(u string) *Attachment          { a.Url = u; return a }
func (a *Attachment) ProxyURL(u string) *Attachment     { a.PURL = u; return a }
func (a *Attachment) Dimensions(w, h int) *Attachment   { a.Width = w; a.Height = h; return a }
func (a *Attachment) Ephemeral(e bool) *Attachment      { a.Eph = e; return a }
func (a *Attachment) Data(d []byte) *Attachment         { a.Dt = d; a.Sz = len(d); return a }
