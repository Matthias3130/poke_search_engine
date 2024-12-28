# Pokemon Search Engine

<video width="1280" height="720" controls>
  <source src="assets/Pokemon Search Engine Demo.mp4" type="video/mp4">
  Your browser does not support the video tag.
</video>

## Table of Contents
- [About](#about)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)
- [Reflection](#reflection)
- [TODOs](#TODOs)

---

## About
I created a simple search engine for a database of Pokémon sourced from PokéAPI. This is a personal project I developed after completing my Web Development course, so some of the approaches reflect what I learned during the course. Instead of using AJAX, I switched to using the Fetch API.  

I used Python to extract all the necessary information into a CSV file, which I then import into a SQLite database.


One important thing to note is that the filters are union-based.  

## Features
- Includes all current Pokémon, from #1 to #1025.  
- Offers filter options to categorize Pokémon by type, generation, or groups such as starters and legendaries.  
- Dynamically retrieves the list of Pokémon without requiring a page refresh.  
- Uses SQLite, with no server or login required.  
- Allows reversing the order of the list from ascending to descending. 

## Installation
1. Clone the repository:  
   ```bash
   git clone https://github.com/username/repository.git
2. Install <a href="https://go.dev/doc/install">GO</a>

3. Install <a href="https://www.sqlite.org/download.html">SQLite</a>

4. Download Dependencies
   ```bash
   go mod tidy

5. Install Air
   ```bash
   go install github.com/cosmtrek/air@latest

## Usage
* Once all the dependencies are installed, you can start the program by typing the command below into your terminal. This will build your project and provide a link to view it on localhost.
    ```bash
    air

## Contributing
You are welcome to clone the project and work with what is given, but I will not be accepting any pull requests that are not my own for this project.

## License
All Pokémon data is obtained from PokéAPI, so please refer to their license policy.

## Reflection
I could have done more, but I'm satisfied with what I have now. The style is minimal, but I quite like it.

## TODOs
If I return to this project, these are the features I plan to add:
1. Dark mode
2. Combine all the stats into one column
3. Add logic to find the max stat and bold it
4. Turn the stats into a colored bar gradient to visually represent the stat's quality
5. Add more tags
6. Find a way to optimize image loading
7. Enable image enlargement on hover
8. Display the definition of abilities when hovered over
9. Might host it
