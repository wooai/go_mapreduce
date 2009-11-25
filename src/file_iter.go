package file_iter

import (
    "os";
    "bufio";
)

func EachLine(filename string) chan string {
    output := make(chan string);
    go func() {
        file, err := os.Open(filename, os.O_RDONLY, 0);
        if err != nil {
            return;
        }
        defer file.Close();
        reader := bufio.NewReader(file);
        for {
            line, err := reader.ReadString('\n');
            output <- line;
            if err == os.EOF {
                break;
            }
        }
        close(output);
    }();
    return output;
}
