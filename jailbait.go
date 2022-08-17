package jailbait

import (
	"net/http"

	"github.com/caddyserver/caddy/v2"
	"github.com/caddyserver/caddy/v2/caddyconfig/caddyfile"
	"github.com/caddyserver/caddy/v2/caddyconfig/httpcaddyfile"
	"github.com/caddyserver/caddy/v2/modules/caddyhttp"
)

func init() {
	caddy.RegisterModule(Jailbait{})
	httpcaddyfile.RegisterHandlerDirective("jailbait", parseCaddyfile)
}

type Jailbait struct {}

func (Jailbait) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "http.handlers.jailbait",
		New: func() caddy.Module {
			return new(Jailbait)
		},
	}
}

func (s *Jailbait) UnmarshalCaddyfile(d *caddyfile.Dispenser) error {
	return nil
}

func parseCaddyfile(h httpcaddyfile.Helper) (caddyhttp.MiddlewareHandler, error) {
	t := new(Jailbait)

	err := t.UnmarshalCaddyfile(h.Dispenser)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (s Jailbait) ServeHTTP(w http.ResponseWriter, r *http.Request, _ caddyhttp.Handler) error {
	if blacklist(r) == true {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hi"))
	} else {
		http.Redirect(w, r, "https://youtu.be/2okC8Sd4_C8", http.StatusFound)
	}
	return nil
}

var (
	_ caddyhttp.MiddlewareHandler 	= (*Jailbait)(nil)
	_ caddyfile.Unmarshaler			= (*Jailbait)(nil)
)