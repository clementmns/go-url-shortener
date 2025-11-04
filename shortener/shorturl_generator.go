package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha2560f(str string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(str))
	return algorithm.Sum(nil)
}

func base58Encoder(bytes []byte) string {
	encoder := base58.BitcoinEncoding
	result, err := encoder.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(result)
}

func GenerateShortUrl(url string) string {
	urlHash := sha2560f(url)
	number := new(big.Int).SetBytes(urlHash).Uint64()
	result := base58Encoder([]byte(fmt.Sprintf("%d", number)))
	return result
}
