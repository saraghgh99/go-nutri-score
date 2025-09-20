# Go Nutri-Score Calculator 🍏

A simple command-line application built in Go to calculate a product's official Nutri-Score.

## Project Purpose

**This project was created as a learning exercise to deepen my understanding of:**

* **Go's folder structure and package management:** Organizing code into a clear, maintainable structure.
* **Custom types and methods:** Using Go's type system to create clear, reusable code (e.g., EnergyKJ, SugarGram, etc.).
* **Algorithm implementation:** Translating a real-world scoring algorithm into code.

As part of my journey to become a backend developer, this project marks the beginning of my hands-on experience with Go.

## Project Structure

The project is organized into a clean folder structure to demonstrate good Go practices:

* **main.go:** The application's entry point, which runs the main calculation logic.
* **nutriscore/:** A reusable Go package containing all the core logic, separated into two files:
    * **types.go:** Defines all custom types and constants for the scoring system.
    * **calculator.go:** Contains all the functions and methods for calculating the score.
