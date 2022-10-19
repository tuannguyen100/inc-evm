package bridge

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/params"
	pUniswapHelper "github.com/incognitochain/bridge-eth/bridge/puniswaphelper"
	puniswap "github.com/incognitochain/bridge-eth/bridge/puniswapproxy"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/stretchr/testify/suite"
	"math/big"
	"strings"
	"testing"
	"time"


	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/require"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PolygonTestSuite struct {
	*TradingTestSuite

	UniswapDeployedAddr      common.Address
	UniswapRouteContractAddr common.Address
	UNiswapQuoteContractAddr common.Address

	IncBUSDTokenIDStr string
	MATICAddr        common.Address
	WMATICAddr        common.Address
	WETHAddr          common.Address
	DAIAddr 	common.Address
	USDCAddr 	common.Address

	// token amounts for tests
	DepositingEther      float64
	DAIBalanceAfterStep1 *big.Int
	MRKBalanceAfterStep2 *big.Int
}

func NewPolygonTestSuite(tradingTestSuite *TradingTestSuite) *PolygonTestSuite {
	return &PolygonTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

const (
	LOW    = 500
	MEDIUM = 3000
	HIGH   = 10000
)

const MAX_PERCENT = 10000


// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *PolygonTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")
	// Polygon testnet env
	tradingSuite.IncBUSDTokenIDStr = "0000000000000000000000000000000000000000000000000000000000000062"
	tradingSuite.UniswapDeployedAddr = common.HexToAddress("0xAe85BB3D2ED209736E4d236DcE24624EA1A04249")  //0xB806dC43E5494845795Ca75BA49406cD0FFEA2e0//0x0C61C7F99DeE0270a5934F1520d109dc44A51d6c//0xB806dC43E5494845795Ca75BA49406cD0FFEA2e0
	tradingSuite.UniswapRouteContractAddr = common.HexToAddress("0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45")
	tradingSuite.UNiswapQuoteContractAddr = common.HexToAddress("0x61ffe014ba17989e743c5f6cb21bf9697530b21e")

	// tokens
	tradingSuite.MATICAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")
	tradingSuite.WMATICAddr = common.HexToAddress("0x9c3c9283d3e44854697cd22d3faa240cfb032889")
	tradingSuite.WETHAddr = common.HexToAddress("0xa6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa")
	tradingSuite.DAIAddr = common.HexToAddress( "0x001b3b4d0f3714ca98ba10f6042daebf0b1b7b6f")
	tradingSuite.USDCAddr = common.HexToAddress( "0x2058A9D7613eEE744279e3856Ef0eAda5FCbaA7e")
	tradingSuite.DepositingEther = float64(0.123456789)
}

func (tradingSuite *PolygonTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.PLGClient.Close()
}

