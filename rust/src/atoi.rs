fn atoi(input: String) -> i32 {
    // Remove leading whitespace
    let input = input.trim_start();

    let (s, sign) = match input.strip_prefix('-') {
        Some(s) => (s, -1_i32),
        None => (input.strip_prefix('+').unwrap_or(input), 1_i32),
    };
    s.chars()
        .filter_map(|c| c.to_digit(10))
        .map(|v| v as i32)
        .fold(0, |acc, digit| {
            acc.saturating_mul(10).saturating_add(digit) * sign
        })
}

#[cfg(test)]
mod tests {
    use super::*;
    #[test]
    fn test_atoi() {
        let input = "word and 987";
        assert!(atoi(input.to_string()) == 987)
    }
}
