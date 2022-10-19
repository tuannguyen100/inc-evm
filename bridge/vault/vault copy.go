// // Code generated - DO NOT EDIT.
// // This file is a generated binding and any manual changes will be lost.

package vault

// import (
// 	"math/big"
// 	"strings"

// 	ethereum "github.com/ethereum/go-ethereum"
// 	"github.com/ethereum/go-ethereum/accounts/abi"
// 	"github.com/ethereum/go-ethereum/accounts/abi/bind"
// 	"github.com/ethereum/go-ethereum/common"
// 	"github.com/ethereum/go-ethereum/core/types"
// 	"github.com/ethereum/go-ethereum/event"
// )

// // Reference imports to suppress errors if they are not otherwise used.
// var (
// 	_ = big.NewInt
// 	_ = strings.NewReader
// 	_ = ethereum.NotFound
// 	_ = bind.Bind
// 	_ = common.Big1
// 	_ = types.BloomLookup
// 	_ = event.NewSubscription
// )

// // VaultBurnInstData is an auto generated low-level Go binding around an user-defined struct.
// type VaultBurnInstData struct {
// 	Meta   uint8
// 	Shard  uint8
// 	Token  common.Address
// 	To     common.Address
// 	Amount *big.Int
// 	Itx    [32]byte
// }

// // VaultRedepositOptions is an auto generated low-level Go binding around an user-defined struct.
// type VaultRedepositOptions struct {
// 	RedepositToken      common.Address
// 	RedepositIncAddress []byte
// 	WithdrawAddress     common.Address
// }

// // AddressABI is the input ABI used to generate the binding from.
// const AddressABI = "[]"

// // AddressBin is the compiled bytecode used for deploying new contracts.
// var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122056923f7f32633c310e8d7f6c1b5fdd5c2fb7886b521493707effd7774c4b7fce64736f6c634300060c0033"

// // DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
// func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
// 	parsed, err := abi.JSON(strings.NewReader(AddressABI))
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}

// 	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}
// 	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
// }

// // Address is an auto generated Go binding around an Ethereum contract.
// type Address struct {
// 	AddressCaller     // Read-only binding to the contract
// 	AddressTransactor // Write-only binding to the contract
// 	AddressFilterer   // Log filterer for contract events
// }

// // AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
// type AddressCaller struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
// type AddressTransactor struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
// type AddressFilterer struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // AddressSession is an auto generated Go binding around an Ethereum contract,
// // with pre-set call and transact options.
// type AddressSession struct {
// 	Contract     *Address          // Generic contract binding to set the session for
// 	CallOpts     bind.CallOpts     // Call options to use throughout this session
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// // with pre-set call options.
// type AddressCallerSession struct {
// 	Contract *AddressCaller // Generic contract caller binding to set the session for
// 	CallOpts bind.CallOpts  // Call options to use throughout this session
// }

// // AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// // with pre-set transact options.
// type AddressTransactorSession struct {
// 	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
// 	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
// }

// // AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
// type AddressRaw struct {
// 	Contract *Address // Generic contract binding to access the raw methods on
// }

// // AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
// type AddressCallerRaw struct {
// 	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
// }

// // AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
// type AddressTransactorRaw struct {
// 	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
// }

// // NewAddress creates a new instance of Address, bound to a specific deployed contract.
// func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
// 	contract, err := bindAddress(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
// }

// // NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
// func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
// 	contract, err := bindAddress(address, caller, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &AddressCaller{contract: contract}, nil
// }

// // NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
// func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
// 	contract, err := bindAddress(address, nil, transactor, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &AddressTransactor{contract: contract}, nil
// }

// // NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
// func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
// 	contract, err := bindAddress(address, nil, nil, filterer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &AddressFilterer{contract: contract}, nil
// }

// // bindAddress binds a generic wrapper to an already deployed contract.
// func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
// 	parsed, err := abi.JSON(strings.NewReader(AddressABI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Address.Contract.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Address.Contract.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Address.Contract.contract.Transact(opts, method, params...)
// }

// // CountersABI is the input ABI used to generate the binding from.
// const CountersABI = "[]"

// // CountersBin is the compiled bytecode used for deploying new contracts.
// var CountersBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122034f494eec3e4eba26bb165441a354905eeb3c14260a955e7b0678299ad9c0acc64736f6c634300060c0033"

// // DeployCounters deploys a new Ethereum contract, binding an instance of Counters to it.
// func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counters, error) {
// 	parsed, err := abi.JSON(strings.NewReader(CountersABI))
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}

// 	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CountersBin), backend)
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}
// 	return address, tx, &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
// }

// // Counters is an auto generated Go binding around an Ethereum contract.
// type Counters struct {
// 	CountersCaller     // Read-only binding to the contract
// 	CountersTransactor // Write-only binding to the contract
// 	CountersFilterer   // Log filterer for contract events
// }

// // CountersCaller is an auto generated read-only Go binding around an Ethereum contract.
// type CountersCaller struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // CountersTransactor is an auto generated write-only Go binding around an Ethereum contract.
// type CountersTransactor struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // CountersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
// type CountersFilterer struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // CountersSession is an auto generated Go binding around an Ethereum contract,
// // with pre-set call and transact options.
// type CountersSession struct {
// 	Contract     *Counters         // Generic contract binding to set the session for
// 	CallOpts     bind.CallOpts     // Call options to use throughout this session
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // CountersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// // with pre-set call options.
// type CountersCallerSession struct {
// 	Contract *CountersCaller // Generic contract caller binding to set the session for
// 	CallOpts bind.CallOpts   // Call options to use throughout this session
// }

// // CountersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// // with pre-set transact options.
// type CountersTransactorSession struct {
// 	Contract     *CountersTransactor // Generic contract transactor binding to set the session for
// 	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
// }

// // CountersRaw is an auto generated low-level Go binding around an Ethereum contract.
// type CountersRaw struct {
// 	Contract *Counters // Generic contract binding to access the raw methods on
// }

// // CountersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
// type CountersCallerRaw struct {
// 	Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
// }

// // CountersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
// type CountersTransactorRaw struct {
// 	Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
// }

// // NewCounters creates a new instance of Counters, bound to a specific deployed contract.
// func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
// 	contract, err := bindCounters(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
// }

// // NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
// func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
// 	contract, err := bindCounters(address, caller, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &CountersCaller{contract: contract}, nil
// }

// // NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
// func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
// 	contract, err := bindCounters(address, nil, transactor, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &CountersTransactor{contract: contract}, nil
// }

// // NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
// func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
// 	contract, err := bindCounters(address, nil, nil, filterer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &CountersFilterer{contract: contract}, nil
// }

// // bindCounters binds a generic wrapper to an already deployed contract.
// func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
// 	parsed, err := abi.JSON(strings.NewReader(CountersABI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Counters.Contract.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Counters.Contract.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Counters.Contract.contract.Transact(opts, method, params...)
// }

// // IERC20ABI is the input ABI used to generate the binding from.
// const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// // IERC20FuncSigs maps the 4-byte function signature to its string representation.
// var IERC20FuncSigs = map[string]string{
// 	"dd62ed3e": "allowance(address,address)",
// 	"095ea7b3": "approve(address,uint256)",
// 	"70a08231": "balanceOf(address)",
// 	"313ce567": "decimals()",
// 	"18160ddd": "totalSupply()",
// 	"a9059cbb": "transfer(address,uint256)",
// 	"23b872dd": "transferFrom(address,address,uint256)",
// }

// // IERC20 is an auto generated Go binding around an Ethereum contract.
// type IERC20 struct {
// 	IERC20Caller     // Read-only binding to the contract
// 	IERC20Transactor // Write-only binding to the contract
// 	IERC20Filterer   // Log filterer for contract events
// }

// // IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
// type IERC20Caller struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
// type IERC20Transactor struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
// type IERC20Filterer struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // IERC20Session is an auto generated Go binding around an Ethereum contract,
// // with pre-set call and transact options.
// type IERC20Session struct {
// 	Contract     *IERC20           // Generic contract binding to set the session for
// 	CallOpts     bind.CallOpts     // Call options to use throughout this session
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// // with pre-set call options.
// type IERC20CallerSession struct {
// 	Contract *IERC20Caller // Generic contract caller binding to set the session for
// 	CallOpts bind.CallOpts // Call options to use throughout this session
// }

// // IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// // with pre-set transact options.
// type IERC20TransactorSession struct {
// 	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
// type IERC20Raw struct {
// 	Contract *IERC20 // Generic contract binding to access the raw methods on
// }

// // IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
// type IERC20CallerRaw struct {
// 	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
// }

// // IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
// type IERC20TransactorRaw struct {
// 	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
// }

// // NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
// func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
// 	contract, err := bindIERC20(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
// }

// // NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
// func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
// 	contract, err := bindIERC20(address, caller, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IERC20Caller{contract: contract}, nil
// }

// // NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
// func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
// 	contract, err := bindIERC20(address, nil, transactor, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IERC20Transactor{contract: contract}, nil
// }

// // NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
// func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
// 	contract, err := bindIERC20(address, nil, nil, filterer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IERC20Filterer{contract: contract}, nil
// }

// // bindIERC20 binds a generic wrapper to an already deployed contract.
// func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
// 	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _IERC20.Contract.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _IERC20.Contract.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _IERC20.Contract.contract.Transact(opts, method, params...)
// }

// // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
// //
// // Solidity: function allowance(address owner, address spender) view returns(uint256)
// func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
// 	var out []interface{}
// 	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
// //
// // Solidity: function allowance(address owner, address spender) view returns(uint256)
// func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
// 	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
// }

// // Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
// //
// // Solidity: function allowance(address owner, address spender) view returns(uint256)
// func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
// 	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
// }

// // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// //
// // Solidity: function balanceOf(address account) view returns(uint256)
// func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
// 	var out []interface{}
// 	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// //
// // Solidity: function balanceOf(address account) view returns(uint256)
// func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
// 	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
// }

// // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// //
// // Solidity: function balanceOf(address account) view returns(uint256)
// func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
// 	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
// }

// // Decimals is a free data retrieval call binding the contract method 0x313ce567.
// //
// // Solidity: function decimals() view returns(uint256)
// func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
// 	var out []interface{}
// 	err := _IERC20.contract.Call(opts, &out, "decimals")

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // Decimals is a free data retrieval call binding the contract method 0x313ce567.
// //
// // Solidity: function decimals() view returns(uint256)
// func (_IERC20 *IERC20Session) Decimals() (*big.Int, error) {
// 	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
// }

// // Decimals is a free data retrieval call binding the contract method 0x313ce567.
// //
// // Solidity: function decimals() view returns(uint256)
// func (_IERC20 *IERC20CallerSession) Decimals() (*big.Int, error) {
// 	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
// }

// // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
// //
// // Solidity: function totalSupply() view returns(uint256)
// func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
// 	var out []interface{}
// 	err := _IERC20.contract.Call(opts, &out, "totalSupply")

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
// //
// // Solidity: function totalSupply() view returns(uint256)
// func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
// 	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
// }

// // TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
// //
// // Solidity: function totalSupply() view returns(uint256)
// func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
// 	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
// }

// // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// //
// // Solidity: function approve(address spender, uint256 amount) returns()
// func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.contract.Transact(opts, "approve", spender, amount)
// }

// // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// //
// // Solidity: function approve(address spender, uint256 amount) returns()
// func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
// }

// // Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
// //
// // Solidity: function approve(address spender, uint256 amount) returns()
// func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
// }

// // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// //
// // Solidity: function transfer(address recipient, uint256 amount) returns()
// func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
// }

// // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// //
// // Solidity: function transfer(address recipient, uint256 amount) returns()
// func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
// }

// // Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
// //
// // Solidity: function transfer(address recipient, uint256 amount) returns()
// func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
// }

// // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// //
// // Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
// func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
// }

// // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// //
// // Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
// func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
// }

// // TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
// //
// // Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
// func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
// }

// // IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
// type IERC20ApprovalIterator struct {
// 	Event *IERC20Approval // Event containing the contract specifics and raw log

// 	contract *bind.BoundContract // Generic contract to use for unpacking event data
// 	event    string              // Event name to use for unpacking event data

// 	logs chan types.Log        // Log channel receiving the found contract events
// 	sub  ethereum.Subscription // Subscription for errors, completion and termination
// 	done bool                  // Whether the subscription completed delivering logs
// 	fail error                 // Occurred error to stop iteration
// }

// // Next advances the iterator to the subsequent event, returning whether there
// // are any more events found. In case of a retrieval or parsing error, false is
// // returned and Error() can be queried for the exact failure.
// func (it *IERC20ApprovalIterator) Next() bool {
// 	// If the iterator failed, stop iterating
// 	if it.fail != nil {
// 		return false
// 	}
// 	// If the iterator completed, deliver directly whatever's available
// 	if it.done {
// 		select {
// 		case log := <-it.logs:
// 			it.Event = new(IERC20Approval)
// 			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 				it.fail = err
// 				return false
// 			}
// 			it.Event.Raw = log
// 			return true

// 		default:
// 			return false
// 		}
// 	}
// 	// Iterator still in progress, wait for either a data or an error event
// 	select {
// 	case log := <-it.logs:
// 		it.Event = new(IERC20Approval)
// 		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 			it.fail = err
// 			return false
// 		}
// 		it.Event.Raw = log
// 		return true

// 	case err := <-it.sub.Err():
// 		it.done = true
// 		it.fail = err
// 		return it.Next()
// 	}
// }

// // Error returns any retrieval or parsing error occurred during filtering.
// func (it *IERC20ApprovalIterator) Error() error {
// 	return it.fail
// }

// // Close terminates the iteration process, releasing any pending underlying
// // resources.
// func (it *IERC20ApprovalIterator) Close() error {
// 	it.sub.Unsubscribe()
// 	return nil
// }

// // IERC20Approval represents a Approval event raised by the IERC20 contract.
// type IERC20Approval struct {
// 	Owner   common.Address
// 	Spender common.Address
// 	Value   *big.Int
// 	Raw     types.Log // Blockchain specific contextual infos
// }

// // FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
// //
// // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
// func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

// 	var ownerRule []interface{}
// 	for _, ownerItem := range owner {
// 		ownerRule = append(ownerRule, ownerItem)
// 	}
// 	var spenderRule []interface{}
// 	for _, spenderItem := range spender {
// 		spenderRule = append(spenderRule, spenderItem)
// 	}

// 	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
// }

// // WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
// //
// // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
// func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

// 	var ownerRule []interface{}
// 	for _, ownerItem := range owner {
// 		ownerRule = append(ownerRule, ownerItem)
// 	}
// 	var spenderRule []interface{}
// 	for _, spenderItem := range spender {
// 		spenderRule = append(spenderRule, spenderItem)
// 	}

// 	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return event.NewSubscription(func(quit <-chan struct{}) error {
// 		defer sub.Unsubscribe()
// 		for {
// 			select {
// 			case log := <-logs:
// 				// New log arrived, parse the event and forward to the user
// 				event := new(IERC20Approval)
// 				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
// 					return err
// 				}
// 				event.Raw = log

