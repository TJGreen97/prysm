package blocks

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	nextWithdrawalIndex = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "next_withdrawal_index",
		Help: "The next withdrawal sweep index",
	})
	nextWithdrawalValidatorIndex = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "next_withdrawal_validator_index",
		Help: "The next validator index to evaluate for withdrawals",
	})
)
