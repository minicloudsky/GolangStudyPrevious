package main

//
//func main() {
//	seen := make(map[string]bool) // a set of strings
//	input := bufio.NewScanner(os.Stdin)
//	for input.Scan() {
//		line := input.Text()
//		if !seen[line] {
//			seen[line] = true
//			fmt.Println(line)
//		}
//	}
//
//	if err := input.Err(); err != nil {
//		_, err := fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
//		if err != nil {
//			return
//		}
//		os.Exit(1)
//	}
//}
