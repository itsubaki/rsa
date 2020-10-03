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

  def Encrypt(x: Int, e: Int, n: Int): Int = {
    var enc = x
    for (i <- 1 until e) {
      enc = (enc * x) % n
    }

    return enc
  }

  def Decrypt(m: Int, d: Int, n: Int): Int = {
    var dec = m
    for (i <- 1 until d) {
      dec = (dec * m) % n
    }

    return dec
  }
}
