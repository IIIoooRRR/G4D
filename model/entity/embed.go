package entity

import (
	"strconv"
	"strings"
	"time"
)

type Embed struct {
	Ttl    string     `json:"title,omitempty"`
	Typ    string     `json:"type,omitempty"`
	Desc   string     `json:"description,omitempty"`
	Url    string     `json:"url,omitempty"`
	Tstamp string     `json:"timestamp,omitempty"`
	Clr    int        `json:"color,omitempty"`
	Ftr    *footer    `json:"footer,omitempty"`
	Img    *image     `json:"image,omitempty"`
	Thmb   *thumbnail `json:"thumbnail,omitempty"`
	Vid    *video     `json:"video,omitempty"`
	Prv    *provider  `json:"provider,omitempty"`
	Auth   *author    `json:"author,omitempty"`
	Flds   []field    `json:"fields,omitempty"`
}

func NewEmbed() *Embed { return &Embed{} }

func (e *Embed) Title(t string) *Embed       { e.Ttl = t; return e }
func (e *Embed) Description(d string) *Embed { e.Desc = d; return e }
func (e *Embed) URL(u string) *Embed         { e.Url = u; return e }
func (e *Embed) Color(hex string) *Embed {
	hex = strings.ReplaceAll(hex, "#", "")
	if v, err := strconv.ParseInt(hex, 16, 64); err == nil {
		e.Clr = int(v)
	}
	return e
}
func (e *Embed) Timestamp(ts string) *Embed { e.Tstamp = ts; return e }
func (e *Embed) TimestampNow() *Embed       { e.Tstamp = time.Now().Format(time.RFC3339); return e }
func (e *Embed) Footer(f *footer) *Embed    { e.Ftr = f; return e }
func (e *Embed) NewFooter(text, iconURL string) *Embed {
	e.Ftr = &footer{Txt: text, IconUrl: iconURL}
	return e
}
func (e *Embed) Image(i *image) *Embed      { e.Img = i; return e }
func (e *Embed) NewImage(url string) *Embed { e.Img = &image{Url: url}; return e }

func (e *Embed) Thumbnail(t *thumbnail) *Embed  { e.Thmb = t; return e }
func (e *Embed) NewThumbnail(url string) *Embed { e.Thmb = &thumbnail{Url: url}; return e }
func (e *Embed) Video(v *video) *Embed          { e.Vid = v; return e }
func (e *Embed) Provider(p *provider) *Embed    { e.Prv = p; return e }
func (e *Embed) Author(a *author) *Embed        { e.Auth = a; return e }

func (e *Embed) NewAuthor(name, iconURL, url string) *Embed {
	e.Auth = &author{Nm: name, IconUrl: iconURL, Url: url}
	return e
}

func (e *Embed) Field(name string, value any, inline bool) *Embed {
	e.Flds = append(e.Flds, field{Nm: name, Val: value, Inl: inline})
	return e
}
func (e *Embed) Fields(fs ...field) *Embed { e.Flds = append(e.Flds, fs...); return e }
func (e *Embed) ClearFields() *Embed       { e.Flds = nil; return e }

// footer
type footer struct {
	Txt       string `json:"text"`
	IconUrl   string `json:"icon_url,omitempty"`
	ProxyIcon string `json:"proxy_icon_url,omitempty"`
}

func Footer(text string) *footer           { return &footer{Txt: text} }
func (f *footer) Text(t string) *footer    { f.Txt = t; return f }
func (f *footer) IconURL(u string) *footer { f.IconUrl = u; return f }

// image
type image struct {
	Url    string `json:"url"`
	PURL   string `json:"proxy_url,omitempty"`
	Height int    `json:"height,omitempty"`
	Wigth  int    `json:"width,omitempty"`
}

func Image(url string) *image         { return &image{Url: url} }
func (i *image) URL(u string) *image  { i.Url = u; return i }
func (i *image) Size(w, h int) *image { i.Wigth = w; i.Height = h; return i }

// thumbnail
type thumbnail struct {
	Url    string `json:"url"`
	PURL   string `json:"proxy_url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

func Thumbnail(url string) *thumbnail         { return &thumbnail{Url: url} }
func (t *thumbnail) URL(u string) *thumbnail  { t.Url = u; return t }
func (t *thumbnail) Size(w, h int) *thumbnail { t.Width = w; t.Height = h; return t }

// video
type video struct {
	Url    string `json:"url,omitempty"`
	PURL   string `json:"proxy_url,omitempty"`
	Height int    `json:"height,omitempty"`
	Width  int    `json:"width,omitempty"`
}

func Video(url string) *video         { return &video{Url: url} }
func (v *video) URL(u string) *video  { v.Url = u; return v }
func (v *video) Size(w, h int) *video { v.Width = w; v.Height = h; return v }

// provider
type provider struct {
	Nm  string `json:"name,omitempty"`
	Url string `json:"url,omitempty"`
}

func Provider(name, url string) *provider   { return &provider{Nm: name, Url: url} }
func (p *provider) Name(n string) *provider { p.Nm = n; return p }
func (p *provider) URL(u string) *provider  { p.Url = u; return p }

// author
type author struct {
	Nm        string `json:"name"`
	Url       string `json:"url,omitempty"`
	IconUrl   string `json:"icon_url,omitempty"`
	ProxyIcon string `json:"proxy_icon_url,omitempty"`
}

func Author(name string) *author           { return &author{Nm: name} }
func (a *author) Name(n string) *author    { a.Nm = n; return a }
func (a *author) URL(u string) *author     { a.Url = u; return a }
func (a *author) IconURL(u string) *author { a.IconUrl = u; return a }

// field
type field struct {
	Nm  string `json:"name"`
	Val any    `json:"value"`
	Inl bool   `json:"inline,omitempty"`
}

func Field(name string, value any) *field { return &field{Nm: name, Val: value} }
func (f *field) Name(n string) *field     { f.Nm = n; return f }
func (f *field) Value(v any) *field       { f.Val = v; return f }
func (f *field) Inline(i bool) *field     { f.Inl = i; return f }
