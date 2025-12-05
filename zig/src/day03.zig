const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Read input
    const input = @embedFile("inputs/day03.txt");

    // Solve
    const result1 = try solvePart1(allocator, input);
    const result2 = try solvePart2(allocator, input);

    std.debug.print("Part 1: {d}\n", .{result1});
    std.debug.print("Part 2: {d}\n", .{result2});
}

fn solvePart1(allocator: std.mem.Allocator, input: []const u8) !i32 {
    _ = allocator;
    _ = input;
    return 0;
}

fn solvePart2(allocator: std.mem.Allocator, input: []const u8) !i32 {
    _ = allocator;
    _ = input;
    return 0;
}

test "day03 part1" {
    const input = "test input";
    const result = try solvePart1(std.testing.allocator, input);
    try std.testing.expectEqual(@as(i32, 0), result);
}

test "day03 part2" {
    const input = "test input";
    const result = try solvePart2(std.testing.allocator, input);
    try std.testing.expectEqual(@as(i32, 0), result);
}
