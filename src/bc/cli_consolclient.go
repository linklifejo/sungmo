package bc

import (
	"fmt"

	"github.com/eiannone/keyboard"
	//	"time"
)

func (cli *CLI) ConsolClient() {

	for {

		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer func() {
			_ = keyboard.Close()
		}()

		CallClear()
		fmt.Println("1. 지갑 생성")
		fmt.Println("2. 지갑 확인")
		fmt.Println("3. 코인 송금")
		fmt.Println("4. 송금 확인")
		fmt.Println("5. 노드 생성")
		fmt.Println("6. 블럭체인 생성")
		fmt.Println("Esc 종료")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			myWalletCreate()
		case "2":
			myWalletLoad()
		case "3":
			myCoinSend()
		case "4":
			myCoinSendAfter()
		case "5":
			nodeCreate()
		case "6":
			blockchainCreate()

		}
		if key == keyboard.KeyEsc {
			break
		}
	}

}

func myWalletCreate() {

	var c = CLI{}
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	CallClear()
	for {
		fmt.Println("1. 3000")
		fmt.Println("2. 3001")
		fmt.Println("3. 3002")
		fmt.Println("4. 3003")
		fmt.Println("5. 3004")
		fmt.Println("6. 3005")
		fmt.Println("7. 3006")
		fmt.Println("8. 3007")
		fmt.Println("9. 3008")
		fmt.Println("Esc 이전화면")
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			c.createWallet("3000")
		case "2":
			c.createWallet("3001")
		case "3":
			c.createWallet("3002")
		case "4":
			c.createWallet("3003")
		case "5":
			c.createWallet("3004")
		case "6":
			c.createWallet("3005")
		case "7":
			c.createWallet("3006")
		case "8":
			c.createWallet("3007")
		case "9":
			c.createWallet("3008")

		}
		if key == keyboard.KeyEsc {
			break
		}

	}
}
func blockchainCreate() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var c = CLI{}

	CallClear()
	var nodeID = "3000"
	data := c.walletAddresses(nodeID)
	for idx, address := range data {
		fmt.Printf("%d. address %s\r\n", idx+1, address)
	}
	char, key, err := keyboard.GetKey()
	if err != nil {
		panic(err)
	}

	fmt.Println("Esc 이전화면")
	switch string(char) {
	case "1":
		c.createBlockchain(data[0], nodeID)
	case "2":
		c.createBlockchain(data[1], nodeID)
	case "3":
		c.createBlockchain(data[2], nodeID)
	case "4":
		c.createBlockchain(data[3], nodeID)
	case "5":
		c.createBlockchain(data[4], nodeID)
	}
	fmt.Println("블럭체인 생성 처리")
	var from, to string
	from = "blockchain_3000.db"
	to = "blockchain_genesis.db"

	err = CopyFile(from, to)
	if err != nil {
		fmt.Printf("CopyFile failed %q\n", err)
	} else {
		fmt.Printf("blockchain_genesis.db CopyFile succeeded\n")
	}
	if key == keyboard.KeyEsc {
		return
	}
}

func myWalletLoad() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var c = CLI{}

	CallClear()
	for {
		fmt.Println("1. 3000")
		fmt.Println("2. 3001")
		fmt.Println("3. 3002")
		fmt.Println("4. 3003")
		fmt.Println("5. 3004")
		fmt.Println("6. 3005")
		fmt.Println("7. 3006")
		fmt.Println("8. 3007")
		fmt.Println("9. 3008")
		fmt.Println("Esc 이전화면")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			c.listAddresses("3000")
		case "2":
			c.listAddresses("3001")
		case "3":
			c.listAddresses("3002")
		case "4":
			c.listAddresses("3003")
		case "5":
			c.listAddresses("3004")
		case "6":
			c.listAddresses("3005")
		case "7":
			c.listAddresses("3006")
		case "8":
			c.listAddresses("3007")
		case "9":
			c.listAddresses("3008")

		}

		if key == keyboard.KeyEsc {
			break
		}

	}
}
func myCoinSendAfter() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		CallClear()

		fmt.Println("1. 코인 송금")
		fmt.Println("2. 송금 확인")
		fmt.Println("3. 노드실행")
		fmt.Println("Esc 이전화면")
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			//clientDownload()
		case "2":
			//clientExc()
		case "3":
			//clientDownload()
		}
		if key == keyboard.KeyEsc {
			break
		}

	}
}
func manageServerCreate() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var c = CLI{}

	CallClear()
	for {
		fmt.Println("1. 3000")
		fmt.Println("2. 3001")
		fmt.Println("3. 3002")
		fmt.Println("4. 3003")
		fmt.Println("5. 3004")
		fmt.Println("6. 3006")
		fmt.Println("7. 3007")
		fmt.Println("8. 3008")
		fmt.Println("9. 3009")
		fmt.Println("Esc 이전화면")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			c.listAddresses("3000")
		case "2":
			c.listAddresses("3001")
		case "3":
			c.listAddresses("3002")
		case "4":
			c.listAddresses("3003")
		case "5":
			c.listAddresses("3004")
		case "6":
			c.listAddresses("3005")
		case "7":
			c.listAddresses("3006")
		case "8":
			c.listAddresses("3007")
		case "9":
			c.listAddresses("3008")

		}

		if key == keyboard.KeyEsc {
			break
		}

	}
}

