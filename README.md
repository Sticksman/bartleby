# longform-text-compiler
Scrivener like text compiler for long form content.

## Goal
Scrivener is one of the best long form text editors because it understands that the author of a long-form piece isn't trying to create a single large file; rather they are trying to create a patchwork of many different pieces to then stitch together. In that sense, Scrivener is a product with a close to perfect use case.

However there are some minor issues with Scrivener. For one thing the text editor is fully featured and very Microsoft Word-like. Further, due to its own internal naming system, Scrivener is not compatible with other editors, not easily anyway. Finally, the compilation step, while fully featured and supporting multiple different formats, is at its heart very simple.

The goal of this project is to create a small binary that is capable of compiling multiple long-form text files into a single document based off a supplied configuration.

## Prereqs
- Go 1.8.1
