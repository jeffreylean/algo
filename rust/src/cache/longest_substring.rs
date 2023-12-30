use std::collections::HashMap;

#[allow(dead_code, unused)]
pub fn longest_sub_str(s: String) -> i32 {
    let mut map = HashMap::<u8, i32>::new();
    let mut start: i32 = 0;
    let mut max_len: i32 = 0;

    for (end, val) in s.bytes().enumerate() {
        if let Some(i) = map.insert(val, end as i32 + 1) {
            start = start.max(i)
        }
        max_len = max_len.max(end as i32 - start)
    }
    max_len
}
