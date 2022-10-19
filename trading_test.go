package bridge

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"time"

	"github.com/incognitochain/bridge-eth/bridge/prveth"
	"github.com/incognitochain/bridge-eth/bridge/vault"
	"github.com/incognitochain/bridge-eth/common/base58"
	"github.com/incognitochain/bridge-eth/consensus/signatureschemes/bridgesig"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/sha3"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"

	"github.com/incognitochain/bridge-eth/erc20"
	"github.com/incognitochain/bridge-eth/rpccaller"
	"github.com/stretchr/testify/require"

	// iCommon "github.com/incognitochain/incognito-chain/common"
)

const (
    EXECUTE_PREFIX         = 0
    REQ_WITHDRAW_PREFIX    = 1
    BSC_EXECUTE_PREFIX     = 2
    BSC_REQ_WITHDRAW_PREFIX = 3
    PLG_EXECUTE_PREFIX     = 4
    PLG_REQ_WITHDRAW_PREFIX = 5
    FTM_EXECUTE_PREFIX     = 6
    FTM_REQ_WITHDRAW_PREFIX = 7
    AURORA_EXECUTE_PREFIX     = 8
    AURORA_REQ_WITHDRAW_PREFIX = 9
    AVAX_EXECUTE_PREFIX     = 10
    AVAX_REQ_WITHDRAW_PREFIX = 11
)


type TnBalanceIncAccount struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}
type StatusBridgeRq struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type IssuingUTokenRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}
type BurningForDepositUTokenToSCRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type IssuingETHRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type BurningForDepositToSCRes struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type callBurningDapp struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}

type ConvertUT struct {
	rpccaller.RPCBaseRes
	Result interface{} `json:"Result"`
}


// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TradingTestSuite struct {
	suite.Suite
	IncBurningAddrStr string
	IncPrivKeyStr     string
	IncPaymentAddrStr string
	IncPrivaKeyReceiverStr string
	IncPaymentReceiverStr  string

	GeneratedPrivKeyForSC ecdsa.PrivateKey
	GeneratedPubKeyForSC  ecdsa.PublicKey

	IncEtherTokenIDStr   string
	IncUTEtherTokenIDStr string
	IncDAITokenIDStr     string
	IncUSDTTokenIDStr    string
	IncUTUSDTTokenIDStr  string
	IncWETHTokenIDStr    string
	IncPRVTokenIDStr     string
	IncPDEXTokenIDStr    string

	IncFTMTokenIDStr   string
	IncUTFTMTokenIDStr string
	IncMATICTokenIDStr   string
	IncUTMATICTokenIDStr string

	IncUTDAITokenIDStr string

	IncBridgeHost string
	IncRPCHost    string

	ETHRegulatorPrivKeyStr string
	ETHRegulatorPrivKey *ecdsa.PrivateKey

	EtherAddressStr string
	DAIAddressStr   string
	SAIAddressStr   string
	USDTAddressStr  string

	ETHPrivKeyStr   string
	ETHOwnerAddrStr string

	ETHHost    string
	ETHPrivKey *ecdsa.PrivateKey
	ETHClient  *ethclient.Client

	BSCHost   string
	BSCClient *ethclient.Client

	PLGHost   string
	PLGClient *ethclient.Client

	FTMHost   string
	FTMClient *ethclient.Client

	AURORAHost   string
	AURORAClient *ethclient.Client

	AVAXHost   string
	AVAXClient *ethclient.Client

	ChainIDETH uint
	ChainIDBSC uint
	ChainIDPLG uint
	ChainIDFTM uint
	ChainIDAURORA uint
	ChainIDAVAX   uint

	VaultAddr            common.Address
	VaultBSCAddr         common.Address
	VaultPLGAddr         common.Address
	KBNTradeDeployedAddr common.Address
	PRVERC20Addr         common.Address
	PRVBEP20Addr         common.Address
	PDEXERC20Addr        common.Address
	PDEXBEP20Addr        common.Address

	VaultFTMAddr         common.Address
	VaultAVAXAddr    	common.Address
	VaultAURORAAddr 	 common.Address

	KyberContractAddr common.Address

}

