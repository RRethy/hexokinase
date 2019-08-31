package main

import (
	"strings"
)

var (
	// This need to be a list and not a map so it is ordered. This ordering
	// allows us to not match "blue" in "aliceblue" since "aliceblue" is a
	// colour that will already be matched.
	// https://www.w3schools.com/colors/colors_names.asp
	webColours = [][2]string{
		[2]string{"lightgoldenrodyellow", "#fafad2"},
		[2]string{"mediumspringgreen", "#00fa9a"},
		[2]string{"mediumaquamarine", "#66cdaa"},
		[2]string{"mediumslateblue", "#7b68ee"},
		[2]string{"mediumvioletred", "#c71585"},
		[2]string{"mediumturquoise", "#48d1cc"},
		[2]string{"lightslategray", "#778899"},
		[2]string{"lightslategrey", "#778899"},
		[2]string{"mediumseagreen", "#3cb371"},
		[2]string{"blanchedalmond", "#ffebcd"},
		[2]string{"lightsteelblue", "#b0c4de"},
		[2]string{"cornflowerblue", "#6495ed"},
		[2]string{"darkolivegreen", "#556b2f"},
		[2]string{"darkgoldenrod", "#b8860b"},
		[2]string{"palegoldenrod", "#eee8aa"},
		[2]string{"paleturquoise", "#afeeee"},
		[2]string{"lavenderblush", "#fff0f5"},
		[2]string{"rebeccapurple", "#663399"},
		[2]string{"darkslategray", "#2f4f4f"},
		[2]string{"darkslategrey", "#2f4f4f"},
		[2]string{"lightseagreen", "#20b2aa"},
		[2]string{"palevioletred", "#db7093"},
		[2]string{"darkturquoise", "#00ced1"},
		[2]string{"darkslateblue", "#483d8b"},
		[2]string{"antiquewhite", "#faebd7"},
		[2]string{"darkseagreen", "#8fbc8f"},
		[2]string{"lightskyblue", "#87cefa"},
		[2]string{"mediumorchid", "#ba55d3"},
		[2]string{"lemonchiffon", "#fffacd"},
		[2]string{"mediumpurple", "#9370db"},
		[2]string{"midnightblue", "#191970"},
		[2]string{"greenyellow", "#adff2f"},
		[2]string{"darkmagenta", "#8b008b"},
		[2]string{"lightsalmon", "#ffa07a"},
		[2]string{"lightyellow", "#ffffe0"},
		[2]string{"deepskyblue", "#00bfff"},
		[2]string{"navajowhite", "#ffdead"},
		[2]string{"saddlebrown", "#8b4513"},
		[2]string{"springgreen", "#00ff7f"},
		[2]string{"forestgreen", "#228b22"},
		[2]string{"floralwhite", "#fffaf0"},
		[2]string{"yellowgreen", "#9acd32"},
		[2]string{"papayawhip", "#ffefd5"},
		[2]string{"aquamarine", "#7fffd4"},
		[2]string{"dodgerblue", "#1e90ff"},
		[2]string{"chartreuse", "#7fff00"},
		[2]string{"blueviolet", "#8a2be2"},
		[2]string{"darkviolet", "#9400d3"},
		[2]string{"darkorange", "#ff8c00"},
		[2]string{"lightgreen", "#90ee90"},
		[2]string{"ghostwhite", "#f8f8ff"},
		[2]string{"whitesmoke", "#f5f5f5"},
		[2]string{"darkorchid", "#9932cc"},
		[2]string{"mediumblue", "#0000cd"},
		[2]string{"powderblue", "#b0e0e6"},
		[2]string{"lightcoral", "#f08080"},
		[2]string{"darksalmon", "#e9967a"},
		[2]string{"sandybrown", "#f4a460"},
		[2]string{"indianred", "#cd5c5c"},
		[2]string{"royalblue", "#4169e1"},
		[2]string{"steelblue", "#4682b4"},
		[2]string{"aliceblue", "#f0f8ff"},
		[2]string{"slategrey", "#708090"},
		[2]string{"mistyrose", "#ffe4e1"},
		[2]string{"turquoise", "#40e0d0"},
		[2]string{"lawngreen", "#7cfc00"},
		[2]string{"mintcream", "#f5fffa"},
		[2]string{"lightblue", "#add8e6"},
		[2]string{"slategray", "#708090"},
		[2]string{"lightcyan", "#e0ffff"},
		[2]string{"goldenrod", "#daa520"},
		[2]string{"lightgray", "#d3d3d3"},
		[2]string{"lightgrey", "#d3d3d3"},
		[2]string{"gainsboro", "#dcdcdc"},
		[2]string{"olivedrab", "#6b8e23"},
		[2]string{"chocolate", "#d2691e"},
		[2]string{"darkgreen", "#006400"},
		[2]string{"peachpuff", "#ffdab9"},
		[2]string{"rosybrown", "#bc8f8f"},
		[2]string{"burlywood", "#deb887"},
		[2]string{"firebrick", "#b22222"},
		[2]string{"slateblue", "#6a5acd"},
		[2]string{"lightpink", "#ffb6c1"},
		[2]string{"limegreen", "#32cd32"},
		[2]string{"orangered", "#ff4500"},
		[2]string{"cadetblue", "#5f9ea0"},
		[2]string{"darkkhaki", "#bdb76b"},
		[2]string{"palegreen", "#98fb98"},
		[2]string{"honeydew", "#f0fff0"},
		[2]string{"seashell", "#fff5ee"},
		[2]string{"seagreen", "#2e8b57"},
		[2]string{"deeppink", "#ff1493"},
		[2]string{"cornsilk", "#fff8dc"},
		[2]string{"darkblue", "#00008b"},
		[2]string{"darkcyan", "#008b8b"},
		[2]string{"darkgray", "#a9a9a9"},
		[2]string{"darkgrey", "#a9a9a9"},
		[2]string{"moccasin", "#ffe4b5"},
		[2]string{"lavender", "#e6e6fa"},
		[2]string{"darkred", "#8b0000"},
		[2]string{"hotpink", "#ff69b4"},
		[2]string{"skyblue", "#87ceeb"},
		[2]string{"oldlace", "#fdf5e6"},
		[2]string{"thistle", "#d8bfd8"},
		[2]string{"fuchsia", "#ff00ff"},
		[2]string{"magenta", "#ff00ff"},
		[2]string{"dimgrey", "#696969"},
		[2]string{"crimson", "#dc143c"},
		[2]string{"dimgray", "#696969"},
		[2]string{"tomato", "#ff6347"},
		[2]string{"bisque", "#ffe4c4"},
		[2]string{"silver", "#c0c0c0"},
		[2]string{"orchid", "#da70d6"},
		[2]string{"orange", "#ffa500"},
		[2]string{"yellow", "#ffff00"},
		[2]string{"sienna", "#a0522d"},
		[2]string{"maroon", "#800000"},
		[2]string{"salmon", "#fa8072"},
		[2]string{"purple", "#800080"},
		[2]string{"indigo", "#4b0082"},
		[2]string{"violet", "#ee82ee"},
		[2]string{"green", "#008000"},
		[2]string{"beige", "#f5f5dc"},
		[2]string{"azure", "#f0ffff"},
		[2]string{"olive", "#808000"},
		[2]string{"ivory", "#fffff0"},
		[2]string{"coral", "#ff7f50"},
		[2]string{"wheat", "#f5deb3"},
		[2]string{"white", "#ffffff"},
		[2]string{"linen", "#faf0e6"},
		[2]string{"brown", "#a52a2a"},
		[2]string{"khaki", "#f0e68c"},
		[2]string{"black", "#000000"},
		[2]string{"cyan", "#00ffff"},
		[2]string{"blue", "#0000ff"},
		[2]string{"aqua", "#00ffff"},
		[2]string{"navy", "#000080"},
		[2]string{"peru", "#cd853f"},
		[2]string{"teal", "#008080"},
		[2]string{"grey", "#808080"},
		[2]string{"snow", "#fffafa"},
		[2]string{"gray", "#808080"},
		[2]string{"gold", "#ffd700"},
		[2]string{"plum", "#dda0dd"},
		[2]string{"pink", "#ffc0cb"},
		[2]string{"lime", "#00ff00"},
		[2]string{"red", "#ff0000"},
		[2]string{"tan", "#d2b48c"},
	}
	webColoursDisabled = false
)

func parseWebColours(line string) colours {
	var clrs colours
	if webColoursDisabled {
		return clrs
	}

	used := make([]bool, len(line))
	line = strings.ToLower(line)
	for _, tuple := range webColours {
		curLine := line
		for len(curLine) > 0 {
			offset := len(line) - len(curLine)
			index := strings.Index(curLine, tuple[0])
			if index != -1 {
				if !used[offset+index] {
					if !checkBoundary || isWord(line, offset+index, offset+index+len(tuple[0])) {
						colour := &Colour{
							ColStart: offset + index + 1,
							ColEnd:   offset + index + len(tuple[0]),
							Hex:      tuple[1],
							Line:     line,
						}
						clrs = append(clrs, colour)
						for i := offset + index; i < offset+index+len(tuple[0]); i++ {
							used[i] = true
						}
					}
				}
				curLine = curLine[index+len(tuple[0]):]
			} else {
				break
			}
		}
	}
	return clrs
}
