package probability


func ChooseHelper(n int, k int, cache map[string]int64) int64 {
  // chooseHelper implements the recursive method of determining binomial
  // coefficients with memoization, as indicated at
  // http://en.wikipedia.org/wiki/Binomial_coefficient

  // if n==k we'll reduce to the k = 0 case and don't need to traverse
  if n == k {
    return 1
  }
  // if n < k, we'll reduce to the n = 0 case and don't need to traverse
  if n < k {
    return 0
  }
  if n == 0 {
    return 0
  }
  if k == 0 {
    return 1
  }

  // take advantage of symetry of pascals triangle and normalize to the
  // "left half" of the tree
  if k > (n / 2) {
    k = n - k
  }
  nCk := string(n) + "C" + string(k)

  value, ok := cache[nCk]
  if ok {
    return value
  }

  value = ChooseHelper(n-1, k-1, cache) + ChooseHelper(n-1, k, cache)
  cache[nCk] = value
  return value
}

func Choose(n, k int) int64 {
  cache := make(map[string]int64)
  return ChooseHelper(n, k, cache)
}

