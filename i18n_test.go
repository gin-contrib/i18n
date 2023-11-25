package i18n

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// newServer ...
func newServer() *gin.Engine {
	router := gin.New()
	router.Use(Localize())

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(context, "welcome"))
	})

	router.GET("/:name", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(context, &i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
		}))
	})
	router.GET("/age/:age", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(context, i18n.LocalizeConfig{
			MessageID: "welcomeWithAge",
			TemplateData: map[string]string{
				"age": context.Param("age"),
			},
		}))
	})

	return router
}

// makeRequest ...
func makeRequest(
	lng language.Tag,
	path string,
) string {
	req, _ := http.NewRequestWithContext(context.Background(), "GET", path, nil)
	req.Header.Add("Accept-Language", lng.String())

	// Perform the request
	w := httptest.NewRecorder()
	r := newServer()
	r.ServeHTTP(w, req)

	return w.Body.String()
}

func TestI18nEN(t *testing.T) {
	type args struct {
		lng  language.Tag
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello world",
			args: args{
				path: "/",
				lng:  language.English,
			},
			want: "hello",
		},
		{
			name: "hello alex",
			args: args{
				path: "/alex",
				lng:  language.English,
			},
			want: "hello alex",
		},
		{
			name: "18 years old",
			args: args{
				path: "/age/18",
				lng:  language.English,
			},
			want: "I am 18 years old",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRequest(tt.args.lng, tt.args.path); got != tt.want {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nDE(t *testing.T) {
	type args struct {
		lng  language.Tag
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hallo",
			args: args{
				path: "/",
				lng:  language.German,
			},
			want: "hallo",
		},
		{
			name: "hallo alex",
			args: args{
				path: "/alex",
				lng:  language.German,
			},
			want: "hallo alex",
		},
		{
			name: "18 jahre alt",
			args: args{
				path: "/age/18",
				lng:  language.German,
			},
			want: "ich bin 18 Jahre alt",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRequest(tt.args.lng, tt.args.path); got != tt.want {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nFR(t *testing.T) {
	type args struct {
		lng  language.Tag
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "bonjour",
			args: args{
				path: "/",
				lng:  language.French,
			},
			want: "bonjour",
		},
		{
			name: "bonjour alex",
			args: args{
				path: "/alex",
				lng:  language.French,
			},
			want: "bonjour alex",
		},
		{
			name: "18 ans",
			args: args{
				path: "/age/18",
				lng:  language.French,
			},
			want: "j'ai 18 ans",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRequest(tt.args.lng, tt.args.path); got != tt.want {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
