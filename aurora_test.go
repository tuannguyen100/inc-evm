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
type AuroraTestSuite struct {
	*TradingTestSuite

	PanackeTradeDeployedAddr common.Address
	PanackeRouteContractAddr common.Address


	IncAUDORATokenIDStr 	string
	IncUTAUDORATokenIDStr 	string
	IncUTUSDTTokenIDStr 	string
	IncN5USDTTokenIDStr 	string
	IncN5TUSDTTokenIDStr 	string

	AUDORAAddressStr string
	N5USDTTokenAddress  common.Address
	N5TUSDTTokenAddress common.Address
	WAURORAAddr    common.Address

	// token amounts for tests
	DepositingEther       float64
	AVAXBalanceAfterStep1 *big.Int
	USDTBalanceAfterStep2 *big.Int
	
}

func NewAuroraTestSuite(tradingTestSuite *TradingTestSuite) *AuroraTestSuite {
	return &AuroraTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *AuroraTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	
	// tradingSuite.PanackeTradeDeployedAddr = common.HexToAddress("0xdfa4df8C6b749004FDF725ff5ddE90feb2bBd7D7")
	// tradingSuite.PanackeRouteContractAddr = common.HexToAddress("0x9168badEE3F2517b5de7Da2E22038fE7727c1F4C")
	// tradingSuite.WAURORAAddr = common.HexToAddress("0x9ef7eB75556E6A77001D778f89D573beFEDe38d7")  // local

	tradingSuite.PanackeTradeDeployedAddr = common.HexToAddress("0xA17b90be4A5F79076c770384332515359D2F6A88")
	tradingSuite.PanackeRouteContractAddr = common.HexToAddress("0x26ec2aFBDFdFB972F106100A3deaE5887353d9B9")
	tradingSuite.WAURORAAddr = common.HexToAddress("0x1b6A3d5B5DCdF7a37CFE35CeBC0C4bD28eA7e946")  


	tradingSuite.N5TUSDTTokenAddress = common.HexToAddress("0xCe7F47c42b1E7aC02d90152D56F5417535c48883")
	tradingSuite.N5USDTTokenAddress = common.HexToAddress("0x30fb06E97a6CD370BCE994A88C428F9F3aB6Ec28")

	tradingSuite.IncN5USDTTokenIDStr = "397883f84afacf6433eb09b3f9cfc51f04b3cc632be93733419a0295b6e585af"
	tradingSuite.IncUTUSDTTokenIDStr = "3a526c0fa9abfc3e3e37becc52c5c10abbb7897b0534ad17018e766fc6133590"
	tradingSuite.IncN5TUSDTTokenIDStr = "c8e149001975b9c636d6fb148ad6c33b6dbe9c665d75b0fe245de71f634ffecc"

	tradingSuite.IncAUDORATokenIDStr = "a189f306574f75733e93a260b168d8a335ba3243899254b3a4bf5f4de69a8e71"
	tradingSuite.IncUTAUDORATokenIDStr = "b366fa400c36e6bbcf24ac3e99c90406ddc64346ab0b7ba21e159b83d938812d"

	tradingSuite.DepositingEther = float64(0.0001)

	tradingSuite.AUDORAAddressStr ="0x0000000000000000000000000000000000000000"
	
}

func (tradingSuite *AuroraTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.AURORAClient.Close()
}

