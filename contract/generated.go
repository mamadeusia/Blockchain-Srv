// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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
	_ = abi.ConvertType
)

// IMessageDomainSeparator is an auto generated low-level Go binding around an user-defined struct.
type IMessageDomainSeparator struct {
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
}

// IMessageMessage is an auto generated low-level Go binding around an user-defined struct.
type IMessageMessage struct {
	ProposalId *big.Int
	Support    uint8
	Reason     string
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212200a2e40664d18f1a48c6736a5c19099f1af1478cbb2bb6d8ee49967e514a23ce064736f6c63430008130033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContextUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ContextUpgradeable contract.
type ContextUpgradeableInitializedIterator struct {
	Event *ContextUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ContextUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContextUpgradeableInitialized)
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
		it.Event = new(ContextUpgradeableInitialized)
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
func (it *ContextUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContextUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContextUpgradeableInitialized represents a Initialized event raised by the ContextUpgradeable contract.
type ContextUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ContextUpgradeableInitializedIterator, error) {

	logs, sub, err := _ContextUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableInitializedIterator{contract: _ContextUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ContextUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ContextUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContextUpgradeableInitialized)
				if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ContextUpgradeable *ContextUpgradeableFilterer) ParseInitialized(log types.Log) (*ContextUpgradeableInitialized, error) {
	event := new(ContextUpgradeableInitialized)
	if err := _ContextUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableMetaData contains all meta data concerning the ERC1967UpgradeUpgradeable contract.
var ERC1967UpgradeUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]",
}

// ERC1967UpgradeUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967UpgradeUpgradeableMetaData.ABI instead.
var ERC1967UpgradeUpgradeableABI = ERC1967UpgradeUpgradeableMetaData.ABI

// ERC1967UpgradeUpgradeable is an auto generated Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeable struct {
	ERC1967UpgradeUpgradeableCaller     // Read-only binding to the contract
	ERC1967UpgradeUpgradeableTransactor // Write-only binding to the contract
	ERC1967UpgradeUpgradeableFilterer   // Log filterer for contract events
}

// ERC1967UpgradeUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967UpgradeUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967UpgradeUpgradeableSession struct {
	Contract     *ERC1967UpgradeUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC1967UpgradeUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967UpgradeUpgradeableCallerSession struct {
	Contract *ERC1967UpgradeUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ERC1967UpgradeUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967UpgradeUpgradeableTransactorSession struct {
	Contract     *ERC1967UpgradeUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ERC1967UpgradeUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableRaw struct {
	Contract *ERC1967UpgradeUpgradeable // Generic contract binding to access the raw methods on
}

// ERC1967UpgradeUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableCallerRaw struct {
	Contract *ERC1967UpgradeUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967UpgradeUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableTransactorRaw struct {
	Contract *ERC1967UpgradeUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967UpgradeUpgradeable creates a new instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC1967UpgradeUpgradeable, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeable{ERC1967UpgradeUpgradeableCaller: ERC1967UpgradeUpgradeableCaller{contract: contract}, ERC1967UpgradeUpgradeableTransactor: ERC1967UpgradeUpgradeableTransactor{contract: contract}, ERC1967UpgradeUpgradeableFilterer: ERC1967UpgradeUpgradeableFilterer{contract: contract}}, nil
}

// NewERC1967UpgradeUpgradeableCaller creates a new read-only instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC1967UpgradeUpgradeableCaller, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableCaller{contract: contract}, nil
}

// NewERC1967UpgradeUpgradeableTransactor creates a new write-only instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967UpgradeUpgradeableTransactor, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableTransactor{contract: contract}, nil
}

// NewERC1967UpgradeUpgradeableFilterer creates a new log filterer instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967UpgradeUpgradeableFilterer, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableFilterer{contract: contract}, nil
}

// bindERC1967UpgradeUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC1967UpgradeUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC1967UpgradeUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC1967UpgradeUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableAdminChangedIterator struct {
	Event *ERC1967UpgradeUpgradeableAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableAdminChanged)
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
		it.Event = new(ERC1967UpgradeUpgradeableAdminChanged)
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
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableAdminChanged represents a AdminChanged event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967UpgradeUpgradeableAdminChangedIterator, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableAdminChangedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableAdminChanged)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseAdminChanged(log types.Log) (*ERC1967UpgradeUpgradeableAdminChanged, error) {
	event := new(ERC1967UpgradeUpgradeableAdminChanged)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableBeaconUpgradedIterator struct {
	Event *ERC1967UpgradeUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableBeaconUpgraded)
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
		it.Event = new(ERC1967UpgradeUpgradeableBeaconUpgraded)
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
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967UpgradeUpgradeableBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableBeaconUpgradedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableBeaconUpgraded)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967UpgradeUpgradeableBeaconUpgraded, error) {
	event := new(ERC1967UpgradeUpgradeableBeaconUpgraded)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableInitializedIterator struct {
	Event *ERC1967UpgradeUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableInitialized)
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
		it.Event = new(ERC1967UpgradeUpgradeableInitialized)
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
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableInitialized represents a Initialized event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*ERC1967UpgradeUpgradeableInitializedIterator, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableInitializedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableInitialized)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseInitialized(log types.Log) (*ERC1967UpgradeUpgradeableInitialized, error) {
	event := new(ERC1967UpgradeUpgradeableInitialized)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableUpgradedIterator struct {
	Event *ERC1967UpgradeUpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableUpgraded)
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
		it.Event = new(ERC1967UpgradeUpgradeableUpgraded)
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
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableUpgraded represents a Upgraded event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967UpgradeUpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableUpgradedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableUpgraded)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseUpgraded(log types.Log) (*ERC1967UpgradeUpgradeableUpgraded, error) {
	event := new(ERC1967UpgradeUpgradeableUpgraded)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EnumerableMapMetaData contains all meta data concerning the EnumerableMap contract.
var EnumerableMapMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a9bb04e456ea7976d112b289525f5fce0ff17ce258d270290d68b90b0944bed764736f6c63430008130033",
}

// EnumerableMapABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableMapMetaData.ABI instead.
var EnumerableMapABI = EnumerableMapMetaData.ABI

// EnumerableMapBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableMapMetaData.Bin instead.
var EnumerableMapBin = EnumerableMapMetaData.Bin

// DeployEnumerableMap deploys a new Ethereum contract, binding an instance of EnumerableMap to it.
func DeployEnumerableMap(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableMap, error) {
	parsed, err := EnumerableMapMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableMapBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// EnumerableMap is an auto generated Go binding around an Ethereum contract.
type EnumerableMap struct {
	EnumerableMapCaller     // Read-only binding to the contract
	EnumerableMapTransactor // Write-only binding to the contract
	EnumerableMapFilterer   // Log filterer for contract events
}

// EnumerableMapCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableMapCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableMapTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableMapFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableMapSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableMapSession struct {
	Contract     *EnumerableMap    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableMapCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableMapCallerSession struct {
	Contract *EnumerableMapCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableMapTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableMapTransactorSession struct {
	Contract     *EnumerableMapTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableMapRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableMapRaw struct {
	Contract *EnumerableMap // Generic contract binding to access the raw methods on
}

// EnumerableMapCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableMapCallerRaw struct {
	Contract *EnumerableMapCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableMapTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableMapTransactorRaw struct {
	Contract *EnumerableMapTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableMap creates a new instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMap(address common.Address, backend bind.ContractBackend) (*EnumerableMap, error) {
	contract, err := bindEnumerableMap(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableMap{EnumerableMapCaller: EnumerableMapCaller{contract: contract}, EnumerableMapTransactor: EnumerableMapTransactor{contract: contract}, EnumerableMapFilterer: EnumerableMapFilterer{contract: contract}}, nil
}

// NewEnumerableMapCaller creates a new read-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapCaller(address common.Address, caller bind.ContractCaller) (*EnumerableMapCaller, error) {
	contract, err := bindEnumerableMap(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapCaller{contract: contract}, nil
}

// NewEnumerableMapTransactor creates a new write-only instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableMapTransactor, error) {
	contract, err := bindEnumerableMap(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapTransactor{contract: contract}, nil
}

// NewEnumerableMapFilterer creates a new log filterer instance of EnumerableMap, bound to a specific deployed contract.
func NewEnumerableMapFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableMapFilterer, error) {
	contract, err := bindEnumerableMap(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableMapFilterer{contract: contract}, nil
}

// bindEnumerableMap binds a generic wrapper to an already deployed contract.
func bindEnumerableMap(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableMapMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.EnumerableMapCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.EnumerableMapTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableMap *EnumerableMapCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableMap.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableMap *EnumerableMapTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableMap.Contract.contract.Transact(opts, method, params...)
}

// EnumerableSetMetaData contains all meta data concerning the EnumerableSet contract.
var EnumerableSetMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220a769f91276e72cf9de6e95ddce7caaf62c3d8a0ac06d034c64e26d3dc9be2afa64736f6c63430008130033",
}

// EnumerableSetABI is the input ABI used to generate the binding from.
// Deprecated: Use EnumerableSetMetaData.ABI instead.
var EnumerableSetABI = EnumerableSetMetaData.ABI

// EnumerableSetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnumerableSetMetaData.Bin instead.
var EnumerableSetBin = EnumerableSetMetaData.Bin

// DeployEnumerableSet deploys a new Ethereum contract, binding an instance of EnumerableSet to it.
func DeployEnumerableSet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EnumerableSet, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnumerableSetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// EnumerableSet is an auto generated Go binding around an Ethereum contract.
type EnumerableSet struct {
	EnumerableSetCaller     // Read-only binding to the contract
	EnumerableSetTransactor // Write-only binding to the contract
	EnumerableSetFilterer   // Log filterer for contract events
}

// EnumerableSetCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnumerableSetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnumerableSetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnumerableSetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnumerableSetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnumerableSetSession struct {
	Contract     *EnumerableSet    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnumerableSetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnumerableSetCallerSession struct {
	Contract *EnumerableSetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// EnumerableSetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnumerableSetTransactorSession struct {
	Contract     *EnumerableSetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// EnumerableSetRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnumerableSetRaw struct {
	Contract *EnumerableSet // Generic contract binding to access the raw methods on
}

// EnumerableSetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnumerableSetCallerRaw struct {
	Contract *EnumerableSetCaller // Generic read-only contract binding to access the raw methods on
}

// EnumerableSetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnumerableSetTransactorRaw struct {
	Contract *EnumerableSetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnumerableSet creates a new instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSet(address common.Address, backend bind.ContractBackend) (*EnumerableSet, error) {
	contract, err := bindEnumerableSet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnumerableSet{EnumerableSetCaller: EnumerableSetCaller{contract: contract}, EnumerableSetTransactor: EnumerableSetTransactor{contract: contract}, EnumerableSetFilterer: EnumerableSetFilterer{contract: contract}}, nil
}

// NewEnumerableSetCaller creates a new read-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetCaller(address common.Address, caller bind.ContractCaller) (*EnumerableSetCaller, error) {
	contract, err := bindEnumerableSet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetCaller{contract: contract}, nil
}

// NewEnumerableSetTransactor creates a new write-only instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetTransactor(address common.Address, transactor bind.ContractTransactor) (*EnumerableSetTransactor, error) {
	contract, err := bindEnumerableSet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetTransactor{contract: contract}, nil
}

// NewEnumerableSetFilterer creates a new log filterer instance of EnumerableSet, bound to a specific deployed contract.
func NewEnumerableSetFilterer(address common.Address, filterer bind.ContractFilterer) (*EnumerableSetFilterer, error) {
	contract, err := bindEnumerableSet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnumerableSetFilterer{contract: contract}, nil
}

// bindEnumerableSet binds a generic wrapper to an already deployed contract.
func bindEnumerableSet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EnumerableSetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.EnumerableSetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.EnumerableSetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnumerableSet *EnumerableSetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnumerableSet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnumerableSet *EnumerableSetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnumerableSet.Contract.contract.Transact(opts, method, params...)
}

// IBeaconUpgradeableMetaData contains all meta data concerning the IBeaconUpgradeable contract.
var IBeaconUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IBeaconUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IBeaconUpgradeableMetaData.ABI instead.
var IBeaconUpgradeableABI = IBeaconUpgradeableMetaData.ABI

// IBeaconUpgradeable is an auto generated Go binding around an Ethereum contract.
type IBeaconUpgradeable struct {
	IBeaconUpgradeableCaller     // Read-only binding to the contract
	IBeaconUpgradeableTransactor // Write-only binding to the contract
	IBeaconUpgradeableFilterer   // Log filterer for contract events
}

// IBeaconUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBeaconUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBeaconUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBeaconUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBeaconUpgradeableSession struct {
	Contract     *IBeaconUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBeaconUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBeaconUpgradeableCallerSession struct {
	Contract *IBeaconUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IBeaconUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBeaconUpgradeableTransactorSession struct {
	Contract     *IBeaconUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IBeaconUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBeaconUpgradeableRaw struct {
	Contract *IBeaconUpgradeable // Generic contract binding to access the raw methods on
}

// IBeaconUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBeaconUpgradeableCallerRaw struct {
	Contract *IBeaconUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IBeaconUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBeaconUpgradeableTransactorRaw struct {
	Contract *IBeaconUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBeaconUpgradeable creates a new instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeable(address common.Address, backend bind.ContractBackend) (*IBeaconUpgradeable, error) {
	contract, err := bindIBeaconUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeable{IBeaconUpgradeableCaller: IBeaconUpgradeableCaller{contract: contract}, IBeaconUpgradeableTransactor: IBeaconUpgradeableTransactor{contract: contract}, IBeaconUpgradeableFilterer: IBeaconUpgradeableFilterer{contract: contract}}, nil
}

// NewIBeaconUpgradeableCaller creates a new read-only instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IBeaconUpgradeableCaller, error) {
	contract, err := bindIBeaconUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableCaller{contract: contract}, nil
}

// NewIBeaconUpgradeableTransactor creates a new write-only instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IBeaconUpgradeableTransactor, error) {
	contract, err := bindIBeaconUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableTransactor{contract: contract}, nil
}

// NewIBeaconUpgradeableFilterer creates a new log filterer instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IBeaconUpgradeableFilterer, error) {
	contract, err := bindIBeaconUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableFilterer{contract: contract}, nil
}

