package blc

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type CLI struct {
	BlockChain *BlockChain
}

func (cli *CLI) Run() {
	isValidArgs()

	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	createGenesisBlockCmd := flag.NewFlagSet("creategenesisblock", flag.ExitOnError)

	flagAddBlockData := addBlockCmd.String("data", "http://xxx.com", "交易数据")
	flagCreateGenesisBlockAddress := createGenesisBlockCmd.String("address", "", "创建创世区块的地址")

	switch os.Args[1] {
	case "addblock":
		err := addBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "printchain":
		err := printChainCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	case "creategenesisblock":
		err := createGenesisBlockCmd.Parse(os.Args[2:])
		if err != nil {
			log.Panic(err)
		}
	default:
		printUsage()
		os.Exit(1)
	}

	if addBlockCmd.Parsed() {
		if *flagAddBlockData == "" {
			printUsage()
			os.Exit(1)
		}
		cli.addBlock([]*Transaction{})
	}

	if printChainCmd.Parsed() {
		cli.printChain()
	}

	if createGenesisBlockCmd.Parsed() {
		if *flagCreateGenesisBlockAddress == "" {
			fmt.Println("地址不能为空....")
			printUsage()
			os.Exit(1)
		}
		cli.createGenesisBlockChain(*flagCreateGenesisBlockAddress)
	}
}

func (cli *CLI) addBlock(txs []*Transaction) {
	if !DBExists() {
		fmt.Println("数据库不存在.......")
		os.Exit(1)
	}
	blockChain := BlockChainObject()
	defer blockChain.DB.Close()
	blockChain.AddBlockToBlockChain(txs)
}

func (cli *CLI) printChain() {
	if !DBExists() {
		fmt.Println("数据库不存在.......")
		os.Exit(1)
	}
	blockChain := BlockChainObject()
	defer blockChain.DB.Close()
	blockChain.PrintChain()
}

func (cli *CLI) createGenesisBlockChain(address string) {
	CreateBlockChainWithGenesisBlock(address)
}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("\tcreategenesisblock -address DATA -- 创建创世区块")
	fmt.Println("\taddblock -data DATA -- 交易数据")
	fmt.Println("\tprintchain -- 输出区块信息")
}

func isValidArgs() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}
}
