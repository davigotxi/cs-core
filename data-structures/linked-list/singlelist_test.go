package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var list *SingleList

func TestSize(t *testing.T) {

	//<group setup code>

	t.Run("Empty list", func(t *testing.T) {
		//when
		list = New()
		res := list.Size()

		//then
		assert.Equal(t, 0, res, "should be 0")
	})

	t.Run("InsertFirst 1 element", func(t *testing.T) {
		//given
		list = New()
		list.InsertFirst(1)

		//when
		res := list.Size()

		//then
		assert.Equal(t, 1, res, "should be 1")
	})

	t.Run("InsertFirst 2 elements", func(t *testing.T) {
		//given
		list = New()
		list.InsertFirst(123)
		list.InsertFirst(456)

		//when
		res := list.Size()

		//then
		assert.Equal(t, 2, res, "should be 2")
	})

	t.Run("InsertLast 1 element", func(t *testing.T) {
		//given
		list = New()
		list.InsertLast(1)

		//when
		res := list.Size()

		//then
		assert.Equal(t, 1, res, "should be 1")
	})

	t.Run("InsertLast 2 elements", func(t *testing.T) {
		//given
		list = New()
		list.InsertLast(123)
		list.InsertLast(456)

		//when
		res := list.Size()

		//then
		assert.Equal(t, 2, res, "should be 2")
	})

	t.Run("Insert first and last elements", func(t *testing.T) {
		//given
		list = New()
		list.InsertFirst(123)
		list.InsertLast(456)

		//when
		res := list.Size()

		//then
		assert.Equal(t, 2, res, "should be 2")
	})

	t.Run("Insert last and first elements", func(t *testing.T) {
		//given
		list = New()
		list.InsertLast(123)
		list.InsertFirst(456)

		//when
		res := list.Size()

		//then
		assert.Equal(t, 2, res, "should be 2")
	})
	// <group tear-down code>
}

func TestFirst(t *testing.T) {

	t.Run("Empty List", func(t *testing.T) {
		list = New()

		_, err := list.First()

		assert.EqualError(t, err, "Single List is empty")
	})

	t.Run("List with 1 element", func(t *testing.T) {
		list = New()
		list.InsertFirst(1)

		res, err := list.First()

		assert.Nil(t, err)
		assert.Equal(t, res, 1)
	})
}

func TestTraverse(t *testing.T) {
	t.Run("Empty List", func(t *testing.T) {
		list = New()

		res := list.Traverse()

		assert.Equal(t, "[]", res)
	})

	t.Run("InsertFirst 1 element", func(t *testing.T) {
		//given
		list = New()
		list.InsertFirst(1)

		//when
		res := list.Traverse()

		//then
		assert.Equal(t, "[1]", res)
	})

	t.Run("InsertFirst 2 elements", func(t *testing.T) {
		//given
		list = New()
		list.InsertFirst(123)
		list.InsertFirst(456)

		//when
		res := list.Traverse()

		//then
		assert.Equal(t, "[456, 123]", res)
	})

	t.Run("InsertLast 1 element", func(t *testing.T) {
		//given
		list = New()
		list.InsertLast(1)

		//when
		res := list.Traverse()

		//then
		assert.Equal(t, "[1]", res)
	})

	t.Run("InsertLast 2 elements", func(t *testing.T) {
		//given
		list = New()
		list.InsertLast(123)
		list.InsertLast(456)

		//when
		res := list.Traverse()

		//then
		assert.Equal(t, "[123, 456]", res)
	})

	t.Run("Insert first and last elements", func(t *testing.T) {
		//given
		list = New()
		list.InsertFirst(123)
		list.InsertLast(456)

		//when
		res := list.Traverse()

		//then
		assert.Equal(t, "[123, 456]", res)
	})

	t.Run("Insert last and first elements", func(t *testing.T) {
		//given
		list = New()
		list.InsertLast(123)
		list.InsertFirst(456)

		//when
		res := list.Traverse()

		//then
		assert.Equal(t, "[456, 123]", res)
	})
}

func setupTestRemove(t *testing.T) func(t *testing.T) {
	t.Log("setup test remove")
	return func(t *testing.T) {
		t.Log("teardown test remove")
	}
}

func setupSubTestRemove(t *testing.T) func(t *testing.T) {
	t.Log("setup sub test remove")

	// list with 3 elements : [1, 2, 3]
	list = New()
	list.InsertLast(1)
	list.InsertLast(2)
	list.InsertLast(3)

	return func(t *testing.T) {
		t.Log("teardown sub test remove")
	}
}
func TestRemove(t *testing.T) {

	teardownTest := setupTestRemove(t)
	defer teardownTest(t)

	t.Run("Remove first 1 time", func(t *testing.T) {
		teardownSubTest := setupSubTestRemove(t)
		defer teardownSubTest(t)

		//when
		list.RemoveFirst()
		res := list.Traverse()

		//then
		assert.Equal(t, "[2, 3]", res)
		assert.Equal(t, 2, list.Size())
	})

	t.Run("Remove first 2 times", func(t *testing.T) {
		teardownSubTest := setupSubTestRemove(t)
		defer teardownSubTest(t)

		//when
		list.RemoveFirst()
		list.RemoveFirst()
		res := list.Traverse()

		//then
		assert.Equal(t, "[3]", res)
		assert.Equal(t, 1, list.Size())
	})

	t.Run("Remove first 3 times", func(t *testing.T) {
		teardownSubTest := setupSubTestRemove(t)
		defer teardownSubTest(t)

		//when
		list.RemoveFirst()
		list.RemoveFirst()
		list.RemoveFirst()
		res := list.Traverse()

		//then
		assert.Equal(t, "[]", res)
		assert.Equal(t, 0, list.Size())
	})

	t.Run("Remove last 1 time", func(t *testing.T) {
		teardownSubTest := setupSubTestRemove(t)
		defer teardownSubTest(t)

		//when
		list.RemoveLast()
		res := list.Traverse()

		//then
		assert.Equal(t, "[1, 2]", res)
		assert.Equal(t, 2, list.Size())
	})

	t.Run("Remove last 2 times", func(t *testing.T) {
		teardownSubTest := setupSubTestRemove(t)
		defer teardownSubTest(t)

		//when
		list.RemoveLast()
		list.RemoveLast()
		res := list.Traverse()

		//then
		assert.Equal(t, "[1]", res)
		assert.Equal(t, 1, list.Size())
	})
}
