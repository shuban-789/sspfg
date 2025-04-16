package main

const (
	def = 0
	des = 2
	pra = 4
	ice = 6
)

var (
	defaultWorldTemplate   [][][2]int
	desertWorldTemplate [][][2]int
	prairieWorldTemplate [][][2]int
	iceWorldTemplate [][][2]int
	layoutType = 1
)

// Layout --> (w: 40 chunks, h: 30 chunks)
// Reserve bottom 2 chunks for the base platform

func genWorldTemplateType1(biome int) [][][2]int {
	var template [][][2]int
	void := [2]int{6, 0}
	template = append(template, repeatRow(void, WidthTiles, 28)...)
	template = append(template, generateUniformRow([2]int{0, biome}, WidthTiles)) // Land Chunk 1
	template = append(template, generateUniformRow([2]int{0, biome}, WidthTiles)) // Land Chunk 2
	return template
}

func initWorldTemplates() {
	switch layoutType {
		case 1:
			defaultWorldTemplate = genWorldTemplateType1(def)
			desertWorldTemplate = genWorldTemplateType1(des)
			prairieWorldTemplate = genWorldTemplateType1(pra)
			iceWorldTemplate = genWorldTemplateType1(ice)
		default:
			defaultWorldTemplate = genWorldTemplateType1(def)
			desertWorldTemplate = genWorldTemplateType1(des)
			prairieWorldTemplate = genWorldTemplateType1(pra)
			iceWorldTemplate = genWorldTemplateType1(ice)
	}
}