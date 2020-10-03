object Number {
  def IsPrime(N: Int): Boolean = {
    if (N < 2) return false
    if (N == 2) return true
    if (N % 2 == 0) return false

    for (i <- 3 to scala.math.sqrt(N).toInt if i % 2 != 0) {
      if (N % i == 0) {
        return false
      }
    }

    return true
  }

  def Coprime(N: Int): Int = {
    val rand = scala.util.Random

    while (true) {
      val v = 2 + rand.nextInt((N - 2) + 1)
      if (GCD(N, v) == 1) {
        return v
      }
    }

    throw new RuntimeException("invalid")
  }

  def GCD(a: Int, b: Int): Int = {
    if (b == 0) return a
    return GCD(b, a % b)
  }
}
