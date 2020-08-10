package xcopy

type Option func(*XCopy)

func WithSource(source interface{}) Option {
	return func(c *XCopy) {
		c.source = source
	}
}

func WithConvert(convert bool) Option {
	return func(c *XCopy) {
		c.convert = convert
	}
}

func WithNext(next bool) Option {
	return func(c *XCopy) {
		c.next = next
	}
}
func WithRecursion(recursion bool) Option {
	return func(c *XCopy) {
		c.recursion = recursion
	}
}

func WithJsonTag(jsonTag bool) Option {
	return func(c *XCopy) {
		c.jsonTag = jsonTag
	}
}

func WithXCM(xcm XConverters) Option {
	return func(c *XCopy) {
		c.xcm = xcm
	}
}
