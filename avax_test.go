package bridge

import (
	"fmt"
	"math/big"
	"testing"
	"time"
	"strings"

	pancakeproxy "github.com/incognitochain/bridge-eth/bridge/pancake"

	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type AvaxTestSuite struct {
	*TradingTestSuite

	PanackeTradeDeployedAddr common.Address
	PanackeRouteContractAddr common.Address

	IncAVAXTokenIDStr 	string
	IncUTAVAXTokenIDStr 	string
	IncUTAXUSDTTokenIDStr 	string
	IncAXUSDTTokenIDStr 	string

	// token amounts for tests
	DepositingEther       float64
	AVAXBalanceAfterStep1 *big.Int
	USDTBalanceAfterStep2 *big.Int

	AVAXAddressStr string
	WAVAXAddr    common.Address
	AXUSDTTokenAddress string
	USDCTokenAddress string

}

func NewAvaxTestSuite(tradingTestSuite *TradingTestSuite) *AvaxTestSuite {
	return &AvaxTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *AvaxTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	

	tradingSuite.PanackeTradeDeployedAddr = common.HexToAddress("0xd17E836453f7DaF2F2d6F8dFdd56449bc97446F4")
	tradingSuite.PanackeRouteContractAddr = common.HexToAddress("0xd7f655E3376cE2D7A2b08fF01Eb3B1023191A901")

	tradingSuite.WAVAXAddr = common.HexToAddress("0xd00ae08403B9bbb9124bB305C09058E32C39A48c")

	tradingSuite.IncUTAVAXTokenIDStr ="5075e4903091b61d2a7a3dd9cd5d369b026900301dadbadefef3e35a77ac4073"
	tradingSuite.IncAVAXTokenIDStr ="c469fb02623a023b469c81e1564193da7d85fe918cd4a4fdd2c64f97f59f60f5"
	tradingSuite.IncUTAXUSDTTokenIDStr = "3a526c0fa9abfc3e3e37becc52c5c10abbb7897b0534ad17018e766fc6133590"
	tradingSuite.IncAXUSDTTokenIDStr = "8133c5a0e01fc03ca77224e653e160ef158ba5d75238816f90d70fa30e3abd63"
	tradingSuite.USDCTokenAddress = "593f9d277eb34de8665c0d38df94b06d9d78c7d87b6d7fd531cc14ec34a7d4b8"

	tradingSuite.AVAXAddressStr ="0x0000000000000000000000000000000000000000"
	tradingSuite.AXUSDTTokenAddress = "0x48601da98F729aE3B4d7DeD8F29777DF102a167d"

	tradingSuite.DepositingEther = float64(0.1)
	
}

func (tradingSuite *AvaxTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.AVAXClient.Close()
}

func (tradingSuite *AvaxTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *AvaxTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAvaxTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for avax test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	avaxSuite := NewAvaxTestSuite(tradingSuite)
	suite.Run(t, avaxSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}


func (tradingSuite *AvaxTestSuite) getExpectedAmount(
	path []common.Address,
	srcQty *big.Int,
) *big.Int {
	c, err := pancakeproxy.NewPancakeproxyrouter(tradingSuite.PanackeRouteContractAddr, tradingSuite.AVAXClient)
	require.Equal(tradingSuite.T(), nil, err)
	amounts, err := c.GetAmountsOut(nil, srcQty, path)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("path: %v\n", path)
	fmt.Printf("intput value: %v\n", srcQty.String())
	fmt.Printf("output value: %v\n", amounts[len(amounts)-1].String())
	fmt.Printf("path output : %v\n", amounts)

	return amounts[len(amounts)-1]
}

func (tradingSuite *AvaxTestSuite) executeWithPancake(
	srcQty *big.Int,
	path []common.Address,
	deadline uint,
	isNative bool,
) {
	require.NotEqual(tradingSuite.T(), 0, len(path))

	tradeAbi, err := abi.JSON(strings.NewReader(pancakeproxy.PancakeproxyMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAVAXAddr, tradingSuite.AVAXClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDAVAX)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(10e10)
	expectOutputAmount := tradingSuite.getExpectedAmount(path, srcQty)
	input, err := tradeAbi.Pack("trade", path, srcQty, expectOutputAmount, big.NewInt(int64(deadline)), isNative)
	require.Equal(tradingSuite.T(), nil, err)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(tradingSuite.T(), nil, err)
	sourceToken := path[0]
	if path[0].String() == tradingSuite.WAVAXAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.AVAXAddressStr)
	}
	destToken := path[len(path)-1]
	if path[len(path)-1].String() == tradingSuite.WAVAXAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.AVAXAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    AVAX_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.PanackeTradeDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.PanackeTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.AVAXClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Dapp trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *AvaxTestSuite) CallData(
	srcQty *big.Int,
	path []common.Address,
	deadline uint,
	isNative bool,
) []byte {
	require.NotEqual(tradingSuite.T(), 0, len(path))

	tradeAbi, err := abi.JSON(strings.NewReader(pancakeproxy.PancakeproxyMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	expectOutputAmount := tradingSuite.getExpectedAmount(path, srcQty)
	amount := big.NewInt(0).Div(expectOutputAmount, big.NewInt(1))
	// amount := big.NewInt(2082011)
	input, err := tradeAbi.Pack("trade", path, srcQty, amount, big.NewInt(int64(deadline)), isNative)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("Call Data : ", common.Bytes2Hex(input))
	return input
}



func (tradingSuite *AvaxTestSuite) Test1Avax() {
return
	fmt.Println("============ TEST SHIELD UNSHIELD NATIVE AVAX ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting AVAX to pAVAX --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.AVAXHost, txHash)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.AVAXHost, common.HexToHash("0xc6f0e2be783a7df2ea32dc243fe941c8fc2992fe0dc3bfb49cb67e833bdcfb87"))
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncAVAXTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(40 * time.Second)
return
	fmt.Println("------------ STEP 2: burning pAVAX to deposit AVAX to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncAVAXTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningavaxfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(40 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		"getburnavaxprooffordeposittosc",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress(tradingSuite.AVAXAddressStr),
		pubKeyToAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	fmt.Println("deposited AVAX: ", deposited)

	fmt.Println("------------ step 3: withdrawing AVAX from SC to pAVAX on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.AVAXAddressStr,
		deposited,
		tradingSuite.AVAXClient,
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		tradingSuite.VaultAVAXAddr,
		AVAX_REQ_WITHDRAW_PREFIX,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.AVAXHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncAVAXTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(40 * time.Second)

	fmt.Println("------------ step 4: withdrawing pAVAX from Incognito to AVAX --------------")
	withdrawingPAVAX := big.NewInt(0).Div(deposited, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncAVAXTokenIDStr,
		withdrawingPAVAX,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningavaxrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(40 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getavaxburnproof",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.AVAXAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
	tradingSuite.AVAXBalanceAfterStep1 = bal
	fmt.Println("AVAX balance after step 1: ", tradingSuite.AVAXBalanceAfterStep1)
}

func (tradingSuite *AvaxTestSuite) Test2AvaxToken() {
return
	fmt.Println("============ TEST SHIELD UNSHIELD TOKEN AVAX ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingUSDT := big.NewInt(3e4)
	burningPUSDT := depositingUSDT

	daibal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.AXUSDTTokenAddress),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
	fmt.Println("balance balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting USDT to pUSDT --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingUSDT,
		common.HexToAddress(tradingSuite.AXUSDTTokenAddress),
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.AVAXHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncAXUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(40 * time.Second)

	fmt.Println("------------ step 2: burning pUSDT to deposit USDT to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncAXUSDTTokenIDStr,
		burningPUSDT,
		pubKeyToAddrStr[2:],
		"createandsendburningavaxfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(40 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		"getburnavaxprooffordeposittosc",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress(tradingSuite.AXUSDTTokenAddress),
		pubKeyToAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	fmt.Println("deposited dai: ", deposited)

	fmt.Println("------------ step 3: withdrawing USDT from SC to pUSDT on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.AXUSDTTokenAddress,
		deposited,
		tradingSuite.AVAXClient,
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		tradingSuite.VaultAVAXAddr,
		AVAX_REQ_WITHDRAW_PREFIX,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.AVAXHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncAXUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(50 * time.Second)

	fmt.Println("------------ step 4: withdrawing pUSDT from Incognito to USDT --------------")
	withdrawingPMRK := deposited
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncAXUSDTTokenIDStr,
		withdrawingPMRK,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningavaxrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(40 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getavaxburnproof",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.AXUSDTTokenAddress),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
	tradingSuite.USDTBalanceAfterStep2 = bal
	fmt.Println("USDT balance after step 2: ", tradingSuite.USDTBalanceAfterStep2)
}

func (tradingSuite *AvaxTestSuite) Test3UTAvax() {
	return
		fmt.Println("============ TEST 3 SHIELD UNSHIELD UNIFIED NATIVE AVAX ===========")
		fmt.Println("------------ STEP 0: declaration & initialization --------------")
		tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
		burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
	
		pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
		fmt.Println("------------ STEP 1: porting AVAX to pAVAX --------------")
		txHash := tradingSuite.depositETH(
			tradingSuite.DepositingEther,
			tradingSuite.IncPaymentAddrStr,
			tradingSuite.VaultAVAXAddr,
			tradingSuite.AVAXClient,
		)
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.AVAXHost, txHash)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
		txhashInC, err := tradingSuite.callIssuingUnifiedPtokenReq(
			tradingSuite.IncUTAVAXTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
			"bridgeaggShield",
			6,
			tradingSuite.IncAVAXTokenIDStr,
		)
		require.Equal(tradingSuite.T(), nil, err)

		time.Sleep(40 * time.Second)
		require.Equal(tradingSuite.T(), 2, tradingSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
		balpBNBS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr,tradingSuite.IncUTAVAXTokenIDStr)
		fmt.Println("[INFO] UT AVAX balance incognito after issuing step 1 : ", balpBNBS1)
	// return
		fmt.Println("------------ STEP 2: burning pAVAX to deposit AVAX to SC --------------")
		// make a burn tx to incognito chain as a result of deposit to SC
		burningRes, err := tradingSuite.callBurningUnifiedPToken(
			tradingSuite.IncUTAVAXTokenIDStr,
			burningPETH,
			big.NewInt(0),
			pubKeyToAddrStr[2:],
			"bridgeaggUnshield",
			tradingSuite.IncPaymentAddrStr,
			0,
			tradingSuite.IncAVAXTokenIDStr,
			true,
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found := burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(40 * time.Second)
	

		tradingSuite.submitBurnProofForDepositToSCV2(
			burningTxID.(string),
			// "e55841caded175ed168c09c2dc2e4ae6c8ba7bbcfb1475216b70c6185dba9371",
			big.NewInt(int64(tradingSuite.ChainIDAVAX)),
			"bridgeaggGetBurnProof",
			tradingSuite.VaultAVAXAddr,
			tradingSuite.AVAXClient,
			0,
			6,
		)

		deposited := tradingSuite.getDepositedBalanceWithParams(
			common.HexToAddress(tradingSuite.AVAXAddressStr),
			pubKeyToAddrStr,
			tradingSuite.VaultAVAXAddr,
			tradingSuite.AVAXClient,
		)
		fmt.Println("deposited AVAX: ", deposited)
	
		fmt.Println("------------ step 3: withdrawing AVAX from SC to pAVAX on Incognito --------------")
		txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
			tradingSuite.AVAXAddressStr,
			deposited,
			tradingSuite.AVAXClient,
			big.NewInt(int64(tradingSuite.ChainIDAVAX)),
			tradingSuite.VaultAVAXAddr,
			AVAX_REQ_WITHDRAW_PREFIX,
		)
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.AVAXHost, txHashByEmittingWithdrawalReq)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
		_, err = tradingSuite.callIssuingUnifiedPtokenReq(
			tradingSuite.IncUTAVAXTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
			"bridgeaggShield",
			6,
			tradingSuite.IncAVAXTokenIDStr,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(40 * time.Second)
	
		fmt.Println("------------ step 4: withdrawing pAVAX from Incognito to AVAX --------------")
		withdrawingPAVAX := big.NewInt(0).Div(deposited, big.NewInt(1000000000))
		burningRes, err = tradingSuite.callBurningUnifiedPToken(
			tradingSuite.IncUTAVAXTokenIDStr,
			withdrawingPAVAX,
			big.NewInt(0),
			pubKeyToAddrStr[2:],
			"bridgeaggUnshield",
			tradingSuite.IncPaymentAddrStr,
			0,
			tradingSuite.IncAVAXTokenIDStr,
			false,
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found = burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(50 * time.Second)
	
		tradingSuite.submitBurnProofForWithdrawalV2(
			burningTxID.(string),
			"bridgeaggGetBurnProof",
			tradingSuite.VaultAVAXAddr,
			tradingSuite.AVAXClient,
			tradingSuite.ChainIDAVAX,
			0,
			6,
		)

	
		bal := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.AVAXAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
			tradingSuite.AVAXClient,
		)
		tradingSuite.AVAXBalanceAfterStep1 = bal
		fmt.Println("AVAX balance after step 1: ", tradingSuite.AVAXBalanceAfterStep1)
}

func (tradingSuite *AvaxTestSuite) Test4UTToken() {
return
	fmt.Println("============ TEST 4 SHIELD UNSHIELD UNIFIED TOKEN AVAX ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingUSDT := big.NewInt(3e4)
	burningPUSDT := depositingUSDT

	daibal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.AXUSDTTokenAddress),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
	fmt.Println("balance balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting USDT to pUSDT --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingUSDT,
		common.HexToAddress(tradingSuite.AXUSDTTokenAddress),
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.AVAXHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	txhashInC, err := tradingSuite.callIssuingUnifiedPtokenReq(
		tradingSuite.IncUTAXUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		6,
		tradingSuite.IncAXUSDTTokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(40 * time.Second)
	require.Equal(tradingSuite.T(), 2, tradingSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr,tradingSuite.IncUTAXUSDTTokenIDStr)
	fmt.Println("[INFO] UT AVAX balance incognito after issuing step 1 : ", balpBNBS1)
// return
	fmt.Println("------------ step 2: burning pUSDT to deposit USDT to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTAXUSDTTokenIDStr,
		burningPUSDT,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncAXUSDTTokenIDStr,
		true,
	)

	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(40 * time.Second)

	tradingSuite.submitBurnProofForDepositToSCV2(
		burningTxID.(string),
		// "e55841caded175ed168c09c2dc2e4ae6c8ba7bbcfb1475216b70c6185dba9371",
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		0,
		6,
	)

	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress(tradingSuite.AXUSDTTokenAddress),
		pubKeyToAddrStr,
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
	)
	fmt.Println("deposited dai: ", deposited)

	fmt.Println("------------ step 3: withdrawing USDT from SC to pUSDT on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.AXUSDTTokenAddress,
		deposited,
		tradingSuite.AVAXClient,
		big.NewInt(int64(tradingSuite.ChainIDAVAX)),
		tradingSuite.VaultAVAXAddr,
		AVAX_REQ_WITHDRAW_PREFIX,
	)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.AVAXHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	_, err = tradingSuite.callIssuingUnifiedPtokenReq(
		tradingSuite.IncUTAXUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		6,
		tradingSuite.IncAXUSDTTokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(50 * time.Second)

	fmt.Println("------------ step 4: withdrawing pUSDT from Incognito to USDT --------------")
	burningRes, err = tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTAXUSDTTokenIDStr,
		burningPUSDT,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncAXUSDTTokenIDStr,
		false,
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(40 * time.Second)

	tradingSuite.submitBurnProofForWithdrawalV2(
		burningTxID.(string),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
		0,
		6,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.AXUSDTTokenAddress),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
	tradingSuite.USDTBalanceAfterStep2 = bal
	fmt.Println("USDT balance after step 2: ", tradingSuite.USDTBalanceAfterStep2)
}

func (tradingSuite *AvaxTestSuite) TestX1_get_call_data() {
	return
	fmt.Println("===== GET call data =====")
	tradingSuite.CallData(
		big.NewInt(100000),
		[]common.Address{
			tradingSuite.WAVAXAddr,
			common.HexToAddress("0x3AFb22cdF460F2299D8b6b4443e1c846882646D9"),
		},
		uint(time.Now().Unix()+60000),
		false)
}

func (tradingSuite *AvaxTestSuite) TestX2_trade_AVAX_USDC_new_flow() {
	return
	fmt.Println("===== TestX2 TRADE AVAX USDC NEW FLOW =====")
	fmt.Println("===== GET call data =====")
	input := tradingSuite.CallData(
		big.NewInt(1000000 * 1e9),
		[]common.Address{
			tradingSuite.WAVAXAddr,
			// common.HexToAddress("0x3AFb22cdF460F2299D8b6b4443e1c846882646D9"),
			common.HexToAddress("0x6346C85e28b1E3276828E8153aFA21BA8F522A5e"),
			
		},
		uint(time.Now().Unix()+60000),
		false)

	fmt.Println(" ===== CALL BURNING ======")
	txhash, err := tradingSuite.callBurningDapp(
		tradingSuite.IncAVAXTokenIDStr,
		big.NewInt(1000000),
		// "0xcE40cE511A5D084017DBee7e3fF3e455ea32D85c",
		"0x0000000000000000000000000000000000000000",
		tradingSuite.PanackeTradeDeployedAddr.String(),
		"bridgeaggBurnForCall",
		tradingSuite.IncPaymentReceiverStr,
		0,
		tradingSuite.IncAVAXTokenIDStr,
		6,
		input,
		"0x6346C85e28b1E3276828E8153aFA21BA8F522A5e",
	)
	require.Equal(tradingSuite.T(), nil, err)
	// return
	time.Sleep(30 * time.Second)
	fmt.Println(" ==== SUBMIT TRADE AVAX USDC ===")
	tx_submit_trade := tradingSuite.submitBurnProofForWithdrawalNewDapp(
		txhash,
		"bridgeaggGetBurnProof",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
		0,
		6,
	)

	time.Sleep(40 * time.Second)
	fmt.Println("==== WITHDRAW ======")

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.AVAXHost, tx_submit_trade)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	// old shield flow
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.USDCTokenAddress,
		// tradingSuite.IncAVAXTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)

	// // // new flow unified shield
	// // _, err := tradingSuite.callIssuingUnifiedPtokenReq(
	// // 	tradingSuite.IncUTEtherTokenIDStr,
	// // 	ethDepositProof,
	// // 	ethBlockHash,
	// // 	ethTxIdx,
	// // 	"bridgeaggShield",
	// // 	3,
	// // 	tradingSuite.IncEtherTokenIDStr,
	// // )

	require.Equal(tradingSuite.T(), nil, err)

	// time.Sleep(60 * time.Second)
}

func (tradingSuite *AvaxTestSuite) TestX3_trade_USDC_AVAX_new_flow() {
	// return
	fmt.Println("===== TestX3 TRADE USDC AVAX NEW FLOW =====")
	fmt.Println("===== GET call data =====")
	input := tradingSuite.CallData(
		big.NewInt(1000),
		[]common.Address{
			// common.HexToAddress("0x3AFb22cdF460F2299D8b6b4443e1c846882646D9"),
			common.HexToAddress("0x6346C85e28b1E3276828E8153aFA21BA8F522A5e"),
			tradingSuite.WAVAXAddr,
			
		},
		uint(time.Now().Unix()+60000),
		false)

	fmt.Println(" ===== CALL BURNING ======")
	txhash, err := tradingSuite.callBurningDapp(
		tradingSuite.USDCTokenAddress,
		big.NewInt(1000),
		// "0xcE40cE511A5D084017DBee7e3fF3e455ea32D85c",
		"0x0000000000000000000000000000000000000000",
		tradingSuite.PanackeTradeDeployedAddr.String(),
		"bridgeaggBurnForCall",
		tradingSuite.IncPaymentReceiverStr,
		0,
		tradingSuite.USDCTokenAddress,
		6,
		input,
		"0x0000000000000000000000000000000000000000",
	)
	require.Equal(tradingSuite.T(), nil, err)
	// return
	time.Sleep(100 * time.Second)
	fmt.Println(" ==== SUBMIT TRADE USDC AVAX  ===")
	tx_submit_trade := tradingSuite.submitBurnProofForWithdrawalNewDapp(
		txhash,
		"bridgeaggGetBurnProof",
		tradingSuite.VaultAVAXAddr,
		tradingSuite.AVAXClient,
		tradingSuite.ChainIDAVAX,
		0,
		6,
	)

	time.Sleep(40 * time.Second)
	fmt.Println("==== WITHDRAW ======")

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.AVAXHost, tx_submit_trade)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(60 * time.Second)

	// old shield flow
	_, err = tradingSuite.callIssuingETHReq(
		// tradingSuite.USDCTokenAddress,
		tradingSuite.IncAVAXTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingavaxreq",
	)

	// // // new flow unified shield
	// _, err = tradingSuite.callIssuingUnifiedPtokenReq(
	// 	tradingSuite.IncUTAVAXTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// 	"bridgeaggShield",
	// 	6,
	// 	tradingSuite.IncAVAXTokenIDStr,
	// )

	require.Equal(tradingSuite.T(), nil, err)

	// time.Sleep(60 * time.Second)
}

