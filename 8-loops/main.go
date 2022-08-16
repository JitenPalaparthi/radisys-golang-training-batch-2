package main

func main() {

	for i := 1; i <= 10; i++ {
		print(i, " ")
	}

	println()

	j := 1

	for ; j <= 10; j++ {
		print(j, " ")
	}

	println()

	k := 1
	for k <= 10 { // like while loop
		print(k, " ")
		k++
	}
	println()
	l := 1

	for {
		if l > 10 {
			break
		}
		print(l, " ")
		l++
	}

	// even numbers
	println()
	for i := 1; i <= 20; i++ {
		if i%2 != 0 {
			continue
		}
		print(i, " ")
	}

	println()
	for i := 2; i <= 20; i += 2 {
		print(i, " ")
	}

	println()
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			print(i, " ")
		}
	}
	println()

	for i := 0; i <= 5; i++ {
		for j := 0; j <= 5; j++ {
			if j > 3 {
				break
			}
			print("i:", i, " j:", j, " ")
		}
		println()
	}

	println()
outer:
	for i := 0; i <= 5; i++ {
		for j := 3; j <= 5; j++ {
			if i == j {
				break outer
			}
			print("i:", i, " j:", j, " ")
		}
		println()
	}
}
