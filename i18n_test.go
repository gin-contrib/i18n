package i18n

import (
	"context"
	"fmt"
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
	router.GET("/exist/:lang", func(context *gin.Context) {
		context.String(http.StatusOK, "%v", HasLang(context, context.Param("lang")))
	})
	router.GET("/lang/default", func(context *gin.Context) {
		context.String(http.StatusOK, "%s", GetDefaultLanguage(context).String())
	})
	router.GET("/lang/current", func(context *gin.Context) {
		context.String(http.StatusOK, "%s", GetCurrentLanguage(context).String())
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
		// German
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
		// French
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
		// has exist
		{
			name: "i81n lang exist",
			args: args{
				path: fmt.Sprintf("/exist/%s", language.English.String()),
				lng:  language.English,
			},
			want: "true",
		},
		{
			name: "i81n lang not exist",
			args: args{
				path: fmt.Sprintf("/exist/%s", language.SimplifiedChinese.String()),
				lng:  language.English,
			},
			want: "false",
		},
		// default lang
		{
			name: "i81n is default " + language.English.String(),
			args: args{
				path: "/lang/default",
				lng:  language.English,
			},
			want: language.English.String(),
		},
		{
			name: "i81n is not default " + language.German.String(),
			args: args{
				path: "/lang/default",
				lng:  language.German,
			},
			want: language.English.String(),
		},
		// current lang
		{
			name: "i81n is current " + language.English.String(),
			args: args{
				path: "/lang/current",
				lng:  language.English,
			},
			want: language.English.String(),
		},
		{
			name: "i81n is not current " + language.English.String(),
			args: args{
				path: "/lang/current",
				lng:  language.German,
			},
			want: language.German.String(),
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
