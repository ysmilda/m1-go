package res

type Attribute uint32

func (a Attribute) TypeOfModule() string {
	switch a & 0xf {
	case 1:
		return "Msys"
	case 2:
		return "VxWorks object"
	case 3:
		return "Library"
	case 4:
		return "PLC Program"
	case 8:
		return "Java Program"
	}
	return "Unknown"
}

func (a Attribute) IsReentrant() bool {
	return a&0x10 == 0x10
}

func (a Attribute) OnlineConfigurationAllowed() bool {
	return !(a&0x20 == 0x20)
}

func (a Attribute) OnlineInstallationAllowed() bool {
	return !(a&0x40 == 0x40)
}

func (a Attribute) OnlineDeinstallationAllowed() bool {
	return !(a&0x80 == 0x80)
}

func (a Attribute) LoadModuleOnBoot() bool {
	return a&0x100 == 0x100
}

func (a Attribute) RetainVariablesInUse() bool {
	return a&0x200 == 0x200
}

func (a Attribute) NoSetToZero() bool {
	return a&0x400 == 0x400
}

func (a Attribute) DebugCode() bool {
	return a&0x800 == 0x800
}

func (a Attribute) DebugInformationIncluded() bool {
	return a&0x1000 == 0x1000
}

func (a Attribute) NoInstallationAtBootTime() bool {
	return a&0x2000 == 0x2000
}

func (a Attribute) IsPartOfSystem() bool {
	return a&0x4000 == 0x4000
}

func (a Attribute) HasJavaClassLoader() bool {
	return a&0x8000 == 0x8000
}

func (a Attribute) IsServiceProgram() bool {
	return a&0x10000 == 0x10000
}

func (a Attribute) HadSecondCRUD() bool {
	return a&0x20000 == 0x20000
}

func (a Attribute) IsRedundant() bool {
	return a&0x40000 == 0x40000
}

func (a Attribute) IsComponent() bool {
	return a&0x80000 == 0x80000
}

func (a Attribute) IsMMPReady() bool {
	return a&0x100000 == 0x100000
}
