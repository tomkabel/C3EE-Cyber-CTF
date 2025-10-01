# C3 Cyber CTF Challenge Write-Up

## Introduction

This repository contains a detailed write-up and solution for a Capture The Flag (CTF) challenge. The challenge was presented as a recruitment exercise for the Estonian Cybercrime Unit (C3), where the job application was embedded within the CTF itself.

## Challenge Overview

The challenge provided a cryptic string of "Number:Number" patterns and a reference to the C3 unit, with no other explicit instructions. The objective was to decode this string to reveal the hidden flag. The solution involved recognizing the pattern as a book cipher, using "The Hacker Manifesto" as the key.

### The Ciphertext

The provided ciphertext was:

```
1:36 0:25 0:6 12:5 0:15 0:3 0:14 y:x 0:33 0:2 0:25 0:3 0:30 0:15 0:28 0:5 0:28 1:36 0:15 4:52 0:5 0:6 1:36 0:15 0:6 0:25 0:3 0:5 0:10 0:5 0:33 0:15 0:28 0:5 0:10 0:14 0:33 0:2 0:25 0:21 0:14 4:52 0:28 0:5 0:5 1:36 0:5 0:28 0:5 0:25 0:30 0:14 0:28 0:5 0:5 0:1 0:21 0:15 0:30 0:30 0:14 0:3 0:14 0:21 0:14 y:x 1:36 0:14 0:1 0:21 0:25 0:21 0:5 0:5 0:6 0:25 1:12 0:25 0:28 0:14 0:33 0:14 0:30 0:21 0:15 0:28 0:5 1:12 0:2 0:25 0:28 0:3 0:14 0:3 0:15 0:28 0:5 0:30 0:14 0:4 0:5 0:1 0:21 0:15 0:28 0:5 1:36 0:14 0:28 0:3 0:25
```

*Source: [C3 Cyber CTF Challenge](https://web.archive.org/web/20250313111130/https://cyber.politsei.ee/m6istatus/)*

## Methodology

The tokens in the ciphertext were identified as coordinates in the format `y:x`, suggesting a book cipher. "The Hacker Manifesto" was used as the source text for decoding.

### 1. Decoding the Tokens

The tokens were interpreted as `part:index` pointers to characters in the manifesto. Specific rules were applied for certain tokens:

### 2. Processing the Source Text

Lines 17-72 of "The Hacker Manifesto" were extracted and all whitespace was removed to create a continuous stream of characters for decoding.

### 3. Solution Implementation

A Go program was developed to automate the decoding process. The code fetches the manifesto, processes the text, and decodes the ciphertext according to the rules described above.

## The Solution

The decoded flag is:

```
kirduta voitluseskuberkuritegevusegavoidatseekeseilaseendullatada kandideeritisavaldusetoistatuselahendusekasti
```

After formatting and adding appropriate spacing and diacritics, the message reads:

```
võitluses küberkuritegevusega võidab see kes ei lase end üllatada. kirjuta kandideerimisavalduse mõistatuse lahendusekasti
```

This translates to:

> In the fight against cybercrime, the one who does not let themselves be surprised wins. Write the solution to the puzzle in the application box.

## Conclusion

This CTF challenge was a creative and engaging way to test problem-solving skills. It required a combination of pattern recognition, knowledge of classic ciphers, and programming to arrive at the solution. This repository serves as a comprehensive walk-through of the solution process.