// bindIBeaconUpgradeable binds a generic wrapper to an already deployed contract.
func bindIBeaconUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IBeaconUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconUpgradeable *IBeaconUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconUpgradeable *IBeaconUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconUpgradeable *IBeaconUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBeaconUpgradeable.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableSession) Implementation() (common.Address, error) {
	return _IBeaconUpgradeable.Contract.Implementation(&_IBeaconUpgradeable.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableCallerSession) Implementation() (common.Address, error) {
	return _IBeaconUpgradeable.Contract.Implementation(&_IBeaconUpgradeable.CallOpts)
}

// IERC1822ProxiableUpgradeableMetaData contains all meta data concerning the IERC1822ProxiableUpgradeable contract.
var IERC1822ProxiableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IERC1822ProxiableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC1822ProxiableUpgradeableMetaData.ABI instead.
var IERC1822ProxiableUpgradeableABI = IERC1822ProxiableUpgradeableMetaData.ABI

// IERC1822ProxiableUpgradeable is an auto generated Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeable struct {
	IERC1822ProxiableUpgradeableCaller     // Read-only binding to the contract
	IERC1822ProxiableUpgradeableTransactor // Write-only binding to the contract
	IERC1822ProxiableUpgradeableFilterer   // Log filterer for contract events
}

// IERC1822ProxiableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC1822ProxiableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC1822ProxiableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC1822ProxiableUpgradeableSession struct {
	Contract     *IERC1822ProxiableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                 // Call options to use throughout this session
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IERC1822ProxiableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC1822ProxiableUpgradeableCallerSession struct {
	Contract *IERC1822ProxiableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                       // Call options to use throughout this session
}

// IERC1822ProxiableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC1822ProxiableUpgradeableTransactorSession struct {
	Contract     *IERC1822ProxiableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                       // Transaction auth options to use throughout this session
}

// IERC1822ProxiableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableRaw struct {
	Contract *IERC1822ProxiableUpgradeable // Generic contract binding to access the raw methods on
}

// IERC1822ProxiableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableCallerRaw struct {
	Contract *IERC1822ProxiableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC1822ProxiableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC1822ProxiableUpgradeableTransactorRaw struct {
	Contract *IERC1822ProxiableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC1822ProxiableUpgradeable creates a new instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeable(address common.Address, backend bind.ContractBackend) (*IERC1822ProxiableUpgradeable, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeable{IERC1822ProxiableUpgradeableCaller: IERC1822ProxiableUpgradeableCaller{contract: contract}, IERC1822ProxiableUpgradeableTransactor: IERC1822ProxiableUpgradeableTransactor{contract: contract}, IERC1822ProxiableUpgradeableFilterer: IERC1822ProxiableUpgradeableFilterer{contract: contract}}, nil
}

// NewIERC1822ProxiableUpgradeableCaller creates a new read-only instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC1822ProxiableUpgradeableCaller, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableCaller{contract: contract}, nil
}

// NewIERC1822ProxiableUpgradeableTransactor creates a new write-only instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC1822ProxiableUpgradeableTransactor, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableTransactor{contract: contract}, nil
}

// NewIERC1822ProxiableUpgradeableFilterer creates a new log filterer instance of IERC1822ProxiableUpgradeable, bound to a specific deployed contract.
func NewIERC1822ProxiableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC1822ProxiableUpgradeableFilterer, error) {
	contract, err := bindIERC1822ProxiableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC1822ProxiableUpgradeableFilterer{contract: contract}, nil
}

// bindIERC1822ProxiableUpgradeable binds a generic wrapper to an already deployed contract.
func bindIERC1822ProxiableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IERC1822ProxiableUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.IERC1822ProxiableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC1822ProxiableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _IERC1822ProxiableUpgradeable.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableSession) ProxiableUUID() ([32]byte, error) {
	return _IERC1822ProxiableUpgradeable.Contract.ProxiableUUID(&_IERC1822ProxiableUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_IERC1822ProxiableUpgradeable *IERC1822ProxiableUpgradeableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _IERC1822ProxiableUpgradeable.Contract.ProxiableUUID(&_IERC1822ProxiableUpgradeable.CallOpts)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

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
	parsed, err := IERC20MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
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
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, amount)
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

// IGovernorMetaData contains all meta data concerning the IGovernor contract.
var IGovernorMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"candidateID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenOffer\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structIMessage.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"}],\"internalType\":\"structIMessage.DomainSeparator\",\"name\":\"domain\",\"type\":\"tuple\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"castVoteWithReason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"candidateID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tokenOffer\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIGovernor.ProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"candidateID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tokenOffer\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"validatorCnt\",\"type\":\"uint32\"},{\"internalType\":\"enumIGovernor.ProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IGovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use IGovernorMetaData.ABI instead.
var IGovernorABI = IGovernorMetaData.ABI

// IGovernor is an auto generated Go binding around an Ethereum contract.
type IGovernor struct {
	IGovernorCaller     // Read-only binding to the contract
	IGovernorTransactor // Write-only binding to the contract
	IGovernorFilterer   // Log filterer for contract events
}

// IGovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type IGovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IGovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IGovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IGovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IGovernorSession struct {
	Contract     *IGovernor        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IGovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IGovernorCallerSession struct {
	Contract *IGovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IGovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IGovernorTransactorSession struct {
	Contract     *IGovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IGovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type IGovernorRaw struct {
	Contract *IGovernor // Generic contract binding to access the raw methods on
}

// IGovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IGovernorCallerRaw struct {
	Contract *IGovernorCaller // Generic read-only contract binding to access the raw methods on
}

// IGovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IGovernorTransactorRaw struct {
	Contract *IGovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIGovernor creates a new instance of IGovernor, bound to a specific deployed contract.
func NewIGovernor(address common.Address, backend bind.ContractBackend) (*IGovernor, error) {
	contract, err := bindIGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IGovernor{IGovernorCaller: IGovernorCaller{contract: contract}, IGovernorTransactor: IGovernorTransactor{contract: contract}, IGovernorFilterer: IGovernorFilterer{contract: contract}}, nil
}

// NewIGovernorCaller creates a new read-only instance of IGovernor, bound to a specific deployed contract.
func NewIGovernorCaller(address common.Address, caller bind.ContractCaller) (*IGovernorCaller, error) {
	contract, err := bindIGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IGovernorCaller{contract: contract}, nil
}

// NewIGovernorTransactor creates a new write-only instance of IGovernor, bound to a specific deployed contract.
func NewIGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*IGovernorTransactor, error) {
	contract, err := bindIGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IGovernorTransactor{contract: contract}, nil
}

// NewIGovernorFilterer creates a new log filterer instance of IGovernor, bound to a specific deployed contract.
func NewIGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*IGovernorFilterer, error) {
	contract, err := bindIGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IGovernorFilterer{contract: contract}, nil
}

