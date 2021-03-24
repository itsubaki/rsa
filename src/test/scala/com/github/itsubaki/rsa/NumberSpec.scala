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

  "divisible" should "returns true when number N is divisible 3 to sqrt(N)" in {
    assert(Number.divisible(3) === false)
    assert(Number.divisible(4) === true)
    assert(Number.divisible(5) === false)
    assert(Number.divisible(6) === true)
    assert(Number.divisible(7) === false)
    assert(Number.divisible(8) === true)
    assert(Number.divisible(9) === true)
    assert(Number.divisible(10) === true)
    assert(Number.divisible(11) === false)
    assert(Number.divisible(12) === true)
    assert(Number.divisible(13) === false)
    assert(Number.divisible(14) === true)
    assert(Number.divisible(15) === true)
  }

  "coprime" should "returns co-prime number of N" in {
    assert(Number.gcd(15, Number.coprime(15)) === 1)
    assert(Number.gcd(21, Number.coprime(21)) === 1)
    assert(Number.gcd(33, Number.coprime(33)) === 1)
  }
}
