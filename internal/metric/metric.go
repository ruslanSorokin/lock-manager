package metric

type appMetric interface {
	SetAppVersion(string)
	SetAppEnvironment(string)
}

type serviceMetric interface {
	IncLockedTotal()
	IncUnlockedTotal()
}

type MetricI interface {
	appMetric
	serviceMetric
}
