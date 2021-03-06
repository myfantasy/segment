// Package segment - segmentation sourse
package segment

import "sort"

// Segment segment info
type Segment struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

// Segments List
// *Segments == nil -> all segments
type Segments struct {
	S []Segment `json:"s"`
}

// MakeSegments make empty Segments
func MakeSegments() (s *Segments) {
	return &Segments{S: make([]Segment, 0)}
}

// In key in Segments
// if s == nil returns true
func (s *Segments) In(key int64) bool {
	if s == nil {
		return true
	}
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
// if s == nil do nothing
func (s *Segments) AddSegment(sg Segment) *Segments {
	if s == nil {
		return s
	}

	l := len(s.S)

	if l == 0 {
		s.S = []Segment{sg}
		return s
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

	return s
}

// Add key in Segments
// if s == nil do nothing (look AddSegment)
func (s *Segments) Add(key int64) *Segments {

	s.AddSegment(Segment{From: key, To: key})

	return s
}

// CutSegment cut Segment in Segments
// if s == nil do nothing
func (s *Segments) CutSegment(sg Segment) *Segments {
	if s == nil {
		return s
	}

	l := len(s.S)

	if l == 0 {
		return s
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

	return s
}

// Cut key in Segments
// if s == nil do nothing (look CutSegment)
func (s *Segments) Cut(key int64) *Segments {

	s.CutSegment(Segment{From: key, To: key})

	return s
}

// Union Segments
// if s == nil do nothing
// if s2 == nil do nothing
// returns s after modification
func (s *Segments) Union(s2 *Segments) *Segments {
	if s == nil {
		return s
	}
	if s2 == nil {
		return s
	}

	for _, v := range s2.S {
		s.AddSegment(v)
	}

	return s
}

// Split - splits Segments to segments list by segment
// nil -> []*Segments{nil}
func (s *Segments) Split() []*Segments {
	if s == nil {
		return []*Segments{nil}
	}

	var res []*Segments

	for _, sg := range s.S {
		res = append(res, &Segments{
			S: []Segment{sg},
		})
	}

	return res
}

// Less - s less then s2 (s starts befor s2)
// nil segments in result head
// empty segments in result tail
func (s *Segments) Less(s2 *Segments) bool {
	if s2 == nil {
		return false
	}
	if s == nil {
		return true
	}
	if len(s.S) == 0 {
		return false
	}
	if len(s2.S) == 0 {
		return true
	}
	return s.S[0].From < s2.S[0].From
}
