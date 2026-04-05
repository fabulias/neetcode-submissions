type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	sb := strings.Builder{}
	sizes := []string{}
	for _, str := range strs {
		sizes = append(sizes, strconv.Itoa(len(str)))
	}
	sb.WriteString(strings.Join(sizes, ","))
	sb.Write([]byte("#"))
	sb.Write([]byte(strings.Join(strs,"")))
	return sb.String()
}

func (s *Solution) Decode(encoded string) []string {
	fmt.Println(encoded)
	if encoded == "" {return []string{}}
	twoPieces := strings.SplitN(encoded, "#", 2) // one piece reference 
	if len(twoPieces) != 2 {
		return []string{}
	}
	onePiece := strings.Split(twoPieces[0],",")
	result := []string{}
	start := 0
	fmt.Println(onePiece)
	for _, length := range onePiece {
		if length == "" {
			fmt.Println("empty string", length)
			continue
		}
		l, err := strconv.Atoi(length)
		if err != nil {
			fmt.Println("error", err)
			continue
		}

		end := start + l
		fmt.Println(start, end)
		result = append(result, twoPieces[1][start:end])
		start += l
	}
	return result
}
