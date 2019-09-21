# set
## Golang implementation of the set datastructure

If you use this package, please consider donating at PayPal: [https://www.paypal.me/jismithpp](https://www.paypal.me/jismithpp)

### Example

	s := set.New()
	s.SetMultiMode(true) // counts the number of times a value is added
	s.SetCap(8) // allocates enough space to add 8 items
	s.Add(0, 1, 2, 3, 4, 5, 6, 7)
	s.Add(7)
	fmt.Println(s) // displays the entries