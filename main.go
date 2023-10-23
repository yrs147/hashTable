package main

import (
	"fmt"

)

//Size of hast table array
const ArraySize = 7

//HashTable Structure
type HashTable struct{
  array [ArraySize]*bucket

}

//bucket strucutre is a linked list in each slot of the hash table array
type bucket struct {
  head *bucketNode
}

//bucketNode is a linkedlist that holds the key

type bucketNode struct {
  key string
  next *bucketNode
}

//Insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string){
  index := hash(key)
  h.array[index].insert(key)
}
  

//Search will take in a key and return true if the key is stored in the hash table
func (h *HashTable) Search(key string) bool{
  index := hash(key)
  return h.array[index].search(key)
}
  

//Delete wil take in a key and delete it from the hash table
func (h *HashTable) Delete(key string){
  index := hash(key)
  h.array[index].delete(key)
}
  

//insert will take in a key , create a node with the key and inset the node in the bucket
func (b *bucket) insert (k string){
  newNode := &bucketNode{key: k}
  if b.head == nil {
    b.head = newNode
  } else {
    currentNode := b.head
    for currentNode != nil {
      if currentNode.key == k {
        fmt.Println("Already Exists")
        return
      }
      if currentNode.next == nil {
        break
      }
      currentNode = currentNode.next
    }
    currentNode.next = newNode
  }
}

//search will take in a key and return true if the bucket has that key
func (b *bucket) search(k string) bool{
  currentNode := b.head
  for currentNode != nil{
    if currentNode.key == k {
      return true
    }
    currentNode = currentNode.next
  }
  return false
}
//delete will take in a ket and delete the node form the bucket 
func (b *bucket) delete(k string){
  
  if b.head.key == k {
    b.head = b.head.next
    return
  }
  previousNode := b.head
  for previousNode.next != nil{

    if previousNode.next.key == k {
      //delete
      previousNode.next = previousNode.next.next
    }
    previousNode = previousNode.next
  }
}

//hash function

func hash(key string) int{
  sum := 0
  for _,v := range key{
    sum+=int(v)
  }

  return sum % ArraySize

}

//Init will create a bucket in each slot of the hash table
func Init() *HashTable{
  result := &HashTable{}

  for i := range result.array{
    result.array[i] = &bucket{}
  }
  return result
}

func main() {
  
  hashTable := Init()
  list := []string{
    "ERIC",
    "AMAN",
    "LESLY",
    "KYLE",
    "KEITH",
    "SNOOP",
    "ADAM",
  }

  for _, v := range list {
    hashTable.Insert(v)
  }
  hashTable.Delete("SNOOP")

  fmt.Println("KEITH", hashTable.Search("KEITH"))
  fmt.Println("LESLY", hashTable.Search("LESLY"))
  fmt.Println("GILGAMESH", hashTable.Search("GILGAMESH"))
  
  // testHashtable := Init()
  // fmt.Println(testHashtable)
  // fmt.Println(hash("RANDY"))
  // testBucket := &bucket{}
  // testBucket.insert("RANDY")
  // testBucket.delete("RANDY")
  // fmt.Println(testBucket.search("RANDY"))
  // fmt.Println(testBucket.search("ERIC"))
}

