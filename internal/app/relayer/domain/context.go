package domain

type Context struct {
	height    uint64
	chainName string
}

func NewContext(height uint64, chainName string) *Context {
	return &Context{height: height, chainName: chainName}
}

// Height return the current height of the chain
func (ctx *Context) Height() uint64 {
	return ctx.height
}

func (ctx *Context) IncrHeight() {
	ctx.height += 1
}

func (ctx *Context) SetHeight(height uint64) {
	ctx.height = height
}

// ChainName return the ChainName of the chain
func (ctx *Context) ChainName() string {
	return ctx.chainName
}
