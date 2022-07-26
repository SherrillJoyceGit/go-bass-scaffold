package bootstrap

import "fmt"

const (
	cBlack   = "\u001b[90m"
	cRed     = "\u001b[91m"
	cCyan    = "\u001b[96m"
	cGreen   = "\u001b[92m"
	cYellow  = "\u001b[93m"
	cBlue    = "\u001b[94m"
	cMagenta = "\u001b[95m"
	cWhite   = "\u001b[97m"
	cReset   = "\u001b[0m"
)

// renderBassIcon 渲染 Bass Icon
func renderBassIcon() {
	fmt.Print(`
-----------------------------------------------
 ` + cBlue + `o` + cReset + `
` + cBlue + `o` + cReset + `      ` + cGreen + `______/~/~/~/__         ` + cRed + `  /((
  ` + cBlue + `o` + cReset + `  ` + cRed + `// __            ====__    /_((
 ` + cBlue + `o` + cReset + `  ` + cRed + `//  @))       ))))      ===/__((
   ` + cRed + ` ))           )))))))        __((
    ` + cRed + `\\     \)     ))))    __===\ _((
    ` + cRed + ` \\_______________====      \_((
     ` + cRed + `                            \((` + cReset + `
-----------------------------------------------


   V0.0.1  URL http://127.0.0.1:8080` + `
================================================

`)
}

// ShowBass 显示 Bass Icon
func ShowBass() {
	renderBassIcon()
}
