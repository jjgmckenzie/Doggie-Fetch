package main

import "gofetch/postedimage"

func makeGitBranch(name postedimage.Image) {

}

func addImageToGitBranch(image postedimage.Image) {

}

func makeGitBranchWithImage(image postedimage.Image) {
	makeGitBranch(image)
	addImageToGitBranch(image)
}

func pushGitBranchToGitHub(image postedimage.Image) string {
	return ""
}

func postImageToGitHub(image postedimage.Image) string {
	makeGitBranchWithImage(image)
	return pushGitBranchToGitHub(image)
}
