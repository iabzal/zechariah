package iskanderzhuma

type (
	options struct {
		debug bool
	}

	Option func(o *options)
)
