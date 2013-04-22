// simple matrix package
package matrix

type Matrix interface {
	Size() (m, n int)
	Row() (m int)
	Col() (n int)
}
