package pagination

func Paginate(currentPage int, perPage int, total int) (int, int, bool, bool) {
	var hasPrev, hasNext bool
	var offset int
	if currentPage == 0 {
		currentPage = 1
	}
	offset = (currentPage - 1) * perPage
	if currentPage > 1 {
		hasPrev = true
	}
	if total > (currentPage * perPage) {
		hasNext = true
	}
	return offset, currentPage, hasPrev, hasNext
}
