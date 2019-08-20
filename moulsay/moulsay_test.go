package moulsay

import "fmt"

func ExampleSay() {
	out, _ := Say("hello world!", 72)
	fmt.Println(out)
	// Output:
	//       ++
	//       ++++
	//        ++++
	//      ++++++++++
	//     +++       |
	//     ++         |
	//     +  -==   ==|
	//    (   <*>   <*>
	//     |           |
	//     |         __|
	//     |      +++
	//      \      =+
	//       \      +                             hello world!
	//       |\++++++
	//       |  ++++      ||//
	//   ____|   |____   _||/__
	//  /     ---     \  \|  |||
	// /  _ _  _     / \   \ /
	// | / / //_//_//  |   | |
}
