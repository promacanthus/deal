package stringMatch

import (
	"crypto/md5"
	"math"
)

func BruteForce(s, p string) bool {
	mBytes := []byte(s)
	pBytes := []byte(p)

	ml := len(mBytes)
	pl := len(pBytes)

	for i, j := 0, 0; i < ml; i++ {
		if mBytes[i] == pBytes[j] {
			j++
			if j == pl {
				return true
			}
		} else {
			j = 0
		}
	}
	return false
}

func RabinKarp(s, p string) bool {
	mBytes := []byte(s)
	pBytes := []byte(p)

	ml := len(mBytes)
	pl := len(pBytes)

	pHash := md5.Sum(pBytes)

	mHashes := make([][16]byte, ml-pl+1)

	for i := 0; i < ml-pl+1; i++ {
		tmp := mBytes[i : i+pl]
		mHashes[i] = md5.Sum(tmp)
	}

	for _, hash := range mHashes {
		if hash == pHash {
			return true
		}
	}
	return false
}

func ImprovedRabinKarp(s, p string) bool {
	mBytes := []byte(s)
	pBytes := []byte(p)

	ml := len(mBytes)
	pl := len(pBytes)

	pHash := irbHash(pBytes, 26, pl)
	mHashes := make([]int, ml-pl+1)

	i := 0
	mHashes[i] = irbHash(mBytes[i:i+pl], 26, pl)
	i++
	for ; i <= ml-pl; i++ {
		mHashes[i] = (mHashes[i-1]-int(mBytes[i-1])*int(math.Pow(float64(26), float64(pl-1))))*26 + int(mBytes[i+pl-1])
	}

	for _, hash := range mHashes {
		if hash == pHash {
			return true
		}
	}

	return false
}

func irbHash(b []byte, base, order int) int {
	res := 0
	for i := 0; i < order; i++ {
		res += int(b[i]) * int(math.Pow(float64(base), float64(order-i-1)))
	}
	return res
}
