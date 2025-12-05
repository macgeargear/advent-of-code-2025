const std = @import("std");

pub fn main() !void {
    var gpa = std.heap.GeneralPurposeAllocator(.{}){};
    defer _ = gpa.deinit();
    const allocator = gpa.allocator();

    // Read input
    const input = @embedFile("inputs/day05.txt");

    // Solve
    const result1 = try solvePart1(allocator, input);
    const result2 = try solvePart2(allocator, input);

    std.debug.print("Part 1: {d}\n", .{result1});
    std.debug.print("Part 2: {d}\n", .{result2});
}

const Range = struct { l: i64, r: i64 };

fn solvePart1(allocator: std.mem.Allocator, input: []const u8) !i64 {
    var sum: i64 = 0;

    var lines = std.mem.splitScalar(u8, input, '\n');

    var ranges: std.ArrayList(Range) = .empty;
    defer ranges.deinit(allocator);

    var nums: std.ArrayList(i64) = .empty;
    defer nums.deinit(allocator);

    var f = false;
    while (lines.next()) |line| {
        if (std.mem.eql(u8, line, "")) {
            f = true;
            continue;
        }
        if (!f) {
            var parts = std.mem.splitScalar(u8, line, '-');
            const l = try std.fmt.parseInt(i64, parts.next().?, 10);
            const r = try std.fmt.parseInt(i64, parts.next().?, 10);
            try ranges.append(allocator, .{ .l = l, .r = r });
        } else {
            const num = try std.fmt.parseInt(i64, line, 10);
            try nums.append(allocator, num);
        }
    }
    std.mem.sort(Range, ranges.items, {}, struct {
        fn compare(_: void, a: Range, b: Range) bool {
            return a.l < b.l;
        }
    }.compare);

    for (nums.items) |num| {
        for (ranges.items) |range| {
            if (range.l <= num and num <= range.r) {
                sum += 1;
                break;
            }
        }
    }

    return sum;
}

fn solvePart2(allocator: std.mem.Allocator, input: []const u8) !i64 {
    _ = allocator;
    _ = input;
    return 0;
}

test "day05 part1 example" {
    const input = @embedFile("inputs/day05_example.txt");

    const expected = 3;
    const actual = try solvePart1(std.testing.allocator, input);

    try std.testing.expectEqual(expected, actual);
}
test "day05 part2 example" {}
