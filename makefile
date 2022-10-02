init: 
	solcjs --bin --abi contract/contract.sol -o build

gen: 
	abigen --bin=build/contract_contract_sol_Store.bin --abi=build/contract_contract_sol_Store.abi --pkg=contract --out=gen/eth-app.go