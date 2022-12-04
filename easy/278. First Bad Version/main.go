package main

var bad = 0

func isBadVersion(version int) bool {
	if version == bad {
		return true
	}

	return false
}

func firstBadVersion(n int) int {
	for i := 1; i <= n; i++ {
		if isBadVersion(i) {
			return i
		}
	}

	return -1
}

func firstBadVersionBinary(n int) int {
	l, r := 1, n

	curBadVersion := -1
	for l <= r {
		mid := l + (r-l)/2
		if isBadVersion(mid) {
			r = mid - 1
			curBadVersion = mid
		} else {
			l = mid + 1
		}
	}

	return curBadVersion
}

func main() {

}
