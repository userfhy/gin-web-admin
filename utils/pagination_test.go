package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	gin.SetMode(gin.TestMode)
}

func newTestContext(rawURL string) *gin.Context {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	req := httptest.NewRequest("GET", rawURL, nil)
	c.Request = req
	return c
}

func TestGetPaginationSuccess(t *testing.T) {
	c := newTestContext("/list?p=2&n=10")
	pg, err := GetPagination(c)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pg.Page != 2 || pg.PageSize != 10 {
		t.Fatalf("page=%d size=%d", pg.Page, pg.PageSize)
	}
	if pg.Offset() != 10 || pg.Limit() != 10 {
		t.Fatalf("offset=%d limit=%d", pg.Offset(), pg.Limit())
	}
}

func TestGetPaginationAliasesAndDefaults(t *testing.T) {
	c := newTestContext("/list?pageNum=0&pageSize=0")
	pg, err := GetPagination(c, WithDefaultPage(3), WithDefaultPageSize(30))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pg.Page != 3 || pg.PageSize != 30 {
		t.Fatalf("default not applied: %#v", pg)
	}
	if pg.Offset() != 60 {
		t.Fatalf("offset mismatch: %d", pg.Offset())
	}
}

func TestGetPaginationMaxSize(t *testing.T) {
	c := newTestContext("/list?page=1&size=999")
	pg, err := GetPagination(c, WithMaxPageSize(200))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if pg.PageSize != 200 {
		t.Fatalf("max size not applied: %d", pg.PageSize)
	}
}

func TestGetPaginationInvalid(t *testing.T) {
	c := newTestContext("/list?p=abc")
	if _, err := GetPagination(c); err == nil {
		t.Fatal("expected error for invalid page")
	}

	c = newTestContext("/list?pageSize=xyz")
	if _, err := GetPagination(c); err == nil {
		t.Fatal("expected error for invalid size")
	}
}
