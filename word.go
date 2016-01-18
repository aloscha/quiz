//Oleksiy Yatsko

package main

import (
	"flag"
	"io/ioutil"
	"log"
	"strings"
	"fmt"
	"sort"
)

type ByLen []string
func (a ByLen) Len() int           { return len(a) }
func (a ByLen) Less(i, j int) bool { return len(a[i]) < len(a[j]) }
func (a ByLen) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Node int
const (
	VISITED Node = 1 + iota
	PREFIX
	ENDWORD
	BADPREFIX
)
var wordMap = make(map[string]Node)

func main() {
	fileName := flag.String("file", "word.list", "flag takes console/file values");
	flag.Parse();
	readDataFromFile(*fileName);
}
func readDataFromFile(fileName string){
	data, err := ioutil.ReadFile(fileName);
	if err != nil {
		log.Fatal("Error occured: ", err);
	}
	
	if len(data) == 0 {
		log.Fatal("Error occured: File is empty");
	}
	
	dataConverted := strings.Fields(string(data));
	engine(dataConverted)
}
func engine(words []string){
	sort.Sort(ByLen(words))
	builCleverMap(words);
	findLongestCompoundWord(words);
}

func builCleverMap(words []string){
	for _,word := range words {
		wordLength := len(word);
		for i := 1; i < wordLength; i++{
			substringWord := word[0:i];
			if _, ok := wordMap[substringWord]; !ok {
				wordMap[substringWord] = VISITED;
			}  else {
				if wordMap[substringWord] == ENDWORD {
					wordMap[substringWord] = PREFIX;
				}
			}
		}
		wordMap[word] = ENDWORD
	}
}
func findLongestCompoundWord(words []string){
	longestWord := "";
	for i := len(words)-1; i > 0; i-- {
		word := words[i]
		if wordMap[word] == ENDWORD {
			if checkCompoundWord(word) {
				longestWord = word;
				break;
			}
		}
	}
	if len(longestWord) == 0 {
		fmt.Printf("No compound word find\n")
	} else {
		fmt.Printf("Longest coumpound word is: %s\n", longestWord)
	}
}
func checkCompoundWord(word string) bool{
	wordtmp := word;
	isCompund := false;
	wordMapTmp := make(map[string]Node)
	for k, v := range wordMap {
		wordMapTmp[k] = v
	}	
	for !isCompund && wordtmp == word {
		currentIndex := len(wordtmp)-1;
		lastPrefix := "";
		for currentIndex != 0 {
			nodeValue := wordMapTmp[wordtmp[0:currentIndex]];
			if nodeValue == PREFIX || nodeValue == ENDWORD {
				lastPrefix = wordtmp[0:currentIndex];
				wordtmp = wordtmp[currentIndex:len(wordtmp)];
				currentIndex = len(wordtmp);
			} else {
				currentIndex--;
			}
		}
		
		if len(wordtmp) == 0 {
			isCompund = true;
		} else if wordtmp == word {
			wordtmp = "";
		} else {
			wordtmp = word
			wordMapTmp[lastPrefix] = BADPREFIX;
		}
	}
	return isCompund;
}









