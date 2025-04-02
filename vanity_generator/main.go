package main

import (
	"encoding/hex"
	"fmt"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	creationCodeHex = "*"		// token byte code
	factoryAddr     = "*"		// factory contract address
	targetSuffix    = "8888"	// vanity addr tail
	threads         = 16		// threas num
)

func FormatSeedForSolidity(seed [32]byte) string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(seed[:]))
}

func worker(startSeed uint64, wg *sync.WaitGroup) {
	defer wg.Done()

	var seed [32]byte
	initCodeHash := crypto.Keccak256(common.FromHex(creationCodeHex))
	factoryAddress := common.HexToAddress(factoryAddr)

	for i := uint64(0); ; i++ {
		// fill seed
		for j := 0; j < 8; j++ {
			seed[31-j] = byte((startSeed + i) >> (8 * j))
		}

		// calc addr
		payload := append([]byte{0xff}, factoryAddress.Bytes()...)
		payload = append(payload, seed[:]...)
		payload = append(payload, initCodeHash...)
		address := common.BytesToAddress(crypto.Keccak256(payload)[12:])

		// check tail of address
		addrHex := address.Hex()[2:] // remove 0x
		if len(addrHex) >= 4 && addrHex[len(addrHex)-len(targetSuffix):] == targetSuffix {
			fmt.Println("seed[", FormatSeedForSolidity(seed), "] addr[", address.Hex(), "]")
		}

		if i % 100000 == 0{
			fmt.Println("i:", i)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// start multi worker
	for i := 0; i < threads; i++ {
		wg.Add(1)
		go worker(uint64(i)<<32, &wg) // each worker start on 2^23
	}

	fmt.Println("wait calc")
	
	wg.Wait()
}
