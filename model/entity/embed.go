package entity

import (
	"strconv"
	"strings"
	"time"
)

type Embed struct {
	Title       string     `json:"title,omitempty"`
	Type        string     `json:"type,omitempty"`
	Description string     `json:"description,omitempty"`
	URL         string     `json:"url,omitempty"`
	Timestamp   string     `json:"timestamp,omitempty"`
	Color       int        `json:"color,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Video       *Video     `json:"video,omitempty"`
	Provider    *Provider  `json:"provider,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	Fields      []Field    `json:"fields,omitempty"`
}

type Footer struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type Image struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type Thumbnail struct {
	URL      string `json:"url"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type Video struct {
	URL      string `json:"url,omitempty"`
	ProxyURL string `json:"proxy_url,omitempty"`
	Height   int    `json:"height,omitempty"`
	Width    int    `json:"width,omitempty"`
}

type Provider struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type Author struct {
	Name         string `json:"name"`
	URL          string `json:"url,omitempty"`
	IconURL      string `json:"icon_url,omitempty"`
	ProxyIconURL string `json:"proxy_icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  any    `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

func NewEmbed() *Embed {
	return &Embed{}
}

func (e *Embed) SetTitle(title string) *Embed {
	e.Title = title
	return e
}

func (e *Embed) SetDescription(desc string) *Embed {
	e.Description = desc
	return e
}

func (e *Embed) SetURL(url string) *Embed {
	e.URL = url
	return e
}

func (e *Embed) SetColor(color string) *Embed {
	color = strings.ReplaceAll(color, "#", "")
	result, err := strconv.ParseInt(color, 16, 64)
	if err != nil {
		return e
	}
	e.Color = int(result)
	return e
}

func (e *Embed) SetTimestamp(ts string) *Embed {
	e.Timestamp = ts
	return e
}

func (e *Embed) SetTimestampNow() *Embed {
	e.Timestamp = time.Now().Format(time.RFC3339)
	return e
}

func (e *Embed) SetFooter(footer *Footer) *Embed {
	e.Footer = footer
	return e
}

func (e *Embed) AddFooter(text, iconURL string) *Embed {
	e.Footer = &Footer{Text: text, IconURL: iconURL}
	return e
}

func (e *Embed) SetImage(image *Image) *Embed {
	e.Image = image
	return e
}

func (e *Embed) AddImage(url string) *Embed {
	e.Image = &Image{URL: url}
	return e
}

func (e *Embed) SetThumbnail(thumbnail *Thumbnail) *Embed {
	e.Thumbnail = thumbnail
	return e
}

func (e *Embed) AddThumbnail(url string) *Embed {
	e.Thumbnail = &Thumbnail{URL: url}
	return e
}

func (e *Embed) SetVideo(video *Video) *Embed {
	e.Video = video
	return e
}

func (e *Embed) SetProvider(provider *Provider) *Embed {
	e.Provider = provider
	return e
}

func (e *Embed) SetAuthor(author *Author) *Embed {
	e.Author = author
	return e
}

func (e *Embed) AddAuthor(name, iconURL, url string) *Embed {
	e.Author = &Author{Name: name, IconURL: iconURL, URL: url}
	return e
}

func (e *Embed) AddField(name string, value any, inline bool) *Embed {
	e.Fields = append(e.Fields, Field{Name: name, Value: value, Inline: inline})
	return e
}

func (e *Embed) AddFields(fields ...Field) *Embed {
	e.Fields = append(e.Fields, fields...)
	return e
}

func (e *Embed) ClearFields() *Embed {
	e.Fields = nil
	return e
}

/* Footer */
func NewFooter(text string) *Footer {
	return &Footer{Text: text}
}

func (f *Footer) SetText(text string) *Footer {
	f.Text = text
	return f
}

func (f *Footer) SetIconURL(url string) *Footer {
	f.IconURL = url
	return f
}

/* Image */
func NewImage(url string) *Image {
	return &Image{URL: url}
}

func (i *Image) SetURL(url string) *Image {
	i.URL = url
	return i
}

func (i *Image) SetSize(width, height int) *Image {
	i.Width = width
	i.Height = height
	return i
}

/* Thumbnail */
func NewThumbnail(url string) *Thumbnail {
	return &Thumbnail{URL: url}
}

func (t *Thumbnail) SetURL(url string) *Thumbnail {
	t.URL = url
	return t
}

func (t *Thumbnail) SetSize(width, height int) *Thumbnail {
	t.Width = width
	t.Height = height
	return t
}

/* Video */
func NewVideo(url string) *Video {
	return &Video{URL: url}
}

func (v *Video) SetURL(url string) *Video {
	v.URL = url
	return v
}

func (v *Video) SetSize(width, height int) *Video {
	v.Width = width
	v.Height = height
	return v
}

/* Provider */
func NewProvider(name, url string) *Provider {
	return &Provider{Name: name, URL: url}
}

func (p *Provider) SetName(name string) *Provider {
	p.Name = name
	return p
}

func (p *Provider) SetURL(url string) *Provider {
	p.URL = url
	return p
}

/* Author */
func NewAuthor(name string) *Author {
	return &Author{Name: name}
}

func (a *Author) SetName(name string) *Author {
	a.Name = name
	return a
}

func (a *Author) SetURL(url string) *Author {
	a.URL = url
	return a
}

func (a *Author) SetIconURL(url string) *Author {
	a.IconURL = url
	return a
}

/* Field */
func NewField(name string, value any) *Field {
	return &Field{Name: name, Value: value}
}

func (f *Field) SetName(name string) *Field {
	f.Name = name
	return f
}

func (f *Field) SetValue(value any) *Field {
	f.Value = value
	return f
}

func (f *Field) SetInline(inline bool) *Field {
	f.Inline = inline
	return f
}
