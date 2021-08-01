//+build ignore
package bc

import (
	"fmt"
	"strconv"

	"github.com/gonutz/wui/v2"
)

var leftport, rightport string
var walletadd []string
var blockchainaddress string
var fromaddress, toaddress, confirmaddress string
var nodeExe []string

func (cli *CLI) WinClient() {
	cc := CLI{}
	winFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -16,
	})

	win := wui.NewWindow()
	win.SetFont(winFont)
	win.SetInnerSize(734, 194)
	win.SetTitle("블럭체인")

	panel1Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -16,
	})

	panel1 := wui.NewPanel()
	panel1.SetFont(panel1Font)
	panel1.SetBounds(4, 7, 724, 180)
	panel1.SetBorderStyle(wui.PanelBorderSingleLine)
	win.Add(panel1)
	initBut := wui.NewButton()
	initBut.SetBounds(500, 5, 85, 25)
	initBut.SetText("초기화실행")
	ib := func() {
		initData()
	}
	initBut.SetOnClick(ib)

	panel1.Add(initBut)

	label2 := wui.NewLabel()
	label2.SetBounds(6, 5, 60, 20)
	label2.SetText("포트입력")
	panel1.Add(label2)
	nodeselect := wui.NewComboBox()
	nodeselect.SetBounds(70, 37, 80, 21)
	nodeselect.SetItems([]string{
		"seed",
		"full",
		"mine",
		"wallet",
	})
	nodeselect.SetSelectedIndex(0)
	panel1.Add(nodeselect)

	minecombo := wui.NewComboBox()
	minecombo.SetBounds(280, 37, 193, 21)
	minecombo.SetItems([]string{"Combo Box"})
	minecombo.SetSelectedIndex(0)
	panel1.Add(minecombo)

	label3 := wui.NewLabel()
	label3.SetBounds(6, 37, 60, 17)
	label3.SetText("노드선택")
	panel1.Add(label3)
	sendleft := wui.NewComboBox()
	sendleft.SetBounds(70, 74, 193, 21)
	sendleft.SetItems([]string{"Combo Box"})
	sendleft.SetSelectedIndex(0)
	sl := func(a int) {
		fromaddress = fmt.Sprintf("%s", sendleft.Text())
	}
	sendleft.SetOnChange(sl)

	panel1.Add(sendleft)

	label4 := wui.NewLabel()
	label4.SetBounds(6, 74, 60, 16)
	label4.SetText("주소선택")
	panel1.Add(label4)

	coinconfirm := wui.NewComboBox()
	coinconfirm.SetBounds(70, 108, 192, 21)
	coinconfirm.SetItems([]string{"Combo Box"})
	coinconfirm.SetSelectedIndex(0)
	ca := func(a int) {
		confirmaddress = fmt.Sprintf("%s", coinconfirm.Text())
	}
	coinconfirm.SetOnChange(ca)

	panel1.Add(coinconfirm)

	sendright := wui.NewComboBox()
	sendright.SetBounds(280, 74, 216, 21)
	sendright.SetItems([]string{"Combo Box"})
	sendright.SetSelectedIndex(0)
	sr := func(a int) {
		toaddress = fmt.Sprintf("%s", sendright.Text())
	}
	sendright.SetOnChange(sr)

	panel1.Add(sendright)
	blockchain := wui.NewComboBox()
	blockchain.SetBounds(70, 143, 193, 21)
	blockchain.SetItems([]string{"Combo Box"})
	blockchain.SetSelectedIndex(0)
	bb := func(a int) {
		blockchainaddress = fmt.Sprintf("%s", blockchain.Text())
	}
	blockchain.SetOnChange(bb)

	panel1.Add(blockchain)

	portleft := wui.NewComboBox()
	portleft.SetBounds(70, 5, 150, 21)
	portleft.SetItems([]string{
		"3000",
		"3001",
		"3002",
		"3003",
		"3004",
		"3005",
		"3006",
		"3007",
		"3008",
		"3009",
	})
	portleft.SetSelectedIndex(0)
	ln := func(a int) {
		leftport = fmt.Sprintf("%s", portleft.Text())
		genesisblockCopy()
		wf := fmt.Sprintf(walletFile, leftport)
		if dbExists(wf) {
			sendleft.SetItems([]string{})
			coinconfirm.SetItems([]string{})
			blockchain.SetItems([]string{})
			minecombo.SetItems([]string{})

			walletadd = cc.walletAddresses(leftport)
			sendleft.SetItems(walletadd)
			coinconfirm.SetItems(walletadd)
			blockchain.SetItems(walletadd)
			minecombo.SetItems(walletadd)

		}

	}
	portleft.SetOnChange(ln)
	panel1.Add(portleft)

	portright := wui.NewComboBox()
	portright.SetBounds(280, 5, 150, 21)
	portright.SetItems([]string{
		"3000",
		"3001",
		"3002",
		"3003",
		"3004",
		"3005",
		"3006",
		"3007",
		"3008",
		"3009",
	})
	portright.SetSelectedIndex(0)
	pr := func(a int) {
		rightport = fmt.Sprintf("%s", portright.Text())
		dbFile := fmt.Sprintf(dbFile, rightport)
		if dbExists(dbFile) {
			walletadd = cc.walletAddresses(rightport)
			sendright.SetItems([]string{})
			sendright.SetItems(walletadd)
		}

	}
	portright.SetOnChange(pr)

	panel1.Add(portright)

	sendcoin := wui.NewEditLine()
	sendcoin.SetBounds(536, 74, 70, 20)
	sendcoin.SetText("0")
	panel1.Add(sendcoin)

	havecoin := wui.NewLabel()
	havecoin.SetBounds(536, 108, 50, 16)
	havecoin.SetText("0")
	havecoin.SetAlignment(wui.AlignRight)
	panel1.Add(havecoin)

	sendBut := wui.NewButton()
	sendBut.SetBounds(627, 74, 85, 25)
	sendBut.SetText("코인전송")
	sb := func() {

		if sv, err := strconv.Atoi(fmt.Sprintf("%s", sendcoin.Text())); err == nil {
			cc.send(fromaddress, toaddress, sv, leftport, false)
		}
	}
	sendBut.SetOnClick(sb)

	panel1.Add(sendBut)

	confirmBut := wui.NewButton()
	confirmBut.SetBounds(628, 108, 85, 25)
	confirmBut.SetText("코인확인")
	cb := func() {
		havecoin.SetText(fmt.Sprintf("%d", cc.getBalanceValue(confirmaddress, leftport)))
	}
	confirmBut.SetOnClick(cb)

	panel1.Add(confirmBut)

	chainBut := wui.NewButton()
	chainBut.SetBounds(628, 143, 85, 25)
	chainBut.SetText("체인생성")
	createchain := func() {
		cc.createBlockchain(blockchainaddress, leftport)
		var from, to string
		from = fmt.Sprintf("blockchain_%s.db", leftport)
		to = "blockchain_genesis.db"
		err := CopyFile(from, to)
		if err != nil {
			fmt.Printf("CopyFile failed %q\n", err)
		} else {
			fmt.Printf("blockchain_genesis.db CopyFile succeeded\n")
		}
	}
	chainBut.SetOnClick(createchain)

	panel1.Add(chainBut)

	walletBut := wui.NewButton()
	walletBut.SetBounds(628, 5, 85, 25)
	walletBut.SetText("지갑생성")
	wb := func() {
		cc.createWallet(fmt.Sprintf(leftport))
	}
	walletBut.SetOnClick(wb)
	panel1.Add(walletBut)

	nodeBut := wui.NewButton()
	nodeBut.SetBounds(628, 40, 85, 25)
	nodeBut.SetText("노드생성")
	nc := func() {
		nodeid := fmt.Sprintf("%s", portleft.Text())
		nodeact := fmt.Sprintf("%s", nodeselect.Text())
		address := fmt.Sprintf("%s", minecombo.Text())

		if !nodeIsExe(nodeid) {
			nodeExe = append(nodeExe, nodeid)
			fmt.Printf("%s 노드 실행\n", nodeact)
			go StartNode(nodeid, nodeact, address)
		}
	}
	nodeBut.SetOnClick(nc)
	panel1.Add(nodeBut)

	win.Show()

}
func initData() {
	var from, to string
	from = "c:/God/src/main/data/"
	to = "c:/God/src/main"

	copyDir(from, to)
	fmt.Println("초기화,......")

}
func genesisblockCopy() {
	var from, to string
	from = "blockchain_genesis.db"
	to = ""
	if !dbExists(from) {
		return
	}
	to = fmt.Sprintf(dbFile, leftport)
	if dbExists(to) {
		return
	}
	err := CopyFile(from, to)
	if err != nil {
		fmt.Printf("CopyFile failed %q\n", err)
	} else {
		fmt.Printf("blockchain_genesis.db CopyFile succeeded\n")
	}

}

func nodeIsExe(addr string) bool {
	for _, node := range nodeExe {
		if node == addr {
			return true
		}
	}

	return false
}
