package query

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// Operator 定义过滤操作符
type Operator string

const (
	OpEqual    Operator = "="
	OpNotEqual Operator = "<>"
	OpLike     Operator = "like"
	OpGte      Operator = ">="
	OpLte      Operator = "<="
	OpIn       Operator = "in"
	OpIs       Operator = "is"
)

type Rule struct {
	Field      string
	Op         Operator
	Parser     func(string) (any, error)
	Pattern    string
	AllowEmpty bool
	Splitter   string
}

type RuleSet map[string]Rule

type Builder struct {
	where map[string]any
}

func NewBuilder() *Builder {
	return &Builder{where: map[string]any{}}
}

func (b *Builder) add(field string, op Operator, value any) *Builder {
	if value == nil && op != OpIs {
		return b
	}
	switch v := value.(type) {
	case string:
		if strings.TrimSpace(v) == "" {
			return b
		}
	}
	key := fmt.Sprintf("%s %s", field, op)
	b.where[key] = value
	return b
}

func (b *Builder) Equal(field string, value any) *Builder {
	return b.add(field, OpEqual, value)
}

func (b *Builder) Like(field string, value string) *Builder {
	return b.add(field, OpLike, fmt.Sprintf("%%%s%%", value))
}

func (b *Builder) GreaterOrEqual(field string, value any) *Builder {
	return b.add(field, OpGte, value)
}

func (b *Builder) LessOrEqual(field string, value any) *Builder {
	return b.add(field, OpLte, value)
}

func (b *Builder) In(field string, values []string) *Builder {
	if len(values) == 0 {
		return b
	}
	cleaned := make([]string, 0, len(values))
	for _, v := range values {
		if v = strings.TrimSpace(v); v != "" {
			cleaned = append(cleaned, v)
		}
	}
	if len(cleaned) == 0 {
		return b
	}
	key := fmt.Sprintf("%s %s", field, OpIn)
	b.where[key] = cleaned
	return b
}

func (b *Builder) IsNull(field string) *Builder {
	return b.add(field, OpIs, nil)
}

func (b *Builder) Build() map[string]any {
	return b.where
}

func (b *Builder) FromQuery(c *gin.Context, rules RuleSet) error {
	for param, rule := range rules {
		raw, exists := c.GetQuery(param)
		if !exists {
			continue
		}
		raw = strings.TrimSpace(raw)
		if raw == "" && !rule.AllowEmpty {
			continue
		}
		value := any(raw)
		if rule.Parser != nil {
			parsed, err := rule.Parser(raw)
			if err != nil {
				return fmt.Errorf("%s 参数无效: %w", param, err)
			}
			value = parsed
		}
		switch rule.Op {
		case OpLike:
			pattern := rule.Pattern
			if pattern == "" {
				pattern = "%%%s%%"
			}
			value = fmt.Sprintf(pattern, value)
		case OpIn:
			splitter := rule.Splitter
			if splitter == "" {
				splitter = ","
			}
			parts := strings.Split(raw, splitter)
			trimmed := make([]string, 0, len(parts))
			for _, p := range parts {
				if s := strings.TrimSpace(p); s != "" {
					trimmed = append(trimmed, s)
				}
			}
			if len(trimmed) == 0 {
				continue
			}
			value = trimmed
		}
		b.add(rule.Field, rule.Op, value)
	}
	return nil
}

// Parser helpers
func IntParser() func(string) (any, error) {
	return func(raw string) (any, error) {
		v, err := strconv.Atoi(strings.TrimSpace(raw))
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func UintParser() func(string) (any, error) {
	return func(raw string) (any, error) {
		v, err := strconv.ParseUint(strings.TrimSpace(raw), 10, 64)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}

func BoolParser() func(string) (any, error) {
	return func(raw string) (any, error) {
		switch strings.ToLower(strings.TrimSpace(raw)) {
		case "1", "true", "yes":
			return 1, nil
		case "0", "false", "no":
			return 0, nil
		default:
			return nil, fmt.Errorf("invalid bool: %s", raw)
		}
	}
}

func IntEnumParser(allowed ...int) func(string) (any, error) {
	allowedSet := make(map[int]struct{}, len(allowed))
	for _, v := range allowed {
		allowedSet[v] = struct{}{}
	}
	return func(raw string) (any, error) {
		v, err := strconv.Atoi(strings.TrimSpace(raw))
		if err != nil {
			return nil, err
		}
		if len(allowedSet) > 0 {
			if _, ok := allowedSet[v]; !ok {
				return nil, fmt.Errorf("值必须为 %v", allowed)
			}
		}
		return v, nil
	}
}
