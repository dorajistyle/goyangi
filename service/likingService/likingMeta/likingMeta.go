package likingMeta

import "github.com/dorajistyle/goyangi/model"

// SetLikingPageMeta set likingList's page meta.
func SetLikingPageMeta(likingList *model.LikingList, currentPage int, hasPrev bool, hasNext bool, count int, currentUserlikedCount int) {
	if len(likingList.Likings) == 0 {
		likingList.Likings = make([]*model.PublicUser, 0)
	}
	if currentUserlikedCount == 1 {
		likingList.IsLiked = true
	}
	likingList.CurrentPage = currentPage
	likingList.HasPrev = hasPrev
	likingList.HasNext = hasNext
	likingList.Count = count
}

// SetLikedPageMeta set likedList's page meta.
func SetLikedPageMeta(likedList *model.LikedList, currentPage int, hasPrev bool, hasNext bool, count int) {
	if len(likedList.Liked) == 0 {
		likedList.Liked = make([]*model.PublicUser, 0)
	}
	likedList.CurrentPage = currentPage
	likedList.HasPrev = hasPrev
	likedList.HasNext = hasNext
	likedList.Count = count
}
