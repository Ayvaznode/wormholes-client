package test

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/wormholes-org/wormholes-client/client"
	"log"
	"strings"
	"testing"
	"time"
)

const (
	endpoint         = "http://192.168.4.237:8574"
	priKey           = "7c6786275d6011adb6288587757653d3f9061275bafc2c35ae62efe0bc4973e9"
	buyerPriKey      = "f616c4d20311a2e73c67ef334630f834b7fb42304a1d4448fb2058e9940ecc0a"
	buyerAddress     = "0x44d952db5dfb4cbb54443554f4bb9cbebee2194c"
	sellerPriKey     = "e04d9e04569d1de38be6b0dbced9413ebf86d33d3670c6db965726b46de0572a"
	sellerAddress    = "0xFFF531a2DA46d051FdE4c47F042eE6322407DF3f"
	exchangerPriKey  = "74960499b76daa6c987fb3872619fe28d875d5c64fd96bbb2b9c0ae676eb2c45"
	exchangeAddress  = "0x83c43f6F7bB4d8E429b21FF303a16b4c99A59b05"
	exchangerPriKey1 = "8c9c4464a685583b1ddcb30bfd991444248eb492d10ceca647fdd41329499b49"
	exchangeAddress1 = "0xB685EB3226d5F0D549607D2cC18672b756fd090c"
)

func TestNewClient(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	_ = worm
}

//Recharge
func TestRecharge(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.NormalTransaction("0x814920c33b1a037F91a16B126282155c6F92A10F", 100, "")
	fmt.Println(rs)
}

//Mint
//NFT mint 0
func TestMint(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.Mint(10, "/ipfs/ddfd90be9408b4", exchangeAddress)
	fmt.Println(rs)
}

//0x8fa2d4b70013407012d002fa395939cb0d322553e4848aaae78d4fad638bef55

//Transfer
//NFT transfer 1
func TestTransfer(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.Transfer("0x0000000000000000000000000000000000000001", sellerAddress)
	fmt.Println(rs)
}

//0x5e8dd659b0ceb95ab53ce32d37daa8688accab601ce58c75e706f08bb47617f4

//Author Single
//NFT authorization 2
func TestAuthor(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.Author("0x0000000000000000000000000000000000000002", exchangeAddress)
	fmt.Println(rs)
}

//0x2657d46f0c4ef16cadbc6842896c1b50f41333d6a247ee43e5085da5d7e3feff

//AuthorRevoke
//Cancel a single authorization 3
func TestAuthorRevoke(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.AuthorRevoke("0x0000000000000000000000000000000000000002", exchangeAddress)
	fmt.Println(rs)
}

//0xe043dc7d8505d01f6cd949b7a7cc4ed685a9e1b640195801c3c6265b7d11efee

//AccountAuthor
//All NFTs under the authorized account 4
func TestAccountAuthor(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.AccountAuthor(exchangeAddress)
	fmt.Println(rs)
}

//0x6b42237b9dad13211d89f1e6c66cf947bb371f407a4621ffcf7fd73e385f6fea

//AccountAuthorRevoke
//Cancel all NFTs under the authorized account 5
func TestAccountAuthorRevoke(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.AccountAuthorRevoke(exchangeAddress)
	fmt.Println(rs)
}

//0x1dee05dff7ea39874ed8401c91288ae627b56ae1df6dc4c26a856fafab0447c5

//SNFTToERB
//Fragment NFT exchange 6
func TestSNFTToERB(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.SNFTToERB("0x8000000000000000000000000000000000000004")
	fmt.Println(rs)
}

//0x77ff920a3a649378e4c7a58644bece643e379113b2bc99257b894a29e220e157

//TokenPledge
//ERB pledge 9
func TestTokenPledge(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.TokenPledge([]byte(""), "")
	fmt.Println(rs)
}

//0x6ceb02802455ab959964866410f37a2f0fcd78e7e64e87d6c9d8102de7f9974b

//TokenRevokesPledge
//ERB revokes pledge 10
func TestTokenRevokesPledge(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.TokenRevokesPledge()
	fmt.Println(rs)
}

