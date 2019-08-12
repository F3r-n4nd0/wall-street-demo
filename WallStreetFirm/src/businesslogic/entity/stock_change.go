package entity


type StockChange struct {
	UUID   string
	StockUUID  string
	OldValue float64
	NewValue float64
}