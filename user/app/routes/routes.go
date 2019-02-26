package routes

import (
	"io/ioutil"

	"assignment/user/app/repository"
	"assignment/user/utils"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/docgen"
	"github.com/go-chi/jwtauth"
)

// All user routes for bicho app will be initialize in NewRouter
func NewRouter(ctrl repository.User) *chi.Mux {
	r := chi.NewRouter()
	// Basic CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	})
	r.Use(cors.Handler)

	r.Group(func(r chi.Router) {
		r.Post("/register", ctrl.CreateUserHandler)
		r.Post("/login", ctrl.GetUserHandler)
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(utils.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Put("/updateuserprofile", ctrl.UpdateUserHandler)
		r.Delete("/delete", ctrl.DeleteUserHandler)
	})

	genDocs(r)
	return r
}

// genDocs will generate documentation for user endpoints
func genDocs(rout *chi.Mux) {
	// Markdown docs
	ioutil.WriteFile("user/user_routes.md", []byte(docgen.MarkdownRoutesDoc(rout,
		docgen.MarkdownOpts{
			ProjectPath:        "assignment/user",
			Intro:              "User API.",
			ForceRelativeLinks: true,
		})), 0666)
}
