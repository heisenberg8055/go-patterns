package proxy

type Ngix struct {
	application       *Application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func NewNgixServer() *Ngix {
	return &Ngix{application: &Application{}, maxAllowedRequest: 3, rateLimiter: make(map[string]int)}
}

func (n *Ngix) HandleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}

func (n *Ngix) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] >= n.maxAllowedRequest {
		return false
	}
	n.rateLimiter[url] += 1
	return true
}
