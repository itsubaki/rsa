object Hello {
  def main(args: Array[String]): Unit = {
    val p = 17
    val q = 19

    if (!Number.isPrime(p)) return
    if (!Number.isPrime(q)) return

    val phi = (p - 1) * (q - 1)
    val N = p * q
    val E = RSA.publicKey(phi)
    val D = RSA.privateKey(phi, E)
    println(s"p=$p, q=$q, N=$N, E=$E, D=$D, phi=$phi")

    val message = scala.util.Random.nextInt((N - 2))
    val enc = RSA.encrypt(message, E, N)
    val dec = RSA.decrypt(enc, D, N)
    println(s"message: $message, encrypted: $enc, decrypted: $dec")
  }
}
