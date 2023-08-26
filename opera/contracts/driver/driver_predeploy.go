package driver

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is NodeDriver contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
// Built from opera-sfc be4e79d5a5a425f08efd6d65b588a72ae90f706f, solc 0.5.17+commit.d19bba13.Emscripten.clang, optimize-runs 10000
func GetContractBin() []byte {
	return hexutil.MustDecode("0x608060405234801561001057600080fd5b50600436106101365760003560e01c806379bead38116100b2578063d6a0c7af11610081578063e08d7e6611610066578063e08d7e6614610676578063e30443bc146106e6578063ebdf104c1461071f57610136565b8063d6a0c7af14610608578063da7fc24f1461064357610136565b806379bead38146103d65780637f52e13e1461040f578063a4066fbe14610575578063b9cc6b1c1461059857610136565b8063242a6e3f1161010957806339e503ab116100ee57806339e503ab146102b1578063485cc955146102f05780634feb92f31461032b57610136565b8063242a6e3f1461021d578063267ab4461461029457610136565b806307690b2a1461013b5780630aeeca001461017857806318f628d4146101955780631e702f83146101fa575b600080fd5b6101766004803603604081101561015157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610885565b005b6101766004803603602081101561018e57600080fd5b5035610989565b61017660048036036101208110156101ac57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e0810135906101000135610a2b565b6101766004803603604081101561021057600080fd5b5080359060200135610b4f565b6101766004803603604081101561023357600080fd5b8135919081019060408101602082013564010000000081111561025557600080fd5b82018360208201111561026757600080fd5b8035906020019184600183028401116401000000008311171561028957600080fd5b509092509050610c1c565b610176600480360360208110156102aa57600080fd5b5035610cee565b610176600480360360608110156102c757600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060400135610d90565b6101766004803603604081101561030657600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516610e9b565b610176600480360361010081101561034257600080fd5b73ffffffffffffffffffffffffffffffffffffffff8235169160208101359181019060608101604082013564010000000081111561037f57600080fd5b82018360208201111561039157600080fd5b803590602001918460018302840111640100000000831117156103b357600080fd5b919350915080359060208101359060408101359060608101359060800135611043565b610176600480360360408110156103ec57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020013561119b565b610176600480360360a081101561042557600080fd5b81019060208101813564010000000081111561044057600080fd5b82018360208201111561045257600080fd5b8035906020019184602083028401116401000000008311171561047457600080fd5b91939092909160208101903564010000000081111561049257600080fd5b8201836020820111156104a457600080fd5b803590602001918460208302840111640100000000831117156104c657600080fd5b9193909290916020810190356401000000008111156104e457600080fd5b8201836020820111156104f657600080fd5b8035906020019184602083028401116401000000008311171561051857600080fd5b91939092909160208101903564010000000081111561053657600080fd5b82018360208201111561054857600080fd5b8035906020019184602083028401116401000000008311171561056a57600080fd5b919350915035611282565b6101766004803603604081101561058b57600080fd5b5080359060200135611411565b610176600480360360208110156105ae57600080fd5b8101906020810181356401000000008111156105c957600080fd5b8201836020820111156105db57600080fd5b803590602001918460018302840111640100000000831117156105fd57600080fd5b5090925090506114b7565b6101766004803603604081101561061e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81358116916020013516611587565b6101766004803603602081101561065957600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661166f565b6101766004803603602081101561068c57600080fd5b8101906020810181356401000000008111156106a757600080fd5b8201836020820111156106b957600080fd5b803590602001918460208302840111640100000000831117156106db57600080fd5b509092509050611763565b610176600480360360408110156106fc57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135611859565b6101766004803603608081101561073557600080fd5b81019060208101813564010000000081111561075057600080fd5b82018360208201111561076257600080fd5b8035906020019184602083028401116401000000008311171561078457600080fd5b9193909290916020810190356401000000008111156107a257600080fd5b8201836020820111156107b457600080fd5b803590602001918460208302840111640100000000831117156107d657600080fd5b9193909290916020810190356401000000008111156107f457600080fd5b82018360208201111561080657600080fd5b8035906020019184602083028401116401000000008311171561082857600080fd5b91939092909160208101903564010000000081111561084657600080fd5b82018360208201111561085857600080fd5b8035906020019184602083028401116401000000008311171561087a57600080fd5b509092509050611940565b60345473ffffffffffffffffffffffffffffffffffffffff1633146108f1576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517f07690b2a00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85811660048301528481166024830152915191909216916307690b2a91604480830192600092919082900301818387803b15801561096d57600080fd5b505af1158015610981573d6000803e3d6000fd5b505050505050565b60345473ffffffffffffffffffffffffffffffffffffffff1633146109f5576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b6040805182815290517f0151256d62457b809bbc891b1f81c6dd0b9987552c70ce915b519750cd434dd19181900360200190a150565b3315610a7e576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603454604080517f18f628d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c81166004830152602482018c9052604482018b9052606482018a90526084820189905260a4820188905260c4820187905260e482018690526101048201859052915191909216916318f628d49161012480830192600092919082900301818387803b158015610b2c57600080fd5b505af1158015610b40573d6000803e3d6000fd5b50505050505050505050505050565b3315610ba2576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603454604080517f1e702f830000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff90921691631e702f839160448082019260009290919082900301818387803b15801561096d57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff163314610c88576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b827f0f0ef1ab97439def0a9d2c6d9dc166207f1b13b99e62b442b2993d6153c63a6e838360405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a2505050565b60345473ffffffffffffffffffffffffffffffffffffffff163314610d5a576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b6040805182815290517f2ccdfd47cf0c1f1069d949f1789bb79b2f12821f021634fc835af1de66ea2feb9181900360200190a150565b60345473ffffffffffffffffffffffffffffffffffffffff163314610dfc576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517f39e503ab00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff86811660048301526024820186905260448201859052915191909216916339e503ab91606480830192600092919082900301818387803b158015610e7e57600080fd5b505af1158015610e92573d6000803e3d6000fd5b50505050505050565b600054610100900460ff1680610eb45750610eb4611af5565b80610ec2575060005460ff16155b610efd5760405162461bcd60e51b815260040180806020018281038252602e815260200180611afc602e913960400191505060405180910390fd5b600054610100900460ff16158015610f6357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff85169081179091556040517f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069390600090a2603580547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416179055801561103e57600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b505050565b3315611096576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b158015610b2c57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff163314611207576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517f79bead3800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015260248201859052915191909216916379bead3891604480830192600092919082900301818387803b15801561096d57600080fd5b33156112d5576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663592fe0c08a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808060200180602001806020018060200186815260200185810385528e8e82818152602001925060200280828437600083820152601f01601f191690910186810385528c8152602090810191508d908d0280828437600083820152601f01601f191690910186810384528a8152602090810191508b908b0280828437600083820152601f01601f19169091018681038352888152602090810191508990890280828437600081840152601f19601f8201169050808301925050509d5050505050505050505050505050600060405180830381600087803b158015610b2c57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff16331461147d576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b60408051828152905183917fb975807576e3b1461be7de07ebf7d20e4790ed802d7a0c4fdd0a1a13df72a935919081900360200190a25050565b60345473ffffffffffffffffffffffffffffffffffffffff163314611523576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b7f47d10eed096a44e3d0abc586c7e3a5d6cb5358cc90e7d437cd0627f7e765fb99828260405180806020018281038252848482818152602001925080828437600083820152604051601f909101601f19169092018290039550909350505050a15050565b60345473ffffffffffffffffffffffffffffffffffffffff1633146115f3576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517fd6a0c7af00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015284811660248301529151919092169163d6a0c7af91604480830192600092919082900301818387803b15801561096d57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff1633146116db576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b60405173ffffffffffffffffffffffffffffffffffffffff8216907f64ee8f7bfc37fc205d7194ee3d64947ab7b57e663cd0d1abd3ef24503583069390600090a2603480547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b33156117b6576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b6034546040517fe08d7e660000000000000000000000000000000000000000000000000000000081526020600482018181526024830185905273ffffffffffffffffffffffffffffffffffffffff9093169263e08d7e6692869286929182916044909101908590850280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b15801561096d57600080fd5b60345473ffffffffffffffffffffffffffffffffffffffff1633146118c5576040805162461bcd60e51b815260206004820152601960248201527f63616c6c6572206973206e6f7420746865206261636b656e6400000000000000604482015290519081900360640190fd5b603554604080517fe30443bc00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590529151919092169163e30443bc91604480830192600092919082900301818387803b15801561096d57600080fd5b3315611993576040805162461bcd60e51b815260206004820152600c60248201527f6e6f742063616c6c61626c650000000000000000000000000000000000000000604482015290519081900360640190fd5b603460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663592fe0c0898989898989898963322adc3a6040518a63ffffffff1660e01b8152600401808060200180602001806020018060200186815260200185810385528e8e82818152602001925060200280828437600083820152601f01601f191690910186810385528c8152602090810191508d908d0280828437600083820152601f01601f191690910186810384528a8152602090810191508b908b0280828437600083820152601f01601f19169091018681038352888152602090810191508990890280828437600081840152601f19601f8201169050808301925050509d5050505050505050505050505050600060405180830381600087803b158015611ad357600080fd5b505af1158015611ae7573d6000803e3d6000fd5b505050505050505050505050565b303b159056fe436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564a265627a7a72315820f5cb1667fae9152cab345389c6f9b54dccee170ca51c179d91361997f298acf464736f6c63430005110032")
}

// ContractAddress is the NodeDriver contract address
var ContractAddress = common.HexToAddress("0xd100a01e00000000000000000000000000000000")
