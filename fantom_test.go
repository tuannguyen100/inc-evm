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
type FantomTestSuite struct {
	*TradingTestSuite

	PanackeTradeDeployedAddr common.Address
	PanackeRouteContractAddr common.Address


	IncBUSDTokenIDStr string
	IncLIQDTokenIDStr string

	WFTMAddr        common.Address
	WMATICAddr        common.Address
	WETHAddr          common.Address
	LIQDAddressStr        string
	FTMAddressStr 	string

	// token amounts for tests
	DepositingFantom  float64
	WithdrawingFantom float64
}

func NewFantomTestSuite(tradingTestSuite *TradingTestSuite) *FantomTestSuite {
	return &FantomTestSuite{
		TradingTestSuite: tradingTestSuite,
	}
}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (fantomSuite *FantomTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	fantomSuite.PanackeTradeDeployedAddr = common.HexToAddress("0x14D0cf3bC307aA15DA40Aa4c8cc2A2a81eF96B3a")
	fantomSuite.PanackeRouteContractAddr = common.HexToAddress("0xa6ad18c2ac47803e193f75c3677b14bf19b94883")


	// Fantom testnet env
	fantomSuite.FTMAddressStr ="0x0000000000000000000000000000000000000000"
	fantomSuite.IncLIQDTokenIDStr = "9f3c0e13bd4c307fd2ac692ed83d2ad4411bd117741010ff91513eea2c2ca7a8"

	// tokens
	fantomSuite.WFTMAddr = common.HexToAddress("0xf1277d1Ed8AD466beddF92ef448A132661956621")
	fantomSuite.WMATICAddr = common.HexToAddress("0x9c3c9283d3e44854697cd22d3faa240cfb032889")
	fantomSuite.WETHAddr = common.HexToAddress("0xa6fa4fb5f76172d178d61b04b0ecd319c5d1c0aa")
	fantomSuite.LIQDAddressStr = "0x8658a7931F9d94180daC7135c627a27B62f199F5"

	fantomSuite.DepositingFantom = float64(0.00123456789)
	fantomSuite.WithdrawingFantom = float64(0.00123456789)
}

func (fantomSuite *FantomTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	fantomSuite.FTMClient.Close()
}

