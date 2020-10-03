object Hello {
  def main(args: Array[String]): Unit = {
    val p = 17
    val q = 19

    if (!Number.IsPrime(p)) return
    if (!Number.IsPrime(q)) return

    val phi = (p - 1) * (q - 1)
    val N = p * q
    val E = Number.Coprime(phi)
    val D = RSA.FindD(phi, E)
    println(s"p=$p, q=$q, N=$N, E=$E, D=$D, phi=$phi")

    val message = scala.util.Random.nextInt((N - 2))
    val enc = RSA.Encrypt(message, E, N)
    val dec = RSA.Decrypt(enc, D, N)
    println(s"message: $message, encrypted: $enc, decrypted: $dec")
  }
}
