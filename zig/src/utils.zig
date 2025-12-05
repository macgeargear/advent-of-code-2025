const std = @import("std");

/// Split input by lines
pub fn splitLines(allocator: std.mem.Allocator, input: []const u8) ![][]const u8 {
    var lines = std.ArrayList([]const u8).init(allocator);
    var iter = std.mem.splitScalar(u8, input, '\n');
    while (iter.next()) |line| {
        if (line.len > 0) {
            try lines.append(line);
        }
    }
    return lines.toOwnedSlice();
}

/// Parse integer from string
pub fn parseInt(comptime T: type, s: []const u8) !T {
    return try std.fmt.parseInt(T, s, 10);
}

/// Read all lines from input, trimming whitespace
pub fn readLines(allocator: std.mem.Allocator, input: []const u8) ![][]const u8 {
    var lines = std.ArrayList([]const u8).init(allocator);
    var iter = std.mem.splitScalar(u8, input, '\n');
    while (iter.next()) |line| {
        const trimmed = std.mem.trim(u8, line, &std.ascii.whitespace);
        if (trimmed.len > 0) {
            try lines.append(trimmed);
        }
    }
    return lines.toOwnedSlice();
}

test "splitLines" {
    const input = "line1\nline2\nline3\n";
    const lines = try splitLines(std.testing.allocator, input);
    defer std.testing.allocator.free(lines);

    try std.testing.expectEqual(@as(usize, 3), lines.len);
    try std.testing.expectEqualStrings("line1", lines[0]);
    try std.testing.expectEqualStrings("line2", lines[1]);
    try std.testing.expectEqualStrings("line3", lines[2]);
}

test "parseInt" {
    try std.testing.expectEqual(@as(i32, 123), try parseInt(i32, "123"));
    try std.testing.expectEqual(@as(i32, -456), try parseInt(i32, "-456"));
}
