package playgrounds

import (
	"log"
	"math/big"
	"net/http"

	"github.com/niccolofant/ic-arb/core/icp"
	"github.com/niccolofant/ic-arb/core/icp/icrc1"
	"github.com/niccolofant/ic-arb/core/icp/icrc2"
	"github.com/niccolofant/ic-arb/core/icp/kongswap/kong"
)

func TestKongswap() {
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
	kong, err := kong.NewWithMetadata(agent)
	if err != nil {
		panic(err)
	}

	amountIn := big.NewInt(81_021_993)

	quoteResult, err := kong.OneStepQuote(icp.DexQuoteParams{
		FromToken: exeToken,
		ToToken:   icpToken,
		AmountIn:  amountIn,
	}, nil)
	if err != nil {
		panic(err)
	}

	log.Println("quote: ", quoteResult)

	result, err := kong.OneStepSwap(icp.DexSwapParams{
		FromToken:    exeToken,
		ToToken:      icpToken,
		AmountIn:     amountIn,
		AmountOutMin: quoteResult,
	}, nil)
	if err != nil {
		panic(err)
	}

	log.Println("swap: ", result)
}
