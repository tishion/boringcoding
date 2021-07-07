package httprouter

import (
	"reflect"
	"testing"
)

func Test_min(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countParams(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want uint8
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countParams(tt.args.path); got != tt.want {
				t.Errorf("countParams() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_incrementChildPrio(t *testing.T) {
	type fields struct {
		path      string
		wildChild bool
		nType     nodeType
		maxParams uint8
		priority  uint32
		indices   string
		children  []*node
		handle    RouteHandler
	}
	type args struct {
		pos int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				path:      tt.fields.path,
				wildChild: tt.fields.wildChild,
				nType:     tt.fields.nType,
				maxParams: tt.fields.maxParams,
				priority:  tt.fields.priority,
				indices:   tt.fields.indices,
				children:  tt.fields.children,
				handle:    tt.fields.handle,
			}
			if got := n.incrementChildPrio(tt.args.pos); got != tt.want {
				t.Errorf("node.incrementChildPrio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_addRoute(t *testing.T) {
	type fields struct {
		path      string
		wildChild bool
		nType     nodeType
		maxParams uint8
		priority  uint32
		indices   string
		children  []*node
		handle    RouteHandler
	}
	type args struct {
		path   string
		handle RouteHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				path:      tt.fields.path,
				wildChild: tt.fields.wildChild,
				nType:     tt.fields.nType,
				maxParams: tt.fields.maxParams,
				priority:  tt.fields.priority,
				indices:   tt.fields.indices,
				children:  tt.fields.children,
				handle:    tt.fields.handle,
			}
			n.addRoute(tt.args.path, tt.args.handle)
		})
	}
}

func Test_node_insertChild(t *testing.T) {
	type fields struct {
		path      string
		wildChild bool
		nType     nodeType
		maxParams uint8
		priority  uint32
		indices   string
		children  []*node
		handle    RouteHandler
	}
	type args struct {
		numParams uint8
		path      string
		fullPath  string
		handle    RouteHandler
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				path:      tt.fields.path,
				wildChild: tt.fields.wildChild,
				nType:     tt.fields.nType,
				maxParams: tt.fields.maxParams,
				priority:  tt.fields.priority,
				indices:   tt.fields.indices,
				children:  tt.fields.children,
				handle:    tt.fields.handle,
			}
			n.insertChild(tt.args.numParams, tt.args.path, tt.args.fullPath, tt.args.handle)
		})
	}
}

func Test_node_getValue(t *testing.T) {
	type fields struct {
		path      string
		wildChild bool
		nType     nodeType
		maxParams uint8
		priority  uint32
		indices   string
		children  []*node
		handle    RouteHandler
	}
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantHandle RouteHandler
		wantP      Params
		wantTsr    bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				path:      tt.fields.path,
				wildChild: tt.fields.wildChild,
				nType:     tt.fields.nType,
				maxParams: tt.fields.maxParams,
				priority:  tt.fields.priority,
				indices:   tt.fields.indices,
				children:  tt.fields.children,
				handle:    tt.fields.handle,
			}
			gotHandle, gotP, gotTsr := n.getValue(tt.args.path)
			if !reflect.DeepEqual(gotHandle, tt.wantHandle) {
				t.Errorf("node.getValue() gotHandle = %v, want %v", gotHandle, tt.wantHandle)
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("node.getValue() gotP = %v, want %v", gotP, tt.wantP)
			}
			if gotTsr != tt.wantTsr {
				t.Errorf("node.getValue() gotTsr = %v, want %v", gotTsr, tt.wantTsr)
			}
		})
	}
}

func Test_node_findCaseInsensitivePath(t *testing.T) {
	type fields struct {
		path      string
		wildChild bool
		nType     nodeType
		maxParams uint8
		priority  uint32
		indices   string
		children  []*node
		handle    RouteHandler
	}
	type args struct {
		path             string
		fixTrailingSlash bool
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantCiPath []byte
		wantFound  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				path:      tt.fields.path,
				wildChild: tt.fields.wildChild,
				nType:     tt.fields.nType,
				maxParams: tt.fields.maxParams,
				priority:  tt.fields.priority,
				indices:   tt.fields.indices,
				children:  tt.fields.children,
				handle:    tt.fields.handle,
			}
			gotCiPath, gotFound := n.findCaseInsensitivePath(tt.args.path, tt.args.fixTrailingSlash)
			if !reflect.DeepEqual(gotCiPath, tt.wantCiPath) {
				t.Errorf("node.findCaseInsensitivePath() gotCiPath = %v, want %v", gotCiPath, tt.wantCiPath)
			}
			if gotFound != tt.wantFound {
				t.Errorf("node.findCaseInsensitivePath() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}

func Test_shiftNRuneBytes(t *testing.T) {
	type args struct {
		rb [4]byte
		n  int
	}
	tests := []struct {
		name string
		args args
		want [4]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftNRuneBytes(tt.args.rb, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shiftNRuneBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_node_findCaseInsensitivePathRec(t *testing.T) {
	type fields struct {
		path      string
		wildChild bool
		nType     nodeType
		maxParams uint8
		priority  uint32
		indices   string
		children  []*node
		handle    RouteHandler
	}
	type args struct {
		path             string
		ciPath           []byte
		rb               [4]byte
		fixTrailingSlash bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []byte
		want1  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &node{
				path:      tt.fields.path,
				wildChild: tt.fields.wildChild,
				nType:     tt.fields.nType,
				maxParams: tt.fields.maxParams,
				priority:  tt.fields.priority,
				indices:   tt.fields.indices,
				children:  tt.fields.children,
				handle:    tt.fields.handle,
			}
			got, got1 := n.findCaseInsensitivePathRec(tt.args.path, tt.args.ciPath, tt.args.rb, tt.args.fixTrailingSlash)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("node.findCaseInsensitivePathRec() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("node.findCaseInsensitivePathRec() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
