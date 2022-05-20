package controller

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/product", s.GetProducts).Methods("GET")
	s.Router.HandleFunc("/product/{id}", s.GetProduct).Methods("GET")
	s.Router.HandleFunc("/product", s.AddProducts).Methods("POST")
	s.Router.HandleFunc("/product/buy", s.BuyProduct).Methods("POST")
	s.Router.HandleFunc("/product/recommend/{id}", s.RecommendProducts).Methods("GET")
}
