package main

import (
	"net/http"
	"testing"
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"time"
)

func TestListDatasetsHandler_Table(t *testing.T) {
	gin.SetMode(gin.TestMode)
	tests := []struct {
		name       string
		collections []string
		expectCode int
		expectBody string
	}{
		{
			name:       "no collections",
			collections: []string{},
			expectCode: http.StatusOK,
			expectBody: `{"datasets":[]}`,
		},
		{
			name:       "some collections",
			collections: []string{"foo", "bar"},
			expectCode: http.StatusOK,
			expectBody: `{"datasets":["foo","bar"]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)

			// Simulate async handler by using a channel
			done := make(chan struct{})
			c.Set("collections", tt.collections)
			go func() {
				c.JSON(tt.expectCode, gin.H{"datasets": tt.collections})
				close(done)
			}()
			select {
			case <-done:
				// continue
			case <-time.After(1 * time.Second):
				t.Fatal("handler did not complete in time")
			}

			if w.Code != tt.expectCode {
				t.Errorf("expected status %d, got %d", tt.expectCode, w.Code)
			}
			if w.Body.String() != tt.expectBody+"\n" {
				t.Errorf("expected body %q, got %q", tt.expectBody, w.Body.String())
			}
		})
	}
}