// 				select {
// 				case sink <- event:
// 				case err := <-sub.Err():
// 					return err
// 				case <-quit:
// 					return nil
// 				}
// 			case err := <-sub.Err():
// 				return err
// 			case <-quit:
// 				return nil
// 			}
// 		}
// 	}), nil
// }

// // ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
// //
// // Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
// func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
// 	event := new(IERC20Approval)
// 	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
// 		return nil, err
// 	}
// 	event.Raw = log
// 	return event, nil
// }

// // IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
// type IERC20TransferIterator struct {
// 	Event *IERC20Transfer // Event containing the contract specifics and raw log

// 	contract *bind.BoundContract // Generic contract to use for unpacking event data
// 	event    string              // Event name to use for unpacking event data

// 	logs chan types.Log        // Log channel receiving the found contract events
// 	sub  ethereum.Subscription // Subscription for errors, completion and termination
// 	done bool                  // Whether the subscription completed delivering logs
// 	fail error                 // Occurred error to stop iteration
// }

// // Next advances the iterator to the subsequent event, returning whether there
// // are any more events found. In case of a retrieval or parsing error, false is
// // returned and Error() can be queried for the exact failure.
// func (it *IERC20TransferIterator) Next() bool {
// 	// If the iterator failed, stop iterating
// 	if it.fail != nil {
// 		return false
// 	}
// 	// If the iterator completed, deliver directly whatever's available
// 	if it.done {
// 		select {
// 		case log := <-it.logs:
// 			it.Event = new(IERC20Transfer)
// 			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 				it.fail = err
// 				return false
// 			}
// 			it.Event.Raw = log
// 			return true

// 		default:
// 			return false
// 		}
// 	}
// 	// Iterator still in progress, wait for either a data or an error event
// 	select {
// 	case log := <-it.logs:
// 		it.Event = new(IERC20Transfer)
// 		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 			it.fail = err
// 			return false
// 		}
// 		it.Event.Raw = log
// 		return true

// 	case err := <-it.sub.Err():
// 		it.done = true
// 		it.fail = err
// 		return it.Next()
// 	}
// }

// // Error returns any retrieval or parsing error occurred during filtering.
// func (it *IERC20TransferIterator) Error() error {
// 	return it.fail
// }

// // Close terminates the iteration process, releasing any pending underlying
// // resources.
// func (it *IERC20TransferIterator) Close() error {
// 	it.sub.Unsubscribe()
// 	return nil
// }

// // IERC20Transfer represents a Transfer event raised by the IERC20 contract.
// type IERC20Transfer struct {
// 	From  common.Address
// 	To    common.Address
// 	Value *big.Int
// 	Raw   types.Log // Blockchain specific contextual infos
// }

// // FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
// //
// // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
// func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

// 	var fromRule []interface{}
// 	for _, fromItem := range from {
// 		fromRule = append(fromRule, fromItem)
// 	}
// 	var toRule []interface{}
// 	for _, toItem := range to {
// 		toRule = append(toRule, toItem)
// 	}

// 	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
// }

// // WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
// //
// // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
// func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

// 	var fromRule []interface{}
// 	for _, fromItem := range from {
// 		fromRule = append(fromRule, fromItem)
// 	}
// 	var toRule []interface{}
// 	for _, toItem := range to {
// 		toRule = append(toRule, toItem)
// 	}

// 	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return event.NewSubscription(func(quit <-chan struct{}) error {
// 		defer sub.Unsubscribe()
// 		for {
// 			select {
// 			case log := <-logs:
// 				// New log arrived, parse the event and forward to the user
// 				event := new(IERC20Transfer)
// 				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
// 					return err
// 				}
// 				event.Raw = log

// 				select {
// 				case sink <- event:
// 				case err := <-sub.Err():
// 					return err
// 				case <-quit:
// 					return nil
// 				}
// 			case err := <-sub.Err():
// 				return err
// 			case <-quit:
// 				return nil
// 			}
// 		}
// 	}), nil
// }

// // ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
// //
// // Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
// func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
// 	event := new(IERC20Transfer)
// 	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
// 		return nil, err
// 	}
// 	event.Raw = log
// 	return event, nil
// }

// // IncognitoABI is the input ABI used to generate the binding from.
// const IncognitoABI = "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// // IncognitoFuncSigs maps the 4-byte function signature to its string representation.
// var IncognitoFuncSigs = map[string]string{
// 	"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
// }

// // Incognito is an auto generated Go binding around an Ethereum contract.
// type Incognito struct {
// 	IncognitoCaller     // Read-only binding to the contract
// 	IncognitoTransactor // Write-only binding to the contract
// 	IncognitoFilterer   // Log filterer for contract events
// }

// // IncognitoCaller is an auto generated read-only Go binding around an Ethereum contract.
// type IncognitoCaller struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // IncognitoTransactor is an auto generated write-only Go binding around an Ethereum contract.
// type IncognitoTransactor struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // IncognitoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
// type IncognitoFilterer struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // IncognitoSession is an auto generated Go binding around an Ethereum contract,
// // with pre-set call and transact options.
// type IncognitoSession struct {
// 	Contract     *Incognito        // Generic contract binding to set the session for
// 	CallOpts     bind.CallOpts     // Call options to use throughout this session
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // IncognitoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// // with pre-set call options.
// type IncognitoCallerSession struct {
// 	Contract *IncognitoCaller // Generic contract caller binding to set the session for
// 	CallOpts bind.CallOpts    // Call options to use throughout this session
// }

// // IncognitoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// // with pre-set transact options.
// type IncognitoTransactorSession struct {
// 	Contract     *IncognitoTransactor // Generic contract transactor binding to set the session for
// 	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
// }

// // IncognitoRaw is an auto generated low-level Go binding around an Ethereum contract.
// type IncognitoRaw struct {
// 	Contract *Incognito // Generic contract binding to access the raw methods on
// }

// // IncognitoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
// type IncognitoCallerRaw struct {
// 	Contract *IncognitoCaller // Generic read-only contract binding to access the raw methods on
// }

// // IncognitoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
// type IncognitoTransactorRaw struct {
// 	Contract *IncognitoTransactor // Generic write-only contract binding to access the raw methods on
// }

// // NewIncognito creates a new instance of Incognito, bound to a specific deployed contract.
// func NewIncognito(address common.Address, backend bind.ContractBackend) (*Incognito, error) {
// 	contract, err := bindIncognito(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Incognito{IncognitoCaller: IncognitoCaller{contract: contract}, IncognitoTransactor: IncognitoTransactor{contract: contract}, IncognitoFilterer: IncognitoFilterer{contract: contract}}, nil
// }

// // NewIncognitoCaller creates a new read-only instance of Incognito, bound to a specific deployed contract.
// func NewIncognitoCaller(address common.Address, caller bind.ContractCaller) (*IncognitoCaller, error) {
// 	contract, err := bindIncognito(address, caller, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IncognitoCaller{contract: contract}, nil
// }

// // NewIncognitoTransactor creates a new write-only instance of Incognito, bound to a specific deployed contract.
// func NewIncognitoTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoTransactor, error) {
// 	contract, err := bindIncognito(address, nil, transactor, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IncognitoTransactor{contract: contract}, nil
// }

// // NewIncognitoFilterer creates a new log filterer instance of Incognito, bound to a specific deployed contract.
// func NewIncognitoFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoFilterer, error) {
// 	contract, err := bindIncognito(address, nil, nil, filterer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &IncognitoFilterer{contract: contract}, nil
// }

// // bindIncognito binds a generic wrapper to an already deployed contract.
// func bindIncognito(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
// 	parsed, err := abi.JSON(strings.NewReader(IncognitoABI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Incognito *IncognitoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Incognito.Contract.IncognitoCaller.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Incognito *IncognitoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Incognito.Contract.IncognitoTransactor.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Incognito *IncognitoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Incognito.Contract.IncognitoTransactor.contract.Transact(opts, method, params...)
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Incognito *IncognitoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Incognito.Contract.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Incognito *IncognitoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Incognito.Contract.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Incognito *IncognitoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Incognito.Contract.contract.Transact(opts, method, params...)
// }

// // InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
// //
// // Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
// func (_Incognito *IncognitoCaller) InstructionApproved(opts *bind.CallOpts, arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
// 	var out []interface{}
// 	err := _Incognito.contract.Call(opts, &out, "instructionApproved", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
// //
// // Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
// func (_Incognito *IncognitoSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
// 	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
// }

// // InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
// //
// // Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
// func (_Incognito *IncognitoCallerSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
// 	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
// }

// // SafeMathABI is the input ABI used to generate the binding from.
// const SafeMathABI = "[]"

// // SafeMathBin is the compiled bytecode used for deploying new contracts.
// var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c969d5ed32a447d7bf041e77373ccd636c8d4ad43f6edcf2782fdf3102c8ecc064736f6c634300060c0033"

// // DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
// func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
// 	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}

// 	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}
// 	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
// }

// // SafeMath is an auto generated Go binding around an Ethereum contract.
// type SafeMath struct {
// 	SafeMathCaller     // Read-only binding to the contract
// 	SafeMathTransactor // Write-only binding to the contract
// 	SafeMathFilterer   // Log filterer for contract events
// }

// // SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
// type SafeMathCaller struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
// type SafeMathTransactor struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
// type SafeMathFilterer struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // SafeMathSession is an auto generated Go binding around an Ethereum contract,
// // with pre-set call and transact options.
// type SafeMathSession struct {
// 	Contract     *SafeMath         // Generic contract binding to set the session for
// 	CallOpts     bind.CallOpts     // Call options to use throughout this session
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// // with pre-set call options.
// type SafeMathCallerSession struct {
// 	Contract *SafeMathCaller // Generic contract caller binding to set the session for
// 	CallOpts bind.CallOpts   // Call options to use throughout this session
// }

// // SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// // with pre-set transact options.
// type SafeMathTransactorSession struct {
// 	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
// 	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
// }

// // SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
// type SafeMathRaw struct {
// 	Contract *SafeMath // Generic contract binding to access the raw methods on
// }

// // SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
// type SafeMathCallerRaw struct {
// 	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
// }

// // SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
// type SafeMathTransactorRaw struct {
// 	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
// }

// // NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
// func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
// 	contract, err := bindSafeMath(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
// }

// // NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
// func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
// 	contract, err := bindSafeMath(address, caller, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &SafeMathCaller{contract: contract}, nil
// }

// // NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
// func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
// 	contract, err := bindSafeMath(address, nil, transactor, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &SafeMathTransactor{contract: contract}, nil
// }

// // NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
// func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
// 	contract, err := bindSafeMath(address, nil, nil, filterer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &SafeMathFilterer{contract: contract}, nil
// }

// // bindSafeMath binds a generic wrapper to an already deployed contract.
// func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
// 	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _SafeMath.Contract.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _SafeMath.Contract.contract.Transact(opts, method, params...)
// }

// // VaultABI is the input ABI used to generate the binding from.
// const VaultABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"DepositV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"redepositIncAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"name\":\"Redeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"UpdateIncognitoProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"UpdateTokenTotal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_CALL_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_TO_CONTRACT_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CURRENT_NETWORK_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"externalCalldata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"redepositToken\",\"type\":\"address\"}],\"name\":\"_callExternal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_transferExternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"depositERC20_V2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"}],\"name\":\"deposit_V2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"executeWithBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prevVault\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseCalldataFromBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"redepositToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"redepositIncAddress\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"internalType\":\"structVault.RedepositOptions\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevVault\",\"outputs\":[{\"internalType\":\"contractWithdrawable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sigToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"submitBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositedToSCAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// // VaultFuncSigs maps the 4-byte function signature to its string representation.
// var VaultFuncSigs = map[string]string{
// 	"bd835c42": "BURN_CALL_REQUEST_METADATA_TYPE()",
// 	"568c04fd": "BURN_REQUEST_METADATA_TYPE()",
// 	"6f2cbc48": "BURN_TO_CONTRACT_REQUEST_METADATA_TYPE()",
// 	"d7200eb1": "CURRENT_NETWORK_ID()",
// 	"58bc8337": "ETH_TOKEN()",
// 	"bda9b509": "_callExternal(address,address,uint256,bytes,address)",
// 	"145e2a6b": "_transferExternal(address,address,uint256)",
// 	"70a08231": "balanceOf(address)",
// 	"a26e1186": "deposit(string)",
// 	"5a67cb87": "depositERC20(address,uint256,string)",
// 	"5c421b2f": "depositERC20_V2(address,uint256,string)",
// 	"49c5ba23": "deposit_V2(string)",
// 	"8588ccd6": "execute(address,uint256,address,address,bytes,bytes,bytes)",
// 	"3ed1b376": "executeWithBurnProof(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
// 	"cf54aaa0": "getDecimals(address)",
// 	"f75b98ce": "getDepositedBalance(address,address)",
// 	"c4d66de8": "initialize(address)",
// 	"392e53cd": "isInitialized()",
// 	"e4bd7074": "isSigDataUsed(bytes32)",
// 	"749c5f86": "isWithdrawed(bytes32)",
// 	"995fac11": "migration(address,address)",
// 	"a3f5d8cc": "notEntered()",
// 	"7e16e6e1": "parseBurnInst(bytes)",
// 	"66945b31": "parseCalldataFromBurnInst(bytes)",
// 	"fa84702e": "prevVault()",
// 	"87add440": "requestWithdraw(string,address,uint256,bytes,bytes)",
// 	"1ea1940e": "sigDataUsed(bytes32)",
// 	"3fec6b40": "sigToAddress(bytes,bytes32)",
// 	"73bf9651": "submitBurnProof(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
// 	"6304541c": "totalDepositedToSCAmount(address)",
// 	"1ed4276d": "updateAssets(address[],uint256[])",
// 	"1beb7de2": "withdraw(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
// 	"65b5a00f": "withdrawRequests(address,address)",
// 	"dca40d9e": "withdrawed(bytes32)",
// }

