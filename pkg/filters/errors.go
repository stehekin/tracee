package filters

import (
	"fmt"
)

func UnsupportedOperator(op Operator) error {
	return fmt.Errorf("failed to add filter: unsupported operator %s", op.String())
}

func InvalidPolicy(policy string) error {
	return fmt.Errorf("invalid policy: %s", policy)
}

func InvalidExpression(expression string) error {
	return fmt.Errorf("invalid filter expression: %s", expression)
}

func InvalidValue(value string) error {
	return fmt.Errorf("invalid filter value %s", value)
}

func InvalidEventName(event string) error {
	return fmt.Errorf("invalid event name in filter: %s", event)
}

func InvalidEventField(data string) error {
	return fmt.Errorf("invalid event data field: %s", data)
}

func InvalidScopeField(field string) error {
	return fmt.Errorf("invalid event scope field: %s", field)
}

func FailedToRetreiveHostNS() error {
	return fmt.Errorf("failed to retrieve host mount namespace")
}
