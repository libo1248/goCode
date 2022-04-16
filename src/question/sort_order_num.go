package main

import (
	"sort"
)

/*
订单号排序订单号有如下规则：AA-BB-CCCC，共8位数，AA表示部门，BB表示类型，CCCC表示有序的订单号。例如12010001，12020002。
相同AA-BB的情况下，CCCC不会重复，不同AA或者BB的情况下，可能重复。希望对订单号以CCCC的大小进行从小到大排序。请实现排序函数。
*/

func sortOrders(orders []int) {
	ous := make(orderUnits, 0, len(orders))
	for _, order := range orders {
		ou := orderUnit{
			head: order / 10000,
			tail: order % 10000,
		}
		ous = append(ous, ou)
	}

	sort.Sort(ous)

	for i, ou := range ous {
		orders[i] = ou.head*10000 + ou.tail
	}
}

type orderUnit struct {
	head int
	tail int
}

type orderUnits []orderUnit

func (ous orderUnits) Len() int {
	return len(ous)
}

func (ous orderUnits) Less(i, j int) bool {
	return ous[i].tail < ous[j].tail || (ous[i].tail == ous[j].tail && ous[i].head < ous[j].head)
}

func (ous orderUnits) Swap(i, j int) {
	ous[i].tail, ous[j].tail = ous[j].tail, ous[i].tail
}
