package request_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/donnigundala/dg-http/request"
	"github.com/gin-gonic/gin"
)

func TestParam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/users/:id", func(c *gin.Context) {
		id := request.Param(c, "id")
		if id != "123" {
			t.Errorf("Expected id '123', got '%s'", id)
		}
		c.String(200, id)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/123", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestParamInt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/users/:id", func(c *gin.Context) {
		id := request.ParamInt(c, "id", 0)
		if id != 456 {
			t.Errorf("Expected id 456, got %d", id)
		}
		c.JSON(200, gin.H{"id": id})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/users/456", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestQuery(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/search", func(c *gin.Context) {
		query := request.Query(c, "q", "default")
		if query != "test" {
			t.Errorf("Expected query 'test', got '%s'", query)
		}
		c.String(200, query)
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/search?q=test", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestQueryInt(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/items", func(c *gin.Context) {
		page := request.QueryInt(c, "page", 1)
		if page != 5 {
			t.Errorf("Expected page 5, got %d", page)
		}
		c.JSON(200, gin.H{"page": page})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/items?page=5", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestQueryWithDefault(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/items", func(c *gin.Context) {
		page := request.QueryInt(c, "page", 1)
		if page != 1 {
			t.Errorf("Expected default page 1, got %d", page)
		}
		c.JSON(200, gin.H{"page": page})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/items", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestQueryBool(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/items", func(c *gin.Context) {
		active := request.QueryBool(c, "active", false)
		if !active {
			t.Error("Expected active to be true")
		}
		c.JSON(200, gin.H{"active": active})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/items?active=true", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}

func TestQueryArray(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()

	router.GET("/items", func(c *gin.Context) {
		tags := request.QueryArray(c, "tags")
		if len(tags) != 3 {
			t.Errorf("Expected 3 tags, got %d", len(tags))
		}
		if tags[0] != "go" || tags[1] != "gin" || tags[2] != "web" {
			t.Errorf("Unexpected tags: %v", tags)
		}
		c.JSON(200, gin.H{"tags": tags})
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/items?tags=go&tags=gin&tags=web", nil)
	router.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status 200, got %d", w.Code)
	}
}
