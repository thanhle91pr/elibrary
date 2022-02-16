# Elibrary service
- The eLibrary service supports to manage the info of: 
1. Books (id, name, description)
2. Songs (id, name, description)
3. Labels (id, name, description)
### Features:
1. List of Labels is seeded in DB with 20 labels+. User cannot change the labels table.
2. User can create new Book, new Song.
3. User can tag a label to any books, 1 book can only be tagged to 1 label, existing tagged label of the book need to be replaced.
4. User can tag a label to any songs, 1 song can only be tagged to 1 label, existing tagged label of the song need to be replaced.
5. User can tag a label to any combo book+song, 1 combo book+song can only be tagged to 1 label, existing tagged label of the combo need to be replaced (shouldn’t replace the label of individual book or individual song)
6. A label can be reused multiple times.
7. A search page: User can select a book, a song, or both to search for the label. There should be only 1 label or no label in the result. (It’s a label of individual book or individual song, or the label of a combo).
8. List 10 most popular labels that are used to tag most.
9. Lower case `name` is unique for each tables.
10. `name` has max 100 characters, `description` is optional and has max 200 characters
11. No authentication required, anyone can use the webapp.
12. No need the update or delete function for books, songs.
