package playgrounds

import (
	"log"
	"math/big"
	"net/http"

	"github.com/niccolofant/ic-arb/icp"
	icpswap_pair "github.com/niccolofant/ic-arb/icp/icpswap/pair"
	"github.com/niccolofant/ic-arb/icp/icrc1"
	"github.com/niccolofant/ic-arb/icp/icrc2"
)

func TestIcpswap() {
	id, err := icp.LoadIntentity("")
	if err != nil {
		panic(err)
	}

	agent, err := icp.NewAgent(id, http.DefaultClient)
	if err != nil {
		panic(err)
	}

	icpToken, err := icrc2.NewWithMetadata(
		agent,
		icp.LedgerPrincipal,
		icp.TokenMetadata{
			Name:     "ICP",
			Symbol:   "ICP",
			Fee:      big.NewInt(0.0001 * icp.E8S),
			Decimals: 8,
			Standard: icp.TokenStandardICP,
		},
	)
	if err != nil {
		panic(err)
	}

	// ckbtcToken, err := icrc2.NewWithMetadata(
	// 	agent,
	// 	icp.MustDecodePrincipal("mxzaz-hqaaa-aaaar-qaada-cai"),
	// 	icp.TokenMetadata{
	// 		Name:     "ckBTC",
	// 		Symbol:   "ckBTC",
	// 		Fee:      big.NewInt(10),
	// 		Decimals: 8,
	// Standard: icp.TokenStandardICRC2,
	// 	},
	// )
	// if err != nil {
	// 	panic(err)
	// }

	exeToken, err := icrc1.NewWithMetadata(
		agent,
		icp.MustDecodePrincipal("rh2pm-ryaaa-aaaan-qeniq-cai"),
		icp.TokenMetadata{
			Name:     "EXE",
			Symbol:   "EXE",
			Fee:      big.NewInt(100000),
			Decimals: 8,
			Standard: icp.TokenStandardICRC1,
		},
	)
	if err != nil {
		panic(err)
	}

	// icpCkbtcPair, err := icpswap_pair.NewWithMetadata(
	// 	agent,
	// 	icp.MustDecodePrincipal("xmiu5-jqaaa-aaaag-qbz7q-cai"),
	// 	ckbtcToken,
	// 	icpToken,
	// )
	icpExePair, err := icpswap_pair.NewWithMetadata(
		agent,
		icp.MustDecodePrincipal("dlfvj-eqaaa-aaaag-qcs3a-cai"),
		exeToken,
		icpToken,
	)
	if err != nil {
		panic(err)
	}

	amountIn := big.NewInt(668552133)

	quoteResult, err := icpExePair.Quote(exeToken, icpToken, amountIn)
	if err != nil {
		panic(err)
	}

	log.Println("quote: ", quoteResult)

	result, err := icpExePair.Swap(exeToken, icpToken, amountIn, quoteResult)
	if err != nil {
		panic(err)
	}

	log.Println("swap: ", result)
}

// dfx canister call --network ic dlfvj-eqaaa-aaaag-qcs3a-cai deposit '(
//                 record {
//                         fee = 100000;
//                         token = "rh2pm-ryaaa-aaaan-qeniq-cai";
//                         amount = 4_795_663_138;
//                 }
//         )'

// dfx canister call --network ic rh2pm-ryaaa-aaaan-qeniq-cai icrc1_transfer '(
// record {
//   to = record {
//     owner = principal "k3ynf-g5k7e-kmmhl-pk6gp-mo6zq-2itpa-4bgo7-7cs54-mwdp3-lsk6u-rqe";
//     subaccount = null;
//   };
//   fee = null;
//   spender_subaccount = null;
//   from = record {
//     owner = principal "2w4lu-h3ht5-x3wbu-udqvf-wqczf-srnyt-stycn-6sci5-5lz77-x3ows-kqe";
//     subaccount = null;
//   };
//   memo = null;
//   created_at_time = null;
//   amount = 668552133;
// })'