func (tradingSuite *PolygonTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *PolygonTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPolygonTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	uniswapTradingSuite := NewPolygonTestSuite(tradingSuite)
	suite.Run(t, uniswapTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *PolygonTestSuite) getExpectedAmount(
	srcQty *big.Int,
	paths []common.Address,
	fees []int64,
) *big.Int {
	c, err := pUniswapHelper.NewPUniswapHelper(tradingSuite.UNiswapQuoteContractAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	var amountOut *big.Int
	var amountIn *big.Int
	if len(fees) > 1 {
		inputParam := &pUniswapHelper.IUinswpaHelperExactInputParams{
			Path:     tradingSuite.buildPath(paths, fees),
			AmountIn: srcQty,
		}
		result, err := c.QuoteExactInput(nil, inputParam.Path, inputParam.AmountOutMinimum)
		require.Equal(tradingSuite.T(), nil, err)
		amountIn = inputParam.AmountIn
		amountOut = result.AmountOut
	} else {
		inputSingleParam := pUniswapHelper.IUinswpaHelperQuoteExactInputSingleParams{
			TokenIn:           paths[0],
			TokenOut:          paths[len(paths)-1],
			Fee:               big.NewInt(fees[0]),
			AmountIn:          srcQty,
			SqrtPriceLimitX96: big.NewInt(0),
		}
		result, err := c.QuoteExactInputSingle(nil, inputSingleParam)
		require.Equal(tradingSuite.T(), nil, err)
		amountIn = inputSingleParam.AmountIn
		amountOut = result.AmountOut
	}
	fmt.Printf("intput value: %v\n", amountIn.String())
	fmt.Printf("output value: %v\n", amountOut.String())
	return amountOut
}


func (tradingSuite *PolygonTestSuite) executeWithPUniswapMultiTrade(
	srcQty *big.Int,
	paths [][]common.Address,
	fees [][]int64,
	percents []int64,
	isNative bool,
	deadline int64,
	isTestMultiPath []bool,
) {
	require.Equal(tradingSuite.T(), true, len(fees) != 0)
	require.Equal(tradingSuite.T(), len(paths), len(fees))
	require.Equal(tradingSuite.T(), len(percents), len(fees))

	tradeAbi1, err := abi.JSON(strings.NewReader(pUniswapHelper.PUniswapHelperMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)
	tradeAbi, err := abi.JSON(strings.NewReader(puniswap.PuniswapMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(3e10)
	var calldata [][]byte
	for i := 0; i < len(fees); i++ {
		var agr interface{}
		amount := big.NewInt(0).Div(big.NewInt(0).Mul(srcQty, big.NewInt(percents[i])), big.NewInt(MAX_PERCENT))
		expectOutputAmount := tradingSuite.getExpectedAmount(
			amount,
			paths[i],
			fees[i],
		)
		recipient := tradingSuite.VaultPLGAddr
		if isNative && bytes.Compare(paths[i][len(paths[i])-1].Bytes(), tradingSuite.WMATICAddr.Bytes()) == 0 {
			recipient = tradingSuite.UniswapDeployedAddr
		}
		var input []byte
		if len(fees[i]) > 1 || isTestMultiPath[i] {
			agr = &pUniswapHelper.IUinswpaHelperExactInputParams{
				Path:             tradingSuite.buildPath(paths[i], fees[i]),
				Recipient:        recipient,
				AmountIn:         amount,
				AmountOutMinimum: expectOutputAmount,
			}
			input, err = tradeAbi1.Pack("exactInput", agr)
			require.Equal(tradingSuite.T(), nil, err)
		} else {
			agr = &pUniswapHelper.IUinswpaHelperExactInputSingleParams{
				TokenIn:           paths[i][0],
				TokenOut:          paths[i][len(paths[i])-1],
				Fee:               big.NewInt(fees[i][0]),
				Recipient:         recipient,
				AmountIn:          amount,
				SqrtPriceLimitX96: big.NewInt(0),
				AmountOutMinimum:  expectOutputAmount,
			}
			input, err = tradeAbi1.Pack("exactInputSingle", agr)
			require.Equal(tradingSuite.T(), nil, err)
		}
		calldata = append(calldata, input)
	}
	input, err := tradeAbi.Pack("multiTrades", big.NewInt(deadline), calldata, paths[0][0], paths[0][len(paths[0])-1], srcQty, isNative)
	require.Equal(tradingSuite.T(), nil, err)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(tradingSuite.T(), nil, err)
	sourceToken := paths[0][0]
	// todo: compare pTokenID
	if paths[0][0].String() == tradingSuite.WMATICAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	destToken := paths[0][len(paths[0])-1]
	if paths[0][len(paths[0])-1].String() == tradingSuite.WMATICAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    PLG_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.UniswapDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.UniswapDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.PLGClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("pUniswap trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *PolygonTestSuite) CallData(
	srcQty *big.Int,
	paths []common.Address,
	fees []int64,
	isNative bool,
	isTestMultiPath bool,
) []byte {
	require.Equal(tradingSuite.T(), true, len(fees) != 0)
	require.Equal(tradingSuite.T(), len(paths), len(fees)+1)

	tradeAbi, err := abi.JSON(strings.NewReader(puniswap.PuniswapMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	var agr interface{}
	expectOutputAmount := tradingSuite.getExpectedAmount(
		srcQty,
		paths,
		fees,
	)
	recipient := tradingSuite.VaultPLGAddr
	if isNative && bytes.Compare(paths[len(paths)-1].Bytes(), tradingSuite.WMATICAddr.Bytes()) == 0 {
		recipient = tradingSuite.UniswapDeployedAddr
	}
	var input []byte
	if len(fees) > 1 || isTestMultiPath {
		agr = &pUniswapHelper.IUinswpaHelperExactInputParams{
			Path:             tradingSuite.buildPath(paths, fees),
			Recipient:        recipient,
			AmountIn:         srcQty,
			AmountOutMinimum: expectOutputAmount,
		}
		input, err = tradeAbi.Pack("tradeInput", agr, isNative)
		fmt.Println("Call Data : ",common.Bytes2Hex(input))
		return input
	} else {
		agr = &pUniswapHelper.IUinswpaHelperExactInputSingleParams{
			TokenIn:           paths[0],
			TokenOut:          paths[len(paths)-1],
			Fee:               big.NewInt(fees[0]),
			Recipient:         recipient,
			AmountIn:          srcQty,
			SqrtPriceLimitX96: big.NewInt(0),
			AmountOutMinimum:  expectOutputAmount,
		}
		input, err = tradeAbi.Pack("tradeInputSingle", agr, isNative)
		fmt.Println("Call Data : ",common.Bytes2Hex(input))
		return input
	}
}


func (tradingSuite *PolygonTestSuite) CallDataMultiTrade(
	srcQty *big.Int,
	paths [][]common.Address,
	fees [][]int64,
	percents []int64,
	isNative bool,
	deadline int64,
	isTestMultiPath []bool,
) []byte {
	require.Equal(tradingSuite.T(), true, len(fees) != 0)
	require.Equal(tradingSuite.T(), len(paths), len(fees))
	require.Equal(tradingSuite.T(), len(percents), len(fees))

	tradeAbi1, err := abi.JSON(strings.NewReader(pUniswapHelper.PUniswapHelperMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)
	tradeAbi, err := abi.JSON(strings.NewReader(puniswap.PuniswapMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)


	var calldata [][]byte
	for i := 0; i < len(fees); i++ {
		var agr interface{}
		amount := big.NewInt(0).Div(big.NewInt(0).Mul(srcQty, big.NewInt(percents[i])), big.NewInt(MAX_PERCENT))
		expectOutputAmount := tradingSuite.getExpectedAmount(
			amount,
			paths[i],
			fees[i],
		)
		recipient := tradingSuite.VaultPLGAddr
		if isNative && bytes.Compare(paths[i][len(paths[i])-1].Bytes(), tradingSuite.WMATICAddr.Bytes()) == 0 {
			recipient = tradingSuite.UniswapDeployedAddr
		}
		fmt.Println("recipient : ",recipient)
		var input []byte
		if len(fees[i]) > 1 || isTestMultiPath[i] {
			agr = &pUniswapHelper.IUinswpaHelperExactInputParams{
				Path:             tradingSuite.buildPath(paths[i], fees[i]),
				Recipient:        recipient,
				AmountIn:         amount,
				AmountOutMinimum: expectOutputAmount,
			}
			input, err = tradeAbi1.Pack("exactInput", agr)
			require.Equal(tradingSuite.T(), nil, err)
		} else {
			agr = &pUniswapHelper.IUinswpaHelperExactInputSingleParams{
				TokenIn:           paths[i][0],
				TokenOut:          paths[i][len(paths[i])-1],
				Fee:               big.NewInt(fees[i][0]),
				Recipient:         recipient,
				AmountIn:          amount,
				SqrtPriceLimitX96: big.NewInt(0),
				AmountOutMinimum:  expectOutputAmount,
			}
			input, err = tradeAbi1.Pack("exactInputSingle", agr)
			require.Equal(tradingSuite.T(), nil, err)
		}
		calldata = append(calldata, input)
	}
	input, err := tradeAbi.Pack("multiTrades", big.NewInt(deadline), calldata, paths[0][0], paths[0][len(paths[0])-1], srcQty, isNative)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("Call Data : ",common.Bytes2Hex(input))
	return input
}

func (tradingSuite *PolygonTestSuite) executeWithPUniswap(
	srcQty *big.Int,
	paths []common.Address,
	fees []int64,
	isNative bool,
	isTestMultiPath bool,
) {
	require.Equal(tradingSuite.T(), true, len(fees) != 0)
	require.Equal(tradingSuite.T(), len(paths), len(fees)+1)

	tradeAbi, err := abi.JSON(strings.NewReader(puniswap.PuniswapMetaData.ABI))
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDPLG)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(3e10)
	// auth.GasLimit = 1000000
	var agr interface{}
	expectOutputAmount := tradingSuite.getExpectedAmount(
		srcQty,
		paths,
		fees,
	)
	recipient := tradingSuite.VaultPLGAddr
	if isNative && bytes.Compare(paths[len(paths)-1].Bytes(), tradingSuite.WMATICAddr.Bytes()) == 0 {
		recipient = tradingSuite.UniswapDeployedAddr
	}
	var input []byte
	if len(fees) > 1 || isTestMultiPath {
		agr = &pUniswapHelper.IUinswpaHelperExactInputParams{
			Path:             tradingSuite.buildPath(paths, fees),
			Recipient:        recipient,
			AmountIn:         srcQty,
			AmountOutMinimum: expectOutputAmount,
		}
		input, err = tradeAbi.Pack("tradeInput", agr, isNative)
		fmt.Println("Call Data : ",common.Bytes2Hex(input))
	} else {
		agr = &pUniswapHelper.IUinswpaHelperExactInputSingleParams{
			TokenIn:           paths[0],
			TokenOut:          paths[len(paths)-1],
			Fee:               big.NewInt(fees[0]),
			Recipient:         recipient,
			AmountIn:          srcQty,
			SqrtPriceLimitX96: big.NewInt(0),
			AmountOutMinimum:  expectOutputAmount,
		}
		input, err = tradeAbi.Pack("tradeInputSingle", agr, isNative)
		fmt.Println("Call Data : ",common.Bytes2Hex(input))
	}
	require.Equal(tradingSuite.T(), nil, err)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(tradingSuite.T(), nil, err)
	sourceToken := paths[0]
	// todo: compare pTokenID
	if paths[0].String() == tradingSuite.WMATICAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	destToken := paths[len(paths)-1]
	if paths[len(paths)-1].String() == tradingSuite.WMATICAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    PLG_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.UniswapDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.UniswapDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.PLGClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("pUniswap trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *PolygonTestSuite) buildPath(paths []common.Address, fees []int64) []byte {
	var temp []byte
	for i := 0; i < len(fees); i++ {
		temp = append(temp, paths[i].Bytes()...)
		fee, err := hex.DecodeString(fmt.Sprintf("%06x", fees[i]))
		require.Equal(tradingSuite.T(), nil, err)
		temp = append(temp, fee...)
	}
	temp = append(temp, paths[len(paths)-1].Bytes()...)

	return temp
}

func (tradingSuite *PolygonTestSuite) Test1TradeEthForDAIWithPancake() {
return
	fmt.Println("============ TEST SHIELD UNSHIELD POLYGON ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("pubKeyToAddrStr: ",pubKeyToAddrStr)

	// privateKeyToAddrStr := crypto.GenerateKey(tradingSuite.GeneratedPrivKeyForSC).Hex()
	// fmt.Println("privateKeyToAddrStr: ",privateKeyToAddrStr)

	fmt.Println("------------ STEP 1: porting MATIC to pMATIC --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, txHash)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("35a13f1d5cf5e129e786a36ddf3ef13fc163309534ca77dff628827d83c8618f"))
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncMATICTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
// return
	fmt.Println("------------ STEP 2: burning pBNB to deposit BNB to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncMATICTokenIDStr,
		burningPETH,
		// "920e561d9FC5D843371d0426C74F6065c1B627DE",
		pubKeyToAddrStr[2:],
		"createandsendburningplgfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(40 * time.Second)
// return
	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		"getburnplgprooffordeposittosc",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	deposited := tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("address own asset %v \n", pubKeyToAddrStr)

	fmt.Println("------------ step 3: execute trade MATIC for USDT through Pancake --------------")
	tradingSuite.executeWithPUniswap(
		deposited,
		[]common.Address{tradingSuite.WMATICAddr, tradingSuite.WETHAddr},
		[]int64{LOW},
		false,
		false,
	)
	time.Sleep(15 * time.Second)
	daiTraded := tradingSuite.getDepositedBalancePLG(
		tradingSuite.WETHAddr,
		pubKeyToAddrStr,
	)

	testCrossPoolTrade := big.NewInt(0).Div(daiTraded, big.NewInt(4))
	tradingSuite.executeWithPUniswap(
		testCrossPoolTrade,
		[]common.Address{tradingSuite.WETHAddr, tradingSuite.WMATICAddr},
		[]int64{LOW},
		true,
		false,
	)

	// tradingSuite.executeWithPUniswap(
	// 	testCrossPoolTrade,
	// 	[]common.Address{tradingSuite.WETHAddr, tradingSuite.DAIAddr},
	// 	[]int64{MEDIUM},
	// 	false,
	// 	true,
	// )

	tradingSuite.executeWithPUniswapMultiTrade(
		testCrossPoolTrade,
		[][]common.Address{{tradingSuite.WETHAddr, tradingSuite.WMATICAddr}, {tradingSuite.WETHAddr, tradingSuite.WMATICAddr}},
		[][]int64{{LOW}, {MEDIUM}},
		[]int64{70, 30},
		false,
		time.Now().Unix()+60000,
		[]bool{false, false},
	)

	daiTraded = tradingSuite.getDepositedBalancePLG(
		tradingSuite.WETHAddr,
		pubKeyToAddrStr,
	)

	fmt.Println("weth: ", daiTraded)

	fmt.Println("------------ step 3: withdrawing WETH from SC to pWETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.WETHAddr.String(),
		daiTraded,
		tradingSuite.PLGClient,
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		tradingSuite.VaultPLGAddr,
		PLG_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncWETHTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)
// return

// fmt.Println("------------ STEP 4.2: burning pBNB to deposit BNB to SC --------------")
// 	// make a burn tx to incognito chain as a result of deposit to SC
// 	withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1e9))
// 	fmt.Printf("IncWETHAddr  %v \n", withdrawingPDAI)

// 	burningRes, err = tradingSuite.callBurningPToken(
// 		tradingSuite.IncWETHTokenIDStr,
// 		withdrawingPDAI,
// 		// "AA3D33840E8d051E5712b643f2527ae027413F3f",
// 		pubKeyToAddrStr[2:],
// 		"createandsendburningplgfordeposittoscrequest",
// 	)

// 	require.Equal(tradingSuite.T(), nil, err)
// 	burningTxID, found = burningRes["TxID"]
// 	require.Equal(tradingSuite.T(), true, found)
// 	time.Sleep(120 * time.Second)
// // return
// 	tradingSuite.submitBurnProofForDepositToSC(
// 		burningTxID.(string),
// 		big.NewInt(int64(tradingSuite.ChainIDPLG)),
// 		"getburnplgprooffordeposittosc",
// 		tradingSuite.VaultPLGAddr,
// 		tradingSuite.PLGClient,
// 	)
// 	daiTraded = tradingSuite.getDepositedBalancePLG(
// 		tradingSuite.WETHAddr,
// 		pubKeyToAddrStr,
// 	)

// 	fmt.Printf("WETHAddr in SC before %v \n", daiTraded)

// 	deposited = tradingSuite.getDepositedBalancePLG(
// 		common.HexToAddress(tradingSuite.EtherAddressStr),
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Printf("MATIC before trade :  %v \n", deposited)

// 	tradingSuite.executeWithPUniswap(
// 		daiTraded,
// 		[]common.Address{tradingSuite.WETHAddr,tradingSuite.WMATICAddr },
// 		[]int64{LOW},
// 		false,
// 		false,
// 	)
// 	time.Sleep(15 * time.Second)

// 	deposited = tradingSuite.getDepositedBalancePLG(
// 		common.HexToAddress(tradingSuite.EtherAddressStr),
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Printf("MATIC after trade :  %v \n", deposited)

// 	daiTraded = tradingSuite.getDepositedBalancePLG(
// 		tradingSuite.WETHAddr,
// 		pubKeyToAddrStr,
// 	)
// 	fmt.Printf("WETHAddr in SC %v \n", daiTraded)

// return
	fmt.Println("------------ step 4: withdrawing pWETH from Incognito to WETH --------------")
	withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1e9))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncWETHTokenIDStr,
		withdrawingPDAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningplgrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getplgburnproof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.WETHAddr,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.PLGClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("WETH balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}


func (tradingSuite *PolygonTestSuite) Test3TradeEthForDAIWithPancake() {
	return
	fmt.Println("============ TEST SHIELD UNSHIELD POLYGON ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting MATIC to pMATIC --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		// "12st42eE1ZZHfdjoKEq9Hq5EQgdQ38A8WTRtumakxWNYi7UBd76D6NEsy6xuFD7hwcu7iWqeJys6Lo282Mqvy8dcjamgX2N68DBaCnCVYuQecZ9ipdTBiWiBSPY4xJd2t2YQnrPiTHMFEFfYtTxs",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	time.Sleep(15 * time.Second)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("cf8522fcbc691b8bc9217e48a4be7cd3567d98239014f10adc4a3c66de7bcde4"))
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(180 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncMATICTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)

	// return
	time.Sleep(120 * time.Second)


// return
	fmt.Println("------------ STEP 2: burning pBNB to deposit BNB to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncMATICTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningplgfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)
	
	// return

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		"getburnplgprooffordeposittosc",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	deposited := tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("deposited %v \n", deposited)


	fmt.Println("------------ step 3: withdrawing MATIC from SC to pMATIC on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		deposited,
		tradingSuite.PLGClient,
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		tradingSuite.VaultPLGAddr,
		PLG_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	deposited = tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("affter burn %v \n", deposited)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(180 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncMATICTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 4: withdrawing pMATIC from Incognito to MATIC --------------")
	// withdrawingPDAI := big.NewInt(0).Div(deposited, big.NewInt(1e9))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncMATICTokenIDStr,
		burningPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningplgrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(180 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getplgburnproof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.PLGClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("USDT balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}

func (tradingSuite *PolygonTestSuite) Test4LINK() {
// return
	fmt.Println("============ TEST 4 SHIELD UNSHIELD LINK on POLYGON network ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting LINK to pLINK --------------")
	txHash := tradingSuite.depositERC20ToBridge(
		tradeAmount,
		common.HexToAddress(tradingSuite.USDTAddressStr),
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
	)
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, txHash)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("0x684c4e448b2e7660c6ef3d79078947ae2973681a8305e9e68f3dd27829353616"))
	
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(200 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(50 * time.Second)
// return
	fmt.Println("------------ STEP 2: burning pLINK to deposit LINK to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningplgfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(50 * time.Second)
	

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		"getburnplgprooffordeposittosc",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	deposited := tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.USDTAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("deposited %v \n", deposited)
// return
	// // fmt.Println("------------ step 3: execute trade BNB for USDT through Pancake --------------")
	// // tradingSuite.executeWithPancake(
	// // 	deposited,
	// // 	[]common.Address{
	// // 		tradingSuite.WBNBAddr,
	// // 		tradingSuite.WBUSDAddr,
	// // 		tradingSuite.WUSDTAddr,
	// // 	},
	// // 	uint(time.Now().Unix()+60000),
	// // 	false,
	// // )
	// // time.Sleep(15 * time.Second)
	// // daiTraded := tradingSuite.getDepositedBalancePLG(
	// // 	tradingSuite.WUSDTAddr,
	// // 	pubKeyToAddrStr,
	// // )
	// // fmt.Println("usdtTraded: ", daiTraded)

	fmt.Println("------------ step 3: withdrawing LINK from SC to pLINK on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.USDTAddressStr,
		deposited,
		tradingSuite.PLGClient,
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		tradingSuite.VaultPLGAddr,
		PLG_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	deposited = tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("after burn %v \n", deposited)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(200 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncUSDTTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(50 * time.Second)

	fmt.Println("------------ step 4: withdrawing pLINK from Incognito to LINK --------------")
	// withdrawingPDAI := big.NewInt(0).Div(deposited, big.NewInt(1e9))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncUSDTTokenIDStr,
		burningPETH,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningplgrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getplgburnproof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.USDTAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.PLGClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("USDT balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}

func (tradingSuite *PolygonTestSuite) Test5UT_MATIC() {
	return
	fmt.Println("============ TEST 5 SHIELD UNSHIELD UT MATIC ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting UT_MATIC to UT_pMATIC --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	time.Sleep(15 * time.Second)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("cf8522fcbc691b8bc9217e48a4be7cd3567d98239014f10adc4a3c66de7bcde4"))
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(250 * time.Second)
	txhashInC, err := tradingSuite.callIssuingUnifiedPtokenReq(
		tradingSuite.IncUTMATICTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		3,
		tradingSuite.IncEtherTokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(80 * time.Second)
	require.Equal(tradingSuite.T(), 2, tradingSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr,tradingSuite.IncUTMATICTokenIDStr)
	fmt.Println("[INFO] UT MATIC balance incognito after issuing step 1 : ", balpBNBS1)
	// // return

	fmt.Println("------------ STEP 2: burning pBNB to deposit BNB to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTMATICTokenIDStr,
		burningPETH,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncMATICTokenIDStr,
		true,
	)

	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(80 * time.Second)
	
	tradingSuite.submitBurnProofForDepositToSCV2(
		burningTxID.(string),
		// "e55841caded175ed168c09c2dc2e4ae6c8ba7bbcfb1475216b70c6185dba9371",
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		0,
		3,

	)
	deposited := tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("deposited %v \n", deposited)


	fmt.Println("------------ step 3: withdrawing MATIC from SC to pMATIC on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.EtherAddressStr,
		deposited,
		tradingSuite.PLGClient,
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		tradingSuite.VaultPLGAddr,
		PLG_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	deposited = tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("affter burn %v \n", deposited)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(250 * time.Second)
	txhashInC, err  =  tradingSuite.callIssuingUnifiedPtokenReq(
		tradingSuite.IncUTMATICTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		3,
		tradingSuite.IncMATICTokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(90 * time.Second)
	require.Equal(tradingSuite.T(), 2, tradingSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS3, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr,tradingSuite.IncUTMATICTokenIDStr)
	fmt.Println("[INFO] UT MATIC balance incognito after issuing step 3 : ", balpBNBS3)

	fmt.Println("------------ step 4: withdrawing pMATIC from Incognito to MATIC --------------")
	// withdrawingPDAI := big.NewInt(0).Div(deposited, big.NewInt(1e9))
	burningRes, err = tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTMATICTokenIDStr,
		burningPETH,
		burningPETH,
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncMATICTokenIDStr,
		false,
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(90 * time.Second)

	tradingSuite.submitBurnProofForWithdrawalV2(
		burningTxID.(string),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
		0,
		3,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.PLGClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("USDT balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}

func (tradingSuite *PolygonTestSuite) Test6UT_MATIC_CrossShard() {
	return
	fmt.Println("============ TEST 6 SHIELD UT MATIC CrossShard ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	fmt.Println("------------ STEP 1: porting UT_MATIC to UT_pMATIC --------------")

	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentReceiverStr,
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	time.Sleep(15 * time.Second)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("cf8522fcbc691b8bc9217e48a4be7cd3567d98239014f10adc4a3c66de7bcde4"))
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(250 * time.Second)
	txhashInC, err := tradingSuite.callIssuingUnifiedPtokenReq(
		tradingSuite.IncUTMATICTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		3,
		tradingSuite.IncMATICTokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(80 * time.Second)
	require.Equal(tradingSuite.T(), 2, tradingSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivaKeyReceiverStr,tradingSuite.IncUTMATICTokenIDStr)
	fmt.Println("[INFO] UT MATIC balance incognito after issuing step 1 : ", balpBNBS1)
}

func (tradingSuite *PolygonTestSuite) Test7UT_LINK() {
	return
		fmt.Println("============ TEST 7 SHIELD UNSHIELD Unified LINK Token  ===========")
		fmt.Println("------------ STEP 0: declaration & initialization --------------")
		tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
		burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))
	
		pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
		fmt.Println("------------ STEP 1: porting LINK to pLINK --------------")
		txHash := tradingSuite.depositERC20ToBridge(
			tradeAmount,
			common.HexToAddress(tradingSuite.USDTAddressStr),
			tradingSuite.IncPaymentAddrStr,
			tradingSuite.VaultPLGAddr,
			tradingSuite.PLGClient,
			tradingSuite.ChainIDPLG,
		)
		time.Sleep(15 * time.Second)
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, txHash)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	// return
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)

	// _, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("0xabe22e9c1da33ffcb980e38cb56cbe7dbea9c4a93635aebabca2541455ac0028"))
		txhashInC, err := tradingSuite.callIssuingUnifiedPtokenReq(
			tradingSuite.IncUTUSDTTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
			"bridgeaggShield",
			3,
			tradingSuite.IncUSDTTokenIDStr,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(50 * time.Second)
	
	require.Equal(tradingSuite.T(), 2, tradingSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr,tradingSuite.IncUTUSDTTokenIDStr)
	fmt.Println("[INFO] UT pLINK balance incognito after issuing step 1 : ", balpBNBS1)
	// return
		fmt.Println("------------ STEP 2: burning pLINK to deposit LINK to SC --------------")
		// make a burn tx to incognito chain as a result of deposit to SC
		burningRes, err := tradingSuite.callBurningUnifiedPToken(
			tradingSuite.IncUTUSDTTokenIDStr,
			burningPETH,
			burningPETH,
			pubKeyToAddrStr[2:],
			"bridgeaggUnshield",
			tradingSuite.IncPaymentAddrStr,
			0,
			tradingSuite.IncUSDTTokenIDStr,
			true,
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found := burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(90 * time.Second)
		
		tradingSuite.submitBurnProofForDepositToSCV2(
			burningTxID.(string),
			big.NewInt(int64(tradingSuite.ChainIDPLG)),
			"bridgeaggGetBurnProof",
			tradingSuite.VaultPLGAddr,
			tradingSuite.PLGClient,
			0,
			3,
		)
		deposited := tradingSuite.getDepositedBalancePLG(
			common.HexToAddress(tradingSuite.USDTAddressStr),
			pubKeyToAddrStr,
		)
		fmt.Printf("deposited %v \n", deposited)
	
		fmt.Println("------------ step 3: withdrawing LINK from SC to pLINK on Incognito --------------")
		txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
			tradingSuite.USDTAddressStr,
			deposited,
			tradingSuite.PLGClient,
			big.NewInt(int64(tradingSuite.ChainIDPLG)),
			tradingSuite.VaultPLGAddr,
			PLG_REQ_WITHDRAW_PREFIX,
		)
		time.Sleep(45 * time.Second)
	
		deposited = tradingSuite.getDepositedBalancePLG(
			common.HexToAddress(tradingSuite.EtherAddressStr),
			pubKeyToAddrStr,
		)
		fmt.Printf("affter burn %v \n", deposited)
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, txHashByEmittingWithdrawalReq)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(250 * time.Second)
		txhashInC, err = tradingSuite.callIssuingUnifiedPtokenReq(
			tradingSuite.IncUTUSDTTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
			"bridgeaggShield",
			3,
			tradingSuite.IncUSDTTokenIDStr,
		)
		require.Equal(tradingSuite.T(), nil, err)
		time.Sleep(80 * time.Second)

		require.Equal(tradingSuite.T(), 2, tradingSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
		balpBNBS3, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr,tradingSuite.IncUTUSDTTokenIDStr)
		fmt.Println("[INFO] UT pLINK balance incognito after issuing step 3 : ", balpBNBS3)
	
		fmt.Println("------------ step 4: withdrawing pLINK from Incognito to LINK --------------")
		// withdrawingPDAI := big.NewInt(0).Div(deposited, big.NewInt(1e9))
		burningRes, err = tradingSuite.callBurningUnifiedPToken(
			tradingSuite.IncUTUSDTTokenIDStr,
			burningPETH,
			burningPETH,
			pubKeyToAddrStr[2:],
			"bridgeaggUnshield",
			tradingSuite.IncPaymentAddrStr,
			0,
			tradingSuite.IncUSDTTokenIDStr,
			false,
		)
		require.Equal(tradingSuite.T(), nil, err)
		burningTxID, found = burningRes["TxID"]
		require.Equal(tradingSuite.T(), true, found)
		time.Sleep(90 * time.Second)
	
		tradingSuite.submitBurnProofForWithdrawalV2(
			burningTxID.(string),
			// "getplgburnproof",
			"bridgeaggGetBurnProof",
			tradingSuite.VaultPLGAddr,
			tradingSuite.PLGClient,
			tradingSuite.ChainIDPLG,
			0,
			3,
		)
	
		bal := tradingSuite.getBalanceOnETHNet(
			common.HexToAddress(tradingSuite.USDTAddressStr),
			common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
			tradingSuite.PLGClient,
		)
		tradingSuite.DAIBalanceAfterStep1 = bal
		fmt.Println("USDT balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}
	
func (tradingSuite *PolygonTestSuite) Test8UT_Trade() {
	  return
	fmt.Println("============ TEST 8 SHIELD UNSHIELD Trade UT POLYGON ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting UT_MATIC to UT_pMATIC --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
	)
	time.Sleep(5 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, txHash)
	// _, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("eaa991123f1ea30a2401499d654d20013946bfb489bca420b5f7d3996dc1b6ac"))

	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	txhashInC, err := tradingSuite.callIssuingUnifiedPtokenReq(
		tradingSuite.IncUTMATICTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		3,
		tradingSuite.IncMATICTokenIDStr,
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	require.Equal(tradingSuite.T(), 2, tradingSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := tradingSuite.getBalanceTokenIncAccount(tradingSuite.IncPrivKeyStr,tradingSuite.IncUTMATICTokenIDStr)
	fmt.Println("[INFO] UT MATIC balance incognito after issuing step 1 : ", balpBNBS1)
	// return


	fmt.Println("------------ STEP 2: burning pMATIC to deposit MATIC to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningUnifiedPToken(
		tradingSuite.IncUTMATICTokenIDStr,
		burningPETH,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		tradingSuite.IncPaymentAddrStr,
		0,
		tradingSuite.IncMATICTokenIDStr,
		true,
	)

	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(60 * time.Second)
	
	tradingSuite.submitBurnProofForDepositToSCV2(
		burningTxID.(string),
		// "c5c3d1da7c28a4ddf0c15ccce4151869d8749d4421f74692056feeade925ed2a",
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		"bridgeaggGetBurnProof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		0,
		3,
	)
	deposited := tradingSuite.getDepositedBalancePLG(
		common.HexToAddress(tradingSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("deposited %v \n", deposited)


	fmt.Println("------------ step 3: execute trade MATIC for ETH through Pancake --------------")
	tradingSuite.executeWithPUniswap(
		deposited,
		[]common.Address{tradingSuite.WMATICAddr, tradingSuite.WETHAddr},
		[]int64{LOW},
		false,
		false,
	)
	time.Sleep(15 * time.Second)
	daiTraded := tradingSuite.getDepositedBalancePLG(
		tradingSuite.WETHAddr,
		pubKeyToAddrStr,
	)

	testCrossPoolTrade := big.NewInt(0).Div(daiTraded, big.NewInt(4))
	tradingSuite.executeWithPUniswap(
		testCrossPoolTrade,
		[]common.Address{tradingSuite.WETHAddr, tradingSuite.WMATICAddr},
		[]int64{LOW},
		true,
		false,
	)


	tradingSuite.executeWithPUniswapMultiTrade(
		testCrossPoolTrade,
		[][]common.Address{{tradingSuite.WETHAddr, tradingSuite.WMATICAddr}, {tradingSuite.WETHAddr, tradingSuite.WMATICAddr}},
		[][]int64{{LOW}, {MEDIUM}},
		[]int64{70, 30},
		false,
		time.Now().Unix()+60000,
		[]bool{false, false},
	)

	daiTraded = tradingSuite.getDepositedBalancePLG(
		tradingSuite.WETHAddr,
		pubKeyToAddrStr,
	)

	fmt.Println("weth: ", daiTraded)
fmt.Println("------------ step 3: withdrawing WETH from SC to pWETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.WETHAddr.String(),
		daiTraded,
		tradingSuite.PLGClient,
		big.NewInt(int64(tradingSuite.ChainIDPLG)),
		tradingSuite.VaultPLGAddr,
		PLG_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.PLGHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncWETHTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 4: withdrawing pWETH from Incognito to WETH --------------")
	withdrawingPDAI := big.NewInt(0).Div(daiTraded, big.NewInt(1e9))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncWETHTokenIDStr,
		withdrawingPDAI,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningplgrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getplgburnproof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		tradingSuite.WETHAddr,
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.PLGClient,
	)
	tradingSuite.DAIBalanceAfterStep1 = bal
	fmt.Println("WETH balance after step 1: ", tradingSuite.DAIBalanceAfterStep1)
}

func (tradingSuite *PolygonTestSuite) TestX1_submit_proof_to_sc() {
return		
tradingSuite.submitBurnProofForWithdrawalNewDapp(
	"edca8d73c1c1ee67831a233d5c36e18e845578ff16f3a3b3828eb1f75c22d9f9",
	"bridgeaggGetBurnProof",
	tradingSuite.VaultPLGAddr,
	tradingSuite.PLGClient,
	tradingSuite.ChainIDPLG,
	0,
	3,
)
}

func (tradingSuite *PolygonTestSuite) TestX2_trade_MATIC_ETH_new_flow(){
return
fmt.Println("===== TestX2 TRADE MATIC ETH NEW FLOW =====")
fmt.Println("===== GET call data =====")
input := tradingSuite.CallData(
	big.NewInt(1000000000000),
	[]common.Address{tradingSuite.WMATICAddr, tradingSuite.WETHAddr},
	[]int64{LOW},
	false,
	false,)

fmt.Println(" ===== CALL BURNING ======")
	txhash,_ := tradingSuite.callBurningDapp(
		"f5d88e2e3c8f02d6dc1e01b54c90f673d730bef7d941aeec81ad1e1db690961f",
		big.NewInt(1000),
		"0x0000000000000000000000000000000000000000",
		tradingSuite.UniswapDeployedAddr.String(),
		"bridgeaggBurnForCall",
		tradingSuite.IncPaymentReceiverStr,
		0,
		"dae027b21d8d57114da11209dce8eeb587d01adf59d4fc356a8be5eedc146859",
		3,
		input,
		"0xa6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa",
	)
// return
	time.Sleep(30 * time.Second)
	fmt.Println(" ==== SUBMIT TRADE MATIC FOR ETH ===")
	tx_submit_trade := tradingSuite.submitBurnProofForWithdrawalNewDapp(
		// burningTxID.(string),
		txhash,
		// "getplgburnproof",
		"bridgeaggGetBurnProof",
		tradingSuite.VaultPLGAddr,
		tradingSuite.PLGClient,
		tradingSuite.ChainIDPLG,
		0,
		3,
	)

	time.Sleep(40 * time.Second)
	fmt.Println("==== WITHDRAW ======")

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, tx_submit_trade)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)

	// old shield flow
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncWETHTokenIDStr,
		// tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)

	// // new flow unified shield
	// _, err := tradingSuite.callIssuingUnifiedPtokenReq(
	// 	tradingSuite.IncUTEtherTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// 	"bridgeaggShield",
	// 	3,
	// 	tradingSuite.IncEtherTokenIDStr,
	// )

	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(60 * time.Second)
}

func (tradingSuite *PolygonTestSuite) TestX2A_trade_MATIC_ETH_new_flow_multi_trade(){
	return
	fmt.Println("===== TestX2A TRADE MATIC ETH NEW FLOW WITH MULTI TRADE =====")
	fmt.Println("===== GET call data =====")
	input := tradingSuite.CallDataMultiTrade(
		big.NewInt(1000 * 1e9),
		[][]common.Address{{ tradingSuite.WMATICAddr,tradingSuite.WETHAddr}, { tradingSuite.WMATICAddr,tradingSuite.WETHAddr}},
		[][]int64{{LOW}, {MEDIUM}},
		[]int64{7000, 3000},
		false,
		time.Now().Unix()+60000,
		[]bool{false, false},
	)
	
	fmt.Println(" ===== CALL BURNING ======")
		txhash,err := tradingSuite.callBurningDapp(
			"f5d88e2e3c8f02d6dc1e01b54c90f673d730bef7d941aeec81ad1e1db690961f",
			big.NewInt(1000),
			"0x0000000000000000000000000000000000000000",
			tradingSuite.UniswapDeployedAddr.String(),
			"bridgeaggBurnForCall",
			tradingSuite.IncPaymentReceiverStr,
			0,
			"dae027b21d8d57114da11209dce8eeb587d01adf59d4fc356a8be5eedc146859",
			3,
			input,
			"0xa6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa",
		)
	
		require.Equal(tradingSuite.T(), nil, err)
	// return
		time.Sleep(30 * time.Second)
		fmt.Println(" ==== SUBMIT TRADE MATIC FOR ETH ===")
		tx_submit_trade := tradingSuite.submitBurnProofForWithdrawalNewDapp(
			txhash,
			"bridgeaggGetBurnProof",
			tradingSuite.VaultPLGAddr,
			tradingSuite.PLGClient,
			tradingSuite.ChainIDPLG,
			0,
			3,
		)
	
		time.Sleep(40 * time.Second)
		fmt.Println("==== WITHDRAW ======")
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, tx_submit_trade)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
	
		// old shield flow
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncWETHTokenIDStr,
			// tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
			"createandsendtxwithissuingplgreq",
		)
	
		// // new flow unified shield
		// _, err = tradingSuite.callIssuingUnifiedPtokenReq(
		// 	tradingSuite.IncWETHTokenIDStr,
		// 	ethDepositProof,
		// 	ethBlockHash,
		// 	ethTxIdx,
		// 	"bridgeaggShield",
		// 	3,
		// 	tradingSuite.IncWETHTokenIDStr,
		// )
	
		require.Equal(tradingSuite.T(), nil, err)
	
		time.Sleep(60 * time.Second)
}
	
func (tradingSuite *PolygonTestSuite) TestX3_trade_ETH_MATIC_new_flow(){
	return
	fmt.Println("===== TestX3 TRADE ETH MATIC NEW FLOW =====")
	fmt.Println("===== GET call data =====")
	input := tradingSuite.CallData(
		big.NewInt(1000 * 1e9),
		[]common.Address{tradingSuite.WETHAddr,tradingSuite.WMATICAddr },
		[]int64{LOW},
		true,
		true,)
	
	fmt.Println(" ===== CALL BURNING ======")
		txhash,_ := tradingSuite.callBurningDapp(
			"a697a5c08d173de37372a20946e37d9e4adeeba68571b29b8ca4a2e1c3fc27fa",
			big.NewInt(1000),
			"0x0000000000000000000000000000000000000000",
			tradingSuite.UniswapDeployedAddr.String(),
			"bridgeaggBurnForCall",
			tradingSuite.IncPaymentReceiverStr,
			0,
			"a697a5c08d173de37372a20946e37d9e4adeeba68571b29b8ca4a2e1c3fc27fa",
			3,
			input,
			"0x0000000000000000000000000000000000000000",
		)
	// return
		time.Sleep(30 * time.Second)
		fmt.Println(" ==== SUBMIT TRADE ETH FOR MATIC ===")
		tx_submit_trade := tradingSuite.submitBurnProofForWithdrawalNewDapp(
			// burningTxID.(string),
			txhash,
			"bridgeaggGetBurnProof",
			tradingSuite.VaultPLGAddr,
			tradingSuite.PLGClient,
			tradingSuite.ChainIDPLG,
			0,
			3,
		)
	
		time.Sleep(40 * time.Second)
		fmt.Println("==== WITHDRAW ======")
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, tx_submit_trade)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
	
		// old shield flow
		_, err = tradingSuite.callIssuingETHReq(
			// tradingSuite.IncWETHTokenIDStr,
			tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
			"createandsendtxwithissuingplgreq",
		)
	
		// // new flow unified shield
		// _, err := tradingSuite.callIssuingUnifiedPtokenReq(
		// 	tradingSuite.IncUTEtherTokenIDStr,
		// 	ethDepositProof,
		// 	ethBlockHash,
		// 	ethTxIdx,
		// 	"bridgeaggShield",
		// 	3,
		// 	tradingSuite.IncEtherTokenIDStr,
		// )
	
		require.Equal(tradingSuite.T(), nil, err)
	
		time.Sleep(60 * time.Second)
}

func (tradingSuite *PolygonTestSuite) TestX3A_trade_ETH_MATIC_new_flow_multi_trade(){
	return
	fmt.Println("===== TestX3A TRADE ETH MATIC NEW FLOW WITH MULTI TRADE =====")
	fmt.Println("===== GET call data =====")
	input := tradingSuite.CallDataMultiTrade(
		big.NewInt(1000 * 1e9),
		[][]common.Address{{ tradingSuite.WETHAddr,tradingSuite.WMATICAddr}, { tradingSuite.WETHAddr,tradingSuite.WMATICAddr}},
		[][]int64{{LOW}, {MEDIUM}},
		[]int64{7000, 3000},
		true,
		time.Now().Unix()+60000,
		[]bool{false, false},
	)
	
	fmt.Println(" ===== CALL BURNING ======")
		txhash,_ := tradingSuite.callBurningDapp(
			"a697a5c08d173de37372a20946e37d9e4adeeba68571b29b8ca4a2e1c3fc27fa",
			big.NewInt(1000),
			"0x0000000000000000000000000000000000000000",
			tradingSuite.UniswapDeployedAddr.String(),
			"bridgeaggBurnForCall",
			tradingSuite.IncPaymentReceiverStr,
			0,
			"a697a5c08d173de37372a20946e37d9e4adeeba68571b29b8ca4a2e1c3fc27fa",
			3,
			input,
			"0x0000000000000000000000000000000000000000",
		)
	// return
		time.Sleep(30 * time.Second)
		fmt.Println(" ==== SUBMIT TRADE ETH FOR MATIC ===")
		tx_submit_trade := tradingSuite.submitBurnProofForWithdrawalNewDapp(
			// burningTxID.(string),
			txhash,
			"bridgeaggGetBurnProof",
			tradingSuite.VaultPLGAddr,
			tradingSuite.PLGClient,
			tradingSuite.ChainIDPLG,
			0,
			3,
		)
	
		time.Sleep(40 * time.Second)
		fmt.Println("==== WITHDRAW ======")
	
		_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, tx_submit_trade)
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)
	
		fmt.Println("Waiting 90s for 15 blocks confirmation")
		time.Sleep(100 * time.Second)
	
		// old shield flow
		_, err = tradingSuite.callIssuingETHReq(
			tradingSuite.IncWETHTokenIDStr,
			// tradingSuite.IncEtherTokenIDStr,
			ethDepositProof,
			ethBlockHash,
			ethTxIdx,
			"createandsendtxwithissuingplgreq",
		)
	
		// // new flow unified shield
		// _, err := tradingSuite.callIssuingUnifiedPtokenReq(
		// 	tradingSuite.IncUTEtherTokenIDStr,
		// 	ethDepositProof,
		// 	ethBlockHash,
		// 	ethTxIdx,
		// 	"bridgeaggShield",
		// 	3,
		// 	tradingSuite.IncEtherTokenIDStr,
		// )
	
		require.Equal(tradingSuite.T(), nil, err)
	
		time.Sleep(60 * time.Second)
}
	
func (tradingSuite *PolygonTestSuite) TestX4_submit_redeposit(){
	return
	fmt.Println("==== TESTX4 SUBMIT REDEPOSIT ======")

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("f3d075f2def7473b97d41d377201dc262b9f7ed29159f3e5c54c6d924755ca7d"))
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(5 * time.Second)

	// old shield flow
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncWETHTokenIDStr,
		// tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingplgreq",
	)


	// // new flow unified shield
	// _, err = tradingSuite.callIssuingUnifiedPtokenReq(
	// 	tradingSuite.IncUTEtherTokenIDStr,
	// 	ethDepositProof,
	// 	ethBlockHash,
	// 	ethTxIdx,
	// 	"bridgeaggShield",
	// 	3,
	// 	tradingSuite.IncEtherTokenIDStr,
	// )
	require.Equal(tradingSuite.T(), nil, err)

	time.Sleep(60 * time.Second)
}

func (tradingSuite *PolygonTestSuite) TestX5_get_call_data(){
	return

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.PLGHost, common.HexToHash("0x85ca18aaa57f2c553e5ec4fa0abf9c5744976d5eaf17f2ae93d48b2fb278dbe4"))
		require.Equal(tradingSuite.T(), nil, err)
		fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

// 	fmt.Println("===== GET call data =====")
//  tradingSuite.CallData(
// 		big.NewInt(100000),
// 		[]common.Address{tradingSuite.WETHAddr,tradingSuite.WMATICAddr},
// 		// []common.Address{common.HexToAddress("0x2058A9D7613eEE744279e3856Ef0eAda5FCbaA7e"),tradingSuite.DAIAddr},
// 		[]int64{LOW},
// 		true,
// 		true,)


	// fmt.Println("===== GET call data multi trade =====")
	// tradingSuite.CallDataMultiTrade(
	// 	big.NewInt(1000000000000),
	// 	[][]common.Address{{ tradingSuite.MATICAddr,tradingSuite.WETHAddr}, {tradingSuite.MATICAddr,tradingSuite.WETHAddr}},
	// 	[][]int64{{LOW}, {MEDIUM}},
	// 	[]int64{70, 30},
	// 	false,
	// 	time.Now().Unix()+60000,
	// 	[]bool{false, false},
	// )
	}