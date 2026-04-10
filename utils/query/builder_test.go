package query

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func TestBuilderFromQuery(t *testing.T) {
	req := httptest.NewRequest("GET", "/users?username=tony&status=1&roles=admin,editor", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	builder := NewBuilder().IsNull("deleted_at")
	err := builder.FromQuery(c, RuleSet{
		"username": {Field: "username", Op: OpLike},
		"status":   {Field: "status", Op: OpEqual, Parser: IntParser()},
		"roles":    {Field: "role_key", Op: OpIn},
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	where := builder.Build()
	if where["username like"].(string) != "%tony%" {
		t.Fatalf("like builder mismatch: %#v", where)
	}
	if where["status ="] != 1 {
		t.Fatalf("equal builder mismatch: %#v", where)
	}

	roles, ok := where["role_key in"].([]string)
	if !ok || len(roles) != 2 {
		t.Fatalf("expected in clause, got %#v", where["role_key in"])
	}
	if _, ok := where["deleted_at is"]; !ok {
		t.Fatal("missing default is null condition")
	}
}

func TestBuilderParserError(t *testing.T) {
	req := httptest.NewRequest("GET", "/users?status=abc", nil)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = req

	builder := NewBuilder()
	err := builder.FromQuery(c, RuleSet{
		"status": {Field: "status", Op: OpEqual, Parser: IntParser()},
	})
	if err == nil {
		t.Fatal("expected parser error")
	}
}

func TestIntEnumParser(t *testing.T) {
	parser := IntEnumParser(0, 1)
	if _, err := parser("2"); err == nil {
		t.Fatal("expected enum parser error")
	}
	val, err := parser("1")
	if err != nil || val.(int) != 1 {
		t.Fatalf("expected 1 got %v err=%v", val, err)
	}
}