// Make sure that VariableThatShouldStartAtFive is set to five
// before each test
func (tradingSuite *TradingTestSuite) SetupSuite() {
	fmt.Println("Setting up the suite...")

	// 0x kovan env
	tradingSuite.IncBurningAddrStr = "12RxahVABnAVCGP3LGwCn8jkQxgw7z1x14wztHzn455TTVpi1wBq9YGwkRMQg3J4e657AbAnCvYCJSdA9czBUNuCKwGSRQt55Xwz8WA"
	// tradingSuite.IncPrivKeyStr = "112t8roafGgHL1rhAP9632Yef3sx5k8xgp8cwK4MCJsCL1UWcxXvpzg97N4dwvcD735iKf31Q2ZgrAvKfVjeSUEvnzKJyyJD3GqqSZdxN4or"
	// tradingSuite.IncPaymentAddrStr = "12svfkP6w5UDJDSCwqH978PvqiqBxKmUnA9em9yAYWYJVRv7wuXY1qhhYpPAm4BDz2mLbFrRmdK3yRhnTqJCZXKHUmoi7NV83HCH2YFpctHNaDdkSiQshsjw2UFUuwdEvcidgaKmF3VJpY5f8RdN"
	tradingSuite.IncPrivKeyStr = "112t8rnX3VTd3MTWMpfbYP8HGY4ToAaLjrmUYzfjJBrAcb8iPLkNqvVDXWrLNiFV5yb2NBpR3FDZj3VW8GcLUwRdQ61hPMWP3YrREZAZ1UbH"
	tradingSuite.IncPaymentAddrStr = "12sx5PaLhb51Hrw28QCJd68EfQF9VVbAMobmqhC7pMzZbStyidvbsZ45gDPvPuN3VfnHxiX4JvtWfvhcefgVPM8TGAWYFgjm4fkA1QmUjAdzYr6zw5h5BEbFeo7mPYTTwxRqwHEbe2FZJXhrLdJj"
	tradingSuite.IncPrivaKeyReceiverStr = "112t8roMtSrKYCL4eA4aXiQ8umGz78znHRTdMpzKAbSgL1Cj6JtbkS9i87jW1KFjbpN9fwM3PY7LNJq3QyawdHX61eTwN6beiiayMjN4yPwC" // shard 1
	tradingSuite.IncPaymentReceiverStr = "12suy7wB2qYzSdUYPrRu2Wys1ZccRDJVwoTY3TyqJjt5nNMUitutD73S5voK5dFmSRpQytiKHRiGs8CwJjx2ZZCjbSCkWSXDwMBpNpytGvKkE6Tg1MFFRNYqLkM5G1Q5a8BeAu9S8g7Dv4niWoJf"



	tradingSuite.IncEtherTokenIDStr = "ffd8d42dc40a8d166ea4848baf8b5f6e9fe0e9c30d60062eb7d44a8df9e00854"   //ETH
	tradingSuite.IncUTEtherTokenIDStr = "b366fa400c36e6bbcf24ac3e99c90406ddc64346ab0b7ba21e159b83d938812d" //UT ETH
	tradingSuite.IncMATICTokenIDStr = "dae027b21d8d57114da11209dce8eeb587d01adf59d4fc356a8be5eedc146859"   //matic
	tradingSuite.IncUTMATICTokenIDStr = "f5d88e2e3c8f02d6dc1e01b54c90f673d730bef7d941aeec81ad1e1db690961f" //UT matic
	// tradingSuite.IncUSDTTokenIDStr = "f2e94c9cd7d003c45ab517e38c1210380195a1088989309d5dd0e421c9ab8f5c"   //DAI xxxx
	// tradingSuite.IncUTUSDTTokenIDStr = "0d953a47a7a488cee562e64c80c25d3dbe29d3b477ccd2b54408c0553a93f126" //DAI UT  xxx
	tradingSuite.IncUSDTTokenIDStr = "d43a67133bba907d04691c2b0e918c48b04db1a6ac2d03dd10f42b70422f12d6"   //link
	tradingSuite.IncUTUSDTTokenIDStr = "b35756452dc1fa1260513fa121c20c2b516a8645f8d496fa4235274dac0b1b52" //link UT
	// tradingSuite.IncUSDTTokenIDStr = "38fc5ad8434ef02ea77c860eb9d6824485de3d68b3be8455842a5bbf7b0940a5" // usdt bcs
	tradingSuite.IncWETHTokenIDStr = "a697a5c08d173de37372a20946e37d9e4adeeba68571b29b8ca4a2e1c3fc27fa"

	tradingSuite.IncDAITokenIDStr = "264cc6c3d58c7f7f0c0570d5aca9e4b33fb3e0b8bcd103df25314660e02d19db"
	tradingSuite.EtherAddressStr = "0x0000000000000000000000000000000000000000"
	tradingSuite.SAIAddressStr = "0xc4375b7de8af5a38a93548eb8453a498222c4ff2"
	tradingSuite.USDTAddressStr = "0x326c977e6efc84e512bb9c30f76e30c160ed06fb"  // LINK PLG
	tradingSuite.DAIAddressStr = "0xfaFedb041c0DD4fA2Dc0d87a6B0979Ee6FA7af5F"  // LINK FTM

	tradingSuite.IncFTMTokenIDStr = "6eed691cb14d11066f939630ff647f5f1c843a8f964d9a4d295fa9cd1111c474"
	tradingSuite.IncUTFTMTokenIDStr = "ebc1c1b5819aa5647192aefd729ef18cd8894d22656e8add678c0aef93e404d4"
	
	tradingSuite.IncDAITokenIDStr = "b5e3a7e43548a442bebab76c715d81e31edf2df0bbbdc89a26ac1f206baafe78" //LINK
	tradingSuite.IncUTDAITokenIDStr = "b35756452dc1fa1260513fa121c20c2b516a8645f8d496fa4235274dac0b1b52" //UT LINK

	tradingSuite.ETHPrivKeyStr = "a5ae26c7154410df235bc8669ffd27c0fc9d3068c21e469a4cc68165c68cd5cb"
	tradingSuite.ETHOwnerAddrStr = "cE40cE511A5D084017DBee7e3fF3e455ea32D85c"
	tradingSuite.ETHRegulatorPrivKeyStr = "98452cb9c013387c2f5806417fe198a0de014594678e2f9d3223d7e7e921b04d"



	tradingSuite.ETHHost = "https://eth-goerli.g.alchemy.com/v2/USn7_Xr8zRh4vcFu0C0d9dSdwcrd3VHF"
	tradingSuite.BSCHost = "https://data-seed-prebsc-1-s1.binance.org:8545"
	tradingSuite.PLGHost = "https://matic-mumbai.chainstacklabs.com"
	// tradingSuite.PLGHost = "https://rpc-mumbai.maticvigil.com"
	// tradingSuite.PLGHost = "https://polygon-mumbai.g.alchemy.com/v2/qOolRKBtov0fk1r657-8qWmLrcfZCw87" 
	tradingSuite.FTMHost = "https://rpc.testnet.fantom.network/"
	tradingSuite.AVAXHost = "https://api.avax-test.network/ext/C/rpc"
	tradingSuite.AURORAHost = "https://aurora-testnet.infura.io/v3/1138a1e99b154b10bae5c382ad894361"

	// tradingSuite.IncBridgeHost = "http://172.105.114.134:8334"
	// tradingSuite.IncRPCHost = "http://172.105.114.134:8334"

	// tradingSuite.IncBridgeHost = "http://51.91.220.58:8334"
	// tradingSuite.IncRPCHost = "http://51.91.220.58:8334"

	tradingSuite.IncBridgeHost = "http://51.83.36.184:9334" // testnet 2
	tradingSuite.IncRPCHost = "http://51.83.36.184:9334"    // testnet 2

	// tradingSuite.IncBridgeHost = "https://lb-fullnode.incognito.org/fullnode" // xxxx
	// tradingSuite.IncRPCHost = "https://lb-fullnode.incognito.org/fullnode"    // xxxx


	tradingSuite.VaultAddr = common.HexToAddress("0xc157CC3077ddfa425bae12d2F3002668971A4e3d") //testnet
	tradingSuite.VaultBSCAddr = common.HexToAddress("0x3534C0a523b3A862c06C8CAF61de230f9b408f51")  //testnet
	tradingSuite.VaultPLGAddr = common.HexToAddress("0x76318093c374e39B260120EBFCe6aBF7f75c8D28")  // testnet 
	tradingSuite.VaultFTMAddr = common.HexToAddress("0x76318093c374e39B260120EBFCe6aBF7f75c8D28") // testnet
	tradingSuite.VaultAVAXAddr = common.HexToAddress("0x01f6549BeF494C8b0B00C2790577AcC1A3Fa0Bd0") // testnet
	tradingSuite.VaultAURORAAddr = common.HexToAddress("0x4cF1d43999606858BaC64B7DbFC196fb4A6853af")  //testnet


	// tradingSuite.VaultAddr = common.HexToAddress("0xBD3b2Ab0c7332C1Ed9D48DaCd5EFa57CdcCDFCe1") //local
	// tradingSuite.VaultBSCAddr = common.HexToAddress("0xB12F9C1399f74F6Dd12fa54a38DE9c1E972D664A")  //local
	// tradingSuite.VaultPLGAddr = common.HexToAddress("0xbad24d417EACA27F32BDC4C1780a18dD9b4F64Ac")  // local 
	// tradingSuite.VaultFTMAddr = common.HexToAddress("0x31BA5CB6f295821A9ef8b9858839ee0e130DBDD2") // local
	// tradingSuite.VaultAVAXAddr = common.HexToAddress("0x5Abb847DA2119eD13F74730476883d4102771c4A") // local
	// tradingSuite.VaultAURORAAddr = common.HexToAddress("0xebB26286fD65221E3F7d7232FFab2FF7693AA93C")  //local

	tradingSuite.PRVERC20Addr = common.HexToAddress("0xf4933b0288644778f6f2264EaB009fD04fF669a1")
	tradingSuite.PRVBEP20Addr = common.HexToAddress("0x5A15626f6beA715870D46f43f50bE9821368963f")
	tradingSuite.PDEXERC20Addr = common.HexToAddress("0x9c59b98fcC33f2859A2aB11BC2aAfDcf513b6c33")
	tradingSuite.PDEXBEP20Addr = common.HexToAddress("0xa43F2911dF4a560A1F687Eba359D047753Cd9BD9")


	tradingSuite.ChainIDBSC = 97
	tradingSuite.ChainIDETH = 5
	tradingSuite.ChainIDPLG = 80001
	tradingSuite.ChainIDFTM = 4002 // testnet
	tradingSuite.ChainIDAURORA = 1313161555
	tradingSuite.ChainIDAVAX = 43113


	// generate a new keys pair for SCZ
	tradingSuite.genKeysPairForSC()

	// connect to ethereum network
	tradingSuite.connectToETH()
}

