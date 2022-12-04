package broker

type options struct {
	bridge Bridge
	auth   Auth
}

type Option func(*options)

func WithBridge(b Bridge) Option {
	return func(o *options) {
		o.bridge = b
	}
}

func WithAuth(a Auth) Option {
	return func(o *options) {
		o.auth = a
	}
}
