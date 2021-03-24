package com.github.itsubaki.rsa

object RSA {
  def publicKey(e: Int): Int = Number.coprime(e)

  def privateKey(e: Int, E: Int): Int =
    LazyList.from(1).take(e).find(E * _ % e == 1).get

  def encrypt(message: Int, E: Int, N: Int): Int = mod(message, E, N, message)

  def decrypt(message: Int, D: Int, N: Int): Int = mod(message, D, N, message)

  @scala.annotation.tailrec
  private def mod(message: Int, e: Int, n: Int, acc: Int): Int =
    if (e == 1) acc else mod(message, e - 1, n, (acc * message) % n)
}
