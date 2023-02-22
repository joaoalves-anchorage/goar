package example

import (
	"github.com/everFinance/goar"
	"github.com/everFinance/goar/types"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func Test_SendData(t *testing.T) {
	arNode := "https://arweave.net"
	w, err := goar.NewWalletFromPath("./wallet/account1.json", arNode) // your wallet private key
	assert.NoError(t, err)

	data, err := ioutil.ReadFile("/Users/local/Downloads/abc.jpeg") // local file path
	if err != nil {
		panic(err)
	}
	tags := []types.Tag{
		{Name: "xxxx", Value: "sssss"},
		{Name: "yyyyyy", Value: "kkkkkk"},
	}
	tx, err := w.SendDataSpeedUp(data, tags, 10)
	assert.NoError(t, err)
	t.Logf("tx hash: %s", tx.ID)
}

func Test_LoadData(t *testing.T) {
	arCli := goar.NewClient("https://arweave.net")

	arId := "r90Z_PuhD-louq6uzLTI-xWMfB5TzIti30o7QvW-6A4"
	data, err := arCli.GetTransactionData(arId)
	assert.NoError(t, err)
	t.Log(len(data))
}

func TestSendDataStream(t *testing.T) {

	arNode := "https://arweave.net"
	w, err := goar.NewWalletFromPath("./testKey.json", arNode) // your wallet private key
	assert.NoError(t, err)
	data, err := os.Open("img.jpeg") // local file path
	defer data.Close()
	if err != nil {
		panic(err)
	}
	tags := []types.Tag{
		{Name: "Content-Type", Value: "img/jpeg"},
		{Name: "test", Value: "kevin-test"},
	}
	tx, err := w.SendDataStreamSpeedUp(data, tags, 10)
	assert.NoError(t, err)
	t.Log(tx.ID)
	// test arId: k5IgHLTag_3bB6Sp5tTUhrFrPPvU5MjevV468dfxNKk
}
