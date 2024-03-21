package algo

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	originColor := image[sr][sc]
	if originColor != color {
		changeColor(sr, sc, originColor, color, image)
	}

	return image
}

func changeColor(row, col, originColor, newColor int, img [][]int) {
	if !(row >= 0 && row < len(img) && col >= 0 && col < len(img[0]) && img[row][col] == originColor) {
		return
	}

	img[row][col] = newColor
	//Upward
	changeColor(row-1, col, originColor, newColor, img)
	// Downward
	changeColor(row+1, col, originColor, newColor, img)
	// Left
	changeColor(row, col-1, originColor, newColor, img)
	// Right
	changeColor(row, col+1, originColor, newColor, img)
}
