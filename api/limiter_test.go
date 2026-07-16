package api_test

import (
	"context"
	"io"
	"net/http"
	"testing"

	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(rate.Limit(5), 1)

func TestClientWithLimiter_DoRequest(t *testing.T) {
	urls := []string{
		"https://cdn.discordapp.com/attachments/1518636444812443659/1518898080466862090/Aki.webp?ex=6a3b9792&is=6a3a4612&hm=1320b9e633592edf5376ff8ae4a0fc047eba18fc910db1466f74f5ed412ffbbc&",
		"https://cdn.discordapp.com/attachments/1518636290583691428/1518898080508805190/Asa.webp?ex=6a3b9792&is=6a3a4612&hm=a4dd99448f08cd8c76ffda8d994b3f464387bb8f2a13cc19b4550c131ddc0a82&",
		"https://cdn.discordapp.com/attachments/1518636480581337189/1518898081238880266/Chainsaws_demon.webp?ex=6a3b9792&is=6a3a4612&hm=9b81a1ba126a72d1443cd38a62950d777e0f10386684044065a109883f8e4e86&",
		"https://cdn.discordapp.com/attachments/1518636268500816013/1518898082044186734/Ghost_devil.webp?ex=6a3b9792&is=6a3a4612&hm=7d162d50b9d3acb7fc78b82bc62f69744b6e383375f0aa340391d7dad807614c&",
		"https://cdn.discordapp.com/attachments/1518636444812443659/1518898084199927898/Angel.webp?ex=6a3b9793&is=6a3a4613&hm=9c21b709b062147af19cd8d1ae5f1c6c902ce4232beb6c31926f28d89c138402&",
		"https://cdn.discordapp.com/attachments/1518636351178936454/1518898085109956648/Aging_devil.webp?ex=6a3b9793&is=6a3a4613&hm=565ff1b47e23fe947e42c9d9fca0f86cb9b1298b4f6f2272cfe6e0b6af9a530e&",
		"https://cdn.discordapp.com/attachments/1518636290583691428/1518898085428859021/Bat_devil.webp?ex=6a3b9793&is=6a3a4613&hm=526a0a0b7a6a92a8068ea99ed818350dd4e66ed940033d490692b007f6b72035&",
		"https://cdn.discordapp.com/attachments/1518636480581337189/1518898087035273348/Death_demon.webp?ex=6a3b9793&is=6a3a4613&hm=ba2784e68ff7ada0aabf193ed18ca06d90cf57e036ca06b44b8c21952ae2a88e&",
		"https://cdn.discordapp.com/attachments/1518636268500816013/1518898090436988948/Nayuta.webp?ex=6a3b9794&is=6a3a4614&hm=2313a81f3ba9690d6bdbfcd6de13559f9ebdae177f6ec423a9de9c81669a52ee&",
		"https://cdn.discordapp.com/attachments/1518636290583691428/1518898092047335534/Bucky.webp?ex=6a3b9794&is=6a3a4614&hm=251d29ddd6c75ed93236c77a3dfcd63b9513a80c4c3dcea88245fd436d8d4019&",
		"https://cdn.discordapp.com/attachments/1518636351178936454/1518898092940984430/Beam.webp?ex=6a3b9795&is=6a3a4615&hm=246718b6586c0fe47aaa61116767f475aa27e563a8f033a30cc8512f00a048af&",
	}
	for _, url := range urls {
		body, err := doRequest(url)
		if err != nil {
			t.Error(err)
		}
		t.Logf("Response size: %d bytes", len(body))
	}
}

func doRequest(url string) ([]byte, error) {
	if err := limiter.Wait(context.Background()); err != nil {
		return nil, err
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
