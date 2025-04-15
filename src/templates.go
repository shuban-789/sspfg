package main

var WidthTiles = ScreenWidth / TileSize

const (
	blue   = 0
	orange = 1
	purple = 2
	cloudy = 3
)

var (
	blueSkyTemplate   [][][2]int
	orangeSkyTemplate [][][2]int
	purpleSkyTemplate [][][2]int
	cloudySkyTemplate [][][2]int
)

func genBackdrop(color int) [][][2]int {
	var template [][][2]int
	template = append(template, repeatRow([2]int{9, color}, WidthTiles, 3)...)
	template = append(template, generateUniformRow([2]int{10, color}, WidthTiles))
	template = append(template, repeatRow([2]int{11, color}, WidthTiles, 8)...)
	template = append(template, generateUniformRow([2]int{12, color}, WidthTiles))
	template = append(template, repeatRow([2]int{13, color}, WidthTiles, 8)...)
	template = append(template, generateUniformRow([2]int{14, color}, WidthTiles))
	template = append(template, repeatRow([2]int{15, color}, WidthTiles, 8)...)
	return template
}

func initTemplates() {
	blueSkyTemplate = genBackdrop(blue)
	orangeSkyTemplate = genBackdrop(orange)
	purpleSkyTemplate = genBackdrop(purple)
	cloudySkyTemplate = genBackdrop(cloudy)
}
