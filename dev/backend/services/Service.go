package services

type Service interface {
	// I don't know any smart method to implement here
	// I just need to have some method to have something
	// to implement in the other sevices
	// In another words
	// Each service will implement this method to
	// be considered a service
	New()
}
