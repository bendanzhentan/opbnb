package bridge

import (
	"github.com/ethereum-optimism/optimism/indexer/config"
	"github.com/ethereum-optimism/optimism/indexer/database"
	"github.com/ethereum-optimism/optimism/indexer/node"
	"github.com/ethereum/go-ethereum/log"
)

func L2ProcessAutoWithdrawEvents(log log.Logger, db *database.DB, metrics L1Metricer, l1EthClient node.EthClient, l2Contracts config.L2Contracts) error {
	limit := 10
	log.Info("bilibili", "start db.BridgeTransactions.L2TransactionWithdrawalsFilter")
	withdrawals, err := db.BridgeTransactions.L2TransactionWithdrawalsFilter(&limit)
	if err != nil {
		return err
	}

	log.Info("bilibili", "start db.BridgeTransactions.L2TransactionWithdrawalsFilter for loop")
	for _, withdrawal := range withdrawals {
		log.Info("bilibili", "withdrawal", withdrawal.InitiatedL2EventGUID)

		err = db.BridgeTransactions.MarkL2TransactionWithdrawalIsAutoWithdraw(withdrawal.WithdrawalHash, false)
		if err != nil {
			log.Crit("bilibili", "MarkL2TransactionWithdrawalIsAutoWithdraw error", err)
		}
	}

	return nil
}
