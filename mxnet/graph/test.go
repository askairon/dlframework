package graph

import (
	"encoding/json"
	"strings"
	"regexp"
	"strconv"
	"fmt"

	"github.com/awalterschulze/gographviz"
	"github.com/pkg/errors"
	"gopkg.in/fatih/set.v0"
)

type NodeEntry struct {
	NodeId  int  `json:"node_id"`
	Index   int  `json:"index"`
	Version *int `json:"version,omitempty"`
}

type Node struct {
	Op                  string            `json:"op"`
	Param               map[string]string `json:"param"`
	Name                string            `json:"name"`
	Inputs              []NodeEntry       `json:"inputs"`
	BackwardSourceID    int               `json:"backward_source_id"`
	ControlDependencies []int             `json:"control_deps,omitempty"`
}

type Graph struct {
	Nodes          []Node                 `json:"nodes"`
	ArgNodes       []int                  `json:"arg_nodes"`
	NodeRowPointer []int                  `json:"node_row_ptr,omitempty"`
	Heads          []NodeEntry            `json:"heads"`
	Attributes     map[string]interface{} `json:"attrs,omitempty"`
}

func (e *NodeEntry) UnmarshalJSON(b []byte) error {
	var s []int
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}
	if len(s) < 2 {
		return errors.New("expecting a node entry length >= 2")
	}
	e.NodeId = s[0]
	e.Index = s[1]
	if len(s) == 3 {
		*e.Version = s[2]
	}
	return nil
}

func (e *NodeEntry) MarshalJSON() ([]byte, error) {
	s := []int{
		e.NodeId,
		e.Index,
	}
	if e.Version != nil {
		s = append(s, *e.Version)
	}

	return json.Marshal(s)
}
func str2tuple(s string) ([]string){

    re :=regexp.MustCompile("[0-9]+")
    return re.FindallString(s,-1)
}
// color map
var fillcolors = []string{
	"#8dd3c7",
	"#fb8072",
	"#ffffb3",
	"#bebada",
	"#80b1d3",
	"#fdb462",
	"#b3de69",
	"#fccde5",
}
var edgecolors = []string{
	"#245b51",
	"#941305",
	"#999900",
	"#3b3564",
	"#275372",
	"#975102",
	"#597d1c",
	"#90094e",
}

func (g *Graph) ToDotGraph() (*gographviz.Escape, error) {

	makeDefaultAttributes := func() map[string]string {
		return map[string]string{
			"shape":     "box",
			"fixedsize": "true",
			"width":     "1.3",
			"height":    "0.8034",
			"style":     "filled",
		}
	}

	isLikeWeight := func(name string) bool {
		if strings.HasSuffix(name, "_weight") {
			return true
		}
		if strings.HasSuffix(name, "_bias") {
			return true
		}
		if strings.HasSuffix(name, "_beta") ||
			strings.HasSuffix(name, "_gamma") ||
			strings.HasSuffix(name, "_moving_var") ||
			strings.HasSuffix(name, "_moving_mean") {
			return true
		}
		return false
	}

	hideWeights := true    // TODO: should be an option
	drawShape := true      // TODO: should be an option
	graphName := "mxnet"   // TODO: should be an option
	layout := "horizontal" // TODO: should be an option

	dg := gographviz.NewEscape()
	dg.SetName(graphName)
	dg.SetDir(true)

	dg.AddAttr(graphName, "nodesep", "1")
	dg.AddAttr(graphName, "ranksep", "1.5 equally")

	switch layout {
	case "vertical":
		dg.AddAttr(graphName, "rankdir", "TB")
	case "horizontal":
		dg.AddAttr(graphName, "rankdir", "RL")
	default:
	}

	hiddenNodes := set.NewNonTS()

	// make nodes
	for _, node := range g.Nodes {
		op := node.Op
		name := node.Name
		attrs := makeDefaultAttributes()
		label := op

		switch op {
			case "null":
				if isLikeWeight(name) {
					if hideWeights {
						hiddenNodes.Add(name)
						continue
					}
				}
				attrs["shape"] = "oval"
				attrs["fillcolor"] = fillcolors[0]
				label = name
			case "Convolution":
			//...
				
				if val, ok := node.Param["stride"]; ok {
		 			stride_info:=str2tuple(val)
		 			label=fmt.Sprintf("Convolution\n%s/%s, %s",strings.Join(str2tuple(node.Param["kernal"]),"x"),strings.Join(stride_info,"x"),node.Param["num_filter"])
				}else{
					label=fmt.Sprintf("Convolution\n%s/%s, %s",strings.Join(str2tuple(node.Param["kernal"]),"x"),"1",node.Param["num_filter"])
				}

				attrs["fillcolor"]=fillcolors[1]

			case "FullyConnected":
				//...
				label=fmt.Sprintf("FullyConnected\n%s",node.Param["num_hidden"])
				attrs["fillcolor"]=fillcolors[1]
			case "BatchNorm":
	            attrs["fillcolor"] = fillcolors[3]
	        case "Activation","LeakyReLU":
	        	label = fmt.Sprintf("%s\n%s",op, node.Param["act_type"])
	            attrs["fillcolor"] = fillcolors[2]
	        case "Pooling":
	        	if val, ok := node.Param["stride"]; ok {
		 			stride_info:=str2tuple(val)
		 			label=fmt.Sprintf("Pooling\n%s, %s/%s",node.Param["pool_type"],strings.Join(str2tuple(node.Param["kernal"]),"x"),strings.Join(stride_info,"x"))
				}else{
					label=fmt.Sprintf("Pooling\n%s, %s/%s",node.Param["pool_type"],strings.Join(str2tuple(node.Param["kernal"]),"x"),"1")
				}
	            attrs["fillcolor"] = fillcolors[4]
	        case "Concat" , "Flatten" , "Reshape":
	            attrs["fillcolor"] = fillcolors[5]
	        case "Softmax":
	            attrs["fillcolor"] = fillcolors[6]
	        default:
	            attrs["fillcolor"] = fillcolors[7]
	            if op == "Custom"{
	                label = node.Param["op_type"]
	            }
			attrs["label"] = label
			dg.AddNode("G", name, attrs)
		}
	}

	// make edges
	for _, node := range g.Nodes {
		op := node.Op
		name := node.Name
		if op == "null" {
			continue
		}
		inputs := node.Inputs
		for _, item := range inputs {
			inputNode := g.Nodes[item.NodeId]
			inputName := inputNode.Name
			if hiddenNodes.Has(inputName) {
				continue
			}
			attrs := map[string]string{
				"dir":       "back",
				"arrowtail": "open",
			}
			if drawShape {
				// ...
				_ = inputNode
				_ = attrs
				key :=inputName
				if inputNode.Op!="null"{
					key="_output"
					if val, ok := inputNode.Param["num_outputs"]; ok {
						key+=strconv.Itoa(strconv.Atoi(inputNode.Param["num_outputs"])-1)
						inputNode.Param["num_outputs"]=strconv.Itoa(strconv.Atoi(inputNode.Param["num_outputs"])-1)
					}

				}
				//shape = shape_dict[key][1:]
                //label = "x".join([str(x) for x in shape])
                //attrs["label"] = label
			}
			dg.AddEdge(name, inputName, true, attrs)
		}

	}

	return dg, nil
}



    
    
