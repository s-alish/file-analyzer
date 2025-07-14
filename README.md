# CLI File Analyzer

A simple command-line tool written in Go that analyzes a text file and reports basic statistics.

## 📌 Features

- Counts:
  - Total lines
  - Total words
  - Total characters (including whitespace and newlines)
  - Unique words (case-insensitive, punctuation-trimmed)
- Handles file-not-found errors
- Outputs results in a clean formatted table

## 🛠️ Usage

```bash
go run main.go <filename>
