package contract_test

import (
	"io/ioutil"
)

var compiledContractsJson = map[string]string{
	"Counter1": `{
		"abi": "[{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
		"bin": "6080604052600160005534801561001557600080fd5b50610173806100256000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806306661abd1461003b578063d09de08a14610059575b600080fd5b610043610063565b604051610050919061009d565b60405180910390f35b610061610069565b005b60005481565b600160008082825461007b91906100e7565b92505081905550565b6000819050919050565b61009781610084565b82525050565b60006020820190506100b2600083018461008e565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006100f282610084565b91506100fd83610084565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610132576101316100b8565b5b82820190509291505056fea2646970667358221220fc4389ee12d6a47e2156c1633e481ef83c9a2f38e942a4b5d166499497afbd9564736f6c634300080f0033",
		"contractName": "Counter1"
	}`,
	"Counter2": `{
		"abi": "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initValue\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"initTxt\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
		"bin": "608060405260016000553480156200001657600080fd5b50604051620009183803806200091883398181016040528101906200003c919062000235565b816000819055508060019081620000549190620004dc565b505050620005c3565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b620000868162000071565b81146200009257600080fd5b50565b600081519050620000a6816200007b565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6200010182620000b6565b810181811067ffffffffffffffff82111715620001235762000122620000c7565b5b80604052505050565b6000620001386200005d565b9050620001468282620000f6565b919050565b600067ffffffffffffffff821115620001695762000168620000c7565b5b6200017482620000b6565b9050602081019050919050565b60005b83811015620001a157808201518184015260208101905062000184565b83811115620001b1576000848401525b50505050565b6000620001ce620001c8846200014b565b6200012c565b905082815260208101848484011115620001ed57620001ec620000b1565b5b620001fa84828562000181565b509392505050565b600082601f8301126200021a5762000219620000ac565b5b81516200022c848260208601620001b7565b91505092915050565b600080604083850312156200024f576200024e62000067565b5b60006200025f8582860162000095565b925050602083015167ffffffffffffffff8111156200028357620002826200006c565b5b620002918582860162000202565b9150509250929050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002ee57607f821691505b602082108103620003045762000303620002a6565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026200036e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826200032f565b6200037a86836200032f565b95508019841693508086168417925050509392505050565b6000819050919050565b6000620003bd620003b7620003b18462000071565b62000392565b62000071565b9050919050565b6000819050919050565b620003d9836200039c565b620003f1620003e882620003c4565b8484546200033c565b825550505050565b600090565b62000408620003f9565b62000415818484620003ce565b505050565b5b818110156200043d5762000431600082620003fe565b6001810190506200041b565b5050565b601f8211156200048c5762000456816200030a565b62000461846200031f565b8101602085101562000471578190505b6200048962000480856200031f565b8301826200041a565b50505b505050565b600082821c905092915050565b6000620004b16000198460080262000491565b1980831691505092915050565b6000620004cc83836200049e565b9150826002028217905092915050565b620004e7826200029b565b67ffffffffffffffff811115620005035762000502620000c7565b5b6200050f8254620002d5565b6200051c82828562000441565b600060209050601f8311600181146200055457600084156200053f578287015190505b6200054b8582620004be565b865550620005bb565b601f19841662000564866200030a565b60005b828110156200058e5784890151825560018201915060208501945060208101905062000567565b86831015620005ae5784890151620005aa601f8916826200049e565b8355505b6001600288020188555050505b505050505050565b61034580620005d36000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806306661abd14610046578063d09de08a14610064578063e21f37ce1461006e575b600080fd5b61004e61008c565b60405161005b9190610154565b60405180910390f35b61006c610092565b005b6100766100ad565b6040516100839190610208565b60405180910390f35b60005481565b60016000808282546100a49190610259565b92505081905550565b600180546100ba906102de565b80601f01602080910402602001604051908101604052809291908181526020018280546100e6906102de565b80156101335780601f1061010857610100808354040283529160200191610133565b820191906000526020600020905b81548152906001019060200180831161011657829003601f168201915b505050505081565b6000819050919050565b61014e8161013b565b82525050565b60006020820190506101696000830184610145565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156101a957808201518184015260208101905061018e565b838111156101b8576000848401525b50505050565b6000601f19601f8301169050919050565b60006101da8261016f565b6101e4818561017a565b93506101f481856020860161018b565b6101fd816101be565b840191505092915050565b6000602082019050818103600083015261022281846101cf565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60006102648261013b565b915061026f8361013b565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156102a4576102a361022a565b5b828201905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806102f657607f821691505b602082108103610309576103086102af565b5b5091905056fea2646970667358221220d58c0d6ae3efc82f0d21e38c812982ece0298a947fed3b5d7fd1ee3ddcdbea3464736f6c634300080f0033",
		"contractName": "Counter2"
	}`,
	"HelloWorld": `{
		"abi": "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"initMessage\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"message\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newMessage\",\"type\":\"string\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
		"bin": "60806040523480156200001157600080fd5b5060405162000c0138038062000c018339818101604052810190620000379190620001ed565b806000908162000048919062000489565b505062000570565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620000b9826200006e565b810181811067ffffffffffffffff82111715620000db57620000da6200007f565b5b80604052505050565b6000620000f062000050565b9050620000fe8282620000ae565b919050565b600067ffffffffffffffff8211156200012157620001206200007f565b5b6200012c826200006e565b9050602081019050919050565b60005b83811015620001595780820151818401526020810190506200013c565b8381111562000169576000848401525b50505050565b600062000186620001808462000103565b620000e4565b905082815260208101848484011115620001a557620001a462000069565b5b620001b284828562000139565b509392505050565b600082601f830112620001d257620001d162000064565b5b8151620001e48482602086016200016f565b91505092915050565b6000602082840312156200020657620002056200005a565b5b600082015167ffffffffffffffff8111156200022757620002266200005f565b5b6200023584828501620001ba565b91505092915050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806200029157607f821691505b602082108103620002a757620002a662000249565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003117fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620002d2565b6200031d8683620002d2565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006200036a620003646200035e8462000335565b6200033f565b62000335565b9050919050565b6000819050919050565b620003868362000349565b6200039e620003958262000371565b848454620002df565b825550505050565b600090565b620003b5620003a6565b620003c28184846200037b565b505050565b5b81811015620003ea57620003de600082620003ab565b600181019050620003c8565b5050565b601f82111562000439576200040381620002ad565b6200040e84620002c2565b810160208510156200041e578190505b620004366200042d85620002c2565b830182620003c7565b50505b505050565b600082821c905092915050565b60006200045e600019846008026200043e565b1980831691505092915050565b60006200047983836200044b565b9150826002028217905092915050565b62000494826200023e565b67ffffffffffffffff811115620004b057620004af6200007f565b5b620004bc825462000278565b620004c9828285620003ee565b600060209050601f831160018114620005015760008415620004ec578287015190505b620004f885826200046b565b86555062000568565b601f1984166200051186620002ad565b60005b828110156200053b5784890151825560018201915060208501945060208101905062000514565b868310156200055b578489015162000557601f8916826200044b565b8355505b6001600288020188555050505b505050505050565b61068180620005806000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80633d7403a31461003b578063e21f37ce14610057575b600080fd5b61005560048036038101906100509190610270565b610075565b005b61005f610088565b60405161006c9190610341565b60405180910390f35b80600090816100849190610579565b5050565b6000805461009590610392565b80601f01602080910402602001604051908101604052809291908181526020018280546100c190610392565b801561010e5780601f106100e35761010080835404028352916020019161010e565b820191906000526020600020905b8154815290600101906020018083116100f157829003601f168201915b505050505081565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61017d82610134565b810181811067ffffffffffffffff8211171561019c5761019b610145565b5b80604052505050565b60006101af610116565b90506101bb8282610174565b919050565b600067ffffffffffffffff8211156101db576101da610145565b5b6101e482610134565b9050602081019050919050565b82818337600083830152505050565b600061021361020e846101c0565b6101a5565b90508281526020810184848401111561022f5761022e61012f565b5b61023a8482856101f1565b509392505050565b600082601f8301126102575761025661012a565b5b8135610267848260208601610200565b91505092915050565b60006020828403121561028657610285610120565b5b600082013567ffffffffffffffff8111156102a4576102a3610125565b5b6102b084828501610242565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156102f35780820151818401526020810190506102d8565b83811115610302576000848401525b50505050565b6000610313826102b9565b61031d81856102c4565b935061032d8185602086016102d5565b61033681610134565b840191505092915050565b6000602082019050818103600083015261035b8184610308565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806103aa57607f821691505b6020821081036103bd576103bc610363565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026104257fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826103e8565b61042f86836103e8565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b600061047661047161046c84610447565b610451565b610447565b9050919050565b6000819050919050565b6104908361045b565b6104a461049c8261047d565b8484546103f5565b825550505050565b600090565b6104b96104ac565b6104c4818484610487565b505050565b5b818110156104e8576104dd6000826104b1565b6001810190506104ca565b5050565b601f82111561052d576104fe816103c3565b610507846103d8565b81016020851015610516578190505b61052a610522856103d8565b8301826104c9565b50505b505050565b600082821c905092915050565b600061055060001984600802610532565b1980831691505092915050565b6000610569838361053f565b9150826002028217905092915050565b610582826102b9565b67ffffffffffffffff81111561059b5761059a610145565b5b6105a58254610392565b6105b08282856104ec565b600060209050601f8311600181146105e357600084156105d1578287015190505b6105db858261055d565b865550610643565b601f1984166105f1866103c3565b60005b82811015610619578489015182556001820191506020850194506020810190506105f4565b868310156106365784890151610632601f89168261053f565b8355505b6001600288020188555050505b50505050505056fea264697066735822122063a799767cdb6f0b8c7d8de594c68a9f16fa62a935a21d9abb8ed592f0c2394664736f6c634300080f0033",
		"contractName": "HelloWorld"
	}`,
	"MyTestToken": `{
		"abi": "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
		"bin": "60806040523480156200001157600080fd5b506040518060400160405280600b81526020017f4d7954657374546f6b656e0000000000000000000000000000000000000000008152506040518060400160405280600381526020017f4d5454000000000000000000000000000000000000000000000000000000000081525081600390816200008f9190620004ed565b508060049081620000a19190620004ed565b505050620000e233620000b9620000e860201b60201c565b600a620000c7919062000764565b612710620000d69190620007b5565b620000f160201b60201c565b62000924565b60006012905090565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160362000163576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200015a9062000877565b60405180910390fd5b62000177600083836200026960201b60201c565b80600260008282546200018b919062000899565b92505081905550806000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254620001e2919062000899565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405162000249919062000907565b60405180910390a362000265600083836200026e60201b60201c565b5050565b505050565b505050565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620002f557607f821691505b6020821081036200030b576200030a620002ad565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620003757fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000336565b62000381868362000336565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620003ce620003c8620003c28462000399565b620003a3565b62000399565b9050919050565b6000819050919050565b620003ea83620003ad565b62000402620003f982620003d5565b84845462000343565b825550505050565b600090565b620004196200040a565b62000426818484620003df565b505050565b5b818110156200044e57620004426000826200040f565b6001810190506200042c565b5050565b601f8211156200049d57620004678162000311565b620004728462000326565b8101602085101562000482578190505b6200049a620004918562000326565b8301826200042b565b50505b505050565b600082821c905092915050565b6000620004c260001984600802620004a2565b1980831691505092915050565b6000620004dd8383620004af565b9150826002028217905092915050565b620004f88262000273565b67ffffffffffffffff8111156200051457620005136200027e565b5b620005208254620002dc565b6200052d82828562000452565b600060209050601f83116001811462000565576000841562000550578287015190505b6200055c8582620004cf565b865550620005cc565b601f198416620005758662000311565b60005b828110156200059f5784890151825560018201915060208501945060208101905062000578565b86831015620005bf5784890151620005bb601f891682620004af565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60008160011c9050919050565b6000808291508390505b600185111562000662578086048111156200063a5762000639620005d4565b5b60018516156200064a5780820291505b80810290506200065a8562000603565b94506200061a565b94509492505050565b6000826200067d576001905062000750565b816200068d576000905062000750565b8160018114620006a65760028114620006b157620006e7565b600191505062000750565b60ff841115620006c657620006c5620005d4565b5b8360020a915084821115620006e057620006df620005d4565b5b5062000750565b5060208310610133831016604e8410600b8410161715620007215782820a9050838111156200071b576200071a620005d4565b5b62000750565b62000730848484600162000610565b925090508184048111156200074a5762000749620005d4565b5b81810290505b9392505050565b600060ff82169050919050565b6000620007718262000399565b91506200077e8362000757565b9250620007ad7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff84846200066b565b905092915050565b6000620007c28262000399565b9150620007cf8362000399565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156200080b576200080a620005d4565b5b828202905092915050565b600082825260208201905092915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006200085f601f8362000816565b91506200086c8262000827565b602082019050919050565b60006020820190508181036000830152620008928162000850565b9050919050565b6000620008a68262000399565b9150620008b38362000399565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115620008eb57620008ea620005d4565b5b828201905092915050565b620009018162000399565b82525050565b60006020820190506200091e6000830184620008f6565b92915050565b61126380620009346000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461016857806370a082311461019857806395d89b41146101c8578063a457c2d7146101e6578063a9059cbb14610216578063dd62ed3e14610246576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100fc57806323b872dd1461011a578063313ce5671461014a575b600080fd5b6100b6610276565b6040516100c39190610b1e565b60405180910390f35b6100e660048036038101906100e19190610bd9565b610308565b6040516100f39190610c34565b60405180910390f35b61010461032b565b6040516101119190610c5e565b60405180910390f35b610134600480360381019061012f9190610c79565b610335565b6040516101419190610c34565b60405180910390f35b610152610364565b60405161015f9190610ce8565b60405180910390f35b610182600480360381019061017d9190610bd9565b61036d565b60405161018f9190610c34565b60405180910390f35b6101b260048036038101906101ad9190610d03565b6103a4565b6040516101bf9190610c5e565b60405180910390f35b6101d06103ec565b6040516101dd9190610b1e565b60405180910390f35b61020060048036038101906101fb9190610bd9565b61047e565b60405161020d9190610c34565b60405180910390f35b610230600480360381019061022b9190610bd9565b6104f5565b60405161023d9190610c34565b60405180910390f35b610260600480360381019061025b9190610d30565b610518565b60405161026d9190610c5e565b60405180910390f35b60606003805461028590610d9f565b80601f01602080910402602001604051908101604052809291908181526020018280546102b190610d9f565b80156102fe5780601f106102d3576101008083540402835291602001916102fe565b820191906000526020600020905b8154815290600101906020018083116102e157829003601f168201915b5050505050905090565b60008061031361059f565b90506103208185856105a7565b600191505092915050565b6000600254905090565b60008061034061059f565b905061034d858285610770565b6103588585856107fc565b60019150509392505050565b60006012905090565b60008061037861059f565b905061039981858561038a8589610518565b6103949190610dff565b6105a7565b600191505092915050565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600480546103fb90610d9f565b80601f016020809104026020016040519081016040528092919081815260200182805461042790610d9f565b80156104745780601f1061044957610100808354040283529160200191610474565b820191906000526020600020905b81548152906001019060200180831161045757829003601f168201915b5050505050905090565b60008061048961059f565b905060006104978286610518565b9050838110156104dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d390610ec7565b60405180910390fd5b6104e982868684036105a7565b60019250505092915050565b60008061050061059f565b905061050d8185856107fc565b600191505092915050565b6000600160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610616576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060d90610f59565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610685576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161067c90610feb565b60405180910390fd5b80600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516107639190610c5e565b60405180910390a3505050565b600061077c8484610518565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff81146107f657818110156107e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107df90611057565b60405180910390fd5b6107f584848484036105a7565b5b50505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361086b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610862906110e9565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d19061117b565b60405180910390fd5b6108e5838383610a7b565b60008060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205490508181101561096b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109629061120d565b60405180910390fd5b8181036000808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550816000808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546109fe9190610dff565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051610a629190610c5e565b60405180910390a3610a75848484610a80565b50505050565b505050565b505050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610abf578082015181840152602081019050610aa4565b83811115610ace576000848401525b50505050565b6000601f19601f8301169050919050565b6000610af082610a85565b610afa8185610a90565b9350610b0a818560208601610aa1565b610b1381610ad4565b840191505092915050565b60006020820190508181036000830152610b388184610ae5565b905092915050565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b7082610b45565b9050919050565b610b8081610b65565b8114610b8b57600080fd5b50565b600081359050610b9d81610b77565b92915050565b6000819050919050565b610bb681610ba3565b8114610bc157600080fd5b50565b600081359050610bd381610bad565b92915050565b60008060408385031215610bf057610bef610b40565b5b6000610bfe85828601610b8e565b9250506020610c0f85828601610bc4565b9150509250929050565b60008115159050919050565b610c2e81610c19565b82525050565b6000602082019050610c496000830184610c25565b92915050565b610c5881610ba3565b82525050565b6000602082019050610c736000830184610c4f565b92915050565b600080600060608486031215610c9257610c91610b40565b5b6000610ca086828701610b8e565b9350506020610cb186828701610b8e565b9250506040610cc286828701610bc4565b9150509250925092565b600060ff82169050919050565b610ce281610ccc565b82525050565b6000602082019050610cfd6000830184610cd9565b92915050565b600060208284031215610d1957610d18610b40565b5b6000610d2784828501610b8e565b91505092915050565b60008060408385031215610d4757610d46610b40565b5b6000610d5585828601610b8e565b9250506020610d6685828601610b8e565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610db757607f821691505b602082108103610dca57610dc9610d70565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e0a82610ba3565b9150610e1583610ba3565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115610e4a57610e49610dd0565b5b828201905092915050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000610eb1602583610a90565b9150610ebc82610e55565b604082019050919050565b60006020820190508181036000830152610ee081610ea4565b9050919050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000610f43602483610a90565b9150610f4e82610ee7565b604082019050919050565b60006020820190508181036000830152610f7281610f36565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000610fd5602283610a90565b9150610fe082610f79565b604082019050919050565b6000602082019050818103600083015261100481610fc8565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000611041601d83610a90565b915061104c8261100b565b602082019050919050565b6000602082019050818103600083015261107081611034565b9050919050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b60006110d3602583610a90565b91506110de82611077565b604082019050919050565b60006020820190508181036000830152611102816110c6565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000611165602383610a90565b915061117082611109565b604082019050919050565b6000602082019050818103600083015261119481611158565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006111f7602683610a90565b91506112028261119b565b604082019050919050565b60006020820190508181036000830152611226816111ea565b905091905056fea26469706673582212200f4b167ce5cf1a23589f027d08f90587a72486f7afaffe792c651606dd0189ee64736f6c634300080f0033",
		"contractName": "MyTestToken"
	}`,
}

func GetContractJsonFilePath(contractName string) (string, error) {
	f, err := ioutil.TempFile("", "compiled-contract-json-")
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write([]byte(compiledContractsJson[contractName]))
	if err != nil {
		return "", err
	}
	return f.Name(), nil
}
