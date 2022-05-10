package zsearch

type ZError struct {
	Err string
}

func Fail(s string) *ZError {
	return &ZError{
		Err: s,
	}
}

func (z *ZError) Error() string {
	return z.Err
}