func (tradingSuite *TradingTestSuite) TearDownSuite() {
	fmt.Println("Tearing down the suite...")
	tradingSuite.ETHClient.Close()
}

func (tradingSuite *TradingTestSuite) SetupTest() {
	fmt.Println("Setting up the test...")
}

func (tradingSuite *TradingTestSuite) TearDownTest() {
	fmt.Println("Tearing down the test...")
}

func (tradingSuite *TradingTestSuite) TestTradingTestSuite() {
	fmt.Println("This is generic test suite")
}

func (tradingSuite *TradingTestSuite) getBalanceOnETHNet(
	tokenAddr common.Address,
	ownerAddr common.Address,
	client *ethclient.Client,
) *big.Int {
	if tokenAddr.Hex() == tradingSuite.EtherAddressStr {
		balance, err := client.BalanceAt(context.Background(), ownerAddr, nil)
		require.Equal(tradingSuite.T(), nil, err)
		return balance
	}
	// erc20 token
	instance, err := erc20.NewErc20(tokenAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	balance, err := instance.BalanceOf(&bind.CallOpts{}, ownerAddr)
	require.Equal(tradingSuite.T(), nil, err)
	return balance
}

func (tradingSuite *TradingTestSuite) connectToETH() {
	privKeyHex := tradingSuite.ETHPrivKeyStr
	privKey, err := crypto.HexToECDSA(privKeyHex)
	require.Equal(tradingSuite.T(), nil, err)

	fmt.Printf("Sign Txs with address: %s\n", crypto.PubkeyToAddress(privKey.PublicKey).Hex())

	network := "development"
	fmt.Printf("Connecting to network %s\n", network)
	client, err := ethclient.Dial(tradingSuite.ETHHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.ETHClient = client

	client, err = ethclient.Dial(tradingSuite.BSCHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.BSCClient = client

	client, err = ethclient.Dial(tradingSuite.PLGHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.PLGClient = client

	client, err = ethclient.Dial(tradingSuite.FTMHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.FTMClient = client

	client, err = ethclient.Dial(tradingSuite.AURORAHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.AURORAClient = client

	client, err = ethclient.Dial(tradingSuite.AVAXHost)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.AVAXClient = client

	tradingSuite.ETHPrivKey = privKey
	privKey, err = crypto.HexToECDSA(tradingSuite.ETHRegulatorPrivKeyStr)
	require.Equal(tradingSuite.T(), nil, err)
	tradingSuite.ETHRegulatorPrivKey = privKey

}

func (tradingSuite *TradingTestSuite) depositETH(
	amt float64,
	incPaymentAddrStr string,
	vaultAddr common.Address,
	client *ethclient.Client,
) common.Hash {
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)
	chainID, err := client.ChainID(auth.Context)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, chainID)
	require.Equal(tradingSuite.T(), nil, err)
	auth.Value = big.NewInt(int64(amt * params.Ether))

	key := make([]byte, 32)
	_, err = rand.Read(key)
	txId := toByte32(key)
	require.Equal(tradingSuite.T(), nil, err)
	signBytes, err := SignDataToShield(txId, tradingSuite.ETHRegulatorPrivKey, auth.From)
	require.Equal(tradingSuite.T(), nil, err)
	tx, err := c.Deposit(auth, incPaymentAddrStr, txId, signBytes)
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()

	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("deposited, txHash: %x\n", txHash[:])
	return txHash
}


func (tradingSuite *TradingTestSuite) depositERC20ToBridge(
	amt *big.Int,
	tokenAddr common.Address,
	incPaymentAddrStr string,
	vaultAddr common.Address,
	client *ethclient.Client,
	chainID uint,
) common.Hash {
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(chainID)))
	require.Equal(tradingSuite.T(), nil, err)
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	erc20Token, _ := erc20.NewErc20(tokenAddr, client)
	auth.GasPrice = big.NewInt(10e10)
	tx2, apprErr := erc20Token.Approve(auth, vaultAddr, amt)
	tx2Hash := tx2.Hash()
	fmt.Printf("Approve tx, txHash: %x\n", tx2Hash[:])
	require.Equal(tradingSuite.T(), nil, apprErr)
	time.Sleep(15 * time.Second)
	auth.GasPrice = big.NewInt(1e10)
	auth.GasLimit = uint64(200000)
	fmt.Println("Starting deposit erc20 to vault contract")
	key := make([]byte, 32)
	_, err = rand.Read(key)
	txId := toByte32(key)
	require.Equal(tradingSuite.T(), nil, err)
	signBytes, err := SignDataToShield(txId, tradingSuite.ETHRegulatorPrivKey, auth.From)
	require.Equal(tradingSuite.T(), nil, err)
	tx, err := c.DepositERC20(auth, tokenAddr, amt, incPaymentAddrStr, txId, signBytes)

	require.Equal(tradingSuite.T(), nil, err)
	fmt.Println("Finished deposit erc20 to vault contract")
	txHash := tx.Hash()

	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("deposited erc20 token to bridge, txHash: %x\n", txHash[:])
	return txHash
}




func (tradingSuite *TradingTestSuite) callIssuingETHReq(
	incTokenIDStr string,
	ethDepositProof []string,
	ethBlockHash string,
	ethTxIdx uint,
	method string,
) (string, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"IncTokenID": incTokenIDStr,
		"BlockHash":  ethBlockHash,
		"ProofStrs":  ethDepositProof,
		"TxIndex":    ethTxIdx,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
	}
	var res IssuingETHRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		method,
		params,
		&res,
	)
	if err != nil {
		return "", err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return "", errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{})["TxID"].(string), nil
}

