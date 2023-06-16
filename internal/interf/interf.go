package interf

type ToolIf interface {
	Initializing()
	Prompting()
	Configuring()
	Default()
	Writing()
	Conflicts()
	Install()
	End()
}
