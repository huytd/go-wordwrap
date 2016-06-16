package main_test

import (
	"github.com/huytd/go-wordwrap"
	"testing"
)

var (
	InputString = "A quick brown-fox jumps over the lazy-dog"
	InputWord = "jumps"
	InputWords = "quick brown"
	InputConnectedWord = "brown-fox"
)

// Wrap Word Test

func TestWrapWordAtMidStart(t *testing.T) {
	// jumps
	//  ^ wrap here
	expect := "\njumps"
	if main.WordWrap(InputWord, 2) != expect {
		t.Fail()
	}
}

func TestWrapWordAtMidEnd(t *testing.T) {
	// jumps
	//    ^ wrap here
	expect := "jumps\n"
	if main.WordWrap(InputWord, 4) != expect {
		t.Fail()
	}
}

// Wrap Multiple Words Test

func TestWrapAtFirstWord(t *testing.T) {
	// quick brown
	//    ^
	expect := "quick\nbrown\n"
	if main.WordWrap(InputWords, 4) != expect {
		t.Fail()
	}
}

func TestWrapAtSpace(t *testing.T) {
	// quick brown
	//      ^
	expect := "quick\nbrown"
	if main.WordWrap(InputWords, 6) != expect {
		t.Fail()
	}
}

func TestwrapAtSecondWord(t *testing.T) {
	// quick brown
	//         ^
	expect := "quick\nbrown"
	if main.WordWrap(InputWords, 9) != expect {
		t.Fail()
	}
}

// Wrap Connected Word Test

func TestWrapAtBeginOfFirstConnectedWord(t *testing.T) {
	// brown-fox
	//  ^
	expect := "\nbrown-fox"
	if main.WordWrap(InputConnectedWord, 2) != expect {
		t.Fail()
	}
}

func TestWrapAtEndOfFirstConnectedWord(t *testing.T) {
	// brown-fox
	//    ^ 
	expect := "brown-\nfox"
	if main.WordWrap(InputConnectedWord, 4) != expect {
		t.Fail()
	}
}

func TestWrapAtConnectCharacter(t *testing.T) {
	// brown-fox
	//      ^
	expect := "brown-\nfox"
	if main.WordWrap(InputConnectedWord, 6) != expect {
		t.Fail()
	}
}

func TestWrapAtBeginOfSecondWord(t *testing.T) {
	// brown-fox
	//       ^
	expect := "brown-\nfox"
	if main.WordWrap(InputConnectedWord, 7) != expect {
		t.Fail()
	}
}

func TestWrapAtEndOfSecondWord(t *testing.T) {
	// brown-fox
	//         ^
	expect := "brown-fox\n"
	if main.WordWrap(InputConnectedWord, 9) != expect {
		t.Fail()
	}
}

// Wrap String Test - Complex

func TestComplexWrapAtSpace(t *testing.T) {
	// brown-fox jumps
	//          ^ wrap here
	expect := "A quick brown-fox\njumps over the\nlazy-dog"
  if main.WordWrap(InputString, 18) != expect {
		t.Fail()
	}
}

func TestComplexWrapAtMidStartWord(t *testing.T) {
	// jumps
	//  ^ wrap here
	expect := "A quick brown-fox\njumps over the lazy-\ndog"
	if main.WordWrap(InputString, 20) != expect {
		t.Fail()
	}
}

func TestComplexWrapAtMidEndWord(t *testing.T) {
	// jumps
  //    ^ wrap here
	expect := "A quick brown-fox jumps\nover the lazy-dog"
	if main.WordWrap(InputString, 22) != expect {
		t.Fail()
	}
}

func TestComplexWrapAtMidStartOfConnectedWord(t *testing.T) {
	// brown-fox
	//  ^ wrap here
	expect := "A quick\nbrown-fox\njumps over\nthe lazy-\ndog"
	if main.WordWrap(InputString, 10) != expect {
		t.Fail()
	}
}

func TestComplexWrapAtConnectCharacter(t *testing.T) {
	// brown-fox
	//      ^ wrap here
	expect := "A quick brown-\nfox jumps over\nthe lazy-dog"
	if main.WordWrap(InputString, 14) != expect {
		t.Fail()
	}
}

func TestComplexWrapAfterConnectCharacter(t *testing.T) {
	// brown-fox
	//        ^ wrap here
	expect := "A quick brown-fox\njumps over the\nlazy-dog"
	if main.WordWrap(InputString, 16) != expect {
		t.Fail()
	}
}

func TestComplexWrapAtTheFirstCharacter(t *testing.T) {
	// The 
	// ^ wrap here
	expect := "A\nq\nu\ni\nc\nk\nb\nr\no\nw\nn\n-\nf\no\nx\nj\nu\nm\np\ns\no\nv\ne\nr\nt\nh\ne\nl\na\nz\ny\n-\nd\no\ng"
	if main.WordWrap(InputString, 1) != expect {
		t.Fail()
	}
}

