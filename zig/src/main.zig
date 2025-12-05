const std = @import("std");
const expect = std.testing.expect;

const constant: i32 = 5;
var variable: u32 = 5000;

pub fn main() void {
    std.debug.print("Hello, {s}!\n", .{"World"});
}

test "always succeeds" {
    try expect(true);
}
