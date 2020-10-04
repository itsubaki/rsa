object Hello {
  def main(args: Array[String]): Unit = {
    val p = 17
    val q = 19

    if (!Number.isPrime(p)) return
    if (!Number.isPrime(q)) return

    val e = Number.euler(p, q)
    val E = RSA.publicKey(e)
    val D = RSA.privateKey(e, E)
    println(s"p=$p, q=$q, E=$E, D=$D, euler=$e")

    val N = p * q
    val message = scala.util.Random.nextInt(N - 4) + 2 // 2 <= message <= N - 2

    val enc = RSA.encrypt(message, E, N)
    val dec = RSA.decrypt(enc, D, N)
    println(s"N: $N, message: $message, encrypted: $enc, decrypted: $dec")
  }
}
