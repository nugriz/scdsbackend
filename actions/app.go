package actions

import (
	"sync"

	"scdsbackend/locales"
	"scdsbackend/models"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo-pop/v3/pop/popmw"
	"github.com/gobuffalo/envy"
	"github.com/gobuffalo/middleware/contenttype"
	"github.com/gobuffalo/middleware/forcessl"
	"github.com/gobuffalo/middleware/i18n"
	"github.com/gobuffalo/middleware/paramlogger"
	tokenauth "github.com/gobuffalo/mw-tokenauth"
	"github.com/gobuffalo/x/sessions"
	"github.com/rs/cors"
	"github.com/unrolled/secure"
)

// ENV is used to help switch settings based on where the
// application is being run. Default is "development".
var ENV = envy.Get("GO_ENV", "development")

var (
	app     *buffalo.App
	appOnce sync.Once
	T       *i18n.Translator
)

// App is where all routes and middleware for buffalo
// should be defined. This is the nerve center of your
// application.
//
// Routing, middleware, groups, etc... are declared TOP -> DOWN.
// This means if you add a middleware to `app` *after* declaring a
// group, that group will NOT have that new middleware. The same
// is true of resource declarations as well.
//
// It also means that routes are checked in the order they are declared.
// `ServeFiles` is a CATCH-ALL route, so it should always be
// placed last in the route declarations, as it will prevent routes
// declared after it to never be called.
func App() *buffalo.App {
	appOnce.Do(func() {
		app = buffalo.New(buffalo.Options{
			Env:          ENV,
			SessionStore: sessions.Null{},
			PreWares: []buffalo.PreWare{
				cors.AllowAll().Handler,
			},
			SessionName: "_scdsbackend_session",
		})

		// Automatically redirect to SSL
		app.Use(forceSSL())

		// Log request parameters (filters apply).
		app.Use(paramlogger.ParameterLogger)

		// Set the request content type to JSON
		app.Use(contenttype.Set("application/json"))

		// Save AuthMiddleware function.
		AuthMiddleware := tokenauth.New(tokenauth.Options{})

		// Adding to my api the function.
		app.Use(AuthMiddleware)

		

		// Disable Auth Middleware in these fuctions
		app.Middleware.Skip(
			AuthMiddleware,
			AuthLogin,
			UsersCreate,
		)

		// Wraps each request in a transaction.
		//   c.Value("tx").(*pop.Connection)
		// Remove to disable this.
		app.Use(popmw.Transaction(models.DB))
		app.GET("/", HomeHandler)
		app.POST("/users", UsersCreate)
		app.POST("/users/auth", AuthLogin)
		app.GET("/users", UsersRead)
		app.GET("/users/{user_id}/", UsersShow)
		app.GET("/users/e/{email}/", UsersShowByEmail)
		app.GET("/users/{supplier_id}/products", ProductsIndexBySupplier)
		app.GET("/users/{buyer_id}/orders", OrdersIndex)
		app.GET("/warehouses/{warehouse_id}/", WarehousesShow)
		app.GET("/warehouses/{warehouse_id}/inventories", InventoriesIndex)
		app.GET("/warehouses", WarehousesIndex)
		app.POST("/warehouses", WarehousesCreate)
		// app.GET("/warehouses/delete", WarehousesDelete)
		app.GET("/products/{product_id}/", ProductsShow)
		app.POST("/products", ProductsCreate)
		app.GET("/products", ProductsIndex)
		// app.GET("/products/delete", ProductsDelete)
		app.POST("/inventories/", InventoriesCreate)
		app.POST("/orders", OrdersCreate)
		app.GET("/orders", OrdersList)

		app.POST("/chats", ChatsCreate)
		app.GET("/chats", ChatsIndex)
	})

	return app
}

// translations will load locale files, set up the translator `actions.T`,
// and will return a middleware to use to load the correct locale for each
// request.
// for more information: https://gobuffalo.io/en/docs/localization
func translations() buffalo.MiddlewareFunc {
	var err error
	if T, err = i18n.New(locales.FS(), "en-US"); err != nil {
		app.Stop(err)
	}
	return T.Middleware()
}

// forceSSL will return a middleware that will redirect an incoming request
// if it is not HTTPS. "http://example.com" => "https://example.com".
// This middleware does **not** enable SSL. for your application. To do that
// we recommend using a proxy: https://gobuffalo.io/en/docs/proxy
// for more information: https://github.com/unrolled/secure/
func forceSSL() buffalo.MiddlewareFunc {
	return forcessl.Middleware(secure.Options{
		SSLRedirect:     ENV == "production",
		SSLProxyHeaders: map[string]string{"X-Forwarded-Proto": "https"},
	})
}
