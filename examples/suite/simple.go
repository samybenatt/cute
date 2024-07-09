package suite

import (
	"context"
	"net/http"
	"net/url"
	"time"

	"github.com/ozontech/allure-go/pkg/framework/provider"
	"github.com/samybenatt/cute"
)

/*
	Example simple request
	Validate:
		1) Execute time
		2) Status code

*/

func (i *ExampleSuite) Test_Simple(t provider.T) {
	var (
		testMaker   = cute.NewHTTPTestMaker()
		testBuilder = testMaker.NewTestBuilder()
	)

	u, _ := url.Parse("https://jsonplaceholder.typicode.com/posts/1/comments")

	testBuilder.
		Title("Super simple test").
		Tags("simple", "suite", "some_local_tag", "json").
		Parallel().
		Create().
		RequestBuilder(
			cute.WithHeaders(map[string][]string{
				"some_header": []string{"something"},
			}),
			cute.WithURL(u),
			cute.WithMethod(http.MethodPost),
		).
		ExpectExecuteTimeout(10*time.Second).
		ExpectStatus(http.StatusCreated).
		ExecuteTest(context.Background(), t)
}
