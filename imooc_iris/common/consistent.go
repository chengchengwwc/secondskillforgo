package common

import (
	"errors"
	"hash/crc32"
	"sort"
	"strconv"
	"sync"
)

//声明西的切片类型
type units []uint32

func (x units) Len() int {
	return len(x)
}

func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

var ErrEmpty = errors.New("Hash No value")

//创建和保存hash一致性
type Consistent struct {
	//hash环
	circle map[uint32]string
	//已经排序的hash节点
	sortedHashes units
	//添加虚拟节点
	VirtualNode int
	//lock 分布式锁
	sync.RWMutex
}

func NewConsistent() *Consistent {
	return &Consistent{
		circle:      make(map[uint32]string),
		VirtualNode: 20,
	}
}

//自动生成key
func (c *Consistent) generateKey(element string, index int) string {
	return element + strconv.Itoa(index)
}

//获取hash位置
func (c *Consistent) hashkey(key string) uint32 {
	if len(key) < 64 {
		//声明一个数组长度为64
		var srcatch [64]byte
		copy(srcatch[:], key)
		//使用IEEE 多项式返回数据的CRC-32校验和
		return crc32.ChecksumIEEE(srcatch[:len(key)])
	}
	return crc32.ChecksumIEEE([]byte(key))
}

//更新排序，方便查找
func (c *Consistent) updateSortedHashes() {
	hashes := c.sortedHashes[:0]
	//判断切片容量是否过大，如果过大则重设置
	if cap(c.sortedHashes)/(c.VirtualNode*4) > len(c.circle) {
		hashes = nil
	}
	for key := range c.circle {
		hashes = append(hashes, key)
	}
	sort.Sort(hashes)
	c.sortedHashes = hashes

}
func (c *Consistent) Add(element string) {
	//加锁
	c.Lock()
	defer c.Unlock()
	c.add(element)
}

func (c *Consistent) add(element string) {
	for i := 0; i < c.VirtualNode; i++ {
		c.circle[c.hashkey(c.generateKey(element, i))] = element
	}
	//
	c.updateSortedHashes()

}

func (c *Consistent) Remove(element string) {
	c.Lock()
	defer c.Unlock()
	c.remove(element)

}

func (c *Consistent) remove(element string) {
	for i := 0; i < c.VirtualNode; i++ {
		delete(c.circle, c.hashkey(c.generateKey(element, i)))
	}
	c.updateSortedHashes()
}

//顺时针查找最近的节点
func (c *Consistent) search(key uint32) int {
	//查找算法
	f := func(x int) bool {
		return c.sortedHashes[x] > key
	}
	// 使用二分查找来搜索指定切片的最小值
	i := sort.Search(len(c.sortedHashes), f)
	if i >= len(c.sortedHashes) {
		i = 0
	}
	return i
}

//根据数据标示，获取最近的hash服务器
func (c *Consistent) Get(name string) (string, error) {
	//添加锁 读写锁
	c.RLock()
	defer c.Unlock()
	if len(c.circle) == 0 {
		return "", ErrEmpty
	}
	//计算hash值
	key := c.hashkey(name)
	i := c.search(key)
	return c.circle[c.sortedHashes[i]], nil

}