func (tradingSuite *AuroraTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *AuroraTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestAuroraTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for aurora test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	auroraSuite := NewAuroraTestSuite(tradingSuite)
	suite.Run(t, auroraSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *AuroraTestSuite) getExpectedAmount(
	path []common.Address,
	srcQty *big.Int,
) *big.Int {
	c, err := pancakeproxy.NewPancakeproxyrouter(tradingSuite.PanackeRouteContractAddr, tradingSuite.AURORAClient)
	require.Equal(tradingSuite.T(), nil, err)
	amounts, err := c.GetAmountsOut(nil, srcQty, path)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("path: %v\n", path)
	fmt.Printf("intput value: %v\n", srcQty.String())
	fmt.Printf("output value: %v\n", amounts[len(amounts)-1].String())
	fmt.Printf("path output : %v\n", amounts)

	return amounts[len(amounts)-1]
}

func (tradingSuite *AuroraTestSuite) executeWithPancake(
	srcQty *big.Int,
	path []common.Address,
	deadline uint,
	isNative bool,
) {
	require.NotEqual(tradingSuite.T(), 0, len(path))

	tradeAbi, err := abi.JSON(strings.NewReader(pancakeproxy.PancakeproxyMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAURORAAddr, tradingSuite.AURORAClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDAURORA)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(10e10)
	auth.GasLimit = uint64(1000000) 
	expectOutputAmount := tradingSuite.getExpectedAmount(path, srcQty)
	input, err := tradeAbi.Pack("trade", path, srcQty, expectOutputAmount, big.NewInt(int64(deadline)), isNative)
	require.Equal(tradingSuite.T(), nil, err)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(tradingSuite.T(), nil, err)
	sourceToken := path[0]
	if path[0].String() == tradingSuite.WAURORAAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.AUDORAAddressStr)
	}
	destToken := path[len(path)-1]
	if path[len(path)-1].String() == tradingSuite.WAURORAAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.AUDORAAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    AURORA_EXECUTE_PREFIX,
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
	if err := wait(tradingSuite.AURORAClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Dapp trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *AuroraTestSuite) CallData(
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

func (tradingSuite *AuroraTestSuite) Test1Aurora() {
	return
	fmt.Println("============ TEST SHIELD UNSHIELD NATIVE AURORA ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting AURORA to pAURORA --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		// "12suy7wB2qYzSdUYPrRu2Wys1ZccRDJVwoTY3TyqJjt5nNMUitutD73S5voK5dFmSRpQytiKHRiGs8CwJjx2ZZCjbSCkWSXDwMBpNpytGvKkE6Tg1MFFRNYqLkM5G1Q5a8BeAu9S8g7Dv4niWoJf",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err := tradingSuite.callIssuingAURORAReq(
		tradingSuite.IncAUDORATokenIDStr,
		txHash.String()[2:],
		"createandsendtxwithissuingaurorareq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)
return
	fmt.Println("------------ STEP 2: burning pAURORA to deposit AURORA to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncAUDORATokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningaurorafordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		// "ca34dc36e2ba7c3b1d1b93adb6dfbd829d9b66914c6a0e9fd9615ec1c80ab10a",
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		"getburnauroraprooffordeposittosc",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	fmt.Println("deposited AURORA: ", deposited)

	fmt.Println("------------ step 3: withdrawing AURORA from SC to pAURORA on Incognito --------------")
	txHash = tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		deposited,
		tradingSuite.AURORAClient,
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		tradingSuite.VaultAURORAAddr,
		AURORA_REQ_WITHDRAW_PREFIX,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)
	_, err = tradingSuite.callIssuingAURORAReq(
		tradingSuite.IncAUDORATokenIDStr,
		txHash.String()[2:],
		"createandsendtxwithissuingaurorareq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)

	fmt.Println("------------ step 4: withdrawing pAURORA from Incognito to AURORA --------------")
	withdrawingPAVAX := big.NewInt(0).Div(deposited, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncAUDORATokenIDStr,
		withdrawingPAVAX,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningaurorarequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getauroraburnproof",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AURORAClient,
	)
	tradingSuite.AVAXBalanceAfterStep1 = bal
	fmt.Println("AVAX balance after step 1: ", tradingSuite.AVAXBalanceAfterStep1)
}

func (tradingSuite *AuroraTestSuite) Test2AuroraToken() {
	return
	fmt.Println("============ TEST SHIELD UNSHIELD TOKEN AURORA ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingUSDT := big.NewInt(3e6)
	burningPUSDT := depositingUSDT

	daibal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.N5USDTTokenAddress,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AURORAClient,
	)
	fmt.Println("usdt balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting USDT to pUSDT --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingUSDT,
		tradingSuite.N5USDTTokenAddress,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err := tradingSuite.callIssuingAURORAReq(
		tradingSuite.IncN5USDTTokenIDStr,
		txHash.String()[2:],
		"createandsendtxwithissuingaurorareq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)
return
	fmt.Println("------------ step 2: burning pUSDT to deposit USDT to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncN5USDTTokenIDStr,
		burningPUSDT,
		pubKeyToAddrStr[2:],
		"createandsendburningaurorafordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		"getburnauroraprooffordeposittosc",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		tradingSuite.N5USDTTokenAddress,
		pubKeyToAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	fmt.Println("deposited usdt: ", deposited)

	fmt.Println("------------ step 3: withdrawing USDT from SC to pUSDT on Incognito --------------")
	txHash = tradingSuite.requestWithdraw(
		tradingSuite.N5USDTTokenAddress.String(),
		deposited,
		tradingSuite.AURORAClient,
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		tradingSuite.VaultAURORAAddr,
		AURORA_REQ_WITHDRAW_PREFIX,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err = tradingSuite.callIssuingAURORAReq(
		tradingSuite.IncN5USDTTokenIDStr,
		txHash.String()[2:],
		"createandsendtxwithissuingaurorareq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)

	fmt.Println("------------ step 4: withdrawing pUSDT from Incognito to USDT --------------")
	withdrawingPMRK := deposited
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncN5USDTTokenIDStr,
		withdrawingPMRK,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningaurorarequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getauroraburnproof",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.N5USDTTokenAddress,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AURORAClient,
	)
	tradingSuite.USDTBalanceAfterStep2 = bal
	fmt.Println("USDT balance after step 2: ", tradingSuite.USDTBalanceAfterStep2)
}

func (tradingSuite *AuroraTestSuite) Test3AuroraUnified() {
	return
	fmt.Println("============ TEST SHIELD UNSHIELD UNIFIED NATIVE AURORA ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err := tradingSuite.callIssuingAUDORAUnifiedPtokenReq(
		tradingSuite.IncUTAUDORATokenIDStr,
		txHash.String()[2:],
		"bridgeaggShield",
		5,
		tradingSuite.IncAUDORATokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(50 * time.Second)
return
	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTAUDORATokenIDStr,
		burningPETH,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncAUDORATokenIDStr,
		true,
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSCV2(
		burningTxID.(string),
		// "ca34dc36e2ba7c3b1d1b93adb6dfbd829d9b66914c6a0e9fd9615ec1c80ab10a",
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		0,
		5,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	fmt.Println("deposited AURORA: ", deposited)

	// tradingSuite.executeWithPancake(
	// 	deposited,
	// 	[]common.Address{
	// 		tradingSuite.WAURORAAddr,
	// 		tradingSuite.N5TUSDTTokenAddress,
	// 	},
	// 	uint(time.Now().Unix()+60000),
	// 	false,
	// )

// return
	fmt.Println("------------ step 3: withdrawing ETH from SC to pETH on Incognito --------------")
	txHash = tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		deposited,
		tradingSuite.AURORAClient,
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		tradingSuite.VaultAURORAAddr,
		AURORA_REQ_WITHDRAW_PREFIX,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err = tradingSuite.callIssuingAUDORAUnifiedPtokenReq(
		tradingSuite.IncUTAUDORATokenIDStr,
		txHash.String()[2:],
		"bridgeaggShield",
		5,
		tradingSuite.IncAUDORATokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)

	fmt.Println("------------ step 4: withdrawing pUSDT from Incognito to USDT --------------")
	burningRes, err = tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTAUDORATokenIDStr,
		burningPETH,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncAUDORATokenIDStr,
		false,
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(40 * time.Second)

	tradingSuite.submitBurnProofForWithdrawalV2(
		burningTxID.(string),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
		0,
		5,
	)

tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.AUDORAAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AVAXClient,
	)
}

func (tradingSuite *AuroraTestSuite) Test4AuroraUnifiedToken() {
	return
	fmt.Println("============ TEST SHIELD UNSHIELD Unified TOKEN AURORA ===========")
	fmt.Println("------------ step 0: declaration & initialization --------------")
	depositingUSDT := big.NewInt(3e5)
	burningPUSDT := depositingUSDT

	daibal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.N5USDTTokenAddress,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AURORAClient,
	)
	fmt.Println("usdt balance of owner: ", daibal)

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ", pubKeyToAddrStr)

	fmt.Println("------------ step 1: porting USDT to pUSDT --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		depositingUSDT,
		tradingSuite.N5USDTTokenAddress,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err := tradingSuite.callIssuingAUDORAUnifiedPtokenReq(
		tradingSuite.IncUTUSDTTokenIDStr,
		txHash.String()[2:],
		"bridgeaggShield",
		5,
		tradingSuite.IncN5USDTTokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)
// return
	fmt.Println("------------ step 2: burning pUSDT to deposit USDT to SC --------------")

	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTUSDTTokenIDStr,
		burningPUSDT,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncN5USDTTokenIDStr,
		true,
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForDepositToSCV2(
		burningTxID.(string),
		// "ca34dc36e2ba7c3b1d1b93adb6dfbd829d9b66914c6a0e9fd9615ec1c80ab10a",
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		0,
		5,
	)
	deposited := tradingSuite.getDepositedBalanceWithParams(
		tradingSuite.N5USDTTokenAddress,
		pubKeyToAddrStr,
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
	)
	fmt.Println("deposited usdt: ", deposited)

// /// test swap
	tradingSuite.executeWithPancake(
		deposited,
		[]common.Address{
			tradingSuite.N5TUSDTTokenAddress,
			tradingSuite.WAURORAAddr,
		},
		uint(time.Now().Unix()+60000),
		true,
	)

// 	time.Sleep(15 * time.Second)

	// deposited = tradingSuite.getDepositedBalanceWithParams(
	// 	common.HexToAddress(tradingSuite.WAURORAAddr),
	// 	pubKeyToAddrStr,
	// 	tradingSuite.VaultAURORAAddr,
	// 	tradingSuite.AURORAClient,
	// )
	// fmt.Println("deposited usdt: ", deposited)

	return

	fmt.Println("------------ step 3: withdrawing USDT from SC to pUSDT on Incognito --------------")
	txHash = tradingSuite.requestWithdraw(
		tradingSuite.N5USDTTokenAddress.String(),
		deposited,
		tradingSuite.AURORAClient,
		big.NewInt(int64(tradingSuite.ChainIDAURORA)),
		tradingSuite.VaultAURORAAddr,
		AURORA_REQ_WITHDRAW_PREFIX,
	)

	fmt.Println("Waiting 50s for 5 blocks confirmation")
	time.Sleep(50 * time.Second)

	_, err = tradingSuite.callIssuingAUDORAUnifiedPtokenReq(
		tradingSuite.IncUTUSDTTokenIDStr,
		txHash.String()[2:],
		"bridgeaggShield",
		5,
		tradingSuite.IncN5USDTTokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(150 * time.Second)

	fmt.Println("------------ step 4: withdrawing pUSDT from Incognito to USDT --------------")
	withdrawingPMRK := deposited
	burningRes, err = tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTUSDTTokenIDStr,
		withdrawingPMRK,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncN5USDTTokenIDStr,
		false,
	)

	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawalV2(
		burningTxID.(string),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultAURORAAddr,
		tradingSuite.AURORAClient,
		tradingSuite.ChainIDAURORA,
		0,
		5,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.N5USDTTokenAddress,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.AURORAClient,
	)
	tradingSuite.USDTBalanceAfterStep2 = bal
	fmt.Println("USDT balance after step 2: ", tradingSuite.USDTBalanceAfterStep2)
}


func (tradingSuite *AuroraTestSuite) TestX1_get_call_data() {
	return
	fmt.Println("===== GET call data =====")
	tradingSuite.CallData(
		big.NewInt(1000000 ),
		[]common.Address{
			common.HexToAddress("0xCe7F47c42b1E7aC02d90152D56F5417535c48883"),
			tradingSuite.WAURORAAddr,
			// common.HexToAddress("0x8711C4728324C9b6264829a2fb92C83c870fd1BE"),
		
		},
		uint(time.Now().Unix()+60000),
		false)
}


func (tradingSuite *AuroraTestSuite) TestX2_trade_AURORA_USDC_new_flow() {
	// return
	fmt.Println("===== TestX2 TRADE AURORA USDC NEW FLOW =====")
	fmt.Println("===== GET call data =====")

	input := tradingSuite.CallData(
		big.NewInt(1000 * 1e9),
		[]common.Address{
			tradingSuite.WAURORAAddr,
			tradingSuite.N5TUSDTTokenAddress,
		},
		uint(time.Now().Unix()+60000),
		false)
	
		fmt.Println(" ===== CALL BURNING ======")
		txhash, err := tradingSuite.callBurningDapp(
			tradingSuite.IncUTAUDORATokenIDStr,
			big.NewInt(1000),
			// "0xcE40cE511A5D084017DBee7e3fF3e455ea32D85c",
			"0x0000000000000000000000000000000000000000",
			tradingSuite.PanackeTradeDeployedAddr.String(),
			"bridgeaggBurnForCall",
			tradingSuite.IncPaymentReceiverStr,
			0,
			tradingSuite.IncAUDORATokenIDStr,
			5,
			input,
			"0xCe7F47c42b1E7aC02d90152D56F5417535c48883",
		)
		require.Equal(tradingSuite.T(), nil, err)
		// return
		time.Sleep(30 * time.Second)
		fmt.Println(" ==== SUBMIT TRADE AURORA USDC ===")
		tx_submit_trade := tradingSuite.submitBurnProofForWithdrawalNewDapp(
			txhash,
			"bridgeaggGetBurnProof",
			tradingSuite.VaultAURORAAddr,
			tradingSuite.AURORAClient,
			tradingSuite.ChainIDAURORA,
			0,
			5,
		)
		
		fmt.Println("Waiting 90s for 5 blocks confirmation")
		time.Sleep(60 * time.Second)
	
		// old shield flow
		 tradingSuite.callIssuingAURORAReq(
			tradingSuite.IncN5USDTTokenIDStr,
			tx_submit_trade.String()[2:],
			"createandsendtxwithissuingaurorareq",
		)
}

func (tradingSuite *AuroraTestSuite) TestX3_trade_USDC_AURORA_new_flow() {
	// return
	fmt.Println("===== TestX3 TRADE USDC AURORA NEW FLOW =====")
	fmt.Println("===== GET call data =====")

	input := tradingSuite.CallData(
		big.NewInt(100000 ),
		[]common.Address{
			tradingSuite.N5TUSDTTokenAddress,
			tradingSuite.WAURORAAddr,
		},
		uint(time.Now().Unix()+60000),
		true)
	
		fmt.Println(" ===== CALL BURNING ======")
		txhash, err := tradingSuite.callBurningDapp(
			tradingSuite.IncN5USDTTokenIDStr,
			big.NewInt(100000),
			// "0xcE40cE511A5D084017DBee7e3fF3e455ea32D85c",
			"0x0000000000000000000000000000000000000000",
			tradingSuite.PanackeTradeDeployedAddr.String(),
			"bridgeaggBurnForCall",
			tradingSuite.IncPaymentReceiverStr,
			0,
			tradingSuite.IncN5USDTTokenIDStr,
			5,
			input,
			"0x0000000000000000000000000000000000000000",
		)
		require.Equal(tradingSuite.T(), nil, err)
		// return
		time.Sleep(30 * time.Second)
		fmt.Println(" ==== SUBMIT TRADE AURORA USDC ===")
		tx_submit_trade := tradingSuite.submitBurnProofForWithdrawalNewDapp(
			txhash,
			"bridgeaggGetBurnProof",
			tradingSuite.VaultAURORAAddr,
			tradingSuite.AURORAClient,
			tradingSuite.ChainIDAURORA,
			0,
			5,
		)
	
		fmt.Println("Waiting 90s for 5 blocks confirmation")
		time.Sleep(60 * time.Second)
	
		// old shield flow
		 tradingSuite.callIssuingAURORAReq(
			// tradingSuite.IncN5TUSDTTokenIDStr,
			tradingSuite.IncAUDORATokenIDStr,
			tx_submit_trade.String()[2:],
			"createandsendtxwithissuingaurorareq",
		)
}
