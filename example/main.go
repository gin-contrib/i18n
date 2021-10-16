package main

import (
	"fmt"
	gini18n "gin-i18n"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	GetGinRouter()
}

func GetGinRouter() *gin.Engine {
	// new gin engine
	router := gin.New()

	// apply i18n middleware
	router.Use(gini18n.Localize())

	// Router Index
	index := router.Group("/")
	{
		index.GET("/", func(context *gin.Context) {
			resp := fmt.Sprintf("Tada")
			context.String(http.StatusOK, resp)
			gini18n.MustGetMessage("welcome")
		})
	}

	return router
}

// func handleGetAPI(r *gin.Engine, t *testing.T, isQuickBuy bool) string {
// 	bodyBytes, err := json.Marshal(body)
// 	if err != nil {
// 		t.Fatalf("Couldn't json marshal: %v\n", err)
// 	}
//
// 	urlStr := "/order-api/order/cart"
// 	if isQuickBuy {
// 		urlStr = "/order-api/order/cart/quick-buy"
// 	}
//
//
// 	// Create a response recorder so you can inspect the response
// 	w := httptest.NewRecorder()
//
// 	fmt.Println("Request: ")
// 	fmt.Println(string(bodyBytes))
//
// 	req, err := http.NewRequest("POST", urlStr, bytes.NewBuffer(bodyBytes))
// 	req.Header.Add("Content-Type", "application/json")
//
// 	if err != nil {
// 		t.Fatalf("Couldn't create request: %v\n", err)
// 	}
//
// 	// Perform the request
// 	r.ServeHTTP(w, req)
//
// 	// Check to see if the response was what you expected
// 	if w.Code != http.StatusOK {
// 		t.Fatalf("Expected to get status %d but instead got %d\n", http.StatusOK, w.Code)
// 	}
//
// 	fmt.Println("Response body: ")
// 	fmt.Println(w.Body.String())
// 	res := new(tranformed_obj.OrderResponseObject)
// 	err = json.Unmarshal(w.Body.Bytes(), &res)
// 	if err != nil {
// 		t.Fatalf("Couldn't parse response body: %v\n", err)
// 	}
//
// 	isOkResponse := res.Status == 1
// 	println("Response message: ", res.Message)
// 	if !isOkResponse {
// 		t.Fatalf("Add to cart failed: %s\n", res.Message)
// 	}
//
// 	cookie := w.Header().Get("Set-Cookie")
// 	return cookie
// }