func (tradingSuite *TradingTestSuite) callBurningPToken(
	incTokenIDStr string,
	amount *big.Int,
	remoteAddrStr string,
	burningMethod string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"TokenID":     incTokenIDStr,
		"TokenTxType": 1,
		"TokenName":   "",
		"TokenSymbol": "",
		"TokenAmount": amount.Uint64(),
		"TokenReceivers": map[string]uint64{
			tradingSuite.IncBurningAddrStr: amount.Uint64(),
		},
		"RemoteAddress": remoteAddrStr,
		"Privacy":       true,
		"TokenFee":      0,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		-1,
		0,
		meta,
		"",
		0,
	}
	var res BurningForDepositToSCRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		burningMethod,
		params,
		&res,
	)
	if err != nil {
		fmt.Println("calling burning ptokens err: ", err)
		return nil, err
	}
	bb, _ := json.Marshal(res)
	fmt.Println("calling burning ptokens res: ", string(bb))
	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (tradingSuite *TradingTestSuite) callBurningPRV(
	amount *big.Int,
	remoteAddrStr string,
	burningMethod string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"TokenID":     tradingSuite.IncPRVTokenIDStr,
		"TokenTxType": 1,
		"TokenName":   "",
		"TokenSymbol": "",
		"TokenAmount": amount.Uint64(),
		"TokenReceivers": map[string]uint64{
			tradingSuite.IncBurningAddrStr: amount.Uint64(),
		},
		"RemoteAddress": remoteAddrStr,
		"Privacy":       true,
		"TokenFee":      0,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		map[string]interface{}{
			tradingSuite.IncBurningAddrStr: amount.Uint64(),
		},
		-1,
		0,
		meta,
		"",
		0,
	}
	var res BurningForDepositToSCRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		burningMethod,
		params,
		&res,
	)
	if err != nil {
		fmt.Println("calling burning ptokens err: ", err)
		return nil, err
	}
	bb, _ := json.Marshal(res)
	fmt.Println("calling burning ptokens res: ", string(bb))
	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (tradingSuite *TradingTestSuite) submitBurnProofForDepositToSC(
	burningTxIDStr string,
	chainID *big.Int,
	method string,
	vaultAddr common.Address,
	client *ethclient.Client,
) {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, method)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, chainID)
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)
	auth.GasLimit = uint64(2000000) 
	tx, err := SubmitBurnProof(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func (tradingSuite *TradingTestSuite) submitBurnProofForWithdrawal(
	burningTxIDStr string,
	method string,
	vaultAddr common.Address,
	client *ethclient.Client,
	chainID uint,
) {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, method)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(chainID)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)
	auth.GasLimit = uint64(200000) // for FTM testnet
	tx, err := Withdraw(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func (tradingSuite *TradingTestSuite) submitBurnProofForDepositToSCV2(
	burningTxIDStr string,
	chainID *big.Int,
	method string,
	vaultAddr common.Address,
	client *ethclient.Client,
	DataIndex int,
	NetworkID int,
) {
	proof, err := getAndDecodeBurnProofV3(tradingSuite.IncBridgeHost, burningTxIDStr, method,DataIndex,NetworkID)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, chainID)
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)
	auth.GasLimit = uint64(200000) // for FTM testnet

	tx, err := SubmitBurnProof(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func (tradingSuite *TradingTestSuite) submitBurnProofForWithdrawalV2(
	burningTxIDStr string,
	method string,
	vaultAddr common.Address,
	client *ethclient.Client,
	chainID uint,
	DataIndex int,
	NetworkID int,
) {

	proof, err := getAndDecodeBurnProofV3(tradingSuite.IncBridgeHost, burningTxIDStr, method,DataIndex,NetworkID)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(chainID)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)
	auth.GasLimit = uint64(2000000) // for FTM testnet
	tx, err := Withdraw(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
}

func (tradingSuite *TradingTestSuite) submitBurnProofForMintPRV(
	burningTxIDStr string,
	contractAddress common.Address,
	method string,
	clientInst *ethclient.Client,
	chainID int64,
) {
	proof, err := getAndDecodeBurnProofV2(tradingSuite.IncBridgeHost, burningTxIDStr, method)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := prveth.NewPrveth(contractAddress, clientInst)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(chainID))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)
	tx, err := SubmitMintPRVProof(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(clientInst, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("mint evm prv, txHash: %x\n", txHash[:])
}

func (tradingSuite *TradingTestSuite) genKeysPairForSC() {
	incPriKeyBytes, _, err := base58.Base58Check{}.Decode(tradingSuite.IncPrivKeyStr)
	require.Equal(tradingSuite.T(), nil, err)

	tradingSuite.GeneratedPrivKeyForSC, tradingSuite.GeneratedPubKeyForSC = bridgesig.KeyGen(incPriKeyBytes)

}

func randomizeTimestamp() string {
	randomTime := rand.Int63n(time.Now().Unix()-94608000) + 94608000
	randomNow := time.Unix(randomTime, 0)
	return randomNow.String()
}

func rawsha3(b []byte) []byte {
	hashF := sha3.NewLegacyKeccak256()
	hashF.Write(b)
	buf := hashF.Sum(nil)
	return buf
}

func rlpHash(x interface{}) (h common.Hash) {
	hw := sha3.NewLegacyKeccak256()
	rlp.Encode(hw, x)
	hw.Sum(h[:0])
	return h
}

func (tradingSuite *TradingTestSuite) getDepositedBalance(
	ethTokenAddrStr string,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultAddr, tradingSuite.ETHClient)
	require.Equal(tradingSuite.T(), nil, err)
	token := common.HexToAddress(ethTokenAddrStr)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) getDepositedBalanceBSC(
	token common.Address,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultBSCAddr, tradingSuite.BSCClient)
	require.Equal(tradingSuite.T(), nil, err)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) getDepositedBalancePLG(
	token common.Address,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultPLGAddr, tradingSuite.PLGClient)
	require.Equal(tradingSuite.T(), nil, err)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) getDepositedBalanceFTM(
	token common.Address,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultFTMAddr, tradingSuite.FTMClient)
	require.Equal(tradingSuite.T(), nil, err)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) requestWithdraw(
	withdrawalETHTokenIDStr string,
	amount *big.Int,
	client *ethclient.Client,
	chainID *big.Int,
	vaultAddrr common.Address,
	signaturePrefix uint8,
) common.Hash {
	c, err := vault.NewVault(vaultAddrr, client)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, chainID)
	require.Equal(tradingSuite.T(), nil, err)
	token := common.HexToAddress(withdrawalETHTokenIDStr)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix:    signaturePrefix,
		Token:     token,
		Timestamp: timestamp,
		Amount:    amount,
	}
	tempData, _ := vaultAbi.Pack("_buildSignRequestWithdraw", psData, tradingSuite.IncPaymentAddrStr)
	data := rawsha3(tempData[4:])
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	auth.GasPrice = big.NewInt(1e10)
	auth.GasLimit = uint64(200000) // for aurora

	key := make([]byte, 32)
	_, err = rand.Read(key)
	txId := toByte32(key)
	require.Equal(tradingSuite.T(), nil, err)
	signBytesRegulator, err := SignDataToShield(txId, tradingSuite.ETHRegulatorPrivKey, auth.From)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.RequestWithdraw(auth, tradingSuite.IncPaymentAddrStr, token, amount, signBytes, timestamp, txId, signBytesRegulator)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("request withdrawal, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) burnPRV(
	amt float64,
	incPaymentAddrStr string,
	contractAddress common.Address,
	clientInst *ethclient.Client,
	chainID int64,
) common.Hash {
	c, err := prveth.NewPrveth(contractAddress, clientInst)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(chainID))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(1e10)

	tx, err := c.Burn(auth, incPaymentAddrStr, big.NewInt(int64(amt*float64(1e9))))
	require.Equal(tradingSuite.T(), nil, err)
	txHash := tx.Hash()

	if err := wait(clientInst, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burn prv token, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) callIssuingPRVReq(
	incTokenIDStr string,
	ethDepositProof []string,
	ethBlockHash string,
	ethTxIdx uint,
	methodName string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"IncTokenID": incTokenIDStr,
		"BlockHash":  ethBlockHash,
		"ProofStrs":  ethDepositProof,
		"TxIndex":    ethTxIdx,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
	}
	var res IssuingETHRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		methodName,
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}


