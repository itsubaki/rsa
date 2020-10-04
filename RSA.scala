object RSA {
  def findD(phi: Int, e: Int): Int =
    LazyList.from(1).take(phi).find(d => (e * d) % phi == 1).get

  @scala.annotation.tailrec
  def mod(message: Int, e: Int, n: Int, acc: Int): Int =
    if (e == 1) acc else mod(message, e - 1, n, (acc * message) % n)

  def encrypt(message: Int, e: Int, n: Int): Int = mod(message, e, n, message)
  def decrypt(message: Int, d: Int, n: Int): Int = mod(message, d, n, message)
}
