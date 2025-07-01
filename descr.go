package bolt

var Description map[string]string = map[string]string{
	"NoView":   "do not view all results",
	"FullView": "show all results",
	"view":     "local variable of view",

	"thk": "thickness of flange. Unit - meter",

	"dia":          "the float number interpretation of diameter",
	"holeDiameter": "diameter of hole. Unit - meter",

	"b":         "bolt property",
	"bc":        "bolt class",
	"bd":        "bolt diameter. Unit - meter",
	"boltPinch": "bolt pinch. Unit - meter",

	"fub": "the ultimate tensile strength. Unit - Pa",
	"fyb": "the yield strength. Unit - Pa",

	"FactorγM2": "the partial safety factor",

	"FtEd": "the design tensile force per bolt for the ultimate limit state. Unit - Pa",
	"FtRd": "the design tension resistance per bolt. Unit - Pa",

	"FvEd": "the design shear force per bolt for the ultimate limit state. Unit - Pa",
	"FvRd": "the shear design resistance per bolt. Unit - Pa",

	"ανThreadShear":   "factor if shear by thread of bolt",
	"ανUnthreadShear": "factor if shear not by thread of bolt",

	"max": "local variable of maximal value",
	"f1":  "local value of ratio",
	"f2":  "local value of ratio",

	"UsuallyBolt":     "type of bolt head",
	"CountersunkBolt": "type of bolt head",

	"ThreadShear":   "location of shear area on thread",
	"UnthreadShear": "location of shear area not on thread",

	// ignore
	"G4p6": "", "G4p8": "", "G5p6": "",
	"G5p8": "", "G6p8": "", "G8p8": "", "G10p9": "",

	"D12": "", "D16": "", "D20": "", "D24": "",
	"D30": "", "D36": "", "D42": "", "D48": "",

	"d": "", "p": "", "pd": "", "s": "",

	"class": "", "fubData": "", "fybData": "", "αν": "",
}