func (tradingSuite *TradingTestSuite) callIssuingUnifiedPtokenReq(
	incUTTokenIDStr string,
	ethDepositProof []string,
	ethBlockHash string,
	ethTxIdx uint,
	method string,
	NetworkID uint,
	incTokenIDStr string,
) (string, error) {
	rpcClient := rpccaller.NewRPCClient()
	metaSub := map[string]interface{}{
		"BlockHash": ethBlockHash,
		"TxIndex":   ethTxIdx,
		"Proof":     ethDepositProof,
		"IncTokenID":incTokenIDStr,
		"NetworkID": NetworkID,
	}
	meta := map[string]interface{}{
		"Data": []interface{}{
			metaSub,
		},
		"UnifiedTokenID": incUTTokenIDStr,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		-1,
		1,
		meta,
	}
	var res IssuingUTokenRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		method,
		params,
		&res,
	)
	if err != nil {
		return "", err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return "", errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{})["TxID"].(string), nil
}

func (tradingSuite *TradingTestSuite) callBurningUnifiedPToken(
	incUnifiedTokenIDStr string,
	amount *big.Int,
	expectedAmount *big.Int,
	remoteAddrStr string,
	burningMethod string,
	receiverFeeAddr string,
	receiverFeeAmt uint64,
	incTokenIDStr string,
	IsDepositToSC bool,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()

	metaSub := map[string]interface{}{
		"BurningAmount":  amount.Uint64(),
		"RemoteAddress":  remoteAddrStr,
		"IncTokenID": incTokenIDStr,
		"ExpectedAmount": expectedAmount.Uint64(),
	}
	meta2 := map[string]interface{}{
		"Data": []interface{}{
			metaSub,
		},
		"UnifiedTokenID": incUnifiedTokenIDStr,
		"IsDepositToSC":  IsDepositToSC,
		"TokenReceivers": map[string]interface{}{
			receiverFeeAddr: receiverFeeAmt,
		},
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		-1,
		1,
		meta2,
	}
	var res BurningForDepositUTokenToSCRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		burningMethod,
		params,
		&res,
	)
	if err != nil {
		fmt.Println("calling burning ptokens err: ", err)
		return nil, err
	}
	bb, _ := json.Marshal(res)
	fmt.Println("calling burning ptokens res: ", string(bb))
	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}


