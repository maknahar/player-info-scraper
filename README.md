# player-info-scraper
A utility to retrieve the football players info

## Environment variables
- LIMIT (Integer) : LIMIT specify the maximum API 
    calls that should be made to scrape the team info.
    Default value is 1000.
 
- TEAM_API_ENDPOINT (string): API endpoint to retrieve
    the team info. Default value is 
    https://vintagemonster.onefootball.com/api/teams/en/%team_id%.json

## Sample Output
Player name in sorted order; Player Age; Team names in Sorted order
```
>  go run main.go 
   Aaron Ramsey; 27; Arsenal
   Achraf Hakimi; 19; Real Madrid
   Adrien Rabiot; 23; France
   Adrià Ortolá; 24; Barcelona
   Ainsley Maitland-Niles; 20; Arsenal
   Aleix Vidal; 28; Barcelona
   Alex Iwobi; 22; Arsenal
   Alexandre Lacazette; 26; Arsenal
   Alexis Sanchez; 29; Manchester Utd
   Alphonse Areola; 25; France
   Alvaro Morata; 25; Chelsea
   Ander Herrera; 28; Manchester Utd
   Andreas Christensen; 22; Chelsea
   André Gomes; 24; Barcelona
   Andrés Iniesta; 34; Barcelona, Spain
   Angel Gomes; 17; Manchester Utd
   Anthony Martial; 22; France, Manchester Utd
   Antoine Griezmann; 27; France
   Antonio Rudiger; 25; Chelsea, Germany
   Antonio Valencia; 32; Manchester Utd
   Arjen Robben; 34; FC Bayern Munich
   Arturo Vidal; 30; FC Bayern Munich
   Ashley Young; 32; England, Manchester Utd
   Benjamin Pavard; 22; France
   Bernd Leno; 26; Germany
   Blaise Matuidi; 31; France
   Borja Mayoral; 21; Real Madrid
   Callum Hudson Odoi; 17; Chelsea
   Calum Chambers; 23; Arsenal
   Cameron Borthwick-Jackson; 21; Manchester Utd
   Carles Aleñá; 20; Barcelona
   Casemiro; 26; Real Madrid
   Cesc Fabregas; 31; Chelsea
   Charlie Gilmour; 19; Arsenal
   Charly Musonda; 21; Chelsea
   Chris Smalling; 28; Manchester Utd
   Christian Fruchtl; 18; FC Bayern Munich
   Corentin Tolisso; 23; FC Bayern Munich, France
   Coutinho; 25; Barcelona
   Cristiano Ronaldo; 33; Real Madrid
   Cristo González; 20; Real Madrid
   César Azpilicueta; 28; Chelsea, Spain
   Daley Blind; 28; Manchester Utd
   Dani Ceballos; 21; Real Madrid
   Daniel Carvajal; 26; Real Madrid, Spain
   Daniel Parejo; 29; Spain
   Danny Drinkwater; 28; Chelsea
   Danny Rose; 27; England
   Danny Welbeck; 27; Arsenal
   Danny Welbeck; 27; England
   David Alaba; 25; FC Bayern Munich
   David Costas; 23; Barcelona
   David De Gea; 27; Manchester Utd, Spain
   David Luiz; 31; Chelsea
   David Ospina; 29; Arsenal
   Davide Zappacosta; 25; Chelsea
   Dele Alli; 22; England
   Denis Suárez; 24; Barcelona
   Deyan Iliev; 23; Arsenal
   Diego Costa; 29; Spain
   Djibril Sidibe; 25; France
   Dujon Sterling; 18; Chelsea
   Eden Hazard; 27; Chelsea
   Eduardo; 35; Chelsea
   Edward Nketiah; 18; Arsenal
   Emerson; 23; Chelsea
   Eric Bailly; 24; Manchester Utd
   Eric Dier; 24; England
   Ethan Ampadu; 17; Chelsea
   Ethan Hamilton; 19; Manchester Utd
   Fabian Benko; 19; FC Bayern Munich
   Fabian Delph; 28; England
   Felix Gotze; 20; FC Bayern Munich
   Florian Thauvin; 25; France
   Franchu; 20; Real Madrid
   Franck Evina; 17; FC Bayern Munich
   Franck Ribery; 35; FC Bayern Munich
   Gareth Bale; 28; Real Madrid
   Gary Cahill; 32; England
   Gary Cahill; 32; Chelsea
   Gerard Piqué; 31; Barcelona, Spain
   Granit Xhaka; 25; Arsenal
   Harry Kane; 24; England
   Harry Maguire; 25; England
   Hector Bellerin; 23; Arsenal
   Henrikh Mkhitaryan; 29; Arsenal
   Hugo Lloris; 31; France
   Iago Aspas; 30; Spain
   Ilkay Gündogan; 27; Germany
   Inigo Ruiz de Galarreta; 24; Barcelona
   Isaiah Brown; 21; Chelsea
   Isco; 26; Real Madrid, Spain
   Ivan Rakitic; 30; Barcelona
   Jack Butland; 25; England
   Jack Wilshere; 26; Arsenal
   Jaime Seoane; 21; Real Madrid
   James Rodriguez; 26; FC Bayern Munich
   Jamie Vardy; 31; England
   Jasper Cillessen; 29; Barcelona
   Javi Martinez; 29; FC Bayern Munich
   Jesse Lingard; 25; Manchester Utd
   Jesse Lingard; 25; England
   Jesús Vallejo; 21; Real Madrid
   Joe Willock; 18; Arsenal
   Joel Pereira; 21; Manchester Utd
   John Stones; 23; England
   Jonas Hector; 27; Germany
   Jonathan Meier; 18; FC Bayern Munich
   Jonathan Tah; 22; Germany
   Jordan Henderson; 27; England
   Jordan Pickford; 24; England
   Jordi Alba; 29; Barcelona, Spain
   Jordi Osei-Tutu; 19; Arsenal
   Jorge Cuenca; 18; Barcelona
   Joshua Da Silva; 19; Arsenal
   Joshua Kimmich; 23; FC Bayern Munich, Germany
   José Arnáiz; 23; Barcelona
   José Reina; 35; Spain
   Juan Bernat; 25; FC Bayern Munich
   Juan Mata; 30; Manchester Utd
   Julian Brandt; 22; Germany
   Julian Draxler; 24; Germany
   Jérôme Boateng; 29; FC Bayern Munich, Germany
   Karim Benzema; 30; Real Madrid
   Kepa Arrizabalaga; 23; Spain
   Kevin Trapp; 27; Germany
   Keylor Navas; 31; Real Madrid
   Kieran Trippier; 27; England
   Kiko Casilla; 31; Real Madrid
   Kingsley Coman; 21; FC Bayern Munich
   Koke; 26; Spain
   Konstantinos Mavropanos; 20; Arsenal
   Kwasi Wriedt; 23; FC Bayern Munich
   Kyle Scott; 20; Chelsea
   Kyle Walker; 27; England
   Kylian Mbappé; 19; France
   Laurent Koscielny; 32; Arsenal, France
   Leo Weinkauf; 21; FC Bayern Munich
   Leon Goretzka; 23; Germany
   Leroy Sané; 22; Germany
   Lionel Messi; 30; Barcelona
   Luca Zidane; 20; Real Madrid
   Lucas Digne; 24; Barcelona, France
   Lucas Hernández; 22; France
   Lucas Vázquez; 26; Real Madrid, Spain
   Luis Suárez; 31; Barcelona
   Luismi Quezada; 22; Real Madrid
   Luka Modric; 32; Real Madrid
   Lukas Mai; 18; FC Bayern Munich
   Luke Shaw; 22; Manchester Utd
   Manu; 19; Real Madrid
   Manuel Neuer; 32; FC Bayern Munich, Germany
   Manuel Wintzheimer; 19; FC Bayern Munich
   Marc Cucurella; 19; Barcelona
   Marc-André ter Stegen; 26; Barcelona, Germany
   Marcelo; 30; Real Madrid
   Marcin Bulka; 18; Chelsea
   Marco Asensio; 22; Real Madrid, Spain
   Marco Reus; 28; Germany
   Marcos Alonso; 27; Chelsea, Spain
   Marcos Llorente; 23; Real Madrid
   Marcos Rojo; 28; Manchester Utd
   Marcus Rashford; 20; England, Manchester Utd
   Mario Gomez; 32; Germany
   Marouane Fellaini; 30; Manchester Utd
   Marvin Plattenhardt; 26; Germany
   Matej Delac; 25; Chelsea
   Mateo Kovacic; 24; Real Madrid
   Mats Hummels; 29; FC Bayern Munich, Germany
   Matt Macey; 23; Arsenal
   Matteo Darmian; 28; Manchester Utd
   Matthias Ginter; 24; Germany
   Meritan Shabani; 19; FC Bayern Munich
   Mesut Ozil; 29; Arsenal, Germany
   Michael Carrick; 36; Manchester Utd
   Milos Pantovic; 21; FC Bayern Munich
   Moha Ramos; 18; Real Madrid
   Mohamed Elneny; 25; Arsenal
   N'Golo Kante; 27; Chelsea, France
   Nacho; 28; Real Madrid, Spain
   Nacho Monreal; 32; Arsenal
   Nemanja Matic; 29; Manchester Utd
   Nick Pope; 26; England
   Niklas Dorsch; 20; FC Bayern Munich
   Niklas Süle; 22; FC Bayern Munich, Germany
   Nils Petersen; 29; Germany
   Nélson Semedo; 24; Barcelona
   Olivier Giroud; 31; Chelsea, France
   Oriol Busquets; 19; Barcelona
   Ousmane Dembélé; 21; Barcelona, France
   Paco Alcácer; 24; Barcelona
   Paul Pogba; 25; France, Manchester Utd
   Paulinho; 29; Barcelona
   Pedro; 30; Chelsea
   Per Mertesacker; 33; Arsenal
   Petr Cech; 36; Arsenal
   Phil Jones; 26; Manchester Utd
   Phil Jones; 26; England
   Pierre-Emerick Aubameyang; 28; Arsenal
   Presnel Kimpembe; 22; France
   Rafinha; 32; FC Bayern Munich
   Raheem Sterling; 23; England
   Raphael Varane; 25; France, Real Madrid
   Reiss Nelson; 18; Arsenal
   Rob Holding; 22; Arsenal
   Robert Lewandowski; 29; FC Bayern Munich
   Rodri; 21; Spain
   Rodrigo Moreno; 27; Spain
   Romelu Lukaku; 25; Manchester Utd
   Ron Thorben Hofmann; 19; FC Bayern Munich
   Ross Barkley; 24; Chelsea
   Ruben Loftus-Cheek; 22; England
   Sami Khedira; 31; Germany
   Samuel Umtiti; 24; Barcelona, France
   Sandro Wagner; 30; FC Bayern Munich
   Santi Cazorla; 33; Arsenal
   Saúl Ñíguez; 23; Spain
   Scott McTominay; 21; Manchester Utd
   Sead Kolasinac; 24; Arsenal
   Sebastian Rudy; 28; FC Bayern Munich, Germany
   Sergi Roberto; 26; Barcelona
   Sergio Busquets; 29; Barcelona
   Sergio Ramos; 32; Real Madrid, Spain
   Sergio Romero; 31; Manchester Utd
   Shkodran Mustafi; 26; Arsenal
   Steve Mandanda; 33; France
   Sven Ulreich; 29; FC Bayern Munich
   Theo Hernández; 20; Real Madrid
   Thiago Alcántara; 27; FC Bayern Munich, Spain
   Thibaut Courtois; 26; Chelsea
   Thomas Lemar; 22; France
   Thomas Müller; 28; FC Bayern Munich, Germany
   Thomas Vermaelen; 32; Barcelona
   Tiemoue Bakayoko; 23; Chelsea
   Timo Werner; 22; Germany
   Tom Starke; 37; FC Bayern Munich
   Toni Kroos; 28; Germany, Real Madrid
   Trent Alexander-Arnold; 19; England
   Trevor Chalobah; 18; Chelsea
   Victor Lindelof; 23; Manchester Utd
   Victor Moses; 27; Chelsea
   Vlad Dragomir; 19; Arsenal
   Wallace Oliveira; 24; Chelsea
   Willian; 29; Chelsea
   Willy Caballero; 36; Chelsea
   Wissam Ben Yedder; 27; France
   Yerry Mina; 23; Barcelona
   Álvaro Odriozola; 22; Spain
   Álvaro Tejero; 21; Real Madrid
   Óscar Rodríguez; 19; Real Madrid
```