package trafikinfo

import (
	"encoding/xml"
	"fmt"
	"reflect"
	"testing"
)

func TestNewQUery(t *testing.T) {}

func TestQueryDistinct(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).Distinct("test")
	if q.query.Distinct != "test" {
		t.Fatalf("Expected distinct value of: test, got: %s", q.query.Distinct)
	}

	q2 := NewQuery(WeatherStation, 1.0).Distinct("test").Distinct("foo")
	if q2.query.Distinct != "foo" {
		t.Fatalf("Expected distinct value of: foo, got: %s", q.query.Distinct)
	}
}

func TestQueryInclude(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output []string
	}{
		{Name: "One", Input: []string{"one"}, Output: []string{"one"}},
		{Name: "Two", Input: []string{"one", "two"}, Output: []string{"one", "two"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			q := NewQuery(WeatherStation, 1.0).Include(tt.Input...)
			if len(q.query.Include) != len(tt.Output) {
				t.Fatalf("Expected: %d elements, got: %d", len(tt.Output), len(q.query.Include))
			}
			if !reflect.DeepEqual(q.query.Include, tt.Output) {
				t.Fatalf("Expected: %#v, got: %#v", tt.Output, q.query.Include)
			}
		})
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Fluent%s", tt.Name), func(t *testing.T) {
			t.Parallel()
			q := NewQuery(WeatherStation, 1.0)
			for _, in := range tt.Input {
				q.Include(in)
			}
			t.Log(q.query.Include)
			if len(q.query.Include) != len(tt.Output) {
				t.Fatalf("Expected: %d elements, got: %d", len(tt.Output), len(q.query.Include))
			}
			if !reflect.DeepEqual(q.query.Include, tt.Output) {
				t.Fatalf("Expected: %#v, got: %#v", tt.Output, q.query.Include)
			}
		})
	}
}

func TestQueryExclude(t *testing.T) {
	tests := []struct {
		Name   string
		Input  []string
		Output []string
	}{
		{Name: "One", Input: []string{"one"}, Output: []string{"one"}},
		{Name: "Two", Input: []string{"one", "two"}, Output: []string{"one", "two"}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			q := NewQuery(WeatherStation, 1.0).Exclude(tt.Input...)
			if len(q.query.Exclude) != len(tt.Output) {
				t.Fatalf("Expected: %d elements, got: %d", len(tt.Output), len(q.query.Exclude))
			}
			if !reflect.DeepEqual(q.query.Exclude, tt.Output) {
				t.Fatalf("Expected: %#v, got: %#v", tt.Output, q.query.Exclude)
			}
		})
	}
	for _, tt := range tests {
		tt := tt
		t.Run(fmt.Sprintf("Fluent%s", tt.Name), func(t *testing.T) {
			t.Parallel()
			q := NewQuery(WeatherStation, 1.0)
			for _, in := range tt.Input {
				q.Exclude(in)
			}
			if len(q.query.Exclude) != len(tt.Output) {
				t.Fatalf("Expected: %d elements, got: %d", len(tt.Output), len(q.query.Exclude))
			}
			if !reflect.DeepEqual(q.query.Exclude, tt.Output) {
				t.Fatalf("Expected: %#v, got: %#v", tt.Output, q.query.Exclude)
			}
		})
	}
}

func TestQueryFilter(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0).Filter(Equal("a", "b"))
	if len(q.query.Filter.filter.Equal) != 1 {
		t.Fatalf("Expected 1 filter, got: %d", len(q.query.Filter.filter.Equal))
	}
	if q.query.Filter.filter.Equal[0].Name != "a" {
		t.Fatalf("Expected Equal filter with name=a, got name=%s", q.query.Filter.filter.Equal[0].Name)
	}
	if q.query.Filter.filter.Equal[0].Value != "b" {
		t.Fatalf("Expected Equal filter with value=b, got value=%s", q.query.Filter.filter.Equal[0].Value)
	}
}

func TestQueryMarshalXML(t *testing.T) {
	q := NewQuery(WeatherStation, 1.0)

	res, err := xml.Marshal(q)
	if err != nil {
		t.Fatal(err)
	}

	exp := `<Query objecttype="WeatherStation" schemaversion="1"></Query>`
	if string(res) != exp {
		t.Fatalf("Expected: %s, got: %s", exp, string(res))
	}
}

