package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"inventaris/database"
	"inventaris/handler"
	"inventaris/repository"
	"inventaris/service"
	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	db := database.NewPostgresDB()

	userRepo := repository.NewUserRepo(db)
	sessionRepo := repository.NewSessionRepository(db)
	srv := service.NewUserService(userRepo, sessionRepo)

	categoryRepo := repository.NewCategoryRepository(db)
	srvCtgr := service.NewCategoryService(categoryRepo)

	itemRepo := repository.NewItemRepository(db)
	srvItm := service.NewItemService(itemRepo)

	h := handler.NewAuthHandler(srv)
	c := handler.NewCategoryHandler(srvCtgr)
	i := handler.NewItemHandler(srvItm)

	r.Use(middleware.Logger)

	// Serve static files
	// fileServer := http.FileServer(http.Dir("./static"))
	// r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Group(func(r chi.Router) {
		r.Post("/login", h.Login)
		r.Post("/register", h.Register)
		r.Post("/logout", h.Logout)
	})

	r.Group(func(r chi.Router) {
		r.Route("/api/categories", func(r chi.Router) {
			// r.Use(middleware_auth.SessionMiddleware(&srv))
			r.Get("/", c.GetCategories)
			r.Get("/{id}", c.GetCategoryByID)
			r.Put("/{id}", c.UpdateCategory)
			r.Delete("/{id}", c.DeleteCategory)
			r.Post("/add", c.CreateCategory)
		})
		r.Route("/api/items", func(r chi.Router) {
			// r.Use(middleware_auth.SessionMiddleware(&srv))
			r.Get("/", i.GetItems)
			r.Get("/{id}", i.GetItemByID)
			r.Put("/{id}", i.UpdateItem)
			r.Delete("/{id}", i.DeleteItem)
			r.Post("/add", i.CreateItem)
			r.Get("/replacement-needed", i.GetItemsReplacementNeeded)

		})
		r.Route("/api/items/investment", func(r chi.Router) {
			// r.Use(middleware_auth.SessionMiddleware(&srv))
			r.Get("/", i.GetTotalInvestment)
			r.Get("/{id}", i.GetItemInvestmentByID)
		})

	})

	return r
}