func (tradingSuite *TradingTestSuite) convertUnifiedToken(
	UnifiedTokenID string,
	IncTokenID string,
	Amount uint64,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"UnifiedTokenID": UnifiedTokenID,
		"TokenID":  IncTokenID,
		"Amount":    Amount,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		-1,
		1,
		meta,
	}
	var res ConvertUT
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		"bridgeaggConvert",
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (tradingSuite *TradingTestSuite) getStatusBridgeRq(txhash string) int {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"TxReqID": txhash,
	}
	params := []interface{}{
		meta,
	}
	var res StatusBridgeRq
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		"getbridgereqwithstatus",
		params,
		&res,
	)
	if err != nil {
		return 0
	}
	return int(res.Result.(float64))
}

func (tradingSuite *TradingTestSuite) getBalanceTokenIncAccount(
	IncPrivKeyStr string,
	ethTokenAddrStr string,
) (uint64, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{
		IncPrivKeyStr,
		ethTokenAddrStr,
	}
	var res TnBalanceIncAccount
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		"getbalanceprivacycustomtoken",
		params,
		&res,
	)
	if err != nil {
		return 0, err
	}
	//fmt.Println(res.Result)
	return uint64(res.Result.(float64)), nil
}

func (tradingSuite *TradingTestSuite) getBalancePrvIncAccount(
	IncPrivKeyStr string,
) (uint64, error) {
	rpcClient := rpccaller.NewRPCClient()
	params := []interface{}{
		IncPrivKeyStr,
	}
	var res TnBalanceIncAccount
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		"getbalancebyprivatekey",
		params,
		&res,
	)
	if err != nil {
		return 0, err
	}
	//fmt.Println(res.Result)
	return uint64(res.Result.(float64)), nil
}

