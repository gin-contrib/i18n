package i18n

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

// newServer ...
func newServer() *gin.Engine {
	router := gin.New()
	router.Use(Localize())

	router.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(context, "welcome"))
	})

	router.GET("/messageId/:name", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(context, &i18n.LocalizeConfig{
			MessageID: msgIDWelcomeName,
			TemplateData: map[string]string{
				keyName: context.Param(keyName),
			},
		}))
	})

	router.GET("/messageType/:name", func(context *gin.Context) {
		context.String(http.StatusOK, MustGetMessage(context, &i18n.LocalizeConfig{
			DefaultMessage: &i18n.Message{
				ID: msgIDWelcomeName,
			},
			TemplateData: map[string]string{
				keyName: context.Param(keyName),
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
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, path, nil)
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
			want: wantHello,
		},
		{
			name: "hello alex - messageId",
			args: args{
				path: pathMessageIDAlex,
				lng:  language.English,
			},
			want: wantHelloAlex,
		},
		{
			name: "hello alex - messageType",
			args: args{
				path: pathMessageTypAlex,
				lng:  language.English,
			},
			want: wantHelloAlex,
		},
		{
			name: "18 years old",
			args: args{
				path: pathAge18,
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
			want: wantHallo,
		},
		{
			name: "hallo alex - messageId",
			args: args{
				path: pathMessageIDAlex,
				lng:  language.German,
			},
			want: wantHalloAlex,
		},
		{
			name: "hallo alex - messageType",
			args: args{
				path: pathMessageTypAlex,
				lng:  language.German,
			},
			want: wantHalloAlex,
		},
		{
			name: "18 jahre alt",
			args: args{
				path: pathAge18,
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
			want: wantBonjour,
		},
		{
			name: "bonjour alex - messageId",
			args: args{
				path: pathMessageIDAlex,
				lng:  language.French,
			},
			want: wantBonjourAlex,
		},
		{
			name: "bonjour alex - messageType",
			args: args{
				path: pathMessageTypAlex,
				lng:  language.French,
			},
			want: wantBonjourAlex,
		},
		{
			name: "18 ans",
			args: args{
				path: pathAge18,
				lng:  language.French,
			},
			want: "j'ai 18 ans",
		},
		// has exist
		{
			name: "i81n lang exist",
			args: args{
				path: "/exist/" + language.English.String(),
				lng:  language.English,
			},
			want: "true",
		},
		{
			name: "i81n lang not exist",
			args: args{
				path: "/exist/" + language.SimplifiedChinese.String(),
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

// newFallbackServer creates a server with FallbackLanguages configured.
// Fallback chain: fr -> en (French first, then English).
// Test data:
//   - en.yaml: welcome, welcomeWithName, welcomeWithAge, englishOnly
//   - de.yaml: welcome, welcomeWithName
//   - fr.yaml: welcome, welcomeWithName, welcomeWithAge
//   - zh.yaml: welcome
func newFallbackServer() *gin.Engine {
	router := gin.New()
	router.Use(Localize(WithBundle(&BundleCfg{
		RootPath: "./testdata/localizeFallback",
		AcceptLanguage: []language.Tag{
			language.English,
			language.German,
			language.French,
			language.Chinese,
		},
		FallbackLanguages: []language.Tag{language.French, language.English},
		DefaultLanguage:   language.English,
		FormatBundleFile:  "yaml",
		UnmarshalFunc:     yaml.Unmarshal,
	})))

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, MustGetMessage(ctx, "welcome"))
	})

	router.GET("/messageId/:name", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, MustGetMessage(ctx, &i18n.LocalizeConfig{
			MessageID: "welcomeWithName",
			TemplateData: map[string]string{
				"name": ctx.Param("name"),
			},
		}))
	})

	router.GET("/age/:age", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, MustGetMessage(ctx, i18n.LocalizeConfig{
			MessageID: "welcomeWithAge",
			TemplateData: map[string]string{
				"age": ctx.Param("age"),
			},
		}))
	})

	router.GET(pathEnglishOnly, func(ctx *gin.Context) {
		ctx.String(http.StatusOK, MustGetMessage(ctx, "englishOnly"))
	})

	return router
}

func makeFallbackRequest(lng language.Tag, path string) string {
	req, _ := http.NewRequestWithContext(context.Background(), http.MethodGet, path, nil)
	req.Header.Add("Accept-Language", lng.String())
	w := httptest.NewRecorder()
	r := newFallbackServer()
	r.ServeHTTP(w, req)
	return w.Body.String()
}

func TestFallbackLocale(t *testing.T) {
	type args struct {
		lng  language.Tag
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// Key exists in target language -> returns target (no fallback needed)
		{
			name: "de has welcome",
			args: args{lng: language.German, path: "/"},
			want: wantHallo,
		},
		{
			name: "fr has welcome",
			args: args{lng: language.French, path: "/"},
			want: wantBonjour,
		},
		{
			name: "zh has welcome",
			args: args{lng: language.Chinese, path: "/"},
			want: "你好",
		},
		{
			name: "en has welcome",
			args: args{lng: language.English, path: "/"},
			want: wantHello,
		},
		// Key missing in target -> fallback to French (first in chain)
		{
			name: "de missing welcomeWithAge falls back to fr",
			args: args{lng: language.German, path: "/age/18"},
			want: "j'ai 18 ans",
		},
		// Key missing in target and French -> fallback to English (second in chain)
		{
			name: "de missing englishOnly falls back to en",
			args: args{lng: language.German, path: pathEnglishOnly},
			want: wantEnglishOnly,
		},
		{
			name: "fr missing englishOnly falls back to en",
			args: args{lng: language.French, path: pathEnglishOnly},
			want: wantEnglishOnly,
		},
		// Multi-step: zh missing welcomeWithAge -> not in zh, try fr (has it) -> returns French
		{
			name: "zh missing welcomeWithAge falls back to fr",
			args: args{lng: language.Chinese, path: "/age/25"},
			want: "j'ai 25 ans",
		},
		// Multi-step: zh missing englishOnly -> not in zh, not in fr, try en -> returns English
		{
			name: "zh missing englishOnly falls back to en",
			args: args{lng: language.Chinese, path: pathEnglishOnly},
			want: wantEnglishOnly,
		},
		// Key exists in target with template data -> no fallback
		{
			name: "de has welcomeWithName",
			args: args{lng: language.German, path: "/messageId/alex"},
			want: wantHalloAlex,
		},
		{
			name: "zh missing welcomeWithName falls back to fr",
			args: args{lng: language.Chinese, path: "/messageId/alex"},
			want: wantBonjourAlex,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeFallbackRequest(tt.args.lng, tt.args.path); got != tt.want {
				t.Errorf("makeFallbackRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLanguagePreferences(t *testing.T) {
	router := newServer()
	router.GET("/language-supported", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "%v", HasLang(ctx, ctx.GetHeader("Accept-Language")))
	})

	tests := []struct {
		name       string
		header     string
		want       string
		registered string
	}{
		{name: "exact language", header: "de", want: wantHallo, registered: "true"},
		{name: "regional language", header: "de-DE", want: wantHallo, registered: "false"},
		{name: "case insensitive language", header: "DE", want: wantHallo, registered: "false"},
		{name: "preference list", header: "fr,en;q=0.9", want: wantBonjour, registered: "false"},
		{
			name:       "quality order",
			header:     "de;q=0.2,fr;q=0.9",
			want:       wantBonjour,
			registered: "false",
		},
		{
			name:       "browser preferences",
			header:     "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7",
			want:       wantHallo,
			registered: "false",
		},
		{
			name:       "unsupported first preference",
			header:     "ja,fr;q=0.9",
			want:       wantBonjour,
			registered: "false",
		},
		{name: "zero quality", header: "de;q=0,en;q=0.5", want: wantHello, registered: "false"},
		{name: "empty header", header: "", want: wantHello, registered: "false"},
		{name: "unsupported language", header: "ja", want: wantHello, registered: "false"},
		{name: "invalid language", header: "not_a_language", want: wantHello, registered: "false"},
		{name: "invalid quality", header: "de;q=invalid", want: wantHello, registered: "false"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			for _, check := range []struct {
				path string
				want string
			}{
				{path: "/", want: tt.want},
				{path: "/language-supported", want: tt.registered},
			} {
				req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, check.path, nil)
				req.Header.Set("Accept-Language", tt.header)
				response := httptest.NewRecorder()
				router.ServeHTTP(response, req)
				if response.Code != http.StatusOK {
					t.Fatalf("%s returned status %d", check.path, response.Code)
				}
				if got := response.Body.String(); got != check.want {
					t.Errorf(
						"%s with Accept-Language %q = %q, want %q",
						check.path,
						tt.header,
						got,
						check.want,
					)
				}
			}
		})
	}
}

func TestLanguagePreferencesFallback(t *testing.T) {
	router := newFallbackServer()
	tests := []struct {
		name   string
		header string
		path   string
		want   string
	}{
		{name: "regional Chinese", header: "zh-CN", path: "/", want: "你好"},
		{name: "first fallback", header: "de-DE,de;q=0.9", path: "/age/18", want: "j'ai 18 ans"},
		{
			name:   "final fallback",
			header: "de-DE,de;q=0.9",
			path:   pathEnglishOnly,
			want:   wantEnglishOnly,
		},
		{
			name:   "requested language before fallback",
			header: "en-US,en;q=0.9",
			path:   "/age/18",
			want:   "I am 18 years old",
		},
		{name: "empty header uses default", header: "", path: "/", want: wantHello},
		{name: "unsupported language uses default", header: "es", path: "/", want: wantHello},
		{
			name:   "invalid language uses default",
			header: "not_a_language",
			path:   "/",
			want:   wantHello,
		},
		{name: "invalid quality uses default", header: "de;q=invalid", path: "/", want: wantHello},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			req := httptest.NewRequestWithContext(t.Context(), http.MethodGet, tt.path, nil)
			req.Header.Set("Accept-Language", tt.header)
			response := httptest.NewRecorder()
			router.ServeHTTP(response, req)
			if response.Code != http.StatusOK {
				t.Fatalf("%s returned status %d", tt.path, response.Code)
			}
			if got := response.Body.String(); got != tt.want {
				t.Errorf(
					"%s with Accept-Language %q = %q, want %q",
					tt.path,
					tt.header,
					got,
					tt.want,
				)
			}
		})
	}
}
