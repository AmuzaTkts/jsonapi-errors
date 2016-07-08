// Errors package implements the JSON API v1.0 format for errors.
// Ref: http://jsonapi.org/format/#errors
package errors

import "fmt"

// Source contain references to the source of the error
type Source struct {
	Pointer   string `json:"pointer,omitempty"`
	Parameter string `json:"parameter,omitempty"`
}

// Link should lead to further details about this error
type Link struct {
	About string `json:"about,omitempty"`
}

// Error provides additional information about problems encountered while
// performing an operation.
type Error struct {
	Code   int                    `json:"code,string,omitempty"`
	Debug  map[string]interface{} `json:"debug,omitempty"`
	Detail string                 `json:"detail"`
	ID     string                 `json:"id,omitempty"`
	Links  *Link                  `json:"links,omitempty"`
	Meta   map[string]interface{} `json:"meta,omitempty"`
	Source *Source                `json:"source,omitempty"`
	Status int                    `json:"status,string"`
	Title  string                 `json:"title,omitempty"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("%d: %s", e.Status, e.Detail)
}

func (e *Error) SetAboutLink(link string) {
	l := &Link{link}
	e.Links = l
}

func (e *Error) AddSourceNode() *Source {
	s := &Source{}
	e.Source = s

	return s
}

func (e *Error) SetPointer(pointer string) *Error {
	if e.Source == nil {
		e.AddSourceNode()
	}

	e.Source.Pointer = pointer

	return e
}

func (e *Error) SetParameter(parameter string) *Error {
	if e.Source == nil {
		e.AddSourceNode()
	}

	e.Source.Parameter = parameter

	return e
}

func NewError(status int, detail string) *Error {
	return &Error{
		Status: status,
		Detail: detail,
	}
}

type Bag struct {
	Errors []*Error `json:"errors"`
	Status int      `json:"status,string,omitempty"`
}

func (b *Bag) Add(err *Error) *Error {
	b.Errors = append(b.Errors, err)

	if len(b.Errors) == 1 {
		b.Status = err.Status

		return err
	}

	if b.Status != err.Status {
		if statusClass := GetErrorClass(b.Status); statusClass == GetErrorClass(err.Status) {
			b.Status = statusClass
		} else {
			b.Status = 400
		}
	}

	return err
}

func (b *Bag) AddError(status int, detail string) *Error {
	return b.Add(NewError(status, detail))
}

// NewBag returns a new Bag
func NewBag() *Bag {
	return &Bag{}
}

// NewBagWithError is a shorthand to `errors.NewBag().Add(status, detail)`
func NewBagWithError(status int, detail string) *Bag {
	bag := &Bag{}
	bag.AddError(status, detail)

	return bag
}

// GetErrorClass returns 400 or 500 depending of the error code passed.
// If the error code is less than 400 or greater or equal to 600 it returns 0.
func GetErrorClass(err int) int {
	if err < 400 || err >= 600 {
		return 0
	}

	if err >= 400 && err < 500 {
		return 400
	}

	return 500
}
