// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abc

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// AbcABI is the input ABI used to generate the binding from.
const AbcABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// AbcBin is the compiled bytecode used for deploying new contracts.
var AbcBin = "0x60806040523480156200001157600080fd5b506000620000246200023560201b60201c565b9050806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3506040518060400160405280600881526020017f414243546f6b656e00000000000000000000000000000000000000000000000081525060069081620001089190620004b7565b506040518060400160405280600381526020017f4142430000000000000000000000000000000000000000000000000000000000815250600590816200014f9190620004b7565b506012600460006101000a81548160ff021916908360ff1602179055506b033b2e3c9fd0803ce8000000600381905550600354600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef600354604051620002279190620005af565b60405180910390a3620005cc565b600033905090565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002bf57607f821691505b602082108103620002d557620002d462000277565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200033f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000300565b6200034b868362000300565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600062000398620003926200038c8462000363565b6200036d565b62000363565b9050919050565b6000819050919050565b620003b48362000377565b620003cc620003c3826200039f565b8484546200030d565b825550505050565b600090565b620003e3620003d4565b620003f0818484620003a9565b505050565b5b8181101562000418576200040c600082620003d9565b600181019050620003f6565b5050565b601f82111562000467576200043181620002db565b6200043c84620002f0565b810160208510156200044c578190505b620004646200045b85620002f0565b830182620003f5565b50505b505050565b600082821c905092915050565b60006200048c600019846008026200046c565b1980831691505092915050565b6000620004a7838362000479565b9150826002028217905092915050565b620004c2826200023d565b67ffffffffffffffff811115620004de57620004dd62000248565b5b620004ea8254620002a6565b620004f78282856200041c565b600060209050601f8311600181146200052f57600084156200051a578287015190505b62000526858262000499565b86555062000596565b601f1984166200053f86620002db565b60005b82811015620005695784890151825560018201915060208501945060208101905062000542565b8683101562000589578489015162000585601f89168262000479565b8355505b6001600288020188555050505b505050505050565b620005a98162000363565b82525050565b6000602082019050620005c660008301846200059e565b92915050565b611bdc80620005dc6000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c8063893d20e811610097578063a457c2d711610066578063a457c2d7146102b3578063a9059cbb146102e3578063dd62ed3e14610313578063f2fde38b1461034357610100565b8063893d20e8146102295780638da5cb5b1461024757806395d89b4114610265578063a0712d681461028357610100565b8063313ce567116100d3578063313ce567146101a157806339509351146101bf57806370a08231146101ef578063715018a61461021f57610100565b806306fdde0314610105578063095ea7b31461012357806318160ddd1461015357806323b872dd14610171575b600080fd5b61010d61035f565b60405161011a9190611353565b60405180910390f35b61013d6004803603810190610138919061140e565b6103f1565b60405161014a9190611469565b60405180910390f35b61015b61040f565b6040516101689190611493565b60405180910390f35b61018b600480360381019061018691906114ae565b610419565b6040516101989190611469565b60405180910390f35b6101a96104f2565b6040516101b6919061151d565b60405180910390f35b6101d960048036038101906101d4919061140e565b610509565b6040516101e69190611469565b60405180910390f35b61020960048036038101906102049190611538565b6105bc565b6040516102169190611493565b60405180910390f35b610227610605565b005b610231610758565b60405161023e9190611574565b60405180910390f35b61024f610767565b60405161025c9190611574565b60405180910390f35b61026d610790565b60405161027a9190611353565b60405180910390f35b61029d6004803603810190610298919061158f565b610822565b6040516102aa9190611469565b60405180910390f35b6102cd60048036038101906102c8919061140e565b6108d3565b6040516102da9190611469565b60405180910390f35b6102fd60048036038101906102f8919061140e565b6109a0565b60405161030a9190611469565b60405180910390f35b61032d600480360381019061032891906115bc565b6109be565b60405161033a9190611493565b60405180910390f35b61035d60048036038101906103589190611538565b610a45565b005b60606006805461036e9061162b565b80601f016020809104026020016040519081016040528092919081815260200182805461039a9061162b565b80156103e75780601f106103bc576101008083540402835291602001916103e7565b820191906000526020600020905b8154815290600101906020018083116103ca57829003601f168201915b5050505050905090565b60006104056103fe610ae6565b8484610aee565b6001905092915050565b6000600354905090565b6000610426848484610cb7565b6104e784610432610ae6565b6104e285604051806060016040528060288152602001611b3460289139600260008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000610498610ae6565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f439092919063ffffffff16565b610aee565b600190509392505050565b6000600460009054906101000a900460ff16905090565b60006105b2610516610ae6565b846105ad8560026000610527610ae6565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610fa790919063ffffffff16565b610aee565b6001905092915050565b6000600160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b61060d610ae6565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461069a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610691906116a8565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a360008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000610762610767565b905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606005805461079f9061162b565b80601f01602080910402602001604051908101604052809291908181526020018280546107cb9061162b565b80156108185780601f106107ed57610100808354040283529160200191610818565b820191906000526020600020905b8154815290600101906020018083116107fb57829003601f168201915b5050505050905090565b600061082c610ae6565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146108b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b0906116a8565b60405180910390fd5b6108ca6108c4610ae6565b83611005565b60019050919050565b60006109966108e0610ae6565b8461099185604051806060016040528060258152602001611b82602591396002600061090a610ae6565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f439092919063ffffffff16565b610aee565b6001905092915050565b60006109b46109ad610ae6565b8484610cb7565b6001905092915050565b6000600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b610a4d610ae6565b73ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610ada576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ad1906116a8565b60405180910390fd5b610ae38161118e565b50565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610b5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b549061173a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610bcc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc3906117cc565b60405180910390fd5b80600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051610caa9190611493565b60405180910390a3505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610d26576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1d9061185e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d8c906118f0565b60405180910390fd5b610e0181604051806060016040528060268152602001611b5c60269139600160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610f439092919063ffffffff16565b600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550610e9681600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610fa790919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610f369190611493565b60405180910390a3505050565b6000838311158290610f8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f829190611353565b60405180910390fd5b5060008385610f9a919061193f565b9050809150509392505050565b6000808284610fb69190611973565b905083811015610ffb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ff290611a15565b60405180910390fd5b8091505092915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611074576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161106b90611a81565b60405180910390fd5b61108981600354610fa790919063ffffffff16565b6003819055506110e181600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054610fa790919063ffffffff16565b600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516111829190611493565b60405180910390a35050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036111fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111f490611b13565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1660008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600081519050919050565b600082825260208201905092915050565b60005b838110156112f45780820151818401526020810190506112d9565b83811115611303576000848401525b50505050565b6000601f19601f8301169050919050565b6000611325826112ba565b61132f81856112c5565b935061133f8185602086016112d6565b61134881611309565b840191505092915050565b6000602082019050818103600083015261136d818461131a565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006113a58261137a565b9050919050565b6113b58161139a565b81146113c057600080fd5b50565b6000813590506113d2816113ac565b92915050565b6000819050919050565b6113eb816113d8565b81146113f657600080fd5b50565b600081359050611408816113e2565b92915050565b6000806040838503121561142557611424611375565b5b6000611433858286016113c3565b9250506020611444858286016113f9565b9150509250929050565b60008115159050919050565b6114638161144e565b82525050565b600060208201905061147e600083018461145a565b92915050565b61148d816113d8565b82525050565b60006020820190506114a86000830184611484565b92915050565b6000806000606084860312156114c7576114c6611375565b5b60006114d5868287016113c3565b93505060206114e6868287016113c3565b92505060406114f7868287016113f9565b9150509250925092565b600060ff82169050919050565b61151781611501565b82525050565b6000602082019050611532600083018461150e565b92915050565b60006020828403121561154e5761154d611375565b5b600061155c848285016113c3565b91505092915050565b61156e8161139a565b82525050565b60006020820190506115896000830184611565565b92915050565b6000602082840312156115a5576115a4611375565b5b60006115b3848285016113f9565b91505092915050565b600080604083850312156115d3576115d2611375565b5b60006115e1858286016113c3565b92505060206115f2858286016113c3565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061164357607f821691505b602082108103611656576116556115fc565b5b50919050565b7f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572600082015250565b60006116926020836112c5565b915061169d8261165c565b602082019050919050565b600060208201905081810360008301526116c181611685565b9050919050565b7f42455032303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b60006117246024836112c5565b915061172f826116c8565b604082019050919050565b6000602082019050818103600083015261175381611717565b9050919050565b7f42455032303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b60006117b66022836112c5565b91506117c18261175a565b604082019050919050565b600060208201905081810360008301526117e5816117a9565b9050919050565b7f42455032303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006118486025836112c5565b9150611853826117ec565b604082019050919050565b600060208201905081810360008301526118778161183b565b9050919050565b7f42455032303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b60006118da6023836112c5565b91506118e58261187e565b604082019050919050565b60006020820190508181036000830152611909816118cd565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061194a826113d8565b9150611955836113d8565b92508282101561196857611967611910565b5b828203905092915050565b600061197e826113d8565b9150611989836113d8565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156119be576119bd611910565b5b828201905092915050565b7f536166654d6174683a206164646974696f6e206f766572666c6f770000000000600082015250565b60006119ff601b836112c5565b9150611a0a826119c9565b602082019050919050565b60006020820190508181036000830152611a2e816119f2565b9050919050565b7f42455032303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b6000611a6b601f836112c5565b9150611a7682611a35565b602082019050919050565b60006020820190508181036000830152611a9a81611a5e565b9050919050565b7f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008201527f6464726573730000000000000000000000000000000000000000000000000000602082015250565b6000611afd6026836112c5565b9150611b0882611aa1565b604082019050919050565b60006020820190508181036000830152611b2c81611af0565b905091905056fe42455032303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636542455032303a207472616e7366657220616d6f756e7420657863656564732062616c616e636542455032303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212202337bbf58420b3adf80b46b88897df1540b293c2a6f632faacd938e1cbbc513864736f6c634300080f0033"

