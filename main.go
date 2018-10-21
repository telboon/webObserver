package main

import (
    "os/exec"
    "os"
    "io/ioutil"
    "fmt"
    "strings"
    "bytes"
    "github.com/yosssi/gohtml"
    "github.com/glaslos/ssdeep"
    "time"
)

const (
    curlfile = "./curlfile.txt"
    resultpath = "./results/"
    diffpath = "./diff/"
    hashTreshold = 90
)

func main() {
    var scoreOutput string

    fullBytes, _ := ioutil.ReadFile(curlfile)
    fullStr := string(fullBytes)
    siteList := make([]string, 0)
    siteList = strings.Split(fullStr, "\n")

    siteNo := len(siteList)/2

    for i:=0; i<siteNo; i++ {
        scoreOutput += runAndCompare(resultpath + siteList[i*2], siteList[i*2 + 1])
    }

    fmt.Println(scoreOutput)
}

func runAndCompare(name string, cmdStr string) string {
    var scoreOutput string
    var diffOutput string
    siteContent := runCmd(fmt.Sprintf("%s", cmdStr))
    siteContent = gohtml.Format(siteContent)
    siteContent = shittyParser(siteContent)


    writeFile(name+"_new", siteContent)

    if _, err := os.Stat(name); !os.IsNotExist(err) {
        hashScore := fuzzyCompare(name + "_new", name)

        scoreOutput += fmt.Sprintf("%s: %d\n", name, hashScore)

        if hashScore < hashTreshold {
            diffOutput = runCmd(fmt.Sprintf("diff %s_new %s", name, name))
        }
        appendFile(diffpath + name[strings.LastIndex(name,"/"):] + "_diff", fmt.Sprintf("_________________________________ %s __________________________________\n", time.Now().Format("2006 Jan 4 06:04:05")))
        appendFile(diffpath + name[strings.LastIndex(name,"/"):] + "_diff", diffOutput)
    }
    runCmd(fmt.Sprintf("mv %s_new %s", name, name))

    return scoreOutput
}

func runCmd(cmdStr string) string {
    var output bytes.Buffer
    cmd := exec.Command("bash", "-c", cmdStr)
    cmd.Stdout = &output
    cmd.Start()
    cmd.Wait()

    return output.String()
}

func writeFile(filename string, text string) {
    ioutil.WriteFile(filename, []byte(text), 0664)
}

func fuzzyCompare(file1 string, file2 string) int {
    tempFile, _ := os.Open(file1)
    hash1, _ := ssdeep.FuzzyFile(tempFile)

    tempFile, _ = os.Open(file2)
    hash2, _ := ssdeep.FuzzyFile(tempFile)

    result, _ := ssdeep.Distance(hash1, hash2)
    return result
}

func appendFile(filename string, text string) {
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0664)
    if err != nil {
        writeFile(filename, text)
    } else {

        defer f.Close()

        if _, err = f.WriteString(text); err != nil {
            panic(err)
        }
    }
}

func shittyParser(jscode string) string {
    var cleaned string
    cleaned = jscode

    cleaned = addNewLine(cleaned, "{")
    cleaned = addNewLine(cleaned, "}")
    cleaned = addNewLine(cleaned, ";")

    return cleaned
}

func addNewLine(code string, iden string) string {
    return strings.Replace(code, iden, iden+"\n", -1)
}
