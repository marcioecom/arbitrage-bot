package strategy

type IStrategy interface {
	// Start starts the strategy
	Start()
	// Stop stops the strategy
	Stop()
}
