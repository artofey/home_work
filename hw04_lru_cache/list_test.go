package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, MakeNativeIntList(l, false))
		require.Equal(t, []int{50, 30, 10, 40, 60, 80, 70}, MakeNativeIntList(l, true))
	})

	t.Run("push front", func(t *testing.T) {
		l := NewList()

		item20 := l.PushFront(20)
		require.Equal(t, 1, l.Len())
		require.Equal(t, item20, l.Front())
		require.Equal(t, item20, l.Back())
		require.Equal(t, []int{20}, MakeNativeIntList(l, false))
		require.Equal(t, []int{20}, MakeNativeIntList(l, true))

		item40 := l.PushFront(40)
		require.Equal(t, 2, l.Len())
		require.Equal(t, item40, l.Front())
		require.Equal(t, item20, l.Back())
		require.Equal(t, []int{40, 20}, MakeNativeIntList(l, false))
		require.Equal(t, []int{20, 40}, MakeNativeIntList(l, true))

		item10 := l.PushFront(10)
		require.Equal(t, 3, l.Len())
		require.Equal(t, item10, l.Front())
		require.Equal(t, item20, l.Back())
		require.NotEqual(t, item40, l.Back())
		require.NotEqual(t, item40, l.Front())
		require.Equal(t, []int{10, 40, 20}, MakeNativeIntList(l, false))
		require.Equal(t, []int{20, 40, 10}, MakeNativeIntList(l, true))
	})

	t.Run("push back", func(t *testing.T) {
		l := NewList()

		item20 := l.PushBack(20) // 20
		require.Equal(t, 1, l.Len())
		require.Equal(t, item20, l.Front())
		require.Equal(t, item20, l.Back())
		require.Empty(t, item20.Prev)
		require.Empty(t, item20.Next)
		require.Equal(t, []int{20}, MakeNativeIntList(l, false))
		require.Equal(t, []int{20}, MakeNativeIntList(l, true))

		item40 := l.PushBack(40) // 20, 40
		require.Equal(t, 2, l.Len())
		require.Equal(t, item20, l.Front())
		require.Equal(t, item40, l.Back())
		require.Equal(t, item20.Next, item40)
		require.Equal(t, item40.Prev, item20)

		require.Equal(t, []int{20, 40}, MakeNativeIntList(l, false))
		require.Equal(t, []int{40, 20}, MakeNativeIntList(l, true))

		item10 := l.PushBack(10) // 20, 40, 10
		require.Equal(t, 3, l.Len())
		require.Equal(t, item20, l.Front())
		require.Equal(t, item10, l.Back())
		require.NotEqual(t, item40, l.Back())
		require.NotEqual(t, item40, l.Front())
		require.Equal(t, item20.Next, item40)
		require.Equal(t, item40.Prev, item20)
		require.Equal(t, item40.Next, item10)
		require.Equal(t, item10.Prev, item40)
		require.Equal(t, []int{20, 40, 10}, MakeNativeIntList(l, false))
		require.Equal(t, []int{10, 40, 20}, MakeNativeIntList(l, true))
	})
}

func TestListMoveToFront(t *testing.T) {
	t.Run("move to front", func(t *testing.T) {
		l := NewList()

		item20 := l.PushFront(20) // 20
		l.MoveToFront(item20)     // 20
		require.Equal(t, 1, l.Len())
		require.Equal(t, item20, l.Front())
		require.Equal(t, item20, l.Back())
		require.Equal(t, []int{20}, MakeNativeIntList(l, false))
		require.Equal(t, []int{20}, MakeNativeIntList(l, true))

		item30 := l.PushFront(30) // 30, 20
		l.MoveToFront(item30)     // 30, 20
		require.Equal(t, 2, l.Len())
		require.Equal(t, item30, l.Front())
		require.Equal(t, item20, l.Back())
		require.Equal(t, []int{30, 20}, MakeNativeIntList(l, false))
		require.Equal(t, []int{20, 30}, MakeNativeIntList(l, true))

		l.MoveToFront(item20) // 20, 30
		require.Equal(t, 2, l.Len())
		require.Equal(t, item20.Value, l.Front().Value)
		require.Equal(t, item30.Value, l.Back().Value)
		require.Equal(t, item20.Next, item30)
		require.Equal(t, item30.Prev, item20)
		require.Equal(t, []int{20, 30}, MakeNativeIntList(l, false))
		require.Equal(t, []int{30, 20}, MakeNativeIntList(l, true))

		item40 := l.PushBack(40) // 20, 30, 40
		l.MoveToFront(item40)    // 40, 20, 30
		require.Equal(t, 3, l.Len())
		require.Equal(t, item40.Value, l.Front().Value)
		require.Equal(t, item30.Value, l.Back().Value)

		require.Equal(t, item40.Next, item20)
		require.Equal(t, item20.Prev, item40)
		require.Equal(t, item20.Next, item30)
		require.Equal(t, item30.Prev, item20)
		require.Equal(t, []int{40, 20, 30}, MakeNativeIntList(l, false))
		require.Equal(t, []int{30, 20, 40}, MakeNativeIntList(l, true))

		l.MoveToFront(item20) // 20, 40, 30
		require.Equal(t, 3, l.Len())
		require.Equal(t, item20.Value, l.Front().Value)
		require.Equal(t, item30.Value, l.Back().Value)

		require.Equal(t, item20.Next, item40)
		require.Equal(t, item40.Prev, item20)
		require.Equal(t, item40.Next, item30)
		require.Equal(t, item30.Prev, item40)
		require.Equal(t, []int{20, 40, 30}, MakeNativeIntList(l, false))
		require.Equal(t, []int{30, 40, 20}, MakeNativeIntList(l, true))
	})
}

func TestListRemove(t *testing.T) {
	t.Run("remove", func(t *testing.T) {
		l := NewList()
		item50 := l.PushBack(50)
		l.Remove(item50)
		require.Nil(t, l.Back())
		require.Nil(t, l.Front())
		require.Equal(t, 0, l.Len())
	})

	t.Run("remove first item", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.PushBack(20)
		l.PushBack(30)
		l.PushBack(40)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 4, l.Len())
		l.Remove(l.Front())
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 3, l.Len())
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{20, 30, 40}, elems)
	})
	t.Run("remove last item", func(t *testing.T) {
		l := NewList()
		l.PushBack(10)
		l.PushBack(20)
		l.PushBack(30)
		l.PushBack(40)
		require.Equal(t, 40, l.Back().Value)
		require.Equal(t, 4, l.Len())
		l.Remove(l.Back())
		require.Equal(t, 30, l.Back().Value)
		require.Equal(t, 3, l.Len())
		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{10, 20, 30}, elems)
	})
}

func MakeNativeIntList(l List, revert bool) []int {
	elems := make([]int, 0, l.Len())
	if revert == false {
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
	} else {
		for i := l.Back(); i != nil; i = i.Prev {
			elems = append(elems, i.Value.(int))
		}
	}
	return elems
}
