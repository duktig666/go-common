from web3 import Web3

# 连接到以太坊节点
w3 = Web3(Web3.HTTPProvider("https://api.chainup.net/base/mainnet/dc59a7e90f05424a81d825d033d0f2de"))

# 以太坊区块号（示例中以最新区块号为例）
block_number = w3.eth.block_number

# 获取指定区块的交易列表
block = w3.eth.get_block(block_identifier=block_number, full_transactions=True)

print("block_number:", block_number)

if block is not None:
    transactions = block['transactions']

    for tx in transactions:
        print("Transaction Hash:", tx['hash'].hex())
        print("From:", tx['from'])
        print("To:", tx['to'])
        print("Value (in Wei):", tx['value'])
        print("Gas Price (in Wei):", tx['gasPrice'])
        print("Gas Used:", tx['gas'])
        print("\n")
else:
    print(f"Block {block_number} not found.")