//0xcbd19c386d8b5944d4a88017680239651edefc527e4ba2c8762ab0df2333a7ca

//Open
//Open an exchange 11
func TestOpen(t *testing.T) {
	worm := client.NewClient(exchangerPriKey, endpoint)
	rs, _ := worm.Open(10, "wormholes", "www.kang123456.com")
	fmt.Println(rs)
}

//0xcccd6c9499224e7d216f3bd230447900550b07345841eebd2e62b613f7bd924f

//Close
//close a exchange 12
func TestClose(t *testing.T) {
	worm := client.NewClient(exchangerPriKey, endpoint)
	rs, _ := worm.Close()
	fmt.Println(rs)
}

//0x7d5b3653b67d298e7cce82b92fa720224e93941315528ef19a36c4daf1efe929

//InsertNFTBlock
//Injecting System NFT Fragments 13
func TestInsertNFTBlock(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.InsertNFTBlock("wormholes2", "0x640001", 6553600, 20, "0xab7624f47fd7dadb6b8e255d06a2f10af55990fe")
	fmt.Println(rs)
}

//0x61cd018d6e70af47c6204fea18db5b33fdecc92162cca66b0089783733809e84

//TransactionNFT 14
func TestTransactionNFT(t *testing.T) {
	worm := client.NewClient(buyerPriKey, endpoint)
	number, _ := worm.BlockNumber(context.Background())
	blockNumber := fmt.Sprintf("0x%x", number+10)
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000002", "0x8b07aff2327a3B7e2876D899caFac99f7AE16B10", blockNumber, "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println("sign ", string(buyer))

	worm1 := client.NewClient(sellerPriKey, endpoint)
	rs, _ := worm1.TransactionNFT(buyer, buyerAddress)
	fmt.Println(rs)
}

//0xc9c4e6652ba411a0435d2e3187f019329b084734f19ae6699ee7f1fa9a92123b

//BuyerInitiatingTransaction 15
func TestBuyerInitiatingTransaction(t *testing.T) {
	worm := client.NewClient(sellerPriKey, "")
	seller1, err := worm.Wallet.SignSeller1("0x38D7EA4C68000", "0x0000000000000000000000000000000000000003", "0x8b07aff2327a3B7e2876D899caFac99f7AE16B10", "0x677")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println("sign ", string(seller1))

	worm1 := client.NewClient(buyerPriKey, endpoint)
	rs, _ := worm1.BuyerInitiatingTransaction(seller1)
	fmt.Println(rs)
}

//0xfb9cf0100340c9bf965fc0f8ef44bb8a75af58175deab0dcff3979a97a8ebefa

//FoundryTradeBuyer 16
func TestFoundryTradeBuyer(t *testing.T) {
	worm := client.NewClient(sellerPriKey, "")
	seller2, err := worm.Wallet.SignSeller2("0x38D7EA4C68000", "0xa", "/ipfs/qqqqqqqqqq", "0", "0x8b07aff2327a3B7e2876D899caFac99f7AE16B10", "0x677")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println("sign ", string(seller2))

	worm1 := client.NewClient(buyerPriKey, endpoint)
	rs, _ := worm1.FoundryTradeBuyer(seller2)
	fmt.Println(rs)
}

//0x4634d6bbc36b9444914a259c2acf0410af0b99122baef30d7a8701a496bc3b6c

