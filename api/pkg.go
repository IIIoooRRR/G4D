package api

import (
	"fmt"
	"unsafe"

	"github.com/IIIoooRRR/G4D/model/_const"
	"github.com/valyala/fasthttp"
)

/*
To reduce the number of allocations,
we’ve made it so that our array, which is stored on the stack, is treated as a string.
This is done to save memory and reduce the load on the garbage collector.
*/

type discordError struct {
	Code    int
	Message []byte
}

func (e *discordError) Error() string {
	return fmt.Sprintf("discord api resp: %d, %s", e.Code, e.Message)
}

//nolint:gosec
func GetURI(strings ...string) string {
	total := 0
	for _, str := range strings {
		total += len(str)
	}
	bytes := make([]byte, 0, total)
	for i := 0; i < len(strings); i++ {
		bytes = append(bytes, strings[i]...)
	}

	// #nosec G103
	return unsafe.String(unsafe.SliceData(bytes), len(bytes))
}

func GetURL(path, uri string) string {
	return path + uri
}

func (c *DiscordClient) doWithBody(method, url string, body []byte) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.Header.Set("Authorization", "Bot "+*c.token)
	req.SetBody(body)

	if err := c.client.Do(req, resp); err != nil {
		return 0, nil, err
	}
	dst := append([]byte(nil), resp.Body()...)
	code := resp.StatusCode()
	return code, dst, nil
}
func (c *DiscordClient) doWithoutBody(method, url string) (int, []byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod(method)
	req.SetRequestURI(url)
	req.Header.Set("Authorization", "Bot "+*c.token)

	if err := c.client.Do(req, resp); err != nil {
		return 0, nil, err
	}

	dst := append([]byte(nil), resp.Body()...)
	code := resp.StatusCode()
	return code, dst, nil
}

func (c *DiscordClient) delete(url string) (int, []byte, error) {
	return c.doWithoutBody("DELETE", url)
}
func (c *DiscordClient) get(url string) (int, []byte, error) {
	return c.doWithoutBody("GET", url)
}

func (c *DiscordClient) post(url string, body []byte) (int, []byte, error) {
	return c.doWithBody("POST", url, body)
}

func (c *DiscordClient) put(url string, body []byte) (int, []byte, error) {
	return c.doWithBody("PUT", url, body)
}

func (c *DiscordClient) patch(url string, body []byte) (int, []byte, error) {
	return c.doWithBody("PATCH", url, body)
}

func (c *DiscordClient) doRequest(method _const.Method, url string, body []byte) (int, []byte, error) {
	switch method {
	case _const.Get:
		return c.get(url)
	case _const.Post:
		return c.post(url, body)
	case _const.Delete:
		return c.delete(url)
	case _const.Put:
		return c.put(url, body)
	case _const.Patch:
		return c.patch(url, body)
	}
	return 0, nil, fmt.Errorf("unsupported method: %d", method)
}
