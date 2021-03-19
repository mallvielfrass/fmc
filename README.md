

# FMC
##### Library  for coloring text in the console.
### powered by https://github.com/fatih/color
### Features

- specify tags directly in the formatting string;
- Support prinft/print/println;
## Color:
-	yst    Yellow
-	ybt    Yellow, Bold

-	rst   Red
-	rbt    Red, Bold

-	gst    Green
-	gbt    Green, Bold

-	bst    Blue
-	bbt    Blue, Bold

-	wst    White
-	wbt    White, Bold
## Color Description:
### in hashtag:
- first char it is Color name(Red,Blue..)
- second char it is font weight slim\bold
- first char it is background color transparent (another colors not implemented yet)
## Examples:
```go
package main
import "github.com/mallvielfrass/fmc"
func main() {
colors := [10]string{"yst", "ybt", "rst", "rbt", "gst", "gbt", "bst", "bbt", "wst", "wbt"}
for _, value := range colors {
	fmc.Printf("color: #"+value+value+" %s\n", "test")
}

fmc.Printf("#ybt I#rbt  love#bbt  Go#gbt!\n")
fmc.Printf("#ybt %s#rbt  %s#bbt  %s#gbt%s\n", "I", "love", "Go", "!")
fmc.Print("#ybt lol")
fmc.Printfln("#ybt  %s", "I")
fmc.Println("#bbt Println")
}
```
# Result

[![Result](https://raw.githubusercontent.com/mallvielfrass/coloredPrint/main/fmc/scrot_2021-01-22-10_1920x1080.png "Result")](https://raw.githubusercontent.com/mallvielfrass/coloredPrint/main/fmc/scrot_2021-01-22-10_1920x1080.png "Result");


