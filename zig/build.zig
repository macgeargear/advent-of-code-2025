const std = @import("std");

pub fn build(b: *std.Build) void {
    const target = b.standardTargetOptions(.{});
    const optimize = b.standardOptimizeOption(.{});

    // Day 01
    const day01 = b.addExecutable(.{
        .name = "day01",
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day01.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    b.installArtifact(day01);

    const run_day01_step = b.step("day01", "Run day 01");
    const run_day01_cmd = b.addRunArtifact(day01);
    run_day01_step.dependOn(&run_day01_cmd.step);

    // Day 02
    const day02 = b.addExecutable(.{
        .name = "day02",
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day02.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    b.installArtifact(day02);

    const run_day02_step = b.step("day02", "Run day 02");
    const run_day02_cmd = b.addRunArtifact(day02);
    run_day02_step.dependOn(&run_day02_cmd.step);

    // Day 03
    const day03 = b.addExecutable(.{
        .name = "day03",
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day03.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    b.installArtifact(day03);

    const run_day03_step = b.step("day03", "Run day 03");
    const run_day03_cmd = b.addRunArtifact(day03);
    run_day03_step.dependOn(&run_day03_cmd.step);

    // Day 04
    const day04 = b.addExecutable(.{
        .name = "day04",
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day04.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    b.installArtifact(day04);

    const run_day04_step = b.step("day04", "Run day 04");
    const run_day04_cmd = b.addRunArtifact(day04);
    run_day04_step.dependOn(&run_day04_cmd.step);

    // Day 05
    const day05 = b.addExecutable(.{
        .name = "day05",
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day05.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    b.installArtifact(day05);

    const run_day05_step = b.step("day05", "Run day 05");
    const run_day05_cmd = b.addRunArtifact(day05);
    run_day05_step.dependOn(&run_day05_cmd.step);

    // Default Main
    const exe = b.addExecutable(.{
        .name = "zig",
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/main.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    b.installArtifact(exe);

    const run_step = b.step("run", "Run the default app");
    const run_cmd = b.addRunArtifact(exe);
    run_step.dependOn(&run_cmd.step);
    run_cmd.step.dependOn(b.getInstallStep());

    if (b.args) |args| {
        run_cmd.addArgs(args);
    }

    // Tests
    const test_step = b.step("test", "Run all tests");

    // Day 01 tests
    const day01_tests = b.addTest(.{
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day01.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    const run_day01_tests = b.addRunArtifact(day01_tests);
    test_step.dependOn(&run_day01_tests.step);

    // Day 02 tests
    const day02_tests = b.addTest(.{
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day02.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    const run_day02_tests = b.addRunArtifact(day02_tests);
    test_step.dependOn(&run_day02_tests.step);

    // Day 03 tests
    const day03_tests = b.addTest(.{
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day03.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    const run_day03_tests = b.addRunArtifact(day03_tests);
    test_step.dependOn(&run_day03_tests.step);

    // Day 04 tests
    const day04_tests = b.addTest(.{
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day04.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    const run_day04_tests = b.addRunArtifact(day04_tests);
    test_step.dependOn(&run_day04_tests.step);

    // Day 05 tests
    const day05_tests = b.addTest(.{
        .root_module = b.createModule(.{
            .root_source_file = b.path("src/day05.zig"),
            .target = target,
            .optimize = optimize,
        }),
    });
    const run_day05_tests = b.addRunArtifact(day05_tests);
    test_step.dependOn(&run_day05_tests.step);

    // Main executable tests
    const exe_tests = b.addTest(.{
        .root_module = exe.root_module,
    });
    const run_exe_tests = b.addRunArtifact(exe_tests);
    test_step.dependOn(&run_exe_tests.step);
}
