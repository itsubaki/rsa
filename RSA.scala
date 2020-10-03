object RSA {
  def FindD(phi: Int, e: Int): Int = {
    for (d <- 1 to phi) {
      val n = (e * d) % phi
      if (n == 1) {
        return d
      }
    }

    throw new RuntimeException("invalid")
  }

  def Encrypt(message: Int, e: Int, n: Int): Int = {
    var enc = message
    for (i <- 1 until e) {
      enc = (enc * message) % n
    }

    return enc
  }

  def Decrypt(message: Int, d: Int, n: Int): Int = {
    var dec = message
    for (i <- 1 until d) {
      dec = (dec * message) % n
    }

    return dec
  }
}