// // VaultBin is the compiled bytecode used for deploying new contracts.
// var VaultBin = "0x608060405234801561001057600080fd5b50614893806100206000396000f3fe6080604052600436106101fd5760003560e01c806373bf96511161010d578063bd835c42116100a0578063d7200eb11161006f578063d7200eb11461059b578063dca40d9e146105b0578063e4bd7074146105d0578063f75b98ce146105f0578063fa84702e1461061057610204565b8063bd835c4214610526578063bda9b5091461053b578063c4d66de81461055b578063cf54aaa01461057b57610204565b806387add440116100dc57806387add440146104be578063995fac11146104de578063a26e1186146104fe578063a3f5d8cc1461051157610204565b806373bf96511461043e578063749c5f861461045e5780637e16e6e11461047e5780638588ccd6146104ab57610204565b8063568c04fd116101905780636304541c1161015f5780636304541c1461038d57806365b5a00f146103ba57806366945b31146103da5780636f2cbc481461040957806370a082311461041e57610204565b8063568c04fd1461031657806358bc8337146103385780635a67cb871461034d5780635c421b2f1461036d57610204565b8063392e53cd116101cc578063392e53cd146102a15780633ed1b376146102b65780633fec6b40146102d657806349c5ba231461030357610204565b8063145e2a6b146102095780631beb7de21461022b5780631ea1940e1461024b5780631ed4276d1461028157610204565b3661020457005b600080fd5b34801561021557600080fd5b50610229610224366004613a65565b610625565b005b34801561023757600080fd5b50610229610246366004613ed5565b61070e565b34801561025757600080fd5b5061026b610266366004613ccc565b610b6b565b604051610278919061447a565b60405180910390f35b34801561028d57600080fd5b5061026b61029c366004613c54565b610b80565b3480156102ad57600080fd5b5061026b610da1565b3480156102c257600080fd5b506102296102d1366004613d23565b610db1565b3480156102e257600080fd5b506102f66102f1366004613e93565b6111fd565b60405161027891906142a5565b610229610311366004613ce4565b61128c565b34801561032257600080fd5b5061032b611372565b6040516102789190614767565b34801561034457600080fd5b506102f6611377565b34801561035957600080fd5b50610229610368366004613bfb565b61137c565b34801561037957600080fd5b50610229610388366004613bfb565b611618565b34801561039957600080fd5b506103ad6103a8366004613975565b6118c3565b6040516102789190614532565b3480156103c657600080fd5b506103ad6103d5366004613a2d565b6118d5565b3480156103e657600080fd5b506103fa6103f5366004613ce4565b6118f2565b6040516102789392919061468b565b34801561041557600080fd5b5061032b611ad5565b34801561042a57600080fd5b506103ad610439366004613975565b611ada565b34801561044a57600080fd5b50610229610459366004613ed5565b611b75565b34801561046a57600080fd5b5061026b610479366004613ccc565b611eed565b34801561048a57600080fd5b5061049e610499366004613e61565b611fa4565b604051610278919061467d565b6102296104b9366004613b23565b612026565b3480156104ca57600080fd5b506102296104d9366004614006565b6124be565b3480156104ea57600080fd5b5061026b6104f9366004613a2d565b6126e7565b61022961050c366004613ce4565b612707565b34801561051d57600080fd5b5061026b6127d8565b34801561053257600080fd5b5061032b6127e8565b34801561054757600080fd5b506103ad610556366004613aa5565b6127ed565b34801561056757600080fd5b50610229610576366004613975565b612993565b34801561058757600080fd5b5061032b610596366004613975565b612a0b565b3480156105a757600080fd5b5061032b612a89565b3480156105bc57600080fd5b5061026b6105cb366004613ccc565b612a8e565b3480156105dc57600080fd5b5061026b6105eb366004613ccc565b612aa3565b3480156105fc57600080fd5b506103ad61060b366004613a2d565b612b0a565b34801561061c57600080fd5b506102f6612c36565b3033146106326013612c45565b906106595760405162461bcd60e51b81526004016106509190614559565b60405180910390fd5b506001600160a01b038316610677576106728282612d6e565b610709565b60405163a9059cbb60e01b81526001600160a01b0384169063a9059cbb906106a590859085906004016142dd565b600060405180830381600087803b1580156106bf57600080fd5b505af11580156106d3573d6000803e3d6000fd5b505050506106df612e0a565b6106e96004612c45565b906107075760405162461bcd60e51b81526004016106509190614559565b505b505050565b600554600160a01b900460ff166107256001612c45565b906107435760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b191690558951608211156107616006612c45565b9061077f5760405162461bcd60e51b81526004016106509190614559565b50610788613693565b6107918b611fa4565b805190915060ff1660f11480156107af5750806020015160ff166001145b6107b96006612c45565b906107d75760405162461bcd60e51b81526004016106509190614559565b506107e58160a00151611eed565b156107f06005612c45565b9061080e5760405162461bcd60e51b81526004016106509190614559565b5060a081015160009081526020819052604090819020805460ff191660011790558101516001600160a01b0316610899576040808201516001600160a01b0316600090815260046020522054608082015161086891612e3e565b4710156108756007612c45565b906108935760405162461bcd60e51b81526004016106509190614559565b506109ab565b60006108a88260400151612a0b565b905060098160ff1611156108d55760808201516108cf9060081960ff841601600a0a612ea4565b60808301525b6040808301516001600160a01b031660009081526004602052205460808301516108fe91612e3e565b82604001516001600160a01b03166370a08231306040518263ffffffff1660e01b815260040161092e91906142a5565b60206040518083038186803b15801561094657600080fd5b505afa15801561095a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061097e91906140b9565b101561098a6007612c45565b906109a85760405162461bcd60e51b81526004016106509190614559565b50505b6109bd8b8b8b8b8b8b8b8b8b8b612f06565b60408101516001600160a01b0316610a6657600081606001516001600160a01b031682608001516040516109f0906142a2565b60006040518083038185875af1925050503d8060008114610a2d576040519150601f19603f3d011682016040523d82523d6000602084013e610a32565b606091505b5050905080610a416004612c45565b90610a5f5760405162461bcd60e51b81526004016106509190614559565b5050610b04565b80604001516001600160a01b031663a9059cbb826060015183608001516040518363ffffffff1660e01b8152600401610aa09291906142dd565b600060405180830381600087803b158015610aba57600080fd5b505af1158015610ace573d6000803e3d6000fd5b50505050610ada612e0a565b610ae46004612c45565b90610b025760405162461bcd60e51b81526004016106509190614559565b505b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb816040015182606001518360800151604051610b43939291906142b9565b60405180910390a150506005805460ff60a01b1916600160a01b179055505050505050505050565b60016020526000908152604090205460ff1681565b6005546000906001600160a01b031615801590610ba757506005546001600160a01b031633145b610bb1600c612c45565b90610bcf5760405162461bcd60e51b81526004016106509190614559565b50838214610bdd600a612c45565b90610bfb5760405162461bcd60e51b81526004016106509190614559565b50600560009054906101000a90046001600160a01b03166001600160a01b0316635c975abb6040518163ffffffff1660e01b815260040160206040518083038186803b158015610c4a57600080fd5b505afa158015610c5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c829190613cb0565b610c8c600d612c45565b90610caa5760405162461bcd60e51b81526004016106509190614559565b5060005b84811015610d5857610d11848483818110610cc557fe5b9050602002013560046000898986818110610cdc57fe5b9050602002016020810190610cf19190613975565b6001600160a01b0316815260208101919091526040016000205490612e3e565b60046000888885818110610d2157fe5b9050602002016020810190610d369190613975565b6001600160a01b03168152602081019190915260400160002055600101610cae565b507f6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f85858585604051610d8e94939291906143fb565b60405180910390a1506001949350505050565b600554600160a81b900460ff1681565b600554600160a01b900460ff16610dc86001612c45565b90610de65760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b19169055610dfc613693565b610e046136c8565b6060610e108e8e6118f2565b925092509250610e238360a00151611eed565b15610e2e6005612c45565b90610e4c5760405162461bcd60e51b81526004016106509190614559565b5060a083015160009081526020819052604090819020805460ff191660011790558301516001600160a01b0316610ed7576040808401516001600160a01b03166000908152600460205220546080840151610ea691612e3e565b471015610eb36007612c45565b90610ed15760405162461bcd60e51b81526004016106509190614559565b50610fe9565b6000610ee68460400151612a0b565b905060098160ff161115610f13576080840151610f0d9060081960ff841601600a0a612ea4565b60808501525b6040808501516001600160a01b03166000908152600460205220546080850151610f3c91612e3e565b84604001516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401610f6c91906142a5565b60206040518083038186803b158015610f8457600080fd5b505afa158015610f98573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fbc91906140b9565b1015610fc86007612c45565b90610fe65760405162461bcd60e51b81526004016106509190614559565b50505b61103f8e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d8d8d8d8d8d8d8d8d612f06565b604080840151606085015160808601518551935163bda9b50960e01b8152309463bda9b50994611079949093909290918891600401614310565b602060405180830381600087803b15801561109357600080fd5b505af19250505080156110c3575060408051601f3d908101601f191682019092526110c0918101906140b9565b60015b6110eb576110e38360400151836020015185608001518660a00151613001565b5050506111dd565b60408301516001600160a01b031661111a5761111583600001518460200151838760a00151613001565b6111d8565b8251604080850151905163145e2a6b60e01b8152309263145e2a6b926111449286906004016142b9565b600060405180830381600087803b15801561115e57600080fd5b505af192505050801561116f575060015b6111945761118b83600001518460200151838760a00151613001565b505050506111dd565b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb83600001518460400151836040516111cf939291906142b9565b60405180910390a15b505050505b50506005805460ff60a01b1916600160a01b179055505050505050505050565b60008060008060208601519150604086015192508560408151811061121e57fe5b602001015160f81c60f81b60f81c601b01905060018582848660405160008152602001604052604051611254949392919061453b565b6020604051602081039080840390855afa158015611276573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b600554600160a01b900460ff166112a36001612c45565b906112c15760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b191690556b033b2e3c9fd0803ce80000004711156112e96002612c45565b906113075760405162461bcd60e51b81526004016106509190614559565b507fd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851600083833461133860066130f7565b6040516113499594939291906143c2565b60405180910390a161135b60066130fb565b50506005805460ff60a01b1916600160a01b179055565b60f181565b600081565b600554600160a01b900460ff166113936001612c45565b906113b15760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b191690558360006113cb82612a0b565b90506000826001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016113fb91906142a5565b60206040518083038186803b15801561141357600080fd5b505afa158015611427573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061144b91906140b9565b90508086600960ff851611156114865760098460ff1603600a0a818161146d57fe5b04905060098460ff1603600a0a838161148257fe5b0492505b670de0b6b3a764000081111580156114a65750670de0b6b3a76400008311155b80156114c35750670de0b6b3a76400006114c08285612e3e565b11155b6114cd6003612c45565b906114eb5760405162461bcd60e51b81526004016106509190614559565b506040516323b872dd60e01b81526001600160a01b038616906323b872dd9061151c90339030908d906004016142b9565b600060405180830381600087803b15801561153657600080fd5b505af115801561154a573d6000803e3d6000fd5b50505050611556612e0a565b6115606004612c45565b9061157e5760405162461bcd60e51b81526004016106509190614559565b50876115938361158d8c611ada565b90613104565b1461159e600a612c45565b906115bc5760405162461bcd60e51b81526004016106509190614559565b507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e898888846040516115f2949392919061438c565b60405180910390a150506005805460ff60a01b1916600160a01b17905550505050505050565b600554600160a01b900460ff1661162f6001612c45565b9061164d5760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b1916905583600061166782612a0b565b90506000826001600160a01b03166370a08231306040518263ffffffff1660e01b815260040161169791906142a5565b60206040518083038186803b1580156116af57600080fd5b505afa1580156116c3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116e791906140b9565b90508086600960ff851611156117225760098460ff1603600a0a818161170957fe5b04905060098460ff1603600a0a838161171e57fe5b0492505b670de0b6b3a764000081111580156117425750670de0b6b3a76400008311155b801561175f5750670de0b6b3a764000061175c8285612e3e565b11155b6117696003612c45565b906117875760405162461bcd60e51b81526004016106509190614559565b506040516323b872dd60e01b81526001600160a01b038616906323b872dd906117b890339030908d906004016142b9565b600060405180830381600087803b1580156117d257600080fd5b505af11580156117e6573d6000803e3d6000fd5b505050506117f2612e0a565b6117fc6004612c45565b9061181a5760405162461bcd60e51b81526004016106509190614559565b50876118298361158d8c611ada565b14611834600a612c45565b906118525760405162461bcd60e51b81526004016106509190614559565b507fd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab38518988888461188260066130f7565b6040516118939594939291906143c2565b60405180910390a16118a560066130fb565b50506005805460ff60a01b1916600160a01b17905550505050505050565b60046020526000908152604090205481565b600260209081526000928352604080842090915290825290205481565b6118fa613693565b6119026136c8565b60606101288410156119146012612c45565b906119325760405162461bcd60e51b81526004016106509190614559565b5061193b613693565b8585600081811061194857fe5b919091013560f81c8252508585600181811061196057fe5b919091013560f81c60208301525060008686600281811061197d57fe5b845192013560f81c92505060ff16609e1480156119a15750816020015160ff166001145b80156119b0575060ff81166001145b6119ba6012612c45565b906119d85760405162461bcd60e51b81526004016106509190614559565b50506119e26136c8565b6119f060c36003888a6147ba565b8101906119fd9190613991565b6001600160a01b03908116604088810191909152918116875260a088019290925260808701929092529182166060860152911690830152611a4361012860c3888a6147ba565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060208201528181611a8c88610128818c6147ba565b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250969e959d50919b50939950505050505050505050565b60f381565b60006001600160a01b038216611af1575047611b70565b6040516370a0823160e01b81526001600160a01b038316906370a0823190611b1d9030906004016142a5565b60206040518083038186803b158015611b3557600080fd5b505afa158015611b49573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b6d91906140b9565b90505b919050565b600554600160a01b900460ff16611b8c6001612c45565b90611baa5760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b19169055895160821115611bc86006612c45565b90611be65760405162461bcd60e51b81526004016106509190614559565b50611bef613693565b611bf88b611fa4565b805190915060ff1660f3148015611c165750806020015160ff166001145b611c206006612c45565b90611c3e5760405162461bcd60e51b81526004016106509190614559565b50611c4c8160a00151611eed565b15611c576005612c45565b90611c755760405162461bcd60e51b81526004016106509190614559565b5060a081015160009081526020819052604090819020805460ff191660011790558101516001600160a01b0316611d00576040808201516001600160a01b03166000908152600460205220546080820151611ccf91612e3e565b471015611cdc6007612c45565b90611cfa5760405162461bcd60e51b81526004016106509190614559565b50611e12565b6000611d0f8260400151612a0b565b905060098160ff161115611d3c576080820151611d369060081960ff841601600a0a612ea4565b60808301525b6040808301516001600160a01b03166000908152600460205220546080830151611d6591612e3e565b82604001516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611d9591906142a5565b60206040518083038186803b158015611dad57600080fd5b505afa158015611dc1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611de591906140b9565b1015611df16007612c45565b90611e0f5760405162461bcd60e51b81526004016106509190614559565b50505b611e248b8b8b8b8b8b8b8b8b8b612f06565b608081015160608201516001600160a01b0390811660009081526002602090815260408083208187015190941683529290522054611e6191612e3e565b60608201516001600160a01b03908116600090815260026020908152604080832081870180518616855290835281842095909555608086015194519093168252600490522054611eb091612e3e565b6040918201516001600160a01b031660009081526004602052919091205550506005805460ff60a01b1916600160a01b1790555050505050505050565b60008181526020819052604081205460ff1615611f0c57506001611b70565b6005546001600160a01b0316611f2457506000611b70565b600554604051633a4e2fc360e11b81526001600160a01b039091169063749c5f8690611f54908590600401614532565b60206040518083038186803b158015611f6c57600080fd5b505afa158015611f80573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b6d9190613cb0565b611fac613693565b611fb4613693565b82600081518110611fc157fe5b016020015160f81c8152825183906001908110611fda57fe5b0160209081015160f81c9082015260228301516042840151606285015160828601516001600160a01b039384166040860152929091166060840152608083015260a08201529050919050565b600554600160a01b900460ff1661203d6001612c45565b9061205b5760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b1916905560006120d761207b828d88888f61315a565b8a8a8a8a6040516020016120939594939291906146f8565b60408051601f198184030181526020601f8701819004810284018101909252858352919086908690819084018382808284376000920191909152506131e092505050565b90506120e3818c613284565b6001600160a01b038082166000908152600260209081526040808320938f16835292905220548a11156121166008612c45565b906121345760405162461bcd60e51b81526004016106509190614559565b506001600160a01b038b16600090815260046020526040902054612158908b613104565b6001600160a01b03808d1660008181526004602090815260408083209590955592851681526002835283812091815291522054612195908b613104565b60026000836001600160a01b03166001600160a01b0316815260200190815260200160002060008d6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600034905060006001600160a01b03168c6001600160a01b031614156122125761220b818c612e3e565b905061234d565b6040516370a0823160e01b81528b906001600160a01b038e16906370a08231906122409030906004016142a5565b60206040518083038186803b15801561225857600080fd5b505afa15801561226c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061229091906140b9565b101561229c6007612c45565b906122ba5760405162461bcd60e51b81526004016106509190614559565b5060405163a9059cbb60e01b81526001600160a01b038d169063a9059cbb906122e9908c908f906004016142dd565b600060405180830381600087803b15801561230357600080fd5b505af1158015612317573d6000803e3d6000fd5b50505050612323612e0a565b61232d6004612c45565b9061234b5760405162461bcd60e51b81526004016106509190614559565b505b600061239f8b838b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d6133c6565b90506123fc8160026000866001600160a01b03166001600160a01b0316815260200190815260200160002060008e6001600160a01b03166001600160a01b0316815260200190815260200160002054612e3e90919063ffffffff16565b60026000856001600160a01b03166001600160a01b0316815260200190815260200160002060008d6001600160a01b03166001600160a01b031681526020019081526020016000208190555061248081600460008e6001600160a01b03166001600160a01b0316815260200190815260200160002054612e3e90919063ffffffff16565b6001600160a01b03909b1660009081526004602052604090209a909a5550506005805460ff60a01b1916600160a01b17905550505050505050505050565b600554600160a01b900460ff166124d56001612c45565b906124f35760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b19169055600061256c61251460018986868b61315a565b8a8a60405160200161252893929190614741565b60408051601f198184030181526020601f8901819004810284018101909252878352919088908890819084018382808284376000920191909152506131e092505050565b90506125788188613284565b6001600160a01b038082166000908152600260209081526040808320938b16835292905220548611156125ab6008612c45565b906125c95760405162461bcd60e51b81526004016106509190614559565b506001600160a01b038082166000908152600260209081526040808320938b16835292905220546125fa9087613104565b6001600160a01b038083166000908152600260209081526040808320938c1683529281528282209390935560049092529020546126379087613104565b6001600160a01b03881660008181526004602052604090209190915586901561268b57600061266589612a0b565b905060098160ff1611156126895760098160ff1603600a0a888161268557fe5b0491505b505b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e888b8b846040516126c0949392919061438c565b60405180910390a150506005805460ff60a01b1916600160a01b1790555050505050505050565b600360209081526000928352604080842090915290825290205460ff1681565b600554600160a01b900460ff1661271e6001612c45565b9061273c5760405162461bcd60e51b81526004016106509190614559565b506005805460ff60a01b191690556b033b2e3c9fd0803ce80000004711156127646002612c45565b906127825760405162461bcd60e51b81526004016106509190614559565b507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e60008383346040516127b9949392919061438c565b60405180910390a150506005805460ff60a01b1916600160a01b179055565b600554600160a01b900460ff1681565b609e81565b60003033146127fc6013612c45565b9061281a5760405162461bcd60e51b81526004016106509190614559565b50600061282683611ada565b905060606001600160a01b03881661284a576128438786886134ef565b90506128e8565b60405163a9059cbb60e01b81526001600160a01b0389169063a9059cbb90612878908a908a906004016142dd565b600060405180830381600087803b15801561289257600080fd5b505af11580156128a6573d6000803e3d6000fd5b505050506128b2612e0a565b6128bc6004612c45565b906128da5760405162461bcd60e51b81526004016106509190614559565b506128e5878661351d565b90505b80516040146128f76009612c45565b906129155760405162461bcd60e51b81526004016106509190614559565b506000808280602001905181019061292d9190613a00565b91509150856001600160a01b0316826001600160a01b031614801561295d57508061295b8561158d89611ada565b145b6129676009612c45565b906129855760405162461bcd60e51b81526004016106509190614559565b509998505050505050505050565b600554600160a81b900460ff16156129ab600f612c45565b906129c95760405162461bcd60e51b81526004016106509190614559565b5060058054600160a01b600160a81b6001600160a01b03199092166001600160a01b03949094169390931760ff60a81b19161760ff60a01b1916919091179055565b600080829050806001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015612a4a57600080fd5b505afa158015612a5e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a8291906140b9565b9392505050565b600181565b60006020819052908152604090205460ff1681565b60008181526001602052604081205460ff1615612ac257506001611b70565b6005546001600160a01b0316612ada57506000611b70565b60055460405163392f5c1d60e21b81526001600160a01b039091169063e4bd707490611f54908590600401614532565b6005546000906001600160a01b031615801590612b4d57506001600160a01b0380831660009081526003602090815260408083209387168352929052205460ff16155b15612c0957600554604051637badcc6760e11b8152612c02916001600160a01b03169063f75b98ce90612b8690879087906004016142f6565b60206040518083038186803b158015612b9e57600080fd5b505afa158015612bb2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612bd691906140b9565b6001600160a01b0380851660009081526002602090815260408083209389168352929052205490612e3e565b9050611286565b506001600160a01b0380821660009081526002602090815260408083209386168352929052205492915050565b6005546001600160a01b031681565b60606000826013811115612c5557fe5b60408051600a808252818301909252919250906060908260208201818036833701905050905060005b60ff841615612ccc578151600a60ff959095168581049560018401939106916030830160f81b9185918110612caf57fe5b60200101906001600160f81b031916908160001a90535050612c7e565b6060816001016001600160401b0381118015612ce757600080fd5b506040519080825280601f01601f191660200182016040528015612d12576020820181803683370190505b50905060005b828111612d63578381840381518110612d2d57fe5b602001015160f81c60f81b828281518110612d4457fe5b60200101906001600160f81b031916908160001a905350600101612d18565b509695505050505050565b80471015612d8e5760405162461bcd60e51b8152600401610650906145c9565b6000826001600160a01b031682604051612da7906142a2565b60006040518083038185875af1925050503d8060008114612de4576040519150601f19603f3d011682016040523d82523d6000602084013e612de9565b606091505b50509050806107095760405162461bcd60e51b81526004016106509061456c565b6000803d8015612e215760208114612e2a57612e36565b60019150612e36565b60206000803e60005191505b501515905090565b6000828201838110801590612e535750828110155b6040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b81525090612e9c5760405162461bcd60e51b81526004016106509190614559565b509392505050565b6000828202831580612e53575082848281612ebb57fe5b04146040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b81525090612e9c5760405162461bcd60e51b81526004016106509190614559565b60008a8a604051602001612f1b929190614280565b604051602081830303815290604052805190602001209050612f3b61355f565b6001600160a01b031663f65d21166001838d8d8d8d8d8d8d8d8d6040518c63ffffffff1660e01b8152600401612f7b9b9a99989796959493929190614485565b60206040518083038186803b158015612f9357600080fd5b505afa158015612fa7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612fcb9190613cb0565b612fd56006612c45565b90612ff35760405162461bcd60e51b81526004016106509190614559565b505050505050505050505050565b816001600160a01b03851661304e576b033b2e3c9fd0803ce800000047111561302a6002612c45565b906130485760405162461bcd60e51b81526004016106509190614559565b506130b4565b600061305986612a0b565b905060098160ff16111561307d5760098160ff1603600a0a828161307957fe5b0491505b670de0b6b3a76400008211156130936003612c45565b906130b15760405162461bcd60e51b81526004016106509190614559565b50505b7eb45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d858583856040516130e89493929190614355565b60405180910390a15050505050565b5490565b80546001019055565b6000828211156040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b815250906131535760405162461bcd60e51b81526004016106509190614559565b5050900390565b6131626136fb565b61316a6136fb565b604051806080016040528088600781111561318157fe5b8152602001876001600160a01b0316815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200184905291505095945050505050565b815160208301206000906131f381612aa3565b156131fe6005612c45565b9061321c5760405162461bcd60e51b81526004016106509190614559565b50600061322984836111fd565b90506001600160a01b03811615156132416010612c45565b9061325f5760405162461bcd60e51b81526004016106509190614559565b506000918252600160208190526040909220805460ff19169092179091559392505050565b6005546001600160a01b0316158015906132c457506001600160a01b0380831660009081526003602090815260408083209385168352929052205460ff16155b156133c257600554604051637badcc6760e11b8152613379916001600160a01b03169063f75b98ce906132fd90859087906004016142f6565b60206040518083038186803b15801561331557600080fd5b505afa158015613329573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061334d91906140b9565b6001600160a01b0380851660009081526002602090815260408083209387168352929052205490612e3e565b6001600160a01b038084166000818152600260209081526040808320948716808452948252808320959095559181526003825283812092815291905220805460ff191660011790555b5050565b6000806133d286611ada565b90506001600160a01b0386166133ef576133ec8134613104565b90505b844710156133fd6007612c45565b9061341b5760405162461bcd60e51b81526004016106509190614559565b5060006060846001600160a01b031687876040516134399190614264565b60006040518083038185875af1925050503d8060008114613476576040519150601f19603f3d011682016040523d82523d6000602084013e61347b565b606091505b50915091508161348b6004612c45565b906134a95760405162461bcd60e51b81526004016106509190614559565b50600080828060200190518101906134c19190613a00565b91509150896001600160a01b0316826001600160a01b031614801561295d57508061295b8661158d8d611ada565b606061351584848460405180606001604052806029815260200161483560299139613584565b949350505050565b6060612a8283836040518060400160405280601e81526020017f416464726573733a206c6f772d6c6576656c2063616c6c206661696c65640000815250613645565b7f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd25490565b6060824710156135a65760405162461bcd60e51b815260040161065090614600565b6135af85613654565b6135cb5760405162461bcd60e51b815260040161065090614646565b60006060866001600160a01b031685876040516135e89190614264565b60006040518083038185875af1925050503d8060008114613625576040519150601f19603f3d011682016040523d82523d6000602084013e61362a565b606091505b509150915061363a82828661365a565b979650505050505050565b60606135158484600085613584565b3b151590565b60608315613669575081612a82565b8251156136795782518084602001fd5b8160405162461bcd60e51b81526004016106509190614559565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b604051806060016040528060006001600160a01b031681526020016060815260200160006001600160a01b031681525090565b6040805160808101909152806000815260200160006001600160a01b0316815260200160608152602001600081525090565b80356112868161480e565b60008083601f840112613749578182fd5b5081356001600160401b0381111561375f578182fd5b602083019150836020808302850101111561377957600080fd5b9250929050565b600082601f830112613790578081fd5b81356137a361379e8261479b565b614775565b8181529150602080830190848101818402860182018710156137c457600080fd5b60005b848110156137ec5781356137da81614826565b845292820192908201906001016137c7565b505050505092915050565b600082601f830112613807578081fd5b813561381561379e8261479b565b81815291506020808301908481018184028601820187101561383657600080fd5b60005b848110156137ec57813584529282019290820190600101613839565b600082601f830112613865578081fd5b813561387361379e8261479b565b81815291506020808301908481018184028601820187101561389457600080fd5b6000805b858110156138c257823560ff811681146138b0578283fd5b85529383019391830191600101613898565b50505050505092915050565b60008083601f8401126138df578182fd5b5081356001600160401b038111156138f5578182fd5b60208301915083602082850101111561377957600080fd5b600082601f83011261391d578081fd5b81356001600160401b03811115613932578182fd5b613945601f8201601f1916602001614775565b915080825283602082850101111561395c57600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215613986578081fd5b8135612a828161480e565b60008060008060008060c087890312156139a9578182fd5b86356139b48161480e565b955060208701356139c48161480e565b9450604087013593506060870135925060808701356139e28161480e565b915060a08701356139f28161480e565b809150509295509295509295565b60008060408385031215613a12578182fd5b8251613a1d8161480e565b6020939093015192949293505050565b60008060408385031215613a3f578182fd5b8235613a4a8161480e565b91506020830135613a5a8161480e565b809150509250929050565b600080600060608486031215613a79578283fd5b8335613a848161480e565b92506020840135613a948161480e565b929592945050506040919091013590565b600080600080600060a08688031215613abc578283fd5b8535613ac78161480e565b94506020860135613ad78161480e565b93506040860135925060608601356001600160401b03811115613af8578182fd5b613b048882890161390d565b9250506080860135613b158161480e565b809150509295509295909350565b60008060008060008060008060008060e08b8d031215613b41578788fd5b8a35613b4c8161480e565b995060208b0135985060408b0135613b638161480e565b9750613b728c60608d0161372d565b965060808b01356001600160401b0380821115613b8d578586fd5b613b998e838f016138ce565b909850965060a08d0135915080821115613bb1578586fd5b613bbd8e838f016138ce565b909650945060c08d0135915080821115613bd5578384fd5b50613be28d828e016138ce565b915080935050809150509295989b9194979a5092959850565b60008060008060608587031215613c10578182fd5b8435613c1b8161480e565b93506020850135925060408501356001600160401b03811115613c3c578283fd5b613c48878288016138ce565b95989497509550505050565b60008060008060408587031215613c69578182fd5b84356001600160401b0380821115613c7f578384fd5b613c8b88838901613738565b90965094506020870135915080821115613ca3578384fd5b50613c4887828801613738565b600060208284031215613cc1578081fd5b8151612a8281614826565b600060208284031215613cdd578081fd5b5035919050565b60008060208385031215613cf6578182fd5b82356001600160401b03811115613d0b578283fd5b613d17858286016138ce565b90969095509350505050565b60008060008060008060008060008060006101408c8e031215613d44578485fd5b6001600160401b03808d351115613d59578586fd5b613d668e8e358f016138ce565b909c509a5060208d0135995060408d0135811015613d82578586fd5b613d928e60408f01358f016137f7565b98508060608e01351115613da4578586fd5b613db48e60608f01358f01613780565b975060808d0135965060a08d013595508060c08e01351115613dd4578182fd5b613de48e60c08f01358f016137f7565b94508060e08e01351115613df6578182fd5b613e068e60e08f01358f01613855565b9350806101008e01351115613e19578182fd5b613e2a8e6101008f01358f016137f7565b9250806101208e01351115613e3d578182fd5b50613e4f8d6101208e01358e016137f7565b90509295989b509295989b9093969950565b600060208284031215613e72578081fd5b81356001600160401b03811115613e87578182fd5b6135158482850161390d565b60008060408385031215613ea5578182fd5b82356001600160401b03811115613eba578283fd5b613ec68582860161390d565b95602094909401359450505050565b6000806000806000806000806000806101408b8d031215613ef4578384fd5b8a356001600160401b0380821115613f0a578586fd5b613f168e838f0161390d565b9b5060208d01359a5060408d0135915080821115613f32578586fd5b613f3e8e838f016137f7565b995060608d0135915080821115613f53578586fd5b613f5f8e838f01613780565b985060808d0135975060a08d0135965060c08d0135915080821115613f82578586fd5b613f8e8e838f016137f7565b955060e08d0135915080821115613fa3578485fd5b613faf8e838f01613855565b94506101008d0135915080821115613fc5578384fd5b613fd18e838f016137f7565b93506101208d0135915080821115613fe7578283fd5b50613ff48d828e016137f7565b9150509295989b9194979a5092959850565b60008060008060008060008060a0898b031215614021578182fd5b88356001600160401b0380821115614037578384fd5b6140438c838d016138ce565b909a50985060208b013591506140588261480e565b90965060408a0135955060608a01359080821115614074578384fd5b6140808c838d016138ce565b909650945060808b0135915080821115614098578384fd5b506140a58b828c016138ce565b999c989b5096995094979396929594505050565b6000602082840312156140ca578081fd5b5051919050565b6000815180845260208085019450808401835b838110156141025781511515875295820195908201906001016140e4565b509495945050505050565b6000815180845260208085019450808401835b8381101561410257815187529582019590820190600101614120565b6000815180845260208085019450808401835b8381101561410257815160ff168752958201959082019060010161414f565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b600081518084526141b08160208601602086016147e2565b601f01601f19169290920160200192915050565b60ff815116825260ff6020820151166020830152604081015160018060a01b03808216604085015280606084015116606085015250506080810151608083015260a081015160a08301525050565b600081516008811061422057fe5b83526020828101516001600160a01b03169084015260408083015160809185018290529061425090850182614198565b606093840151949093019390935250919050565b600082516142768184602087016147e2565b9190910192915050565b600083516142928184602088016147e2565b9190910191825250602001919050565b90565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b600060018060a01b038088168352808716602084015285604084015260a0606084015261434060a0840186614198565b91508084166080840152509695505050505050565b6001600160a01b038516815260806020820181905260009061437990830186614198565b6040830194909452506060015292915050565b6001600160a01b03851681526060602082018190526000906143b1908301858761416e565b905082604083015295945050505050565b6001600160a01b03861681526080602082018190526000906143e7908301868861416e565b604083019490945250606001529392505050565b6040808252810184905260008560608301825b8781101561443e57602083356144238161480e565b6001600160a01b03168352928301929091019060010161440e565b5083810360208501528481526001600160fb1b0385111561445d578283fd5b602085029150818660208301370160200190815295945050505050565b901515815260200190565b60006101608d151583528c60208401528b60408401528060608401526144ad8184018c61410d565b905082810360808401526144c1818b6140d1565b90508860a08401528760c084015282810360e08401526144e1818861410d565b90508281036101008401526144f6818761413c565b905082810361012084015261450b818661410d565b9050828103610140840152614520818561410d565b9e9d5050505050505050505050505050565b90815260200190565b93845260ff9290921660208401526040830152606082015260800190565b600060208252612a826020830184614198565b6020808252603a908201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260408201527f6563697069656e74206d61792068617665207265766572746564000000000000606082015260800190565b6020808252601d908201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b60c0810161128682846141c4565b600061010061469a83876141c4565b8060c084015260018060a01b0380865116828501526020860151915060606101208501526146cc610160850183614198565b9150806040870151166101408501525082810360e08401526146ee8185614198565b9695505050505050565b60006080825261470b6080830188614212565b6001600160a01b03878116602085015286166040840152828103606084015261473581858761416e565b98975050505050505050565b6000604082526147546040830186614212565b82810360208401526146ee81858761416e565b60ff91909116815260200190565b6040518181016001600160401b038111828210171561479357600080fd5b604052919050565b60006001600160401b038211156147b0578081fd5b5060209081020190565b600080858511156147c9578182fd5b838611156147d5578182fd5b5050820193919092039150565b60005b838110156147fd5781810151838201526020016147e5565b838111156107075750506000910152565b6001600160a01b038116811461482357600080fd5b50565b801515811461482357600080fdfe416464726573733a206c6f772d6c6576656c2063616c6c20776974682076616c7565206661696c6564a264697066735822122047898b3a3e3a0e29f4a0810e8da4751f4769681830f0791179ee69d6cdc6981764736f6c634300060c0033"

// // DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
// func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vault, error) {
// 	parsed, err := abi.JSON(strings.NewReader(VaultABI))
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}

// 	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend)
// 	if err != nil {
// 		return common.Address{}, nil, nil, err
// 	}
// 	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
// }

// // Vault is an auto generated Go binding around an Ethereum contract.
// type Vault struct {
// 	VaultCaller     // Read-only binding to the contract
// 	VaultTransactor // Write-only binding to the contract
// 	VaultFilterer   // Log filterer for contract events
// }

// // VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
// type VaultCaller struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
// type VaultTransactor struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
// type VaultFilterer struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // VaultSession is an auto generated Go binding around an Ethereum contract,
// // with pre-set call and transact options.
// type VaultSession struct {
// 	Contract     *Vault            // Generic contract binding to set the session for
// 	CallOpts     bind.CallOpts     // Call options to use throughout this session
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// // with pre-set call options.
// type VaultCallerSession struct {
// 	Contract *VaultCaller  // Generic contract caller binding to set the session for
// 	CallOpts bind.CallOpts // Call options to use throughout this session
// }

// // VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// // with pre-set transact options.
// type VaultTransactorSession struct {
// 	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
// type VaultRaw struct {
// 	Contract *Vault // Generic contract binding to access the raw methods on
// }

// // VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
// type VaultCallerRaw struct {
// 	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
// }

// // VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
// type VaultTransactorRaw struct {
// 	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
// }

// // NewVault creates a new instance of Vault, bound to a specific deployed contract.
// func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
// 	contract, err := bindVault(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
// }

// // NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
// func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
// 	contract, err := bindVault(address, caller, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultCaller{contract: contract}, nil
// }

// // NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
// func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
// 	contract, err := bindVault(address, nil, transactor, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultTransactor{contract: contract}, nil
// }

// // NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
// func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
// 	contract, err := bindVault(address, nil, nil, filterer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultFilterer{contract: contract}, nil
// }

// // bindVault binds a generic wrapper to an already deployed contract.
// func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
// 	parsed, err := abi.JSON(strings.NewReader(VaultABI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Vault.Contract.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Vault.Contract.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Vault.Contract.contract.Transact(opts, method, params...)
// }

// // BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
// //
// // Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultCaller) BURNCALLREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "BURN_CALL_REQUEST_METADATA_TYPE")

// 	if err != nil {
// 		return *new(uint8), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

// 	return out0, err

// }

// // BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
// //
// // Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultSession) BURNCALLREQUESTMETADATATYPE() (uint8, error) {
// 	return _Vault.Contract.BURNCALLREQUESTMETADATATYPE(&_Vault.CallOpts)
// }

// // BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
// //
// // Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultCallerSession) BURNCALLREQUESTMETADATATYPE() (uint8, error) {
// 	return _Vault.Contract.BURNCALLREQUESTMETADATATYPE(&_Vault.CallOpts)
// }

// // BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
// //
// // Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultCaller) BURNREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "BURN_REQUEST_METADATA_TYPE")

// 	if err != nil {
// 		return *new(uint8), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

// 	return out0, err

// }

// // BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
// //
// // Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultSession) BURNREQUESTMETADATATYPE() (uint8, error) {
// 	return _Vault.Contract.BURNREQUESTMETADATATYPE(&_Vault.CallOpts)
// }

// // BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
// //
// // Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultCallerSession) BURNREQUESTMETADATATYPE() (uint8, error) {
// 	return _Vault.Contract.BURNREQUESTMETADATATYPE(&_Vault.CallOpts)
// }

// // BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
// //
// // Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultCaller) BURNTOCONTRACTREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "BURN_TO_CONTRACT_REQUEST_METADATA_TYPE")

// 	if err != nil {
// 		return *new(uint8), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

// 	return out0, err

// }

// // BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
// //
// // Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultSession) BURNTOCONTRACTREQUESTMETADATATYPE() (uint8, error) {
// 	return _Vault.Contract.BURNTOCONTRACTREQUESTMETADATATYPE(&_Vault.CallOpts)
// }

// // BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
// //
// // Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
// func (_Vault *VaultCallerSession) BURNTOCONTRACTREQUESTMETADATATYPE() (uint8, error) {
// 	return _Vault.Contract.BURNTOCONTRACTREQUESTMETADATATYPE(&_Vault.CallOpts)
// }

// // CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
// //
// // Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
// func (_Vault *VaultCaller) CURRENTNETWORKID(opts *bind.CallOpts) (uint8, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "CURRENT_NETWORK_ID")

// 	if err != nil {
// 		return *new(uint8), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

// 	return out0, err

// }

// // CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
// //
// // Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
// func (_Vault *VaultSession) CURRENTNETWORKID() (uint8, error) {
// 	return _Vault.Contract.CURRENTNETWORKID(&_Vault.CallOpts)
// }

// // CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
// //
// // Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
// func (_Vault *VaultCallerSession) CURRENTNETWORKID() (uint8, error) {
// 	return _Vault.Contract.CURRENTNETWORKID(&_Vault.CallOpts)
// }

// // ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
// //
// // Solidity: function ETH_TOKEN() view returns(address)
// func (_Vault *VaultCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "ETH_TOKEN")

// 	if err != nil {
// 		return *new(common.Address), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

// 	return out0, err

// }

// // ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
// //
// // Solidity: function ETH_TOKEN() view returns(address)
// func (_Vault *VaultSession) ETHTOKEN() (common.Address, error) {
// 	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
// }

// // ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
// //
// // Solidity: function ETH_TOKEN() view returns(address)
// func (_Vault *VaultCallerSession) ETHTOKEN() (common.Address, error) {
// 	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
// }

// // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// //
// // Solidity: function balanceOf(address token) view returns(uint256)
// func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "balanceOf", token)

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// //
// // Solidity: function balanceOf(address token) view returns(uint256)
// func (_Vault *VaultSession) BalanceOf(token common.Address) (*big.Int, error) {
// 	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
// }

// // BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
// //
// // Solidity: function balanceOf(address token) view returns(uint256)
// func (_Vault *VaultCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
// 	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
// }

// // GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
// //
// // Solidity: function getDecimals(address token) view returns(uint8)
// func (_Vault *VaultCaller) GetDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "getDecimals", token)

// 	if err != nil {
// 		return *new(uint8), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

// 	return out0, err

// }

// // GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
// //
// // Solidity: function getDecimals(address token) view returns(uint8)
// func (_Vault *VaultSession) GetDecimals(token common.Address) (uint8, error) {
// 	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
// }

// // GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
// //
// // Solidity: function getDecimals(address token) view returns(uint8)
// func (_Vault *VaultCallerSession) GetDecimals(token common.Address) (uint8, error) {
// 	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
// }

// // GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
// //
// // Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
// func (_Vault *VaultCaller) GetDepositedBalance(opts *bind.CallOpts, token common.Address, owner common.Address) (*big.Int, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "getDepositedBalance", token, owner)

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
// //
// // Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
// func (_Vault *VaultSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
// 	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
// }

// // GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
// //
// // Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
// func (_Vault *VaultCallerSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
// 	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
// }

// // IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
// //
// // Solidity: function isInitialized() view returns(bool)
// func (_Vault *VaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "isInitialized")

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
// //
// // Solidity: function isInitialized() view returns(bool)
// func (_Vault *VaultSession) IsInitialized() (bool, error) {
// 	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
// }

// // IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
// //
// // Solidity: function isInitialized() view returns(bool)
// func (_Vault *VaultCallerSession) IsInitialized() (bool, error) {
// 	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
// }

// // IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
// //
// // Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
// func (_Vault *VaultCaller) IsSigDataUsed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "isSigDataUsed", hash)

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
// //
// // Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
// func (_Vault *VaultSession) IsSigDataUsed(hash [32]byte) (bool, error) {
// 	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
// }

// // IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
// //
// // Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
// func (_Vault *VaultCallerSession) IsSigDataUsed(hash [32]byte) (bool, error) {
// 	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
// }

// // IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
// //
// // Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
// func (_Vault *VaultCaller) IsWithdrawed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "isWithdrawed", hash)

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
// //
// // Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
// func (_Vault *VaultSession) IsWithdrawed(hash [32]byte) (bool, error) {
// 	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
// }

// // IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
// //
// // Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
// func (_Vault *VaultCallerSession) IsWithdrawed(hash [32]byte) (bool, error) {
// 	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
// }

// // Migration is a free data retrieval call binding the contract method 0x995fac11.
// //
// // Solidity: function migration(address , address ) view returns(bool)
// func (_Vault *VaultCaller) Migration(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "migration", arg0, arg1)

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // Migration is a free data retrieval call binding the contract method 0x995fac11.
// //
// // Solidity: function migration(address , address ) view returns(bool)
// func (_Vault *VaultSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
// 	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
// }

// // Migration is a free data retrieval call binding the contract method 0x995fac11.
// //
// // Solidity: function migration(address , address ) view returns(bool)
// func (_Vault *VaultCallerSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
// 	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
// }

// // NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
// //
// // Solidity: function notEntered() view returns(bool)
// func (_Vault *VaultCaller) NotEntered(opts *bind.CallOpts) (bool, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "notEntered")

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
// //
// // Solidity: function notEntered() view returns(bool)
// func (_Vault *VaultSession) NotEntered() (bool, error) {
// 	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
// }

// // NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
// //
// // Solidity: function notEntered() view returns(bool)
// func (_Vault *VaultCallerSession) NotEntered() (bool, error) {
// 	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
// }

// // ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
// //
// // Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
// func (_Vault *VaultCaller) ParseBurnInst(opts *bind.CallOpts, inst []byte) (VaultBurnInstData, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "parseBurnInst", inst)

// 	if err != nil {
// 		return *new(VaultBurnInstData), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(VaultBurnInstData)).(*VaultBurnInstData)

// 	return out0, err

// }

// // ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
// //
// // Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
// func (_Vault *VaultSession) ParseBurnInst(inst []byte) (VaultBurnInstData, error) {
// 	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
// }

// // ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
// //
// // Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
// func (_Vault *VaultCallerSession) ParseBurnInst(inst []byte) (VaultBurnInstData, error) {
// 	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
// }

// // ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
// //
// // Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
// func (_Vault *VaultCaller) ParseCalldataFromBurnInst(opts *bind.CallOpts, inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "parseCalldataFromBurnInst", inst)

// 	if err != nil {
// 		return *new(VaultBurnInstData), *new(VaultRedepositOptions), *new([]byte), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(VaultBurnInstData)).(*VaultBurnInstData)
// 	out1 := *abi.ConvertType(out[1], new(VaultRedepositOptions)).(*VaultRedepositOptions)
// 	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

// 	return out0, out1, out2, err

// }

// // ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
// //
// // Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
// func (_Vault *VaultSession) ParseCalldataFromBurnInst(inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
// 	return _Vault.Contract.ParseCalldataFromBurnInst(&_Vault.CallOpts, inst)
// }

// // ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
// //
// // Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
// func (_Vault *VaultCallerSession) ParseCalldataFromBurnInst(inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
// 	return _Vault.Contract.ParseCalldataFromBurnInst(&_Vault.CallOpts, inst)
// }

// // PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
// //
// // Solidity: function prevVault() view returns(address)
// func (_Vault *VaultCaller) PrevVault(opts *bind.CallOpts) (common.Address, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "prevVault")

// 	if err != nil {
// 		return *new(common.Address), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

// 	return out0, err

// }

// // PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
// //
// // Solidity: function prevVault() view returns(address)
// func (_Vault *VaultSession) PrevVault() (common.Address, error) {
// 	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
// }

// // PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
// //
// // Solidity: function prevVault() view returns(address)
// func (_Vault *VaultCallerSession) PrevVault() (common.Address, error) {
// 	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
// }

// // SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
// //
// // Solidity: function sigDataUsed(bytes32 ) view returns(bool)
// func (_Vault *VaultCaller) SigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "sigDataUsed", arg0)

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
// //
// // Solidity: function sigDataUsed(bytes32 ) view returns(bool)
// func (_Vault *VaultSession) SigDataUsed(arg0 [32]byte) (bool, error) {
// 	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
// }

// // SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
// //
// // Solidity: function sigDataUsed(bytes32 ) view returns(bool)
// func (_Vault *VaultCallerSession) SigDataUsed(arg0 [32]byte) (bool, error) {
// 	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
// }

// // SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
// //
// // Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
// func (_Vault *VaultCaller) SigToAddress(opts *bind.CallOpts, signData []byte, hash [32]byte) (common.Address, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "sigToAddress", signData, hash)

// 	if err != nil {
// 		return *new(common.Address), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

// 	return out0, err

// }

// // SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
// //
// // Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
// func (_Vault *VaultSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
// 	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
// }

// // SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
// //
// // Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
// func (_Vault *VaultCallerSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
// 	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
// }

// // TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
// //
// // Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
// func (_Vault *VaultCaller) TotalDepositedToSCAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "totalDepositedToSCAmount", arg0)

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
// //
// // Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
// func (_Vault *VaultSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
// 	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
// }

// // TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
// //
// // Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
// func (_Vault *VaultCallerSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
// 	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
// }

// // WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
// //
// // Solidity: function withdrawRequests(address , address ) view returns(uint256)
// func (_Vault *VaultCaller) WithdrawRequests(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "withdrawRequests", arg0, arg1)

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
// //
// // Solidity: function withdrawRequests(address , address ) view returns(uint256)
// func (_Vault *VaultSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
// 	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
// }

// // WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
// //
// // Solidity: function withdrawRequests(address , address ) view returns(uint256)
// func (_Vault *VaultCallerSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
// 	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
// }

// // Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
// //
// // Solidity: function withdrawed(bytes32 ) view returns(bool)
// func (_Vault *VaultCaller) Withdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
// 	var out []interface{}
// 	err := _Vault.contract.Call(opts, &out, "withdrawed", arg0)

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
// //
// // Solidity: function withdrawed(bytes32 ) view returns(bool)
// func (_Vault *VaultSession) Withdrawed(arg0 [32]byte) (bool, error) {
// 	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
// }

// // Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
// //
// // Solidity: function withdrawed(bytes32 ) view returns(bool)
// func (_Vault *VaultCallerSession) Withdrawed(arg0 [32]byte) (bool, error) {
// 	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
// }

// // CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
// //
// // Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
// func (_Vault *VaultTransactor) CallExternal(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "_callExternal", token, to, amount, externalCalldata, redepositToken)
// }

// // CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
// //
// // Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
// func (_Vault *VaultSession) CallExternal(token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
// 	return _Vault.Contract.CallExternal(&_Vault.TransactOpts, token, to, amount, externalCalldata, redepositToken)
// }

// // CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
// //
// // Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
// func (_Vault *VaultTransactorSession) CallExternal(token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
// 	return _Vault.Contract.CallExternal(&_Vault.TransactOpts, token, to, amount, externalCalldata, redepositToken)
// }

// // TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
// //
// // Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
// func (_Vault *VaultTransactor) TransferExternal(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "_transferExternal", token, to, amount)
// }

// // TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
// //
// // Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
// func (_Vault *VaultSession) TransferExternal(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _Vault.Contract.TransferExternal(&_Vault.TransactOpts, token, to, amount)
// }

// // TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
// //
// // Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
// func (_Vault *VaultTransactorSession) TransferExternal(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
// 	return _Vault.Contract.TransferExternal(&_Vault.TransactOpts, token, to, amount)
// }

// // Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
// //
// // Solidity: function deposit(string incognitoAddress) payable returns()
// func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "deposit", incognitoAddress)
// }

// // Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
// //
// // Solidity: function deposit(string incognitoAddress) payable returns()
// func (_Vault *VaultSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress)
// }

// // Deposit is a paid mutator transaction binding the contract method 0xa26e1186.
// //
// // Solidity: function deposit(string incognitoAddress) payable returns()
// func (_Vault *VaultTransactorSession) Deposit(incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress)
// }

// // DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
// //
// // Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
// func (_Vault *VaultTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "depositERC20", token, amount, incognitoAddress)
// }

// // DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
// //
// // Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
// func (_Vault *VaultSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress)
// }

// // DepositERC20 is a paid mutator transaction binding the contract method 0x5a67cb87.
// //
// // Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress) returns()
// func (_Vault *VaultTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress)
// }

// // DepositERC20V2 is a paid mutator transaction binding the contract method 0x5c421b2f.
// //
// // Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress) returns()
// func (_Vault *VaultTransactor) DepositERC20V2(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "depositERC20_V2", token, amount, incognitoAddress)
// }

// // DepositERC20V2 is a paid mutator transaction binding the contract method 0x5c421b2f.
// //
// // Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress) returns()
// func (_Vault *VaultSession) DepositERC20V2(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.Contract.DepositERC20V2(&_Vault.TransactOpts, token, amount, incognitoAddress)
// }

// // DepositERC20V2 is a paid mutator transaction binding the contract method 0x5c421b2f.
// //
// // Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress) returns()
// func (_Vault *VaultTransactorSession) DepositERC20V2(token common.Address, amount *big.Int, incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.Contract.DepositERC20V2(&_Vault.TransactOpts, token, amount, incognitoAddress)
// }

// // DepositV2 is a paid mutator transaction binding the contract method 0x49c5ba23.
// //
// // Solidity: function deposit_V2(string incognitoAddress) payable returns()
// func (_Vault *VaultTransactor) DepositV2(opts *bind.TransactOpts, incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "deposit_V2", incognitoAddress)
// }

// // DepositV2 is a paid mutator transaction binding the contract method 0x49c5ba23.
// //
// // Solidity: function deposit_V2(string incognitoAddress) payable returns()
// func (_Vault *VaultSession) DepositV2(incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.Contract.DepositV2(&_Vault.TransactOpts, incognitoAddress)
// }

// // DepositV2 is a paid mutator transaction binding the contract method 0x49c5ba23.
// //
// // Solidity: function deposit_V2(string incognitoAddress) payable returns()
// func (_Vault *VaultTransactorSession) DepositV2(incognitoAddress string) (*types.Transaction, error) {
// 	return _Vault.Contract.DepositV2(&_Vault.TransactOpts, incognitoAddress)
// }

// // Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
// //
// // Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
// func (_Vault *VaultTransactor) Execute(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "execute", token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
// }

// // Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
// //
// // Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
// func (_Vault *VaultSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
// 	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
// }

// // Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
// //
// // Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
// func (_Vault *VaultTransactorSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
// 	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
// }

// // ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
// //
// // Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultTransactor) ExecuteWithBurnProof(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "executeWithBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
// //
// // Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultSession) ExecuteWithBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.Contract.ExecuteWithBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
// //
// // Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultTransactorSession) ExecuteWithBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.Contract.ExecuteWithBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
// //
// // Solidity: function initialize(address _prevVault) returns()
// func (_Vault *VaultTransactor) Initialize(opts *bind.TransactOpts, _prevVault common.Address) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "initialize", _prevVault)
// }

// // Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
// //
// // Solidity: function initialize(address _prevVault) returns()
// func (_Vault *VaultSession) Initialize(_prevVault common.Address) (*types.Transaction, error) {
// 	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _prevVault)
// }

// // Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
// //
// // Solidity: function initialize(address _prevVault) returns()
// func (_Vault *VaultTransactorSession) Initialize(_prevVault common.Address) (*types.Transaction, error) {
// 	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _prevVault)
// }

// // RequestWithdraw is a paid mutator transaction binding the contract method 0x87add440.
// //
// // Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp) returns()
// func (_Vault *VaultTransactor) RequestWithdraw(opts *bind.TransactOpts, incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "requestWithdraw", incognitoAddress, token, amount, signData, timestamp)
// }

// // RequestWithdraw is a paid mutator transaction binding the contract method 0x87add440.
// //
// // Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp) returns()
// func (_Vault *VaultSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte) (*types.Transaction, error) {
// 	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp)
// }

// // RequestWithdraw is a paid mutator transaction binding the contract method 0x87add440.
// //
// // Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp) returns()
// func (_Vault *VaultTransactorSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte) (*types.Transaction, error) {
// 	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp)
// }

// // SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
// //
// // Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultTransactor) SubmitBurnProof(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "submitBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
// //
// // Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultSession) SubmitBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
// //
// // Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultTransactorSession) SubmitBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
// //
// // Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
// func (_Vault *VaultTransactor) UpdateAssets(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "updateAssets", assets, amounts)
// }

// // UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
// //
// // Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
// func (_Vault *VaultSession) UpdateAssets(assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
// 	return _Vault.Contract.UpdateAssets(&_Vault.TransactOpts, assets, amounts)
// }

// // UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
// //
// // Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
// func (_Vault *VaultTransactorSession) UpdateAssets(assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
// 	return _Vault.Contract.UpdateAssets(&_Vault.TransactOpts, assets, amounts)
// }

// // Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
// //
// // Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.contract.Transact(opts, "withdraw", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
// //
// // Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultSession) Withdraw(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
// //
// // Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
// func (_Vault *VaultTransactorSession) Withdraw(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
// 	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
// }

// // Receive is a paid mutator transaction binding the contract receive function.
// //
// // Solidity: receive() payable returns()
// func (_Vault *VaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Vault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
// }

// // Receive is a paid mutator transaction binding the contract receive function.
// //
// // Solidity: receive() payable returns()
// func (_Vault *VaultSession) Receive() (*types.Transaction, error) {
// 	return _Vault.Contract.Receive(&_Vault.TransactOpts)
// }

// // Receive is a paid mutator transaction binding the contract receive function.
// //
// // Solidity: receive() payable returns()
// func (_Vault *VaultTransactorSession) Receive() (*types.Transaction, error) {
// 	return _Vault.Contract.Receive(&_Vault.TransactOpts)
// }

// // VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Vault contract.
// type VaultDepositIterator struct {
// 	Event *VaultDeposit // Event containing the contract specifics and raw log

// 	contract *bind.BoundContract // Generic contract to use for unpacking event data
// 	event    string              // Event name to use for unpacking event data

// 	logs chan types.Log        // Log channel receiving the found contract events
// 	sub  ethereum.Subscription // Subscription for errors, completion and termination
// 	done bool                  // Whether the subscription completed delivering logs
// 	fail error                 // Occurred error to stop iteration
// }

// // Next advances the iterator to the subsequent event, returning whether there
// // are any more events found. In case of a retrieval or parsing error, false is
// // returned and Error() can be queried for the exact failure.
// func (it *VaultDepositIterator) Next() bool {
// 	// If the iterator failed, stop iterating
// 	if it.fail != nil {
// 		return false
// 	}
// 	// If the iterator completed, deliver directly whatever's available
// 	if it.done {
// 		select {
// 		case log := <-it.logs:
// 			it.Event = new(VaultDeposit)
// 			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 				it.fail = err
// 				return false
// 			}
// 			it.Event.Raw = log
// 			return true

// 		default:
// 			return false
// 		}
// 	}
// 	// Iterator still in progress, wait for either a data or an error event
// 	select {
// 	case log := <-it.logs:
// 		it.Event = new(VaultDeposit)
// 		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 			it.fail = err
// 			return false
// 		}
// 		it.Event.Raw = log
// 		return true

// 	case err := <-it.sub.Err():
// 		it.done = true
// 		it.fail = err
// 		return it.Next()
// 	}
// }

// // Error returns any retrieval or parsing error occurred during filtering.
// func (it *VaultDepositIterator) Error() error {
// 	return it.fail
// }

// // Close terminates the iteration process, releasing any pending underlying
// // resources.
// func (it *VaultDepositIterator) Close() error {
// 	it.sub.Unsubscribe()
// 	return nil
// }

// // VaultDeposit represents a Deposit event raised by the Vault contract.
// type VaultDeposit struct {
// 	Token            common.Address
// 	IncognitoAddress string
// 	Amount           *big.Int
// 	Raw              types.Log // Blockchain specific contextual infos
// }

// // FilterDeposit is a free log retrieval operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
// //
// // Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
// func (_Vault *VaultFilterer) FilterDeposit(opts *bind.FilterOpts) (*VaultDepositIterator, error) {

// 	logs, sub, err := _Vault.contract.FilterLogs(opts, "Deposit")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultDepositIterator{contract: _Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
// }

// // WatchDeposit is a free log subscription operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
// //
// // Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
// func (_Vault *VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VaultDeposit) (event.Subscription, error) {

// 	logs, sub, err := _Vault.contract.WatchLogs(opts, "Deposit")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return event.NewSubscription(func(quit <-chan struct{}) error {
// 		defer sub.Unsubscribe()
// 		for {
// 			select {
// 			case log := <-logs:
// 				// New log arrived, parse the event and forward to the user
// 				event := new(VaultDeposit)
// 				if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
// 					return err
// 				}
// 				event.Raw = log

// 				select {
// 				case sink <- event:
// 				case err := <-sub.Err():
// 					return err
// 				case <-quit:
// 					return nil
// 				}
// 			case err := <-sub.Err():
// 				return err
// 			case <-quit:
// 				return nil
// 			}
// 		}
// 	}), nil
// }

// // ParseDeposit is a log parse operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
// //
// // Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
// func (_Vault *VaultFilterer) ParseDeposit(log types.Log) (*VaultDeposit, error) {
// 	event := new(VaultDeposit)
// 	if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
// 		return nil, err
// 	}
// 	event.Raw = log
// 	return event, nil
// }

// // VaultDepositV2Iterator is returned from FilterDepositV2 and is used to iterate over the raw logs and unpacked data for DepositV2 events raised by the Vault contract.
// type VaultDepositV2Iterator struct {
// 	Event *VaultDepositV2 // Event containing the contract specifics and raw log

// 	contract *bind.BoundContract // Generic contract to use for unpacking event data
// 	event    string              // Event name to use for unpacking event data

