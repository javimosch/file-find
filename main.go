package main
import ("fmt";"os";"path/filepath";"strings")
func main() {
	root := "."; pattern := ""
	if len(os.Args) > 1 { root = os.Args[1] }
	if len(os.Args) > 2 { pattern = os.Args[2] }
	fmt.Println("[")
	first := true
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil { return nil }
		if info.IsDir() { return nil }
		if pattern != "" && !strings.Contains(info.Name(), pattern) && !globMatch(pattern, info.Name()) { return nil }
		if !first { fmt.Println(",") }; first = false
		fmt.Printf(`{"file":"%s","size":%d,"mode":"%s"}`, path, info.Size(), info.Mode().String())
		return nil
	})
	fmt.Println("\n]")
}
func globMatch(pattern, name string) bool {
	if !strings.Contains(pattern, "*") { return strings.HasPrefix(name, pattern) }
	parts := strings.Split(pattern, "*")
	if len(parts) == 1 { return strings.HasPrefix(name, parts[0]) }
	if !strings.HasPrefix(name, parts[0]) { return false }
	n := name[len(parts[0]):]
	for _, part := range parts[1:] {
		if part == "" { return true }
		idx := strings.Index(n, part)
		if idx < 0 { return false }
		n = n[idx+len(part):]
	}
	return true
}
