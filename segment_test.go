package segment

import (
	"testing"
)

func TestSegment(t *testing.T) {
	s := Segments{}
	r := s.In(43)
	if r {
		t.Fatal("Empty segment in != true ")
	}

	s.Add(43)
	r = s.In(43)
	if !r {
		t.Fatal(s, "NOT empty segment in = true ")
	}

	s.Add(45)

	s.Add(47)

	s.Add(41)

	s.Add(44)

	if len(s.S) != 3 {
		t.Fatal(s, " not Correct Union")
	}

	s.AddSegment(Segment{From: 20, To: 46})

	if len(s.S) != 1 {
		t.Fatal(s, " not Correct Union s2")
	}

	s.Cut(33)
	s.Cut(32)
	s.Cut(34)
	s.Cut(20)
	s.Cut(47)
	s.Cut(48)
	s.Cut(19)

	if len(s.S) != 2 {
		t.Fatal(s, " not Correct Cut")
	}

	s.CutSegment(Segment{From: 20, To: 46})

	if len(s.S) != 0 {
		t.Fatal(s, " not Correct Cut 2")
	}
}
