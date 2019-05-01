package proxy

type subject interface {
	request() string
}

type realSubject struct {
}

func (rs *realSubject) request() string {
	return "realRequest"
}

type proxy struct {
	rs *realSubject
}

func newProxy() *proxy {
	return &proxy{
		rs: &realSubject{},
	}
}

func (p *proxy) request() string {
	return p.preRequest() + p.rs.request() + p.postRequest()
}

func (p *proxy) preRequest() string {
	return "preRequest "
}

func (p *proxy) postRequest() string {
	return " postRequest"
}
