package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/rs/cors"
)

type Server struct {
	donatugee *Donatugee
}

func NewServer(donatugee *Donatugee) *Server {
	s := &Server{
		donatugee: donatugee,
	}

	return s
}

func (s *Server) start() error {
	addr := "8081"
	if os.Getenv("ENV") == "production" {
		addr = os.Getenv("PORT")
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/challenges", s.challenges)
	mux.HandleFunc("/api/v1/insert-techfugee", s.insertTechfugee)
	mux.HandleFunc("/api/v1/insert-donator", s.insertDonator)
	mux.HandleFunc("/api/v1/techfugees", s.techfugees)
	mux.HandleFunc("/api/v1/techfugee", s.techfugee)
	mux.HandleFunc("/api/v1/login", s.loginTechfugee)
	mux.HandleFunc("/api/v1/login-donator", s.loginDonator)
	mux.HandleFunc("/api/v1/challenge", s.challenge)
	mux.HandleFunc("/api/v1/donator", s.donator)
	mux.HandleFunc("/api/v1/update-auth", s.updateAuth)
	mux.HandleFunc("/api/v1/add-skills", s.addSkills)
	mux.HandleFunc("/api/v1/insert-challenge", s.insertChallenge)
	mux.HandleFunc("/api/v1/update-techfugee", s.updateTechfugee)
	mux.HandleFunc("/api/v1/insert-application", s.insertApplication)
	mux.HandleFunc("/api/v1/accept-application", s.acceptApplication)
	mux.HandleFunc("/api/v1/application-by-techfugee", s.applicationByTechfugee)
	mux.HandleFunc("/api/v1/challenges-by-donator", s.challengesByDonator)

	mux.Handle("/public", http.FileServer(http.Dir("./frontend/public")))
	mux.Handle("/dist", http.FileServer(http.Dir("./frontend/dist")))
	mux.Handle("/", http.FileServer(http.Dir("./frontend")))

	handler := cors.Default().Handler(mux)

	return http.ListenAndServe(":"+addr, handler)
}

func IndexHandler(entrypoint string) func(w http.ResponseWriter, r *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, entrypoint)
	}

	return http.HandlerFunc(fn)
}

func (s *Server) applicationByTechfugee(resp http.ResponseWriter, r *http.Request) {
	idTechfugee := r.FormValue("id")

	applications, errs := s.donatugee.ChallengesByTechfugee(idTechfugee)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("query: %v", errs), http.StatusInternalServerError)
	}

	js, err := json.Marshal(applications)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) insertChallenge(resp http.ResponseWriter, r *http.Request) {
	idDonator := r.FormValue("id_donator")
	name := r.FormValue("name")
	description := r.FormValue("description")
	laptopType := r.FormValue("laptop_type")
	hardwareProvided := r.FormValue("hardware_provided")
	amount := r.FormValue("amount")
	duration := r.FormValue("duration")

	challenge, errs := s.donatugee.InsertChallenge(idDonator, name, description, laptopType, amount, hardwareProvided, duration)
	if len(errs) > 0 {
		http.Error(resp, fmt.Sprintf("insert: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(challenge)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) updateAuth(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	passed := r.FormValue("passed")

	techfugee, errs := s.donatugee.UpdateAuth(id, passed)
	if len(errs) > 0 {
		http.Error(resp, fmt.Sprintf("update: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", errs), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) updateTechfugee(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	city := r.FormValue("city")
	introduction := r.FormValue("introduction")

	techfugee, errs := s.donatugee.UpdateTechfugee(id, city, introduction)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("query: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) challengesByDonator(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	challenges, errs := s.donatugee.ChallengesByDonator(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("query: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(challenges)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) techfugees(resp http.ResponseWriter, r *http.Request) {
	techfugees, errs := s.donatugee.Techfugees()
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("query: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugees)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) insertTechfugee(resp http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	skills := r.FormValue("skills")

	techfugee, errs := s.donatugee.InsertTechfugee(name, email, skills)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("insert: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) insertDonator(resp http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	email := r.FormValue("email")
	website := r.FormValue("website")
	address := r.FormValue("address")

	donator, errs := s.donatugee.InsertDonator(name, email, website, address)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("insert: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(donator)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) insertApplication(resp http.ResponseWriter, r *http.Request) {
	techfugeeID := r.FormValue("techfugee_id")
	challengeID := r.FormValue("challenge_id")

	application, errs := s.donatugee.InsertApplication(techfugeeID, challengeID)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("insert: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(application)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) techfugee(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	techfugee, errs := s.donatugee.Techfugee(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("get: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) loginTechfugee(resp http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	techfugee, errs := s.donatugee.LoginTechfugee(email)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("get: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) loginDonator(resp http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")

	donator, errs := s.donatugee.LoginDonator(email)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("get: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(donator)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) donator(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	donator, errs := s.donatugee.Donator(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("get: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(donator)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) challenge(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	challenge, errs := s.donatugee.Challenge(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("get: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(challenge)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) acceptApplication(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	application, errs := s.donatugee.AcceptApplication(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("get: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(application)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}

func (s *Server) addSkills(resp http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	skills := r.FormValue("skills")

	techfugee, errs := s.donatugee.Techfugee(id)
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("getfugee: %v", errs), http.StatusInternalServerError)
		return
	}

	techfugee, errs = s.donatugee.UpdateTechfugeeSkills(techfugee, skills)

	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("addskille: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(techfugee)
	if err != nil {
		http.Error(resp, fmt.Sprintf("marshal: %v", err), http.StatusInternalServerError)
	}

	_, _ = resp.Write(js)
}

func (s *Server) challenges(resp http.ResponseWriter, r *http.Request) {
	challenges, errs := s.donatugee.Challenges()
	if len(errs) != 0 {
		http.Error(resp, fmt.Sprintf("challenges: %v", errs), http.StatusInternalServerError)
		return
	}

	js, err := json.Marshal(challenges)
	if err != nil {
		http.Error(resp, fmt.Sprintf("json: %v", err), http.StatusInternalServerError)
		return
	}

	_, _ = resp.Write(js)
}
