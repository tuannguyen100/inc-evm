// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VaultBurnInstData is an auto generated low-level Go binding around an user-defined struct.
type VaultBurnInstData struct {
	Meta   uint8
	Shard  uint8
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Itx    [32]byte
}

// VaultRedepositOptions is an auto generated low-level Go binding around an user-defined struct.
type VaultRedepositOptions struct {
	RedepositToken      common.Address
	RedepositIncAddress []byte
	WithdrawAddress     common.Address
}

// AddressMetaData contains all meta data concerning the Address contract.
var AddressMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122056923f7f32633c310e8d7f6c1b5fdd5c2fb7886b521493707effd7774c4b7fce64736f6c634300060c0033",
}

// AddressABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressMetaData.ABI instead.
var AddressABI = AddressMetaData.ABI

// AddressBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressMetaData.Bin instead.
var AddressBin = AddressMetaData.Bin

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := AddressMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// CountersMetaData contains all meta data concerning the Counters contract.
var CountersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122052eead925374c42118905eec33485a28f5701bd939bc23becb57b0a115d066b464736f6c634300060c0033",
}

// CountersABI is the input ABI used to generate the binding from.
// Deprecated: Use CountersMetaData.ABI instead.
var CountersABI = CountersMetaData.ABI

// CountersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CountersMetaData.Bin instead.
var CountersBin = CountersMetaData.Bin

// DeployCounters deploys a new Ethereum contract, binding an instance of Counters to it.
func DeployCounters(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counters, error) {
	parsed, err := CountersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CountersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// Counters is an auto generated Go binding around an Ethereum contract.
type Counters struct {
	CountersCaller     // Read-only binding to the contract
	CountersTransactor // Write-only binding to the contract
	CountersFilterer   // Log filterer for contract events
}

// CountersCaller is an auto generated read-only Go binding around an Ethereum contract.
type CountersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CountersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CountersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CountersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CountersSession struct {
	Contract     *Counters         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CountersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CountersCallerSession struct {
	Contract *CountersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CountersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CountersTransactorSession struct {
	Contract     *CountersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CountersRaw is an auto generated low-level Go binding around an Ethereum contract.
type CountersRaw struct {
	Contract *Counters // Generic contract binding to access the raw methods on
}

// CountersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CountersCallerRaw struct {
	Contract *CountersCaller // Generic read-only contract binding to access the raw methods on
}

// CountersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CountersTransactorRaw struct {
	Contract *CountersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounters creates a new instance of Counters, bound to a specific deployed contract.
func NewCounters(address common.Address, backend bind.ContractBackend) (*Counters, error) {
	contract, err := bindCounters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counters{CountersCaller: CountersCaller{contract: contract}, CountersTransactor: CountersTransactor{contract: contract}, CountersFilterer: CountersFilterer{contract: contract}}, nil
}

// NewCountersCaller creates a new read-only instance of Counters, bound to a specific deployed contract.
func NewCountersCaller(address common.Address, caller bind.ContractCaller) (*CountersCaller, error) {
	contract, err := bindCounters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CountersCaller{contract: contract}, nil
}

// NewCountersTransactor creates a new write-only instance of Counters, bound to a specific deployed contract.
func NewCountersTransactor(address common.Address, transactor bind.ContractTransactor) (*CountersTransactor, error) {
	contract, err := bindCounters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CountersTransactor{contract: contract}, nil
}

// NewCountersFilterer creates a new log filterer instance of Counters, bound to a specific deployed contract.
func NewCountersFilterer(address common.Address, filterer bind.ContractFilterer) (*CountersFilterer, error) {
	contract, err := bindCounters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CountersFilterer{contract: contract}, nil
}

// bindCounters binds a generic wrapper to an already deployed contract.
func bindCounters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CountersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.CountersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.CountersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counters *CountersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counters *CountersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counters *CountersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counters.Contract.contract.Transact(opts, method, params...)
}

// ExecutorMetaData contains all meta data concerning the Executor contract.
var ExecutorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fns\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"1cff79cd": "execute(address,bytes)",
	},
}

// ExecutorABI is the input ABI used to generate the binding from.
// Deprecated: Use ExecutorMetaData.ABI instead.
var ExecutorABI = ExecutorMetaData.ABI

// Deprecated: Use ExecutorMetaData.Sigs instead.
// ExecutorFuncSigs maps the 4-byte function signature to its string representation.
var ExecutorFuncSigs = ExecutorMetaData.Sigs

// Executor is an auto generated Go binding around an Ethereum contract.
type Executor struct {
	ExecutorCaller     // Read-only binding to the contract
	ExecutorTransactor // Write-only binding to the contract
	ExecutorFilterer   // Log filterer for contract events
}

// ExecutorCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExecutorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExecutorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExecutorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExecutorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExecutorSession struct {
	Contract     *Executor         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExecutorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExecutorCallerSession struct {
	Contract *ExecutorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExecutorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExecutorTransactorSession struct {
	Contract     *ExecutorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExecutorRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExecutorRaw struct {
	Contract *Executor // Generic contract binding to access the raw methods on
}

// ExecutorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExecutorCallerRaw struct {
	Contract *ExecutorCaller // Generic read-only contract binding to access the raw methods on
}

// ExecutorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExecutorTransactorRaw struct {
	Contract *ExecutorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExecutor creates a new instance of Executor, bound to a specific deployed contract.
func NewExecutor(address common.Address, backend bind.ContractBackend) (*Executor, error) {
	contract, err := bindExecutor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Executor{ExecutorCaller: ExecutorCaller{contract: contract}, ExecutorTransactor: ExecutorTransactor{contract: contract}, ExecutorFilterer: ExecutorFilterer{contract: contract}}, nil
}

// NewExecutorCaller creates a new read-only instance of Executor, bound to a specific deployed contract.
func NewExecutorCaller(address common.Address, caller bind.ContractCaller) (*ExecutorCaller, error) {
	contract, err := bindExecutor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorCaller{contract: contract}, nil
}

// NewExecutorTransactor creates a new write-only instance of Executor, bound to a specific deployed contract.
func NewExecutorTransactor(address common.Address, transactor bind.ContractTransactor) (*ExecutorTransactor, error) {
	contract, err := bindExecutor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExecutorTransactor{contract: contract}, nil
}

// NewExecutorFilterer creates a new log filterer instance of Executor, bound to a specific deployed contract.
func NewExecutorFilterer(address common.Address, filterer bind.ContractFilterer) (*ExecutorFilterer, error) {
	contract, err := bindExecutor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExecutorFilterer{contract: contract}, nil
}

// bindExecutor binds a generic wrapper to an already deployed contract.
func bindExecutor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExecutorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.ExecutorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.ExecutorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Executor *ExecutorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Executor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Executor *ExecutorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Executor *ExecutorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Executor.Contract.contract.Transact(opts, method, params...)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address fns, bytes data) payable returns(bytes)
func (_Executor *ExecutorTransactor) Execute(opts *bind.TransactOpts, fns common.Address, data []byte) (*types.Transaction, error) {
	return _Executor.contract.Transact(opts, "execute", fns, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address fns, bytes data) payable returns(bytes)
func (_Executor *ExecutorSession) Execute(fns common.Address, data []byte) (*types.Transaction, error) {
	return _Executor.Contract.Execute(&_Executor.TransactOpts, fns, data)
}

// Execute is a paid mutator transaction binding the contract method 0x1cff79cd.
//
// Solidity: function execute(address fns, bytes data) payable returns(bytes)
func (_Executor *ExecutorTransactorSession) Execute(fns common.Address, data []byte) (*types.Transaction, error) {
	return _Executor.Contract.Execute(&_Executor.TransactOpts, fns, data)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_IERC20 *IERC20Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_IERC20 *IERC20Session) Decimals() (*big.Int, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_IERC20 *IERC20CallerSession) Decimals() (*big.Int, error) {
	return _IERC20.Contract.Decimals(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns()
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns()
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns()
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns()
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns()
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IERC20Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IncognitoMetaData contains all meta data concerning the Incognito contract.
var IncognitoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"name\":\"instructionApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f65d2116": "instructionApproved(bool,bytes32,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
	},
}

// IncognitoABI is the input ABI used to generate the binding from.
// Deprecated: Use IncognitoMetaData.ABI instead.
var IncognitoABI = IncognitoMetaData.ABI

// Deprecated: Use IncognitoMetaData.Sigs instead.
// IncognitoFuncSigs maps the 4-byte function signature to its string representation.
var IncognitoFuncSigs = IncognitoMetaData.Sigs

// Incognito is an auto generated Go binding around an Ethereum contract.
type Incognito struct {
	IncognitoCaller     // Read-only binding to the contract
	IncognitoTransactor // Write-only binding to the contract
	IncognitoFilterer   // Log filterer for contract events
}

// IncognitoCaller is an auto generated read-only Go binding around an Ethereum contract.
type IncognitoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IncognitoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IncognitoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IncognitoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IncognitoSession struct {
	Contract     *Incognito        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IncognitoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IncognitoCallerSession struct {
	Contract *IncognitoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IncognitoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IncognitoTransactorSession struct {
	Contract     *IncognitoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IncognitoRaw is an auto generated low-level Go binding around an Ethereum contract.
type IncognitoRaw struct {
	Contract *Incognito // Generic contract binding to access the raw methods on
}

// IncognitoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IncognitoCallerRaw struct {
	Contract *IncognitoCaller // Generic read-only contract binding to access the raw methods on
}

// IncognitoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IncognitoTransactorRaw struct {
	Contract *IncognitoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIncognito creates a new instance of Incognito, bound to a specific deployed contract.
func NewIncognito(address common.Address, backend bind.ContractBackend) (*Incognito, error) {
	contract, err := bindIncognito(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Incognito{IncognitoCaller: IncognitoCaller{contract: contract}, IncognitoTransactor: IncognitoTransactor{contract: contract}, IncognitoFilterer: IncognitoFilterer{contract: contract}}, nil
}

// NewIncognitoCaller creates a new read-only instance of Incognito, bound to a specific deployed contract.
func NewIncognitoCaller(address common.Address, caller bind.ContractCaller) (*IncognitoCaller, error) {
	contract, err := bindIncognito(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoCaller{contract: contract}, nil
}

// NewIncognitoTransactor creates a new write-only instance of Incognito, bound to a specific deployed contract.
func NewIncognitoTransactor(address common.Address, transactor bind.ContractTransactor) (*IncognitoTransactor, error) {
	contract, err := bindIncognito(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IncognitoTransactor{contract: contract}, nil
}

// NewIncognitoFilterer creates a new log filterer instance of Incognito, bound to a specific deployed contract.
func NewIncognitoFilterer(address common.Address, filterer bind.ContractFilterer) (*IncognitoFilterer, error) {
	contract, err := bindIncognito(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IncognitoFilterer{contract: contract}, nil
}

// bindIncognito binds a generic wrapper to an already deployed contract.
func bindIncognito(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IncognitoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incognito *IncognitoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Incognito.Contract.IncognitoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incognito *IncognitoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognito.Contract.IncognitoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incognito *IncognitoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incognito.Contract.IncognitoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Incognito *IncognitoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Incognito.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Incognito *IncognitoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Incognito.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Incognito *IncognitoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Incognito.Contract.contract.Transact(opts, method, params...)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
func (_Incognito *IncognitoCaller) InstructionApproved(opts *bind.CallOpts, arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	var out []interface{}
	err := _Incognito.contract.Call(opts, &out, "instructionApproved", arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
func (_Incognito *IncognitoSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// InstructionApproved is a free data retrieval call binding the contract method 0xf65d2116.
//
// Solidity: function instructionApproved(bool , bytes32 , uint256 , bytes32[] , bool[] , bytes32 , bytes32 , uint256[] , uint8[] , bytes32[] , bytes32[] ) view returns(bool)
func (_Incognito *IncognitoCallerSession) InstructionApproved(arg0 bool, arg1 [32]byte, arg2 *big.Int, arg3 [][32]byte, arg4 []bool, arg5 [32]byte, arg6 [32]byte, arg7 []*big.Int, arg8 []uint8, arg9 [][32]byte, arg10 [][32]byte) (bool, error) {
	return _Incognito.Contract.InstructionApproved(&_Incognito.CallOpts, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202dc5c92709e9348a27c5656f1b5eb20ce8376d6c0c1cecc93fbe589f292cc80864736f6c634300060c0033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// TradeUtilsMetaData contains all meta data concerning the TradeUtils contract.
var TradeUtilsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ETH_CONTRACT_ADDRESS\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"72e94bf6": "ETH_CONTRACT_ADDRESS()",
	},
	Bin: "0x6080604052348015600f57600080fd5b50608a8061001e6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c806372e94bf614602d575b600080fd5b6033604f565b604080516001600160a01b039092168252519081900360200190f35b60008156fea26469706673582212200b6ece422dde4ee6862b0bf1c81e839133801e217820818bd33158d845a6f50b64736f6c634300060c0033",
}

// TradeUtilsABI is the input ABI used to generate the binding from.
// Deprecated: Use TradeUtilsMetaData.ABI instead.
var TradeUtilsABI = TradeUtilsMetaData.ABI

// Deprecated: Use TradeUtilsMetaData.Sigs instead.
// TradeUtilsFuncSigs maps the 4-byte function signature to its string representation.
var TradeUtilsFuncSigs = TradeUtilsMetaData.Sigs

// TradeUtilsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TradeUtilsMetaData.Bin instead.
var TradeUtilsBin = TradeUtilsMetaData.Bin

// DeployTradeUtils deploys a new Ethereum contract, binding an instance of TradeUtils to it.
func DeployTradeUtils(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TradeUtils, error) {
	parsed, err := TradeUtilsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TradeUtilsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TradeUtils{TradeUtilsCaller: TradeUtilsCaller{contract: contract}, TradeUtilsTransactor: TradeUtilsTransactor{contract: contract}, TradeUtilsFilterer: TradeUtilsFilterer{contract: contract}}, nil
}

// TradeUtils is an auto generated Go binding around an Ethereum contract.
type TradeUtils struct {
	TradeUtilsCaller     // Read-only binding to the contract
	TradeUtilsTransactor // Write-only binding to the contract
	TradeUtilsFilterer   // Log filterer for contract events
}

// TradeUtilsCaller is an auto generated read-only Go binding around an Ethereum contract.
type TradeUtilsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeUtilsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TradeUtilsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeUtilsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TradeUtilsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TradeUtilsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TradeUtilsSession struct {
	Contract     *TradeUtils       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TradeUtilsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TradeUtilsCallerSession struct {
	Contract *TradeUtilsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TradeUtilsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TradeUtilsTransactorSession struct {
	Contract     *TradeUtilsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TradeUtilsRaw is an auto generated low-level Go binding around an Ethereum contract.
type TradeUtilsRaw struct {
	Contract *TradeUtils // Generic contract binding to access the raw methods on
}

// TradeUtilsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TradeUtilsCallerRaw struct {
	Contract *TradeUtilsCaller // Generic read-only contract binding to access the raw methods on
}

// TradeUtilsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TradeUtilsTransactorRaw struct {
	Contract *TradeUtilsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTradeUtils creates a new instance of TradeUtils, bound to a specific deployed contract.
func NewTradeUtils(address common.Address, backend bind.ContractBackend) (*TradeUtils, error) {
	contract, err := bindTradeUtils(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TradeUtils{TradeUtilsCaller: TradeUtilsCaller{contract: contract}, TradeUtilsTransactor: TradeUtilsTransactor{contract: contract}, TradeUtilsFilterer: TradeUtilsFilterer{contract: contract}}, nil
}

// NewTradeUtilsCaller creates a new read-only instance of TradeUtils, bound to a specific deployed contract.
func NewTradeUtilsCaller(address common.Address, caller bind.ContractCaller) (*TradeUtilsCaller, error) {
	contract, err := bindTradeUtils(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TradeUtilsCaller{contract: contract}, nil
}

// NewTradeUtilsTransactor creates a new write-only instance of TradeUtils, bound to a specific deployed contract.
func NewTradeUtilsTransactor(address common.Address, transactor bind.ContractTransactor) (*TradeUtilsTransactor, error) {
	contract, err := bindTradeUtils(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TradeUtilsTransactor{contract: contract}, nil
}

// NewTradeUtilsFilterer creates a new log filterer instance of TradeUtils, bound to a specific deployed contract.
func NewTradeUtilsFilterer(address common.Address, filterer bind.ContractFilterer) (*TradeUtilsFilterer, error) {
	contract, err := bindTradeUtils(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TradeUtilsFilterer{contract: contract}, nil
}

// bindTradeUtils binds a generic wrapper to an already deployed contract.
func bindTradeUtils(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TradeUtilsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeUtils *TradeUtilsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradeUtils.Contract.TradeUtilsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeUtils *TradeUtilsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeUtils.Contract.TradeUtilsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeUtils *TradeUtilsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeUtils.Contract.TradeUtilsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TradeUtils *TradeUtilsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TradeUtils.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TradeUtils *TradeUtilsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TradeUtils.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TradeUtils *TradeUtilsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TradeUtils.Contract.contract.Transact(opts, method, params...)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_TradeUtils *TradeUtilsCaller) ETHCONTRACTADDRESS(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TradeUtils.contract.Call(opts, &out, "ETH_CONTRACT_ADDRESS")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_TradeUtils *TradeUtilsSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _TradeUtils.Contract.ETHCONTRACTADDRESS(&_TradeUtils.CallOpts)
}

// ETHCONTRACTADDRESS is a free data retrieval call binding the contract method 0x72e94bf6.
//
// Solidity: function ETH_CONTRACT_ADDRESS() view returns(address)
func (_TradeUtils *TradeUtilsCallerSession) ETHCONTRACTADDRESS() (common.Address, error) {
	return _TradeUtils.Contract.ETHCONTRACTADDRESS(&_TradeUtils.CallOpts)
}

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositID\",\"type\":\"uint256\"}],\"name\":\"DepositV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"phaseID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"errorData\",\"type\":\"bytes\"}],\"name\":\"ExecuteFnLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"redepositIncAddress\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"name\":\"Redeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newIncognitoProxy\",\"type\":\"address\"}],\"name\":\"UpdateIncognitoProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"UpdateTokenTotal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURN_CALL_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_TO_CONTRACT_REQUEST_METADATA_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CURRENT_NETWORK_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ETH_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"externalCalldata\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"redepositToken\",\"type\":\"address\"}],\"name\":\"_callExternal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"_transferExternal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"depositERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"depositERC20_V2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"deposit_V2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipientToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exchangeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"executeWithBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_prevVault\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_regulator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"migration\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"}],\"name\":\"parseCalldataFromBurnInst\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"meta\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"shard\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"itx\",\"type\":\"bytes32\"}],\"internalType\":\"structVault.BurnInstData\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"redepositToken\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"redepositIncAddress\",\"type\":\"bytes\"},{\"internalType\":\"addresspayable\",\"name\":\"withdrawAddress\",\"type\":\"address\"}],\"internalType\":\"structVault.RedepositOptions\",\"name\":\"\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prevVault\",\"outputs\":[{\"internalType\":\"contractWithdrawable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"regulator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"incognitoAddress\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"timestamp\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"txId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"regulatorSig\",\"type\":\"bytes\"}],\"name\":\"requestWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_regulator\",\"type\":\"address\"}],\"name\":\"setRegulator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"sigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"signData\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"sigToAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"storageLayoutVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"submitBurnProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"totalDepositedToSCAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_regulator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_executor\",\"type\":\"address\"}],\"name\":\"upgradeVaultStorage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"inst\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"heights\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"instPaths\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool[]\",\"name\":\"instPathIsLefts\",\"type\":\"bool[]\"},{\"internalType\":\"bytes32\",\"name\":\"instRoots\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"blkData\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"sigIdxs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint8[]\",\"name\":\"sigVs\",\"type\":\"uint8[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigRs\",\"type\":\"bytes32[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"sigSs\",\"type\":\"bytes32[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"withdrawRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Sigs: map[string]string{
		"bd835c42": "BURN_CALL_REQUEST_METADATA_TYPE()",
		"568c04fd": "BURN_REQUEST_METADATA_TYPE()",
		"6f2cbc48": "BURN_TO_CONTRACT_REQUEST_METADATA_TYPE()",
		"d7200eb1": "CURRENT_NETWORK_ID()",
		"58bc8337": "ETH_TOKEN()",
		"bda9b509": "_callExternal(address,address,uint256,bytes,address)",
		"145e2a6b": "_transferExternal(address,address,uint256)",
		"70a08231": "balanceOf(address)",
		"c791d705": "deposit(string,bytes32,bytes)",
		"a807b5bb": "depositERC20(address,uint256,string,bytes32,bytes)",
		"b8237dbb": "depositERC20_V2(address,uint256,string,bytes32,bytes)",
		"84b3ac03": "deposit_V2(string,bytes32,bytes)",
		"8588ccd6": "execute(address,uint256,address,address,bytes,bytes,bytes)",
		"3ed1b376": "executeWithBurnProof(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
		"c34c08e5": "executor()",
		"cf54aaa0": "getDecimals(address)",
		"f75b98ce": "getDepositedBalance(address,address)",
		"c0c53b8b": "initialize(address,address,address)",
		"392e53cd": "isInitialized()",
		"e4bd7074": "isSigDataUsed(bytes32)",
		"749c5f86": "isWithdrawed(bytes32)",
		"995fac11": "migration(address,address)",
		"a3f5d8cc": "notEntered()",
		"7e16e6e1": "parseBurnInst(bytes)",
		"66945b31": "parseCalldataFromBurnInst(bytes)",
		"fa84702e": "prevVault()",
		"dd8fee14": "regulator()",
		"fee8efda": "requestWithdraw(string,address,uint256,bytes,bytes,bytes32,bytes)",
		"cde0a4f8": "setRegulator(address)",
		"1ea1940e": "sigDataUsed(bytes32)",
		"3fec6b40": "sigToAddress(bytes,bytes32)",
		"d6a1fe3b": "storageLayoutVersion()",
		"73bf9651": "submitBurnProof(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
		"6304541c": "totalDepositedToSCAmount(address)",
		"1ed4276d": "updateAssets(address[],uint256[])",
		"a73b1532": "upgradeVaultStorage(address,address)",
		"1beb7de2": "withdraw(bytes,uint256,bytes32[],bool[],bytes32,bytes32,uint256[],uint8[],bytes32[],bytes32[])",
		"65b5a00f": "withdrawRequests(address,address)",
		"dca40d9e": "withdrawed(bytes32)",
	},
	Bin: "0x608060405234801561001057600080fd5b50614f01806100206000396000f3fe6080604052600436106102345760003560e01c8063995fac111161012e578063cde0a4f8116100ab578063dd8fee141161006f578063dd8fee1414610651578063e4bd707414610666578063f75b98ce14610686578063fa84702e146106a6578063fee8efda146106bb5761023b565b8063cde0a4f8146105c7578063cf54aaa0146105e7578063d6a1fe3b14610607578063d7200eb11461061c578063dca40d9e146106315761023b565b8063bd835c42116100f2578063bd835c421461054a578063bda9b5091461055f578063c0c53b8b1461057f578063c34c08e51461059f578063c791d705146105b45761023b565b8063995fac11146104b5578063a3f5d8cc146104d5578063a73b1532146104ea578063a807b5bb1461050a578063b8237dbb1461052a5761023b565b80636304541c116101bc57806373bf96511161018057806373bf965114610422578063749c5f86146104425780637e16e6e11461046257806384b3ac031461048f5780638588ccd6146104a25761023b565b80636304541c1461037157806365b5a00f1461039e57806366945b31146103be5780636f2cbc48146103ed57806370a08231146104025761023b565b8063392e53cd11610203578063392e53cd146102d85780633ed1b376146102ed5780633fec6b401461030d578063568c04fd1461033a57806358bc83371461035c5761023b565b8063145e2a6b146102405780631beb7de2146102625780631ea1940e146102825780631ed4276d146102b85761023b565b3661023b57005b600080fd5b34801561024c57600080fd5b5061026061025b366004613e6b565b6106db565b005b34801561026e57600080fd5b5061026061027d3660046143e4565b6107c4565b34801561028e57600080fd5b506102a261029d366004614161565b610c21565b6040516102af9190614a57565b60405180910390f35b3480156102c457600080fd5b506102a26102d33660046140dd565b610c36565b3480156102e457600080fd5b506102a2610e57565b3480156102f957600080fd5b506102606103083660046141b8565b610e67565b34801561031957600080fd5b5061032d6103283660046143a2565b611395565b6040516102af919061485e565b34801561034657600080fd5b5061034f611424565b6040516102af9190614ddb565b34801561036857600080fd5b5061032d611429565b34801561037d57600080fd5b5061039161038c366004613d74565b61142e565b6040516102af9190614b0f565b3480156103aa57600080fd5b506103916103b9366004613e33565b611440565b3480156103ca57600080fd5b506103de6103d9366004614179565b61145d565b6040516102af93929190614cdf565b3480156103f957600080fd5b5061034f611640565b34801561040e57600080fd5b5061039161041d366004613d74565b611645565b34801561042e57600080fd5b5061026061043d3660046143e4565b611719565b34801561044e57600080fd5b506102a261045d366004614161565b611a91565b34801561046e57600080fd5b5061048261047d3660046142f6565b611b48565b6040516102af9190614cd1565b61026061049d3660046145fc565b611bca565b6102606104b0366004613f73565b611cf3565b3480156104c157600080fd5b506102a26104d0366004613e33565b61218b565b3480156104e157600080fd5b506102a26121ab565b3480156104f657600080fd5b50610260610505366004613e33565b6121bb565b34801561051657600080fd5b5061026061052536600461404b565b612251565b34801561053657600080fd5b5061026061054536600461404b565b6124f0565b34801561055657600080fd5b5061034f6127de565b34801561056b57600080fd5b5061039161057a366004613ef5565b6127e3565b34801561058b57600080fd5b5061026061059a366004613eab565b612a07565b3480156105ab57600080fd5b5061032d612ad7565b6102606105c23660046145fc565b612ae6565b3480156105d357600080fd5b506102606105e2366004613d74565b612bba565b3480156105f357600080fd5b5061034f610602366004613d74565b612c3b565b34801561061357600080fd5b50610391612cf5565b34801561062857600080fd5b5061034f612cfb565b34801561063d57600080fd5b506102a261064c366004614161565b612d00565b34801561065d57600080fd5b5061032d612d15565b34801561067257600080fd5b506102a2610681366004614161565b612d24565b34801561069257600080fd5b506103916106a1366004613e33565b612d8b565b3480156106b257600080fd5b5061032d612eb7565b3480156106c757600080fd5b506102606106d6366004614515565b612ec6565b3033146106e86014613132565b9061070f5760405162461bcd60e51b81526004016107069190614b55565b60405180910390fd5b506001600160a01b03831661072d57610728828261325b565b6107bf565b60405163a9059cbb60e01b81526001600160a01b0384169063a9059cbb9061075b9085908590600401614896565b600060405180830381600087803b15801561077557600080fd5b505af1158015610789573d6000803e3d6000fd5b505050506107956132f7565b61079f6004613132565b906107bd5760405162461bcd60e51b81526004016107069190614b55565b505b505050565b600554600160a01b900460ff166107db6001613132565b906107f95760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b191690558951608211156108176006613132565b906108355760405162461bcd60e51b81526004016107069190614b55565b5061083e613aac565b6108478b611b48565b805190915060ff1660f11480156108655750806020015160ff166001145b61086f6006613132565b9061088d5760405162461bcd60e51b81526004016107069190614b55565b5061089b8160a00151611a91565b156108a66005613132565b906108c45760405162461bcd60e51b81526004016107069190614b55565b5060a081015160009081526020819052604090819020805460ff191660011790558101516001600160a01b031661094f576040808201516001600160a01b0316600090815260046020522054608082015161091e9161332b565b47101561092b6007613132565b906109495760405162461bcd60e51b81526004016107069190614b55565b50610a61565b600061095e8260400151612c3b565b905060098160ff16111561098b5760808201516109859060081960ff841601600a0a613391565b60808301525b6040808301516001600160a01b031660009081526004602052205460808301516109b49161332b565b82604001516001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016109e4919061485e565b60206040518083038186803b1580156109fc57600080fd5b505afa158015610a10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a349190614672565b1015610a406007613132565b90610a5e5760405162461bcd60e51b81526004016107069190614b55565b50505b610a738b8b8b8b8b8b8b8b8b8b6133f3565b60408101516001600160a01b0316610b1c57600081606001516001600160a01b03168260800151604051610aa69061485b565b60006040518083038185875af1925050503d8060008114610ae3576040519150601f19603f3d011682016040523d82523d6000602084013e610ae8565b606091505b5050905080610af76004613132565b90610b155760405162461bcd60e51b81526004016107069190614b55565b5050610bba565b80604001516001600160a01b031663a9059cbb826060015183608001516040518363ffffffff1660e01b8152600401610b56929190614896565b600060405180830381600087803b158015610b7057600080fd5b505af1158015610b84573d6000803e3d6000fd5b50505050610b906132f7565b610b9a6004613132565b90610bb85760405162461bcd60e51b81526004016107069190614b55565b505b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb816040015182606001518360800151604051610bf993929190614872565b60405180910390a150506005805460ff60a01b1916600160a01b179055505050505050505050565b60016020526000908152604090205460ff1681565b6005546000906001600160a01b031615801590610c5d57506005546001600160a01b031633145b610c67600c613132565b90610c855760405162461bcd60e51b81526004016107069190614b55565b50838214610c93600a613132565b90610cb15760405162461bcd60e51b81526004016107069190614b55565b50600560009054906101000a90046001600160a01b03166001600160a01b0316635c975abb6040518163ffffffff1660e01b815260040160206040518083038186803b158015610d0057600080fd5b505afa158015610d14573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d389190614145565b610d42600d613132565b90610d605760405162461bcd60e51b81526004016107069190614b55565b5060005b84811015610e0e57610dc7848483818110610d7b57fe5b9050602002013560046000898986818110610d9257fe5b9050602002016020810190610da79190613d74565b6001600160a01b031681526020810191909152604001600020549061332b565b60046000888885818110610dd757fe5b9050602002016020810190610dec9190613d74565b6001600160a01b03168152602081019190915260400160002055600101610d64565b507f6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f85858585604051610e4494939291906149d8565b60405180910390a1506001949350505050565b600554600160a81b900460ff1681565b600554600160a01b900460ff16610e7e6001613132565b90610e9c5760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b19169055610eb2613aac565b610eba613ae1565b6060610ec68e8e61145d565b925092509250610ed98360a00151611a91565b15610ee46005613132565b90610f025760405162461bcd60e51b81526004016107069190614b55565b5060a083015160009081526020819052604090819020805460ff191660011790558301516001600160a01b0316610f8d576040808401516001600160a01b03166000908152600460205220546080840151610f5c9161332b565b471015610f696007613132565b90610f875760405162461bcd60e51b81526004016107069190614b55565b5061109f565b6000610f9c8460400151612c3b565b905060098160ff161115610fc9576080840151610fc39060081960ff841601600a0a613391565b60808501525b6040808501516001600160a01b03166000908152600460205220546080850151610ff29161332b565b84604001516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611022919061485e565b60206040518083038186803b15801561103a57600080fd5b505afa15801561104e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110729190614672565b101561107e6007613132565b9061109c5760405162461bcd60e51b81526004016107069190614b55565b50505b6110f58e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d8d8d8d8d8d8d8d8d6133f3565b604080840151606085015160808601518551935163bda9b50960e01b8152309463bda9b5099461112f9490939092909188916004016148c9565b602060405180830381600087803b15801561114957600080fd5b505af1925050508015611179575060408051601f3d908101601f1916820190925261117691810190614672565b60015b611212573d8080156111a7576040519150601f19603f3d011682016040523d82523d6000602084013e6111ac565b606091505b506111c98460400151846020015186608001518760a001516134ee565b7fdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d4168460a0015160008360405161120193929190614b18565b60405180910390a150505050611375565b60408301516001600160a01b03166112415761123c83600001518460200151838760a001516134ee565b611370565b8251604080850151905163145e2a6b60e01b8152309263145e2a6b9261126b928690600401614872565b600060405180830381600087803b15801561128557600080fd5b505af1925050508015611296575060015b61132c573d8080156112c4576040519150601f19603f3d011682016040523d82523d6000602084013e6112c9565b606091505b506112e284600001518560200151848860a001516134ee565b7fdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d4168560a0015160018360405161131a93929190614b18565b60405180910390a15050505050611375565b7f9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb836000015184604001518360405161136793929190614872565b60405180910390a15b505050505b50506005805460ff60a01b1916600160a01b179055505050505050505050565b6000806000806020860151915060408601519250856040815181106113b657fe5b602001015160f81c60f81b60f81c601b019050600185828486604051600081526020016040526040516113ec9493929190614b37565b6020604051602081039080840390855afa15801561140e573d6000803e3d6000fd5b5050506020604051035193505050505b92915050565b60f181565b600081565b60046020526000908152604090205481565b600260209081526000928352604080842090915290825290205481565b611465613aac565b61146d613ae1565b606061012884101561147f6013613132565b9061149d5760405162461bcd60e51b81526004016107069190614b55565b506114a6613aac565b858560008181106114b357fe5b919091013560f81c825250858560018181106114cb57fe5b919091013560f81c6020830152506000868660028181106114e857fe5b845192013560f81c92505060ff16609e14801561150c5750816020015160ff166001145b801561151b575060ff81166001145b6115256013613132565b906115435760405162461bcd60e51b81526004016107069190614b55565b505061154d613ae1565b61155b60c36003888a614e51565b8101906115689190613d97565b6001600160a01b03908116604088810191909152918116875260a0880192909252608087019290925291821660608601529116908301526115ae61012860c3888a614e51565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505050602082015281816115f788610128818c614e51565b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250969e959d50919b50939950505050505050505050565b60f381565b60006001600160a01b03821661165c575047611714565b611665826135e4565b6116815760405162461bcd60e51b815260040161070690614c33565b6040516370a0823160e01b81526001600160a01b038316906370a08231906116ad90309060040161485e565b60206040518083038186803b1580156116c557600080fd5b505afa9250505080156116f5575060408051601f3d908101601f191682019092526116f291810190614672565b60015b6117115760405162461bcd60e51b815260040161070690614c9a565b90505b919050565b600554600160a01b900460ff166117306001613132565b9061174e5760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b1916905589516082111561176c6006613132565b9061178a5760405162461bcd60e51b81526004016107069190614b55565b50611793613aac565b61179c8b611b48565b805190915060ff1660f31480156117ba5750806020015160ff166001145b6117c46006613132565b906117e25760405162461bcd60e51b81526004016107069190614b55565b506117f08160a00151611a91565b156117fb6005613132565b906118195760405162461bcd60e51b81526004016107069190614b55565b5060a081015160009081526020819052604090819020805460ff191660011790558101516001600160a01b03166118a4576040808201516001600160a01b031660009081526004602052205460808201516118739161332b565b4710156118806007613132565b9061189e5760405162461bcd60e51b81526004016107069190614b55565b506119b6565b60006118b38260400151612c3b565b905060098160ff1611156118e05760808201516118da9060081960ff841601600a0a613391565b60808301525b6040808301516001600160a01b031660009081526004602052205460808301516119099161332b565b82604001516001600160a01b03166370a08231306040518263ffffffff1660e01b8152600401611939919061485e565b60206040518083038186803b15801561195157600080fd5b505afa158015611965573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119899190614672565b10156119956007613132565b906119b35760405162461bcd60e51b81526004016107069190614b55565b50505b6119c88b8b8b8b8b8b8b8b8b8b6133f3565b608081015160608201516001600160a01b0390811660009081526002602090815260408083208187015190941683529290522054611a059161332b565b60608201516001600160a01b03908116600090815260026020908152604080832081870180518616855290835281842095909555608086015194519093168252600490522054611a549161332b565b6040918201516001600160a01b031660009081526004602052919091205550506005805460ff60a01b1916600160a01b1790555050505050505050565b60008181526020819052604081205460ff1615611ab057506001611714565b6005546001600160a01b0316611ac857506000611714565b600554604051633a4e2fc360e11b81526001600160a01b039091169063749c5f8690611af8908590600401614b0f565b60206040518083038186803b158015611b1057600080fd5b505afa158015611b24573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117119190614145565b611b50613aac565b611b58613aac565b82600081518110611b6557fe5b016020015160f81c8152825183906001908110611b7e57fe5b0160209081015160f81c9082015260228301516042840151606285015160828601516001600160a01b039384166040860152929091166060840152608083015260a08201529050919050565b600554600160a01b900460ff16611be16001613132565b90611bff5760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b191690556b033b2e3c9fd0803ce8000000471115611c276002613132565b90611c455760405162461bcd60e51b81526004016107069190614b55565b50611c868383838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506135ea92505050565b7fd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab38516000868634611cb66006613657565b604051611cc795949392919061499f565b60405180910390a1611cd9600661365b565b50506005805460ff60a01b1916600160a01b179055505050565b600554600160a01b900460ff16611d0a6001613132565b90611d285760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b191690556000611da4611d48828d88888f613664565b8a8a8a8a604051602001611d60959493929190614d4c565b60408051601f198184030181526020601f8701819004810284018101909252858352919086908690819084018382808284376000920191909152506136ea92505050565b9050611db0818c61378e565b6001600160a01b038082166000908152600260209081526040808320938f16835292905220548a1115611de36008613132565b90611e015760405162461bcd60e51b81526004016107069190614b55565b506001600160a01b038b16600090815260046020526040902054611e25908b6138d0565b6001600160a01b03808d1660008181526004602090815260408083209590955592851681526002835283812091815291522054611e62908b6138d0565b60026000836001600160a01b03166001600160a01b0316815260200190815260200160002060008d6001600160a01b03166001600160a01b0316815260200190815260200160002081905550600034905060006001600160a01b03168c6001600160a01b03161415611edf57611ed8818c61332b565b905061201a565b6040516370a0823160e01b81528b906001600160a01b038e16906370a0823190611f0d90309060040161485e565b60206040518083038186803b158015611f2557600080fd5b505afa158015611f39573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f5d9190614672565b1015611f696007613132565b90611f875760405162461bcd60e51b81526004016107069190614b55565b5060405163a9059cbb60e01b81526001600160a01b038d169063a9059cbb90611fb6908c908f90600401614896565b600060405180830381600087803b158015611fd057600080fd5b505af1158015611fe4573d6000803e3d6000fd5b50505050611ff06132f7565b611ffa6004613132565b906120185760405162461bcd60e51b81526004016107069190614b55565b505b600061206c8b838b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508d613926565b90506120c98160026000866001600160a01b03166001600160a01b0316815260200190815260200160002060008e6001600160a01b03166001600160a01b031681526020019081526020016000205461332b90919063ffffffff16565b60026000856001600160a01b03166001600160a01b0316815260200190815260200160002060008d6001600160a01b03166001600160a01b031681526020019081526020016000208190555061214d81600460008e6001600160a01b03166001600160a01b031681526020019081526020016000205461332b90919063ffffffff16565b6001600160a01b03909b1660009081526004602052604090209a909a5550506005805460ff60a01b1916600160a01b17905550505050505050505050565b600360209081526000928352604080842090915290825290205460ff1681565b600554600160a01b900460ff1681565b600854156121c96012613132565b906121e75760405162461bcd60e51b81526004016107069190614b55565b5060026008556007546001600160a01b0316156122046011613132565b906122225760405162461bcd60e51b81526004016107069190614b55565b50600780546001600160a01b039384166001600160a01b03199182161790915560098054929093169116179055565b600554600160a01b900460ff166122686001613132565b906122865760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b191690558660006122a082612c3b565b90506000826001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016122d0919061485e565b60206040518083038186803b1580156122e857600080fd5b505afa1580156122fc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123209190614672565b90508089600960ff8516111561235b5760098460ff1603600a0a818161234257fe5b04905060098460ff1603600a0a838161235757fe5b0492505b670de0b6b3a7640000811115801561237b5750670de0b6b3a76400008311155b80156123985750670de0b6b3a7640000612395828561332b565b11155b6123a26003613132565b906123c05760405162461bcd60e51b81526004016107069190614b55565b50846001600160a01b03166323b872dd33308e6040518463ffffffff1660e01b81526004016123f193929190614872565b600060405180830381600087803b15801561240b57600080fd5b505af115801561241f573d6000803e3d6000fd5b5050505061242b6132f7565b6124356004613132565b906124535760405162461bcd60e51b81526004016107069190614b55565b508a612468836124628f611645565b906138d0565b14612473600a613132565b906124915760405162461bcd60e51b81526004016107069190614b55565b507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e8c8b8b846040516124c79493929190614969565b60405180910390a150506005805460ff60a01b1916600160a01b17905550505050505050505050565b600554600160a01b900460ff166125076001613132565b906125255760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b19169055604080516020601f84018190048102820181019092528281526125739185919085908590819084018382808284376000920191909152506135ea92505050565b86600061257f82612c3b565b90506000826001600160a01b03166370a08231306040518263ffffffff1660e01b81526004016125af919061485e565b60206040518083038186803b1580156125c757600080fd5b505afa1580156125db573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125ff9190614672565b90508089600960ff8516111561263a5760098460ff1603600a0a818161262157fe5b04905060098460ff1603600a0a838161263657fe5b0492505b670de0b6b3a7640000811115801561265a5750670de0b6b3a76400008311155b80156126775750670de0b6b3a7640000612674828561332b565b11155b6126816003613132565b9061269f5760405162461bcd60e51b81526004016107069190614b55565b50846001600160a01b03166323b872dd33308e6040518463ffffffff1660e01b81526004016126d093929190614872565b600060405180830381600087803b1580156126ea57600080fd5b505af11580156126fe573d6000803e3d6000fd5b5050505061270a6132f7565b6127146004613132565b906127325760405162461bcd60e51b81526004016107069190614b55565b508a612741836124628f611645565b1461274c600a613132565b9061276a5760405162461bcd60e51b81526004016107069190614b55565b507fd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab38518c8b8b8461279a6006613657565b6040516127ab95949392919061499f565b60405180910390a16127bd600661365b565b50506005805460ff60a01b1916600160a01b17905550505050505050505050565b609e81565b60003033146127f26014613132565b906128105760405162461bcd60e51b81526004016107069190614b55565b50600061281c83611645565b9050606060006001600160a01b0389166128375750856128cf565b60095460405163a9059cbb60e01b81526001600160a01b038b81169263a9059cbb9261286b92909116908b90600401614896565b600060405180830381600087803b15801561288557600080fd5b505af1158015612899573d6000803e3d6000fd5b505050506128a56132f7565b6128af6004613132565b906128cd5760405162461bcd60e51b81526004016107069190614b55565b505b600954604051631cff79cd60e01b81526001600160a01b0390911690631cff79cd908390612903908c908b9060040161490e565b6000604051808303818588803b15801561291c57600080fd5b505af1158015612930573d6000803e3d6000fd5b50505050506040513d6000823e601f3d908101601f191682016040526129599190810190614330565b9150815160401461296a6009613132565b906129885760405162461bcd60e51b81526004016107069190614b55565b50600080838060200190518101906129a09190613e06565b91509150866001600160a01b0316826001600160a01b03161480156129d05750806129ce866124628a611645565b145b6129da6009613132565b906129f85760405162461bcd60e51b81526004016107069190614b55565b509a9950505050505050505050565b600554600160a81b900460ff1615612a1f600f613132565b90612a3d5760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b1960ff60a81b196001600160a01b038088166001600160a01b03199094169390931716600160a81b1716600160a01b179091556007541615612a896011613132565b90612aa75760405162461bcd60e51b81526004016107069190614b55565b50600780546001600160a01b039384166001600160a01b0319918216179091556009805492909316911617905550565b6009546001600160a01b031681565b600554600160a01b900460ff16612afd6001613132565b90612b1b5760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b191690556b033b2e3c9fd0803ce8000000471115612b436002613132565b90612b615760405162461bcd60e51b81526004016107069190614b55565b507f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e6000868634604051612b989493929190614969565b60405180910390a150506005805460ff60a01b1916600160a01b179055505050565b6007546001600160a01b03161580612bdc57506007546001600160a01b031633145b8015612bf057506001600160a01b03811615155b612bfa6011613132565b90612c185760405162461bcd60e51b81526004016107069190614b55565b50600780546001600160a01b0319166001600160a01b0392909216919091179055565b6000612c46826135e4565b612c625760405162461bcd60e51b815260040161070690614c63565b6000829050806001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015612ca057600080fd5b505afa925050508015612cd0575060408051601f3d908101601f19168201909252612ccd91810190614672565b60015b612cec5760405162461bcd60e51b815260040161070690614b68565b91506117149050565b60085481565b600181565b60006020819052908152604090205460ff1681565b6007546001600160a01b031681565b60008181526001602052604081205460ff1615612d4357506001611714565b6005546001600160a01b0316612d5b57506000611714565b60055460405163392f5c1d60e21b81526001600160a01b039091169063e4bd707490611af8908590600401614b0f565b6005546000906001600160a01b031615801590612dce57506001600160a01b0380831660009081526003602090815260408083209387168352929052205460ff16155b15612e8a57600554604051637badcc6760e11b8152612e83916001600160a01b03169063f75b98ce90612e0790879087906004016148af565b60206040518083038186803b158015612e1f57600080fd5b505afa158015612e33573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612e579190614672565b6001600160a01b038085166000908152600260209081526040808320938916835292905220549061332b565b905061141e565b506001600160a01b0380821660009081526002602090815260408083209386168352929052205492915050565b6005546001600160a01b031681565b600554600160a01b900460ff16612edd6001613132565b90612efb5760405162461bcd60e51b81526004016107069190614b55565b506005805460ff60a01b19169055604080516020601f8401819004810282018101909252828152612f499185919085908590819084018382808284376000920191909152506135ea92505050565b6000612fb4612f5c60018c89898e613664565b8d8d604051602001612f7093929190614d95565b60408051601f198184030181526020601f8c018190048102840181019092528a835291908b908b90819084018382808284376000920191909152506136ea92505050565b9050612fc0818b61378e565b6001600160a01b038082166000908152600260209081526040808320938e1683529290522054891115612ff36008613132565b906130115760405162461bcd60e51b81526004016107069190614b55565b506001600160a01b038082166000908152600260209081526040808320938e1683529290522054613042908a6138d0565b6001600160a01b038083166000908152600260209081526040808320938f16835292815282822093909355600490925290205461307f908a6138d0565b6001600160a01b038b166000818152600460205260409020919091558990156130d35760006130ad8c612c3b565b905060098160ff1611156130d15760098160ff1603600a0a8b816130cd57fe5b0491505b505b7f2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e8b8e8e846040516131089493929190614969565b60405180910390a150506005805460ff60a01b1916600160a01b1790555050505050505050505050565b6060600082601481111561314257fe5b60408051600a808252818301909252919250906060908260208201818036833701905050905060005b60ff8416156131b9578151600a60ff959095168581049560018401939106916030830160f81b918591811061319c57fe5b60200101906001600160f81b031916908160001a9053505061316b565b6060816001016001600160401b03811180156131d457600080fd5b506040519080825280601f01601f1916602001820160405280156131ff576020820181803683370190505b50905060005b82811161325057838184038151811061321a57fe5b602001015160f81c60f81b82828151811061323157fe5b60200101906001600160f81b031916908160001a905350600101613205565b509695505050505050565b8047101561327b5760405162461bcd60e51b815260040161070690614bfc565b6000826001600160a01b0316826040516132949061485b565b60006040518083038185875af1925050503d80600081146132d1576040519150601f19603f3d011682016040523d82523d6000602084013e6132d6565b606091505b50509050806107bf5760405162461bcd60e51b815260040161070690614b9f565b6000803d801561330e576020811461331757613323565b60019150613323565b60206000803e60005191505b501515905090565b60008282018381108015906133405750828110155b6040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b815250906133895760405162461bcd60e51b81526004016107069190614b55565b509392505050565b60008282028315806133405750828482816133a857fe5b04146040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b815250906133895760405162461bcd60e51b81526004016107069190614b55565b60008a8a604051602001613408929190614839565b604051602081830303815290604052805190602001209050613428613a87565b6001600160a01b031663f65d21166001838d8d8d8d8d8d8d8d8d6040518c63ffffffff1660e01b81526004016134689b9a99989796959493929190614a62565b60206040518083038186803b15801561348057600080fd5b505afa158015613494573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134b89190614145565b6134c26006613132565b906134e05760405162461bcd60e51b81526004016107069190614b55565b505050505050505050505050565b816001600160a01b03851661353b576b033b2e3c9fd0803ce80000004711156135176002613132565b906135355760405162461bcd60e51b81526004016107069190614b55565b506135a1565b600061354686612c3b565b905060098160ff16111561356a5760098160ff1603600a0a828161356657fe5b0491505b670de0b6b3a76400008211156135806003613132565b9061359e5760405162461bcd60e51b81526004016107069190614b55565b50505b7eb45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d858583856040516135d59493929190614932565b60405180910390a15050505050565b3b151590565b6000613639826040518060400160405280336001600160a01b031681526020018681525060405160200161361e9190614dbb565b60405160208183030381529060405280519060200120611395565b6007549091506001600160a01b0380831691161461079f6010613132565b5490565b80546001019055565b61366c613b14565b613674613b14565b604051806080016040528088600781111561368b57fe5b8152602001876001600160a01b0316815260200186868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060200184905291505095945050505050565b815160208301206000906136fd81612d24565b156137086005613132565b906137265760405162461bcd60e51b81526004016107069190614b55565b5060006137338483611395565b90506001600160a01b038116151561374b6010613132565b906137695760405162461bcd60e51b81526004016107069190614b55565b506000918252600160208190526040909220805460ff19169092179091559392505050565b6005546001600160a01b0316158015906137ce57506001600160a01b0380831660009081526003602090815260408083209385168352929052205460ff16155b156138cc57600554604051637badcc6760e11b8152613883916001600160a01b03169063f75b98ce9061380790859087906004016148af565b60206040518083038186803b15801561381f57600080fd5b505afa158015613833573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906138579190614672565b6001600160a01b038085166000908152600260209081526040808320938716835292905220549061332b565b6001600160a01b038084166000818152600260209081526040808320948716808452948252808320959095559181526003825283812092815291905220805460ff191660011790555b5050565b6000828211156040518060400160405280601281526020017129b0b332a6b0ba341032bc31b2b83a34b7b760711b8152509061391f5760405162461bcd60e51b81526004016107069190614b55565b5050900390565b60008061393286611645565b90506001600160a01b03861661394f5761394c81346138d0565b90505b8447101561395d6007613132565b9061397b5760405162461bcd60e51b81526004016107069190614b55565b5060006060846001600160a01b03168787604051613999919061481d565b60006040518083038185875af1925050503d80600081146139d6576040519150601f19603f3d011682016040523d82523d6000602084013e6139db565b606091505b5091509150816139eb6004613132565b90613a095760405162461bcd60e51b81526004016107069190614b55565b5060008082806020019051810190613a219190613e06565b91509150896001600160a01b0316826001600160a01b0316148015613a51575080613a4f866124628d611645565b145b613a5b6009613132565b90613a795760405162461bcd60e51b81526004016107069190614b55565b509998505050505050505050565b7f62135fc083646fdb4e1a9d700e351b886a4a5a39da980650269edd1ade91ffd25490565b6040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a081019190915290565b604051806060016040528060006001600160a01b031681526020016060815260200160006001600160a01b031681525090565b6040805160808101909152806000815260200160006001600160a01b0316815260200160608152602001600081525090565b803561141e81614ea5565b60008083601f840112613b62578182fd5b5081356001600160401b03811115613b78578182fd5b6020830191508360208083028501011115613b9257600080fd5b9250929050565b600082601f830112613ba9578081fd5b8135613bbc613bb782614e0f565b614de9565b818152915060208083019084810181840286018201871015613bdd57600080fd5b60005b84811015613c05578135613bf381614ebd565b84529282019290820190600101613be0565b505050505092915050565b600082601f830112613c20578081fd5b8135613c2e613bb782614e0f565b818152915060208083019084810181840286018201871015613c4f57600080fd5b60005b84811015613c0557813584529282019290820190600101613c52565b600082601f830112613c7e578081fd5b8135613c8c613bb782614e0f565b818152915060208083019084810181840286018201871015613cad57600080fd5b6000805b85811015613cdb57823560ff81168114613cc9578283fd5b85529383019391830191600101613cb1565b50505050505092915050565b60008083601f840112613cf8578182fd5b5081356001600160401b03811115613d0e578182fd5b602083019150836020828501011115613b9257600080fd5b600082601f830112613d36578081fd5b8135613d44613bb782614e2e565b9150808252836020828501011115613d5b57600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215613d85578081fd5b8135613d9081614ea5565b9392505050565b60008060008060008060c08789031215613daf578182fd5b8635613dba81614ea5565b95506020870135613dca81614ea5565b945060408701359350606087013592506080870135613de881614ea5565b915060a0870135613df881614ea5565b809150509295509295509295565b60008060408385031215613e18578182fd5b8251613e2381614ea5565b6020939093015192949293505050565b60008060408385031215613e45578182fd5b8235613e5081614ea5565b91506020830135613e6081614ea5565b809150509250929050565b600080600060608486031215613e7f578081fd5b8335613e8a81614ea5565b92506020840135613e9a81614ea5565b929592945050506040919091013590565b600080600060608486031215613ebf578081fd5b8335613eca81614ea5565b92506020840135613eda81614ea5565b91506040840135613eea81614ea5565b809150509250925092565b600080600080600060a08688031215613f0c578283fd5b8535613f1781614ea5565b94506020860135613f2781614ea5565b93506040860135925060608601356001600160401b03811115613f48578182fd5b613f5488828901613d26565b9250506080860135613f6581614ea5565b809150509295509295909350565b60008060008060008060008060008060e08b8d031215613f91578788fd5b8a35613f9c81614ea5565b995060208b0135985060408b0135613fb381614ea5565b9750613fc28c60608d01613b46565b965060808b01356001600160401b0380821115613fdd578586fd5b613fe98e838f01613ce7565b909850965060a08d0135915080821115614001578586fd5b61400d8e838f01613ce7565b909650945060c08d0135915080821115614025578384fd5b506140328d828e01613ce7565b915080935050809150509295989b9194979a5092959850565b600080600080600080600060a0888a031215614065578081fd5b873561407081614ea5565b96506020880135955060408801356001600160401b0380821115614092578283fd5b61409e8b838c01613ce7565b909750955060608a0135945060808a01359150808211156140bd578283fd5b506140ca8a828b01613ce7565b989b979a50959850939692959293505050565b600080600080604085870312156140f2578182fd5b84356001600160401b0380821115614108578384fd5b61411488838901613b51565b9096509450602087013591508082111561412c578384fd5b5061413987828801613b51565b95989497509550505050565b600060208284031215614156578081fd5b8151613d9081614ebd565b600060208284031215614172578081fd5b5035919050565b6000806020838503121561418b578182fd5b82356001600160401b038111156141a0578283fd5b6141ac85828601613ce7565b90969095509350505050565b60008060008060008060008060008060006101408c8e0312156141d9578485fd5b6001600160401b03808d3511156141ee578586fd5b6141fb8e8e358f01613ce7565b909c509a5060208d0135995060408d0135811015614217578586fd5b6142278e60408f01358f01613c10565b98508060608e01351115614239578586fd5b6142498e60608f01358f01613b99565b975060808d0135965060a08d013595508060c08e01351115614269578182fd5b6142798e60c08f01358f01613c10565b94508060e08e0135111561428b578182fd5b61429b8e60e08f01358f01613c6e565b9350806101008e013511156142ae578182fd5b6142bf8e6101008f01358f01613c10565b9250806101208e013511156142d2578182fd5b506142e48d6101208e01358e01613c10565b90509295989b509295989b9093969950565b600060208284031215614307578081fd5b81356001600160401b0381111561431c578182fd5b61432884828501613d26565b949350505050565b600060208284031215614341578081fd5b81516001600160401b03811115614356578182fd5b8201601f81018413614366578182fd5b8051614374613bb782614e2e565b818152856020838501011115614388578384fd5b614399826020830160208601614e79565b95945050505050565b600080604083850312156143b4578182fd5b82356001600160401b038111156143c9578283fd5b6143d585828601613d26565b95602094909401359450505050565b6000806000806000806000806000806101408b8d031215614403578384fd5b8a356001600160401b0380821115614419578586fd5b6144258e838f01613d26565b9b5060208d01359a5060408d0135915080821115614441578586fd5b61444d8e838f01613c10565b995060608d0135915080821115614462578586fd5b61446e8e838f01613b99565b985060808d0135975060a08d0135965060c08d0135915080821115614491578586fd5b61449d8e838f01613c10565b955060e08d01359150808211156144b2578485fd5b6144be8e838f01613c6e565b94506101008d01359150808211156144d4578384fd5b6144e08e838f01613c10565b93506101208d01359150808211156144f6578283fd5b506145038d828e01613c10565b9150509295989b9194979a5092959850565b600080600080600080600080600080600060e08c8e031215614535578485fd5b6001600160401b03808d35111561454a578586fd5b6145578e8e358f01613ce7565b909c509a506145698e60208f01613b46565b995060408d013598508060608e01351115614582578586fd5b6145928e60608f01358f01613ce7565b909850965060808d01358110156145a7578586fd5b6145b78e60808f01358f01613ce7565b909650945060a08d0135935060c08d01358110156145d3578182fd5b506145e48d60c08e01358e01613ce7565b81935080925050509295989b509295989b9093969950565b600080600080600060608688031215614613578283fd5b85356001600160401b0380821115614629578485fd5b61463589838a01613ce7565b9097509550602088013594506040880135915080821115614654578283fd5b5061466188828901613ce7565b969995985093965092949392505050565b600060208284031215614683578081fd5b5051919050565b6000815180845260208085019450808401835b838110156146bb57815115158752958201959082019060010161469d565b509495945050505050565b6000815180845260208085019450808401835b838110156146bb578151875295820195908201906001016146d9565b6000815180845260208085019450808401835b838110156146bb57815160ff1687529582019590820190600101614708565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b60008151808452614769816020860160208601614e79565b601f01601f19169290920160200192915050565b60ff815116825260ff6020820151166020830152604081015160018060a01b03808216604085015280606084015116606085015250506080810151608083015260a081015160a08301525050565b60008151600881106147d957fe5b83526020828101516001600160a01b03169084015260408083015160809185018290529061480990850182614751565b606093840151949093019390935250919050565b6000825161482f818460208701614e79565b9190910192915050565b6000835161484b818460208801614e79565b9190910191825250602001919050565b90565b6001600160a01b0391909116815260200190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b0392831681529116602082015260400190565b600060018060a01b038088168352808716602084015285604084015260a060608401526148f960a0840186614751565b91508084166080840152509695505050505050565b6001600160a01b038316815260406020820181905260009061432890830184614751565b6001600160a01b038516815260806020820181905260009061495690830186614751565b6040830194909452506060015292915050565b6001600160a01b038516815260606020820181905260009061498e9083018587614727565b905082604083015295945050505050565b6001600160a01b03861681526080602082018190526000906149c49083018688614727565b604083019490945250606001529392505050565b6040808252810184905260008560608301825b87811015614a1b5760208335614a0081614ea5565b6001600160a01b0316835292830192909101906001016149eb565b5083810360208501528481526001600160fb1b03851115614a3a578283fd5b602085029150818660208301370160200190815295945050505050565b901515815260200190565b60006101608d151583528c60208401528b6040840152806060840152614a8a8184018c6146c6565b90508281036080840152614a9e818b61468a565b90508860a08401528760c084015282810360e0840152614abe81886146c6565b9050828103610100840152614ad381876146f5565b9050828103610120840152614ae881866146c6565b9050828103610140840152614afd81856146c6565b9e9d5050505050505050505050505050565b90815260200190565b6000848252836020830152606060408301526143996060830184614751565b93845260ff9290921660208401526040830152606082015260800190565b600060208252613d906020830184614751565b60208082526018908201527f67657420455243323020646563696d616c206661696c65640000000000000000604082015260600190565b6020808252603a908201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260408201527f6563697069656e74206d61792068617665207265766572746564000000000000606082015260800190565b6020808252601d908201527f416464726573733a20696e73756666696369656e742062616c616e6365000000604082015260600190565b60208082526016908201527518985b185b98d953d9881b9bdb8b58dbdb9d1c9858dd60521b604082015260600190565b60208082526018908201527f676574446563696d616c73206e6f6e2d636f6e74726163740000000000000000604082015260600190565b60208082526018908201527f6765742045524332302062616c616e6365206661696c65640000000000000000604082015260600190565b60c0810161141e828461477d565b6000610100614cee838761477d565b8060c084015260018060a01b038086511682850152602086015191506060610120850152614d20610160850183614751565b9150806040870151166101408501525082810360e0840152614d428185614751565b9695505050505050565b600060808252614d5f60808301886147cb565b6001600160a01b038781166020850152861660408401528281036060840152614d89818587614727565b98975050505050505050565b600060408252614da860408301866147cb565b8281036020840152614d42818587614727565b81516001600160a01b031681526020918201519181019190915260400190565b60ff91909116815260200190565b6040518181016001600160401b0381118282101715614e0757600080fd5b604052919050565b60006001600160401b03821115614e24578081fd5b5060209081020190565b60006001600160401b03821115614e43578081fd5b50601f01601f191660200190565b60008085851115614e60578182fd5b83861115614e6c578182fd5b5050820193919092039150565b60005b83811015614e94578181015183820152602001614e7c565b838111156107bd5750506000910152565b6001600160a01b0381168114614eba57600080fd5b50565b8015158114614eba57600080fdfea264697066735822122064e4b7b2a01ba46c1207efce5f4c6afbaf214090a7322a8c6e48153b4f5a6d4864736f6c634300060c0033",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Deprecated: Use VaultMetaData.Sigs instead.
// VaultFuncSigs maps the 4-byte function signature to its string representation.
var VaultFuncSigs = VaultMetaData.Sigs

// VaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VaultMetaData.Bin instead.
var VaultBin = VaultMetaData.Bin

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := VaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VaultBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNCALLREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_CALL_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNCALLREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNCALLREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNCALLREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0xbd835c42.
//
// Solidity: function BURN_CALL_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNCALLREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNCALLREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x568c04fd.
//
// Solidity: function BURN_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCaller) BURNTOCONTRACTREQUESTMETADATATYPE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "BURN_TO_CONTRACT_REQUEST_METADATA_TYPE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultSession) BURNTOCONTRACTREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNTOCONTRACTREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// BURNTOCONTRACTREQUESTMETADATATYPE is a free data retrieval call binding the contract method 0x6f2cbc48.
//
// Solidity: function BURN_TO_CONTRACT_REQUEST_METADATA_TYPE() view returns(uint8)
func (_Vault *VaultCallerSession) BURNTOCONTRACTREQUESTMETADATATYPE() (uint8, error) {
	return _Vault.Contract.BURNTOCONTRACTREQUESTMETADATATYPE(&_Vault.CallOpts)
}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultCaller) CURRENTNETWORKID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "CURRENT_NETWORK_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultSession) CURRENTNETWORKID() (uint8, error) {
	return _Vault.Contract.CURRENTNETWORKID(&_Vault.CallOpts)
}

// CURRENTNETWORKID is a free data retrieval call binding the contract method 0xd7200eb1.
//
// Solidity: function CURRENT_NETWORK_ID() view returns(uint8)
func (_Vault *VaultCallerSession) CURRENTNETWORKID() (uint8, error) {
	return _Vault.Contract.CURRENTNETWORKID(&_Vault.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultCaller) ETHTOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "ETH_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultSession) ETHTOKEN() (common.Address, error) {
	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
}

// ETHTOKEN is a free data retrieval call binding the contract method 0x58bc8337.
//
// Solidity: function ETH_TOKEN() view returns(address)
func (_Vault *VaultCallerSession) ETHTOKEN() (common.Address, error) {
	return _Vault.Contract.ETHTOKEN(&_Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "balanceOf", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address token) view returns(uint256)
func (_Vault *VaultCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Vault *VaultCaller) Executor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "executor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Vault *VaultSession) Executor() (common.Address, error) {
	return _Vault.Contract.Executor(&_Vault.CallOpts)
}

// Executor is a free data retrieval call binding the contract method 0xc34c08e5.
//
// Solidity: function executor() view returns(address)
func (_Vault *VaultCallerSession) Executor() (common.Address, error) {
	return _Vault.Contract.Executor(&_Vault.CallOpts)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultCaller) GetDecimals(opts *bind.CallOpts, token common.Address) (uint8, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDecimals", token)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultSession) GetDecimals(token common.Address) (uint8, error) {
	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
}

// GetDecimals is a free data retrieval call binding the contract method 0xcf54aaa0.
//
// Solidity: function getDecimals(address token) view returns(uint8)
func (_Vault *VaultCallerSession) GetDecimals(token common.Address) (uint8, error) {
	return _Vault.Contract.GetDecimals(&_Vault.CallOpts, token)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
func (_Vault *VaultCaller) GetDepositedBalance(opts *bind.CallOpts, token common.Address, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getDepositedBalance", token, owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
func (_Vault *VaultSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address token, address owner) view returns(uint256)
func (_Vault *VaultCallerSession) GetDepositedBalance(token common.Address, owner common.Address) (*big.Int, error) {
	return _Vault.Contract.GetDepositedBalance(&_Vault.CallOpts, token, owner)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_Vault *VaultCallerSession) IsInitialized() (bool, error) {
	return _Vault.Contract.IsInitialized(&_Vault.CallOpts)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultCaller) IsSigDataUsed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isSigDataUsed", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 hash) view returns(bool)
func (_Vault *VaultCallerSession) IsSigDataUsed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsSigDataUsed(&_Vault.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
func (_Vault *VaultCaller) IsWithdrawed(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "isWithdrawed", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
func (_Vault *VaultSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 hash) view returns(bool)
func (_Vault *VaultCallerSession) IsWithdrawed(hash [32]byte) (bool, error) {
	return _Vault.Contract.IsWithdrawed(&_Vault.CallOpts, hash)
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultCaller) Migration(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "migration", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
}

// Migration is a free data retrieval call binding the contract method 0x995fac11.
//
// Solidity: function migration(address , address ) view returns(bool)
func (_Vault *VaultCallerSession) Migration(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Vault.Contract.Migration(&_Vault.CallOpts, arg0, arg1)
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultCaller) NotEntered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "notEntered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultSession) NotEntered() (bool, error) {
	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
}

// NotEntered is a free data retrieval call binding the contract method 0xa3f5d8cc.
//
// Solidity: function notEntered() view returns(bool)
func (_Vault *VaultCallerSession) NotEntered() (bool, error) {
	return _Vault.Contract.NotEntered(&_Vault.CallOpts)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultCaller) ParseBurnInst(opts *bind.CallOpts, inst []byte) (VaultBurnInstData, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "parseBurnInst", inst)

	if err != nil {
		return *new(VaultBurnInstData), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultBurnInstData)).(*VaultBurnInstData)

	return out0, err

}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultSession) ParseBurnInst(inst []byte) (VaultBurnInstData, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// ParseBurnInst is a free data retrieval call binding the contract method 0x7e16e6e1.
//
// Solidity: function parseBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32))
func (_Vault *VaultCallerSession) ParseBurnInst(inst []byte) (VaultBurnInstData, error) {
	return _Vault.Contract.ParseBurnInst(&_Vault.CallOpts, inst)
}

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultCaller) ParseCalldataFromBurnInst(opts *bind.CallOpts, inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "parseCalldataFromBurnInst", inst)

	if err != nil {
		return *new(VaultBurnInstData), *new(VaultRedepositOptions), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(VaultBurnInstData)).(*VaultBurnInstData)
	out1 := *abi.ConvertType(out[1], new(VaultRedepositOptions)).(*VaultRedepositOptions)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultSession) ParseCalldataFromBurnInst(inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	return _Vault.Contract.ParseCalldataFromBurnInst(&_Vault.CallOpts, inst)
}

// ParseCalldataFromBurnInst is a free data retrieval call binding the contract method 0x66945b31.
//
// Solidity: function parseCalldataFromBurnInst(bytes inst) pure returns((uint8,uint8,address,address,uint256,bytes32), (address,bytes,address), bytes)
func (_Vault *VaultCallerSession) ParseCalldataFromBurnInst(inst []byte) (VaultBurnInstData, VaultRedepositOptions, []byte, error) {
	return _Vault.Contract.ParseCalldataFromBurnInst(&_Vault.CallOpts, inst)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultCaller) PrevVault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "prevVault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultSession) PrevVault() (common.Address, error) {
	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
}

// PrevVault is a free data retrieval call binding the contract method 0xfa84702e.
//
// Solidity: function prevVault() view returns(address)
func (_Vault *VaultCallerSession) PrevVault() (common.Address, error) {
	return _Vault.Contract.PrevVault(&_Vault.CallOpts)
}

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultCaller) Regulator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "regulator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultSession) Regulator() (common.Address, error) {
	return _Vault.Contract.Regulator(&_Vault.CallOpts)
}

// Regulator is a free data retrieval call binding the contract method 0xdd8fee14.
//
// Solidity: function regulator() view returns(address)
func (_Vault *VaultCallerSession) Regulator() (common.Address, error) {
	return _Vault.Contract.Regulator(&_Vault.CallOpts)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultCaller) SigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "sigDataUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
}

// SigDataUsed is a free data retrieval call binding the contract method 0x1ea1940e.
//
// Solidity: function sigDataUsed(bytes32 ) view returns(bool)
func (_Vault *VaultCallerSession) SigDataUsed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.SigDataUsed(&_Vault.CallOpts, arg0)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
func (_Vault *VaultCaller) SigToAddress(opts *bind.CallOpts, signData []byte, hash [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "sigToAddress", signData, hash)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
func (_Vault *VaultSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
}

// SigToAddress is a free data retrieval call binding the contract method 0x3fec6b40.
//
// Solidity: function sigToAddress(bytes signData, bytes32 hash) pure returns(address)
func (_Vault *VaultCallerSession) SigToAddress(signData []byte, hash [32]byte) (common.Address, error) {
	return _Vault.Contract.SigToAddress(&_Vault.CallOpts, signData, hash)
}

// StorageLayoutVersion is a free data retrieval call binding the contract method 0xd6a1fe3b.
//
// Solidity: function storageLayoutVersion() view returns(uint256)
func (_Vault *VaultCaller) StorageLayoutVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "storageLayoutVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StorageLayoutVersion is a free data retrieval call binding the contract method 0xd6a1fe3b.
//
// Solidity: function storageLayoutVersion() view returns(uint256)
func (_Vault *VaultSession) StorageLayoutVersion() (*big.Int, error) {
	return _Vault.Contract.StorageLayoutVersion(&_Vault.CallOpts)
}

// StorageLayoutVersion is a free data retrieval call binding the contract method 0xd6a1fe3b.
//
// Solidity: function storageLayoutVersion() view returns(uint256)
func (_Vault *VaultCallerSession) StorageLayoutVersion() (*big.Int, error) {
	return _Vault.Contract.StorageLayoutVersion(&_Vault.CallOpts)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultCaller) TotalDepositedToSCAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "totalDepositedToSCAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
}

// TotalDepositedToSCAmount is a free data retrieval call binding the contract method 0x6304541c.
//
// Solidity: function totalDepositedToSCAmount(address ) view returns(uint256)
func (_Vault *VaultCallerSession) TotalDepositedToSCAmount(arg0 common.Address) (*big.Int, error) {
	return _Vault.Contract.TotalDepositedToSCAmount(&_Vault.CallOpts, arg0)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
func (_Vault *VaultCaller) WithdrawRequests(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "withdrawRequests", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
func (_Vault *VaultSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
}

// WithdrawRequests is a free data retrieval call binding the contract method 0x65b5a00f.
//
// Solidity: function withdrawRequests(address , address ) view returns(uint256)
func (_Vault *VaultCallerSession) WithdrawRequests(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Vault.Contract.WithdrawRequests(&_Vault.CallOpts, arg0, arg1)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) view returns(bool)
func (_Vault *VaultCaller) Withdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "withdrawed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) view returns(bool)
func (_Vault *VaultSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// Withdrawed is a free data retrieval call binding the contract method 0xdca40d9e.
//
// Solidity: function withdrawed(bytes32 ) view returns(bool)
func (_Vault *VaultCallerSession) Withdrawed(arg0 [32]byte) (bool, error) {
	return _Vault.Contract.Withdrawed(&_Vault.CallOpts, arg0)
}

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultTransactor) CallExternal(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "_callExternal", token, to, amount, externalCalldata, redepositToken)
}

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultSession) CallExternal(token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.CallExternal(&_Vault.TransactOpts, token, to, amount, externalCalldata, redepositToken)
}

// CallExternal is a paid mutator transaction binding the contract method 0xbda9b509.
//
// Solidity: function _callExternal(address token, address to, uint256 amount, bytes externalCalldata, address redepositToken) returns(uint256)
func (_Vault *VaultTransactorSession) CallExternal(token common.Address, to common.Address, amount *big.Int, externalCalldata []byte, redepositToken common.Address) (*types.Transaction, error) {
	return _Vault.Contract.CallExternal(&_Vault.TransactOpts, token, to, amount, externalCalldata, redepositToken)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactor) TransferExternal(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "_transferExternal", token, to, amount)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultSession) TransferExternal(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferExternal(&_Vault.TransactOpts, token, to, amount)
}

// TransferExternal is a paid mutator transaction binding the contract method 0x145e2a6b.
//
// Solidity: function _transferExternal(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactorSession) TransferExternal(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferExternal(&_Vault.TransactOpts, token, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactor) Deposit(opts *bind.TransactOpts, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit", incognitoAddress, txId, signData)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultSession) Deposit(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// Deposit is a paid mutator transaction binding the contract method 0xc791d705.
//
// Solidity: function deposit(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactorSession) Deposit(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Deposit(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactor) DepositERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20", token, amount, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositERC20 is a paid mutator transaction binding the contract method 0xa807b5bb.
//
// Solidity: function depositERC20(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactorSession) DepositERC20(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0xb8237dbb.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactor) DepositERC20V2(opts *bind.TransactOpts, token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "depositERC20_V2", token, amount, incognitoAddress, txId, signData)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0xb8237dbb.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultSession) DepositERC20V2(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20V2(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositERC20V2 is a paid mutator transaction binding the contract method 0xb8237dbb.
//
// Solidity: function depositERC20_V2(address token, uint256 amount, string incognitoAddress, bytes32 txId, bytes signData) returns()
func (_Vault *VaultTransactorSession) DepositERC20V2(token common.Address, amount *big.Int, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositERC20V2(&_Vault.TransactOpts, token, amount, incognitoAddress, txId, signData)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x84b3ac03.
//
// Solidity: function deposit_V2(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactor) DepositV2(opts *bind.TransactOpts, incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "deposit_V2", incognitoAddress, txId, signData)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x84b3ac03.
//
// Solidity: function deposit_V2(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultSession) DepositV2(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositV2(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// DepositV2 is a paid mutator transaction binding the contract method 0x84b3ac03.
//
// Solidity: function deposit_V2(string incognitoAddress, bytes32 txId, bytes signData) payable returns()
func (_Vault *VaultTransactorSession) DepositV2(incognitoAddress string, txId [32]byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.DepositV2(&_Vault.TransactOpts, incognitoAddress, txId, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultTransactor) Execute(opts *bind.TransactOpts, token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "execute", token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// Execute is a paid mutator transaction binding the contract method 0x8588ccd6.
//
// Solidity: function execute(address token, uint256 amount, address recipientToken, address exchangeAddress, bytes callData, bytes timestamp, bytes signData) payable returns()
func (_Vault *VaultTransactorSession) Execute(token common.Address, amount *big.Int, recipientToken common.Address, exchangeAddress common.Address, callData []byte, timestamp []byte, signData []byte) (*types.Transaction, error) {
	return _Vault.Contract.Execute(&_Vault.TransactOpts, token, amount, recipientToken, exchangeAddress, callData, timestamp, signData)
}

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) ExecuteWithBurnProof(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "executeWithBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) ExecuteWithBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteWithBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// ExecuteWithBurnProof is a paid mutator transaction binding the contract method 0x3ed1b376.
//
// Solidity: function executeWithBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) ExecuteWithBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.ExecuteWithBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _prevVault, address _regulator, address _executor) returns()
func (_Vault *VaultTransactor) Initialize(opts *bind.TransactOpts, _prevVault common.Address, _regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "initialize", _prevVault, _regulator, _executor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _prevVault, address _regulator, address _executor) returns()
func (_Vault *VaultSession) Initialize(_prevVault common.Address, _regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _prevVault, _regulator, _executor)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _prevVault, address _regulator, address _executor) returns()
func (_Vault *VaultTransactorSession) Initialize(_prevVault common.Address, _regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.Initialize(&_Vault.TransactOpts, _prevVault, _regulator, _executor)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultTransactor) RequestWithdraw(opts *bind.TransactOpts, incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "requestWithdraw", incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// RequestWithdraw is a paid mutator transaction binding the contract method 0xfee8efda.
//
// Solidity: function requestWithdraw(string incognitoAddress, address token, uint256 amount, bytes signData, bytes timestamp, bytes32 txId, bytes regulatorSig) returns()
func (_Vault *VaultTransactorSession) RequestWithdraw(incognitoAddress string, token common.Address, amount *big.Int, signData []byte, timestamp []byte, txId [32]byte, regulatorSig []byte) (*types.Transaction, error) {
	return _Vault.Contract.RequestWithdraw(&_Vault.TransactOpts, incognitoAddress, token, amount, signData, timestamp, txId, regulatorSig)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultTransactor) SetRegulator(opts *bind.TransactOpts, _regulator common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setRegulator", _regulator)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultSession) SetRegulator(_regulator common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetRegulator(&_Vault.TransactOpts, _regulator)
}

// SetRegulator is a paid mutator transaction binding the contract method 0xcde0a4f8.
//
// Solidity: function setRegulator(address _regulator) returns()
func (_Vault *VaultTransactorSession) SetRegulator(_regulator common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetRegulator(&_Vault.TransactOpts, _regulator)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) SubmitBurnProof(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "submitBurnProof", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) SubmitBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// SubmitBurnProof is a paid mutator transaction binding the contract method 0x73bf9651.
//
// Solidity: function submitBurnProof(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) SubmitBurnProof(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.SubmitBurnProof(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultTransactor) UpdateAssets(opts *bind.TransactOpts, assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "updateAssets", assets, amounts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultSession) UpdateAssets(assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpdateAssets(&_Vault.TransactOpts, assets, amounts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] assets, uint256[] amounts) returns(bool)
func (_Vault *VaultTransactorSession) UpdateAssets(assets []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Vault.Contract.UpdateAssets(&_Vault.TransactOpts, assets, amounts)
}

// UpgradeVaultStorage is a paid mutator transaction binding the contract method 0xa73b1532.
//
// Solidity: function upgradeVaultStorage(address _regulator, address _executor) returns()
func (_Vault *VaultTransactor) UpgradeVaultStorage(opts *bind.TransactOpts, _regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "upgradeVaultStorage", _regulator, _executor)
}

// UpgradeVaultStorage is a paid mutator transaction binding the contract method 0xa73b1532.
//
// Solidity: function upgradeVaultStorage(address _regulator, address _executor) returns()
func (_Vault *VaultSession) UpgradeVaultStorage(_regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpgradeVaultStorage(&_Vault.TransactOpts, _regulator, _executor)
}

// UpgradeVaultStorage is a paid mutator transaction binding the contract method 0xa73b1532.
//
// Solidity: function upgradeVaultStorage(address _regulator, address _executor) returns()
func (_Vault *VaultTransactorSession) UpgradeVaultStorage(_regulator common.Address, _executor common.Address) (*types.Transaction, error) {
	return _Vault.Contract.UpgradeVaultStorage(&_Vault.TransactOpts, _regulator, _executor)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactor) Withdraw(opts *bind.TransactOpts, inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "withdraw", inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultSession) Withdraw(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Withdraw is a paid mutator transaction binding the contract method 0x1beb7de2.
//
// Solidity: function withdraw(bytes inst, uint256 heights, bytes32[] instPaths, bool[] instPathIsLefts, bytes32 instRoots, bytes32 blkData, uint256[] sigIdxs, uint8[] sigVs, bytes32[] sigRs, bytes32[] sigSs) returns()
func (_Vault *VaultTransactorSession) Withdraw(inst []byte, heights *big.Int, instPaths [][32]byte, instPathIsLefts []bool, instRoots [32]byte, blkData [32]byte, sigIdxs []*big.Int, sigVs []uint8, sigRs [][32]byte, sigSs [][32]byte) (*types.Transaction, error) {
	return _Vault.Contract.Withdraw(&_Vault.TransactOpts, inst, heights, instPaths, instPathIsLefts, instRoots, blkData, sigIdxs, sigVs, sigRs, sigSs)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// VaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Vault contract.
type VaultDepositIterator struct {
	Event *VaultDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDeposit represents a Deposit event raised by the Vault contract.
type VaultDeposit struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) FilterDeposit(opts *bind.FilterOpts) (*VaultDepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &VaultDepositIterator{contract: _Vault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *VaultDeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDeposit)
				if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e.
//
// Solidity: event Deposit(address token, string incognitoAddress, uint256 amount)
func (_Vault *VaultFilterer) ParseDeposit(log types.Log) (*VaultDeposit, error) {
	event := new(VaultDeposit)
	if err := _Vault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultDepositV2Iterator is returned from FilterDepositV2 and is used to iterate over the raw logs and unpacked data for DepositV2 events raised by the Vault contract.
type VaultDepositV2Iterator struct {
	Event *VaultDepositV2 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultDepositV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultDepositV2)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultDepositV2)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultDepositV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultDepositV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultDepositV2 represents a DepositV2 event raised by the Vault contract.
type VaultDepositV2 struct {
	Token            common.Address
	IncognitoAddress string
	Amount           *big.Int
	DepositID        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDepositV2 is a free log retrieval operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) FilterDepositV2(opts *bind.FilterOpts) (*VaultDepositV2Iterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "DepositV2")
	if err != nil {
		return nil, err
	}
	return &VaultDepositV2Iterator{contract: _Vault.contract, event: "DepositV2", logs: logs, sub: sub}, nil
}

// WatchDepositV2 is a free log subscription operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) WatchDepositV2(opts *bind.WatchOpts, sink chan<- *VaultDepositV2) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "DepositV2")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultDepositV2)
				if err := _Vault.contract.UnpackLog(event, "DepositV2", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDepositV2 is a log parse operation binding the contract event 0xd30df8040a1092415b49422a02dbd8cdd5915a596abcba02cd0f65dd86ab3851.
//
// Solidity: event DepositV2(address token, string incognitoAddress, uint256 amount, uint256 depositID)
func (_Vault *VaultFilterer) ParseDepositV2(log types.Log) (*VaultDepositV2, error) {
	event := new(VaultDepositV2)
	if err := _Vault.contract.UnpackLog(event, "DepositV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultExecuteFnLogIterator is returned from FilterExecuteFnLog and is used to iterate over the raw logs and unpacked data for ExecuteFnLog events raised by the Vault contract.
type VaultExecuteFnLogIterator struct {
	Event *VaultExecuteFnLog // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultExecuteFnLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultExecuteFnLog)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultExecuteFnLog)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultExecuteFnLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultExecuteFnLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultExecuteFnLog represents a ExecuteFnLog event raised by the Vault contract.
type VaultExecuteFnLog struct {
	Id        [32]byte
	PhaseID   *big.Int
	ErrorData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExecuteFnLog is a free log retrieval operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) FilterExecuteFnLog(opts *bind.FilterOpts) (*VaultExecuteFnLogIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "ExecuteFnLog")
	if err != nil {
		return nil, err
	}
	return &VaultExecuteFnLogIterator{contract: _Vault.contract, event: "ExecuteFnLog", logs: logs, sub: sub}, nil
}

// WatchExecuteFnLog is a free log subscription operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) WatchExecuteFnLog(opts *bind.WatchOpts, sink chan<- *VaultExecuteFnLog) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "ExecuteFnLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultExecuteFnLog)
				if err := _Vault.contract.UnpackLog(event, "ExecuteFnLog", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecuteFnLog is a log parse operation binding the contract event 0xdbbb883f24557adf486292429863dcfd4ac5d4db168ae94921da8e3d9a95d416.
//
// Solidity: event ExecuteFnLog(bytes32 id, uint256 phaseID, bytes errorData)
func (_Vault *VaultFilterer) ParseExecuteFnLog(log types.Log) (*VaultExecuteFnLog, error) {
	event := new(VaultExecuteFnLog)
	if err := _Vault.contract.UnpackLog(event, "ExecuteFnLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultRedepositIterator is returned from FilterRedeposit and is used to iterate over the raw logs and unpacked data for Redeposit events raised by the Vault contract.
type VaultRedepositIterator struct {
	Event *VaultRedeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultRedepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultRedeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultRedeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultRedepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultRedepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultRedeposit represents a Redeposit event raised by the Vault contract.
type VaultRedeposit struct {
	Token               common.Address
	RedepositIncAddress []byte
	Amount              *big.Int
	Itx                 [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterRedeposit is a free log retrieval operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) FilterRedeposit(opts *bind.FilterOpts) (*VaultRedepositIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Redeposit")
	if err != nil {
		return nil, err
	}
	return &VaultRedepositIterator{contract: _Vault.contract, event: "Redeposit", logs: logs, sub: sub}, nil
}

// WatchRedeposit is a free log subscription operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) WatchRedeposit(opts *bind.WatchOpts, sink chan<- *VaultRedeposit) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Redeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultRedeposit)
				if err := _Vault.contract.UnpackLog(event, "Redeposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRedeposit is a log parse operation binding the contract event 0x00b45d95b5117447e2fafe7f34def913ff3ba220e4b8688acf37ae2328af7a3d.
//
// Solidity: event Redeposit(address token, bytes redepositIncAddress, uint256 amount, bytes32 itx)
func (_Vault *VaultFilterer) ParseRedeposit(log types.Log) (*VaultRedeposit, error) {
	event := new(VaultRedeposit)
	if err := _Vault.contract.UnpackLog(event, "Redeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdateIncognitoProxyIterator is returned from FilterUpdateIncognitoProxy and is used to iterate over the raw logs and unpacked data for UpdateIncognitoProxy events raised by the Vault contract.
type VaultUpdateIncognitoProxyIterator struct {
	Event *VaultUpdateIncognitoProxy // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultUpdateIncognitoProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdateIncognitoProxy)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultUpdateIncognitoProxy)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultUpdateIncognitoProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdateIncognitoProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdateIncognitoProxy represents a UpdateIncognitoProxy event raised by the Vault contract.
type VaultUpdateIncognitoProxy struct {
	NewIncognitoProxy common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterUpdateIncognitoProxy is a free log retrieval operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) FilterUpdateIncognitoProxy(opts *bind.FilterOpts) (*VaultUpdateIncognitoProxyIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateIncognitoProxy")
	if err != nil {
		return nil, err
	}
	return &VaultUpdateIncognitoProxyIterator{contract: _Vault.contract, event: "UpdateIncognitoProxy", logs: logs, sub: sub}, nil
}

// WatchUpdateIncognitoProxy is a free log subscription operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) WatchUpdateIncognitoProxy(opts *bind.WatchOpts, sink chan<- *VaultUpdateIncognitoProxy) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateIncognitoProxy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdateIncognitoProxy)
				if err := _Vault.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateIncognitoProxy is a log parse operation binding the contract event 0x204252dfe190ad6ef63db40a490f048b39f661de74628408f13cd0bb2d4c3446.
//
// Solidity: event UpdateIncognitoProxy(address newIncognitoProxy)
func (_Vault *VaultFilterer) ParseUpdateIncognitoProxy(log types.Log) (*VaultUpdateIncognitoProxy, error) {
	event := new(VaultUpdateIncognitoProxy)
	if err := _Vault.contract.UnpackLog(event, "UpdateIncognitoProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUpdateTokenTotalIterator is returned from FilterUpdateTokenTotal and is used to iterate over the raw logs and unpacked data for UpdateTokenTotal events raised by the Vault contract.
type VaultUpdateTokenTotalIterator struct {
	Event *VaultUpdateTokenTotal // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultUpdateTokenTotalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUpdateTokenTotal)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultUpdateTokenTotal)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultUpdateTokenTotalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUpdateTokenTotalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUpdateTokenTotal represents a UpdateTokenTotal event raised by the Vault contract.
type VaultUpdateTokenTotal struct {
	Assets  []common.Address
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUpdateTokenTotal is a free log retrieval operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) FilterUpdateTokenTotal(opts *bind.FilterOpts) (*VaultUpdateTokenTotalIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "UpdateTokenTotal")
	if err != nil {
		return nil, err
	}
	return &VaultUpdateTokenTotalIterator{contract: _Vault.contract, event: "UpdateTokenTotal", logs: logs, sub: sub}, nil
}

// WatchUpdateTokenTotal is a free log subscription operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) WatchUpdateTokenTotal(opts *bind.WatchOpts, sink chan<- *VaultUpdateTokenTotal) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "UpdateTokenTotal")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUpdateTokenTotal)
				if err := _Vault.contract.UnpackLog(event, "UpdateTokenTotal", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpdateTokenTotal is a log parse operation binding the contract event 0x6a7fbbcddfd518bb8c56b28ac6c7acb0f7ca093ed232eb3306e53d14e469895f.
//
// Solidity: event UpdateTokenTotal(address[] assets, uint256[] amounts)
func (_Vault *VaultFilterer) ParseUpdateTokenTotal(log types.Log) (*VaultUpdateTokenTotal, error) {
	event := new(VaultUpdateTokenTotal)
	if err := _Vault.contract.UnpackLog(event, "UpdateTokenTotal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Vault contract.
type VaultWithdrawIterator struct {
	Event *VaultWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *VaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(VaultWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *VaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultWithdraw represents a Withdraw event raised by the Vault contract.
type VaultWithdraw struct {
	Token  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) FilterWithdraw(opts *bind.FilterOpts) (*VaultWithdrawIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &VaultWithdrawIterator{contract: _Vault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *VaultWithdraw) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultWithdraw)
				if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x9b1bfa7fa9ee420a16e124f794c35ac9f90472acc99140eb2f6447c714cad8eb.
//
// Solidity: event Withdraw(address token, address to, uint256 amount)
func (_Vault *VaultFilterer) ParseWithdraw(log types.Log) (*VaultWithdraw, error) {
	event := new(VaultWithdraw)
	if err := _Vault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WithdrawableMetaData contains all meta data concerning the Withdrawable contract.
var WithdrawableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"getDepositedBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isSigDataUsed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"isWithdrawed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"updateAssets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"f75b98ce": "getDepositedBalance(address,address)",
		"e4bd7074": "isSigDataUsed(bytes32)",
		"749c5f86": "isWithdrawed(bytes32)",
		"5c975abb": "paused()",
		"1ed4276d": "updateAssets(address[],uint256[])",
	},
}

// WithdrawableABI is the input ABI used to generate the binding from.
// Deprecated: Use WithdrawableMetaData.ABI instead.
var WithdrawableABI = WithdrawableMetaData.ABI

// Deprecated: Use WithdrawableMetaData.Sigs instead.
// WithdrawableFuncSigs maps the 4-byte function signature to its string representation.
var WithdrawableFuncSigs = WithdrawableMetaData.Sigs

// Withdrawable is an auto generated Go binding around an Ethereum contract.
type Withdrawable struct {
	WithdrawableCaller     // Read-only binding to the contract
	WithdrawableTransactor // Write-only binding to the contract
	WithdrawableFilterer   // Log filterer for contract events
}

// WithdrawableCaller is an auto generated read-only Go binding around an Ethereum contract.
type WithdrawableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WithdrawableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WithdrawableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WithdrawableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WithdrawableSession struct {
	Contract     *Withdrawable     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WithdrawableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WithdrawableCallerSession struct {
	Contract *WithdrawableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WithdrawableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WithdrawableTransactorSession struct {
	Contract     *WithdrawableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WithdrawableRaw is an auto generated low-level Go binding around an Ethereum contract.
type WithdrawableRaw struct {
	Contract *Withdrawable // Generic contract binding to access the raw methods on
}

// WithdrawableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WithdrawableCallerRaw struct {
	Contract *WithdrawableCaller // Generic read-only contract binding to access the raw methods on
}

// WithdrawableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WithdrawableTransactorRaw struct {
	Contract *WithdrawableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWithdrawable creates a new instance of Withdrawable, bound to a specific deployed contract.
func NewWithdrawable(address common.Address, backend bind.ContractBackend) (*Withdrawable, error) {
	contract, err := bindWithdrawable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Withdrawable{WithdrawableCaller: WithdrawableCaller{contract: contract}, WithdrawableTransactor: WithdrawableTransactor{contract: contract}, WithdrawableFilterer: WithdrawableFilterer{contract: contract}}, nil
}

// NewWithdrawableCaller creates a new read-only instance of Withdrawable, bound to a specific deployed contract.
func NewWithdrawableCaller(address common.Address, caller bind.ContractCaller) (*WithdrawableCaller, error) {
	contract, err := bindWithdrawable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawableCaller{contract: contract}, nil
}

// NewWithdrawableTransactor creates a new write-only instance of Withdrawable, bound to a specific deployed contract.
func NewWithdrawableTransactor(address common.Address, transactor bind.ContractTransactor) (*WithdrawableTransactor, error) {
	contract, err := bindWithdrawable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WithdrawableTransactor{contract: contract}, nil
}

// NewWithdrawableFilterer creates a new log filterer instance of Withdrawable, bound to a specific deployed contract.
func NewWithdrawableFilterer(address common.Address, filterer bind.ContractFilterer) (*WithdrawableFilterer, error) {
	contract, err := bindWithdrawable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WithdrawableFilterer{contract: contract}, nil
}

// bindWithdrawable binds a generic wrapper to an already deployed contract.
func bindWithdrawable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WithdrawableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawable *WithdrawableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawable.Contract.WithdrawableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawable *WithdrawableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawable.Contract.WithdrawableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawable *WithdrawableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawable.Contract.WithdrawableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Withdrawable *WithdrawableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Withdrawable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Withdrawable *WithdrawableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Withdrawable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Withdrawable *WithdrawableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Withdrawable.Contract.contract.Transact(opts, method, params...)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address , address ) view returns(uint256)
func (_Withdrawable *WithdrawableCaller) GetDepositedBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Withdrawable.contract.Call(opts, &out, "getDepositedBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address , address ) view returns(uint256)
func (_Withdrawable *WithdrawableSession) GetDepositedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Withdrawable.Contract.GetDepositedBalance(&_Withdrawable.CallOpts, arg0, arg1)
}

// GetDepositedBalance is a free data retrieval call binding the contract method 0xf75b98ce.
//
// Solidity: function getDepositedBalance(address , address ) view returns(uint256)
func (_Withdrawable *WithdrawableCallerSession) GetDepositedBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Withdrawable.Contract.GetDepositedBalance(&_Withdrawable.CallOpts, arg0, arg1)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableCaller) IsSigDataUsed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Withdrawable.contract.Call(opts, &out, "isSigDataUsed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableSession) IsSigDataUsed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsSigDataUsed(&_Withdrawable.CallOpts, arg0)
}

// IsSigDataUsed is a free data retrieval call binding the contract method 0xe4bd7074.
//
// Solidity: function isSigDataUsed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableCallerSession) IsSigDataUsed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsSigDataUsed(&_Withdrawable.CallOpts, arg0)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableCaller) IsWithdrawed(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Withdrawable.contract.Call(opts, &out, "isWithdrawed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableSession) IsWithdrawed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsWithdrawed(&_Withdrawable.CallOpts, arg0)
}

// IsWithdrawed is a free data retrieval call binding the contract method 0x749c5f86.
//
// Solidity: function isWithdrawed(bytes32 ) view returns(bool)
func (_Withdrawable *WithdrawableCallerSession) IsWithdrawed(arg0 [32]byte) (bool, error) {
	return _Withdrawable.Contract.IsWithdrawed(&_Withdrawable.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Withdrawable *WithdrawableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Withdrawable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Withdrawable *WithdrawableSession) Paused() (bool, error) {
	return _Withdrawable.Contract.Paused(&_Withdrawable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Withdrawable *WithdrawableCallerSession) Paused() (bool, error) {
	return _Withdrawable.Contract.Paused(&_Withdrawable.CallOpts)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
func (_Withdrawable *WithdrawableTransactor) UpdateAssets(opts *bind.TransactOpts, arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
	return _Withdrawable.contract.Transact(opts, "updateAssets", arg0, arg1)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
func (_Withdrawable *WithdrawableSession) UpdateAssets(arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
	return _Withdrawable.Contract.UpdateAssets(&_Withdrawable.TransactOpts, arg0, arg1)
}

// UpdateAssets is a paid mutator transaction binding the contract method 0x1ed4276d.
//
// Solidity: function updateAssets(address[] , uint256[] ) returns(bool)
func (_Withdrawable *WithdrawableTransactorSession) UpdateAssets(arg0 []common.Address, arg1 []*big.Int) (*types.Transaction, error) {
	return _Withdrawable.Contract.UpdateAssets(&_Withdrawable.TransactOpts, arg0, arg1)
}