func (tradingSuite *TradingTestSuite) callBurningDapp(
	incUnifiedTokenIDStr string,
	amount *big.Int,
	remoteAddrStr string,
	ExternalCallAddress string,
	burningMethod string,
	receiverFeeAddr string,
	receiverFeeAmt uint64,
	incTokenIDStr string,
	networkId uint64,
	ExternalCalldata []byte,
	ReceiveToken string,
) (string, error) {
	rpcClient := rpccaller.NewRPCClient()
	metaSub := map[string]interface{}{
		"BurningAmount":  amount.Uint64(),
		"ExternalCallAddress":ExternalCallAddress[2:],
		"IncTokenID": incTokenIDStr,
		"ExternalNetworkID": networkId,
		"ExternalCalldata": string(common.Bytes2Hex(ExternalCalldata)),
		"ReceiveToken" : ReceiveToken[2:],
		"WithdrawAddress":  remoteAddrStr[2:],
	}
	meta2 := map[string]interface{}{
		"Data": []interface{}{
			metaSub,
		},
		"BurnTokenID": incUnifiedTokenIDStr,
		"TokenReceivers": map[string]interface{}{
			receiverFeeAddr: receiverFeeAmt,
		},
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		5,
		1,
		meta2,
	}
	var res callBurningDapp
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		burningMethod,
		params,
		&res,
	)
	if err != nil {
		fmt.Println("calling burning ptokens err: ", err)
		return "", err
	}
	bb, _ := json.Marshal(res)
	fmt.Println("calling burning ptokens res: ", string(bb))
	if res.RPCError != nil {
		return "", errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{})["TxID"].(string), nil
}

func (tradingSuite *TradingTestSuite) submitBurnProofForWithdrawalNewDapp(
	burningTxIDStr string,
	method string,
	vaultAddr common.Address,
	client *ethclient.Client,
	chainID uint,
	DataIndex int,
	NetworkID int,
) common.Hash {

	proof, err := getAndDecodeBurnProofV3(tradingSuite.IncBridgeHost, burningTxIDStr, method,DataIndex,NetworkID)
	require.Equal(tradingSuite.T(), nil, err)

	// Get contract instance
	c, err := vault.NewVault(vaultAddr, client)
	require.Equal(tradingSuite.T(), nil, err)

	// Burn
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, big.NewInt(int64(chainID)))
	require.Equal(tradingSuite.T(), nil, err)
	auth.GasPrice = big.NewInt(10e10)
	auth.GasLimit = uint64(500000) 
	tx, err := ExecuteWithBurnProof(c, auth, proof)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("burned, txHash: %x\n", txHash[:])
	return txHash
}

