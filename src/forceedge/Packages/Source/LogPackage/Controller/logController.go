/*
Package logService would ideally deal with displaying, storing logs

The methods provided by this package are:
Print()
*/
package logPackage

import "forceedge/Packages/Source/LogPackage/Struct"

// Print will simply print text set to the display buffer
func Print(text string) {
	l := logPackage.NewLog()
	l.SetText(text).PrintLog()
}
