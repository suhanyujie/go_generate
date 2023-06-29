package service

// GoProGen go 项目生成器
type GoProGen struct {
	srcPath, dstPath string
	promptList       []*Prompt

	projectName string
}

// NewGoProGen 实例化项目生成器
func NewGoProGen(srcPath, dstPath string) any {
	ret := &GoProGen{
		srcPath: srcPath,
		dstPath: dstPath,
	}
	promptList := PromptBuiltInConfigList
	ret.promptList = promptList

	return ret
}

func (g *GoProGen) Initializing() {
	//TODO implement me
	return
}

func (g *GoProGen) Prompting() {
	for _, prompt := range g.promptList {
		prompt.Run()
		switch prompt.Name {
		case "name":
			g.projectName = prompt.Value
		}
	}
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