func seedServerCreate() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var c = CLI{}

	CallClear()
	for {
		fmt.Println("1. 3000")
		fmt.Println("2. 3001")
		fmt.Println("3. 3002")
		fmt.Println("4. 3003")
		fmt.Println("5. 3004")
		fmt.Println("6. 3005")
		fmt.Println("7. 3006")
		fmt.Println("8. 3007")
		fmt.Println("9. 3008")
		fmt.Println("Esc 이전화면")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			c.listAddresses("3000")
		case "2":
			c.listAddresses("3001")
		case "3":
			c.listAddresses("3002")
		case "4":
			c.listAddresses("3003")
		case "5":
			c.listAddresses("3004")
		case "6":
			c.listAddresses("3005")
		case "7":
			c.listAddresses("3006")
		case "8":
			c.listAddresses("3007")
		case "9":
			c.listAddresses("3008")

		}

		if key == keyboard.KeyEsc {
			break
		}

	}
}
func downloadServerCreate() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var c = CLI{}

	CallClear()
	for {
		fmt.Println("1. 3000")
		fmt.Println("2. 3001")
		fmt.Println("3. 3002")
		fmt.Println("4. 3003")
		fmt.Println("5. 3004")
		fmt.Println("6. 3005")
		fmt.Println("7. 3006")
		fmt.Println("8. 3007")
		fmt.Println("9. 3008")
		fmt.Println("Esc 이전화면")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			c.listAddresses("3000")
		case "2":
			c.listAddresses("3001")
		case "3":
			c.listAddresses("3002")
		case "4":
			c.listAddresses("3003")
		case "5":
			c.listAddresses("3004")
		case "6":
			c.listAddresses("3005")
		case "7":
			c.listAddresses("3006")
		case "8":
			c.listAddresses("3007")
		case "9":
			c.listAddresses("3008")

		}

		if key == keyboard.KeyEsc {
			break
		}

	}
}
func fullNodeCreate() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var c = CLI{}

	CallClear()
	for {
		fmt.Println("1. 3000")
		fmt.Println("2. 3001")
		fmt.Println("3. 3002")
		fmt.Println("4. 3003")
		fmt.Println("5. 3004")
		fmt.Println("6. 3005")
		fmt.Println("7. 3006")
		fmt.Println("8. 3007")
		fmt.Println("9. 3008")
		fmt.Println("Esc 이전화면")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			c.listAddresses("3000")
		case "2":
			c.listAddresses("3001")
		case "3":
			c.listAddresses("3002")
		case "4":
			c.listAddresses("3003")
		case "5":
			c.listAddresses("3004")
		case "6":
			c.listAddresses("3005")
		case "7":
			c.listAddresses("3006")
		case "8":
			c.listAddresses("3007")
		case "9":
			c.listAddresses("3008")

		}

		if key == keyboard.KeyEsc {
			break
		}

	}
}
func walletNodeCreate() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var c = CLI{}

	CallClear()
	for {
		fmt.Println("1. 3000")
		fmt.Println("2. 3001")
		fmt.Println("3. 3002")
		fmt.Println("4. 3003")
		fmt.Println("5. 3004")
		fmt.Println("6. 3005")
		fmt.Println("7. 3006")
		fmt.Println("8. 3007")
		fmt.Println("9. 3008")
		fmt.Println("Esc 이전화면")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			c.listAddresses("3000")
		case "2":
			c.listAddresses("3001")
		case "3":
			c.listAddresses("3002")
		case "4":
			c.listAddresses("3003")
		case "5":
			c.listAddresses("3004")
		case "6":
			c.listAddresses("3005")
		case "7":
			c.listAddresses("3006")
		case "8":
			c.listAddresses("3007")
		case "9":
			c.listAddresses("3008")

		}

		if key == keyboard.KeyEsc {
			break
		}

	}
}
func miningNodeCreate() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	var c = CLI{}

	CallClear()
	for {
		fmt.Println("1. 3000")
		fmt.Println("2. 3001")
		fmt.Println("3. 3002")
		fmt.Println("4. 3003")
		fmt.Println("5. 3004")
		fmt.Println("6. 3005")
		fmt.Println("7. 3006")
		fmt.Println("8. 3007")
		fmt.Println("9. 3008")

		fmt.Println("Esc 이전화면")

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			c.listAddresses("3000")
		case "2":
			c.listAddresses("3001")
		case "3":
			c.listAddresses("3002")
		case "4":
			c.listAddresses("3003")
		case "5":
			c.listAddresses("3004")
		case "6":
			c.listAddresses("3005")
		case "7":
			c.listAddresses("3006")
		case "8":
			c.listAddresses("3007")
		case "9":
			c.listAddresses("3008")

		}

		if key == keyboard.KeyEsc {
			break
		}

	}
}

func nodeCreate() {

	for {
		if err := keyboard.Open(); err != nil {
			panic(err)
		}
		defer func() {
			_ = keyboard.Close()
		}()

		CallClear()

		fmt.Println("1. 시드서버    생성")
		fmt.Println("2. 관리서버    생성")
		fmt.Println("3. 다운로드서버 생성")
		fmt.Println("4. 풀노드      생성")
		fmt.Println("5. 지갑노드    생성")
		fmt.Println("6. 채굴노드    생성")
		fmt.Println("Esc 이전화면")
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			seedServerCreate()
		case "2":
			manageServerCreate()
		case "3":
			downloadServerCreate()
		case "4":
			fullNodeCreate()
		case "5":
			walletNodeCreate()
		case "6":
			miningNodeCreate()
		}
		if key == keyboard.KeyEsc {
			break
		}

	}
}

func myCoinSend() {
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		CallClear()

		fmt.Println("1. 코인 송금")
		fmt.Println("2. 송금 확인")
		fmt.Println("3. 노드실행")
		fmt.Println("Esc 이전화면")
		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}

		switch string(char) {
		case "1":
			//clientDownload()
		case "2":
			//clientExc()
		case "3":
			//clientDownload()
		}
		if key == keyboard.KeyEsc {
			break
		}

	}
}