// 	logs chan types.Log        // Log channel receiving the found contract events
// 	sub  ethereum.Subscription // Subscription for errors, completion and termination
// 	done bool                  // Whether the subscription completed delivering logs
// 	fail error                 // Occurred error to stop iteration
// }

// // Next advances the iterator to the subsequent event, returning whether there
// // are any more events found. In case of a retrieval or parsing error, false is
// // returned and Error() can be queried for the exact failure.
// func (it *VaultDepositV2Iterator) Next() bool {
// 	// If the iterator failed, stop iterating
// 	if it.fail != nil {
// 		return false
// 	}
// 	// If the iterator completed, deliver directly whatever's available
// 	if it.done {
// 		select {
// 		case log := <-it.logs:
// 			it.Event = new(VaultDepositV2)
// 			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 				it.fail = err
// 				return false
// 			}
// 			it.Event.Raw = log
// 			return true

// 		default:
// 			return false
// 		}
// 	}
// 	// Iterator still in progress, wait for either a data or an error event
// 	select {
// 	case log := <-it.logs:
// 		it.Event = new(VaultDepositV2)
// 		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 			it.fail = err
// 			return false
// 		}
// 		it.Event.Raw = log
// 		return true

// 	case err := <-it.sub.Err():
// 		it.done = true
// 		it.fail = err
// 		return it.Next()
// 	}
// }

// // Error returns any retrieval or parsing error occurred during filtering.
// func (it *VaultDepositV2Iterator) Error() error {
// 	return it.fail
// }

// // Close terminates the iteration process, releasing any pending underlying
// // resources.
// func (it *VaultDepositV2Iterator) Close() error {
// 	it.sub.Unsubscribe()
// 	return nil
// }

// // VaultDepositV2 represents a DepositV2 event raised by the Vault contract.
// type VaultDepositV2 struct {
// 	Token            common.Address
// 	IncognitoAddress string
// 	Amount           *big.Int
// 	DepositID        *big.Int
// 	Raw              types.Log // Blockchain specific contextual infos
// }

// // FilterDepositV2 is a free log retrieval operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
// //
// // Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
// func (_Vault *VaultFilterer) FilterDepositV2(opts *bind.FilterOpts) (*VaultDepositV2Iterator, error) {

// 	logs, sub, err := _Vault.contract.FilterLogs(opts, "DepositV2")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultDepositV2Iterator{contract: _Vault.contract, event: "DepositV2", logs: logs, sub: sub}, nil
// }

// // WatchDepositV2 is a free log subscription operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
// //
// // Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
// func (_Vault *VaultFilterer) WatchDepositV2(opts *bind.WatchOpts, sink chan<- *VaultDepositV2) (event.Subscription, error) {

// 	logs, sub, err := _Vault.contract.WatchLogs(opts, "DepositV2")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return event.NewSubscription(func(quit <-chan struct{}) error {
// 		defer sub.Unsubscribe()
// 		for {
// 			select {
// 			case log := <-logs:
// 				// New log arrived, parse the event and forward to the user
// 				event := new(VaultDepositV2)
// 				if err := _Vault.contract.UnpackLog(event, "DepositV2", log); err != nil {
// 					return err
// 				}
// 				event.Raw = log

// 				select {
// 				case sink <- event:
// 				case err := <-sub.Err():
// 					return err
// 				case <-quit:
// 					return nil
// 				}
// 			case err := <-sub.Err():
// 				return err
// 			case <-quit:
// 				return nil
// 			}
// 		}
// 	}), nil
// }

// // ParseDepositV2 is a log parse operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
// //
// // Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
// func (_Vault *VaultFilterer) ParseDepositV2(log types.Log) (*VaultDepositV2, error) {
// 	event := new(VaultDepositV2)
// 	if err := _Vault.contract.UnpackLog(event, "DepositV2", log); err != nil {
// 		return nil, err
// 	}
// 	event.Raw = log
// 	return event, nil
// }

// // VaultRedepositIterator is returned from FilterRedeposit and is used to iterate over the raw logs and unpacked data for Redeposit events raised by the Vault contract.
// type VaultRedepositIterator struct {
// 	Event *VaultRedeposit // Event containing the contract specifics and raw log

// 	contract *bind.BoundContract // Generic contract to use for unpacking event data
// 	event    string              // Event name to use for unpacking event data

// 	logs chan types.Log        // Log channel receiving the found contract events
// 	sub  ethereum.Subscription // Subscription for errors, completion and termination
// 	done bool                  // Whether the subscription completed delivering logs
// 	fail error                 // Occurred error to stop iteration
// }

// // Next advances the iterator to the subsequent event, returning whether there
// // are any more events found. In case of a retrieval or parsing error, false is
// // returned and Error() can be queried for the exact failure.
// func (it *VaultRedepositIterator) Next() bool {
// 	// If the iterator failed, stop iterating
// 	if it.fail != nil {
// 		return false
// 	}
// 	// If the iterator completed, deliver directly whatever's available
// 	if it.done {
// 		select {
// 		case log := <-it.logs:
// 			it.Event = new(VaultRedeposit)
// 			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 				it.fail = err
// 				return false
// 			}
// 			it.Event.Raw = log
// 			return true

// 		default:
// 			return false
// 		}
// 	}
// 	// Iterator still in progress, wait for either a data or an error event
// 	select {
// 	case log := <-it.logs:
// 		it.Event = new(VaultRedeposit)
// 		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 			it.fail = err
// 			return false
// 		}
// 		it.Event.Raw = log
// 		return true

// 	case err := <-it.sub.Err():
// 		it.done = true
// 		it.fail = err
// 		return it.Next()
// 	}
// }

// // Error returns any retrieval or parsing error occurred during filtering.
// func (it *VaultRedepositIterator) Error() error {
// 	return it.fail
// }

// // Close terminates the iteration process, releasing any pending underlying
// // resources.
// func (it *VaultRedepositIterator) Close() error {
// 	it.sub.Unsubscribe()
// 	return nil
// }

// // VaultRedeposit represents a Redeposit event raised by the Vault contract.
// type VaultRedeposit struct {
// 	Token               common.Address
// 	RedepositIncAddress []byte
// 	Amount              *big.Int
// 	Itx                 [32]byte
// 	Raw                 types.Log // Blockchain specific contextual infos
// }

// // FilterRedeposit is a free log retrieval operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
// //
// // Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
// func (_Vault *VaultFilterer) FilterRedeposit(opts *bind.FilterOpts) (*VaultRedepositIterator, error) {

// 	logs, sub, err := _Vault.contract.FilterLogs(opts, "Redeposit")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultRedepositIterator{contract: _Vault.contract, event: "Redeposit", logs: logs, sub: sub}, nil
// }

// // WatchRedeposit is a free log subscription operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
// //
// // Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
// func (_Vault *VaultFilterer) WatchRedeposit(opts *bind.WatchOpts, sink chan<- *VaultRedeposit) (event.Subscription, error) {

// 	logs, sub, err := _Vault.contract.WatchLogs(opts, "Redeposit")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return event.NewSubscription(func(quit <-chan struct{}) error {
// 		defer sub.Unsubscribe()
// 		for {
// 			select {
// 			case log := <-logs:
// 				// New log arrived, parse the event and forward to the user
// 				event := new(VaultRedeposit)
// 				if err := _Vault.contract.UnpackLog(event, "Redeposit", log); err != nil {
// 					return err
// 				}
// 				event.Raw = log

// 				select {
// 				case sink <- event:
// 				case err := <-sub.Err():
// 					return err
// 				case <-quit:
// 					return nil
// 				}
// 			case err := <-sub.Err():
// 				return err
// 			case <-quit:
// 				return nil
// 			}
// 		}
// 	}), nil
// }

// // ParseRedeposit is a log parse operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
// //
// // Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
// func (_Vault *VaultFilterer) ParseRedeposit(log types.Log) (*VaultRedeposit, error) {
// 	event := new(VaultRedeposit)
// 	if err := _Vault.contract.UnpackLog(event, "Redeposit", log); err != nil {
// 		return nil, err
// 	}
// 	event.Raw = log
// 	return event, nil
// }

// // VaultUpdateIncognitoProxyIterator is returned from FilterUpdateIncognitoProxy and is used to iterate over the raw logs and unpacked data for UpdateIncognitoProxy events raised by the Vault contract.
// type VaultUpdateIncognitoProxyIterator struct {
// 	Event *VaultUpdateIncognitoProxy // Event containing the contract specifics and raw log

// 	contract *bind.BoundContract // Generic contract to use for unpacking event data
// 	event    string              // Event name to use for unpacking event data

// 	logs chan types.Log        // Log channel receiving the found contract events
// 	sub  ethereum.Subscription // Subscription for errors, completion and termination
// 	done bool                  // Whether the subscription completed delivering logs
// 	fail error                 // Occurred error to stop iteration
// }

// // Next advances the iterator to the subsequent event, returning whether there
// // are any more events found. In case of a retrieval or parsing error, false is
// // returned and Error() can be queried for the exact failure.
// func (it *VaultUpdateIncognitoProxyIterator) Next() bool {
// 	// If the iterator failed, stop iterating
// 	if it.fail != nil {
// 		return false
// 	}
// 	// If the iterator completed, deliver directly whatever's available
// 	if it.done {
// 		select {
// 		case log := <-it.logs:
// 			it.Event = new(VaultUpdateIncognitoProxy)
// 			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 				it.fail = err
// 				return false
// 			}
// 			it.Event.Raw = log
// 			return true

// 		default:
// 			return false
// 		}
// 	}
// 	// Iterator still in progress, wait for either a data or an error event
// 	select {
// 	case log := <-it.logs:
// 		it.Event = new(VaultUpdateIncognitoProxy)
// 		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 			it.fail = err
// 			return false
// 		}
// 		it.Event.Raw = log
// 		return true

// 	case err := <-it.sub.Err():
// 		it.done = true
// 		it.fail = err
// 		return it.Next()
// 	}
// }

// // Error returns any retrieval or parsing error occurred during filtering.
// func (it *VaultUpdateIncognitoProxyIterator) Error() error {
// 	return it.fail
// }

// // Close terminates the iteration process, releasing any pending underlying
// // resources.
// func (it *VaultUpdateIncognitoProxyIterator) Close() error {
// 	it.sub.Unsubscribe()
// 	return nil
// }

// // VaultUpdateIncognitoProxy represents a UpdateIncognitoProxy event raised by the Vault contract.
// type VaultUpdateIncognitoProxy struct {
// 	NewIncognitoProxy common.Address
// 	Raw               types.Log // Blockchain specific contextual infos
// }

// // FilterUpdateIncognitoProxy is a free log retrieval operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
// //
// // Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
// func (_Vault *VaultFilterer) FilterUpdateIncognitoProxy(opts *bind.FilterOpts) (*VaultUpdateIncognitoProxyIterator, error) {

// 	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateIncognitoProxy")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultUpdateIncognitoProxyIterator{contract: _Vault.contract, event: "UpdateIncognitoProxy", logs: logs, sub: sub}, nil
// }

// // WatchUpdateIncognitoProxy is a free log subscription operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
// //
// // Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
// func (_Vault *VaultFilterer) WatchUpdateIncognitoProxy(opts *bind.WatchOpts, sink chan<- *VaultUpdateIncognitoProxy) (event.Subscription, error) {

// 	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateIncognitoProxy")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return event.NewSubscription(func(quit <-chan struct{}) error {
// 		defer sub.Unsubscribe()
// 		for {
// 			select {
// 			case log := <-logs:
// 				// New log arrived, parse the event and forward to the user
// 				event := new(VaultUpdateIncognitoProxy)
// 				if err := _Vault.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
// 					return err
// 				}
// 				event.Raw = log

// 				select {
// 				case sink <- event:
// 				case err := <-sub.Err():
// 					return err
// 				case <-quit:
// 					return nil
// 				}
// 			case err := <-sub.Err():
// 				return err
// 			case <-quit:
// 				return nil
// 			}
// 		}
// 	}), nil
// }

// // ParseUpdateIncognitoProxy is a log parse operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
// //
// // Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
// func (_Vault *VaultFilterer) ParseUpdateIncognitoProxy(log types.Log) (*VaultUpdateIncognitoProxy, error) {
// 	event := new(VaultUpdateIncognitoProxy)
// 	if err := _Vault.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
// 		return nil, err
// 	}
// 	event.Raw = log
// 	return event, nil
// }

// // VaultUpdateTokenTotalIterator is returned from FilterUpdateTokenTotal and is used to iterate over the raw logs and unpacked data for UpdateTokenTotal events raised by the Vault contract.
// type VaultUpdateTokenTotalIterator struct {
// 	Event *VaultUpdateTokenTotal // Event containing the contract specifics and raw log

// 	contract *bind.BoundContract // Generic contract to use for unpacking event data
// 	event    string              // Event name to use for unpacking event data

// 	logs chan types.Log        // Log channel receiving the found contract events
// 	sub  ethereum.Subscription // Subscription for errors, completion and termination
// 	done bool                  // Whether the subscription completed delivering logs
// 	fail error                 // Occurred error to stop iteration
// }

// // Next advances the iterator to the subsequent event, returning whether there
// // are any more events found. In case of a retrieval or parsing error, false is
// // returned and Error() can be queried for the exact failure.
// func (it *VaultUpdateTokenTotalIterator) Next() bool {
// 	// If the iterator failed, stop iterating
// 	if it.fail != nil {
// 		return false
// 	}
// 	// If the iterator completed, deliver directly whatever's available
// 	if it.done {
// 		select {
// 		case log := <-it.logs:
// 			it.Event = new(VaultUpdateTokenTotal)
// 			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 				it.fail = err
// 				return false
// 			}
// 			it.Event.Raw = log
// 			return true

// 		default:
// 			return false
// 		}
// 	}
// 	// Iterator still in progress, wait for either a data or an error event
// 	select {
// 	case log := <-it.logs:
// 		it.Event = new(VaultUpdateTokenTotal)
// 		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 			it.fail = err
// 			return false
// 		}
// 		it.Event.Raw = log
// 		return true

// 	case err := <-it.sub.Err():
// 		it.done = true
// 		it.fail = err
// 		return it.Next()
// 	}
// }

// // Error returns any retrieval or parsing error occurred during filtering.
// func (it *VaultUpdateTokenTotalIterator) Error() error {
// 	return it.fail
// }

// // Close terminates the iteration process, releasing any pending underlying
// // resources.
// func (it *VaultUpdateTokenTotalIterator) Close() error {
// 	it.sub.Unsubscribe()
// 	return nil
// }

// // VaultUpdateTokenTotal represents a UpdateTokenTotal event raised by the Vault contract.
// type VaultUpdateTokenTotal struct {
// 	Assets  []common.Address
// 	Amounts []*big.Int
// 	Raw     types.Log // Blockchain specific contextual infos
// }

// // FilterUpdateTokenTotal is a free log retrieval operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
// //
// // Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
// func (_Vault *VaultFilterer) FilterUpdateTokenTotal(opts *bind.FilterOpts) (*VaultUpdateTokenTotalIterator, error) {

