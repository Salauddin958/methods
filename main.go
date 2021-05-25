package main

import "fmt"

type Age int

func (age Age) LargerThan(a Age) bool {
	return age > a
}

func (age *Age) IsNil() bool {
	return age == nil
}

func (age *Age) Increase() {
	*age++
}

// Receiver of custom defined function type
type FilterFunc func(in int) bool

func (ff FilterFunc) Filter(in int) bool {
	return ff(in)
}

// Receiver of custom defined map type.
type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}

func (ss StringSet) Add(key string) {
	ss[key] = struct{}{}
}

func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// Receiver of custom defined struct type.
type Book struct {
	pages int
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(p int) {
	b.pages = p
}

type Books []Book

func (books Books) Modify() {
	books[0].pages = 100
	books = append(books, Book{400})
}

func (books *Books) ModifyBooks() {
	*books = append(*books, Book{789})
	(*books)[0].pages = 500
}

func main() {
	var book Book

	fmt.Printf("%T\n", book.Pages)
	fmt.Printf("%T\n", (&book).SetPages)

	fmt.Printf("%T \n", (&book).Pages)
	// Calling the three methods.
	(&book).SetPages(123)
	book.SetPages(123)
	fmt.Println(book.Pages())
	fmt.Println((&book).Pages())

	_ = StringSet(nil).Has  // will not panic
	_ = ((*Age)(nil)).IsNil // will not panic
	_ = ((*Age)(nil))       // will not panic

	_ = (StringSet(nil)).Has("key") // will not panic
	_ = ((*Age)(nil)).IsNil()       // will not panic

	var b1 = Books{{123}, {456}}
	b1.Modify()
	fmt.Println(b1)
	b1.ModifyBooks()
	fmt.Println(b1)

}
