use std::{assert_eq, cell::RefCell, collections::HashMap, rc::Rc};

type Link = Option<Rc<RefCell<Node>>>;

struct LRUCache {
    map: HashMap<i32, Rc<RefCell<Node>>>,
    head: Rc<RefCell<Node>>,
    tail: Rc<RefCell<Node>>,
    size: i32,
}

struct Node {
    key: i32,
    val: i32,
    prev: Link,
    next: Link,
}

//head->1->2->3->tail
//
impl LRUCache {
    fn new(capacity: i32) -> Self {
        let lru = LRUCache {
            map: HashMap::with_capacity(capacity as usize),
            head: Rc::new(RefCell::new(Node {
                key: 0,
                val: 0,
                prev: None,
                next: None,
            })),
            tail: Rc::new(RefCell::new(Node {
                key: 0,
                val: 0,
                prev: None,
                next: None,
            })),
            size: capacity,
        };
        lru.head.borrow_mut().next = Some(Rc::clone(&lru.tail));
        lru.tail.borrow_mut().prev = Some(Rc::clone(&lru.head));
        lru
    }

    fn get(&mut self, key: i32) -> i32 {
        if !self.map.contains_key(&key) {
            return -1;
        }
        let node = self.map[&key].clone();
        self.remove(node.clone());
        self.add_to_front(node.clone());

        let value = node.borrow().val;
        value
    }

    fn put(&mut self, key: i32, value: i32) {
        if self.map.contains_key(&key) {
            let node = &self.map[&key];

            node.borrow_mut().val = value;
            self.remove(node.clone());
            self.add_to_front(node.clone());
            return;
        }
        if self.map.len() == self.size as usize {
            let to_remove = self.tail.as_ref().borrow().prev.as_ref().unwrap().clone();
            self.remove(to_remove.clone());
            self.map.remove(&to_remove.borrow().key);
        }
        let new_node = Rc::new(RefCell::new(Node {
            key,
            val: value,
            prev: None,
            next: None,
        }));
        self.add_to_front(new_node.clone());
        self.map.insert(key, new_node);
    }

    fn remove(&self, node: Rc<RefCell<Node>>) {
        let prev = node.borrow_mut().prev.take().unwrap();
        let next = node.borrow_mut().next.take().unwrap();
        prev.borrow_mut().next = Some(Rc::clone(&next));
        next.borrow_mut().prev = Some(Rc::clone(&prev));
    }

    fn add_to_front(&self, node: Rc<RefCell<Node>>) {
        let next = self.head.borrow_mut().next.take();
        self.head.borrow_mut().next = Some(Rc::clone(&node));
        next.as_ref().unwrap().borrow_mut().prev = Some(Rc::clone(&node));
        node.borrow_mut().prev = Some(Rc::clone(&self.head));
        node.borrow_mut().next = Some(Rc::clone(&next.unwrap()));
    }
}
#[test]
fn test_lru() {
    let mut lru = LRUCache::new(2);
    lru.put(1, 1);
    lru.put(2, 2);
    let val = lru.get(1);
    assert_eq!(val, 1);
    lru.put(3, 3);
    let val = lru.get(2);
    assert_eq!(val, -1);
    lru.put(4, 4);
    let val = lru.get(1);
    assert_eq!(val, -1);
    let val = lru.get(3);
    assert_eq!(val, 3);
    let val = lru.get(4);
    assert_eq!(val, 4);
}
