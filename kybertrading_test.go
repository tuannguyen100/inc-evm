package bridge

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"
	"bytes"
	"encoding/hex"

	"github.com/incognitochain/bridge-eth/bridge/kbntrade"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/stretchr/testify/suite"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"github.com/stretchr/testify/require"
	pUniswapHelper "github.com/incognitochain/bridge-eth/bridge/puniswaphelper"
	puniswap "github.com/incognitochain/bridge-eth/bridge/puniswapproxy"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type KyberTradingTestSuite struct {
	*TradingTestSuite

	KyberTradeDeployedAddr      common.Address
	KyberMultiTradeDeployedAddr common.Address
	KyberContractAddr           common.Address
	WETHAddr                    common.Address

	ETHUniswapDeployedAddr      common.Address
	ETHUniswapRouteContractAddr common.Address
	ETHUNiswapQuoteContractAddr common.Address

	IncKBNTokenIDStr  string
	IncSALTTokenIDStr string
	IncOMGTokenIDStr  string
	IncSNTTokenIDStr  string

	EtherAddressStrKyber string
	KBNAddressStr        string
	SALTAddressStr       string
	OMGAddressStr        string
	SNTAddressStr        string
	WEtherAddr        common.Address
	DAIETHAddr 	common.Address

	// token amounts for tests
	DepositingEther       float64
	KBNBalanceAfterStep1  *big.Int
	SALTBalanceAfterStep2 *big.Int
}

func NewKyberTradingTestSuite(tradingTestSuite *TradingTestSuite) *KyberTradingTestSuite {
	return &KyberTradingTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *KyberTradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	tradingSuite.WETHAddr = common.HexToAddress("0x0Bb7509324cE409F7bbC4b701f932eAca9736AB7")
	tradingSuite.DAIETHAddr = common.HexToAddress("0x73967c6a0904aA032C103b4104747E88c566B1A2")

	
	tradingSuite.ETHUniswapDeployedAddr = common.HexToAddress("0x1ec63144756FC4905341ef5907fB7873cCDdb798") 
	tradingSuite.ETHUniswapRouteContractAddr = common.HexToAddress("0x68b3465833fb72a70ecdf485e0e4c7bd8665fc45")
	tradingSuite.ETHUNiswapQuoteContractAddr = common.HexToAddress("0x61ffe014ba17989e743c5f6cb21bf9697530b21e")

	tradingSuite.DepositingEther = float64(0.0005)
	
}

