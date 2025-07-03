package rabbitmq

const (
	//交換機
	TUT_LOSE_BETSLIP_EXCHANGE = "tut.lose_betslip"
	TUT_WIN_BETSLIP_EXCHANGE  = "tut.win_betslip"

	//隊列
	LOSE_BETSLIP_QUEUE = "lose_betslip_queue"
	WIN_BETSLIP_QUEUE  = "win_betslip_queue"

	//綁定RoutingKey
	LOSE_BETSLIP_ROUTING_KEY = "lose_betslip-routing-key"

	WIN_BETSLIP_ROUTING_KEY = "win_betslip-routing-key"
)

type ExchangeConfig struct {
	Name string
	Kind string // topic, fanout, direct
}

type QueueBinding struct {
	QueueName    string
	RoutingKey   string
	ExchangeName string
	HandlerFunc  func([]byte) error
}
type MQConfig struct {
	URL       string
	Exchanges []ExchangeConfig
	Bindings  []QueueBinding
}