func (tradingSuite *TradingTestSuite) requestWithdrawCompliance(
	withdrawalETHTokenIDStr string,
	amount *big.Int,
	client *ethclient.Client,
	chainID *big.Int,
	vaultAddrr common.Address,
	signaturePrefix uint8,
) common.Hash {
	c, err := vault.NewVault(vaultAddrr, client)
	require.Equal(tradingSuite.T(), nil, err)
	auth, err := bind.NewKeyedTransactorWithChainID(tradingSuite.ETHPrivKey, chainID)
	require.Equal(tradingSuite.T(), nil, err)
	token := common.HexToAddress(withdrawalETHTokenIDStr)
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix:    signaturePrefix,
		Token:     token,
		Timestamp: timestamp,
		Amount:    amount,
	}
	tempData, _ := vaultAbi.Pack("_buildSignRequestWithdraw", psData, tradingSuite.IncPaymentAddrStr)
	data := rawsha3(tempData[4:])
	signBytes, _ := crypto.Sign(data, &tradingSuite.GeneratedPrivKeyForSC)
	key := make([]byte, 32)
	_, err = rand.Read(key)
	txId := toByte32(key)
	require.Equal(tradingSuite.T(), nil, err)
	keySign, err := crypto.HexToECDSA("")
	require.Equal(tradingSuite.T(), nil, err)
	signBytesRegulator, err := SignDataToShield(txId, keySign, auth.From)
	require.Equal(tradingSuite.T(), nil, err)

	tx, err := c.RequestWithdraw(auth, tradingSuite.IncPaymentAddrStr, token, amount, signBytes, timestamp, txId, signBytesRegulator)
	require.Equal(tradingSuite.T(), nil, err)

	txHash := tx.Hash()
	if err := wait(client, txHash); err != nil {
		require.Equal(tradingSuite.T(), nil, err)
	}
	fmt.Printf("request withdrawal, txHash: %x\n", txHash[:])
	return txHash
}

func SignDataToWithdraw(token common.Address, key *ecdsa.PrivateKey, amount *big.Int) ([]byte, []byte, error) {
	timestamp := []byte(randomizeTimestamp())
	vaultAbi, _ := abi.JSON(strings.NewReader(vault.VaultHelperMetaData.ABI))
	// vaultAbi,  _ := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	psData := vault.VaultHelperPreSignData{
		Prefix:    REQ_WITHDRAW_PREFIX,
		Token:     token,
		Timestamp: timestamp,
		Amount:    amount,
	}
	tempData, err := vaultAbi.Pack("_buildSignRequestWithdraw", psData, "")
	if err != nil {
		return nil, nil, err
	}
	data := rawsha3(tempData[4:])
	signBytes, _ := crypto.Sign(data, key)
	return signBytes, timestamp, nil
}

func SignDataToShield(txId [32]byte, key *ecdsa.PrivateKey, from common.Address) ([]byte, error) {
	vaultHelperAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperMetaData.ABI))
	// vaultHelperAbi, err := abi.JSON(strings.NewReader(vault.VaultHelperABI))
	if err != nil {
		return nil, err
	}
	tempData, err := vaultHelperAbi.Pack("_buildSignShield", from, txId)
	if err != nil {
		return nil, err
	}
	data := rawsha3(tempData[4:])
	signBytes, err := crypto.Sign(data, key)
	if err != nil {
		return nil, err
	}
	return signBytes, nil
}

func (tradingSuite *TradingTestSuite) getDepositedBalanceAVAX(
	token common.Address,
	ownerAddrStr string,
) *big.Int {
	c, err := vault.NewVault(tradingSuite.VaultAVAXAddr, tradingSuite.AVAXClient)
	require.Equal(tradingSuite.T(), nil, err)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) getDepositedBalanceWithParams(
	token common.Address,
	ownerAddrStr string,
	vaultAddress common.Address,
	client *ethclient.Client,
) *big.Int {
	c, err := vault.NewVault(vaultAddress, client)
	require.Equal(tradingSuite.T(), nil, err)
	owner := common.HexToAddress(ownerAddrStr)
	bal, err := c.GetDepositedBalance(nil, token, owner)
	require.Equal(tradingSuite.T(), nil, err)
	fmt.Printf("deposited balance: %d\n", bal)
	return bal
}

func (tradingSuite *TradingTestSuite) callIssuingAURORAReq(
	incTokenIDStr string,
	txHash string,
	method string,
) (map[string]interface{}, error) {
	rpcClient := rpccaller.NewRPCClient()
	meta := map[string]interface{}{
		"IncTokenID": incTokenIDStr,
		"TxHash":     txHash,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		5,
		-1,
		meta,
	}
	var res IssuingETHRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		method,
		params,
		&res,
	)
	if err != nil {
		return nil, err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return nil, errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{}), nil
}

func (tradingSuite *TradingTestSuite) callIssuingAUDORAUnifiedPtokenReq(
	incUTTokenIDStr string,
	txhash string,
	method string,
	NetworkID uint,
	incTokenIDStr string,
) (string, error) {
	rpcClient := rpccaller.NewRPCClient()
	metaSub2 := []interface{}{
		txhash,
	}
	metaSub := map[string]interface{}{
		"Proof":     metaSub2,
		"IncTokenID":incTokenIDStr,
		"NetworkID": NetworkID,
	}
	meta := map[string]interface{}{
		"Data": []interface{}{
			metaSub,
		},
		"UnifiedTokenID": incUTTokenIDStr,
	}
	params := []interface{}{
		tradingSuite.IncPrivKeyStr,
		nil,
		-1,
		1,
		meta,
	}
	var res IssuingUTokenRes
	err := rpcClient.RPCCall(
		"",
		tradingSuite.IncRPCHost,
		"",
		method,
		params,
		&res,
	)
	if err != nil {
		return "", err
	}

	response, _ := json.Marshal(res)
	fmt.Println("get response", string(response))

	if res.RPCError != nil {
		return "", errors.New(res.RPCError.Message)
	}
	return res.Result.(map[string]interface{})["TxID"].(string), nil
}