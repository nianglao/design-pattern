package main

import "fmt"

const (
	_ int8 = iota
	FILE
	QUEUE
	PATHWAY
)

type PersistenceAttribute struct {
	Type  int8
	Value string
}

type DistrWorkPackage struct {
	Desc string
}

func NewDistrWorkPackage(Type string) *DistrWorkPackage {
	return &DistrWorkPackage{Desc: fmt.Sprintf("Distributed Work Package for: %s", Type)}
}

func (d *DistrWorkPackage) SetFile(f, v string) {
	d.Desc = d.Desc + fmt.Sprintf("\n File(%s): %s", f, v)
}

func (d *DistrWorkPackage) SetQueue(q, v string) {
	d.Desc = d.Desc + fmt.Sprintf("\n Queue(%s): %s", q, v)
}

func (d *DistrWorkPackage) SetPathWay(p, v string) {
	d.Desc = d.Desc + fmt.Sprintf("\n PathWay(%s): %s", p, v)
}

func (d *DistrWorkPackage) GetState() string {
	return d.Desc
}

type Builder interface {
	ConfigureFile(string)
	ConfigureQueue(string)
	ConfigurePathWay(string)
}

type UnixBuilder struct {
	Result *DistrWorkPackage
}

func NewUnixBuilder() *UnixBuilder {
	return &UnixBuilder{Result: NewDistrWorkPackage("Unix")}
}

func (u *UnixBuilder) ConfigureFile(name string) {
	u.Result.SetFile("flatFile", name)
}

func (u *UnixBuilder) ConfigureQueue(queue string) {
	u.Result.SetQueue("FIFO", queue)
}

func (u *UnixBuilder) ConfigurePathWay(pathWay string) {
	u.Result.SetPathWay("thread", pathWay)
}

type VmsBuilder struct {
	Result *DistrWorkPackage
}

func NewVmsBuilder() *VmsBuilder {
	return &VmsBuilder{Result: NewDistrWorkPackage("Vms")}
}

func (v *VmsBuilder) ConfigureFile(name string) {
	v.Result.SetFile("ISAM", name)
}

func (v *VmsBuilder) ConfigureQueue(queue string) {
	v.Result.SetQueue("priority", queue)
}

func (v *VmsBuilder) ConfigurePathWay(pathWay string) {
	v.Result.SetPathWay("LWP", pathWay)
}

type Reader struct {
	builder Builder
}

func (r *Reader) SetBuilder(b Builder) {
	r.builder = b
}

func (r *Reader) Construct(attrs []PersistenceAttribute) {
	for i := 0; i < len(attrs); i++ {
		attr := attrs[i]
		if attr.Type == FILE {
			r.builder.ConfigureFile(attr.Value)
		} else if attr.Type == QUEUE {
			r.builder.ConfigureQueue(attr.Value)
		} else if attr.Type == PATHWAY {
			r.builder.ConfigurePathWay(attr.Value)
		}
	}
}

var Attrs []PersistenceAttribute = []PersistenceAttribute{
	PersistenceAttribute{Type: FILE, Value: "state.dat"},
	PersistenceAttribute{Type: FILE, Value: "config.sys"},
	PersistenceAttribute{Type: QUEUE, Value: "compute"},
	PersistenceAttribute{Type: QUEUE, Value: "log"},
	PersistenceAttribute{Type: PATHWAY, Value: "authentication"},
	PersistenceAttribute{Type: PATHWAY, Value: "error processing"},
}

func main() {
	ub := NewUnixBuilder()
	vb := NewVmsBuilder()
	var r Reader

	r.SetBuilder(ub)
	r.Construct(Attrs)
	fmt.Println(ub.Result.GetState())

	r.SetBuilder(vb)
	r.Construct(Attrs)
	fmt.Println(vb.Result.GetState())
}
