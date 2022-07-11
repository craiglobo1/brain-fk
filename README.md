Brainfuck
===========

## Usage
$ brainfuck [-hb] file [-m]
	-h --help	show a help message.
    -b --build  build into a executable [output path]
    -m --memory [no of bytes]
$


## Instruction Set 

dp <- data pointer
ip <- instruction pointer

Brainfuck interpreter written in go.
\>	Increment the dp (to point to the next cell to the right).
<	Decrement the dp (to point to the next cell to the left).
+	Increment (increase by one) the byte at the dp.
-	Decrement (decrease by one) the byte at the dp.
.	Output the byte at the dp.
,	Accept one byte of input, storing its value in the byte at the dp.
[	If the byte at the dp is zero, jump ip after the matching ] command.
]	If the byte at the dp is nonzero jump ip after the matching [ command.