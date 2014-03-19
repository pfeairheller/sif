package gomethius

import (
	"time"
	"fmt"
	"log"
	"encoding/json"
	"os/exec"
	"bufio"
	"strings"
)


type ShellSpout struct {
	Commands []string
	Fields []string
	Output chan []Value
	Emitter chan Values
}

func NewShellSpout(commands[] string, fields []string) (*ShellSpout){
	out := new(ShellSpout)
	out.Commands = commands
	out.Fields = fields
	out.Output = make(chan []Value)
	return out
}

func (s *ShellSpout) Open(conf map[string]string, context *TopologyContext, emitter chan Values) {
	s.Emitter = emitter
	go s.ExecuteShell()
}


func (s *ShellSpout) ExecuteShell() {
	cmd := exec.Command(s.Commands[0], s.Commands[1:]...)
	go showErr(cmd)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		line = strings.TrimSuffix(line, "\n")
		if line == "Started" {
			break
		} else {
			fmt.Println(line)
		}
	}

	decoder := json.NewDecoder(stdout)
	m := []interface{}{}
	for {
		if err := decoder.Decode(&m); err != nil {
			log.Println(err);
			break
		}

		vals := []Value{}
		for _, val := range m {
			switch t := val.(type) {
			default:
				fmt.Printf("unexpected type %T", t)       // %T prints whatever type t has
			case bool:
				//vals = append(vals, NewBoolValue(val))
			case string:
				vals = append(vals, NewStringValue(val.(string)))
			case float64:
				vals = append(vals, NewFloat32Value(float32(val.(float64))))
			}
		}

		s.Output <- vals
	}

}

func showErr(cmd *exec.Cmd) {
	stdout, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			return
		} else {
			log.Println(line)
		}
	}
}

func (s *ShellSpout) NextTuple() {
	var vals [] Value
	select {
	case vals = <- s.Output:
		vals = append(vals, NewTimeValue(time.Now()))
		s.Emitter <- *NewValues(vals...)
	default:
		time.Sleep(time.Millisecond * 50)
	}
}

func (s *ShellSpout) DeclareOutputFields() *Fields {
	return NewFields(s.Fields...)
}

