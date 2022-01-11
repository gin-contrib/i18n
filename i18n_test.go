package i18n

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func init() {
	gin.SetMode(gin.ReleaseMode)
}

func newPartCustomerConfigServer() *gin.Engine {
	router := gin.New()
	router.Use(Localize(WithBundle(&BundleCfg{
		RootPath:         "./_example/i18n",
		FormatBundleFile: "json",
	})))

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage("welcome"))
	})

	router.GET("/:name", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
		}))
	})

	return router
}
func newCustomerServer() *gin.Engine {
	router := gin.New()
	router.Use(Localize(WithBundle(&BundleCfg{
		RootPath:         "./_example/i18n",
		AcceptLanguage:   []language.Tag{language.Chinese, language.English},
		DefaultLanguage:  language.English,
		FormatBundleFile: "json",
		UnmarshalFunc:    json.Unmarshal,
	})))

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage("welcome"))
	})

	router.GET("/:name/:food", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithNameAndFood",
			TemplateData: map[string]string{
				"name": context.Param("name"),
				"food": context.Param("food"),
			},
		}))
	})

	return router
}

// newServer ...
func newServer() *gin.Engine {
	router := gin.New()
	router.Use(Localize())

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage("welcome"))
	})

	router.GET("/:name", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(&i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": context.Param("name"),
			},
		}))
	})

	return router
}

func makeRequest3(
	lng language.Tag,
	name string,
) string {
	path := "/" + name
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Add("Accept-Language", lng.String())

	// Perform the request
	w := httptest.NewRecorder()
	r := newPartCustomerConfigServer()
	r.ServeHTTP(w, req)

	return w.Body.String()
}

func makeRequest2(
	lng language.Tag,
	name string,
	food string,
) string {
	path := "/" + name + "/" + food
	req, _ := http.NewRequest("GET", path, nil)
	req.Header.Add("Accept-Language", lng.String())

	// Perform the request
	w := httptest.NewRecorder()
	r := newCustomerServer()
	r.ServeHTTP(w, req)

	return w.Body.String()
}

// makeRequest ...
func makeRequest(
	lng language.Tag,
	name string,
) string {
	path := "/" + name
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
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello world",
			args: args{
				name: "",
				lng:  language.English,
			},
			want: "hello",
		},
		{
			name: "hello alex",
			args: args{
				name: "alex",
				lng:  language.English,
			},
			want: "hello alex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRequest(tt.args.lng, tt.args.name); got != tt.want {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nDE(t *testing.T) {
	type args struct {
		lng  language.Tag
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hallo",
			args: args{
				name: "",
				lng:  language.German,
			},
			want: "hallo",
		},
		{
			name: "hallo alex",
			args: args{
				name: "alex",
				lng:  language.German,
			},
			want: "hallo alex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRequest(tt.args.lng, tt.args.name); got != tt.want {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestI18nFR(t *testing.T) {
	type args struct {
		lng  language.Tag
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "bonjour",
			args: args{
				name: "",
				lng:  language.French,
			},
			want: "bonjour",
		},
		{
			name: "bonjour alex",
			args: args{
				name: "alex",
				lng:  language.French,
			},
			want: "bonjour alex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRequest(tt.args.lng, tt.args.name); got != tt.want {
				t.Errorf("makeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomerI18n(t *testing.T) {
	type args struct {
		lng  language.Tag
		name string
		food string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1st",
			args: args{
				name: "watch",
				food: "apple",
				lng:  language.Chinese,
			},
			want: "欢迎宝贝watch, 我喜欢吃 apple",
		},
		{
			name: "2nd",
			args: args{
				name: "issues",
				food: "orange",
				lng:  language.English,
			},
			want: "welcome my baby issues，I like orange",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRequest2(tt.args.lng, tt.args.name, tt.args.food); got != tt.want {
				t.Errorf("makeRequest2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPartCustomerConfigI18n(t *testing.T) {
	type args struct {
		lng  language.Tag
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hello world",
			args: args{
				name: "",
				lng:  language.English,
			},
			want: "hello",
		},
		{
			name: "hello alex",
			args: args{
				name: "alex",
				lng:  language.English,
			},
			want: "hello alex",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeRequest3(tt.args.lng, tt.args.name); got != tt.want {
				t.Errorf("makeRequest3() = %v, want %v", got, tt.want)
			}
		})
	}
}
