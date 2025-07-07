# 🔓 Cracking the C3 Cyber CTF 

![image](https://github.com/user-attachments/assets/5cae576d-6bbf-459f-ba90-c82df9c6e568)

trap
So, picture this: The Estonian Cybercrime Unit (C3) needs a new Chief. Their ingenious recruitment method? Embed a job application within a Capture The Flag (CTF) challenge! 🕵️‍♂️ Why? Because clearly, the ultimate test of leadership is deciphering obscure references from The Hacker Manifesto in Estonian.

Get ready to answer your future subordinate's burning question: "Boss, how do I decode 0:70 again?" Let's dive in!

# 🎯 Challenge Overview

The challenge site offers zero hints—just a cryptic string, and a sly nod to the C3 unit. Subtle, right? It’s your job to connect the dots from mysterious "Number:Number" patterns to something tangible. And those funky tokens like y:x? Clearly, top-tier steganography…

>TheConscienceofaHacker
>
> `1:36 0:25 0:6 12:5 0:15 0:3 0:14 y:x 0:33 0:2 0:25 0:3 0:30 0:15 0:28 0:5 0:28 1:36 0:15 4:52 0:5 0:6 1:36 0:15 0:6 0:25 0:3 0:5 0:10 0:5 0:33 0:15 0:28 0:5 0:10 0:14 0:33 0:2 0:25 0:21 0:14 4:52 0:28 0:5 0:5 1:36 0:5 0:28 0:5 0:25 0:30 0:14 0:28 0:5 0:5 0:1 0:21 0:15 0:30 0:30 0:14 0:3 0:14 0:21 0:14 y:x 1:36 0:14 0:1 0:21 0:25 0:21 0:5 0:5 0:6 0:25 1:12 0:25 0:28 0:14 0:33 0:14 0:30 0:21 0:15 0:28 0:5 1:12 0:2 0:25 0:28 0:3 0:14 0:3 0:15 0:28 0:5 0:30 0:14 0:4 0:5 0:1 0:21 0:15 0:28 0:5 1:36 0:14 0:28 0:3 0:25`
>
> *Source: https://web.archive.org/web/20250313111130/https://cyber.politsei.ee/m6istatus/*


# 🔎 Recon & Ideation

Upon first glance, these tokens scream "coordinates." The format y:x looks suspiciously like a pointer—think "book cipher," but instead of a book, we have 70 lines of the classic Hacker's Manifesto at our disposal.

When you spot "1:36," doesn’t "section:index" just leap out at you? This is a classic CTF staple, and a nod to anyone familiar with coordinate-based ciphers!

# 🛠️ How I Approached the Challenge
## 1️⃣ Deciphering the Tokens
Tokens are structured in a straightforward part:index pattern.

0:25 – Take the 25th character, no fuss.

1:36 – But wait, add 183 because obviously, that’s fun. 🧮
### 🦄 Special Cases
0:70 → ü (Because why use boring letters?)

0:2 → õ (Estonian: home of extra vowels!)



## 2️⃣ Extracting the Manifesto
Pulled the text from GitHub, processed lines 17-72, and stripped the whitespace—resulting in one seamless character stream. Here’s the code snippet:

<details>
  <summary>Golang Code</summary>
  
```go
// Fetches and processes a specific range of lines from the Hacker Manifesto.
func fetchManifestoContent() (string, error) {
	resp, err := http.Get("https://raw.githubusercontent.com/greyscalepress/manifestos/refs/heads/master/content/manifestos/1986-hacker-manifesto.txt")
	if err != nil {
		return "", fmt.Errorf("failed to fetch manifesto: %w", err)
	}
	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)
	var processedContent strings.Builder
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		if lineNum < 17 {
			continue
		}
		if lineNum > 72 {
			break
		}
		line := scanner.Text()
		for _, r := range line {
			if !unicode.IsSpace(r) {
				processedContent.WriteRune(r)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("failed during scanning: %w", err)
	}
	return processedContent.String(), nil
}
```

</details>

## 3️⃣ Decoding Logic
<details>
  <summary>Golang Code</summary>

```go
func decodeCoordinate(coords, sourceText string) string {
	switch coords {
	case "y:x":
		return ". "
	case "0:0":
		return " "
	case "0:70":
		return "ü"
	case "0:2":
		return "õ"
	case "4:52":
		coords = "2:18" // Handle specific case
	}
	parts := strings.Split(coords, ":")
	p, err1 := strconv.Atoi(parts[0])
	i, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return "?" // Invalid coordinate values
	}
	if p > 0 {
		i += 183
	}
	runes := []rune(sourceText)
	if i < 0 || i >= len(runes) {
		return "#" // Coordinate out of bounds
	}
	return string(runes[i])
}
```
</details>


# 🎉 The Big Reveal!
Running the code produces... 🥁 the flag!


`kirduta voitluseskuberkuritegevusegavoidatseekeseilaseendullatada kandideeritisavaldusetoistatuselahendusekasti`

Naturally, the organizers left a challenge: decipher and smooth out the output into a beautiful, meaningful sentence:

`võitluses küberkuritegevusega võidab see kes ei lase end üllatada. kirjuta kandideerimisavalduse mõistatuse lahendusekasti`


# 🏆 Conclusion: Leadership Material?
Forget your CISSP, OSCP, or years managing security teams. If you navigated this labyrinth of obscure offsets and Estonian diacritics hidden within a hacker classic, congratulations! 🎉 The C3 leadership awaits your meticulously decoded application.

Honestly, this challenge was... something. A unique blend of hacker nostalgia and a cipher so straightforward it barely qualifies as crypto. It's the perfect entry point – a confidence-booster designed for aspiring cyber legends who are perhaps still mastering grep.

Kudos to the C3 organizers for this unforgettable, character-building introduction to their unit! It certainly sets a tone. 😉

# ⭐️ Who Should Try This?

- No IT background? No problem!
- IT-conversion students? Yes, please!
- CTF rookies or veterans—it’s a party for all.
- Anyone testing if C3 is a good fit for them.
- Be sure to thank the organizers—this challenge is a rite of passage for Estonia's cyber-defenders! 🛡️✨

