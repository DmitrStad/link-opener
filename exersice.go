package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/fumiama/go-docx"
)

var mut sync.Mutex

func main() {

	reader, writer := io.Pipe()

	// 2. Запускаем горутину, которая пишет в writer
	go func() {
		defer writer.Close() // закрываем, когда вывод закончен
		WriteDocxContent(writer)
	}()

	// 3. Читаем из reader и строим map
	counts := countLinesFromReader(reader)

	// 4. Выводим результат
	fmt.Println("Частота строк:")
	for line, cnt := range counts {
		fmt.Printf("%q: %d\n", line, cnt)
	}
	defer func() {
		word := Map_to_String(counts)
		fmt.Println(word)
	}()
}

func Map_Append(answer map[string]int) {
	for _, arg := range os.Args[1:] {
		answer[arg]++
	}

}

func Map_Print(answer map[string]int) {
	for word, count := range answer {
		fmt.Println("Word: ", word, "count: ", count)
	}
}

func Map_to_String(answer map[string]int) string {

	mut.Lock()
	defer mut.Unlock()
	sentence := ""
	keys := make([]string, 0, len(answer))
	for index := range answer {
		keys = append(keys, index)
	}
	sort.Strings(keys)

	for _, key := range keys {

		fmt.Printf("Открываем: %s\n", hyperlink(key, key))
		err := openBrowser(key)
		if err != nil {
			fmt.Printf("Ошибка при открытии %s: %v\n", key, err)
		}

		time.Sleep(500 * time.Millisecond) // задержка между открытиями

	}

	return sentence
}

func hyperlink(url, text string) string {
	// escape-последовательность: \x1b]8;;URL\x1b\\TEXT\x1b]8;;\x1b\\
	return fmt.Sprintf("\x1b]8;;%s\x1b\\%s\x1b]8;;\x1b\\", url, text)
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "darwin":
		cmd = exec.Command("open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	// Запускаем процесс в фоне и не ждём его завершения
	return cmd.Start()
}

func WriteDocxContent(w io.Writer) {
	readFile, err := os.Open("sample.docx")
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileinfo, err := readFile.Stat()
	if err != nil {
		panic(err)
	}
	doc, err := docx.Parse(readFile, fileinfo.Size())
	if err != nil {
		panic(err)
	}

	for _, it := range doc.Document.Body.Items {
		switch v := it.(type) {
		case *docx.Paragraph:
			// Извлекаем текст параграфа
			fmt.Fprintln(w, it)
			//fmt.Println(it)
		case *docx.Table:
			// Используем ваш метод String() для таблицы
			fmt.Fprintln(w, v.String())
		}
	}

}

func countLinesFromReader(r io.Reader) map[string]int {
	scanner := bufio.NewScanner(r)
	counts := make(map[string]int)

	var line string
	for scanner.Scan() {
		line += scanner.Text() + " "
	}
	re := regexp.MustCompile(`http?:\\[^\s<>"']+`)
	urls := re.FindAllString(line, -1)

	for _, url := range urls {
		url = strings.ToLower(url)
		counts[url]++
	}
	return counts
}
