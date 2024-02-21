struct MyCircularQueue {
    queue: Vec<i32>,
    head: usize,
    tail: usize,
}

/**
 * `&self` means the method takes an immutable reference.
 * If you need a mutable reference, change it to `&mut self` instead.
 */
impl MyCircularQueue {
    fn new(k: i32) -> Self {
        MyCircularQueue {
            queue: Vec::with_capacity(k as usize),
            head: 0,
            tail: 0,
        }
    }

    fn en_queue(&mut self, value: i32) -> bool {
        if self.queue.len() == self.queue.capacity() {
            return false;
        }

        if self.queue.is_empty() {
            self.queue.insert(0, value);
            return true;
        }

        self.tail += 1 % self.queue.capacity();
        self.queue.insert(self.tail, value);

        true
    }

    fn de_queue(&mut self) -> bool {
        if self.queue.is_empty() {
            return false;
        }

        let _ = self.queue.remove(self.head);

        if self.queue.is_empty() {
            self.head = 0;
            self.tail = 0;

            return true;
        }
        self.tail -= 1;

        true
    }

    fn front(&self) -> i32 {
        if self.queue.is_empty() {
            return -1;
        }
        self.queue[self.head]
    }

    fn rear(&self) -> i32 {
        if self.queue.is_empty() {
            return -1;
        }
        self.queue[self.tail]
    }

    fn is_empty(&self) -> bool {
        self.queue.is_empty()
    }

    fn is_full(&self) -> bool {
        self.queue.len() == self.queue.capacity()
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_queue() {
        let mut queue = MyCircularQueue::new(8);
        assert!(queue.en_queue(3));
        dbg!(queue.head);
        dbg!(queue.tail);
        assert!(queue.en_queue(9));
        dbg!(queue.head);
        dbg!(queue.tail);
        assert!(queue.en_queue(5));
        dbg!(queue.head);
        dbg!(queue.tail);
        assert!(queue.en_queue(0));
        dbg!(queue.head);
        dbg!(queue.tail);
        assert!(queue.de_queue());
        dbg!(queue.head);
        dbg!(queue.tail);
        assert!(queue.de_queue());
        dbg!(queue.head);
        dbg!(queue.tail);
        assert!(!queue.is_empty());
        assert!(!queue.is_empty());
        assert_eq!(queue.rear(), 0);
        assert_eq!(queue.rear(), 0);
        assert!(queue.de_queue());
    }
}
