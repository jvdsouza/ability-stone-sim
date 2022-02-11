package main

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jvdsouza/ability-stone-sim/server-app/abilitystone"
)

func main() {
	abilityStoneAPI := abilitystone.NewApi()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Route("/abilitystone", func(r chi.Router) {
		r.Get("/new-stone", abilityStoneAPI.NewStone)
		r.Post("/roll-ability/{line}", abilitystone.RollAbility)
		r.Post("/current-probability", abilitystone.CurrentProbabilty)
		r.Post("/current-abilities", abilitystone.CurrentAbilities)
	})
	http.ListenAndServe(":3000", r)
}
