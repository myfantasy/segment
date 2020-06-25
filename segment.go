// Package segment - segmentation sourse
package segment

import "sort"

// Segment segment info
type Segment struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

// Segments List
type Segments struct {
	S []Segment `json:"s"`
}

// In key in Segments
func (s *Segments) In(key int64) bool {
	l := len(s.S)
	if l == 0 {
		return false
	}
	el := sort.Search(l, func(i int) bool {
		return s.S[i].To >= key
	})

	return s.S[el].From <= key && s.S[el].To >= key
}

// AddSegment add Segment in Segments
func (s *Segments) AddSegment(sg Segment) {

	l := len(s.S)

	if l == 0 {
		s.S = []Segment{sg}
		return
	}

	res := make([]Segment, 0, l)
	ok := false

	for i := 0; i < l; i++ {

		if ok {
			res = append(res, s.S[i])
			continue
		}

		if sg.From > s.S[i].To+1 {
			res = append(res, s.S[i])
			continue
		}

		if sg.To+1 < s.S[i].From {
			res = append(res, sg, s.S[i])
			ok = true
			continue
		}

		if sg.From > s.S[i].From {
			sg.From = s.S[i].From
		}

		if sg.To < s.S[i].To {
			sg.To = s.S[i].To
		}
	}

	if !ok {
		res = append(res, sg)
	}

	s.S = res
}

// Add key in Segments
func (s *Segments) Add(key int64) {

	s.AddSegment(Segment{From: key, To: key})
}

// CutSegment cut Segment in Segments
func (s *Segments) CutSegment(sg Segment) {

	l := len(s.S)

	if l == 0 {
		return
	}

	res := make([]Segment, 0, l)
	ok := false

	for i := 0; i < l; i++ {

		if ok {
			res = append(res, s.S[i])
			continue
		}

		if sg.From > s.S[i].To {
			res = append(res, s.S[i])
			continue
		}

		if sg.To < s.S[i].From {
			res = append(res, s.S[i])
			ok = true
			continue
		}

		if sg.From > s.S[i].From {
			res = append(res, Segment{From: s.S[i].From, To: sg.From - 1})
		}

		if sg.To < s.S[i].To {
			res = append(res, Segment{From: sg.To + 1, To: s.S[i].To})
		}
	}

	s.S = res
}

// Cut key in Segments
func (s *Segments) Cut(key int64) {

	s.CutSegment(Segment{From: key, To: key})
}