func (fantomSuite *FantomTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (fantomSuite *FantomTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestFantomTestSuite(t *testing.T) {
	fmt.Println("Starting entry point for Fantom bridge test suite...")

	tradingSuite := new(TradingTestSuite)
	suite.Run(t, tradingSuite)

	fantomTestSuite := NewFantomTestSuite(tradingSuite)
	suite.Run(t, fantomTestSuite)

	fmt.Println("Finishing entry point for 0x test suite...")
}



func (fantomSuite *FantomTestSuite) getExpectedAmount(
	path []common.Address,
	srcQty *big.Int,
) *big.Int {
	c, err := pancakeproxy.NewPancakeproxyrouter(fantomSuite.PanackeRouteContractAddr, fantomSuite.FTMClient)
	require.Equal(fantomSuite.T(), nil, err)
	amounts, err := c.GetAmountsOut(nil, srcQty, path)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Printf("path: %v\n", path)
	fmt.Printf("intput value: %v\n", srcQty.String())
	fmt.Printf("output value: %v\n", amounts[len(amounts)-1].String())
	fmt.Printf("path output : %v\n", amounts)

	return amounts[len(amounts)-1]
}

func (fantomSuite *FantomTestSuite) executeWithPancake(
	srcQty *big.Int,
	path []common.Address,
	deadline uint,
	isNative bool,
) {
	require.NotEqual(fantomSuite.T(), 0, len(path))

	tradeAbi, err := abi.JSON(strings.NewReader(pancakeproxy.PancakeproxyMetaData.ABI))
	require.Equal(fantomSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(fantomSuite.VaultFTMAddr, fantomSuite.FTMClient)
	require.Equal(fantomSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(fantomSuite.ETHPrivKey, big.NewInt(int64(fantomSuite.ChainIDFTM)))
	require.Equal(fantomSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(10e10)
	expectOutputAmount := fantomSuite.getExpectedAmount(path, srcQty)
	input, err := tradeAbi.Pack("trade", path, srcQty, expectOutputAmount, big.NewInt(int64(deadline)), isNative)
	require.Equal(fantomSuite.T(), nil, err)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	require.Equal(fantomSuite.T(), nil, err)
	sourceToken := path[0]
	if path[0].String() == fantomSuite.WFTMAddr.String() {
		sourceToken = common.HexToAddress(fantomSuite.FTMAddressStr)
	}
	destToken := path[len(path)-1]
	if path[len(path)-1].String() == fantomSuite.WFTMAddr.String() && isNative {
		destToken = common.HexToAddress(fantomSuite.FTMAddressStr)
	}
	psData := vault.VaultHelperPreSignData{
		Prefix:    FTM_EXECUTE_PREFIX,
		Token:     sourceToken,
		Timestamp: timestamp,
		Amount:    srcQty,
	}
	tempData, err := vaultAbi.Pack("_buildSignExecute", psData, destToken, fantomSuite.PanackeTradeDeployedAddr, input)
	require.Equal(fantomSuite.T(), nil, err)
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, &fantomSuite.GeneratedPrivKeyForSC)
	require.Equal(fantomSuite.T(), nil, err)

	tx, err := c.Execute(
		auth,
		sourceToken,
		srcQty,
		destToken,
		fantomSuite.PanackeTradeDeployedAddr,
		input,
		timestamp,
		signBytes,
	)
	require.Equal(fantomSuite.T(), nil, err)
	txHash := tx.Hash()
	if err := wait(fantomSuite.FTMClient, txHash); err != nil {
		require.Equal(fantomSuite.T(), nil, err)
	}
	fmt.Printf("Dapp trade executed , txHash: %x\n", txHash[:])
}

func (fantomSuite *FantomTestSuite) CallData(
	srcQty *big.Int,
	path []common.Address,
	deadline uint,
	isNative bool,
) []byte {
	require.NotEqual(fantomSuite.T(), 0, len(path))

	tradeAbi, err := abi.JSON(strings.NewReader(pancakeproxy.PancakeproxyMetaData.ABI))
	require.Equal(fantomSuite.T(), nil, err)

	expectOutputAmount := fantomSuite.getExpectedAmount(path, srcQty)
	amount := big.NewInt(0).Div(expectOutputAmount, big.NewInt(1))
	// amount := big.NewInt(2082011)
	input, err := tradeAbi.Pack("trade", path, srcQty, amount, big.NewInt(int64(deadline)), isNative)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("Call Data : ", common.Bytes2Hex(input))
	return input
}



func (fantomSuite *FantomTestSuite) Test1ShieldUnshieldFTM() {
	return
	fmt.Println("============ TEST 1 SHIELD UNSHIELD FANTOM ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	// depositAmount := big.NewInt(int64(fantomSuite.DepositingFantom * params.Ether))
	withdrawAmount := big.NewInt(int64(fantomSuite.WithdrawingFantom * params.Ether))
	withdrawingAmt := big.NewInt(0).Div(withdrawAmount, big.NewInt(1e9))

	pubKeyToAddrStr := crypto.PubkeyToAddress(fantomSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting FTM to pFTM --------------")
	// create Fantom tx to send FTM to Vault FTM
	txHash := fantomSuite.depositETH(
		fantomSuite.DepositingFantom,
		fantomSuite.IncPaymentAddrStr,
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
	)
	// txHash := common.HexToHash("0x0b297b019fde220d2b9a49231f853ab616507ec1f4791899b0b937257875fe1e")
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, txHash)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	fmt.Println("Waiting 3s for 1 block confirmation")
	time.Sleep(120 * time.Second)
	txhashInC, err := fantomSuite.callIssuingETHReq(
		fantomSuite.IncFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingftmreq",
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(50 * time.Second)
	require.Equal(fantomSuite.T(), 2, fantomSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncFTMTokenIDStr)
	fmt.Println("[INFO] FTM balance incognito after issuing step 1 : ", balpBNBS1)
// return
	fmt.Println("------------ STEP 2: burning pFTM to deposit FTM to SC --------------")
	// // make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := fantomSuite.callBurningPToken(
		fantomSuite.IncFTMTokenIDStr,
		withdrawingAmt,
		pubKeyToAddrStr[2:],
		"createandsendburningftmfordeposittoscrequest",
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(100 * time.Second)

	fantomSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		"getburnftmprooffordeposittosc",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
	)
	deposited := fantomSuite.getDepositedBalanceFTM(
		common.HexToAddress(fantomSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("deposited %v \n", deposited)
	
	fmt.Println("------------ step 3: withdrawing WETH from SC to pWETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := fantomSuite.requestWithdraw(
		fantomSuite.EtherAddressStr,
		deposited,
		fantomSuite.FTMClient,
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		fantomSuite.VaultFTMAddr,
		FTM_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(fantomSuite.FTMHost, txHashByEmittingWithdrawalReq)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	txhashInC, err = fantomSuite.callIssuingETHReq(
		fantomSuite.IncFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingftmreq",
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	require.Equal(fantomSuite.T(), 2, fantomSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncFTMTokenIDStr)
	fmt.Println("[INFO] FTM balance incognito after issuing step 3 : ", balpBNBS1)

	fmt.Println("------------ step 4: withdrawing pFTM from Incognito to FTM --------------")
	burningRes, err = fantomSuite.callBurningPToken(
		fantomSuite.IncFTMTokenIDStr,
		withdrawingAmt,
		fantomSuite.ETHOwnerAddrStr,
		"createandsendburningftmrequest",
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	fantomSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getftmburnproof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
	)

	bal := fantomSuite.getBalanceOnETHNet(
		common.HexToAddress(fantomSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", fantomSuite.ETHOwnerAddrStr)),
		fantomSuite.FTMClient,
	)
	fmt.Println("FTM balance: ", bal)
}

func (fantomSuite *FantomTestSuite) Test2shieldUnshieldLINK() {
	return
	fmt.Println("============ TEST 2 SHIELD UNSHIELD LINK ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	// depositAmount := big.NewInt(int64(fantomSuite.DepositingFantom * params.Ether))
	withdrawAmount := big.NewInt(int64(fantomSuite.DepositingFantom * params.Ether))
	withdrawingAmt := big.NewInt(0).Div(withdrawAmount, big.NewInt(1e9))

	pubKeyToAddrStr := crypto.PubkeyToAddress(fantomSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting LINK to pLINK --------------")
	// create Fantom tx to send FTM to Vault FTM
	txHash := fantomSuite.depositERC20ToBridge(
		withdrawAmount,
		common.HexToAddress(fantomSuite.DAIAddressStr),
		fantomSuite.IncPaymentAddrStr,
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
	)
	// txHash := common.HexToHash("0x7a5c68e7c78b1e8a0c2fe99864a536af4000965e8fbb8f5c04ba6278c6d6d63d")
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, txHash)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	fmt.Println("Waiting 3s for 6 block confirmation")
	time.Sleep(100 * time.Second)
	txhashInC, err := fantomSuite.callIssuingETHReq(
		fantomSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingftmreq",
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	require.Equal(fantomSuite.T(), 2, fantomSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pLINK  balance incognito after issuing step 1 : ", balpBNBS1)

	fmt.Println("------------ STEP 2: burning pLINK to deposit LINK to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := fantomSuite.callBurningPToken(
		fantomSuite.IncDAITokenIDStr,
		withdrawingAmt,
		pubKeyToAddrStr[2:],
		"createandsendburningftmfordeposittoscrequest",
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(120 * time.Second)

	fantomSuite.submitBurnProofForDepositToSC(
		burningTxID.(string),
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		"getburnftmprooffordeposittosc",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
	)
	deposited := fantomSuite.getDepositedBalanceFTM(
		common.HexToAddress(fantomSuite.DAIAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("LINK balance %v \n", deposited)
	

	fmt.Println("------------ step 3: withdrawing LINK from SC to pLINK on Incognito --------------")
	txHashByEmittingWithdrawalReq := fantomSuite.requestWithdraw(
		fantomSuite.DAIAddressStr,
		deposited,
		fantomSuite.FTMClient,
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		fantomSuite.VaultFTMAddr,
		FTM_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(fantomSuite.FTMHost, txHashByEmittingWithdrawalReq)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(100 * time.Second)
	txhashInC, err = fantomSuite.callIssuingETHReq(
		fantomSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingftmreq",
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	require.Equal(fantomSuite.T(), 2, fantomSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pLINK token balance incognito after issuing step 3 : ", balpBNBS1)

	fmt.Println("------------ step 4: withdrawing pLINK from Incognito to LINK --------------")
	burningRes, err = fantomSuite.callBurningPToken(
		fantomSuite.IncDAITokenIDStr,
		withdrawingAmt,
		fantomSuite.ETHOwnerAddrStr,
		"createandsendburningftmrequest",
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(40 * time.Second)

	fantomSuite.submitBurnProofForWithdrawal(
		burningTxID.(string),
		"getftmburnproof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
	)

	bal := fantomSuite.getBalanceOnETHNet(
		common.HexToAddress(fantomSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", fantomSuite.ETHOwnerAddrStr)),
		fantomSuite.FTMClient,
	)
	fmt.Println("LINK balance: ", bal)
}

func (fantomSuite *FantomTestSuite) Test3UnifiedFTM() {
	return
	fmt.Println("============ TEST 3 SHIELD UNSHIELD UNIFIED FANTOM ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")

	withdrawAmount := big.NewInt(int64(fantomSuite.WithdrawingFantom * params.Ether))
	withdrawingAmt := big.NewInt(0).Div(withdrawAmount, big.NewInt(1e9))

	pubKeyToAddrStr := crypto.PubkeyToAddress(fantomSuite.GeneratedPubKeyForSC).Hex()
	
	fmt.Println("------------ STEP 1: porting FTM to pFTM --------------")
	// create Fantom tx to send FTM to Vault FTM
	txHash := fantomSuite.depositETH(
		fantomSuite.DepositingFantom,
		fantomSuite.IncPaymentAddrStr,
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
	)
	// txHash := common.HexToHash("0xf1449fefc4c1cdb95d70753050597957e2936ad794cdabcddbae02399d53f45f")
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, txHash)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 3s for 1 block confirmation")
	time.Sleep(120 * time.Second)
	txhashInC, err := fantomSuite.callIssuingUnifiedPtokenReq(
		fantomSuite.IncUTFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		4,
		fantomSuite.IncFTMTokenIDStr,
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(50 * time.Second)
	require.Equal(fantomSuite.T(), 2, fantomSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncUTFTMTokenIDStr)
	fmt.Println("[INFO] UTFTM balance incognito after issuing step 1 : ", balpBNBS1)
	// return
	fmt.Println("------------ STEP 2: burning pFTM to deposit FTM to SC --------------")
	// // make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := fantomSuite.callBurningUnifiedPToken(
		fantomSuite.IncUTFTMTokenIDStr,
		big.NewInt(int64(balpBNBS1)),
		// withdrawingAmt,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		fantomSuite.IncPaymentAddrStr,
		0,
		fantomSuite.IncFTMTokenIDStr,
		true,
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)

	time.Sleep(80 * time.Second)
	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncUTFTMTokenIDStr)
	fmt.Println("[INFO] UTFTM balance incognito after issuing step 2 : ", balpBNBS1)

	fantomSuite.submitBurnProofForDepositToSCV2(
		burningTxID.(string),
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		"bridgeaggGetBurnProof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		0,
		4,
	)
	deposited := fantomSuite.getDepositedBalanceFTM(
		common.HexToAddress(fantomSuite.EtherAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("deposited %v \n", deposited)
	
// return
	fmt.Println("------------ step 3: withdrawing WETH from SC to pWETH on Incognito --------------")
	txHashByEmittingWithdrawalReq := fantomSuite.requestWithdraw(
		fantomSuite.EtherAddressStr,
		deposited,
		fantomSuite.FTMClient,
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		fantomSuite.VaultFTMAddr,
		FTM_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(45 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(fantomSuite.FTMHost, txHashByEmittingWithdrawalReq)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(90 * time.Second)
	_, err = fantomSuite.callIssuingUnifiedPtokenReq(
		fantomSuite.IncUTFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		4,
		fantomSuite.IncFTMTokenIDStr,
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(120 * time.Second)

	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncUTFTMTokenIDStr)
	fmt.Println("[INFO] UTFTM balance incognito after issuing step 3 : ", balpBNBS1)

	fmt.Println("------------ step 4: withdrawing pFTM from Incognito to FTM --------------")
	burningRes, err = fantomSuite.callBurningUnifiedPToken(
		fantomSuite.IncUTFTMTokenIDStr,
		withdrawingAmt,
		big.NewInt(609),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		fantomSuite.IncPaymentAddrStr,
		0,
		fantomSuite.IncFTMTokenIDStr,
		false,
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)

	time.Sleep(120 * time.Second)
	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncUTFTMTokenIDStr)
	fmt.Println("[INFO] UTFTM balance incognito after issuing step 4 : ", balpBNBS1)

	fantomSuite.submitBurnProofForWithdrawalV2(
		burningTxID.(string),
		"bridgeaggGetBurnProof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
		0,
		4,
	
	)	
	bal := fantomSuite.getBalanceOnETHNet(
		common.HexToAddress(fantomSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", fantomSuite.ETHOwnerAddrStr)),
		fantomSuite.FTMClient,
	)
	fmt.Println("FTM balance: ", bal)
}

func (fantomSuite *FantomTestSuite) Test4UnifiedLINK() {
	return
	fmt.Println("============ TEST 4 SHIELD UNSHIELD UNIFIED LINK ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	depositAmount := big.NewInt(int64(fantomSuite.DepositingFantom * params.Ether))
	withdrawAmount := big.NewInt(int64(fantomSuite.WithdrawingFantom * params.Ether))
	withdrawingAmt := big.NewInt(0).Div(withdrawAmount, big.NewInt(1e9))

	pubKeyToAddrStr := crypto.PubkeyToAddress(fantomSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting LINK to pLINK --------------")
	// create Fantom tx to send FTM to Vault FTM
	txHash := fantomSuite.depositERC20ToBridge(
		depositAmount,
		common.HexToAddress(fantomSuite.DAIAddressStr),
		fantomSuite.IncPaymentAddrStr,
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
	)
	// txHash := common.HexToHash("0x7a5c68e7c78b1e8a0c2fe99864a536af4000965e8fbb8f5c04ba6278c6d6d63d")
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, txHash)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
// return
	fmt.Println("Waiting 3s for 6 block confirmation")
	time.Sleep(80 * time.Second)
	txhashInC, err := fantomSuite.callIssuingUnifiedPtokenReq(
		fantomSuite.IncUTDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		4,
		fantomSuite.IncDAITokenIDStr,
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(70 * time.Second)
	require.Equal(fantomSuite.T(), 2, fantomSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncUTDAITokenIDStr)
	fmt.Println("[INFO] UTLINK balance incognito after issuing step 1 : ", balpBNBS1)

return
	fmt.Println("------------ STEP 2: burning pLINK to deposit LINK to SC --------------")
	// make a burn tx to incognito chain as a result of deposit to SC
	burningRes, err := fantomSuite.callBurningUnifiedPToken(
		fantomSuite.IncUTDAITokenIDStr,
		withdrawingAmt,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		fantomSuite.IncPaymentAddrStr,
		0,
		fantomSuite.IncDAITokenIDStr,
		true,
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(70 * time.Second)
// return
	fantomSuite.submitBurnProofForDepositToSCV2(
		burningTxID.(string),
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		"bridgeaggGetBurnProof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		0,
		4,
	)
	deposited := fantomSuite.getDepositedBalanceFTM(
		common.HexToAddress(fantomSuite.DAIAddressStr),
		pubKeyToAddrStr,
	)
	fmt.Printf("LINK balance %v \n", deposited)
// return	

	fmt.Println("------------ step 3: withdrawing LINK from SC to pLINK on Incognito --------------")
	txHashByEmittingWithdrawalReq := fantomSuite.requestWithdraw(
		fantomSuite.DAIAddressStr,
		deposited,
		fantomSuite.FTMClient,
		big.NewInt(int64(fantomSuite.ChainIDFTM)),
		fantomSuite.VaultFTMAddr,
		FTM_REQ_WITHDRAW_PREFIX,
	)
	time.Sleep(15 * time.Second)

	_, ethBlockHash, ethTxIdx, ethDepositProof, err = getETHDepositProof(fantomSuite.FTMHost, txHashByEmittingWithdrawalReq)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 6 blocks confirmation")
	time.Sleep(80 * time.Second)
	_, err = fantomSuite.callIssuingUnifiedPtokenReq(
		fantomSuite.IncUTDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		4,
		fantomSuite.IncDAITokenIDStr,
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(80 * time.Second)

	fmt.Println("------------ step 4: withdrawing pLINK from Incognito to LINK --------------")
	burningRes, err = fantomSuite.callBurningUnifiedPToken(
		fantomSuite.IncUTDAITokenIDStr,
		withdrawingAmt,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		fantomSuite.IncPaymentAddrStr,
		0,
		fantomSuite.IncDAITokenIDStr,
		false,
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found = burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(30 * time.Second)
// return
	fantomSuite.submitBurnProofForWithdrawalV2(
		burningTxID.(string),
		"bridgeaggGetBurnProof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
		0,
		4,
	)

	bal := fantomSuite.getBalanceOnETHNet(
		common.HexToAddress(fantomSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", fantomSuite.ETHOwnerAddrStr)),
		fantomSuite.FTMClient,
	)
	fmt.Println("LINK balance: ", bal)
}

func (fantomSuite *FantomTestSuite) Test5ShieldConvertUnshieldFTM() {
	return
	fmt.Println("============ TEST 5 SHIELD UNSHIELD CONVERT FANTOM ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	// depositAmount := big.NewInt(int64(fantomSuite.DepositingFantom * params.Ether))
	withdrawAmount := big.NewInt(int64(fantomSuite.WithdrawingFantom * params.Ether))
	withdrawingAmt := big.NewInt(0).Div(withdrawAmount, big.NewInt(1e9))

	pubKeyToAddrStr := crypto.PubkeyToAddress(fantomSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting FTM to pFTM --------------")
	// create Fantom tx to send FTM to Vault FTM
	txHash := fantomSuite.depositETH(
		fantomSuite.DepositingFantom,
		fantomSuite.IncPaymentAddrStr,
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
	)
	// txHash := common.HexToHash("0x0b297b019fde220d2b9a49231f853ab616507ec1f4791899b0b937257875fe1e")
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, txHash)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
// return
	fmt.Println("Waiting 3s for 1 block confirmation")
	time.Sleep(120 * time.Second)
	txhashInC, err := fantomSuite.callIssuingETHReq(
		fantomSuite.IncFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingftmreq",
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(50 * time.Second)
	require.Equal(fantomSuite.T(), 2, fantomSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncFTMTokenIDStr)
	fmt.Println("[INFO] FTM balance incognito after issuing step 1 : ", balpBNBS1)

	fmt.Println("------------ step 2: CONVET TO UT FTM --------------")
	fantomSuite.convertUnifiedToken(fantomSuite.IncUTFTMTokenIDStr,fantomSuite.IncFTMTokenIDStr,balpBNBS1)
	time.Sleep(30 * time.Second)
	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncFTMTokenIDStr)
	fmt.Println("[INFO] FTM balance incognito after convert : ", balpBNBS1)

	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncUTFTMTokenIDStr)
	fmt.Println("[INFO] UTFTM balance incognito after convert : ", balpBNBS1)

	fmt.Println("------------ step 3: withdrawing pFTM from Incognito to FTM --------------")
	burningRes, err := fantomSuite.callBurningUnifiedPToken(
		fantomSuite.IncUTFTMTokenIDStr,
		withdrawingAmt,
		big.NewInt(609),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		fantomSuite.IncPaymentAddrStr,
		0,
		fantomSuite.IncFTMTokenIDStr,
		false,
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)

	time.Sleep(120 * time.Second)
	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncUTFTMTokenIDStr)
	fmt.Println("[INFO] UTFTM balance incognito after issuing step 3 : ", balpBNBS1)

	fantomSuite.submitBurnProofForWithdrawalV2(
		burningTxID.(string),
		"bridgeaggGetBurnProof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
		0,
		4,
	
	)

	bal := fantomSuite.getBalanceOnETHNet(
		common.HexToAddress(fantomSuite.EtherAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", fantomSuite.ETHOwnerAddrStr)),
		fantomSuite.FTMClient,
	)
	fmt.Println("FTM balance: ", bal)
}

func (fantomSuite *FantomTestSuite) Test6ShieldConvertUnshieldLINK() {
	return
	fmt.Println("============ TEST 6 SHIELD UNSHIELD FANTOM ===========")
	fmt.Println("------------ STEP 0: declaration & initialization --------------")
	// depositAmount := big.NewInt(int64(fantomSuite.DepositingFantom * params.Ether))
	withdrawAmount := big.NewInt(int64(fantomSuite.DepositingFantom * params.Ether))
	withdrawingAmt := big.NewInt(0).Div(withdrawAmount, big.NewInt(1e9))

	pubKeyToAddrStr := crypto.PubkeyToAddress(fantomSuite.GeneratedPubKeyForSC).Hex()
	fmt.Println("------------ STEP 1: porting LINK to pLINK --------------")
	// create Fantom tx to send FTM to Vault FTM
	txHash := fantomSuite.depositERC20ToBridge(
		withdrawAmount,
		common.HexToAddress(fantomSuite.DAIAddressStr),
		fantomSuite.IncPaymentAddrStr,
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
	)
	// txHash := common.HexToHash("0x7a5c68e7c78b1e8a0c2fe99864a536af4000965e8fbb8f5c04ba6278c6d6d63d")
	time.Sleep(15 * time.Second)
	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, txHash)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof ---- : ", ethBlockHash, ethTxIdx, ethDepositProof)
	fmt.Println("Waiting 3s for 6 block confirmation")
	time.Sleep(100 * time.Second)
	txhashInC, err := fantomSuite.callIssuingETHReq(
		fantomSuite.IncDAITokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingftmreq",
	)
	require.Equal(fantomSuite.T(), nil, err)
	time.Sleep(60 * time.Second)
	require.Equal(fantomSuite.T(), 2, fantomSuite.getStatusBridgeRq(txhashInC), "Mint transaction rejected")
	balpBNBS1, _ := fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] pLINK  balance incognito after issuing step 1 : ", balpBNBS1)

	fmt.Println("------------ step 2: CONVET TO UT LINK --------------")
	fantomSuite.convertUnifiedToken(fantomSuite.IncUTDAITokenIDStr,fantomSuite.IncDAITokenIDStr,balpBNBS1)
	time.Sleep(30 * time.Second)
	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncDAITokenIDStr)
	fmt.Println("[INFO] LINK balance incognito after convert : ", balpBNBS1)

	balpBNBS1, _ = fantomSuite.getBalanceTokenIncAccount(fantomSuite.IncPrivKeyStr,fantomSuite.IncUTDAITokenIDStr)
	fmt.Println("[INFO] UT LINK balance incognito after convert : ", balpBNBS1)

	amount := big.NewInt(0).Div(withdrawingAmt, big.NewInt(5))

	fmt.Println("------------ step 3: unshield to queue --------------")
	burningRes, err := fantomSuite.callBurningUnifiedPToken(
		fantomSuite.IncUTDAITokenIDStr,
		amount,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		fantomSuite.IncPaymentAddrStr,
		0,
		"264cc6c3d58c7f7f0c0570d5aca9e4b33fb3e0b8bcd103df25314660e02d19db",
		false,
	)
	require.Equal(fantomSuite.T(), nil, err)


	burningRes, err = fantomSuite.callBurningUnifiedPToken(
		fantomSuite.IncUTDAITokenIDStr,
		amount,
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		fantomSuite.IncPaymentAddrStr,
		0,
		"264cc6c3d58c7f7f0c0570d5aca9e4b33fb3e0b8bcd103df25314660e02d19db",
		false,
	)
	require.Equal(fantomSuite.T(), nil, err)

	fmt.Println("------------ step 4: withdrawing pLINK from Incognito to LINK --------------")
	burningRes, err = fantomSuite.callBurningUnifiedPToken(
		fantomSuite.IncUTDAITokenIDStr,
		big.NewInt(0).Mul(amount, big.NewInt(3)),
		big.NewInt(0),
		pubKeyToAddrStr[2:],
		"bridgeaggUnshield",
		fantomSuite.IncPaymentAddrStr,
		0,
		fantomSuite.IncDAITokenIDStr,
		false,
	)
	require.Equal(fantomSuite.T(), nil, err)
	burningTxID, found := burningRes["TxID"]
	require.Equal(fantomSuite.T(), true, found)
	time.Sleep(30 * time.Second)
// return
	fantomSuite.submitBurnProofForWithdrawalV2(
		burningTxID.(string),
		"bridgeaggGetBurnProof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
		0,
		4,
	)

	bal := fantomSuite.getBalanceOnETHNet(
		common.HexToAddress(fantomSuite.DAIAddressStr),
		common.HexToAddress(fmt.Sprintf("0x%s", fantomSuite.ETHOwnerAddrStr)),
		fantomSuite.FTMClient,
	)
	fmt.Println("LINK balance: ", bal)

}

func (fantomSuite *FantomTestSuite) TestXXXSubmitProofToSC() {
return
// fantomSuite.submitBurnProofForDepositToSCV2(
// 	"f83e448b4248b10a4e3a5bb65a23f101fe152a5015db4054c207351d8dac295f",
// 	big.NewInt(int64(fantomSuite.ChainIDFTM)),
// 	"bridgeaggGetBurnProof",
// 	fantomSuite.VaultFTMAddr,
// 	fantomSuite.FTMClient,
// 	2,
// 	4,
// )

fantomSuite.submitBurnProofForWithdrawalV2(
	"f83e448b4248b10a4e3a5bb65a23f101fe152a5015db4054c207351d8dac295f",
	"bridgeaggGetBurnProof",
	fantomSuite.VaultFTMAddr,
	fantomSuite.FTMClient,
	fantomSuite.ChainIDFTM,
	0,
	4,
)
}


func (fantomSuite *FantomTestSuite) TestX1_get_call_data() {
	return
	fmt.Println("===== GET call data =====")
	fantomSuite.CallData(
		big.NewInt(1000000 * 1e9),
		[]common.Address{
			fantomSuite.WFTMAddr,
			common.HexToAddress("0x8658a7931F9d94180daC7135c627a27B62f199F5"),
		},
		uint(time.Now().Unix()+60000),
		false)
}


func (fantomSuite *FantomTestSuite) TestX2_trade_FTM_LIQD_new_flow() {
	// return
	fmt.Println("===== TestX2 TRADE FTM LIQD NEW FLOW =====")
	fmt.Println("===== GET call data =====")
	input := fantomSuite.CallData(
		big.NewInt(10000 * 1e9),
		[]common.Address{
			fantomSuite.WFTMAddr,
			common.HexToAddress("0x8658a7931F9d94180daC7135c627a27B62f199F5"),
			// common.HexToAddress("0x6346C85e28b1E3276828E8153aFA21BA8F522A5e"),
			
		},
		uint(time.Now().Unix()+60000),
		false)

	fmt.Println(" ===== CALL BURNING ======")
	txhash, err := fantomSuite.callBurningDapp(
		fantomSuite.IncUTFTMTokenIDStr,
		big.NewInt(10000),
		"0xcE40cE511A5D084017DBee7e3fF3e455ea32D85c",
		// "0x0000000000000000000000000000000000000000",
		fantomSuite.PanackeTradeDeployedAddr.String(),
		"bridgeaggBurnForCall",
		fantomSuite.IncPaymentReceiverStr,
		0,
		fantomSuite.IncFTMTokenIDStr,
		4,
		input,
		"0x8658a7931F9d94180daC7135c627a27B62f199F5",
	)
	require.Equal(fantomSuite.T(), nil, err)
	// return
	time.Sleep(30 * time.Second)
	fmt.Println(" ==== SUBMIT TRADE FTM LIQD ===")
	tx_submit_trade := fantomSuite.submitBurnProofForWithdrawalNewDapp(
		txhash,
		"bridgeaggGetBurnProof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
		0,
		4,
	)

	time.Sleep(40 * time.Second)
	fmt.Println("==== WITHDRAW ======")

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, tx_submit_trade)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(150 * time.Second)

	// old shield flow
	_, err = fantomSuite.callIssuingETHReq(
		fantomSuite.IncLIQDTokenIDStr,
		// tradingSuite.IncAVAXTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingftmreq",
	)

}


func (fantomSuite *FantomTestSuite) TestX3_trade_LIQD_FTM_new_flow() {
	// return
	fmt.Println("===== TestX2 TRADE LIQD FTM  NEW FLOW =====")
	fmt.Println("===== GET call data =====")
	input := fantomSuite.CallData(
		big.NewInt(10000 * 1e9),
		[]common.Address{
			common.HexToAddress("0x8658a7931F9d94180daC7135c627a27B62f199F5"),
			fantomSuite.WFTMAddr,
		},
		uint(time.Now().Unix()+60000),
		false)

	fmt.Println(" ===== CALL BURNING ======")
	txhash, err := fantomSuite.callBurningDapp(
		fantomSuite.IncLIQDTokenIDStr,
		big.NewInt(10000),
		// "0xcE40cE511A5D084017DBee7e3fF3e455ea32D85c",
		"0x0000000000000000000000000000000000000000",
		fantomSuite.PanackeTradeDeployedAddr.String(),
		"bridgeaggBurnForCall",
		fantomSuite.IncPaymentReceiverStr,
		0,
		fantomSuite.IncLIQDTokenIDStr,
		4,
		input,
		"0x0000000000000000000000000000000000000000",
	)
	require.Equal(fantomSuite.T(), nil, err)
	// return
	time.Sleep(30 * time.Second)
	fmt.Println(" ==== SUBMIT TRADE LIQD FTM ===")
	tx_submit_trade := fantomSuite.submitBurnProofForWithdrawalNewDapp(
		txhash,
		"bridgeaggGetBurnProof",
		fantomSuite.VaultFTMAddr,
		fantomSuite.FTMClient,
		fantomSuite.ChainIDFTM,
		0,
		4,
	)

	time.Sleep(40 * time.Second)
	fmt.Println("==== WITHDRAW ======")

	_, ethBlockHash, ethTxIdx, ethDepositProof, err := getETHDepositProof(fantomSuite.FTMHost, tx_submit_trade)
	require.Equal(fantomSuite.T(), nil, err)
	fmt.Println("depositProof by emitting withdarawal req: ", ethBlockHash, ethTxIdx, ethDepositProof)

	fmt.Println("Waiting 90s for 15 blocks confirmation")
	time.Sleep(150 * time.Second)

	// old shield flow
	_, err = fantomSuite.callIssuingETHReq(
		fantomSuite.IncLIQDTokenIDStr,
		// tradingSuite.IncFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"createandsendtxwithissuingftmreq",
	)


	// // // new flow unified shield
	_, err = fantomSuite.callIssuingUnifiedPtokenReq(
		fantomSuite.IncUTFTMTokenIDStr,
		ethDepositProof,
		ethBlockHash,
		ethTxIdx,
		"bridgeaggShield",
		4,
		fantomSuite.IncFTMTokenIDStr,
	)
	require.Equal(fantomSuite.T(), nil, err)
	// time.Sleep(60 * time.Second)
}