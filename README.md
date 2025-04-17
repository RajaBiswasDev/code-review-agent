# 🤖 Code Review Agent

Welcome to `code-review-agent` — a sample repository created to **showcase the capabilities of an AI agent that performs automated code reviews**.

This repo contains intentionally poor Go code <<test_before.go>> that serves as a testbed for demonstrating how an AI agent can:

- Review source code written in Go
- Provide structured feedback
- Suggest actionable improvements
- Help developers learn best practices through guided review

---

## 🧪 What's Inside?

### 🔹 `main.go`
This contains the AI Agent code which can scan files in a location, read the content and review the file (Go / Scala / Java) and based on the review can accommodate the updates in the file automatically.

### 🔹 `test_before.go`
This file is purposefully written with common Go anti-patterns and bad practices such as:

- Poor naming conventions
- Lack of error handling
- Use of global variables
- Hardcoded values
- Mixed concerns (I/O, business logic, etc.)
- No documentation or comments
- Inconsistent formatting
- Inefficient logic

This code serves as a great candidate for AI-assisted review sessions.

### 🔹 `test_after.go`
This file contains the updated code which has been updated by the AI Agent itself; after reviewing the source file <<test_before.go>> and accomodating the revoew comments it has generated.

### 🔹 `sample_console_output.txt`
This contains the conversation done with the AI Agent once its run through ```go run main.go```.

---

## 🎯 Project Goal

The primary objective of this project is to demonstrate the **AI agent’s ability to understand and critique source code** in natural language.

This can help in:

- Teaching junior developers best practices
- Automating code review workflows
- Generating PR review comments
- Enabling conversational code walkthroughs

---

## 🚀 Usage

To run the code (for demo purposes only):

```bash
go run main.go
