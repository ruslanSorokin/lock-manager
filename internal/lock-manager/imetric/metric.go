package imetric

type ServiceMetricI interface {
	IncLockedTotal()
	IncUnlockedTotal()
}
