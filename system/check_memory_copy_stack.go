package main

const SIZE = 1024

func main() {
	s := "HELLO"
	deepCopy(&s, 0, [SIZE]int{})
}

func deepCopy(s *string, c int, r [SIZE]int) {
	println(c, s, *s)
	//fmt.Printf("%d %p %v\n", c, s, *s)
	/*
		String will be in a heap if you are using fmt.Printf
		its happens because fmt creates a new struct
		and String cannot be saved in a local stack memory
		at the end of the file you can see runtime debug memory trace
	*/
	c++

	if c == 16 {
		return
	}

	deepCopy(s, c, r)
}

/*
0xc00006c1e0: 0xc00006c1e0 - 1
0xc00006c1e0: 0xc00006c1e0 - 2
runtime: newstack sp=0xc0000a7e80 stack=[0xc0000a6000, 0xc0000ae000]
        morebuf={pc:0x48edec sp:0xc0000a7e90 lr:0x0}
        sched={pc:0x48ee01 sp:0xc0000a7e88 lr:0x0 ctxt:0x0}
stackalloc 65536
  allocated 0xc0000b2000
copystack gp=0xc000000180 [0xc0000a6000 0xc0000a7e88 0xc0000ae000] -> [0xc0000b2000 0xc0000bbe88 0xc0000c2000]/65536
stackfree 0xc0000a6000 32768
stack grow done
0xc00006c1e0: 0xc00006c1e0 - 3
0xc00006c1e0: 0xc00006c1e0 - 4
0xc00006c1e0: 0xc00006c1e0 - 5
0xc00006c1e0: 0xc00006c1e0 - 6
runtime: newstack sp=0xc0000b3d20 stack=[0xc0000b2000, 0xc0000c2000]
        morebuf={pc:0x48edec sp:0xc0000b3d30 lr:0x0}
        sched={pc:0x48ee01 sp:0xc0000b3d28 lr:0x0 ctxt:0x0}
stackalloc 131072
  allocated 0xc0000c2000
copystack gp=0xc000000180 [0xc0000b2000 0xc0000b3d28 0xc0000c2000] -> [0xc0000c2000 0xc0000d3d28 0xc0000e2000]/131072
stackfree 0xc0000b2000 65536
stack grow done
0xc00006c1e0: 0xc00006c1e0 - 7
0xc00006c1e0: 0xc00006c1e0 - 8
0xc00006c1e0: 0xc00006c1e0 - 9
0xc00006c1e0: 0xc00006c1e0 - 10
0xc00006c1e0: 0xc00006c1e0 - 11
0xc00006c1e0: 0xc00006c1e0 - 12
0xc00006c1e0: 0xc00006c1e0 - 13
0xc00006c1e0: 0xc00006c1e0 - 14
runtime: newstack sp=0xc0000c3a60 stack=[0xc0000c2000, 0xc0000e2000]
        morebuf={pc:0x48edec sp:0xc0000c3a70 lr:0x0}
        sched={pc:0x48ee01 sp:0xc0000c3a68 lr:0x0 ctxt:0x0}
stackalloc 262144
  allocated 0xc0000e2000
copystack gp=0xc000000180 [0xc0000c2000 0xc0000c3a68 0xc0000e2000] -> [0xc0000e2000 0xc000103a68 0xc000122000]/262144
stackfree 0xc0000c2000 131072
stack grow done
0xc00006c1e0: 0xc00006c1e0 - 15
0xc00006c1e0: 0xc00006c1e0 - 16

*/
