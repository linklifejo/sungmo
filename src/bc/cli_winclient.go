package bc

import (
	"fmt"
	"strconv"

	"github.com/gonutz/wui/v2"
	//	"time"
)

var walletadd []string
var blockchainaddress string
var fromaddress, toaddress, confirmaddress string
var node09, fromNode string
var nodefmw string
var serversm string

func (cli *CLI) WinClient() {
	cc := CLI{}
	winFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	win := wui.NewWindow()
	win.SetFont(winFont)
	win.SetInnerSize(728, 596)
	win.SetTitle("코인클라이언트")

	panel1Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -14,
	})

	panel1 := wui.NewPanel()
	panel1.SetFont(panel1Font)
	panel1.SetBounds(8, 12, 224, 297)
	panel1.SetBorderStyle(wui.PanelBorderSingleLine)
	win.Add(panel1)

	node3000 := wui.NewRadioButton()
	node3000.SetBounds(10, 36, 50, 17)
	node3000.SetText("3000")
	chk3000 := func(c bool) {
		if c == true {
			node09 = "3000"
			genesisblockCopy()
		}
	}
	node3000.SetOnCheck(chk3000)
	panel1.Add(node3000)

	nodeselect := wui.NewLabel()
	nodeselect.SetBounds(7, 6, 70, 13)
	nodeselect.SetText("노드 선택")
	panel1.Add(nodeselect)

	node3001 := wui.NewRadioButton()
	node3001.SetBounds(9, 62, 50, 17)
	node3001.SetText("3001")
	chk3001 := func(c bool) {
		if c == true {
			node09 = "3001"
			genesisblockCopy()
		}
	}
	node3001.SetOnCheck(chk3001)
	panel1.Add(node3001)

	node3002 := wui.NewRadioButton()
	node3002.SetBounds(8, 88, 50, 17)
	node3002.SetText("3002")
	chk3002 := func(c bool) {
		if c == true {
			node09 = "3002"
			genesisblockCopy()
		}
	}
	node3002.SetOnCheck(chk3002)
	panel1.Add(node3002)

	node3003 := wui.NewRadioButton()
	node3003.SetBounds(8, 114, 50, 17)
	node3003.SetText("3003")
	chk3003 := func(c bool) {
		if c == true {
			node09 = "3003"
			genesisblockCopy()
		}
	}
	node3003.SetOnCheck(chk3003)
	panel1.Add(node3003)

	node3004 := wui.NewRadioButton()
	node3004.SetBounds(7, 140, 50, 17)
	node3004.SetText("3004")
	chk3004 := func(c bool) {
		if c == true {
			node09 = "3004"
			genesisblockCopy()
		}
	}
	node3004.SetOnCheck(chk3004)

	panel1.Add(node3004)

	node3005 := wui.NewRadioButton()
	node3005.SetBounds(7, 165, 50, 17)
	node3005.SetText("3005")
	chk3005 := func(c bool) {
		if c == true {
			node09 = "3005"
			genesisblockCopy()
		}
	}
	node3005.SetOnCheck(chk3005)
	panel1.Add(node3005)

	node3006 := wui.NewRadioButton()
	node3006.SetBounds(7, 191, 50, 17)
	node3006.SetText("3006")
	chk3006 := func(c bool) {
		if c == true {
			node09 = "3006"
			genesisblockCopy()
		}
	}
	node3006.SetOnCheck(chk3006)

	panel1.Add(node3006)

	node3007 := wui.NewRadioButton()
	node3007.SetBounds(7, 215, 50, 17)
	node3007.SetText("3007")
	chk3007 := func(c bool) {
		if c == true {
			node09 = "3007"
			genesisblockCopy()
		}
	}
	node3007.SetOnCheck(chk3007)
	panel1.Add(node3007)

	node3008 := wui.NewRadioButton()
	node3008.SetBounds(7, 237, 50, 17)
	node3008.SetText("3008")
	chk3008 := func(c bool) {
		if c == true {
			node09 = "3008"
			genesisblockCopy()
		}
	}
	node3008.SetOnCheck(chk3008)
	panel1.Add(node3008)

	node3009 := wui.NewRadioButton()
	node3009.SetBounds(7, 261, 50, 17)
	node3009.SetText("3009")
	chk3009 := func(c bool) {
		if c == true {
			node09 = "3009"
			genesisblockCopy()
		}
	}
	node3009.SetOnCheck(chk3009)
	panel1.Add(node3009)

	walletbutton := wui.NewButton()
	walletbutton.SetBounds(88, 131, 85, 25)
	walletbutton.SetText("지갑생성")
	walletb := func() {
		fmt.Println(node09)
		cc.createWallet(node09)
	}
	walletbutton.SetOnClick(walletb)

	panel1.Add(walletbutton)

	panel2Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	panel2 := wui.NewPanel()
	panel2.SetFont(panel2Font)
	panel2.SetBounds(242, 13, 233, 297)
	panel2.SetBorderStyle(wui.PanelBorderSingleLine)
	win.Add(panel2)

	nodecreate := wui.NewLabel()
	nodecreate.SetBounds(8, 8, 50, 13)
	nodecreate.SetText("노드 생성")
	panel2.Add(nodecreate)

	fullnode := wui.NewRadioButton()
	fullnode.SetBounds(8, 29, 70, 17)
	fullnode.SetText("풀 노드")
	chkfull := func(c bool) {
		if c == true {
			nodefmw = "fullnode"

			fmt.Println("full node")
		}
	}
	fullnode.SetOnCheck(chkfull)

	panel2.Add(fullnode)

	minenode := wui.NewRadioButton()
	minenode.SetBounds(8, 55, 70, 17)
	minenode.SetText("채굴노드")
	chkmine := func(c bool) {
		if c == true {
			nodefmw = "minenode"

			fmt.Println("minenode")
		}
	}
	minenode.SetOnCheck(chkmine)

	panel2.Add(minenode)

	walletnode := wui.NewRadioButton()
	walletnode.SetBounds(9, 81, 70, 17)
	walletnode.SetText("지갑노드")
	chkwallet := func(c bool) {
		if c == true {
			nodefmw = "walletnode"

			fmt.Println("walletnode")
		}
	}
	walletnode.SetOnCheck(chkwallet)

	panel2.Add(walletnode)

	nodebutton := wui.NewButton()
	nodebutton.SetBounds(10, 131, 85, 25)
	nodebutton.SetText("노드생성")
	nodeb := func() {
		fmt.Println("노드생성")
	}
	nodebutton.SetOnClick(nodeb)

	panel2.Add(nodebutton)

	label1 := wui.NewLabel()
	label1.SetBounds(94, 8, 70, 13)
	label1.SetText("생성수")
	panel2.Add(label1)

	label2 := wui.NewLabel()
	label2.SetBounds(164, 8, 50, 13)
	label2.SetText("실행수")
	panel2.Add(label2)

	fullcreate := wui.NewLabel()
	fullcreate.SetBounds(97, 30, 30, 13)
	fullcreate.SetText("0")
	fullcreate.SetAlignment(wui.AlignRight)
	panel2.Add(fullcreate)

	minecreate := wui.NewLabel()
	minecreate.SetBounds(96, 56, 30, 13)
	minecreate.SetText("0")
	minecreate.SetAlignment(wui.AlignRight)
	panel2.Add(minecreate)

	walletcreate := wui.NewLabel()
	walletcreate.SetBounds(97, 81, 30, 13)
	walletcreate.SetText("0")
	walletcreate.SetAlignment(wui.AlignRight)
	panel2.Add(walletcreate)

	fullexec := wui.NewLabel()
	fullexec.SetBounds(162, 30, 30, 13)
	fullexec.SetText("0")
	fullexec.SetAlignment(wui.AlignRight)
	panel2.Add(fullexec)

	mineexec := wui.NewLabel()
	mineexec.SetBounds(161, 57, 30, 13)
	mineexec.SetText("0")
	mineexec.SetAlignment(wui.AlignRight)
	panel2.Add(mineexec)

	walletexec := wui.NewLabel()
	walletexec.SetBounds(160, 79, 30, 13)
	walletexec.SetText("0")
	walletexec.SetAlignment(wui.AlignRight)
	panel2.Add(walletexec)

	panel4Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	panel4 := wui.NewPanel()
	panel4.SetFont(panel4Font)
	panel4.SetBounds(483, 14, 233, 296)
	panel4.SetBorderStyle(wui.PanelBorderSingleLine)
	win.Add(panel4)

	servercreate := wui.NewLabel()
	servercreate.SetBounds(7, 8, 70, 13)
	servercreate.SetText("서버생성")
	panel4.Add(servercreate)

	seedserver := wui.NewRadioButton()
	seedserver.SetBounds(7, 27, 70, 17)
	seedserver.SetText("시드서버")
	chkseed := func(c bool) {
		if c == true {
			serversm = "seedserver"

			fmt.Println("seedserver")
		}
	}
	seedserver.SetOnCheck(chkseed)

	panel4.Add(seedserver)

	manageserver := wui.NewRadioButton()
	manageserver.SetBounds(7, 49, 70, 17)
	manageserver.SetText("관리서버")
	chkmanage := func(c bool) {
		if c == true {
			serversm = "manageserver"

			fmt.Println("manageserver")
		}
	}
	manageserver.SetOnCheck(chkmanage)

	panel4.Add(manageserver)

	severbutton := wui.NewButton()
	severbutton.SetBounds(9, 131, 85, 25)
	severbutton.SetText("서버생성")
	severb := func() {
		fmt.Println("server")
	}
	severbutton.SetOnClick(severb)
	panel4.Add(severbutton)

	label4 := wui.NewLabel()
	label4.SetBounds(93, 5, 70, 13)
	label4.SetText("생성수")
	panel4.Add(label4)

	label5 := wui.NewLabel()
	label5.SetBounds(182, 7, 70, 13)
	label5.SetText("실행수")
	panel4.Add(label5)

	seedcreate := wui.NewLabel()
	seedcreate.SetBounds(95, 30, 30, 13)
	seedcreate.SetText("0")
	seedcreate.SetAlignment(wui.AlignRight)
	panel4.Add(seedcreate)

	seedexec := wui.NewLabel()
	seedexec.SetBounds(180, 27, 30, 13)
	seedexec.SetText("0")
	seedexec.SetAlignment(wui.AlignRight)
	panel4.Add(seedexec)

	managecreate := wui.NewLabel()
	managecreate.SetBounds(94, 51, 30, 13)
	managecreate.SetText("0")
	managecreate.SetAlignment(wui.AlignRight)
	panel4.Add(managecreate)

	manageexec := wui.NewLabel()
	manageexec.SetBounds(180, 49, 30, 13)
	manageexec.SetText("0")
	manageexec.SetAlignment(wui.AlignRight)
	panel4.Add(manageexec)

	panel5Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	panel5 := wui.NewPanel()
	panel5.SetFont(panel5Font)
	panel5.SetBounds(8, 320, 709, 120)
	panel5.SetBorderStyle(wui.PanelBorderSingleLine)
	win.Add(panel5)

	// coinsend := wui.NewLabel()
	// coinsend.SetBounds(6, 7, 50, 13)
	// coinsend.SetText("코인전송")
	// panel5.Add(coinsend)
	fromcombo := wui.NewComboBox()
	fromcombo.SetBounds(75, 7, 250, 21)
	fromcombo.SetItems([]string{"ComboBox"})
	fromcombo.SetSelectedIndex(0)
	fromadd := func(a int) {
		fromaddress = fromcombo.Text()
	}
	fromcombo.SetOnChange(fromadd)
	panel5.Add(fromcombo)

	fromnode := wui.NewLabel()
	fromnode.SetBounds(75, 30, 100, 13)
	fromnode.SetAlignment(wui.AlignRight)
	panel5.Add(fromnode)

	coinfrom := wui.NewRadioButton()
	coinfrom.SetBounds(6, 7, 40, 17)
	coinfrom.SetText("from")
	chkfrom := func(c bool) {
		if c == true && node09 != "" {
			walletadd = cc.walletAddresses(node09)
			fromcombo.SetItems(walletadd)
			fromnode.SetText(node09)
			fromNode = node09
		}
	}
	coinfrom.SetOnCheck(chkfrom)
	panel5.Add(coinfrom)
	tocombo := wui.NewComboBox()
	tocombo.SetBounds(440, 7, 250, 21)
	tocombo.SetItems([]string{"ComboBox"})
	tocombo.SetSelectedIndex(0)
	toadd := func(a int) {
		toaddress = tocombo.Text()
	}
	tocombo.SetOnChange(toadd)

	panel5.Add(tocombo)
	tonode := wui.NewLabel()
	tonode.SetBounds(440, 30, 100, 13)
	tonode.SetText(node09)
	tonode.SetAlignment(wui.AlignRight)
	panel5.Add(tonode)

	cointo := wui.NewRadioButton()
	cointo.SetBounds(390, 7, 25, 17)
	cointo.SetText("to")
	chkto := func(c bool) {
		if c == true && node09 != "" {
			walletadd = cc.walletAddresses(node09)
			tocombo.SetItems(walletadd)
			tonode.SetText(node09)

		}
	}
	cointo.SetOnCheck(chkto)
	panel5.Add(cointo)
	sendcoin := wui.NewEditLine()
	sendcoin.SetHorizontalAnchor(wui.AnchorMax)
	sendcoin.SetBounds(448, 83, 150, 20)
	sendcoin.SetText("0")
	panel5.Add(sendcoin)

	sendbutton := wui.NewButton()
	sendbutton.SetBounds(608, 81, 85, 25)
	sendbutton.SetText("코인전송")
	sendb := func() {
		if sv, err := strconv.Atoi(fmt.Sprintf("%s", sendcoin.Text())); err == nil {
			cc.send(fromaddress, toaddress, sv, fromNode, true)
		}

		fmt.Println("코인전송")
	}
	sendbutton.SetOnClick(sendb)

	panel5.Add(sendbutton)

	panel6Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	panel6 := wui.NewPanel()
	panel6.SetFont(panel6Font)
	panel6.SetBounds(7, 452, 709, 58)
	panel6.SetBorderStyle(wui.PanelBorderSingleLine)
	win.Add(panel6)

	coinconfirm := wui.NewLabel()
	coinconfirm.SetBounds(9, 5, 70, 13)
	coinconfirm.SetText("코인확인")
	panel6.Add(coinconfirm)
	confirmcombo := wui.NewComboBox()
	confirmcombo.SetBounds(80, 23, 250, 21)
	confirmcombo.SetItems([]string{"ComboBox"})
	confirmcombo.SetSelectedIndex(0)
	confirmadd := func(a int) {
		confirmaddress = confirmcombo.Text()
	}
	confirmcombo.SetOnChange(confirmadd)

	panel6.Add(confirmcombo)

	confirm := wui.NewRadioButton()
	confirm.SetBounds(80, 5, 50, 17)
	confirm.SetText("confirm")
	chkcon := func(c bool) {
		if c == true && node09 != "" {
			walletadd = cc.walletAddresses(node09)
			confirmcombo.SetItems(walletadd)

		}
	}
	confirm.SetOnCheck(chkcon)
	panel6.Add(confirm)

	havecoin := wui.NewLabel()
	havecoin.SetBounds(591, 25, 100, 15)
	havecoin.SetText("0")
	havecoin.SetAlignment(wui.AlignRight)
	panel6.Add(havecoin)

	confirmsu := wui.NewLabel()
	confirmsu.SetBounds(653, 3, 50, 19)
	confirmsu.SetText("보유수량")
	panel6.Add(confirmsu)

	label10 := wui.NewLabel()
	label10.SetBounds(9, 23, 50, 15)
	label10.SetText("주소선택")
	panel6.Add(label10)

	confirmbutton := wui.NewButton()
	confirmbutton.SetBounds(500, 5, 85, 25)
	confirmbutton.SetText("코인확인")
	confirmb := func() {
		havecoin.SetText(fmt.Sprintf("%d", cc.getBalanceValue(confirmaddress, node09)))
	}
	confirmbutton.SetOnClick(confirmb)
	panel6.Add(confirmbutton)

	////////////////////////////////////////////////////////////
	panel7Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	panel7 := wui.NewPanel()
	panel7.SetFont(panel7Font)
	panel7.SetBounds(7, 525, 709, 58)
	panel7.SetBorderStyle(wui.PanelBorderSingleLine)
	win.Add(panel7)

	blockcreate := wui.NewLabel()
	blockcreate.SetBounds(9, 5, 70, 13)
	blockcreate.SetText("블럭생성")
	panel7.Add(blockcreate)
	blockcombo := wui.NewComboBox()
	blockcombo.SetBounds(80, 23, 250, 21)
	blockcombo.SetItems([]string{"ComboBox"})
	blockcombo.SetSelectedIndex(0)
	badd := func(a int) {
		blockchainaddress = blockcombo.Text()
	}
	blockcombo.SetOnChange(badd)

	panel7.Add(blockcombo)

	bccreate := wui.NewRadioButton()
	bccreate.SetBounds(80, 5, 50, 17)
	bccreate.SetText("confirm")
	chkblock := func(c bool) {
		if c == true && node09 != "" {
			walletadd = cc.walletAddresses(node09)
			blockcombo.SetItems(walletadd)

		}
	}
	bccreate.SetOnCheck(chkblock)
	panel7.Add(bccreate)

	label11 := wui.NewLabel()
	label11.SetBounds(9, 23, 50, 15)
	label11.SetText("주소선택")
	panel7.Add(label11)

	createbutton := wui.NewButton()
	createbutton.SetBounds(608, 21, 85, 25)
	createbutton.SetText("체인생성")
	createb := func() {
		cc.createBlockchain(blockchainaddress, node09)
		var from, to string
		from = fmt.Sprintf("blockchain_%s.db", node09)
		to = "blockchain_genesis.db"
		err := CopyFile(from, to)
		if err != nil {
			fmt.Printf("CopyFile failed %q\n", err)
		} else {
			fmt.Printf("blockchain_genesis.db CopyFile succeeded\n")
		}
	}
	createbutton.SetOnClick(createb)

	panel7.Add(createbutton)

	win.Show()
}
func genesisblockCopy() {
	var from, to string
	from = "blockchain_genesis.db"
	if !dbExists(from) {
		return
	}
	to = fmt.Sprintf(dbFile, node09)
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
