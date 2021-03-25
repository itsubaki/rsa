import org.scalatest.flatspec.AnyFlatSpec
import org.scalatest.diagrams.Diagrams
import com.github.itsubaki.rsa.Number

class NumberSpec extends AnyFlatSpec with Diagrams {

  "gcd" should "returns the largest positive integer that divides each of the integers" in {
    assert(Number.gcd(15, 3) === 3)
    assert(Number.gcd(15, 5) === 5)
    assert(Number.gcd(15, 4) === 1)
    assert(Number.gcd(15, 7) === 1)
    assert(Number.gcd(15, 11) === 1)
    assert(Number.gcd(15, 13) === 1)
  }

  "isPrime" should "returns true when number N is prime" in {
    assert(Number.isPrime(2) === true)
    assert(Number.isPrime(3) === true)
    assert(Number.isPrime(4) === false)
    assert(Number.isPrime(5) === true)
    assert(Number.isPrime(6) === false)
    assert(Number.isPrime(7) === true)
    assert(Number.isPrime(8) === false)
    assert(Number.isPrime(9) === false)
    assert(Number.isPrime(10) === false)
    assert(Number.isPrime(11) === true)
    assert(Number.isPrime(12) === false)
    assert(Number.isPrime(13) === true)
    assert(Number.isPrime(14) === false)
    assert(Number.isPrime(15) === false)
  }

  "individable" should "returns true when number N > 8 is not divide by 3 to sqrt(N)" in {
    assert(Number.individable(9) === false)
    assert(Number.individable(10) === false)
    assert(Number.individable(11) === true)
    assert(Number.individable(12) === false)
    assert(Number.individable(13) === true)
    assert(Number.individable(14) === false)
    assert(Number.individable(15) === false)
  }

  "coprime" should "returns co-prime number of N" in {
    assert(Number.gcd(15, Number.coprime(15)) === 1)
    assert(Number.gcd(21, Number.coprime(21)) === 1)
    assert(Number.gcd(33, Number.coprime(33)) === 1)
  }
}
