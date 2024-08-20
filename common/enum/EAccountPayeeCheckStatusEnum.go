package enum

type EAccountPayeeCheckStatusEnum int

const (
	WAIT EAccountPayeeCheckStatusEnum = iota
	AGREE
	REJECT
)

func (c EAccountPayeeCheckStatusEnum) GetDesc() string {
	switch c {
	case WAIT:
		return "未审合"
	case AGREE:
		return "以同意"
	case REJECT:
		return "以拒绝"
	default:
		return ""
	}
}
