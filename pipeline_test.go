package pipeline

import (
	"testing"
	"fmt"
	"encoding/json"
)

func Test_PipeLine(t *testing.T) {
	p := NewPipeLine()
	p.AddStage(New())
	p.Run(nil)
	b, err := json.Marshal(p)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(b))
	newpipe := NewPipeLine()
	if err := json.Unmarshal(b, newpipe); err != nil {
		t.Error(err)
	}
	fmt.Println(newpipe)
}