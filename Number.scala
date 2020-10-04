import scala.annotation.switch
object Number {
  def gcd(a: Int, b: Int): Int = if (b == 0) a else gcd(b, a % b)

  def isPrime(N: Int): Boolean = {
    if (N == 2) true else if (N < 2 || N % 2 == 0) false
    if (
      LazyList
        .from(3, 2)
        .take(scala.math.sqrt(N).toInt)
        .filter(N % _ != 0)
        .size > 0
    ) true
    else false
  }

  def coprime(N: Int): Int =
    Iterator
      .continually(scala.util.Random.nextInt((N - 2) + 1))
      .find(gcd(N, _) == 1)
      .get
}
