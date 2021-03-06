# Golang Data Struct

### 2. 链表

链表是一种链式存储的线性表，所有元素的`内存地址不一定是连续的`；

#### 2.1 单向链表

> ⭐ 单链表的结构和构造方法：**
>
> ```go
> // define node
> type Node struct {
>     Val interface{}
>     Next *Node
> }
> 
> // define singal linked List
> type SLL struct {
>     Size int
>     first *Node
> }
> 
> // SSL constructed function
> func newSSL() *SSL {
>     list := new(SLL)
>     list.first = &Node{
>          Val : nil,
>          Next: nil
>      }
>      return list
> }
> ```
>
> **⭐ 单链表的常用方法：**
>
> 



#### 2.2 双向链表

>   **⭐ 单链表的结构和构造方法：**
>
>   ```go
>   type DNode struct{
>       Val interface{}
>       prev *DNode
>       Next *DNode
>   }
>   
>   type DLL struct {
>       Size int
>       first *DNode
>       last *DNode
>   }
>   
>   func newDLL() *DLL {
>       list = new(DLL)
>       list.first = &DNode{
>           Val: nil,
>           Prev: nil,
>           Next: nil,
>       }
>       return list
>   }
>   ```
>
>   



*   双向链表 VS 动态数组

    动态数组：开辟、销毁内存空间的次数相对较少，但可能造成内存空间浪费（可以通过缩容解决）

    双向链表：开辟、销毁内存空间的次数相对较多，但不会造成内存空间的浪费

*   数据结构的选取

    如果频繁在==尾部==进行添加、删除操作，==动态数组、双向链表==均可选择
    如果频繁在==头部==进行添加、删除操作，建议选择使用==双向链表==
    如果有频繁的（在==任意位置==）添加、删除操作，建议选择使用==双向链表==
    如果有频繁的==查询操作==（随机访问操作），建议选择使用==动态数组==

*   有了双向链表，单向链表是否就没有任何用处了？

    并非如此，在哈希表的设计中就用到了单链表



### 4. 树

   节点、根节点、父节点、子节点、兄弟节点

   <img src="D:\Development\WorkSpace\GoProjects\src\code.golang.com\Datastruct\photo\tree.png" alt="tree" style="zoom:67%;" />

   >   节点的度（degree）：子树的个数；
   >
   >   树的度：Max(节点的度)；
   >
   >   叶子节点：没有任何子树（度为0）；
   >
   >   非叶子节点： 度不为0的节点；
   
   >   层数（level）：根节点在第一层，其子节点在第二层；
   >
   >   节点的深度（depth）：往上数包含自己在内的节点数；
   >
   >   节点的高度（height）：往下数包含自己在内的节点数；
   >
   >   树的深度：max（节点的深度）；
   >
   >   树的高度：max（节点的高度）；
   >
   >   ==树的深度 ====== =树的高度==

*   二叉树

    -   每个节点的度最大为2（最多只有两颗子树）

    -   左右子树是有序的；（==有序树==）
    -   即使某节点只有一棵子树，也要区分左右子树；
    -   非空二叉树的第i层，最多有2^(i - 1)个节点；
    -   在高度为h的二叉树上最多有2^h^ - 1；
    -   对于任何一棵非空二叉树，如果叶子节点的个数为n~0~，度为2的节点个数为n~2~， 则有n~0~ = n~2~ + 1；
    -   二叉树的边数T = n~1~ + 2 * n~2~ = n - 1 = n~0~ + n~1~ + n~2~  - 1 ----> n~0~ = n~2~ + 1

    

*   真二叉树

    -   所有节点要么为0要么为2；

    

*   满二叉树，

    -   最后一层节点度均为0， 其他节点度均为2；

    -   第i层节点数：2 ^(i - 1)， 叶子节点数 2^(h - 1)，总节点数 n = 2^h^ -1, 高度 h = log~2~^(n + 1)
    -   在同样高度的二叉树中，满二叉树的叶子节点数量最多、总结点数最多；
    -   满二叉树一定是真二叉树，真二叉树不一定是满二叉树；

    

*   完全二叉树

    -   叶子节点只会出现在最后2层，最后一层叶子节点靠左对齐；
    -   完全二叉树从==根节点==到==倒数第二层==是一棵满二叉树；
    -   满二叉树一定是完全二叉树，完全二叉树不一定是满二叉树；

    >   性质：
    >
    >   1.  度为1的节点只有左子树；
    >   2.  度为1的节点要么只有1个要么只有0个；
    >   3.  同样节点数量，完全二叉树高度最小；
    >   4.  至少有2^(h - 1),  至多有 2^h^ - 1(满二叉树);
    >   5.  h = floor(log~2~n) + 1;
    >   6.  完全二叉树总结点数为n，
    >       -   叶子节点 = ceiling(n / 2) = (n+1)  >> 1
    >       -   非叶子节点数 = ceiling((n-1)) / 2 = n >> 1




#### 4.1 BST 二叉搜索树

>   二叉搜索树是二叉树的一种：
>
>   -   任意一节点的值都小于其左子树所有节点的值；
>   -   任意一节点的值都大于其右子树所有节点的值；
>   -   其左右子树也是BST
>
>   注意：
>
>   -   BST储存的元素必须具备可比较性；
>   -   如果是自定义类型需要指定比较方式；
>   -   不允许为null；
>
>   ```golang
>   /*
>   	接口设计：
>   	Size() int
>   	IsEMPTY() bool
>   	Clear()
>   	Add(e interface{})
>   	Remove(index int)
>   	Contains(E interface{}) bool
>   	!!!! 对于当前的BST没有索引的概念 !!!!
>   	*/
>   	func (b *BST) Add(e interface{}) error {
>   		if e == nil {
>   			return errors.New("element must be not nil")
>   		}
>   		
>   		if b.root == nil {
>   			b.root = &node{
>   				val: e,
>   			}
>   			b.size++
>   			return nill
>   		}
>   		
>   		parentNode, tmpNode, ret := new(node), b.root, 0
>   		for tmpNode != nil {
>   			ret = b.Comparate(e, tmpNode.val)
>   			parentNode = tmpNode
>   			if ret > 0 {
>   				tmpNode = tmpNode.right
>   			} else if ret < 0 {
>   				tmpNode = tmpNode.left
>   			} else {
>   			tmpNode.val = e
>   			return nil
>   			}
>   		}
>   		
>   		newNode = &node{
>               val: e,
>               parent: parentNode
>   		}
>   		
>   		if ret > 0 {
>               parentNode.right = newNode
>   		} else {
>   			parentNode.right = newNode
>   		}
>   		b.size++
>   		return nil
>   	}
>   ```
>
>   
