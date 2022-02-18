package lt341

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (this NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (this NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (this *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (this NestedInteger) GetList() []*NestedInteger {}
 */

 type NestedIterator struct {
    NestedList []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
    return &NestedIterator{NestedList: nestedList}
}

// 返回第一个元素，如果第一个元素不是单个整数，则取出列表，插入列表起始处，直到第一个元素是单个整数为止
func (this *NestedIterator) Next() int {
    // 因为HasNext判断过了是否还存在数据，所以这里不再判断列表长度
    firstInt := this.NestedList[0].GetInteger()
    this.NestedList = this.NestedList[1:]

    return firstInt
}

// 判断NestedList长度
func (this *NestedIterator) HasNext() bool {
    for len(this.NestedList) > 0 && !(this.NestedList[0]).IsInteger() {
        list := this.NestedList[0].GetList()
        this.NestedList = append(list, this.NestedList[1:]...)
    }

    if len(this.NestedList) == 0 {
        return false
    }

    return true
}