//FoundryExchange 17
func TestFoundryExchange(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "", exchangeAddress, "0xa", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm1 := client.NewClient(sellerPriKey, "")
	seller2, err := worm1.Wallet.SignSeller2("0x38D7EA4C68000", "0xa", "/ipfs/qqqqqqqqqq", "0", exchangeAddress, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm2 := client.NewClient(exchangerPriKey, endpoint)
	rs, _ := worm2.FoundryExchange(buyer, seller2, buyerAddress)
	fmt.Println(rs)
}

//0x70853466fdf5dc4476fab34b79f9be2e66f0448789937094de0b0aa5f3345e8c

//ftExchangeMatch  18
func TestNftExchangeMatch(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000004", exchangeAddress, "0xa", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm1 := client.NewClient(sellerPriKey, "")
	seller, err := worm1.Wallet.SignSeller1("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000004", exchangeAddress, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm2 := client.NewClient(exchangerPriKey, "")
	exchangeAuth, err := worm2.Wallet.SignExchanger(exchangeAddress, exchangeAddress1, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm3 := client.NewClient(exchangerPriKey1, endpoint)
	rs, _ := worm3.NftExchangeMatch(buyer, seller, exchangeAuth, buyerAddress)
	fmt.Println(rs)
}

//0xf11e024297b89e6dfd02bc2da4680cea353ea6956c3ea9084afa40d58477932f

//FoundryExchangeInitiated 19
func TestFoundryExchangeInitiated(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "", exchangeAddress, "0xa", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println(string(buyer))

	worm1 := client.NewClient(sellerPriKey, "")
	seller2, err := worm1.Wallet.SignSeller2("0x38D7EA4C68000", "0xa", "/ipfs/qqqqqqqqqq", "0", exchangeAddress, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println(string(seller2))

	worm2 := client.NewClient(exchangerPriKey, "")
	exchangeAuth, err := worm2.Wallet.SignExchanger(exchangeAddress, exchangeAddress1, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println(string(exchangeAuth))

	worm3 := client.NewClient(exchangerPriKey1, endpoint)
	rs, _ := worm3.FoundryExchangeInitiated(buyer, seller2, exchangeAuth, buyerAddress)
	fmt.Println(rs)
}

//0xc9cc570057faf1edd83f48833520f9d546e4972083ee705152b5f35630f1588d

//FtDoesNotAuthorizeExchanges 20
func TestNFTDoesNotAuthorizeExchanges(t *testing.T) {
	worm := client.NewClient(buyerPriKey, "")
	buyer, err := worm.Wallet.SignBuyer("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000001", exchangeAddress, "0xa", "")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm1 := client.NewClient(sellerPriKey, "")
	seller1, err := worm1.Wallet.SignSeller1("0xde0b6b3a7640000", "0x0000000000000000000000000000000000000001", exchangeAddress, "0xa")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	worm2 := client.NewClient(exchangerPriKey, endpoint)

	rs, _ := worm2.NFTDoesNotAuthorizeExchanges(buyer, seller1, buyerAddress)
	fmt.Println(rs)
}

//0x95615a6c7a164537257492c112a9fcd99907315893706a1b104456d9e3aa8af6

//AdditionalPledgeAmount 21
func TestAdditionalPledgeAmount(t *testing.T) {
	worm := client.NewClient(exchangerPriKey, endpoint)
	rs, _ := worm.AdditionalPledgeAmount(100)
	fmt.Println(rs)
}

//0x25f2ed8cf5f1041be9e71d483a32b01fd3f7820ec59e0c060830214c53fea5f9

//AdditionalPledgeAmount 22
func TestRevokesPledgeAmount(t *testing.T) {
	worm := client.NewClient(exchangerPriKey, endpoint)
	rs, _ := worm.RevokesPledgeAmount(100)
	fmt.Println(rs)
}

//0xd2c7f943f0f5364b0928c518e7b6de7491c0e8efb6abf912a17e6860f70ebec1

//VoteOfficialNFT
func TestVoteOfficialNFT(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.VoteOfficialNFT("wormholes2", "0x640001", 6553600, 20, "0xab7624f47fd7dadb6b8e255d06a2f10af55990fe")
	fmt.Println(rs)
}

//VoteOfficialNFTByApprovedExchanger
func TestVoteOfficialNFTByApprovedExchanger(t *testing.T) {
	worm := client.NewClient(exchangerPriKey, "")
	exchangeAuth, err := worm.Wallet.SignExchanger(exchangeAddress, exchangeAddress1, "0x0")
	if err != nil {
		log.Fatalln("Signing failed")
	}

	fmt.Println(string(exchangeAuth))
	worm1 := client.NewClient(exchangeAddress1, endpoint)
	rs, _ := worm1.VoteOfficialNFTByApprovedExchanger("wormholes2", "0x640001", 6553600, 20, "0xab7624f47fd7dadb6b8e255d06a2f10af55990fe", exchangeAuth)
	fmt.Println(rs)
}

//ChangeRewardsType
//change revenue model
func TestUnforzenAccount(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	rs, _ := worm.UnforzenAccount()
	fmt.Println(rs)
}

//AccountDelegate
//Delegate large accounts to small accounts
func TestAccountDelegate(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	proxySign, _ := worm.Wallet.SignDelegate("address", "pledgeAccount")
	rs, _ := worm.AccountDelegate(proxySign, buyerAddress)
	fmt.Println(rs)
}

func TestGetBalance(t *testing.T) {
	worm := client.NewClient(priKey, endpoint)
	balance, _ := worm.Balance(context.Background(), exchangeAddress)
	fmt.Println(balance)
}

func TestCheckNFTPool(t *testing.T) {
	worm := client.NewClient("c1e74da8e26c5a60870089f59695a1b243887f9d23571d24c7f011b8eb068768", "http://192.168.4.240:8561")

	var flag bool
	num := int64(22)
	for {
		if flag {
			break
		} else {
			current, _ := worm.BlockNumber(context.Background())
			if uint64(num) > current {
				time.Sleep(time.Second * 5)
			} else {
				fmt.Println("num: ", num)
				//res1, _ := worm.NFT.GetBlockBeneficiaryAddressByNumber(context.Background(), num)
				//for _, miners := range *res1 {
				//	if miners.Address == common.HexToAddress("0xEE3168308949237d395202F134C4243630ebB4A8") {
				//		fmt.Println("miner", miners.Address)
				//		flag = true
				//		//break
				//	}
				//}

				//res1, _ := worm.NFT.GetValidators(context.Background(), num)
				//for _, validator := range res1.Validators {
				//	fmt.Println("validator", validator)
				//	if validator.Addr == common.HexToAddress("0xA7aa3f181aebE59ca697D803B2197cfA50A3913E") {
				//		fmt.Println("miner", validator) //0x0F7094Cf6391273AAC99b478b8Eca9D636BBbf0c
				//		flag = true
				//		break
				//	}
				//}

				res1, _ := worm.GetActiveLivePool(context.Background(), uint64(num))
				for _, miners := range res1.ActiveMiners {
					fmt.Println("miner", miners)
					if miners.Address == common.HexToAddress("0xA7aa3f181aebE59ca697D803B2197cfA50A3913E") {
						fmt.Println("miner", miners) //0x0F7094Cf6391273AAC99b478b8Eca9D636BBbf0c
						flag = true
						break
					}
				}

				//res1, _ := worm.NFT.QueryMinerProxy(context.Background(), num, "0xA7F60Adc80E09F71a7A56044003a2B606Ed1Cac2")
				//for _, miners := range res1 {
				//	if miners.Address == common.HexToAddress("0x279c59A0DC597276bac3D160Cb1596beFA46bad2") {
				//		fmt.Println("miner", miners)
				//		flag = true
				//		break
				//	}
				//}
				num++
			}
		}
	}
}

func TestGetSNFT(t *testing.T) {
	worm := client.NewClient("c1e74da8e26c5a60870089f59695a1b243887f9d23571d24c7f011b8eb068768", "http://129.226.154.223:8560")
	num := int64(40000)
	snfts := make([]string, 0)
	for {
		fmt.Print("blockNumber: ", num)
		if len(snfts) != 0 {
			fmt.Print(" snft:")
			for _, v := range snfts {
				fmt.Print(" ", v)
			}
		} else {
			fmt.Print(" mei")
		}
		fmt.Println()
		current, _ := worm.BlockNumber(context.Background())
		if num > int64(current) {
			time.Sleep(time.Second * 10)
		}

		res1, _ := worm.GetBlockBeneficiaryAddressByNumber(context.Background(), num)
		if res1 != nil {
			for _, info := range *res1 {
				if strings.EqualFold(info.Address.String(), "0xEb1432C1066D2919ec8C1af5a6Dad5b088A11ae1") {
					snfts = append(snfts, info.NftAddress.String())
				}
			}
		}
		num++
	}
}
