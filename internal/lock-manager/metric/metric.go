package metric

type ServiceMetricI interface {
	IncLockedTotal()
	IncUnlockedTotal()
}
