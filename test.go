package matrix

/*
 * A matrix is valid if it has rows, columns and if all its rows have the same
 * amount of columns.
 */
func ( matrix Matrix ) Valid() bool {
	return len( matrix ) > 2 && len( matrix ) == int(matrix[0]) * int(matrix[1]) + 2
}

/*
 * Tells if two matrices have the same dimensions and same
 * values in each cell.
 */
func ( matrix Matrix ) EqualTo( otherMatrix Matrix ) bool {
	for i := 0 ; i < len( matrix ) ; i++ {
		if matrix[i] != otherMatrix[i] {
			return false
		}
	}

	return true
}

/*
 * True if both matrices are valid and have the same dimensions
 */
func ( matrix Matrix ) SameDimensions( otherMatrix Matrix ) bool {
	if ! matrix.Valid() || ! otherMatrix.Valid() {
		return false
	}

	return matrix[0] == otherMatrix[0] && matrix[1] == otherMatrix[1]
}