func TestCompositeFilter(t *testing.T) {
	tests := []struct {
		Name    string
		Func    func(f1, f2 Filter, filters ...Filter) Filter
		Filters []Filter
		Result  Filter
	}{
		{
			Name:    "And",
			Func:    And,
			Filters: []Filter{Equal("a", "b"), Equal("c", "d")},
			Result: Filter{filter: filter{
				And: []filter{{
					Equal: []comparison{
						{Name: "a", Value: "b"},
						{Name: "c", Value: "d"},
					},
				}},
			}},
		},
		{
			Name:    "Or",
			Func:    Or,
			Filters: []Filter{Equal("a", "b"), Equal("c", "d")},
			Result: Filter{filter: filter{
				Or: []filter{{
					Equal: []comparison{
						{Name: "a", Value: "b"},
						{Name: "c", Value: "d"},
					},
				}},
			}},
		},
		{
			Name:    "Not",
			Func:    Not,
			Filters: []Filter{Equal("a", "b"), Equal("c", "d")},
			Result: Filter{filter: filter{
				Not: []filter{{
					Equal: []comparison{
						{Name: "a", Value: "b"},
						{Name: "c", Value: "d"},
					},
				}},
			}},
		},
		{
			Name:    "ElementMatch",
			Func:    ElementMatch,
			Filters: []Filter{Equal("a", "b"), Equal("c", "d")},
			Result: Filter{filter: filter{
				ElementMatch: []filter{{
					Equal: []comparison{
						{Name: "a", Value: "b"},
						{Name: "c", Value: "d"},
					},
				}},
			}},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			var f Filter
			if len(tt.Filters) == 2 {
				f = tt.Func(tt.Filters[0], tt.Filters[1])
			} else {
				f = tt.Func(tt.Filters[0], tt.Filters[1], tt.Filters[2:]...)
			}

			if !reflect.DeepEqual(tt.Result, f) {
				t.Fatalf("Did not get expected filter set: %+v", f)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	f := Equal("field", "something")
	if len(f.filter.Equal) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Equal))
	}
	comp := f.filter.Equal[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestExists(t *testing.T) {
	f := Exists("field", true)
	if len(f.filter.Exists) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Exists))
	}
	comp := f.filter.Exists[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "true" {
		t.Fatalf("Expected filter with value=true, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestGreaterThan(t *testing.T) {
	f := GreaterThan("field", "something")
	if len(f.filter.GreaterThan) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.GreaterThan))
	}
	comp := f.filter.GreaterThan[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=true, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestGreaterThanOrEqual(t *testing.T) {
	f := GreaterThanOrEqual("field", "something")
	if len(f.filter.GreaterThanOrEqual) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.GreaterThanOrEqual))
	}
	comp := f.filter.GreaterThanOrEqual[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestLessThan(t *testing.T) {
	f := LessThan("field", "something")
	if len(f.filter.LessThan) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.LessThan))
	}
	comp := f.filter.LessThan[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestLeesThanOrEqual(t *testing.T) {
	f := LessThanOrEqual("field", "something")
	if len(f.filter.LessThanOrEqual) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.LessThanOrEqual))
	}
	comp := f.filter.LessThanOrEqual[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestNotEqual(t *testing.T) {
	f := NotEqual("field", "something")
	if len(f.filter.NotEqual) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.NotEqual))
	}
	comp := f.filter.NotEqual[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestLike(t *testing.T) {
	f := Like("field", "something")
	if len(f.filter.Like) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Like))
	}
	comp := f.filter.Like[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestNotLike(t *testing.T) {
	f := NotLike("field", "something")
	if len(f.filter.NotLike) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.NotLike))
	}
	comp := f.filter.NotLike[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestIn(t *testing.T) {
	f := In("field", "something")
	if len(f.filter.In) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.In))
	}
	comp := f.filter.In[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}

func TestNotIn(t *testing.T) {
	f := NotIn("field", "something")
	if len(f.filter.NotIn) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.NotIn))
	}
	comp := f.filter.NotIn[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestWithin(t *testing.T) {
	f := Within("field", "something", "box", 2.5)
	if len(f.filter.Within) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Within))
	}
	comp := f.filter.Within[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 2.5 {
		t.Fatalf("Expected filter with radius=2.5, but got radius=%f", radius)
	}
	if shape := comp.Shape; shape != "box" {
		t.Fatalf("Expected filter with shape=box, but got shape=%s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestIntersects(t *testing.T) {
	f := Intersects("field", "something")
	if len(f.filter.Intersects) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Intersects))
	}
	comp := f.filter.Intersects[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := comp.MinDistance; dst != nil {
		t.Fatalf("Did not expect filter to have mindistance, got: %d", dst)
	}
	if dst := comp.MaxDistance; dst != nil {
		t.Fatalf("Did not expect filter to have maxdistance, got: %d", dst)
	}
}
func TestNear(t *testing.T) {
	f := Near("field", "something", 1, 10)
	if len(f.filter.Near) != 1 {
		t.Fatalf("Expected 1 element, got: %d", len(f.filter.Near))
	}
	comp := f.filter.Near[0]
	if name := comp.Name; name != "field" {
		t.Fatalf("Expected filter with name=field, but got name=%s", name)
	}
	if value := comp.Value; value != "something" {
		t.Fatalf("Expected filter with value=something, but got value=%s", value)
	}
	if radius := comp.Radius; radius != 0 {
		t.Fatalf("Did not expect filter to have radius, got: %f", radius)
	}
	if shape := comp.Shape; shape != "" {
		t.Fatalf("Did not expect filter to have shape, got: %s", shape)
	}
	if dst := *comp.MinDistance; dst != 1 {
		t.Fatalf("Expected filter with mindistance=1, but got mindistance=%d", dst)
	}
	if dst := *comp.MaxDistance; dst != 10 {
		t.Fatalf("Expected filter with maxdistance=10, but got maxdistance=%d", dst)
	}
}

func TestFilterMarshalXML(t *testing.T) {
	f := Equal("a", "b")

	res, err := xml.Marshal(&f)
	if err != nil {
		t.Fatal(err)
	}

	exp := `<Filter><EQ name="a" value="b"></EQ></Filter>`
	if string(res) != exp {
		t.Fatalf("Expected: %s, got: %s", exp, string(res))
	}
}