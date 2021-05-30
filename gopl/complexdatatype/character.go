// Charcount computes counts of Unicode characters.
package main

//func main() {
//	s := "sdfjhdf487453984f47hr38fJHBHGDdhgvdggTFCdhfvshgfvsghdcsdhgvJHVghdvcgfcvdgsfFFVOOHh38rgfrb"
//	dataTypes := make(map[string]int)
//	for _, v := range s {
//		if unicode.IsDigit(rune(v)) {
//			dataTypes["digit"] += 1
//		}
//		if unicode.IsLower(rune(v)) {
//			dataTypes["lower"] += 1
//		}
//		if unicode.IsUpper(rune(v)) {
//			dataTypes["upper"] += 1
//		}
//	}
//	fmt.Println(dataTypes)
//	fmt.Println(unicode.IsLetter(rune(s[0])))
//	counts := make(map[rune]int)    // counts of Unicode characters
//	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
//	invalid := 0                    // count of invalid UTF-8 characters
//
//	in := bufio.NewReader(os.Stdin)
//	for {
//		r, n, err := in.ReadRune() // returns rune, nbytes, error
//		if err == io.EOF {
//			break
//		}
//		if err != nil {
//			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
//			os.Exit(1)
//		}
//		if r == unicode.ReplacementChar && n == 1 {
//			invalid++
//			continue
//		}
//		counts[r]++
//		utflen[n]++
//	}
//	fmt.Printf("rune\tcount\n")
//	for c, n := range counts {
//		fmt.Printf("%q\t%d\n", c, n)
//	}
//	fmt.Print("\nlen\tcount\n")
//	for i, n := range utflen {
//		if i > 0 {
//			fmt.Printf("%d\t%d\n", i, n)
//		}
//	}
//	if invalid > 0 {
//		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
//	}
//}