// bindIGovernor binds a generic wrapper to an already deployed contract.
func bindIGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGovernor *IGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGovernor.Contract.IGovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGovernor *IGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovernor.Contract.IGovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGovernor *IGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGovernor.Contract.IGovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IGovernor *IGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IGovernor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IGovernor *IGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IGovernor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IGovernor *IGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IGovernor.Contract.contract.Transact(opts, method, params...)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_IGovernor *IGovernorCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _IGovernor.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_IGovernor *IGovernorSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _IGovernor.Contract.HasVoted(&_IGovernor.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_IGovernor *IGovernorCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _IGovernor.Contract.HasVoted(&_IGovernor.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0x3a4b0e3b.
//
// Solidity: function hashProposal(bytes32 descriptionHash, address proposer, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint8 proposalType) pure returns(uint256)
func (_IGovernor *IGovernorCaller) HashProposal(opts *bind.CallOpts, descriptionHash [32]byte, proposer common.Address, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, proposalType uint8) (*big.Int, error) {
	var out []interface{}
	err := _IGovernor.contract.Call(opts, &out, "hashProposal", descriptionHash, proposer, candidateID, candidate, tokenOffer, merkleRoot, proposalType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0x3a4b0e3b.
//
// Solidity: function hashProposal(bytes32 descriptionHash, address proposer, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint8 proposalType) pure returns(uint256)
func (_IGovernor *IGovernorSession) HashProposal(descriptionHash [32]byte, proposer common.Address, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, proposalType uint8) (*big.Int, error) {
	return _IGovernor.Contract.HashProposal(&_IGovernor.CallOpts, descriptionHash, proposer, candidateID, candidate, tokenOffer, merkleRoot, proposalType)
}

// HashProposal is a free data retrieval call binding the contract method 0x3a4b0e3b.
//
// Solidity: function hashProposal(bytes32 descriptionHash, address proposer, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint8 proposalType) pure returns(uint256)
func (_IGovernor *IGovernorCallerSession) HashProposal(descriptionHash [32]byte, proposer common.Address, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, proposalType uint8) (*big.Int, error) {
	return _IGovernor.Contract.HashProposal(&_IGovernor.CallOpts, descriptionHash, proposer, candidateID, candidate, tokenOffer, merkleRoot, proposalType)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_IGovernor *IGovernorCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _IGovernor.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_IGovernor *IGovernorSession) State(proposalId *big.Int) (uint8, error) {
	return _IGovernor.Contract.State(&_IGovernor.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_IGovernor *IGovernorCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _IGovernor.Contract.State(&_IGovernor.CallOpts, proposalId)
}

// CastVote is a paid mutator transaction binding the contract method 0xd92202f8.
//
// Solidity: function castVote(uint256 proposalId, uint8 support, bytes32[] proof) returns()
func (_IGovernor *IGovernorTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8, proof [][32]byte) (*types.Transaction, error) {
	return _IGovernor.contract.Transact(opts, "castVote", proposalId, support, proof)
}

// CastVote is a paid mutator transaction binding the contract method 0xd92202f8.
//
// Solidity: function castVote(uint256 proposalId, uint8 support, bytes32[] proof) returns()
func (_IGovernor *IGovernorSession) CastVote(proposalId *big.Int, support uint8, proof [][32]byte) (*types.Transaction, error) {
	return _IGovernor.Contract.CastVote(&_IGovernor.TransactOpts, proposalId, support, proof)
}

// CastVote is a paid mutator transaction binding the contract method 0xd92202f8.
//
// Solidity: function castVote(uint256 proposalId, uint8 support, bytes32[] proof) returns()
func (_IGovernor *IGovernorTransactorSession) CastVote(proposalId *big.Int, support uint8, proof [][32]byte) (*types.Transaction, error) {
	return _IGovernor.Contract.CastVote(&_IGovernor.TransactOpts, proposalId, support, proof)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x02ce3b1a.
//
// Solidity: function castVoteBySig((uint256,uint8,string) message, bytes signature, (string,string,uint256,address) domain) returns(uint256 balance)
func (_IGovernor *IGovernorTransactor) CastVoteBySig(opts *bind.TransactOpts, message IMessageMessage, signature []byte, domain IMessageDomainSeparator) (*types.Transaction, error) {
	return _IGovernor.contract.Transact(opts, "castVoteBySig", message, signature, domain)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x02ce3b1a.
//
// Solidity: function castVoteBySig((uint256,uint8,string) message, bytes signature, (string,string,uint256,address) domain) returns(uint256 balance)
func (_IGovernor *IGovernorSession) CastVoteBySig(message IMessageMessage, signature []byte, domain IMessageDomainSeparator) (*types.Transaction, error) {
	return _IGovernor.Contract.CastVoteBySig(&_IGovernor.TransactOpts, message, signature, domain)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x02ce3b1a.
//
// Solidity: function castVoteBySig((uint256,uint8,string) message, bytes signature, (string,string,uint256,address) domain) returns(uint256 balance)
func (_IGovernor *IGovernorTransactorSession) CastVoteBySig(message IMessageMessage, signature []byte, domain IMessageDomainSeparator) (*types.Transaction, error) {
	return _IGovernor.Contract.CastVoteBySig(&_IGovernor.TransactOpts, message, signature, domain)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x1d46e7b9.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason, bytes32[] proof) returns()
func (_IGovernor *IGovernorTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, proof [][32]byte) (*types.Transaction, error) {
	return _IGovernor.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason, proof)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x1d46e7b9.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason, bytes32[] proof) returns()
func (_IGovernor *IGovernorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string, proof [][32]byte) (*types.Transaction, error) {
	return _IGovernor.Contract.CastVoteWithReason(&_IGovernor.TransactOpts, proposalId, support, reason, proof)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x1d46e7b9.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason, bytes32[] proof) returns()
func (_IGovernor *IGovernorTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string, proof [][32]byte) (*types.Transaction, error) {
	return _IGovernor.Contract.CastVoteWithReason(&_IGovernor.TransactOpts, proposalId, support, reason, proof)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns(bool)
func (_IGovernor *IGovernorTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _IGovernor.contract.Transact(opts, "execute", proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns(bool)
func (_IGovernor *IGovernorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _IGovernor.Contract.Execute(&_IGovernor.TransactOpts, proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns(bool)
func (_IGovernor *IGovernorTransactorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _IGovernor.Contract.Execute(&_IGovernor.TransactOpts, proposalId)
}

// Propose is a paid mutator transaction binding the contract method 0xa618b8c1.
//
// Solidity: function propose(string description, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint32 validatorCnt, uint8 proposalType) returns(uint256)
func (_IGovernor *IGovernorTransactor) Propose(opts *bind.TransactOpts, description string, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, validatorCnt uint32, proposalType uint8) (*types.Transaction, error) {
	return _IGovernor.contract.Transact(opts, "propose", description, candidateID, candidate, tokenOffer, merkleRoot, validatorCnt, proposalType)
}

// Propose is a paid mutator transaction binding the contract method 0xa618b8c1.
//
// Solidity: function propose(string description, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint32 validatorCnt, uint8 proposalType) returns(uint256)
func (_IGovernor *IGovernorSession) Propose(description string, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, validatorCnt uint32, proposalType uint8) (*types.Transaction, error) {
	return _IGovernor.Contract.Propose(&_IGovernor.TransactOpts, description, candidateID, candidate, tokenOffer, merkleRoot, validatorCnt, proposalType)
}

// Propose is a paid mutator transaction binding the contract method 0xa618b8c1.
//
// Solidity: function propose(string description, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint32 validatorCnt, uint8 proposalType) returns(uint256)
func (_IGovernor *IGovernorTransactorSession) Propose(description string, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, validatorCnt uint32, proposalType uint8) (*types.Transaction, error) {
	return _IGovernor.Contract.Propose(&_IGovernor.TransactOpts, description, candidateID, candidate, tokenOffer, merkleRoot, validatorCnt, proposalType)
}

// IGovernorProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the IGovernor contract.
type IGovernorProposalCanceledIterator struct {
	Event *IGovernorProposalCanceled // Event containing the contract specifics and raw log

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
func (it *IGovernorProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovernorProposalCanceled)
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
		it.Event = new(IGovernorProposalCanceled)
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
func (it *IGovernorProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovernorProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovernorProposalCanceled represents a ProposalCanceled event raised by the IGovernor contract.
type IGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_IGovernor *IGovernorFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*IGovernorProposalCanceledIterator, error) {

	logs, sub, err := _IGovernor.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &IGovernorProposalCanceledIterator{contract: _IGovernor.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_IGovernor *IGovernorFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *IGovernorProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _IGovernor.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovernorProposalCanceled)
				if err := _IGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_IGovernor *IGovernorFilterer) ParseProposalCanceled(log types.Log) (*IGovernorProposalCanceled, error) {
	event := new(IGovernorProposalCanceled)
	if err := _IGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGovernorProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the IGovernor contract.
type IGovernorProposalCreatedIterator struct {
	Event *IGovernorProposalCreated // Event containing the contract specifics and raw log

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
func (it *IGovernorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovernorProposalCreated)
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
		it.Event = new(IGovernorProposalCreated)
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
func (it *IGovernorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovernorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovernorProposalCreated represents a ProposalCreated event raised by the IGovernor contract.
type IGovernorProposalCreated struct {
	ProposalId     *big.Int
	CandidateID    [32]byte
	Candidate      common.Address
	TokenOffer     uint32
	MerkleRoot     [32]byte
	StartTimestamp uint64
	Description    string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x76336534418640ae4cd606c430d9287a4790c09310247a7a9fe90435f297acf9.
//
// Solidity: event ProposalCreated(uint256 proposalId, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint64 startTimestamp, string description)
func (_IGovernor *IGovernorFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*IGovernorProposalCreatedIterator, error) {

	logs, sub, err := _IGovernor.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &IGovernorProposalCreatedIterator{contract: _IGovernor.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x76336534418640ae4cd606c430d9287a4790c09310247a7a9fe90435f297acf9.
//
// Solidity: event ProposalCreated(uint256 proposalId, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint64 startTimestamp, string description)
func (_IGovernor *IGovernorFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *IGovernorProposalCreated) (event.Subscription, error) {

	logs, sub, err := _IGovernor.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovernorProposalCreated)
				if err := _IGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x76336534418640ae4cd606c430d9287a4790c09310247a7a9fe90435f297acf9.
//
// Solidity: event ProposalCreated(uint256 proposalId, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint64 startTimestamp, string description)
func (_IGovernor *IGovernorFilterer) ParseProposalCreated(log types.Log) (*IGovernorProposalCreated, error) {
	event := new(IGovernorProposalCreated)
	if err := _IGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGovernorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the IGovernor contract.
type IGovernorProposalExecutedIterator struct {
	Event *IGovernorProposalExecuted // Event containing the contract specifics and raw log

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
func (it *IGovernorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovernorProposalExecuted)
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
		it.Event = new(IGovernorProposalExecuted)
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
func (it *IGovernorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovernorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovernorProposalExecuted represents a ProposalExecuted event raised by the IGovernor contract.
type IGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_IGovernor *IGovernorFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*IGovernorProposalExecutedIterator, error) {

	logs, sub, err := _IGovernor.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &IGovernorProposalExecutedIterator{contract: _IGovernor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_IGovernor *IGovernorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *IGovernorProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _IGovernor.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovernorProposalExecuted)
				if err := _IGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_IGovernor *IGovernorFilterer) ParseProposalExecuted(log types.Log) (*IGovernorProposalExecuted, error) {
	event := new(IGovernorProposalExecuted)
	if err := _IGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IGovernorVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the IGovernor contract.
type IGovernorVoteCastIterator struct {
	Event *IGovernorVoteCast // Event containing the contract specifics and raw log

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
func (it *IGovernorVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IGovernorVoteCast)
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
		it.Event = new(IGovernorVoteCast)
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
func (it *IGovernorVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IGovernorVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IGovernorVoteCast represents a VoteCast event raised by the IGovernor contract.
type IGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x5be2678cab71c7d94235767eb1df3c57673f19e805cdc0edbd2e74854ae9f0f9.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, string reason)
func (_IGovernor *IGovernorFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*IGovernorVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IGovernor.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &IGovernorVoteCastIterator{contract: _IGovernor.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x5be2678cab71c7d94235767eb1df3c57673f19e805cdc0edbd2e74854ae9f0f9.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, string reason)
func (_IGovernor *IGovernorFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *IGovernorVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _IGovernor.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IGovernorVoteCast)
				if err := _IGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0x5be2678cab71c7d94235767eb1df3c57673f19e805cdc0edbd2e74854ae9f0f9.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, string reason)
func (_IGovernor *IGovernorFilterer) ParseVoteCast(log types.Log) (*IGovernorVoteCast, error) {
	event := new(IGovernorVoteCast)
	if err := _IGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMessageMetaData contains all meta data concerning the IMessage contract.
var IMessageMetaData = &bind.MetaData{
	ABI: "[]",
}

// IMessageABI is the input ABI used to generate the binding from.
// Deprecated: Use IMessageMetaData.ABI instead.
var IMessageABI = IMessageMetaData.ABI

// IMessage is an auto generated Go binding around an Ethereum contract.
type IMessage struct {
	IMessageCaller     // Read-only binding to the contract
	IMessageTransactor // Write-only binding to the contract
	IMessageFilterer   // Log filterer for contract events
}

// IMessageCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMessageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMessageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMessageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMessageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMessageSession struct {
	Contract     *IMessage         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMessageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMessageCallerSession struct {
	Contract *IMessageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IMessageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMessageTransactorSession struct {
	Contract     *IMessageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IMessageRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMessageRaw struct {
	Contract *IMessage // Generic contract binding to access the raw methods on
}

// IMessageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMessageCallerRaw struct {
	Contract *IMessageCaller // Generic read-only contract binding to access the raw methods on
}

// IMessageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMessageTransactorRaw struct {
	Contract *IMessageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMessage creates a new instance of IMessage, bound to a specific deployed contract.
func NewIMessage(address common.Address, backend bind.ContractBackend) (*IMessage, error) {
	contract, err := bindIMessage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMessage{IMessageCaller: IMessageCaller{contract: contract}, IMessageTransactor: IMessageTransactor{contract: contract}, IMessageFilterer: IMessageFilterer{contract: contract}}, nil
}

// NewIMessageCaller creates a new read-only instance of IMessage, bound to a specific deployed contract.
func NewIMessageCaller(address common.Address, caller bind.ContractCaller) (*IMessageCaller, error) {
	contract, err := bindIMessage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageCaller{contract: contract}, nil
}

// NewIMessageTransactor creates a new write-only instance of IMessage, bound to a specific deployed contract.
func NewIMessageTransactor(address common.Address, transactor bind.ContractTransactor) (*IMessageTransactor, error) {
	contract, err := bindIMessage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMessageTransactor{contract: contract}, nil
}

// NewIMessageFilterer creates a new log filterer instance of IMessage, bound to a specific deployed contract.
func NewIMessageFilterer(address common.Address, filterer bind.ContractFilterer) (*IMessageFilterer, error) {
	contract, err := bindIMessage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMessageFilterer{contract: contract}, nil
}

// bindIMessage binds a generic wrapper to an already deployed contract.
func bindIMessage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IMessageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessage *IMessageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessage.Contract.IMessageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessage *IMessageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessage.Contract.IMessageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessage *IMessageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessage.Contract.IMessageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMessage *IMessageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMessage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMessage *IMessageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMessage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMessage *IMessageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMessage.Contract.contract.Transact(opts, method, params...)
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"}]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InitializableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// InitializableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Initializable contract.
type InitializableInitializedIterator struct {
	Event *InitializableInitialized // Event containing the contract specifics and raw log

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
func (it *InitializableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(InitializableInitialized)
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
		it.Event = new(InitializableInitialized)
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
func (it *InitializableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *InitializableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// InitializableInitialized represents a Initialized event raised by the Initializable contract.
type InitializableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) FilterInitialized(opts *bind.FilterOpts) (*InitializableInitializedIterator, error) {

	logs, sub, err := _Initializable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &InitializableInitializedIterator{contract: _Initializable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *InitializableInitialized) (event.Subscription, error) {

	logs, sub, err := _Initializable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(InitializableInitialized)
				if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Initializable *InitializableFilterer) ParseInitialized(log types.Log) (*InitializableInitialized, error) {
	event := new(InitializableInitialized)
	if err := _Initializable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MerkleProofMetaData contains all meta data concerning the MerkleProof contract.
var MerkleProofMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203d2a6285f3867e0981bfea164614f1c12f37823fac95a8c53669fd6d54cf9e9c64736f6c63430008130033",
}

// MerkleProofABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleProofMetaData.ABI instead.
var MerkleProofABI = MerkleProofMetaData.ABI

// MerkleProofBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleProofMetaData.Bin instead.
var MerkleProofBin = MerkleProofMetaData.Bin

// DeployMerkleProof deploys a new Ethereum contract, binding an instance of MerkleProof to it.
func DeployMerkleProof(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleProof, error) {
	parsed, err := MerkleProofMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleProofBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleProof{MerkleProofCaller: MerkleProofCaller{contract: contract}, MerkleProofTransactor: MerkleProofTransactor{contract: contract}, MerkleProofFilterer: MerkleProofFilterer{contract: contract}}, nil
}

// MerkleProof is an auto generated Go binding around an Ethereum contract.
type MerkleProof struct {
	MerkleProofCaller     // Read-only binding to the contract
	MerkleProofTransactor // Write-only binding to the contract
	MerkleProofFilterer   // Log filterer for contract events
}

// MerkleProofCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleProofCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleProofTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleProofFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleProofSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleProofSession struct {
	Contract     *MerkleProof      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleProofCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleProofCallerSession struct {
	Contract *MerkleProofCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MerkleProofTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleProofTransactorSession struct {
	Contract     *MerkleProofTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MerkleProofRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleProofRaw struct {
	Contract *MerkleProof // Generic contract binding to access the raw methods on
}

// MerkleProofCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleProofCallerRaw struct {
	Contract *MerkleProofCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleProofTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleProofTransactorRaw struct {
	Contract *MerkleProofTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleProof creates a new instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProof(address common.Address, backend bind.ContractBackend) (*MerkleProof, error) {
	contract, err := bindMerkleProof(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleProof{MerkleProofCaller: MerkleProofCaller{contract: contract}, MerkleProofTransactor: MerkleProofTransactor{contract: contract}, MerkleProofFilterer: MerkleProofFilterer{contract: contract}}, nil
}

// NewMerkleProofCaller creates a new read-only instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofCaller(address common.Address, caller bind.ContractCaller) (*MerkleProofCaller, error) {
	contract, err := bindMerkleProof(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleProofCaller{contract: contract}, nil
}

// NewMerkleProofTransactor creates a new write-only instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleProofTransactor, error) {
	contract, err := bindMerkleProof(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleProofTransactor{contract: contract}, nil
}

// NewMerkleProofFilterer creates a new log filterer instance of MerkleProof, bound to a specific deployed contract.
func NewMerkleProofFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleProofFilterer, error) {
	contract, err := bindMerkleProof(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleProofFilterer{contract: contract}, nil
}

// bindMerkleProof binds a generic wrapper to an already deployed contract.
func bindMerkleProof(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleProofMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleProof *MerkleProofRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleProof.Contract.MerkleProofCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleProof *MerkleProofRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleProof.Contract.MerkleProofTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleProof *MerkleProofRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleProof.Contract.MerkleProofTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleProof *MerkleProofCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleProof.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleProof *MerkleProofTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleProof.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleProof *MerkleProofTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleProof.Contract.contract.Transact(opts, method, params...)
}

// OwnableUpgradeableMetaData contains all meta data concerning the OwnableUpgradeable contract.
var OwnableUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableUpgradeableMetaData.ABI instead.
var OwnableUpgradeableABI = OwnableUpgradeableMetaData.ABI

// OwnableUpgradeable is an auto generated Go binding around an Ethereum contract.
type OwnableUpgradeable struct {
	OwnableUpgradeableCaller     // Read-only binding to the contract
	OwnableUpgradeableTransactor // Write-only binding to the contract
	OwnableUpgradeableFilterer   // Log filterer for contract events
}

// OwnableUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableUpgradeableSession struct {
	Contract     *OwnableUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OwnableUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableUpgradeableCallerSession struct {
	Contract *OwnableUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// OwnableUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableUpgradeableTransactorSession struct {
	Contract     *OwnableUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// OwnableUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableUpgradeableRaw struct {
	Contract *OwnableUpgradeable // Generic contract binding to access the raw methods on
}

// OwnableUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableUpgradeableCallerRaw struct {
	Contract *OwnableUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableUpgradeableTransactorRaw struct {
	Contract *OwnableUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableUpgradeable creates a new instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeable(address common.Address, backend bind.ContractBackend) (*OwnableUpgradeable, error) {
	contract, err := bindOwnableUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeable{OwnableUpgradeableCaller: OwnableUpgradeableCaller{contract: contract}, OwnableUpgradeableTransactor: OwnableUpgradeableTransactor{contract: contract}, OwnableUpgradeableFilterer: OwnableUpgradeableFilterer{contract: contract}}, nil
}

// NewOwnableUpgradeableCaller creates a new read-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*OwnableUpgradeableCaller, error) {
	contract, err := bindOwnableUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableCaller{contract: contract}, nil
}

// NewOwnableUpgradeableTransactor creates a new write-only instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableUpgradeableTransactor, error) {
	contract, err := bindOwnableUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableTransactor{contract: contract}, nil
}

// NewOwnableUpgradeableFilterer creates a new log filterer instance of OwnableUpgradeable, bound to a specific deployed contract.
func NewOwnableUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableUpgradeableFilterer, error) {
	contract, err := bindOwnableUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableFilterer{contract: contract}, nil
}

// bindOwnableUpgradeable binds a generic wrapper to an already deployed contract.
func bindOwnableUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.OwnableUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableUpgradeable *OwnableUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableUpgradeable *OwnableUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableUpgradeable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableUpgradeable *OwnableUpgradeableCallerSession) Owner() (common.Address, error) {
	return _OwnableUpgradeable.Contract.Owner(&_OwnableUpgradeable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.RenounceOwnership(&_OwnableUpgradeable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OwnableUpgradeable *OwnableUpgradeableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OwnableUpgradeable.Contract.TransferOwnership(&_OwnableUpgradeable.TransactOpts, newOwner)
}

// OwnableUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitializedIterator struct {
	Event *OwnableUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableInitialized)
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
		it.Event = new(OwnableUpgradeableInitialized)
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
func (it *OwnableUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableInitialized represents a Initialized event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*OwnableUpgradeableInitializedIterator, error) {

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableInitializedIterator{contract: _OwnableUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableInitialized)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseInitialized(log types.Log) (*OwnableUpgradeableInitialized, error) {
	event := new(OwnableUpgradeableInitialized)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableUpgradeableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferredIterator struct {
	Event *OwnableUpgradeableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableUpgradeableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableUpgradeableOwnershipTransferred)
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
		it.Event = new(OwnableUpgradeableOwnershipTransferred)
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
func (it *OwnableUpgradeableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableUpgradeableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableUpgradeableOwnershipTransferred represents a OwnershipTransferred event raised by the OwnableUpgradeable contract.
type OwnableUpgradeableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableUpgradeableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableUpgradeableOwnershipTransferredIterator{contract: _OwnableUpgradeable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableUpgradeableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OwnableUpgradeable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableUpgradeableOwnershipTransferred)
				if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OwnableUpgradeable *OwnableUpgradeableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableUpgradeableOwnershipTransferred, error) {
	event := new(OwnableUpgradeableOwnershipTransferred)
	if err := _OwnableUpgradeable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorageSlotUpgradeableMetaData contains all meta data concerning the StorageSlotUpgradeable contract.
var StorageSlotUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205dc86503018cae360ed6a66f2b03b5a4407be502bbb65f21491a466fd4ba208064736f6c63430008130033",
}

// StorageSlotUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSlotUpgradeableMetaData.ABI instead.
var StorageSlotUpgradeableABI = StorageSlotUpgradeableMetaData.ABI

// StorageSlotUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSlotUpgradeableMetaData.Bin instead.
var StorageSlotUpgradeableBin = StorageSlotUpgradeableMetaData.Bin

// DeployStorageSlotUpgradeable deploys a new Ethereum contract, binding an instance of StorageSlotUpgradeable to it.
func DeployStorageSlotUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSlotUpgradeable, error) {
	parsed, err := StorageSlotUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSlotUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSlotUpgradeable{StorageSlotUpgradeableCaller: StorageSlotUpgradeableCaller{contract: contract}, StorageSlotUpgradeableTransactor: StorageSlotUpgradeableTransactor{contract: contract}, StorageSlotUpgradeableFilterer: StorageSlotUpgradeableFilterer{contract: contract}}, nil
}

// StorageSlotUpgradeable is an auto generated Go binding around an Ethereum contract.
type StorageSlotUpgradeable struct {
	StorageSlotUpgradeableCaller     // Read-only binding to the contract
	StorageSlotUpgradeableTransactor // Write-only binding to the contract
	StorageSlotUpgradeableFilterer   // Log filterer for contract events
}

// StorageSlotUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSlotUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSlotUpgradeableSession struct {
	Contract     *StorageSlotUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StorageSlotUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageSlotUpgradeableCallerSession struct {
	Contract *StorageSlotUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StorageSlotUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageSlotUpgradeableTransactorSession struct {
	Contract     *StorageSlotUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StorageSlotUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageSlotUpgradeableRaw struct {
	Contract *StorageSlotUpgradeable // Generic contract binding to access the raw methods on
}

// StorageSlotUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableCallerRaw struct {
	Contract *StorageSlotUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSlotUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableTransactorRaw struct {
	Contract *StorageSlotUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSlotUpgradeable creates a new instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeable(address common.Address, backend bind.ContractBackend) (*StorageSlotUpgradeable, error) {
	contract, err := bindStorageSlotUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeable{StorageSlotUpgradeableCaller: StorageSlotUpgradeableCaller{contract: contract}, StorageSlotUpgradeableTransactor: StorageSlotUpgradeableTransactor{contract: contract}, StorageSlotUpgradeableFilterer: StorageSlotUpgradeableFilterer{contract: contract}}, nil
}

// NewStorageSlotUpgradeableCaller creates a new read-only instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*StorageSlotUpgradeableCaller, error) {
	contract, err := bindStorageSlotUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableCaller{contract: contract}, nil
}

// NewStorageSlotUpgradeableTransactor creates a new write-only instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSlotUpgradeableTransactor, error) {
	contract, err := bindStorageSlotUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableTransactor{contract: contract}, nil
}

// NewStorageSlotUpgradeableFilterer creates a new log filterer instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSlotUpgradeableFilterer, error) {
	contract, err := bindStorageSlotUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableFilterer{contract: contract}, nil
}

// bindStorageSlotUpgradeable binds a generic wrapper to an already deployed contract.
func bindStorageSlotUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StorageSlotUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlotUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// StrikeGovernorMetaData contains all meta data concerning the StrikeGovernor contract.
var StrikeGovernorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"candidateID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"tokenOffer\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"startTimestamp\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"castVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structIMessage.Message\",\"name\":\"message\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"}],\"internalType\":\"structIMessage.DomainSeparator\",\"name\":\"domain\",\"type\":\"tuple\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"castVoteWithReason\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"candidateID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tokenOffer\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"enumIGovernor.ProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_votingDelay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_votingPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"candidateID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"tokenOffer\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"merkleRoot\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"validatorCnt\",\"type\":\"uint32\"},{\"internalType\":\"enumIGovernor.ProposalType\",\"name\":\"proposalType\",\"type\":\"uint8\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newValue\",\"type\":\"address\"}],\"name\":\"updateRewardToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"updateVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"updateVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff1681525034801561004357600080fd5b5060805161428961007b60003960008181610585015281816106130152818161097501528181610a030152610ab301526142896000f3fe6080604052600436106101095760003560e01c8063715018a611610095578063d92202f811610064578063d92202f81461036c578063ef00ef4314610395578063f2fde38b146103be578063f8cf31cb146103e7578063fe0d94c11461041057610109565b8063715018a6146102c45780638da5cb5b146102db578063a618b8c114610306578063a6ab36f21461034357610109565b80633e4f49e6116100dc5780633e4f49e6146101da57806343859632146102175780634f1ef2861461025457806352d1902d1461027057806363d61a191461029b57610109565b806302ce3b1a1461010e5780631d46e7b91461014b5780633659cfe6146101745780633a4b0e3b1461019d575b600080fd5b34801561011a57600080fd5b506101356004803603810190610130919061287c565b610440565b6040516101429190612932565b60405180910390f35b34801561015757600080fd5b50610172600480360381019061016d9190612a03565b61047c565b005b34801561018057600080fd5b5061019b60048036038101906101969190612aaa565b610583565b005b3480156101a957600080fd5b506101c460048036038101906101bf9190612b6e565b61070b565b6040516101d19190612932565b60405180910390f35b3480156101e657600080fd5b5061020160048036038101906101fc9190612c10565b610750565b60405161020e9190612cb4565b60405180910390f35b34801561022357600080fd5b5061023e60048036038101906102399190612ccf565b610941565b60405161024b9190612d2a565b60405180910390f35b61026e60048036038101906102699190612d45565b610973565b005b34801561027c57600080fd5b50610285610aaf565b6040516102929190612db0565b60405180910390f35b3480156102a757600080fd5b506102c260048036038101906102bd9190612c10565b610b68565b005b3480156102d057600080fd5b506102d9610bbe565b005b3480156102e757600080fd5b506102f0610bd2565b6040516102fd9190612dda565b60405180910390f35b34801561031257600080fd5b5061032d60048036038101906103289190612df5565b610bfc565b60405161033a9190612932565b60405180910390f35b34801561034f57600080fd5b5061036a60048036038101906103659190612eb3565b610e44565b005b34801561037857600080fd5b50610393600480360381019061038e9190612f06565b610fe2565b005b3480156103a157600080fd5b506103bc60048036038101906103b79190612c10565b6110b2565b005b3480156103ca57600080fd5b506103e560048036038101906103e09190612aaa565b611108565b005b3480156103f357600080fd5b5061040e60048036038101906104099190612aaa565b61118b565b005b61042a60048036038101906104259190612c10565b611264565b6040516104379190612d2a565b60405180910390f35b600061044a61138c565b600061045785848661140a565b905061047185600001518287602001518860400151611449565b479150509392505050565b8561048561167d565b8383600060036000868152602001908152602001600020905060006104a861167d565b6040516020016104b89190612fc2565b60405160208183030381529060405280519060200120905060006104e28585856004015485611685565b905080610524576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051b9061303a565b60405180910390fd5b6105748d338e8e8e8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050611449565b50505050505050505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1603610611576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610608906130cc565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1661065061169e565b73ffffffffffffffffffffffffffffffffffffffff16146106a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069d9061315e565b60405180910390fd5b6106af816116f5565b61070881600067ffffffffffffffff8111156106ce576106cd6124a4565b5b6040519080825280601f01601f1916602001820160405280156107005781602001600182028036833780820191505090505b506000611775565b50565b60008787878787878760405160200161072a97969594939291906131d5565b6040516020818303038152906040528051906020012060001c9050979650505050505050565b6000806005600084815260200190815260200160002090508060020160009054906101000a900460ff161561078957600891505061093c565b8060020160019054906101000a900460ff16156107aa57600391505061093c565b426107f4826000016040518060200160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250506118e3565b67ffffffffffffffff161061080d57600191505061093c565b610856816001016040518060200160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff16815250506118e3565b67ffffffffffffffff1642111580156108755750610873836118f1565b155b1561088457600291505061093c565b6108cd816001016040518060200160405290816000820160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1667ffffffffffffffff1681525050611986565b806108dd57506108dc836118f1565b5b15610901576108eb836119b2565b6108f65760046108f9565b60055b91505061093c565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093390613290565b60405180910390fd5b919050565b600061096b8260046000868152602001908152602001600020600101611a2290919063ffffffff16565b905092915050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1603610a01576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109f8906130cc565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16610a4061169e565b73ffffffffffffffffffffffffffffffffffffffff1614610a96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8d9061315e565b60405180910390fd5b610a9f826116f5565b610aab82826001611775565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614610b3f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b3690613322565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b610b7061138c565b8060015403610bb4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bab9061338e565b60405180910390fd5b8060018190555050565b610bc661138c565b610bd06000611a52565b565b6000603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000610c0661138c565b6000610c1e8980519060200120338a8a8a8a8961070b565b9050600060056000838152602001908152602001600020905060006001546001610c4891906133dd565b42610c539190613433565b905060006002546001610c6691906133dd565b42610c719190613433565b9050610c898284600001611b1890919063ffffffff16565b610c9f8184600101611b1890919063ffffffff16565b5050506000898051906020012090506000600360008481526020019081526020016000209050818160000181905550898160020181905550338160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550888160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550878160030160146101000a81548163ffffffff021916908363ffffffff160217905550868160040181905550848160050160006101000a81548160ff02191690836001811115610db057610daf612c3d565b5b0217905550505060006004600083815260200190815260200160002090508481600001600c6101000a81548163ffffffff021916908363ffffffff160217905550507f76336534418640ae4cd606c430d9287a4790c09310247a7a9fe90435f297acf98189898989428f604051610e2d97969594939291906134ec565b60405180910390a180915050979650505050505050565b6000600660019054906101000a900460ff16159050808015610e7857506001600660009054906101000a900460ff1660ff16105b80610ea75750610e8730611b47565b158015610ea657506001600660009054906101000a900460ff1660ff16145b5b610ee6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610edd906135d4565b60405180910390fd5b6001600660006101000a81548160ff021916908360ff1602179055508015610f24576001600660016101000a81548160ff0219169083151502179055505b8360018190555082600281905550816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550610f7a611b6a565b610f82611bc3565b8015610fdc576000600660016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024986001604051610fd39190613639565b60405180910390a15b50505050565b83610feb61167d565b83836000600360008681526020019081526020016000209050600061100e61167d565b60405160200161101e9190612fc2565b60405160208183030381529060405280519060200120905060006110488585856004015485611685565b90508061108a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110819061303a565b60405180910390fd5b6110a58b338c60405180602001604052806000815250611449565b5050505050505050505050565b6110ba61138c565b80600254036110fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110f5906136a0565b60405180910390fd5b8060028190555050565b61111061138c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361117f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161117690613732565b60405180910390fd5b61118881611a52565b50565b61119361138c565b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1603611221576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112189061379e565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008061127083610750565b90506005600881111561128657611285612c3d565b5b81600881111561129957611298612c3d565b5b14806112c95750600660088111156112b4576112b3612c3d565b5b8160088111156112c7576112c6612c3d565b5b145b611308576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112ff90613830565b60405180910390fd5b60016005600085815260200190815260200160002060020160006101000a81548160ff021916908315150217905550600061134284611c14565b90508015611382577f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f846040516113799190612932565b60405180910390a15b8092505050919050565b61139461167d565b73ffffffffffffffffffffffffffffffffffffffff166113b2610bd2565b73ffffffffffffffffffffffffffffffffffffffff1614611408576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113ff9061389c565b60405180910390fd5b565b60008061141684611d02565b9050600061142386611db6565b905060006114318383611e3e565b905061143d8186611e71565b93505050509392505050565b6002600881111561145d5761145c612c3d565b5b61146685610750565b600881111561147857611477612c3d565b5b146114b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114af9061392e565b60405180910390fd5b60006004600086815260200190815260200160002090506114d98585610941565b15611519576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611510906139c0565b60405180910390fd5b61152f8482600101611ee090919063ffffffff16565b5060018360ff16036115825760018160000160008282829054906101000a900463ffffffff1661155f91906139e0565b92506101000a81548163ffffffff021916908363ffffffff160217905550611624565b60028360ff16036115d45760018160000160048282829054906101000a900463ffffffff166115b191906139e0565b92506101000a81548163ffffffff021916908363ffffffff160217905550611623565b60038360ff16036116225760018160000160088282829054906101000a900463ffffffff1661160391906139e0565b92506101000a81548163ffffffff021916908363ffffffff1602179055505b5b5b8373ffffffffffffffffffffffffffffffffffffffff167f5be2678cab71c7d94235767eb1df3c57673f19e805cdc0edbd2e74854ae9f0f986858560405161166e93929190613a27565b60405180910390a25050505050565b600033905090565b600082611693868685611f10565b149050949350505050565b60006116cc7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611f68565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6116fd61138c565b61170561169e565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603611772576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161176990613ad7565b60405180910390fd5b50565b6117a17f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b611f72565b60000160009054906101000a900460ff16156117c5576117c083611f7c565b6118de565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561182d57506040513d601f19601f8201168201806040525081019061182a9190613b0c565b60015b61186c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161186390613bab565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b81146118d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118c890613c3d565b60405180910390fd5b506118dd838383612035565b5b505050565b600081600001519050919050565b6000806004600084815260200190815260200160002090508060000160089054906101000a900463ffffffff168160000160009054906101000a900463ffffffff168260000160049054906101000a900463ffffffff1661195291906139e0565b61195c91906139e0565b63ffffffff1681600001600c9054906101000a900463ffffffff1663ffffffff1614915050919050565b600061199182612061565b80156119ab575042826000015167ffffffffffffffff1611155b9050919050565b6000806004600084815260200190815260200160002090506001600282600001600c9054906101000a900463ffffffff166119ed9190613c8c565b6119f791906139e0565b63ffffffff168160000160049054906101000a900463ffffffff1663ffffffff161015915050919050565b6000611a4a836000018373ffffffffffffffffffffffffffffffffffffffff1660001b61207b565b905092915050565b6000603960009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081603960006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b808260000160006101000a81548167ffffffffffffffff021916908367ffffffffffffffff1602179055505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600660019054906101000a900460ff16611bb9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bb090613d2f565b60405180910390fd5b611bc161209e565b565b600660019054906101000a900460ff16611c12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611c0990613d2f565b60405180910390fd5b565b6000806003600084815260200190815260200160002090506000806001811115611c4157611c40612c3d565b5b8260050160009054906101000a900460ff166001811115611c6557611c64612c3d565b5b1480611ca55750600180811115611c7f57611c7e612c3d565b5b8260050160009054906101000a900460ff166001811115611ca357611ca2612c3d565b5b145b15611cf857611cf28260030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168360030160149054906101000a900463ffffffff1663ffffffff166120ff565b50600190505b8092505050919050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f8260000151604051602001611d3a9190613d8b565b604051602081830303815290604052805190602001208360200151604051602001611d659190613d8b565b6040516020818303038152906040528051906020012084604001518560600151604051602001611d99959493929190613da2565b604051602081830303815290604052805190602001209050919050565b60007f6677d4b4933cc0bf8a08b99aac6d774d9bc0e7d9ed28a3a3cdbd139470842b41826000015183602001518460400151604051602001611df89190613d8b565b60405160208183030381529060405280519060200120604051602001611e219493929190613df5565b604051602081830303815290604052805190602001209050919050565b60008282604051602001611e53929190613ea7565b60405160208183030381529060405280519060200120905092915050565b600080600080611e80856121bb565b92509250925060018682858560405160008152602001604052604051611ea99493929190613ede565b6020604051602081039080840390855afa158015611ecb573d6000803e3d6000fd5b50505060206040510351935050505092915050565b6000611f08836000018373ffffffffffffffffffffffffffffffffffffffff1660001b612223565b905092915050565b60008082905060005b85859050811015611f5c57611f4782878784818110611f3b57611f3a613f23565b5b90506020020135612293565b91508080611f5490613f52565b915050611f19565b50809150509392505050565b6000819050919050565b6000819050919050565b611f8581611b47565b611fc4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fbb9061400c565b60405180910390fd5b80611ff17f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b611f68565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b61203e836122be565b60008251118061204b5750805b1561205c5761205a838361230d565b505b505050565b600080826000015167ffffffffffffffff16119050919050565b600080836001016000848152602001908152602001600020541415905092915050565b600660019054906101000a900460ff166120ed576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120e490613d2f565b60405180910390fd5b6120fd6120f861167d565b611a52565b565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84670de0b6b3a76400008561215391906133dd565b6040518363ffffffff1660e01b815260040161217092919061402c565b6020604051808303816000875af115801561218f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121b39190614081565b905092915050565b60008060006041845114612204576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121fb906140fa565b60405180910390fd5b6020840151925060408401519150606084015160001a90509193909250565b600061222f838361207b565b61228857826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905061228d565b600090505b92915050565b60008183106122ab576122a682846123f1565b6122b6565b6122b583836123f1565b5b905092915050565b6122c781611f7c565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b606061231883611b47565b612357576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161234e9061418c565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff168460405161237f91906141f3565b600060405180830381855af49150503d80600081146123ba576040519150601f19603f3d011682016040523d82523d6000602084013e6123bf565b606091505b50915091506123e7828260405180606001604052806027815260200161422d60279139612408565b9250505092915050565b600082600052816020526040600020905092915050565b6060831561241857829050612423565b612422838361242a565b5b9392505050565b60008251111561243d5781518083602001fd5b806040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612471919061420a565b60405180910390fd5b6000604051905090565b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6124dc82612493565b810181811067ffffffffffffffff821117156124fb576124fa6124a4565b5b80604052505050565b600061250e61247a565b905061251a82826124d3565b919050565b600080fd5b6000819050919050565b61253781612524565b811461254257600080fd5b50565b6000813590506125548161252e565b92915050565b600060ff82169050919050565b6125708161255a565b811461257b57600080fd5b50565b60008135905061258d81612567565b92915050565b600080fd5b600080fd5b600067ffffffffffffffff8211156125b8576125b76124a4565b5b6125c182612493565b9050602081019050919050565b82818337600083830152505050565b60006125f06125eb8461259d565b612504565b90508281526020810184848401111561260c5761260b612598565b5b6126178482856125ce565b509392505050565b600082601f83011261263457612633612593565b5b81356126448482602086016125dd565b91505092915050565b6000606082840312156126635761266261248e565b5b61266d6060612504565b9050600061267d84828501612545565b60008301525060206126918482850161257e565b602083015250604082013567ffffffffffffffff8111156126b5576126b461251f565b5b6126c18482850161261f565b60408301525092915050565b600067ffffffffffffffff8211156126e8576126e76124a4565b5b6126f182612493565b9050602081019050919050565b600061271161270c846126cd565b612504565b90508281526020810184848401111561272d5761272c612598565b5b6127388482856125ce565b509392505050565b600082601f83011261275557612754612593565b5b81356127658482602086016126fe565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006127998261276e565b9050919050565b6127a98161278e565b81146127b457600080fd5b50565b6000813590506127c6816127a0565b92915050565b6000608082840312156127e2576127e161248e565b5b6127ec6080612504565b9050600082013567ffffffffffffffff81111561280c5761280b61251f565b5b6128188482850161261f565b600083015250602082013567ffffffffffffffff81111561283c5761283b61251f565b5b6128488482850161261f565b602083015250604061285c84828501612545565b6040830152506060612870848285016127b7565b60608301525092915050565b60008060006060848603121561289557612894612484565b5b600084013567ffffffffffffffff8111156128b3576128b2612489565b5b6128bf8682870161264d565b935050602084013567ffffffffffffffff8111156128e0576128df612489565b5b6128ec86828701612740565b925050604084013567ffffffffffffffff81111561290d5761290c612489565b5b612919868287016127cc565b9150509250925092565b61292c81612524565b82525050565b60006020820190506129476000830184612923565b92915050565b600080fd5b600080fd5b60008083601f84011261296d5761296c612593565b5b8235905067ffffffffffffffff81111561298a5761298961294d565b5b6020830191508360018202830111156129a6576129a5612952565b5b9250929050565b60008083601f8401126129c3576129c2612593565b5b8235905067ffffffffffffffff8111156129e0576129df61294d565b5b6020830191508360208202830111156129fc576129fb612952565b5b9250929050565b60008060008060008060808789031215612a2057612a1f612484565b5b6000612a2e89828a01612545565b9650506020612a3f89828a0161257e565b955050604087013567ffffffffffffffff811115612a6057612a5f612489565b5b612a6c89828a01612957565b9450945050606087013567ffffffffffffffff811115612a8f57612a8e612489565b5b612a9b89828a016129ad565b92509250509295509295509295565b600060208284031215612ac057612abf612484565b5b6000612ace848285016127b7565b91505092915050565b6000819050919050565b612aea81612ad7565b8114612af557600080fd5b50565b600081359050612b0781612ae1565b92915050565b600063ffffffff82169050919050565b612b2681612b0d565b8114612b3157600080fd5b50565b600081359050612b4381612b1d565b92915050565b60028110612b5657600080fd5b50565b600081359050612b6881612b49565b92915050565b600080600080600080600060e0888a031215612b8d57612b8c612484565b5b6000612b9b8a828b01612af8565b9750506020612bac8a828b016127b7565b9650506040612bbd8a828b01612af8565b9550506060612bce8a828b016127b7565b9450506080612bdf8a828b01612b34565b93505060a0612bf08a828b01612af8565b92505060c0612c018a828b01612b59565b91505092959891949750929550565b600060208284031215612c2657612c25612484565b5b6000612c3484828501612545565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60098110612c7d57612c7c612c3d565b5b50565b6000819050612c8e82612c6c565b919050565b6000612c9e82612c80565b9050919050565b612cae81612c93565b82525050565b6000602082019050612cc96000830184612ca5565b92915050565b60008060408385031215612ce657612ce5612484565b5b6000612cf485828601612545565b9250506020612d05858286016127b7565b9150509250929050565b60008115159050919050565b612d2481612d0f565b82525050565b6000602082019050612d3f6000830184612d1b565b92915050565b60008060408385031215612d5c57612d5b612484565b5b6000612d6a858286016127b7565b925050602083013567ffffffffffffffff811115612d8b57612d8a612489565b5b612d9785828601612740565b9150509250929050565b612daa81612ad7565b82525050565b6000602082019050612dc56000830184612da1565b92915050565b612dd48161278e565b82525050565b6000602082019050612def6000830184612dcb565b92915050565b600080600080600080600060e0888a031215612e1457612e13612484565b5b600088013567ffffffffffffffff811115612e3257612e31612489565b5b612e3e8a828b0161261f565b9750506020612e4f8a828b01612af8565b9650506040612e608a828b016127b7565b9550506060612e718a828b01612b34565b9450506080612e828a828b01612af8565b93505060a0612e938a828b01612b34565b92505060c0612ea48a828b01612b59565b91505092959891949750929550565b600080600060608486031215612ecc57612ecb612484565b5b6000612eda86828701612545565b9350506020612eeb86828701612545565b9250506040612efc868287016127b7565b9150509250925092565b60008060008060608587031215612f2057612f1f612484565b5b6000612f2e87828801612545565b9450506020612f3f8782880161257e565b935050604085013567ffffffffffffffff811115612f6057612f5f612489565b5b612f6c878288016129ad565b925092505092959194509250565b60008160601b9050919050565b6000612f9282612f7a565b9050919050565b6000612fa482612f87565b9050919050565b612fbc612fb78261278e565b612f99565b82525050565b6000612fce8284612fab565b60148201915081905092915050565b600082825260208201905092915050565b7f63617374566f74653a20496e76616c69642070726f6f662e0000000000000000600082015250565b6000613024601883612fdd565b915061302f82612fee565b602082019050919050565b6000602082019050818103600083015261305381613017565b9050919050565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b60006130b6602c83612fdd565b91506130c18261305a565b604082019050919050565b600060208201905081810360008301526130e5816130a9565b9050919050565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b6000613148602c83612fdd565b9150613153826130ec565b604082019050919050565b600060208201905081810360008301526131778161313b565b9050919050565b61318781612b0d565b82525050565b6002811061319e5761319d612c3d565b5b50565b60008190506131af8261318d565b919050565b60006131bf826131a1565b9050919050565b6131cf816131b4565b82525050565b600060e0820190506131ea600083018a612da1565b6131f76020830189612dcb565b6132046040830188612da1565b6132116060830187612dcb565b61321e608083018661317e565b61322b60a0830185612da1565b61323860c08301846131c6565b98975050505050505050565b7f476f7665726e6f723a20756e6b6e6f776e2070726f706f73616c206964000000600082015250565b600061327a601d83612fdd565b915061328582613244565b602082019050919050565b600060208201905081810360008301526132a98161326d565b9050919050565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b600061330c603883612fdd565b9150613317826132b0565b604082019050919050565b6000602082019050818103600083015261333b816132ff565b9050919050565b7f496c6c6567616c2076616c756520666f7220766f74696e672064656c61792100600082015250565b6000613378601f83612fdd565b915061338382613342565b602082019050919050565b600060208201905081810360008301526133a78161336b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006133e882612524565b91506133f383612524565b925082820261340181612524565b91508282048414831517613418576134176133ae565b5b5092915050565b600067ffffffffffffffff82169050919050565b600061343e8261341f565b91506134498361341f565b9250828201905067ffffffffffffffff811115613469576134686133ae565b5b92915050565b6134788161341f565b82525050565b600081519050919050565b60005b838110156134a757808201518184015260208101905061348c565b60008484015250505050565b60006134be8261347e565b6134c88185612fdd565b93506134d8818560208601613489565b6134e181612493565b840191505092915050565b600060e082019050613501600083018a612923565b61350e6020830189612da1565b61351b6040830188612dcb565b613528606083018761317e565b6135356080830186612da1565b61354260a083018561346f565b81810360c083015261355481846134b3565b905098975050505050505050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b60006135be602e83612fdd565b91506135c982613562565b604082019050919050565b600060208201905081810360008301526135ed816135b1565b9050919050565b6000819050919050565b6000819050919050565b600061362361361e613619846135f4565b6135fe565b61255a565b9050919050565b61363381613608565b82525050565b600060208201905061364e600083018461362a565b92915050565b7f496c6c6567616c2076616c756520666f7220766f74696e6720706572696f6421600082015250565b600061368a602083612fdd565b915061369582613654565b602082019050919050565b600060208201905081810360008301526136b98161367d565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b600061371c602683612fdd565b9150613727826136c0565b604082019050919050565b6000602082019050818103600083015261374b8161370f565b9050919050565b7f496c6c6567616c2076616c756520666f722072657761726420746f6b656e2100600082015250565b6000613788601f83612fdd565b915061379382613752565b602082019050919050565b600060208201905081810360008301526137b78161377b565b9050919050565b7f476f7665726e6f723a2070726f706f73616c206e6f742073756363657373667560008201527f6c00000000000000000000000000000000000000000000000000000000000000602082015250565b600061381a602183612fdd565b9150613825826137be565b604082019050919050565b600060208201905081810360008301526138498161380d565b9050919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b6000613886602083612fdd565b915061389182613850565b602082019050919050565b600060208201905081810360008301526138b581613879565b9050919050565b7f476f7665726e6f723a20766f7465206e6f742063757272656e746c792061637460008201527f6976650000000000000000000000000000000000000000000000000000000000602082015250565b6000613918602383612fdd565b9150613923826138bc565b604082019050919050565b600060208201905081810360008301526139478161390b565b9050919050565b7f476f7665726e6f72566f74696e6753696d706c653a20766f746520616c72656160008201527f6479206361737400000000000000000000000000000000000000000000000000602082015250565b60006139aa602783612fdd565b91506139b58261394e565b604082019050919050565b600060208201905081810360008301526139d98161399d565b9050919050565b60006139eb82612b0d565b91506139f683612b0d565b9250828201905063ffffffff811115613a1257613a116133ae565b5b92915050565b613a218161255a565b82525050565b6000606082019050613a3c6000830186612923565b613a496020830185613a18565b8181036040830152613a5b81846134b3565b9050949350505050565b7f50726f78793a20496c6c6567616c204e657720496d706c656d656e746174696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000613ac1602183612fdd565b9150613acc82613a65565b604082019050919050565b60006020820190508181036000830152613af081613ab4565b9050919050565b600081519050613b0681612ae1565b92915050565b600060208284031215613b2257613b21612484565b5b6000613b3084828501613af7565b91505092915050565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b6000613b95602e83612fdd565b9150613ba082613b39565b604082019050919050565b60006020820190508181036000830152613bc481613b88565b9050919050565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b6000613c27602983612fdd565b9150613c3282613bcb565b604082019050919050565b60006020820190508181036000830152613c5681613c1a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000613c9782612b0d565b9150613ca283612b0d565b925082613cb257613cb1613c5d565b5b828204905092915050565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b6000613d19602b83612fdd565b9150613d2482613cbd565b604082019050919050565b60006020820190508181036000830152613d4881613d0c565b9050919050565b600081905092915050565b6000613d658261347e565b613d6f8185613d4f565b9350613d7f818560208601613489565b80840191505092915050565b6000613d978284613d5a565b915081905092915050565b600060a082019050613db76000830188612da1565b613dc46020830187612da1565b613dd16040830186612da1565b613dde6060830185612923565b613deb6080830184612dcb565b9695505050505050565b6000608082019050613e0a6000830187612da1565b613e176020830186612923565b613e246040830185613a18565b613e316060830184612da1565b95945050505050565b7f1901000000000000000000000000000000000000000000000000000000000000600082015250565b6000613e70600283613d4f565b9150613e7b82613e3a565b600282019050919050565b6000819050919050565b613ea1613e9c82612ad7565b613e86565b82525050565b6000613eb282613e63565b9150613ebe8285613e90565b602082019150613ece8284613e90565b6020820191508190509392505050565b6000608082019050613ef36000830187612da1565b613f006020830186613a18565b613f0d6040830185612da1565b613f1a6060830184612da1565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000613f5d82612524565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203613f8f57613f8e6133ae565b5b600182019050919050565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b6000613ff6602d83612fdd565b915061400182613f9a565b604082019050919050565b6000602082019050818103600083015261402581613fe9565b9050919050565b60006040820190506140416000830185612dcb565b61404e6020830184612923565b9392505050565b61405e81612d0f565b811461406957600080fd5b50565b60008151905061407b81614055565b92915050565b60006020828403121561409757614096612484565b5b60006140a58482850161406c565b91505092915050565b7f696e76616c6964207369676e6174757265206c656e6774680000000000000000600082015250565b60006140e4601883612fdd565b91506140ef826140ae565b602082019050919050565b60006020820190508181036000830152614113816140d7565b9050919050565b7f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60008201527f6e74726163740000000000000000000000000000000000000000000000000000602082015250565b6000614176602683612fdd565b91506141818261411a565b604082019050919050565b600060208201905081810360008301526141a581614169565b9050919050565b600081519050919050565b600081905092915050565b60006141cd826141ac565b6141d781856141b7565b93506141e7818560208601613489565b80840191505092915050565b60006141ff82846141c2565b915081905092915050565b6000602082019050818103600083015261422481846134b3565b90509291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220053566f29e7c0628c42a4902ecb1b6beccd09af64e1f09508ce08b40584756ec64736f6c63430008130033",
}

// StrikeGovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use StrikeGovernorMetaData.ABI instead.
var StrikeGovernorABI = StrikeGovernorMetaData.ABI

// StrikeGovernorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrikeGovernorMetaData.Bin instead.
var StrikeGovernorBin = StrikeGovernorMetaData.Bin

// DeployStrikeGovernor deploys a new Ethereum contract, binding an instance of StrikeGovernor to it.
func DeployStrikeGovernor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StrikeGovernor, error) {
	parsed, err := StrikeGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrikeGovernorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrikeGovernor{StrikeGovernorCaller: StrikeGovernorCaller{contract: contract}, StrikeGovernorTransactor: StrikeGovernorTransactor{contract: contract}, StrikeGovernorFilterer: StrikeGovernorFilterer{contract: contract}}, nil
}

// StrikeGovernor is an auto generated Go binding around an Ethereum contract.
type StrikeGovernor struct {
	StrikeGovernorCaller     // Read-only binding to the contract
	StrikeGovernorTransactor // Write-only binding to the contract
	StrikeGovernorFilterer   // Log filterer for contract events
}

// StrikeGovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrikeGovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrikeGovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrikeGovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrikeGovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrikeGovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrikeGovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrikeGovernorSession struct {
	Contract     *StrikeGovernor   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StrikeGovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrikeGovernorCallerSession struct {
	Contract *StrikeGovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// StrikeGovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrikeGovernorTransactorSession struct {
	Contract     *StrikeGovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// StrikeGovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrikeGovernorRaw struct {
	Contract *StrikeGovernor // Generic contract binding to access the raw methods on
}

// StrikeGovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrikeGovernorCallerRaw struct {
	Contract *StrikeGovernorCaller // Generic read-only contract binding to access the raw methods on
}

// StrikeGovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrikeGovernorTransactorRaw struct {
	Contract *StrikeGovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrikeGovernor creates a new instance of StrikeGovernor, bound to a specific deployed contract.
func NewStrikeGovernor(address common.Address, backend bind.ContractBackend) (*StrikeGovernor, error) {
	contract, err := bindStrikeGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernor{StrikeGovernorCaller: StrikeGovernorCaller{contract: contract}, StrikeGovernorTransactor: StrikeGovernorTransactor{contract: contract}, StrikeGovernorFilterer: StrikeGovernorFilterer{contract: contract}}, nil
}

// NewStrikeGovernorCaller creates a new read-only instance of StrikeGovernor, bound to a specific deployed contract.
func NewStrikeGovernorCaller(address common.Address, caller bind.ContractCaller) (*StrikeGovernorCaller, error) {
	contract, err := bindStrikeGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorCaller{contract: contract}, nil
}

// NewStrikeGovernorTransactor creates a new write-only instance of StrikeGovernor, bound to a specific deployed contract.
func NewStrikeGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*StrikeGovernorTransactor, error) {
	contract, err := bindStrikeGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorTransactor{contract: contract}, nil
}

// NewStrikeGovernorFilterer creates a new log filterer instance of StrikeGovernor, bound to a specific deployed contract.
func NewStrikeGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*StrikeGovernorFilterer, error) {
	contract, err := bindStrikeGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorFilterer{contract: contract}, nil
}

// bindStrikeGovernor binds a generic wrapper to an already deployed contract.
func bindStrikeGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrikeGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrikeGovernor *StrikeGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrikeGovernor.Contract.StrikeGovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrikeGovernor *StrikeGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.StrikeGovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrikeGovernor *StrikeGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.StrikeGovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrikeGovernor *StrikeGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrikeGovernor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrikeGovernor *StrikeGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrikeGovernor *StrikeGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.contract.Transact(opts, method, params...)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_StrikeGovernor *StrikeGovernorCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _StrikeGovernor.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_StrikeGovernor *StrikeGovernorSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _StrikeGovernor.Contract.HasVoted(&_StrikeGovernor.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_StrikeGovernor *StrikeGovernorCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _StrikeGovernor.Contract.HasVoted(&_StrikeGovernor.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0x3a4b0e3b.
//
// Solidity: function hashProposal(bytes32 descriptionHash, address proposer, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint8 proposalType) pure returns(uint256)
func (_StrikeGovernor *StrikeGovernorCaller) HashProposal(opts *bind.CallOpts, descriptionHash [32]byte, proposer common.Address, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, proposalType uint8) (*big.Int, error) {
	var out []interface{}
	err := _StrikeGovernor.contract.Call(opts, &out, "hashProposal", descriptionHash, proposer, candidateID, candidate, tokenOffer, merkleRoot, proposalType)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0x3a4b0e3b.
//
// Solidity: function hashProposal(bytes32 descriptionHash, address proposer, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint8 proposalType) pure returns(uint256)
func (_StrikeGovernor *StrikeGovernorSession) HashProposal(descriptionHash [32]byte, proposer common.Address, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, proposalType uint8) (*big.Int, error) {
	return _StrikeGovernor.Contract.HashProposal(&_StrikeGovernor.CallOpts, descriptionHash, proposer, candidateID, candidate, tokenOffer, merkleRoot, proposalType)
}

// HashProposal is a free data retrieval call binding the contract method 0x3a4b0e3b.
//
// Solidity: function hashProposal(bytes32 descriptionHash, address proposer, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint8 proposalType) pure returns(uint256)
func (_StrikeGovernor *StrikeGovernorCallerSession) HashProposal(descriptionHash [32]byte, proposer common.Address, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, proposalType uint8) (*big.Int, error) {
	return _StrikeGovernor.Contract.HashProposal(&_StrikeGovernor.CallOpts, descriptionHash, proposer, candidateID, candidate, tokenOffer, merkleRoot, proposalType)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrikeGovernor *StrikeGovernorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StrikeGovernor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrikeGovernor *StrikeGovernorSession) Owner() (common.Address, error) {
	return _StrikeGovernor.Contract.Owner(&_StrikeGovernor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StrikeGovernor *StrikeGovernorCallerSession) Owner() (common.Address, error) {
	return _StrikeGovernor.Contract.Owner(&_StrikeGovernor.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StrikeGovernor *StrikeGovernorCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StrikeGovernor.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StrikeGovernor *StrikeGovernorSession) ProxiableUUID() ([32]byte, error) {
	return _StrikeGovernor.Contract.ProxiableUUID(&_StrikeGovernor.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StrikeGovernor *StrikeGovernorCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StrikeGovernor.Contract.ProxiableUUID(&_StrikeGovernor.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_StrikeGovernor *StrikeGovernorCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _StrikeGovernor.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_StrikeGovernor *StrikeGovernorSession) State(proposalId *big.Int) (uint8, error) {
	return _StrikeGovernor.Contract.State(&_StrikeGovernor.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_StrikeGovernor *StrikeGovernorCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _StrikeGovernor.Contract.State(&_StrikeGovernor.CallOpts, proposalId)
}

// CastVote is a paid mutator transaction binding the contract method 0xd92202f8.
//
// Solidity: function castVote(uint256 proposalId, uint8 support, bytes32[] proof) returns()
func (_StrikeGovernor *StrikeGovernorTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8, proof [][32]byte) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "castVote", proposalId, support, proof)
}

// CastVote is a paid mutator transaction binding the contract method 0xd92202f8.
//
// Solidity: function castVote(uint256 proposalId, uint8 support, bytes32[] proof) returns()
func (_StrikeGovernor *StrikeGovernorSession) CastVote(proposalId *big.Int, support uint8, proof [][32]byte) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.CastVote(&_StrikeGovernor.TransactOpts, proposalId, support, proof)
}

// CastVote is a paid mutator transaction binding the contract method 0xd92202f8.
//
// Solidity: function castVote(uint256 proposalId, uint8 support, bytes32[] proof) returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) CastVote(proposalId *big.Int, support uint8, proof [][32]byte) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.CastVote(&_StrikeGovernor.TransactOpts, proposalId, support, proof)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x02ce3b1a.
//
// Solidity: function castVoteBySig((uint256,uint8,string) message, bytes signature, (string,string,uint256,address) domain) returns(uint256 balance)
func (_StrikeGovernor *StrikeGovernorTransactor) CastVoteBySig(opts *bind.TransactOpts, message IMessageMessage, signature []byte, domain IMessageDomainSeparator) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "castVoteBySig", message, signature, domain)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x02ce3b1a.
//
// Solidity: function castVoteBySig((uint256,uint8,string) message, bytes signature, (string,string,uint256,address) domain) returns(uint256 balance)
func (_StrikeGovernor *StrikeGovernorSession) CastVoteBySig(message IMessageMessage, signature []byte, domain IMessageDomainSeparator) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.CastVoteBySig(&_StrikeGovernor.TransactOpts, message, signature, domain)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x02ce3b1a.
//
// Solidity: function castVoteBySig((uint256,uint8,string) message, bytes signature, (string,string,uint256,address) domain) returns(uint256 balance)
func (_StrikeGovernor *StrikeGovernorTransactorSession) CastVoteBySig(message IMessageMessage, signature []byte, domain IMessageDomainSeparator) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.CastVoteBySig(&_StrikeGovernor.TransactOpts, message, signature, domain)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x1d46e7b9.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason, bytes32[] proof) returns()
func (_StrikeGovernor *StrikeGovernorTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, proof [][32]byte) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason, proof)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x1d46e7b9.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason, bytes32[] proof) returns()
func (_StrikeGovernor *StrikeGovernorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string, proof [][32]byte) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.CastVoteWithReason(&_StrikeGovernor.TransactOpts, proposalId, support, reason, proof)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x1d46e7b9.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason, bytes32[] proof) returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string, proof [][32]byte) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.CastVoteWithReason(&_StrikeGovernor.TransactOpts, proposalId, support, reason, proof)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns(bool)
func (_StrikeGovernor *StrikeGovernorTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "execute", proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns(bool)
func (_StrikeGovernor *StrikeGovernorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.Execute(&_StrikeGovernor.TransactOpts, proposalId)
}

// Execute is a paid mutator transaction binding the contract method 0xfe0d94c1.
//
// Solidity: function execute(uint256 proposalId) payable returns(bool)
func (_StrikeGovernor *StrikeGovernorTransactorSession) Execute(proposalId *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.Execute(&_StrikeGovernor.TransactOpts, proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 _votingDelay, uint256 _votingPeriod, address _rewardToken) returns()
func (_StrikeGovernor *StrikeGovernorTransactor) Initialize(opts *bind.TransactOpts, _votingDelay *big.Int, _votingPeriod *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "initialize", _votingDelay, _votingPeriod, _rewardToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 _votingDelay, uint256 _votingPeriod, address _rewardToken) returns()
func (_StrikeGovernor *StrikeGovernorSession) Initialize(_votingDelay *big.Int, _votingPeriod *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.Initialize(&_StrikeGovernor.TransactOpts, _votingDelay, _votingPeriod, _rewardToken)
}

// Initialize is a paid mutator transaction binding the contract method 0xa6ab36f2.
//
// Solidity: function initialize(uint256 _votingDelay, uint256 _votingPeriod, address _rewardToken) returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) Initialize(_votingDelay *big.Int, _votingPeriod *big.Int, _rewardToken common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.Initialize(&_StrikeGovernor.TransactOpts, _votingDelay, _votingPeriod, _rewardToken)
}

// Propose is a paid mutator transaction binding the contract method 0xa618b8c1.
//
// Solidity: function propose(string description, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint32 validatorCnt, uint8 proposalType) returns(uint256)
func (_StrikeGovernor *StrikeGovernorTransactor) Propose(opts *bind.TransactOpts, description string, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, validatorCnt uint32, proposalType uint8) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "propose", description, candidateID, candidate, tokenOffer, merkleRoot, validatorCnt, proposalType)
}

// Propose is a paid mutator transaction binding the contract method 0xa618b8c1.
//
// Solidity: function propose(string description, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint32 validatorCnt, uint8 proposalType) returns(uint256)
func (_StrikeGovernor *StrikeGovernorSession) Propose(description string, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, validatorCnt uint32, proposalType uint8) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.Propose(&_StrikeGovernor.TransactOpts, description, candidateID, candidate, tokenOffer, merkleRoot, validatorCnt, proposalType)
}

// Propose is a paid mutator transaction binding the contract method 0xa618b8c1.
//
// Solidity: function propose(string description, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint32 validatorCnt, uint8 proposalType) returns(uint256)
func (_StrikeGovernor *StrikeGovernorTransactorSession) Propose(description string, candidateID [32]byte, candidate common.Address, tokenOffer uint32, merkleRoot [32]byte, validatorCnt uint32, proposalType uint8) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.Propose(&_StrikeGovernor.TransactOpts, description, candidateID, candidate, tokenOffer, merkleRoot, validatorCnt, proposalType)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrikeGovernor *StrikeGovernorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrikeGovernor *StrikeGovernorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrikeGovernor.Contract.RenounceOwnership(&_StrikeGovernor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StrikeGovernor.Contract.RenounceOwnership(&_StrikeGovernor.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrikeGovernor *StrikeGovernorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrikeGovernor *StrikeGovernorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.TransferOwnership(&_StrikeGovernor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.TransferOwnership(&_StrikeGovernor.TransactOpts, newOwner)
}

// UpdateRewardToken is a paid mutator transaction binding the contract method 0xf8cf31cb.
//
// Solidity: function updateRewardToken(address newValue) returns()
func (_StrikeGovernor *StrikeGovernorTransactor) UpdateRewardToken(opts *bind.TransactOpts, newValue common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "updateRewardToken", newValue)
}

// UpdateRewardToken is a paid mutator transaction binding the contract method 0xf8cf31cb.
//
// Solidity: function updateRewardToken(address newValue) returns()
func (_StrikeGovernor *StrikeGovernorSession) UpdateRewardToken(newValue common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpdateRewardToken(&_StrikeGovernor.TransactOpts, newValue)
}

// UpdateRewardToken is a paid mutator transaction binding the contract method 0xf8cf31cb.
//
// Solidity: function updateRewardToken(address newValue) returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) UpdateRewardToken(newValue common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpdateRewardToken(&_StrikeGovernor.TransactOpts, newValue)
}

// UpdateVotingDelay is a paid mutator transaction binding the contract method 0x63d61a19.
//
// Solidity: function updateVotingDelay(uint256 newValue) returns()
func (_StrikeGovernor *StrikeGovernorTransactor) UpdateVotingDelay(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "updateVotingDelay", newValue)
}

// UpdateVotingDelay is a paid mutator transaction binding the contract method 0x63d61a19.
//
// Solidity: function updateVotingDelay(uint256 newValue) returns()
func (_StrikeGovernor *StrikeGovernorSession) UpdateVotingDelay(newValue *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpdateVotingDelay(&_StrikeGovernor.TransactOpts, newValue)
}

// UpdateVotingDelay is a paid mutator transaction binding the contract method 0x63d61a19.
//
// Solidity: function updateVotingDelay(uint256 newValue) returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) UpdateVotingDelay(newValue *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpdateVotingDelay(&_StrikeGovernor.TransactOpts, newValue)
}

// UpdateVotingPeriod is a paid mutator transaction binding the contract method 0xef00ef43.
//
// Solidity: function updateVotingPeriod(uint256 newValue) returns()
func (_StrikeGovernor *StrikeGovernorTransactor) UpdateVotingPeriod(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "updateVotingPeriod", newValue)
}

// UpdateVotingPeriod is a paid mutator transaction binding the contract method 0xef00ef43.
//
// Solidity: function updateVotingPeriod(uint256 newValue) returns()
func (_StrikeGovernor *StrikeGovernorSession) UpdateVotingPeriod(newValue *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpdateVotingPeriod(&_StrikeGovernor.TransactOpts, newValue)
}

// UpdateVotingPeriod is a paid mutator transaction binding the contract method 0xef00ef43.
//
// Solidity: function updateVotingPeriod(uint256 newValue) returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) UpdateVotingPeriod(newValue *big.Int) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpdateVotingPeriod(&_StrikeGovernor.TransactOpts, newValue)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StrikeGovernor *StrikeGovernorTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StrikeGovernor *StrikeGovernorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpgradeTo(&_StrikeGovernor.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpgradeTo(&_StrikeGovernor.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StrikeGovernor *StrikeGovernorTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StrikeGovernor.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StrikeGovernor *StrikeGovernorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpgradeToAndCall(&_StrikeGovernor.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StrikeGovernor *StrikeGovernorTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StrikeGovernor.Contract.UpgradeToAndCall(&_StrikeGovernor.TransactOpts, newImplementation, data)
}

// StrikeGovernorAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the StrikeGovernor contract.
type StrikeGovernorAdminChangedIterator struct {
	Event *StrikeGovernorAdminChanged // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorAdminChanged)
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
		it.Event = new(StrikeGovernorAdminChanged)
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
func (it *StrikeGovernorAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorAdminChanged represents a AdminChanged event raised by the StrikeGovernor contract.
type StrikeGovernorAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*StrikeGovernorAdminChangedIterator, error) {

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorAdminChangedIterator{contract: _StrikeGovernor.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *StrikeGovernorAdminChanged) (event.Subscription, error) {

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorAdminChanged)
				if err := _StrikeGovernor.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseAdminChanged(log types.Log) (*StrikeGovernorAdminChanged, error) {
	event := new(StrikeGovernorAdminChanged)
	if err := _StrikeGovernor.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the StrikeGovernor contract.
type StrikeGovernorBeaconUpgradedIterator struct {
	Event *StrikeGovernorBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorBeaconUpgraded)
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
		it.Event = new(StrikeGovernorBeaconUpgraded)
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
func (it *StrikeGovernorBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorBeaconUpgraded represents a BeaconUpgraded event raised by the StrikeGovernor contract.
type StrikeGovernorBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*StrikeGovernorBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorBeaconUpgradedIterator{contract: _StrikeGovernor.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *StrikeGovernorBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorBeaconUpgraded)
				if err := _StrikeGovernor.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseBeaconUpgraded(log types.Log) (*StrikeGovernorBeaconUpgraded, error) {
	event := new(StrikeGovernorBeaconUpgraded)
	if err := _StrikeGovernor.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StrikeGovernor contract.
type StrikeGovernorInitializedIterator struct {
	Event *StrikeGovernorInitialized // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorInitialized)
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
		it.Event = new(StrikeGovernorInitialized)
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
func (it *StrikeGovernorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorInitialized represents a Initialized event raised by the StrikeGovernor contract.
type StrikeGovernorInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterInitialized(opts *bind.FilterOpts) (*StrikeGovernorInitializedIterator, error) {

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorInitializedIterator{contract: _StrikeGovernor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StrikeGovernorInitialized) (event.Subscription, error) {

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorInitialized)
				if err := _StrikeGovernor.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseInitialized(log types.Log) (*StrikeGovernorInitialized, error) {
	event := new(StrikeGovernorInitialized)
	if err := _StrikeGovernor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StrikeGovernor contract.
type StrikeGovernorOwnershipTransferredIterator struct {
	Event *StrikeGovernorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorOwnershipTransferred)
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
		it.Event = new(StrikeGovernorOwnershipTransferred)
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
func (it *StrikeGovernorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorOwnershipTransferred represents a OwnershipTransferred event raised by the StrikeGovernor contract.
type StrikeGovernorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StrikeGovernorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorOwnershipTransferredIterator{contract: _StrikeGovernor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StrikeGovernorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorOwnershipTransferred)
				if err := _StrikeGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseOwnershipTransferred(log types.Log) (*StrikeGovernorOwnershipTransferred, error) {
	event := new(StrikeGovernorOwnershipTransferred)
	if err := _StrikeGovernor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the StrikeGovernor contract.
type StrikeGovernorProposalCanceledIterator struct {
	Event *StrikeGovernorProposalCanceled // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorProposalCanceled)
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
		it.Event = new(StrikeGovernorProposalCanceled)
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
func (it *StrikeGovernorProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorProposalCanceled represents a ProposalCanceled event raised by the StrikeGovernor contract.
type StrikeGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*StrikeGovernorProposalCanceledIterator, error) {

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorProposalCanceledIterator{contract: _StrikeGovernor.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *StrikeGovernorProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorProposalCanceled)
				if err := _StrikeGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
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

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseProposalCanceled(log types.Log) (*StrikeGovernorProposalCanceled, error) {
	event := new(StrikeGovernorProposalCanceled)
	if err := _StrikeGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the StrikeGovernor contract.
type StrikeGovernorProposalCreatedIterator struct {
	Event *StrikeGovernorProposalCreated // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorProposalCreated)
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
		it.Event = new(StrikeGovernorProposalCreated)
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
func (it *StrikeGovernorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorProposalCreated represents a ProposalCreated event raised by the StrikeGovernor contract.
type StrikeGovernorProposalCreated struct {
	ProposalId     *big.Int
	CandidateID    [32]byte
	Candidate      common.Address
	TokenOffer     uint32
	MerkleRoot     [32]byte
	StartTimestamp uint64
	Description    string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x76336534418640ae4cd606c430d9287a4790c09310247a7a9fe90435f297acf9.
//
// Solidity: event ProposalCreated(uint256 proposalId, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint64 startTimestamp, string description)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*StrikeGovernorProposalCreatedIterator, error) {

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorProposalCreatedIterator{contract: _StrikeGovernor.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x76336534418640ae4cd606c430d9287a4790c09310247a7a9fe90435f297acf9.
//
// Solidity: event ProposalCreated(uint256 proposalId, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint64 startTimestamp, string description)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *StrikeGovernorProposalCreated) (event.Subscription, error) {

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorProposalCreated)
				if err := _StrikeGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x76336534418640ae4cd606c430d9287a4790c09310247a7a9fe90435f297acf9.
//
// Solidity: event ProposalCreated(uint256 proposalId, bytes32 candidateID, address candidate, uint32 tokenOffer, bytes32 merkleRoot, uint64 startTimestamp, string description)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseProposalCreated(log types.Log) (*StrikeGovernorProposalCreated, error) {
	event := new(StrikeGovernorProposalCreated)
	if err := _StrikeGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the StrikeGovernor contract.
type StrikeGovernorProposalExecutedIterator struct {
	Event *StrikeGovernorProposalExecuted // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorProposalExecuted)
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
		it.Event = new(StrikeGovernorProposalExecuted)
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
func (it *StrikeGovernorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorProposalExecuted represents a ProposalExecuted event raised by the StrikeGovernor contract.
type StrikeGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*StrikeGovernorProposalExecutedIterator, error) {

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorProposalExecutedIterator{contract: _StrikeGovernor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *StrikeGovernorProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorProposalExecuted)
				if err := _StrikeGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseProposalExecuted(log types.Log) (*StrikeGovernorProposalExecuted, error) {
	event := new(StrikeGovernorProposalExecuted)
	if err := _StrikeGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StrikeGovernor contract.
type StrikeGovernorUpgradedIterator struct {
	Event *StrikeGovernorUpgraded // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorUpgraded)
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
		it.Event = new(StrikeGovernorUpgraded)
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
func (it *StrikeGovernorUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorUpgraded represents a Upgraded event raised by the StrikeGovernor contract.
type StrikeGovernorUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StrikeGovernorUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorUpgradedIterator{contract: _StrikeGovernor.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StrikeGovernorUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorUpgraded)
				if err := _StrikeGovernor.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseUpgraded(log types.Log) (*StrikeGovernorUpgraded, error) {
	event := new(StrikeGovernorUpgraded)
	if err := _StrikeGovernor.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the StrikeGovernor contract.
type StrikeGovernorVoteCastIterator struct {
	Event *StrikeGovernorVoteCast // Event containing the contract specifics and raw log

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
func (it *StrikeGovernorVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StrikeGovernorVoteCast)
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
		it.Event = new(StrikeGovernorVoteCast)
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
func (it *StrikeGovernorVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StrikeGovernorVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StrikeGovernorVoteCast represents a VoteCast event raised by the StrikeGovernor contract.
type StrikeGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0x5be2678cab71c7d94235767eb1df3c57673f19e805cdc0edbd2e74854ae9f0f9.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, string reason)
func (_StrikeGovernor *StrikeGovernorFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*StrikeGovernorVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _StrikeGovernor.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorVoteCastIterator{contract: _StrikeGovernor.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0x5be2678cab71c7d94235767eb1df3c57673f19e805cdc0edbd2e74854ae9f0f9.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, string reason)
func (_StrikeGovernor *StrikeGovernorFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *StrikeGovernorVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _StrikeGovernor.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StrikeGovernorVoteCast)
				if err := _StrikeGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0x5be2678cab71c7d94235767eb1df3c57673f19e805cdc0edbd2e74854ae9f0f9.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, string reason)
func (_StrikeGovernor *StrikeGovernorFilterer) ParseVoteCast(log types.Log) (*StrikeGovernorVoteCast, error) {
	event := new(StrikeGovernorVoteCast)
	if err := _StrikeGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StrikeGovernorStorageMetaData contains all meta data concerning the StrikeGovernorStorage contract.
var StrikeGovernorStorageMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x6080604052348015600f57600080fd5b50603f80601d6000396000f3fe6080604052600080fdfea26469706673582212209abe8f49cb1774556e2762bd36a98a2f7201a2faa5a65214cc88cd609799b07d64736f6c63430008130033",
}

// StrikeGovernorStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use StrikeGovernorStorageMetaData.ABI instead.
var StrikeGovernorStorageABI = StrikeGovernorStorageMetaData.ABI

// StrikeGovernorStorageBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StrikeGovernorStorageMetaData.Bin instead.
var StrikeGovernorStorageBin = StrikeGovernorStorageMetaData.Bin

// DeployStrikeGovernorStorage deploys a new Ethereum contract, binding an instance of StrikeGovernorStorage to it.
func DeployStrikeGovernorStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StrikeGovernorStorage, error) {
	parsed, err := StrikeGovernorStorageMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StrikeGovernorStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StrikeGovernorStorage{StrikeGovernorStorageCaller: StrikeGovernorStorageCaller{contract: contract}, StrikeGovernorStorageTransactor: StrikeGovernorStorageTransactor{contract: contract}, StrikeGovernorStorageFilterer: StrikeGovernorStorageFilterer{contract: contract}}, nil
}

// StrikeGovernorStorage is an auto generated Go binding around an Ethereum contract.
type StrikeGovernorStorage struct {
	StrikeGovernorStorageCaller     // Read-only binding to the contract
	StrikeGovernorStorageTransactor // Write-only binding to the contract
	StrikeGovernorStorageFilterer   // Log filterer for contract events
}

// StrikeGovernorStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type StrikeGovernorStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrikeGovernorStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StrikeGovernorStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrikeGovernorStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StrikeGovernorStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StrikeGovernorStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StrikeGovernorStorageSession struct {
	Contract     *StrikeGovernorStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StrikeGovernorStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StrikeGovernorStorageCallerSession struct {
	Contract *StrikeGovernorStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// StrikeGovernorStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StrikeGovernorStorageTransactorSession struct {
	Contract     *StrikeGovernorStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// StrikeGovernorStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type StrikeGovernorStorageRaw struct {
	Contract *StrikeGovernorStorage // Generic contract binding to access the raw methods on
}

// StrikeGovernorStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StrikeGovernorStorageCallerRaw struct {
	Contract *StrikeGovernorStorageCaller // Generic read-only contract binding to access the raw methods on
}

// StrikeGovernorStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StrikeGovernorStorageTransactorRaw struct {
	Contract *StrikeGovernorStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrikeGovernorStorage creates a new instance of StrikeGovernorStorage, bound to a specific deployed contract.
func NewStrikeGovernorStorage(address common.Address, backend bind.ContractBackend) (*StrikeGovernorStorage, error) {
	contract, err := bindStrikeGovernorStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorStorage{StrikeGovernorStorageCaller: StrikeGovernorStorageCaller{contract: contract}, StrikeGovernorStorageTransactor: StrikeGovernorStorageTransactor{contract: contract}, StrikeGovernorStorageFilterer: StrikeGovernorStorageFilterer{contract: contract}}, nil
}

// NewStrikeGovernorStorageCaller creates a new read-only instance of StrikeGovernorStorage, bound to a specific deployed contract.
func NewStrikeGovernorStorageCaller(address common.Address, caller bind.ContractCaller) (*StrikeGovernorStorageCaller, error) {
	contract, err := bindStrikeGovernorStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorStorageCaller{contract: contract}, nil
}

// NewStrikeGovernorStorageTransactor creates a new write-only instance of StrikeGovernorStorage, bound to a specific deployed contract.
func NewStrikeGovernorStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*StrikeGovernorStorageTransactor, error) {
	contract, err := bindStrikeGovernorStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorStorageTransactor{contract: contract}, nil
}

// NewStrikeGovernorStorageFilterer creates a new log filterer instance of StrikeGovernorStorage, bound to a specific deployed contract.
func NewStrikeGovernorStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*StrikeGovernorStorageFilterer, error) {
	contract, err := bindStrikeGovernorStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StrikeGovernorStorageFilterer{contract: contract}, nil
}

// bindStrikeGovernorStorage binds a generic wrapper to an already deployed contract.
func bindStrikeGovernorStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StrikeGovernorStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrikeGovernorStorage *StrikeGovernorStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrikeGovernorStorage.Contract.StrikeGovernorStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrikeGovernorStorage *StrikeGovernorStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrikeGovernorStorage.Contract.StrikeGovernorStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrikeGovernorStorage *StrikeGovernorStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrikeGovernorStorage.Contract.StrikeGovernorStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StrikeGovernorStorage *StrikeGovernorStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StrikeGovernorStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StrikeGovernorStorage *StrikeGovernorStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StrikeGovernorStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StrikeGovernorStorage *StrikeGovernorStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StrikeGovernorStorage.Contract.contract.Transact(opts, method, params...)
}

// TimersMetaData contains all meta data concerning the Timers contract.
var TimersMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220ab68f7d27046130293706ab8c72c69225f2a1825c5d35b35b1bce7289372e03564736f6c63430008130033",
}

// TimersABI is the input ABI used to generate the binding from.
// Deprecated: Use TimersMetaData.ABI instead.
var TimersABI = TimersMetaData.ABI

// TimersBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TimersMetaData.Bin instead.
var TimersBin = TimersMetaData.Bin

// DeployTimers deploys a new Ethereum contract, binding an instance of Timers to it.
func DeployTimers(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Timers, error) {
	parsed, err := TimersMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TimersBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Timers{TimersCaller: TimersCaller{contract: contract}, TimersTransactor: TimersTransactor{contract: contract}, TimersFilterer: TimersFilterer{contract: contract}}, nil
}

// Timers is an auto generated Go binding around an Ethereum contract.
type Timers struct {
	TimersCaller     // Read-only binding to the contract
	TimersTransactor // Write-only binding to the contract
	TimersFilterer   // Log filterer for contract events
}

// TimersCaller is an auto generated read-only Go binding around an Ethereum contract.
type TimersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TimersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TimersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TimersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TimersSession struct {
	Contract     *Timers           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TimersCallerSession struct {
	Contract *TimersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TimersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TimersTransactorSession struct {
	Contract     *TimersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TimersRaw is an auto generated low-level Go binding around an Ethereum contract.
type TimersRaw struct {
	Contract *Timers // Generic contract binding to access the raw methods on
}

// TimersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TimersCallerRaw struct {
	Contract *TimersCaller // Generic read-only contract binding to access the raw methods on
}

// TimersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TimersTransactorRaw struct {
	Contract *TimersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTimers creates a new instance of Timers, bound to a specific deployed contract.
func NewTimers(address common.Address, backend bind.ContractBackend) (*Timers, error) {
	contract, err := bindTimers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Timers{TimersCaller: TimersCaller{contract: contract}, TimersTransactor: TimersTransactor{contract: contract}, TimersFilterer: TimersFilterer{contract: contract}}, nil
}

// NewTimersCaller creates a new read-only instance of Timers, bound to a specific deployed contract.
func NewTimersCaller(address common.Address, caller bind.ContractCaller) (*TimersCaller, error) {
	contract, err := bindTimers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TimersCaller{contract: contract}, nil
}

// NewTimersTransactor creates a new write-only instance of Timers, bound to a specific deployed contract.
func NewTimersTransactor(address common.Address, transactor bind.ContractTransactor) (*TimersTransactor, error) {
	contract, err := bindTimers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TimersTransactor{contract: contract}, nil
}

// NewTimersFilterer creates a new log filterer instance of Timers, bound to a specific deployed contract.
func NewTimersFilterer(address common.Address, filterer bind.ContractFilterer) (*TimersFilterer, error) {
	contract, err := bindTimers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TimersFilterer{contract: contract}, nil
}

// bindTimers binds a generic wrapper to an already deployed contract.
func bindTimers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TimersMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timers *TimersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Timers.Contract.TimersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timers *TimersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timers.Contract.TimersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timers *TimersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timers.Contract.TimersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Timers *TimersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Timers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Timers *TimersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Timers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Timers *TimersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Timers.Contract.contract.Transact(opts, method, params...)
}

// UUPSUpgradeableMetaData contains all meta data concerning the UUPSUpgradeable contract.
var UUPSUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// UUPSUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use UUPSUpgradeableMetaData.ABI instead.
var UUPSUpgradeableABI = UUPSUpgradeableMetaData.ABI

// UUPSUpgradeable is an auto generated Go binding around an Ethereum contract.
type UUPSUpgradeable struct {
	UUPSUpgradeableCaller     // Read-only binding to the contract
	UUPSUpgradeableTransactor // Write-only binding to the contract
	UUPSUpgradeableFilterer   // Log filterer for contract events
}

// UUPSUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type UUPSUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UUPSUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UUPSUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UUPSUpgradeableSession struct {
	Contract     *UUPSUpgradeable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UUPSUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UUPSUpgradeableCallerSession struct {
	Contract *UUPSUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UUPSUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UUPSUpgradeableTransactorSession struct {
	Contract     *UUPSUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UUPSUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type UUPSUpgradeableRaw struct {
	Contract *UUPSUpgradeable // Generic contract binding to access the raw methods on
}

// UUPSUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UUPSUpgradeableCallerRaw struct {
	Contract *UUPSUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// UUPSUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UUPSUpgradeableTransactorRaw struct {
	Contract *UUPSUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUUPSUpgradeable creates a new instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeable(address common.Address, backend bind.ContractBackend) (*UUPSUpgradeable, error) {
	contract, err := bindUUPSUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeable{UUPSUpgradeableCaller: UUPSUpgradeableCaller{contract: contract}, UUPSUpgradeableTransactor: UUPSUpgradeableTransactor{contract: contract}, UUPSUpgradeableFilterer: UUPSUpgradeableFilterer{contract: contract}}, nil
}

// NewUUPSUpgradeableCaller creates a new read-only instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*UUPSUpgradeableCaller, error) {
	contract, err := bindUUPSUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableCaller{contract: contract}, nil
}

// NewUUPSUpgradeableTransactor creates a new write-only instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*UUPSUpgradeableTransactor, error) {
	contract, err := bindUUPSUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableTransactor{contract: contract}, nil
}

// NewUUPSUpgradeableFilterer creates a new log filterer instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*UUPSUpgradeableFilterer, error) {
	contract, err := bindUUPSUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableFilterer{contract: contract}, nil
}

// bindUUPSUpgradeable binds a generic wrapper to an already deployed contract.
func bindUUPSUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UUPSUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSUpgradeable *UUPSUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UUPSUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSUpgradeable *UUPSUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSUpgradeable *UUPSUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UUPSUpgradeable.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableSession) ProxiableUUID() ([32]byte, error) {
	return _UUPSUpgradeable.Contract.ProxiableUUID(&_UUPSUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_UUPSUpgradeable *UUPSUpgradeableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _UUPSUpgradeable.Contract.ProxiableUUID(&_UUPSUpgradeable.CallOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeTo(&_UUPSUpgradeable.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeTo(&_UUPSUpgradeable.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeToAndCall(&_UUPSUpgradeable.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeToAndCall(&_UUPSUpgradeable.TransactOpts, newImplementation, data)
}

// UUPSUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableAdminChangedIterator struct {
	Event *UUPSUpgradeableAdminChanged // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableAdminChanged)
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
		it.Event = new(UUPSUpgradeableAdminChanged)
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
func (it *UUPSUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableAdminChanged represents a AdminChanged event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UUPSUpgradeableAdminChangedIterator, error) {

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableAdminChangedIterator{contract: _UUPSUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableAdminChanged) (event.Subscription, error) {

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableAdminChanged)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseAdminChanged(log types.Log) (*UUPSUpgradeableAdminChanged, error) {
	event := new(UUPSUpgradeableAdminChanged)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableBeaconUpgradedIterator struct {
	Event *UUPSUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableBeaconUpgraded)
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
		it.Event = new(UUPSUpgradeableBeaconUpgraded)
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
func (it *UUPSUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UUPSUpgradeableBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableBeaconUpgradedIterator{contract: _UUPSUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableBeaconUpgraded)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*UUPSUpgradeableBeaconUpgraded, error) {
	event := new(UUPSUpgradeableBeaconUpgraded)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableInitializedIterator struct {
	Event *UUPSUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableInitialized)
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
		it.Event = new(UUPSUpgradeableInitialized)
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
func (it *UUPSUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableInitialized represents a Initialized event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*UUPSUpgradeableInitializedIterator, error) {

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableInitializedIterator{contract: _UUPSUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableInitialized)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseInitialized(log types.Log) (*UUPSUpgradeableInitialized, error) {
	event := new(UUPSUpgradeableInitialized)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableUpgradedIterator struct {
	Event *UUPSUpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableUpgraded)
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
		it.Event = new(UUPSUpgradeableUpgraded)
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
func (it *UUPSUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableUpgraded represents a Upgraded event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UUPSUpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableUpgradedIterator{contract: _UUPSUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableUpgraded)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseUpgraded(log types.Log) (*UUPSUpgradeableUpgraded, error) {
	event := new(UUPSUpgradeableUpgraded)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifySignatureMetaData contains all meta data concerning the VerifySignature contract.
var VerifySignatureMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566050600b82828239805160001a6073146043577f4e487b7100000000000000000000000000000000000000000000000000000000600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205377c39dcf2afa8bcb35f292370339d2df619763fc56aadbf61921246334858a64736f6c63430008130033",
}

// VerifySignatureABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifySignatureMetaData.ABI instead.
var VerifySignatureABI = VerifySignatureMetaData.ABI

// VerifySignatureBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifySignatureMetaData.Bin instead.
var VerifySignatureBin = VerifySignatureMetaData.Bin

// DeployVerifySignature deploys a new Ethereum contract, binding an instance of VerifySignature to it.
func DeployVerifySignature(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VerifySignature, error) {
	parsed, err := VerifySignatureMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifySignatureBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VerifySignature{VerifySignatureCaller: VerifySignatureCaller{contract: contract}, VerifySignatureTransactor: VerifySignatureTransactor{contract: contract}, VerifySignatureFilterer: VerifySignatureFilterer{contract: contract}}, nil
}

// VerifySignature is an auto generated Go binding around an Ethereum contract.
type VerifySignature struct {
	VerifySignatureCaller     // Read-only binding to the contract
	VerifySignatureTransactor // Write-only binding to the contract
	VerifySignatureFilterer   // Log filterer for contract events
}

// VerifySignatureCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifySignatureCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifySignatureTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifySignatureFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifySignatureSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifySignatureSession struct {
	Contract     *VerifySignature  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifySignatureCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifySignatureCallerSession struct {
	Contract *VerifySignatureCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// VerifySignatureTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifySignatureTransactorSession struct {
	Contract     *VerifySignatureTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// VerifySignatureRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifySignatureRaw struct {
	Contract *VerifySignature // Generic contract binding to access the raw methods on
}

// VerifySignatureCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifySignatureCallerRaw struct {
	Contract *VerifySignatureCaller // Generic read-only contract binding to access the raw methods on
}

// VerifySignatureTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifySignatureTransactorRaw struct {
	Contract *VerifySignatureTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifySignature creates a new instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignature(address common.Address, backend bind.ContractBackend) (*VerifySignature, error) {
	contract, err := bindVerifySignature(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VerifySignature{VerifySignatureCaller: VerifySignatureCaller{contract: contract}, VerifySignatureTransactor: VerifySignatureTransactor{contract: contract}, VerifySignatureFilterer: VerifySignatureFilterer{contract: contract}}, nil
}

// NewVerifySignatureCaller creates a new read-only instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureCaller(address common.Address, caller bind.ContractCaller) (*VerifySignatureCaller, error) {
	contract, err := bindVerifySignature(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureCaller{contract: contract}, nil
}

// NewVerifySignatureTransactor creates a new write-only instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifySignatureTransactor, error) {
	contract, err := bindVerifySignature(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureTransactor{contract: contract}, nil
}

// NewVerifySignatureFilterer creates a new log filterer instance of VerifySignature, bound to a specific deployed contract.
func NewVerifySignatureFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifySignatureFilterer, error) {
	contract, err := bindVerifySignature(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifySignatureFilterer{contract: contract}, nil
}

// bindVerifySignature binds a generic wrapper to an already deployed contract.
func bindVerifySignature(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifySignatureMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifySignature *VerifySignatureRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifySignature.Contract.VerifySignatureCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifySignature *VerifySignatureRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifySignature.Contract.VerifySignatureTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifySignature *VerifySignatureRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifySignature.Contract.VerifySignatureTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VerifySignature *VerifySignatureCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VerifySignature.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VerifySignature *VerifySignatureTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VerifySignature.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VerifySignature *VerifySignatureTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VerifySignature.Contract.contract.Transact(opts, method, params...)
}
