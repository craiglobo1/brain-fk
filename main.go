package main
	
import (
    "fmt"
    "os"
	"io/ioutil"
	"bufio"
	"strings"
)

const MAX_OPS = 100000
const MAX_MEM  = 30000

var debug bool = false
var max_ip = 0

const (
	PTR_INC int = iota
	PTR_DEC
	VAL_INC
	VAL_DEC
	OUTPUT_CHAR
	INPUT_CHAR
	CJUMP
	JUMP
)

type operand struct {
	typ int
	jmp_val int
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

func print_help()  {
	fmt.Println("$ brainfuck [-hb] file [-m]")
	fmt.Println("	-h --help	show a help message.")
	fmt.Println("	-b --build  build into a executable [output path]")
	fmt.Println("	-m --memory [no of bytes]")
}

func main()  {
	// take cmd args for file and flags
	filepath := ""
	if len(os.Args) < 2 {
		fmt.Println("Error: no cmd args")
		print_help()
		os.Exit(1)
	} else {
		for i := 1; i < len(os.Args);i++ {
			if os.Args[i][0] == '-'{
				if strings.Contains(os.Args[i], "h"){
					print_help()
					os.Exit(0)
				} else if strings.Contains(os.Args[i], "b"){
					if len(os.Args) < i+1{
						fmt.Print("Error: no size given for -b flag")
						os.Exit(1)
					}
					// add building option detail
				}
			} else {
				filepath = os.Args[i]
			}
		}
	}
	// parse file and generate 
	program := parse_and_gen_program(filepath)
	// interpret the program or generate assembly and build the program
	interpret_program(program)
}

func parse_and_gen_program(filepath string) [MAX_OPS]operand {
	file_text, err := ioutil.ReadFile(filepath)
	check(err)

	var program [MAX_OPS]operand

	var stack [100]int
	var stack_ip int = 0

	ip := 0
	for _, char := range file_text {
		switch char {
		case '>':
			program[ip] = operand{PTR_INC, 0}
			ip++
		case '<':
			program[ip] = operand{PTR_DEC, 0}
			ip++
		case '+':
			program[ip] = operand{VAL_INC, 0}
			ip++
		case '-':
			program[ip] = operand{VAL_DEC, 0}
			ip++
		case '.':
			program[ip] = operand{OUTPUT_CHAR, 0}
			ip++
		case ',':
			program[ip] = operand{INPUT_CHAR, 0}
			ip++
		case '[':
			program[ip] = operand{CJUMP, 0}
			stack[stack_ip] = ip
			stack_ip++
			ip++
		case ']':
			program[ip] = operand{JUMP, stack[stack_ip-1]}
			program[stack[stack_ip-1]].jmp_val = ip + 1
			stack_ip--
			ip++
		}
	}
	max_ip = ip-1
	return program
}

func interpret_program(program [MAX_OPS]operand)  {
	reader := bufio.NewReader(os.Stdin)
	var memory [MAX_MEM]int8
	dp := 0
	i := 0
	for i <= max_ip{
		cur_op := program[i]
		switch cur_op.typ {
			case PTR_INC:
				dp++
				i++	
			case PTR_DEC:
				dp--
				i++
			case VAL_INC:
				memory[dp] += 1
				i++
			case VAL_DEC:
				memory[dp] -= 1
				i++
			case OUTPUT_CHAR:
				fmt.Print(string(int(memory[dp])))
				i++
			case INPUT_CHAR:
				in_char, _, _ := reader.ReadRune()
				memory[dp] = int8(in_char)
				i++
			case CJUMP:
				if memory[dp] == 0{
					i = cur_op.jmp_val
				} else {
					i++
				}
			case JUMP:
				i = cur_op.jmp_val
			default:
				fmt.Println("Error: invalid operand type")
				os.Exit(1)
		}
	}
}