func (tradingSuite *KyberTradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *KyberTradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *KyberTradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestKyberTradingTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Kyber test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	kyberTradingSuite := NewKyberTradingTestSuite(tradingSuite)
	suite.Run(t, kyberTradingSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}

func (tradingSuite *KyberTradingTestSuite) getExpectedRate(
	srcToken string,
	destToken string,
	srcQty *big.Int,
) *big.Int {
	if srcToken == tradingSuite.EtherAddressStr {
		srcToken = tradingSuite.EtherAddressStrKyber
	}
	if destToken == tradingSuite.EtherAddressStr {
		destToken = tradingSuite.EtherAddressStrKyber
	}
	c, err := kbntrade.NewKBNTrade(tradingSuite.KyberTradeDeployedAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	expectRate, slippageRate, err := c.GetConversionRates(nil, common.HexToAddress(srcToken), srcQty, common.HexToAddress(destToken))
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("slippageRate value: %d\n", slippageRate)
	fmt.Printf("expectRate value: %d\n", expectRate)
	return expectRate
}

func (tradingSuite *KyberTradingTestSuite) executeWithKyber(
	srcQty *big.Int,
	srcTokenIDStr string,
	destTokenIDStr string,
) {
	tradeAbi, _ := abi.JSON(strings.NewReader(kbntrade.KBNTradeABI))

	// Get contract instance
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth := bind.NewKeyedTransactor(tradingSuite.ETHPrivKey)
	auth.GasPrice = big.NewInt(50000000000)
	// auth.GasLimit = 2000000
	srcToken := common.HexToAddress(srcTokenIDStr)
	destToken := common.HexToAddress(destTokenIDStr)
	expectRate := tradingSuite.getExpectedRate(srcTokenIDStr, destTokenIDStr, srcQty)
	input, _ := tradeAbi.Pack("trade", srcToken, srcQty, destToken, expectRate)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix: EXECUTE_PREFIX,
		Token: srcToken,
		Timestamp: timestamp,
		Amount: srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.KyberTradeDeployedAddr, input)
	if err != nil{
		panic(err)
	}
	data := rawsha3(tempData[4:])
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)

	tx, err := c.Execute(
		auth,
		srcToken,
		srcQty,
		destToken,
		tradingSuite.KyberTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("Kyber trade executed , txHash: %x\n", txHash[:])
}


func (tradingSuite *KyberTradingTestSuite) buildPath(paths []common.Address, fees []int64) []byte {
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

func (tradingSuite *KyberTradingTestSuite) getExpectedAmount(
	srcQty *big.Int,
	paths []common.Address,
	fees []int64,
) *big.Int {
	c, err := pUniswapHelper.NewPUniswapHelper(tradingSuite.ETHUNiswapQuoteContractAddr, tradingSuite.ETHClient)
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

func (tradingSuite *KyberTradingTestSuite) executeWithPUniswap(
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
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDETH)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(3e10)
	// auth.GasLimit = 1000000
	var agr interface{}
	expectOutputAmount := tradingSuite.getExpectedAmount(
		srcQty,
		paths,
		fees,
	)
	recipient := tradingSuite.VaultAddr
	if isNative && bytes.Compare(paths[len(paths)-1].Bytes(), tradingSuite.WEtherAddr.Bytes()) == 0 {
		recipient = tradingSuite.ETHUniswapDeployedAddr
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
	if paths[0].String() == tradingSuite.WEtherAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	destToken := paths[len(paths)-1]
	if paths[len(paths)-1].String() == tradingSuite.WEtherAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.ETHUniswapDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.ETHUniswapDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("pUniswap trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *KyberTradingTestSuite) executeWithPUniswapMultiTrade(
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
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(tradingSuite.ChainIDETH)))
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
		recipient := tradingSuite.VaultAddr
		if isNative && bytes.Compare(paths[i][len(paths[i])-1].Bytes(), tradingSuite.WEtherAddr.Bytes()) == 0 {
			recipient = tradingSuite.ETHUniswapDeployedAddr
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
	if paths[0][0].String() == tradingSuite.WEtherAddr.String() {
		sourceToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	destToken := paths[0][len(paths[0])-1]
	if paths[0][len(paths[0])-1].String() == tradingSuite.WEtherAddr.String() && isNative {
		destToken = common.HexToAddress(tradingSuite.EtherAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, tradingSuite.ETHUniswapDeployedAddr, input)
	require.Equal(tradingSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		tradingSuite.ETHUniswapDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(tradingSuite.ETHClient, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("pUniswap trade executed , txHash: %x\n", txHash[:])
}

func (tradingSuite *KyberTradingTestSuite) CallData(
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
	recipient := tradingSuite.VaultAddr
	if isNative && bytes.Compare(paths[len(paths)-1].Bytes(), tradingSuite.WEtherAddr.Bytes()) == 0 {
		recipient = tradingSuite.ETHUniswapDeployedAddr
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


func (tradingSuite *KyberTradingTestSuite) CallDataMultiTrade(
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
		recipient := tradingSuite.VaultAddr
		if isNative && bytes.Compare(paths[i][len(paths[i])-1].Bytes(), tradingSuite.WEtherAddr.Bytes()) == 0 {
			recipient = tradingSuite.ETHUniswapDeployedAddr
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



func (tradingSuite *KyberTradingTestSuite) Test1TradeEthForKBNWithKyber() {
	return
	fmt.Println("============ TEST TRADE ETHER FOR KBN WITH Kyber AGGREGATOR ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	tradeAmount := big.NewInt(int64(tradingSuite.DepositingEther * params.Ether))
	burningPETH := big.NewInt(0).Div(tradeAmount, big.NewInt(1000000000))

	pubKeyToAddrStr := crypto.PubkeyToAddress(tradingSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting ETH to pETH --------------")
	txHash := tradingSuite.depositETH(
		tradingSuite.DepositingEther,
		tradingSuite.IncPaymentAddrStr,
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
	)
	// time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(tradingSuite.ETHHost, txHash)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(230 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncEtherTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(50 * time.Second)

	fmt.Println("------------ STEP 2: burning pETH to deposit ETH to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := tradingSuite.callBurningPToken(
		tradingSuite.IncEtherTokenIDStr,
		burningPETH,
		pubKeyToAddrStr[2:],
		"createandsendburningfordeposittoscrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(50 * time.Second)

	tradingSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(tradingSuite.ChainIDETH)),
		"getburnprooffordeposittosc",
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
	)
	deposited := tradingSuite.getDepositedBalance(
		tradingSuite.EtherAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("deposited EHT: ", deposited)
	// require.Equal(tradingSuite.T(), big.NewInt(0).Mul(burningPETH, big.NewInt(1000000000)), deposited)
return
	fmt.Println("------------ step 3: execute trade ETH for KBN through Kyber aggregator --------------")
	tradingSuite.executeWithKyber(
		deposited,
		tradingSuite.EtherAddressStr,
		tradingSuite.KBNAddressStr,
	)
	time.Sleep(15 * time.Second)
	kbnTraded := tradingSuite.getDepositedBalance(
		tradingSuite.KBNAddressStr,
		pubKeyToAddrStr,
	)
	fmt.Println("kbnTraded: ", kbnTraded)

	fmt.Println("------------ step 4: withdrawing KBN from SC to pKBN on Incognito --------------")
	txHashByEmittingWithdrawalReq := tradingSuite.requestWithdraw(
		tradingSuite.KBNAddressStr,
		kbnTraded,
		tradingSuite.BSCClient,
		big.NewInt(int64(tradingSuite.ChainIDETH)),
		tradingSuite.VaultAddr,
		REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(tradingSuite.ETHHost, txHashByEmittingWithdrawalReq)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	_, err = tradingSuite.callIssuingETHReq(
		tradingSuite.IncKBNTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingethreq",
	)
	require.Equal(tradingSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	fmt.Println("------------ step 5: withdrawing pKBN from Incognito to KBN --------------")
	withdrawingPKBN := big.NewInt(0).Div(kbnTraded, big.NewInt(1000000000))
	// withdrawingPKBN.Mul(withdrawingPKBN, big.NewInt(1000000000))
	burningRes, err = tradingSuite.callBurningPToken(
		tradingSuite.IncKBNTokenIDStr,
		withdrawingPKBN,
		tradingSuite.ETHOwnerAddrStr,
		"createandsendburningrequest",
	)
	require.Equal(tradingSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(tradingSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	tradingSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getburnproof",
		tradingSuite.VaultAddr,
		tradingSuite.ETHClient,
		tradingSuite.ChainIDETH,
	)

	bal := tradingSuite.getBalanceOnETHNet(
		common.HexToAddress(tradingSuite.KBNAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", tradingSuite.ETHOwnerAddrStr)),
		tradingSuite.ETHClient,
	)
	tradingSuite.KBNBalanceAfterStep1 = bal
	fmt.Println("KBN balance after step 1: ", tradingSuite.KBNBalanceAfterStep1)
	// require.Equal(tradingSuite.T(), withdrawingPKBN.Uint64(), bal.Div(bal, big.NewInt(1000000000)).Uint64())
}


func (tradingSuite *KyberTradingTestSuite) Testxxxxx() {
return
tradingSuite.submitBurnProofForDepositToSC(
	"7b96e0187ea1beb5290a4eea2b49cfac6c3f126814c7e4010387b71c9aae8e8d",
	big.NewInt(int64(tradingSuite.ChainIDETH)),
	"getburnprooffordeposittosc",
	tradingSuite.VaultAddr,
	tradingSuite.ETHClient,
)
}

func (tradingSuite *KyberTradingTestSuite) TestX5_get_call_data(){
	// return
	fmt.Println("===== GET call data =====")
 tradingSuite.CallData(
		big.NewInt(100000),
		[]common.Address{tradingSuite.WEtherAddr,tradingSuite.DAIETHAddr},
		// []common.Address{common.HexToAddress("0x2058A9D7613eEE744279e3856Ef0eAda5FCbaA7e"),tradingSuite.DAIAddr},
		[]int64{LOW},
		true,
		true,)


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