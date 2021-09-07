// SPDX-FileCopyrightText: 2021 None
//
// SPDX-License-Identifier: CC0-1.0

module.exports = {
  parserPreset: "conventional-changelog-conventionalcommits",
  rules: {
    "body-leading-blank": [2, "always"],
    "body-max-line-length": [2, "always", 120],
    "footer-leading-blank": [2, "always"],
    "footer-max-line-length": [2, "always", 120],
    "header-max-length": [2, "always", 120],
    "scope-case": [2, "always", "lower-case"],
    "scope-enum": [
      2,
      "always",
      [
        "codegen",
        "db",
        "sql",
        "logger",
        "config",
        "api",
        "ci",
        "docs",

        "auth",
        "auth/federation",

        "batch",
        "chat",
        "emote",
        "mediaproxy",
        "profile",
        "stream",
        "sync",
        "voice",
      ],
    ],
    "subject-case": [
      2,
      "never",
      ["sentence-case", "start-case", "pascal-case", "upper-case"],
    ],
    "scope-empty": [2, "never"],
    "subject-empty": [2, "never"],
    "subject-full-stop": [2, "never", "."],
    "type-case": [2, "always", "lower-case"],
    "type-empty": [2, "never"],
    "type-enum": [
      2,
      "always",
      [
        "build",
        "feat",
        "fix",
        "perf",
        "refactor",
        "revert",
        "style",
        "test",
      ],
    ],
  },
};
