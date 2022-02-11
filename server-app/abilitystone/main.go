package abilitystone

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type API struct {
	NewStone           func(http.ResponseWriter, *http.Request)
	RollAbility        func(http.ResponseWriter, *http.Request)
	CurrentProbability func(http.ResponseWriter, *http.Request)
	CurrentAbilities   func(http.ResponseWriter, *http.Request)
}

func NewApi() API {
	return API{
		NewStone,
		RollAbility,
		CurrentProbabilty,
		CurrentAbilities,
	}
}

func NewStone(w http.ResponseWriter, r *http.Request) {
	abilityStone := NewAbilityStone()
	json.NewEncoder(w).Encode(abilityStone)
}

func RollAbility(w http.ResponseWriter, r *http.Request) {
	lineNumberJSON := chi.URLParam(r, "line")
	lineNumber, err := strconv.Atoi(lineNumberJSON)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	line := Line(lineNumber)

	var abilityStone AbilityStone
	err = json.NewDecoder(r.Body).Decode(&abilityStone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	abilityStone.RollAbilityLine(line)
	json.NewEncoder(w).Encode(abilityStone)
}

func CurrentProbabilty(w http.ResponseWriter, r *http.Request) {
	var abilityStone AbilityStone
	err := json.NewDecoder(r.Body).Decode(&abilityStone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	probability := strconv.Itoa(abilityStone.CurrentProbability)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(probability))
}

func CurrentAbilities(w http.ResponseWriter, r *http.Request) {
	var abilityStone AbilityStone
	err := json.NewDecoder(r.Body).Decode(&abilityStone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(abilityStone.Abilities)
}