// DeployAbc deploys a new Ethereum contract, binding an instance of Abc to it.
func DeployAbc(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Abc, error) {
	parsed, err := abi.JSON(strings.NewReader(AbcABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AbcBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Abc{AbcCaller: AbcCaller{contract: contract}, AbcTransactor: AbcTransactor{contract: contract}, AbcFilterer: AbcFilterer{contract: contract}}, nil
}

// Abc is an auto generated Go binding around an Ethereum contract.
type Abc struct {
	AbcCaller     // Read-only binding to the contract
	AbcTransactor // Write-only binding to the contract
	AbcFilterer   // Log filterer for contract events
}

// AbcCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbcSession struct {
	Contract     *Abc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbcCallerSession struct {
	Contract *AbcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AbcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbcTransactorSession struct {
	Contract     *AbcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AbcRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbcRaw struct {
	Contract *Abc // Generic contract binding to access the raw methods on
}

// AbcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbcCallerRaw struct {
	Contract *AbcCaller // Generic read-only contract binding to access the raw methods on
}

// AbcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbcTransactorRaw struct {
	Contract *AbcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbc creates a new instance of Abc, bound to a specific deployed contract.
func NewAbc(address common.Address, backend bind.ContractBackend) (*Abc, error) {
	contract, err := bindAbc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Abc{AbcCaller: AbcCaller{contract: contract}, AbcTransactor: AbcTransactor{contract: contract}, AbcFilterer: AbcFilterer{contract: contract}}, nil
}

// NewAbcCaller creates a new read-only instance of Abc, bound to a specific deployed contract.
func NewAbcCaller(address common.Address, caller bind.ContractCaller) (*AbcCaller, error) {
	contract, err := bindAbc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbcCaller{contract: contract}, nil
}

// NewAbcTransactor creates a new write-only instance of Abc, bound to a specific deployed contract.
func NewAbcTransactor(address common.Address, transactor bind.ContractTransactor) (*AbcTransactor, error) {
	contract, err := bindAbc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbcTransactor{contract: contract}, nil
}

// NewAbcFilterer creates a new log filterer instance of Abc, bound to a specific deployed contract.
func NewAbcFilterer(address common.Address, filterer bind.ContractFilterer) (*AbcFilterer, error) {
	contract, err := bindAbc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbcFilterer{contract: contract}, nil
}

// bindAbc binds a generic wrapper to an already deployed contract.
func bindAbc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AbcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abc *AbcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abc.Contract.AbcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abc *AbcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abc.Contract.AbcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abc *AbcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abc.Contract.AbcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Abc *AbcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Abc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Abc *AbcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Abc *AbcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Abc.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abc *AbcCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abc *AbcSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Abc.Contract.Allowance(&_Abc.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Abc *AbcCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Abc.Contract.Allowance(&_Abc.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Abc *AbcCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Abc *AbcSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Abc.Contract.BalanceOf(&_Abc.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Abc *AbcCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Abc.Contract.BalanceOf(&_Abc.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abc *AbcCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abc *AbcSession) Decimals() (uint8, error) {
	return _Abc.Contract.Decimals(&_Abc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Abc *AbcCallerSession) Decimals() (uint8, error) {
	return _Abc.Contract.Decimals(&_Abc.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Abc *AbcCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Abc *AbcSession) GetOwner() (common.Address, error) {
	return _Abc.Contract.GetOwner(&_Abc.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Abc *AbcCallerSession) GetOwner() (common.Address, error) {
	return _Abc.Contract.GetOwner(&_Abc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abc *AbcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abc *AbcSession) Name() (string, error) {
	return _Abc.Contract.Name(&_Abc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Abc *AbcCallerSession) Name() (string, error) {
	return _Abc.Contract.Name(&_Abc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abc *AbcCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abc *AbcSession) Owner() (common.Address, error) {
	return _Abc.Contract.Owner(&_Abc.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Abc *AbcCallerSession) Owner() (common.Address, error) {
	return _Abc.Contract.Owner(&_Abc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abc *AbcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abc *AbcSession) Symbol() (string, error) {
	return _Abc.Contract.Symbol(&_Abc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Abc *AbcCallerSession) Symbol() (string, error) {
	return _Abc.Contract.Symbol(&_Abc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abc *AbcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Abc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abc *AbcSession) TotalSupply() (*big.Int, error) {
	return _Abc.Contract.TotalSupply(&_Abc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Abc *AbcCallerSession) TotalSupply() (*big.Int, error) {
	return _Abc.Contract.TotalSupply(&_Abc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abc *AbcTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abc *AbcSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.Approve(&_Abc.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Abc *AbcTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.Approve(&_Abc.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Abc *AbcTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Abc *AbcSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.DecreaseAllowance(&_Abc.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Abc *AbcTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.DecreaseAllowance(&_Abc.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Abc *AbcTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Abc *AbcSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.IncreaseAllowance(&_Abc.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Abc *AbcTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.IncreaseAllowance(&_Abc.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Abc *AbcTransactor) Mint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Abc *AbcSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.Mint(&_Abc.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Abc *AbcTransactorSession) Mint(amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.Mint(&_Abc.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abc *AbcTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abc *AbcSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abc.Contract.RenounceOwnership(&_Abc.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Abc *AbcTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Abc.Contract.RenounceOwnership(&_Abc.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Abc *AbcTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Abc *AbcSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.Transfer(&_Abc.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Abc *AbcTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.Transfer(&_Abc.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Abc *AbcTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Abc *AbcSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.TransferFrom(&_Abc.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Abc *AbcTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Abc.Contract.TransferFrom(&_Abc.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abc *AbcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Abc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abc *AbcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abc.Contract.TransferOwnership(&_Abc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Abc *AbcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Abc.Contract.TransferOwnership(&_Abc.TransactOpts, newOwner)
}

// AbcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Abc contract.
type AbcApprovalIterator struct {
	Event *AbcApproval // Event containing the contract specifics and raw log

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
func (it *AbcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbcApproval)
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
		it.Event = new(AbcApproval)
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
func (it *AbcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbcApproval represents a Approval event raised by the Abc contract.
type AbcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Abc *AbcFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AbcApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abc.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AbcApprovalIterator{contract: _Abc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Abc *AbcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AbcApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Abc.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbcApproval)
				if err := _Abc.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Abc *AbcFilterer) ParseApproval(log types.Log) (*AbcApproval, error) {
	event := new(AbcApproval)
	if err := _Abc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbcOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Abc contract.
type AbcOwnershipTransferredIterator struct {
	Event *AbcOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbcOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbcOwnershipTransferred)
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
		it.Event = new(AbcOwnershipTransferred)
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
func (it *AbcOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbcOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbcOwnershipTransferred represents a OwnershipTransferred event raised by the Abc contract.
type AbcOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abc *AbcFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbcOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abc.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbcOwnershipTransferredIterator{contract: _Abc.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Abc *AbcFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbcOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Abc.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbcOwnershipTransferred)
				if err := _Abc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Abc *AbcFilterer) ParseOwnershipTransferred(log types.Log) (*AbcOwnershipTransferred, error) {
	event := new(AbcOwnershipTransferred)
	if err := _Abc.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Abc contract.
type AbcTransferIterator struct {
	Event *AbcTransfer // Event containing the contract specifics and raw log

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
func (it *AbcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbcTransfer)
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
		it.Event = new(AbcTransfer)
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
func (it *AbcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbcTransfer represents a Transfer event raised by the Abc contract.
type AbcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Abc *AbcFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AbcTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abc.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AbcTransferIterator{contract: _Abc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Abc *AbcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AbcTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Abc.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbcTransfer)
				if err := _Abc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Abc *AbcFilterer) ParseTransfer(log types.Log) (*AbcTransfer, error) {
	event := new(AbcTransfer)
	if err := _Abc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
