package main

import (
	"fmt"
	"bufio"
	"bytes"
)




/* Exercise 7.1
*/
type wordCounter int
type lineCounter int

func (wc *wordCounter) Write(p []byte) (int, error){
	bReader := bytes.NewReader(p)
	bScanner := bufio.NewScanner(bReader)
	bScanner.Split(bufio.ScanWords)
	for bScanner.Scan(){
		*wc++
	}
	if bScanner.Err() != nil {
		return 0, bScanner.Err()
	}
	return int(*wc), nil

}

func (lw *lineCounter) Write(p []byte) (int, error){
	bScanner := bufio.NewScanner(bytes.NewReader(p))
	bScanner.Split(bufio.ScanLines)

	for bScanner.Scan(){
		*lw++
	}
	if bScanner.Err() != nil {
		return 0, bScanner.Err()
	}
	return int(*lw), nil

}



func main(){

	words := [] byte("what a wonderful word")
	lines := [] byte(`I see trees of green, red roses too
I see them bloom for me and you
And I think to myself, what a wonderful world
I see skies of blue and clouds of white
The bright blessed day, the dark sacred night
And I think to myself, what a wonderful world
The colors of the rainbow, so pretty in the sky
Are also on the faces of people going by
I see friends shaking hands, saying' "How do you do?"
They're really saying "I love you"
I hear babies cry, I watch them grow
They'll learn much more than I'll ever know
And I think to myself, what a wonderful world
And I think to myself, what a wonderful world`)
	var wc wordCounter
	var lc lineCounter
	wordNum, _ := wc.Write(words)
	lineNum, _ := lc.Write(lines)

	fmt.Println(wordNum)
	fmt.Println(lineNum)
}