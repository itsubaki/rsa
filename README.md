# rsa

```scala
object Hello {
  def main(args: Array[String]): Unit = {
    val p = 17
    val q = 19

    if (!Number.isPrime(p)) return
    if (!Number.isPrime(q)) return

    val phi = Number.euler(p, q)
    val E = RSA.publicKey(phi)
    val D = RSA.privateKey(phi, E)
    println(s"p=$p, q=$q, E=$E, D=$D, phi=$phi")

    val N = p * q
    val message = scala.util.Random.nextInt(N - 2) + 1

    val enc = RSA.encrypt(message, E, N)
    val dec = RSA.decrypt(enc, D, N)
    println(s"message: $message, encrypted: $enc, decrypted: $dec")
  }
}

// p=17, q=19, N=323, E=185, D=137, phi=288
// message: 177, encrypted: 214, decrypted: 177
```