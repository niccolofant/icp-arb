package playgrounds

import (
	"log"
	"math/big"
	"net/http"

	"github.com/niccolofant/ic-arb/core/icp"
	"github.com/niccolofant/ic-arb/core/icp/icrc2"
	"github.com/niccolofant/ic-arb/core/icp/ntn/ntnswap"
)

func TestNtnswap() {
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

	ntnToken, err := icrc2.NewWithMetadata(
		agent,
		icp.MustDecodePrincipal("f54if-eqaaa-aaaaq-aacea-cai"),
		icp.TokenMetadata{
			Name:     "NTN",
			Symbol:   "NTN",
			Fee:      big.NewInt(10000),
			Decimals: 8,
			Standard: icp.TokenStandardICRC2,
		},
	)
	if err != nil {
		panic(err)
	}

	_ = ntnToken

	ntnswap, err := ntnswap.NewWithMetadata(agent)
	if err != nil {
		panic(err)
	}

	balanceBefore, err := icpToken.BalanceOf(agent.Sender())
	if err != nil {
		panic(err)
	}

	log.Println("balance before: ", balanceBefore)

	// withdrawnAmount, err := ntnswap.Withdraw(
	// 	icpToken,
	// 	big.NewInt(0.1*icp.E8S),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println("withdrawn amount: ", withdrawnAmount)

	// deposited, err := ntnswap.Deposit(
	// 	icpToken,
	// 	big.NewInt(0.01*icp.E8S),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	//log.Println(deposited)

	amountIn := big.NewInt(79792708)

	quoteResult, err := ntnswap.OneStepQuote(icp.DexQuoteParams{
		FromToken: icpToken,
		ToToken:   ntnToken,
		AmountIn:  amountIn,
	}, nil)
	if err != nil {
		panic(err)
	}

	log.Println("quote result: ", quoteResult)

	swappedResult, err := ntnswap.OneStepSwap(icp.DexSwapParams{
		FromToken:    icpToken,
		ToToken:      ntnToken,
		AmountIn:     amountIn,
		AmountOutMin: quoteResult,
	}, nil)
	if err != nil {
		panic(err)
	}

	log.Println("swap result: ", swappedResult)

	balanceAfter, err := icpToken.BalanceOf(agent.Sender())
	if err != nil {
		panic(err)
	}

	log.Println("balance after: ", balanceAfter)

	// quoteResult, err := ntnswap.Quote(ntnToken, icpToken, big.NewInt(119493788))
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println("quote result: ", quoteResult)

	// swappedResult, err := ntnswap.Swap(ntnToken, icpToken, big.NewInt(119493788), quoteResult)
	// if err != nil {
	// 	panic(err)
	// }

	// log.Println("swap result: ", swappedResult)
}
