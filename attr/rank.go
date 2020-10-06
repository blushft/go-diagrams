package attr

type RankType string

func (r RankType) String() string {
	return string(r)
}

const (
	Same   RankType = "same"
	Min    RankType = "min"
	Source RankType = "source"
	Max    RankType = "max"
	Sink   RankType = "sink"
)

type RankDirection string

func (r RankDirection) String() string {
	return string(r)
}

const (
	TopToBottom RankDirection = "TB"
	LeftToRight RankDirection = "LR"
	BottomToTop RankDirection = "BT"
	RightToLeft RankDirection = "RL"
)
