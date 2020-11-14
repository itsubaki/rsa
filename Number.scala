object Number {
  def euler(p: Int, q: Int): Int = (p - 1) * (q - 1)

  @scala.annotation.tailrec
  def gcd(a: Int, b: Int): Int = if (b == 0) a else gcd(b, a % b)

  def rand(min: Int, max: Int): Int = scala.util.Random.nextInt(max - min) + min

  def isPrime(N: Int): Boolean =
    if (N == 2) true
    else if (N < 2 || N % 2 == 0) false
    else if (indivisible(N)) true
    else false

  def indivisible(N: Int): Boolean =
    LazyList
      .from(3, 2)
      .take(scala.math.sqrt(N).toInt)
      .find(N % _ == 0)
      .isEmpty

  def coprime(N: Int): Int =
    Iterator
      .continually(rand(2, N - 1))
      .find(gcd(N, _) == 1)
      .get
}