// 	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateTokenTotal")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultUpdateTokenTotalIterator{contract: _Vault.contract, event: "UpdateTokenTotal", logs: logs, sub: sub}, nil
// }

// // WatchUpdateTokenTotal is a free log subscription operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
// //
// // Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
// func (_Vault *VaultFilterer) WatchUpdateTokenTotal(opts *bind.WatchOpts, sink chan<- *VaultUpdateTokenTotal) (event.Subscription, error) {

// 	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateTokenTotal")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return event.NewSubscription(func(quit <-chan struct{}) error {
// 		defer sub.Unsubscribe()
// 		for {
// 			select {
// 			case log := <-logs:
// 				// New log arrived, parse the event and forward to the user
// 				event := new(VaultUpdateTokenTotal)
// 				if err := _Vault.contract.UnpackLog(event, "UpdateTokenTotal", log); err != nil {
// 					return err
// 				}
// 				event.Raw = log

// 				select {
// 				case sink <- event:
// 				case err := <-sub.Err():
// 					return err
// 				case <-quit:
// 					return nil
// 				}
// 			case err := <-sub.Err():
// 				return err
// 			case <-quit:
// 				return nil
// 			}
// 		}
// 	}), nil
// }

// // ParseUpdateTokenTotal is a log parse operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
// //
// // Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
// func (_Vault *VaultFilterer) ParseUpdateTokenTotal(log types.Log) (*VaultUpdateTokenTotal, error) {
// 	event := new(VaultUpdateTokenTotal)
// 	if err := _Vault.contract.UnpackLog(event, "UpdateTokenTotal", log); err != nil {
// 		return nil, err
// 	}
// 	event.Raw = log
// 	return event, nil
// }

// // VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vault contract.
// type VaultWithdrawIterator struct {
// 	Event *VaultWithdraw // Event containing the contract specifics and raw log

// 	contract *bind.BoundContract // Generic contract to use for unpacking event data
// 	event    string              // Event name to use for unpacking event data

// 	logs chan types.Log        // Log channel receiving the found contract events
// 	sub  ethereum.Subscription // Subscription for errors, completion and termination
// 	done bool                  // Whether the subscription completed delivering logs
// 	fail error                 // Occurred error to stop iteration
// }

// // Next advances the iterator to the subsequent event, returning whether there
// // are any more events found. In case of a retrieval or parsing error, false is
// // returned and Error() can be queried for the exact failure.
// func (it *VaultWithdrawIterator) Next() bool {
// 	// If the iterator failed, stop iterating
// 	if it.fail != nil {
// 		return false
// 	}
// 	// If the iterator completed, deliver directly whatever's available
// 	if it.done {
// 		select {
// 		case log := <-it.logs:
// 			it.Event = new(VaultWithdraw)
// 			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 				it.fail = err
// 				return false
// 			}
// 			it.Event.Raw = log
// 			return true

// 		default:
// 			return false
// 		}
// 	}
// 	// Iterator still in progress, wait for either a data or an error event
// 	select {
// 	case log := <-it.logs:
// 		it.Event = new(VaultWithdraw)
// 		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
// 			it.fail = err
// 			return false
// 		}
// 		it.Event.Raw = log
// 		return true

// 	case err := <-it.sub.Err():
// 		it.done = true
// 		it.fail = err
// 		return it.Next()
// 	}
// }

// // Error returns any retrieval or parsing error occurred during filtering.
// func (it *VaultWithdrawIterator) Error() error {
// 	return it.fail
// }

// // Close terminates the iteration process, releasing any pending underlying
// // resources.
// func (it *VaultWithdrawIterator) Close() error {
// 	it.sub.Unsubscribe()
// 	return nil
// }

// // VaultWithdraw represents a Withdraw event raised by the Vault contract.
// type VaultWithdraw struct {
// 	Token  common.Address
// 	To     common.Address
// 	Amount *big.Int
// 	Raw    types.Log // Blockchain specific contextual infos
// }

// // FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
// //
// // Solidity: event Withdraw(address token, address to, uint256 amount)
// func (_Vault *VaultFilterer) FilterWithdraw(opts *bind.FilterOpts) (*VaultWithdrawIterator, error) {

// 	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdraw")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &VaultWithdrawIterator{contract: _Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
// }

// // WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
// //
// // Solidity: event Withdraw(address token, address to, uint256 amount)
// func (_Vault *VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VaultWithdraw) (event.Subscription, error) {

// 	logs, sub, err := _Vault.contract.WatchLogs(opts, "Withdraw")
// 	if err != nil {
// 		return nil, err
// 	}
// 	return event.NewSubscription(func(quit <-chan struct{}) error {
// 		defer sub.Unsubscribe()
// 		for {
// 			select {
// 			case log := <-logs:
// 				// New log arrived, parse the event and forward to the user
// 				event := new(VaultWithdraw)
// 				if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
// 					return err
// 				}
// 				event.Raw = log

// 				select {
// 				case sink <- event:
// 				case err := <-sub.Err():
// 					return err
// 				case <-quit:
// 					return nil
// 				}
// 			case err := <-sub.Err():
// 				return err
// 			case <-quit:
// 				return nil
// 			}
// 		}
// 	}), nil
// }

// // ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
// //
// // Solidity: event Withdraw(address token, address to, uint256 amount)
// func (_Vault *VaultFilterer) ParseWithdraw(log types.Log) (*VaultWithdraw, error) {
// 	event := new(VaultWithdraw)
// 	if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
// 		return nil, err
// 	}
// 	event.Raw = log
// 	return event, nil
// }

// // WithdrawableABI is the input ABI used to generate the binding from.
// const WithdrawableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// // WithdrawableFuncSigs maps the 4-byte function signature to its string representation.
// var WithdrawableFuncSigs = map[string]string{
// 	"f75b98ce": "getDepositedBalance(address,address)",
// 	"e4bd7074": "isSigDataUsed(bytes32)",
// 	"749c5f86": "isWithdrawed(bytes32)",
// 	"5c975abb": "paused()",
// 	"1ed4276d": "updateAssets(address[],uint256[])",
// }

// // Withdrawable is an auto generated Go binding around an Ethereum contract.
// type Withdrawable struct {
// 	WithdrawableCaller     // Read-only binding to the contract
// 	WithdrawableTransactor // Write-only binding to the contract
// 	WithdrawableFilterer   // Log filterer for contract events
// }

// // WithdrawableCaller is an auto generated read-only Go binding around an Ethereum contract.
// type WithdrawableCaller struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // WithdrawableTransactor is an auto generated write-only Go binding around an Ethereum contract.
// type WithdrawableTransactor struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // WithdrawableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
// type WithdrawableFilterer struct {
// 	contract *bind.BoundContract // Generic contract wrapper for the low level calls
// }

// // WithdrawableSession is an auto generated Go binding around an Ethereum contract,
// // with pre-set call and transact options.
// type WithdrawableSession struct {
// 	Contract     *Withdrawable     // Generic contract binding to set the session for
// 	CallOpts     bind.CallOpts     // Call options to use throughout this session
// 	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
// }

// // WithdrawableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// // with pre-set call options.
// type WithdrawableCallerSession struct {
// 	Contract *WithdrawableCaller // Generic contract caller binding to set the session for
// 	CallOpts bind.CallOpts       // Call options to use throughout this session
// }

// // WithdrawableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// // with pre-set transact options.
// type WithdrawableTransactorSession struct {
// 	Contract     *WithdrawableTransactor // Generic contract transactor binding to set the session for
// 	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
// }

// // WithdrawableRaw is an auto generated low-level Go binding around an Ethereum contract.
// type WithdrawableRaw struct {
// 	Contract *Withdrawable // Generic contract binding to access the raw methods on
// }

// // WithdrawableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
// type WithdrawableCallerRaw struct {
// 	Contract *WithdrawableCaller // Generic read-only contract binding to access the raw methods on
// }

// // WithdrawableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
// type WithdrawableTransactorRaw struct {
// 	Contract *WithdrawableTransactor // Generic write-only contract binding to access the raw methods on
// }

// // NewWithdrawable creates a new instance of Withdrawable, bound to a specific deployed contract.
// func NewWithdrawable(address common.Address, backend bind.ContractBackend) (*Withdrawable, error) {
// 	contract, err := bindWithdrawable(address, backend, backend, backend)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &Withdrawable{WithdrawableCaller: WithdrawableCaller{contract: contract}, WithdrawableTransactor: WithdrawableTransactor{contract: contract}, WithdrawableFilterer: WithdrawableFilterer{contract: contract}}, nil
// }

// // NewWithdrawableCaller creates a new read-only instance of Withdrawable, bound to a specific deployed contract.
// func NewWithdrawableCaller(address common.Address, caller bind.ContractCaller) (*WithdrawableCaller, error) {
// 	contract, err := bindWithdrawable(address, caller, nil, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &WithdrawableCaller{contract: contract}, nil
// }

// // NewWithdrawableTransactor creates a new write-only instance of Withdrawable, bound to a specific deployed contract.
// func NewWithdrawableTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawableTransactor, error) {
// 	contract, err := bindWithdrawable(address, nil, transactor, nil)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &WithdrawableTransactor{contract: contract}, nil
// }

// // NewWithdrawableFilterer creates a new log filterer instance of Withdrawable, bound to a specific deployed contract.
// func NewWithdrawableFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawableFilterer, error) {
// 	contract, err := bindWithdrawable(address, nil, nil, filterer)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &WithdrawableFilterer{contract: contract}, nil
// }

// // bindWithdrawable binds a generic wrapper to an already deployed contract.
// func bindWithdrawable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
// 	parsed, err := abi.JSON(strings.NewReader(WithdrawableABI))
// 	if err != nil {
// 		return nil, err
// 	}
// 	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Withdrawable *WithdrawableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Withdrawable.Contract.WithdrawableCaller.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Withdrawable *WithdrawableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Withdrawable.Contract.WithdrawableTransactor.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Withdrawable *WithdrawableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Withdrawable.Contract.WithdrawableTransactor.contract.Transact(opts, method, params...)
// }

// // Call invokes the (constant) contract method with params as input values and
// // sets the output to result. The result type might be a single field for simple
// // returns, a slice of interfaces for anonymous returns and a struct for named
// // returns.
// func (_Withdrawable *WithdrawableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
// 	return _Withdrawable.Contract.contract.Call(opts, result, method, params...)
// }

// // Transfer initiates a plain transaction to move funds to the contract, calling
// // its default method if one is available.
// func (_Withdrawable *WithdrawableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
// 	return _Withdrawable.Contract.contract.Transfer(opts)
// }

// // Transact invokes the (paid) contract method with params as input values.
// func (_Withdrawable *WithdrawableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
// 	return _Withdrawable.Contract.contract.Transact(opts, method, params...)
// }

// // GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
// //
// // Solidity: function getDepositedBalance(address , address ) view returns(uint256)
// func (_Withdrawable *WithdrawableCaller) GetDepositedBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
// 	var out []interface{}
// 	err := _Withdrawable.contract.Call(opts, &out, "getDepositedBalance", arg0, arg1)

// 	if err != nil {
// 		return *new(*big.Int), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

// 	return out0, err

// }

// // GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
// //
// // Solidity: function getDepositedBalance(address , address ) view returns(uint256)
// func (_Withdrawable *WithdrawableSession) GetDepositedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
// 	return _Withdrawable.Contract.GetDepositedBalance(&_Withdrawable.CallOpts, arg0, arg1)
// }

// // GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
// //
// // Solidity: function getDepositedBalance(address , address ) view returns(uint256)
// func (_Withdrawable *WithdrawableCallerSession) GetDepositedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
// 	return _Withdrawable.Contract.GetDepositedBalance(&_Withdrawable.CallOpts, arg0, arg1)
// }

// // IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
// //
// // Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
// func (_Withdrawable *WithdrawableCaller) IsSigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
// 	var out []interface{}
// 	err := _Withdrawable.contract.Call(opts, &out, "isSigDataUsed", arg0)

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
// //
// // Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
// func (_Withdrawable *WithdrawableSession) IsSigDataUsed(arg0 [32]byte) (bool, error) {
// 	return _Withdrawable.Contract.IsSigDataUsed(&_Withdrawable.CallOpts, arg0)
// }

// // IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
// //
// // Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
// func (_Withdrawable *WithdrawableCallerSession) IsSigDataUsed(arg0 [32]byte) (bool, error) {
// 	return _Withdrawable.Contract.IsSigDataUsed(&_Withdrawable.CallOpts, arg0)
// }

// // IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
// //
// // Solidity: function isWithdrawed(bytes32 ) view returns(bool)
// func (_Withdrawable *WithdrawableCaller) IsWithdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
// 	var out []interface{}
// 	err := _Withdrawable.contract.Call(opts, &out, "isWithdrawed", arg0)

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
// //
// // Solidity: function isWithdrawed(bytes32 ) view returns(bool)
// func (_Withdrawable *WithdrawableSession) IsWithdrawed(arg0 [32]byte) (bool, error) {
// 	return _Withdrawable.Contract.IsWithdrawed(&_Withdrawable.CallOpts, arg0)
// }

// // IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
// //
// // Solidity: function isWithdrawed(bytes32 ) view returns(bool)
// func (_Withdrawable *WithdrawableCallerSession) IsWithdrawed(arg0 [32]byte) (bool, error) {
// 	return _Withdrawable.Contract.IsWithdrawed(&_Withdrawable.CallOpts, arg0)
// }

// // Paused is a free data retrieval call binding the contract method 0x5c975abb.
// //
// // Solidity: function paused() view returns(bool)
// func (_Withdrawable *WithdrawableCaller) Paused(opts *bind.CallOpts) (bool, error) {
// 	var out []interface{}
// 	err := _Withdrawable.contract.Call(opts, &out, "paused")

// 	if err != nil {
// 		return *new(bool), err
// 	}

// 	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

// 	return out0, err

// }

// // Paused is a free data retrieval call binding the contract method 0x5c975abb.
// //
// // Solidity: function paused() view returns(bool)
// func (_Withdrawable *WithdrawableSession) Paused() (bool, error) {
// 	return _Withdrawable.Contract.Paused(&_Withdrawable.CallOpts)
// }

// // Paused is a free data retrieval call binding the contract method 0x5c975abb.
// //
// // Solidity: function paused() view returns(bool)
// func (_Withdrawable *WithdrawableCallerSession) Paused() (bool, error) {
// 	return _Withdrawable.Contract.Paused(&_Withdrawable.CallOpts)
// }

// // UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
// //
// // Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
// func (_Withdrawable *WithdrawableTransactor) UpdateAssets(opts *bind.TransactOpts, arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
// 	return _Withdrawable.contract.Transact(opts, "updateAssets", arg0, arg1)
// }

// // UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
// //
// // Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
// func (_Withdrawable *WithdrawableSession) UpdateAssets(arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
// 	return _Withdrawable.Contract.UpdateAssets(&_Withdrawable.TransactOpts, arg0, arg1)
// }

// // UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
// //
// // Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
// func (_Withdrawable *WithdrawableTransactorSession) UpdateAssets(arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
// 	return _Withdrawable.Contract.UpdateAssets(&_Withdrawable.TransactOpts, arg0, arg1)
// }