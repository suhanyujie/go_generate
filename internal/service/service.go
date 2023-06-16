package service

// GoProGen go 项目生成器
type GoProGen struct {
	srcPath, dstPath string
}

// NewGoProGen 实例化项目生成器
func NewGoProGen(srcPath, dstPath string) any {
	ret := &GoProGen{
		srcPath: srcPath,
		dstPath: dstPath,
	}
	return ret
}

func (g *GoProGen) Initializing() {
	//TODO implement me
	return
}

func (g *GoProGen) Prompting() {

	return
}

func (g *GoProGen) Configuring() {
	//TODO implement me
	return
}

func (g *GoProGen) Default() {
	//TODO implement me
	return
}

func (g *GoProGen) Writing() {
	//TODO implement me
	return
}

func (g *GoProGen) Conflicts() {
	//TODO implement me
	return
}

func (g *GoProGen) Install() {
	//TODO implement me
	return
}

func (g *GoProGen) End() {
	//TODO implement me
	return
}
